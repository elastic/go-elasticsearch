// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package estransport

import (
	"fmt"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestMetrics(t *testing.T) {
	t.Run("Metrics()", func(t *testing.T) {
		tp := New(
			Config{
				URLs: []*url.URL{
					{Scheme: "http", Host: "foo1"},
					{Scheme: "http", Host: "foo2"},
					{Scheme: "http", Host: "foo3"},
				},
				DisableRetry:  true,
				EnableMetrics: true,
			},
		)

		tp.metrics.requests = 3
		tp.metrics.failures = 4
		tp.metrics.responses[200] = 1
		tp.metrics.responses[404] = 2

		req, _ := http.NewRequest("HEAD", "/", nil)
		tp.Perform(req)

		m, err := tp.Metrics()
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		fmt.Println(m)

		if m.Requests != 4 {
			t.Errorf("Unexpected output, want=4, got=%d", m.Requests)
		}
		if m.Failures != 5 {
			t.Errorf("Unexpected output, want=5, got=%d", m.Failures)
		}
		if len(m.Responses) != 2 {
			t.Errorf("Unexpected output: %+v", m.Responses)
		}
		if len(m.Live) != 2 {
			t.Errorf("Unexpected output: %+v", m.Live)
		}
		if len(m.Dead) != 1 {
			t.Errorf("Unexpected output: %+v", m.Dead)
		}
	})

	t.Run("Metrics() when not enabled", func(t *testing.T) {
		tp := New(Config{})

		_, err := tp.Metrics()
		if err == nil {
			t.Fatalf("Expected error, got: %v", err)
		}
	})

	t.Run("String()", func(t *testing.T) {
		var m connectionMetric

		m = connectionMetric{
			URL:         "http://foo1",
			Failures:    0,
			DeadSince:   nullableTime{time.Time{}},
			ResurrectIn: time.Duration(0),
		}

		if m.String() != "{http://foo1}" {
			t.Errorf("Unexpected output: %s", m)
		}

		tt, _ := time.Parse(time.RFC3339, "2010-11-11T11:00:00Z")
		m = connectionMetric{
			URL:         "http://foo2",
			Failures:    123,
			DeadSince:   nullableTime{tt},
			ResurrectIn: time.Duration(100),
		}

		match, err := regexp.MatchString(
			`{http://foo2 failures=123 dead_since=Nov 11 \d+:00:00 resurrect_in=100ns}`,
			m.String(),
		)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if !match {
			t.Errorf("Unexpected output: %s", m)
		}
	})

	t.Run("nullableTime.MarshalJSON()", func(t *testing.T) {
		var (
			j   []byte
			err error
		)

		j, err = nullableTime{time.Now()}.MarshalJSON()
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if !strings.HasPrefix(string(j), `"`+strconv.Itoa(time.Now().Year())) {
			t.Errorf("Unexpected value: %s", j)
		}

		j, err = nullableTime{}.MarshalJSON()
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if string(j) != `null` {
			t.Errorf("Unexpected value: %s", j)
		}
	})
}
