// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

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
		pool := &singleConnectionPool{connection: &Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}}}

		b.Run("Single          ", func(b *testing.B) {
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

	b.Run("OnFailure()", func(b *testing.B) {
		pool := &singleConnectionPool{connection: &Connection{URL: &url.URL{Scheme: "http", Host: "foo1"}}}

		b.Run("Single     ", func(b *testing.B) {
			c, _ := pool.Next()

			for i := 0; i < b.N; i++ {
				if err := pool.OnFailure(c); err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
			}
		})

		b.Run("Parallel (1000)", func(b *testing.B) {
			b.SetParallelism(1000)
			b.RunParallel(func(pb *testing.PB) {
				c, _ := pool.Next()

				for pb.Next() {
					if err := pool.OnFailure(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})
	})
}

func BenchmarkStatusConnectionPool(b *testing.B) {
	b.ReportAllocs()

	var conns []*Connection
	for i := 0; i < 1000; i++ {
		conns = append(conns, &Connection{URL: &url.URL{Scheme: "http", Host: fmt.Sprintf("foo%d", i)}})
	}

	b.Run("Next()", func(b *testing.B) {
		pool := &statusConnectionPool{
			live:     conns,
			selector: &roundRobinSelector{curr: -1},
		}

		b.Run("Single     ", func(b *testing.B) {
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

	b.Run("OnFailure()", func(b *testing.B) {
		pool := &statusConnectionPool{
			live:     conns,
			selector: &roundRobinSelector{curr: -1},
		}

		b.Run("Single     ", func(b *testing.B) {
			c, err := pool.Next()
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}

			for i := 0; i < b.N; i++ {
				if err := pool.OnFailure(c); err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
			}
		})

		b.Run("Parallel (10)", func(b *testing.B) {
			b.SetParallelism(10)
			b.RunParallel(func(pb *testing.PB) {
				c, err := pool.Next()
				if err != nil {
					b.Fatalf("Unexpected error: %s", err)
				}

				for pb.Next() {
					if err := pool.OnFailure(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})

		b.Run("Parallel (100)", func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				c, err := pool.Next()
				if err != nil {
					b.Fatalf("Unexpected error: %s", err)
				}

				for pb.Next() {
					if err := pool.OnFailure(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})
	})

	b.Run("OnSuccess()", func(b *testing.B) {
		pool := &statusConnectionPool{
			live:     conns,
			selector: &roundRobinSelector{curr: -1},
		}

		b.Run("Single     ", func(b *testing.B) {
			c, err := pool.Next()
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}

			for i := 0; i < b.N; i++ {
				if err := pool.OnSuccess(c); err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
			}
		})

		b.Run("Parallel (10)", func(b *testing.B) {
			b.SetParallelism(10)
			b.RunParallel(func(pb *testing.PB) {
				c, err := pool.Next()
				if err != nil {
					b.Fatalf("Unexpected error: %s", err)
				}

				for pb.Next() {
					if err := pool.OnSuccess(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})

		b.Run("Parallel (100)", func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				c, err := pool.Next()
				if err != nil {
					b.Fatalf("Unexpected error: %s", err)
				}

				for pb.Next() {
					if err := pool.OnSuccess(c); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
				}
			})
		})
	})

	b.Run("resurrect()", func(b *testing.B) {
		pool := &statusConnectionPool{
			live:     conns,
			selector: &roundRobinSelector{curr: -1},
		}

		b.Run("Single", func(b *testing.B) {
			c, err := pool.Next()
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
			err = pool.OnFailure(c)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}

			for i := 0; i < b.N; i++ {
				pool.Lock()
				if err := pool.resurrect(c, true); err != nil {
					b.Errorf("Unexpected error: %v", err)
				}
				pool.Unlock()
			}
		})

		b.Run("Parallel (10)", func(b *testing.B) {
			b.SetParallelism(10)
			b.RunParallel(func(pb *testing.PB) {
				c, err := pool.Next()
				if err != nil {
					b.Fatalf("Unexpected error: %s", err)
				}
				err = pool.OnFailure(c)
				if err != nil {
					b.Fatalf("Unexpected error: %s", err)
				}

				for pb.Next() {
					pool.Lock()
					if err := pool.resurrect(c, true); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
					pool.Unlock()
				}
			})
		})

		b.Run("Parallel (100)", func(b *testing.B) {
			b.SetParallelism(100)
			b.RunParallel(func(pb *testing.PB) {
				c, err := pool.Next()
				if err != nil {
					b.Fatalf("Unexpected error: %s", err)
				}
				err = pool.OnFailure(c)
				if err != nil {
					b.Fatalf("Unexpected error: %s", err)
				}

				for pb.Next() {
					pool.Lock()
					if err := pool.resurrect(c, true); err != nil {
						b.Errorf("Unexpected error: %v", err)
					}
					pool.Unlock()
				}
			})
		})
	})
}
