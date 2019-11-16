// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package estransport

import (
	"net/url"
	"regexp"
	"testing"
	"time"
)

func TestSingleConnectionPoolNext(t *testing.T) {
	t.Run("Single URL", func(t *testing.T) {
		pool := &singleConnectionPool{
			connection: &Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}},
		}

		for i := 0; i < 7; i++ {
			c, err := pool.Next()
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if c.URL.String() != "http://foo1" {
				t.Errorf("Unexpected URL, want=http://foo1, got=%s", c.URL)
			}
		}
	})
}

func TestSingleConnectionPoolOnFailure(t *testing.T) {
	t.Run("Noop", func(t *testing.T) {
		pool := &singleConnectionPool{
			connection: &Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}},
		}

		if err := pool.OnFailure(&Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}}); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})
}

func TestStatusConnectionPoolNext(t *testing.T) {
	t.Run("No URL", func(t *testing.T) {
		pool := &statusConnectionPool{}

		c, err := pool.Next()
		if err == nil {
			t.Errorf("Expected error, but got: %s", c.URL)
		}
	})

	t.Run("Two URLs", func(t *testing.T) {
		var c *Connection

		pool := &statusConnectionPool{
			live: []*Connection{
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}},
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo2"}},
			},
			selector: &roundRobinSelector{curr: -1},
		}

		c, _ = pool.Next()

		if c.URL.String() != "http://foo1" {
			t.Errorf("Unexpected URL, want=foo1, got=%s", c.URL)
		}

		c, _ = pool.Next()
		if c.URL.String() != "http://foo2" {
			t.Errorf("Unexpected URL, want=http://foo2, got=%s", c.URL)
		}

		c, _ = pool.Next()
		if c.URL.String() != "http://foo1" {
			t.Errorf("Unexpected URL, want=http://foo1, got=%s", c.URL)
		}
	})

	t.Run("Three URLs", func(t *testing.T) {
		pool := &statusConnectionPool{
			live: []*Connection{
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}},
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo2"}},
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo3"}},
			},
			selector: &roundRobinSelector{curr: -1},
		}

		var expected string
		for i := 0; i < 11; i++ {
			c, err := pool.Next()

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			switch i % len(pool.live) {
			case 0:
				expected = "http://foo1"
			case 1:
				expected = "http://foo2"
			case 2:
				expected = "http://foo3"
			default:
				t.Fatalf("Unexpected i %% 3: %d", i%3)
			}

			if c.URL.String() != expected {
				t.Errorf("Unexpected URL, want=%s, got=%s", expected, c.URL)
			}
		}
	})

	t.Run("Resurrect dead connection when no live is available", func(t *testing.T) {
		pool := &statusConnectionPool{
			live: []*Connection{},
			dead: []*Connection{
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}, Failures: 3},
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo2"}, Failures: 1},
			},
			selector: &roundRobinSelector{curr: -1},
		}

		c, err := pool.Next()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if c == nil {
			t.Errorf("Expected connection, got nil: %s", c)
		}

		if c.URL.String() != "http://foo2" {
			t.Errorf("Expected <http://foo2>, got: %s", c.URL.String())
		}

		if c.IsDead {
			t.Errorf("Expected connection to be live, got: %s", c)
		}

		if len(pool.live) != 1 {
			t.Errorf("Expected 1 connection in live list, got: %s", pool.live)
		}

		if len(pool.dead) != 1 {
			t.Errorf("Expected 1 connection in dead list, got: %s", pool.dead)
		}
	})
}

func TestStatusConnectionPoolOnSuccess(t *testing.T) {
	t.Run("Move connection to live list and mark it as healthy", func(t *testing.T) {
		pool := &statusConnectionPool{
			dead: []*Connection{
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}, Failures: 3, IsDead: true},
			},
			selector: &roundRobinSelector{curr: -1},
		}

		conn := pool.dead[0]

		if err := pool.OnSuccess(conn); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if conn.IsDead {
			t.Errorf("Expected the connection to be live; %s", conn)
		}

		if !conn.DeadSince.IsZero() {
			t.Errorf("Unexpected value for DeadSince: %s", conn.DeadSince)
		}

		if len(pool.live) != 1 {
			t.Errorf("Expected 1 live connection, got: %d", len(pool.live))
		}

		if len(pool.dead) != 0 {
			t.Errorf("Expected 0 dead connections, got: %d", len(pool.dead))
		}
	})
}

