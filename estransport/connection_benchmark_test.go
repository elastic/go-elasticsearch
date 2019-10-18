// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package estransport

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"testing"

	_ "net/http/pprof"
)

func init() {
	go func() { log.Fatalln(http.ListenAndServe("localhost:6060", nil)) }()
}

func BenchmarkSingleConnectionPool(b *testing.B) {
	b.ReportAllocs()

	b.Run("Next()", func(b *testing.B) {
		pool := newSingleConnectionPool(&url.URL{Scheme: "http", Host: "foo1"})

		b.Run("Single    ", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := pool.Next()
				if err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
			}
		})

		b.Run("Parallel (1000)", func(b *testing.B) {
			b.SetParallelism(1000)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := pool.Next()
					if err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})
	})

	b.Run("Remove()", func(b *testing.B) {
		pool := newSingleConnectionPool(&url.URL{Scheme: "http", Host: "foo1"})

		b.Run("Single    ", func(b *testing.B) {
			c, _ := pool.Next()

			for i := 0; i < b.N; i++ {
				if err := pool.Remove(c); err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
			}
		})

		b.Run("Parallel (1000)", func(b *testing.B) {
			b.SetParallelism(1000)
			b.RunParallel(func(pb *testing.PB) {
				c, _ := pool.Next()

				for pb.Next() {
					if err := pool.Remove(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})
	})
}

func BenchmarkRoundRobinConnectionPool(b *testing.B) {
	b.ReportAllocs()

	var urls []*url.URL
	for i := 0; i < 1000; i++ {
		urls = append(urls, &url.URL{Scheme: "http", Host: fmt.Sprintf("foo%d", i)})
	}

	b.Run("Next()", func(b *testing.B) {
		pool := newRoundRobinConnectionPool(urls...)

		b.Run("Single    ", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := pool.Next()
				if err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
			}
		})

		b.Run("Parallel (100)", func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := pool.Next()
					if err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})

		b.Run("Parallel (1000)", func(b *testing.B) {
			b.SetParallelism(1000)
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					_, err := pool.Next()
					if err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})
	})

	b.Run("Remove()", func(b *testing.B) {
		pool := newRoundRobinConnectionPool(urls...)

		b.Run("Single    ", func(b *testing.B) {
			c, _ := pool.Next()

			for i := 0; i < b.N; i++ {
				if err := pool.Remove(c); err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
			}
		})

		b.Run("Parallel (10)", func(b *testing.B) {
			b.SetParallelism(10)
			b.RunParallel(func(pb *testing.PB) {
				c, _ := pool.Next()

				for pb.Next() {
					if err := pool.Remove(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})

		b.Run("Parallel (100)", func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				c, _ := pool.Next()

				for pb.Next() {
					if err := pool.Remove(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})
	})

	b.Run("Resurrect()", func(b *testing.B) {
		pool := newRoundRobinConnectionPool(urls...)

		b.Run("Single       ", func(b *testing.B) {
			c, _ := pool.Next()
			pool.Remove(c)

			for i := 0; i < b.N; i++ {
				if err := pool.Resurrect(c); err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
			}
		})

		b.Run("Parallel (10)", func(b *testing.B) {
			b.SetParallelism(10)
			b.RunParallel(func(pb *testing.PB) {
				c, _ := pool.Next()
				pool.Remove(c)

				for pb.Next() {
					if err := pool.Resurrect(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})

		b.Run("Parallel (100)", func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				c, _ := pool.Next()
				pool.Remove(c)

				for pb.Next() {
					if err := pool.Resurrect(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})
	})
}
