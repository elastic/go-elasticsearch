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

	URL       *url.URL  `json:"url"`
	Dead      bool      `json:"dead"`
	DeadSince time.Time `json:"dead_since"`
	Failures  int       `json:"failures"`

	// ID         string
	// Name       string
	// Version    string
	// Roles      []string
	// Attributes map[string]interface{}
}

type singleConnectionPool struct {
	connection *Connection
}

type roundRobinConnectionPool struct {
	sync.Mutex

	curr int           // Index of the current connection
	live []*Connection // List of live connections
	dead []*Connection // List of dead connections
	orig []*url.URL    // List of original URLs, passed in during initialization
}

// newSingleConnectionPool creates a new SingleConnectionPool.
//
func newSingleConnectionPool(u *url.URL) *singleConnectionPool {
	return &singleConnectionPool{connection: &Connection{URL: u}}
}

// newRoundRobinConnectionPool creates a new roundRobinConnectionPool.
//
func newRoundRobinConnectionPool(u ...*url.URL) *roundRobinConnectionPool {
	var conns []*Connection
	for _, url := range u {
		conns = append(conns, &Connection{URL: url})
	}

	cp := roundRobinConnectionPool{live: conns, orig: u}

	if metrics != nil {
		metrics.Lock()
		metrics.Pool = cp.live
		metrics.Dead = cp.dead
		metrics.Unlock()
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

	// fmt.Println("Next()", cp.live)

	// Try to resurrect a dead connection if no healthy connections are available
	//
	if len(cp.live) < 1 {
		if len(cp.dead) > 0 {
			fmt.Println("Next(), Dead:", cp.dead)
			fmt.Printf("Resurrecting connection ")
			c, cp.dead = cp.dead[len(cp.dead)-1], cp.dead[:len(cp.dead)-1] // Pop item
			c.Lock()
			fmt.Println(c.URL)
			c.markAsLive()
			cp.live = append(cp.live, c)
			c.Unlock()

			if metrics != nil {
				metrics.Lock()
				metrics.Pool = cp.live
				metrics.Dead = cp.dead
				metrics.Unlock()
			}

			return c, nil
		}
		return nil, errors.New("no connection available")
	}

	if cp.curr >= len(cp.live) {
		cp.curr = len(cp.live) - 1
	}

	if cp.curr < 0 {
		return nil, errors.New("no connection available")
	}

	c = cp.live[cp.curr]
	cp.curr = (cp.curr + 1) % len(cp.live)

	return c, nil
}

// Remove removes a connection from the pool.
//
func (cp *roundRobinConnectionPool) Remove(c *Connection) error {
	c.Lock()

	if c.Dead {
		fmt.Printf("Already removed %s\n", c.URL)
		c.Unlock()
		return nil
	}

	fmt.Printf("Removing %s...\n", c.URL)
	c.markAsDead()
	c.scheduleResurrect(cp)
	c.Unlock()

	cp.Lock()
	defer cp.Unlock()

	// Push item to dead list and sort slice by number of failures
	cp.dead = append(cp.dead, c)
	sort.Slice(cp.dead, func(i, j int) bool {
		c1 := cp.dead[i]
		c2 := cp.dead[j]
		c1.Lock()
		if c1 != c2 {
			c2.Lock()
		}
		defer c1.Unlock()
		if c1 != c2 {
			defer c2.Unlock()
		}

		res := c1.Failures > c2.Failures
		return res
	})

	if metrics != nil {
		metrics.Lock()
		metrics.Dead = cp.dead
		metrics.Unlock()
	}

	// Check if connection exists in the list. Return nil if it doesn't, because another
	// goroutine might have already removed it.
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

	if metrics != nil {
		metrics.Lock()
		metrics.Pool = cp.live
		metrics.Unlock()
	}

	return nil
}

// Resurrect removes the connection from the dead list and adds it to the pool.
//
// TODO(karmi): Add a pluggable strategy as argument, eg. "optimistic", "ping".
//
func (c *Connection) Resurrect(cp *roundRobinConnectionPool) error {
	cp.Lock()
	defer cp.Unlock()

	c.Lock()
	defer c.Unlock()

	if !c.Dead {
		fmt.Printf("Already resurrected %s\n", c.URL)
		return nil
	}

	fmt.Printf("Resurrecting %s, timeout passed\n", c.URL)

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

	return nil
}

// scheduleResurrect schedules the connection to be resurrected.
//
func (c *Connection) scheduleResurrect(cp *roundRobinConnectionPool) {
	factor := func(a, b int) float64 {
		if a > b {
			return float64(b)
		}
		return float64(a)
	}(c.Failures-1, defaultResurrectTimeoutFactorCutoff)

	timeout := time.Duration(defaultResurrectTimeoutInitial.Seconds() * math.Exp2(factor) * float64(time.Second))
	fmt.Printf("Resurrect %s (failures=%d, factor=%1.1f, timeout=%s) in %s\n", c.URL, c.Failures, factor, timeout, c.DeadSince.Add(timeout).Sub(time.Now().UTC()).Truncate(time.Second))

	time.AfterFunc(timeout, func() { c.Resurrect(cp) })
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
