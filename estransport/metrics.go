package estransport

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

// Measurable defines the interface for transports supporting metrics.
//
type Measurable interface {
	Metrics() (Metrics, error)
}

// Metrics represents the transport metrics.
//
type Metrics struct {
	Requests int `json:"requests"`
	Failures int `json:"failures"`

	Live []connectionMetric `json:"live,omitempty"`
	Dead []connectionMetric `json:"dead,omitempty"`
}

// TODO(karmi): Fix empty value for DeadSince
type connectionMetric struct {
	URL         string        `json:"url"`
	Failures    int           `json:"failures,omitempty"`
	DeadSince   time.Time     `json:"dead_since,omitempty"`
	ResurrectIn time.Duration `json:"resurrect_in,omitempty"`
}

type metrics struct {
	sync.RWMutex

	requests int
	failures int
	live     []*Connection
	dead     []*Connection
}

// Metrics returns the transport metrics.
//
func (c *Client) Metrics() (Metrics, error) {
	c.metrics.RLock()
	defer c.metrics.RUnlock()

	if lockable, ok := c.pool.(sync.Locker); ok {
		lockable.Lock()
		defer lockable.Unlock()
	}

	m := Metrics{
		Requests: c.metrics.requests,
		Failures: c.metrics.failures,
	}

	// FIXME(karmi): Type assertion to interface
	if pool, ok := c.pool.(*roundRobinConnectionPool); ok {
		for _, c := range pool.live {
			c.Lock()
			m.Live = append(m.Live, connectionMetric{URL: c.URL.String()})
			c.Unlock()
		}

		for _, c := range pool.dead {
			c.Lock()
			m.Dead = append(m.Dead, connectionMetric{
				URL:       c.URL.String(),
				Failures:  c.Failures,
				DeadSince: c.DeadSince,
				// FIXME(karmi): Take factor into account
				ResurrectIn: c.DeadSince.Add(defaultResurrectTimeoutInitial).Sub(time.Now().UTC()).Truncate(time.Second),
			})
			c.Unlock()
		}
	}

	return m, nil
}

// String returns the connection information as a string.
//
func (m connectionMetric) String() string {
	var b strings.Builder
	b.WriteString("{")
	b.WriteString(m.URL)
	if m.Failures > 0 {
		fmt.Fprintf(&b, " failures=%d", m.Failures)
	}
	if m.ResurrectIn > time.Duration(0) {
		fmt.Fprintf(&b, " resurrect_in=%s", m.ResurrectIn)
	}
	if !m.DeadSince.IsZero() {
		fmt.Fprintf(&b, " dead_since=%s", m.DeadSince.Local().Format(time.Stamp))
	}
	b.WriteString("}")
	return b.String()
}

// For expvar, do something like this:
//
// expvar.Publish("goelasticsearch", expvar.Func(func() interface{} {
// 		m, _ := es.Metrics()
// 		return m
// 	}))