func TestStatusConnectionPoolOnFailure(t *testing.T) {
	t.Run("Remove connection, mark it, and sort dead connections", func(t *testing.T) {
		pool := &statusConnectionPool{
			live: []*Connection{
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}},
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo2"}},
			},
			dead: []*Connection{
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo3"}, Failures: 0},
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo4"}, Failures: 99},
			},
			selector: &roundRobinSelector{curr: -1},
		}

		conn := pool.live[0]

		if err := pool.OnFailure(conn); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if !conn.IsDead {
			t.Errorf("Expected the connection to be dead; %s", conn)
		}

		if conn.DeadSince.IsZero() {
			t.Errorf("Unexpected value for DeadSince: %s", conn.DeadSince)
		}

		if len(pool.live) != 1 {
			t.Errorf("Expected 1 live connection, got: %d", len(pool.live))
		}

		if len(pool.dead) != 3 {
			t.Errorf("Expected 3 dead connections, got: %d", len(pool.dead))
		}

		expected := []string{
			"http://foo4",
			"http://foo1",
			"http://foo3",
		}

		for i, u := range expected {
			if pool.dead[i].URL.String() != u {
				t.Errorf("Unexpected value for item %d in pool.dead: %s", i, pool.dead[i].URL.String())
			}
		}
	})

	t.Run("Short circuit when the connection is already dead", func(t *testing.T) {
		pool := &statusConnectionPool{
			live: []*Connection{
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}},
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo2"}},
				&Connection{URL: &url.URL{Scheme: "http", Host: "foo3"}},
			},
			selector: &roundRobinSelector{curr: -1},
		}

		conn := pool.live[0]
		conn.IsDead = true

		if err := pool.OnFailure(conn); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if len(pool.dead) != 0 {
			t.Errorf("Expected the dead list to be empty, got: %s", pool.dead)
		}
	})
}

func TestStatusConnectionPoolResurrect(t *testing.T) {
	t.Run("Mark the connection as dead and add/remove it to the lists", func(t *testing.T) {
		pool := &statusConnectionPool{
			live:     []*Connection{},
			dead:     []*Connection{&Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}, IsDead: true}},
			selector: &roundRobinSelector{curr: -1},
		}

		conn := pool.dead[0]

		if err := pool.resurrect(conn, true); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if conn.IsDead {
			t.Errorf("Expected connection to be dead, got: %s", conn)
		}

		if len(pool.dead) != 0 {
			t.Errorf("Expected no dead connections, got: %s", pool.dead)
		}

		if len(pool.live) != 1 {
			t.Errorf("Expected 1 live connection, got: %s", pool.live)
		}
	})

	t.Run("Short circuit removal when the connection is not in the dead list", func(t *testing.T) {
		pool := &statusConnectionPool{
			dead:     []*Connection{&Connection{URL: &url.URL{Scheme: "http", Host: "bar"}, IsDead: true}},
			selector: &roundRobinSelector{curr: -1},
		}

		conn := &Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}, IsDead: true}

		if err := pool.resurrect(conn, true); err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if len(pool.live) != 1 {
			t.Errorf("Expected 1 live connection, got: %s", pool.live)
		}

		if len(pool.dead) != 1 {
			t.Errorf("Expected 1 dead connection, got: %s", pool.dead)
		}
	})

	t.Run("Schedule resurrect", func(t *testing.T) {
		defaultResurrectTimeoutInitial = 0
		defer func() { defaultResurrectTimeoutInitial = 60 * time.Second }()

		pool := &statusConnectionPool{
			live: []*Connection{},
			dead: []*Connection{
				&Connection{
					URL:       &url.URL{Scheme: "http", Host: "foo1"},
					Failures:  100,
					IsDead:    true,
					DeadSince: time.Now().UTC(),
				},
			},
			selector: &roundRobinSelector{curr: -1},
		}

		conn := pool.dead[0]
		pool.scheduleResurrect(conn)
		time.Sleep(50 * time.Millisecond)

		pool.Lock()
		defer pool.Unlock()

		if len(pool.live) != 1 {
			t.Errorf("Expected 1 live connection, got: %s", pool.live)
		}
		if len(pool.dead) != 0 {
			t.Errorf("Expected no dead connections, got: %s", pool.dead)
		}
	})
}

func TestConnection(t *testing.T) {
	t.Run("String", func(t *testing.T) {
		conn := &Connection{
			URL:       &url.URL{Scheme: "http", Host: "foo1"},
			Failures:  10,
			IsDead:    true,
			DeadSince: time.Now().UTC(),
		}

		match, err := regexp.MatchString(
			`<http://foo1> dead=true failures=10`,
			conn.String(),
		)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if !match {
			t.Errorf("Unexpected output: %s", conn)
		}
	})
}
