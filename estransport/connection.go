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

// Selector defines the interface for selecting connections from the pool.
//
type Selector interface {
	Select([]*Connection) (*Connection, error)
}

// ConnectionPool defines the interface for the connection pool.
//
type ConnectionPool interface {
	Next() (*Connection, error)  // Next returns the next available connection.
	OnSuccess(*Connection) error // OnSuccess reports that the connection behaved successfully.
	OnFailure(*Connection) error // OnFailure reports that the connection failed.
	Len() int                    // Len returns the number of available connections.
}

// Connection represents a connection to a node.
//
type Connection struct {
	sync.Mutex

	URL       *url.URL
	IsDead    bool
	DeadSince time.Time
	Failures  int
}

type singleConnectionPool struct {
	connection *Connection

	metrics     *metrics
	debugLogger DebuggingLogger
}

type roundRobinConnectionPool struct {
	sync.Mutex

	live     []*Connection // List of live connections
	dead     []*Connection // List of dead connections
	orig     []*Connection // List of the original connections, passed in during initialization
	selector Selector

	metrics     *metrics
	debugLogger DebuggingLogger
}

type roundRobinSelector struct {
	sync.Mutex

	curr int // Index of the current connection
}

// NewDefaultConnectionPool creates the default connection pool.
//
func NewDefaultConnectionPool(conns []*Connection, selector Selector) (ConnectionPool, error) {
	if len(conns) == 1 {
		return &singleConnectionPool{connection: conns[0]}, nil
	}
	if selector == nil {
		selector = &roundRobinSelector{curr: -1}
	}
	return &roundRobinConnectionPool{live: conns, selector: selector}, nil
}

// Next returns the connection from pool.
//
func (cp *singleConnectionPool) Next() (*Connection, error) {
	return cp.connection, nil
}

// OnSuccess is a no-op for single connection pool.
func (cp *singleConnectionPool) OnSuccess(c *Connection) error { return nil }

// OnFailure is a no-op for single connection pool.
func (cp *singleConnectionPool) OnFailure(c *Connection) error { return nil }

// Len returns the number of available connections.
func (cp *singleConnectionPool) Len() int { return 1 }

// Next returns a connection from pool, or an error.
//
func (cp *roundRobinConnectionPool) Next() (*Connection, error) {
	cp.Lock()
	defer cp.Unlock()

	// Return next live connection
	if len(cp.live) > 0 {
		return cp.selector.Select(cp.live)
	} else if len(cp.dead) > 0 {
		// No live connection is available, resurrect one of the dead ones.
		c := cp.dead[len(cp.dead)-1]
		cp.dead = cp.dead[:len(cp.dead)-1]
		c.Lock()
		defer c.Unlock()
		cp.resurrect(c, false)
		return c, nil
	}
	return nil, errors.New("no connection available")
}

// OnFailure marks the connection as failed.
//
func (cp *roundRobinConnectionPool) OnFailure(c *Connection) error {
	cp.Lock()
	defer cp.Unlock()

	c.Lock()

	if c.IsDead {
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

	// Check if connection exists in the list, return error if not.
	index := -1
	for i, conn := range cp.live {
		if conn == c {
			index = i
		}
	}
	if index < 0 {
		return errors.New("connection not in live list")
	}

	// Remove item; https://github.com/golang/go/wiki/SliceTricks
	copy(cp.live[index:], cp.live[index+1:])
	cp.live = cp.live[:len(cp.live)-1]

	if cp.metrics != nil {
		cp.metrics.Lock()
		cp.metrics.dead = cp.dead
		cp.metrics.live = cp.live
		cp.metrics.Unlock()
	}

	return nil
}

// OnSuccess marks the connection as successful.
//
func (cp *roundRobinConnectionPool) OnSuccess(c *Connection) error {
	c.Lock()
	defer c.Unlock()

	// Short-circuit for live connection
	if !c.IsDead {
		return nil
	}

	cp.Lock()
	defer cp.Unlock()
	return cp.resurrect(c, true)
}

func (cp *roundRobinConnectionPool) Len() int {
	cp.Lock()
	defer cp.Unlock()

	return len(cp.live)
}

// resurrect adds the connection to the list of available connections.
// When removeDead is true, it also removes it from the dead list.
// The calling code is responsible for locking.
//
func (cp *roundRobinConnectionPool) resurrect(c *Connection, removeDead bool) error {
	if cp.debugLogger != nil {
		cp.debugLogger.Logf("Resurrecting %s\n", c.URL)
	}

	c.markAsLive()

	cp.live = append(cp.live, c)
	if removeDead {
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
	}

	if cp.metrics != nil {
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

	time.AfterFunc(timeout, func() {
		cp.Lock()
		defer cp.Unlock()

		c.Lock()
		defer c.Unlock()

		if !c.IsDead {
			if cp.debugLogger != nil {
				cp.debugLogger.Logf("Already resurrected %s\n", c.URL)
			}
			return
		}

		cp.resurrect(c, true)
	})
}

// Select returns the connection in a round-robin fashion.
//
func (s *roundRobinSelector) Select(conns []*Connection) (*Connection, error) {
	s.Lock()
	defer s.Unlock()

	s.curr = (s.curr + 1) % len(conns)
	return conns[s.curr], nil
}

// markAsDead marks the connection as dead.
//
func (c *Connection) markAsDead() {
	c.IsDead = true
	c.DeadSince = time.Now().UTC()
	c.Failures++
}

// markAsLive marks the connection as alive.
//
func (c *Connection) markAsLive() {
	c.IsDead = false
}

// markAsHealthy marks the connection as healthy.
//
func (c *Connection) markAsHealthy() {
	c.IsDead = false
	c.DeadSince = time.Time{}
	c.Failures = 0
}

// String returns a readable connection representation.
//
func (c *Connection) String() string {
	c.Lock()
	defer c.Unlock()
	return fmt.Sprintf("<%s> dead=%v failures=%d", c.URL, c.IsDead, c.Failures)
}
