// +build !integration

package estransport

import (
	"net/url"
	"testing"
)

func TestConnectionPoolNext(t *testing.T) {
	t.Run("No URL", func(t *testing.T) {
		tp := New(Config{})

		c, err := tp.pool.Next()
		if err == nil {
			t.Errorf("Expected error, but got %s", c.URL)
		}
	})

	t.Run("Single URL", func(t *testing.T) {
		tp := New(Config{URLs: []*url.URL{{Scheme: "http", Host: "localhost:9200"}}})

		if _, ok := tp.pool.(*singleConnectionPool); !ok {
			t.Errorf("Expected connection to be singleConnectionPool, got: %T", tp)
		}

		for i := 0; i < 7; i++ {
			c, err := tp.pool.Next()
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if c.URL.String() != "http://localhost:9200" {
				t.Errorf("Unexpected URL, want=http://localhost:9200, got=%s", c.URL)
			}
		}
	})

	t.Run("Two URLs", func(t *testing.T) {
		var c *Connection

		tp := New(Config{URLs: []*url.URL{
			{Scheme: "http", Host: "localhost:9200"},
			{Scheme: "http", Host: "localhost:9201"},
		}})

		if _, ok := tp.pool.(*roundRobinConnectionPool); !ok {
			t.Errorf("Expected connection to be roundRobinConnectionPool, got: %T", tp)
		}

		c, _ = tp.pool.Next()
		if c.URL.String() != "http://localhost:9200" {
			t.Errorf("Unexpected URL, want=http://localhost:9200, got=%s", c.URL)
		}

		c, _ = tp.pool.Next()
		if c.URL.String() != "http://localhost:9201" {
			t.Errorf("Unexpected URL, want=http://localhost:9201, got=%s", c.URL)
		}

		c, _ = tp.pool.Next()
		if c.URL.String() != "http://localhost:9200" {
			t.Errorf("Unexpected URL, want=http://localhost:9200, got=%s", c.URL)
		}
	})

	t.Run("Three URLs", func(t *testing.T) {
		tp := New(Config{URLs: []*url.URL{
			{Scheme: "http", Host: "localhost:9200"},
			{Scheme: "http", Host: "localhost:9201"},
			{Scheme: "http", Host: "localhost:9202"},
		}})

		var expected string
		for i := 0; i < 11; i++ {
			c, err := tp.pool.Next()

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			switch i % 3 {
			case 0:
				expected = "http://localhost:9200"
			case 1:
				expected = "http://localhost:9201"
			case 2:
				expected = "http://localhost:9202"
			default:
				t.Fatalf("Unexpected i %% 3: %d", i%3)
			}

			if c.URL.String() != expected {
				t.Errorf("Unexpected URL, want=%s, got=%s", expected, c.URL)
			}
		}
	})
}
