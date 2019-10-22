// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package estransport

import (
	"errors"
	"fmt"
	"math"
	"net/url"
	"sort"
	"sync"
	"time"
)

var (
	defaultResurrectTimeoutInitial      = 60 * time.Second
	defaultResurrectTimeoutFactorCutoff = 5
)

// ConnectionPool defines the interface for the connection pool.
//
type ConnectionPool interface {
	Next() (*Connection, error)
	Remove(*Connection) error
}

// Connection represents a connection to a node.
//
type Connection struct {
	sync.Mutex

	URL       *url.URL
	Dead      bool
	DeadSince time.Time
	Failures  int
}

type singleConnectionPool struct {
	connection *Connection

	metrics       *metrics
	enableMetrics bool

	debugLogger DebuggingLogger
}

type roundRobinConnectionPool struct {
	sync.Mutex

	curr int           // Index of the current connection
	live []*Connection // List of live connections
	dead []*Connection // List of dead connections
	orig []*url.URL    // List of original URLs, passed in during initialization

	metrics       *metrics
	enableMetrics bool

	debugLogger DebuggingLogger
}

func newSingleConnectionPool(u *url.URL) *singleConnectionPool {
	cp := singleConnectionPool{connection: &Connection{URL: u}}

	return &cp
}

func newRoundRobinConnectionPool(u ...*url.URL) *roundRobinConnectionPool {
	var conns []*Connection
	for _, url := range u {
		conns = append(conns, &Connection{URL: url})
	}

	cp := roundRobinConnectionPool{live: conns, orig: u, curr: -1}

	if cp.enableMetrics {
		cp.metrics.Lock()
		cp.metrics.live = cp.live
		cp.metrics.dead = cp.dead
		cp.metrics.Unlock()
	}

	return &cp
}

// Next returns the connection from pool.
//
func (cp *singleConnectionPool) Next() (*Connection, error) {
	return cp.connection, nil
}

// Remove is a no-op for single connection pool.
//
func (cp *singleConnectionPool) Remove(c *Connection) error {
	return nil
}

// Next returns a connection from pool, or an error.
//
func (cp *roundRobinConnectionPool) Next() (*Connection, error) {
	var c *Connection

	cp.Lock()
	defer cp.Unlock()

	// Return next live connection
	if len(cp.live) > 0 {
		cp.curr = (cp.curr + 1) % len(cp.live)
		return cp.live[cp.curr], nil

	} else if len(cp.dead) > 0 {
		// Try to resurrect a dead connection if no live connections are available
		if cp.debugLogger != nil {
			cp.debugLogger.Log("Resurrecting connection ")
		}
		c, cp.dead = cp.dead[len(cp.dead)-1], cp.dead[:len(cp.dead)-1] // Pop item
		c.Lock()
		if cp.debugLogger != nil {
			cp.debugLogger.Log(c.URL.String() + "\n")
		}
		c.markAsLive()
		cp.live = append(cp.live, c)
		c.Unlock()

		return c, nil
	}

	return nil, errors.New("no connection available")
}

// Remove removes a connection from the pool.
//
func (cp *roundRobinConnectionPool) Remove(c *Connection) error {
	cp.Lock()
	defer cp.Unlock()

	c.Lock()

	if c.Dead {
		if cp.debugLogger != nil {
			cp.debugLogger.Logf("Already removed %s\n", c.URL)
		}
		c.Unlock()
		return nil
	}

	if cp.debugLogger != nil {
		cp.debugLogger.Logf("Removing %s...\n", c.URL)
	}
	c.markAsDead()
	cp.scheduleResurrect(c)
	c.Unlock()

	// Push item to dead list and sort slice by number of failures
	cp.dead = append(cp.dead, c)
	sort.Slice(cp.dead, func(i, j int) bool {
		c1 := cp.dead[i]
		c2 := cp.dead[j]
		c1.Lock()
		c2.Lock()
		defer c1.Unlock()
		defer c2.Unlock()

		res := c1.Failures > c2.Failures
		return res
	})

	// Check if connection exists in the list. Return nil if it doesn't,
	// because another goroutine might have already removed it.
	index := -1
	for i, conn := range cp.live {
		if conn == c {
			index = i
		}
	}
	if index < 0 {
		return nil
	}

	// Remove item; https://github.com/golang/go/wiki/SliceTricks
	copy(cp.live[index:], cp.live[index+1:])
	cp.live = cp.live[:len(cp.live)-1]

	if cp.enableMetrics {
		cp.metrics.Lock()
		cp.metrics.dead = cp.dead
		cp.metrics.live = cp.live
		cp.metrics.Unlock()
	}

	return nil
}

// Resurrect removes the connection from the dead list and adds it to the pool.
//
// TODO(karmi): Add a pluggable strategy as argument, eg. "optimistic", "ping".
//
func (cp *roundRobinConnectionPool) Resurrect(c *Connection) error {
	cp.Lock()
	defer cp.Unlock()

	c.Lock()
	defer c.Unlock()

	if !c.Dead {
		if cp.debugLogger != nil {
			cp.debugLogger.Logf("Already resurrected %s\n", c.URL)
		}
		return nil
	}

	if cp.debugLogger != nil {
		cp.debugLogger.Logf("Resurrecting %s\n", c.URL)
	}

	c.markAsLive()

	cp.live = append(cp.live, c)
	index := -1
	for i, conn := range cp.dead {
		if conn == c {
			index = i
		}
	}
	if index >= 0 {
		// Remove item; https://github.com/golang/go/wiki/SliceTricks
		copy(cp.dead[index:], cp.dead[index+1:])
		cp.dead = cp.dead[:len(cp.dead)-1]
	}

	if cp.enableMetrics {
		cp.metrics.Lock()
		cp.metrics.dead = cp.dead
		cp.metrics.live = cp.live
		cp.metrics.Unlock()
	}

	return nil
}

// scheduleResurrect schedules the connection to be resurrected.
//
func (cp *roundRobinConnectionPool) scheduleResurrect(c *Connection) {
	factor := math.Min(float64(c.Failures-1), float64(defaultResurrectTimeoutFactorCutoff))
	timeout := time.Duration(defaultResurrectTimeoutInitial.Seconds() * math.Exp2(factor) * float64(time.Second))
	if cp.debugLogger != nil {
		cp.debugLogger.Logf("Resurrect %s (failures=%d, factor=%1.1f, timeout=%s) in %s\n", c.URL, c.Failures, factor, timeout, c.DeadSince.Add(timeout).Sub(time.Now().UTC()).Truncate(time.Second))
	}

	time.AfterFunc(timeout, func() { cp.Resurrect(c) })
}

// markAsDead marks the connection as dead.
//
func (c *Connection) markAsDead() {
	c.Dead = true
	c.DeadSince = time.Now().UTC()
	c.Failures++
}

// markAsLive marks the connection as alive.
//
func (c *Connection) markAsLive() {
	c.Dead = false
}

// markAsHealthy marks the connection as healthy.
//
func (c *Connection) markAsHealthy() {
	c.Dead = false
	c.DeadSince = time.Time{}
	c.Failures = 0
}

func (c *Connection) String() string {
	c.Lock()
	defer c.Unlock()
	return fmt.Sprintf("<%s> dead=%v failures=%d", c.URL, c.Dead, c.Failures)
}
