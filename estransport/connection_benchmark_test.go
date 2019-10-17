// +build !integration

package estransport

import (
	"net/url"
	"testing"
)

func BenchmarkRoundRobinConnectionPool(b *testing.B) {
	b.ReportAllocs()

	b.Run("Next()", func(b *testing.B) {
		pool := newRoundRobinConnectionPool(
			&url.URL{Scheme: "http", Host: "foo1"},
			&url.URL{Scheme: "http", Host: "foo2"})

		for i := 0; i < b.N; i++ {
			_, err := pool.Next()
			if err != nil {
				b.Errorf("Unexpected error: %v", err)
			}
		}
	})

	b.Run("Remove()", func(b *testing.B) {
		pool := newRoundRobinConnectionPool(
			&url.URL{Scheme: "http", Host: "foo1"},
			&url.URL{Scheme: "http", Host: "foo2"})

		c, _ := pool.Next()

		for i := 0; i < b.N; i++ {
			if err := pool.Remove(c); err != nil {
				b.Errorf("Unexpected error: %v", err)
			}

			pool.live = append(pool.live, c)
		}
	})

	b.Run("Resurrect()", func(b *testing.B) {
		pool := newRoundRobinConnectionPool(
			&url.URL{Scheme: "http", Host: "foo1"},
			&url.URL{Scheme: "http", Host: "foo2"})

		c, _ := pool.Next()
		pool.Remove(c)

		for i := 0; i < b.N; i++ {
			if err := pool.Resurrect(c); err != nil {
				b.Errorf("Unexpected error: %v", err)
			}
		}
	})
}
