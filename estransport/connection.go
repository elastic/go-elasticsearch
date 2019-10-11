package estransport

import (
	"errors"
	"fmt"
	"net/url"
	"sort"
	"sync"
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

	URL      *url.URL
	Dead     bool
	Failures int

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

	list []*Connection
	curr int

	dead []*Connection
}

// newRoundRobinConnectionPool creates a new roundRobinConnectionPool.
//
func newRoundRobinConnectionPool(connections ...*Connection) *roundRobinConnectionPool {
	cp := roundRobinConnectionPool{
		list: connections,
	}

	return &cp
}

// Next returns a connection from pool, or an error.
//
func (cp *roundRobinConnectionPool) Next() (*Connection, error) {
	var c *Connection

	cp.Lock()
	defer cp.Unlock()

	// fmt.Println("Next()", cp.list)

	// Try to resurrect a dead connection if no healthy connections are available
	//
	if len(cp.list) < 1 {
		if len(cp.dead) > 0 {
			fmt.Println("Next()", cp.dead)
			fmt.Printf("Resurrecting connection...")
			c, cp.dead = cp.dead[len(cp.dead)-1], cp.dead[:len(cp.dead)-1] // Pop item
			fmt.Println(c)
			c.Dead = false
			cp.list = append(cp.list, c)
			return c, nil
		}
		return nil, errors.New("no connection available")
	}

	if cp.curr >= len(cp.list) {
		cp.curr = len(cp.list) - 1
	}

	if cp.curr < 0 {
		return nil, errors.New("no connection available")
	}

	c = cp.list[cp.curr]
	cp.curr = (cp.curr + 1) % len(cp.list)

	return c, nil
}

// Remove removes a connection from the pool.
//
func (cp *roundRobinConnectionPool) Remove(c *Connection) error {
	c.Lock()
	fmt.Printf("Removing %s...\n", c.URL)
	c.Dead = true
	c.Failures++
	c.Unlock()

	cp.Lock()
	defer cp.Unlock()

	// Push item to dead list and sort slice by number of failures
	cp.dead = append(cp.dead, c)
	sort.Slice(cp.dead, func(i, j int) bool {
		c1 := cp.dead[i]
		c2 := cp.dead[j]
		c1.Lock()
		c2.Lock()

		res := c1.Failures > c2.Failures
		c1.Unlock()
		c2.Unlock()
		return res
	})

	// Check if connection exists in the list. Return nil if it doesn't, because another
	// goroutine might have already removed it.
	index := -1
	for i, conn := range cp.list {
		if conn == c {
			index = i
		}
	}
	if index < 0 {
		return nil
	}

	// Remove item; https://github.com/golang/go/wiki/SliceTricks
	copy(cp.list[index:], cp.list[index+1:])
	cp.list[len(cp.list)-1] = nil
	cp.list = cp.list[:len(cp.list)-1]

	// fmt.Println("Remove()", cp.list)

	return nil
}

func (c *Connection) String() string {
	c.Lock()
	defer c.Unlock()
	return fmt.Sprintf("<%s> dead=%v failures=%d", c.URL, c.Dead, c.Failures)
}
