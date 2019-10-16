package estransport

import (
	"sync"
)

// Metrics represents the transport metrics.
//
type Metrics struct {
	sync.RWMutex

	NumRequests int `json:"num_requests"`
	NumFailures int `json:"num_failures"`

	// TODO(karmi): Serialize connection URL as plain string
	Live []*Connection `json:"live"`
	Dead []*Connection `json:"dead"`
}

// Metrics returns the transport metrics.
//
func (c *Client) Metrics() *Metrics {
	c.metrics.RLock()
	c.metrics.RUnlock()
	return c.metrics
}

// For expvar, do something like this:
//
// expvar.Publish("goelasticsearch", expvar.Func(func() interface{} {
// 		m, _ := es.Metrics()
// 		return m
// 	}))
