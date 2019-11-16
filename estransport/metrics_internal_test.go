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
		if len(m.Connections) != 3 {
			t.Errorf("Unexpected output: %+v", m.Connections)
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
		var m ConnectionMetric

		m = ConnectionMetric{URL: "http://foo1"}

		if m.String() != "{http://foo1}" {
			t.Errorf("Unexpected output: %s", m)
		}

		tt, _ := time.Parse(time.RFC3339, "2010-11-11T11:00:00Z")
		m = ConnectionMetric{
			URL:       "http://foo2",
			IsDead:    true,
			Failures:  123,
			DeadSince: &tt,
		}

		match, err := regexp.MatchString(
			`{http://foo2 dead=true failures=123 dead_since=Nov 11 \d+:00:00}`,
			m.String(),
		)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if !match {
			t.Errorf("Unexpected output: %s", m)
		}
	})
}
