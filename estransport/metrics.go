// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

package estransport

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

// For expvar, do something like this:
//
// expvar.Publish("go-elasticsearch", expvar.Func(func() interface{} {
// 		m, _ := es.Metrics()
// 		return m
// 	}))

// Measurable defines the interface for transports supporting metrics.
//
type Measurable interface {
	Metrics() (Metrics, error)
}

// Metrics represents the transport metrics.
//
type Metrics struct {
	Requests  int         `json:"requests"`
	Failures  int         `json:"failures"`
	Responses map[int]int `json:"responses"`

	Live []connectionMetric `json:"live,omitempty"`
	Dead []connectionMetric `json:"dead,omitempty"`
}

// connectionMetric represents metric information for a connection.
//
type connectionMetric struct {
	URL         string        `json:"url"`
	Failures    int           `json:"failures,omitempty"`
	DeadSince   nullableTime  `json:"dead_since,omitempty"`
	ResurrectIn time.Duration `json:"resurrect_in,omitempty"`
}

// nullableTime allows to return time zero value as nil.
//
type nullableTime struct{ time.Time }

// metrics represents the inner state of metrics.
//
type metrics struct {
	sync.RWMutex

	requests  int
	failures  int
	responses map[int]int

	live []*Connection
	dead []*Connection
}

// Metrics returns the transport metrics.
//
func (c *Client) Metrics() (Metrics, error) {
	if !c.enableMetrics {
		return Metrics{}, errors.New("transport metrics not enabled")
	}
	c.metrics.RLock()
	defer c.metrics.RUnlock()

	if lockable, ok := c.pool.(sync.Locker); ok {
		lockable.Lock()
		defer lockable.Unlock()
	}

	m := Metrics{
		Requests:  c.metrics.requests,
		Failures:  c.metrics.failures,
		Responses: c.metrics.responses,
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
				DeadSince: nullableTime{c.DeadSince},
				// FIXME(karmi): Take factor into account
				ResurrectIn: c.DeadSince.Add(defaultResurrectTimeoutInitial).Sub(time.Now().UTC()).Truncate(time.Second),
			})
			c.Unlock()
		}
	}

	return m, nil
}

// String returns the metrics as a string.
//
func (m Metrics) String() string {
	var (
		i int
		b strings.Builder
	)
	b.WriteString("{")

	b.WriteString("Requests:")
	b.WriteString(strconv.Itoa(m.Requests))

	b.WriteString(" Failures:")
	b.WriteString(strconv.Itoa(m.Failures))

	if len(m.Responses) > 0 {
		b.WriteString(" Responses: ")
		b.WriteString("[")

		for code, num := range m.Responses {
			b.WriteString(strconv.Itoa(code))
			b.WriteString(":")
			b.WriteString(strconv.Itoa(num))
			if i+1 < len(m.Responses) {
				b.WriteString(", ")
			}
			i++
		}
		b.WriteString("]")
	}

	b.WriteString(" Live: [")
	for i, c := range m.Live {
		b.WriteString(c.String())
		if i+1 < len(m.Live) {
			b.WriteString(", ")
		}
		i++
	}
	b.WriteString("]")

	b.WriteString(" Dead: [")
	for i, c := range m.Dead {
		b.WriteString(c.String())
		if i+1 < len(m.Dead) {
			b.WriteString(", ")
		}
		i++
	}
	b.WriteString("]")

	b.WriteString("}")
	return b.String()
}

// String returns the connection information as a string.
//
func (cm connectionMetric) String() string {
	var b strings.Builder
	b.WriteString("{")
	b.WriteString(cm.URL)
	if cm.Failures > 0 {
		fmt.Fprintf(&b, " failures=%d", cm.Failures)
	}
	if !cm.DeadSince.IsZero() {
		fmt.Fprintf(&b, " dead_since=%s", cm.DeadSince.Local().Format(time.Stamp))
	}
	if cm.ResurrectIn > time.Duration(0) {
		fmt.Fprintf(&b, " resurrect_in=%s", cm.ResurrectIn)
	}
	b.WriteString("}")
	return b.String()
}

// MarshallJSON encodes zero value of time as nil.
//
// NOTE: isEmptyValue() doesn't handle time values.
//       https://golang.org/src/encoding/json/encode.go?s=10804:10846#L318
//
func (t nullableTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`null`), nil
	}
	return json.Marshal(t.Time)
}
