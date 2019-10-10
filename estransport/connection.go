package estransport

import (
	"container/ring"
	"errors"
	"net/url"
	"sync"
)

// ConnectionPool defines the interface for the connection pool.
//
type ConnectionPool interface {
	GetConnection() (*Connection, error)
}

// Connection represents a connection to a node.
//
type Connection struct {
	URL *url.URL

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
	ring *ring.Ring
}

// newRoundRobinConnectionPool creates a new roundRobinConnectionPool.
//
func newRoundRobinConnectionPool(connections ...*Connection) *roundRobinConnectionPool {
	cp := roundRobinConnectionPool{}

	cp.ring = ring.New(len(connections))
	for _, c := range connections {
		cp.ring.Value = c
		cp.ring = cp.ring.Next()
	}

	return &cp
}

// GetConnection returns a connection from pool, or an error.
//
func (cp *roundRobinConnectionPool) GetConnection() (*Connection, error) {
	cp.Lock()
	defer cp.Unlock()

	if cp.ring.Len() < 1 {
		return nil, errors.New("No connection available")
	}

	v := cp.ring.Value
	if c, ok := v.(*Connection); !ok || c == nil {
		return nil, errors.New("No connection available")
	}

	cp.ring = cp.ring.Next()
	return v.(*Connection), nil
}
