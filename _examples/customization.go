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

//go:build ignore
// +build ignore

package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v9"
)

// This example demonstrates how to observe or manipulate each request and
// response using an Interceptor on the typed client. Interceptors are the
// recommended way to add cross-cutting behaviour (custom headers, logging,
// per-request auth, metrics, tracing) without replacing the HTTP transport.
//
// CountingInterceptor adds custom headers to each request, logs the request
// and response to stdout, and counts the total number of requests.
func CountingInterceptor(count *atomic.Uint64) elastictransport.InterceptorFunc {
	return func(next elastictransport.RoundTripFunc) elastictransport.RoundTripFunc {
		return func(req *http.Request) (*http.Response, error) {
			count.Add(1)

			req.Header.Set("Accept", "application/yaml")
			req.Header.Set("X-Request-ID", "foo-123")

			res, err := next(req)

			var b bytes.Buffer
			b.WriteString(strings.Repeat("-", 80) + "\n")
			fmt.Fprintf(&b, "%s %s", req.Method, req.URL.String())
			if err == nil {
				fmt.Fprintf(&b, " [%s] %s\n", res.Status, res.Header.Get("Content-Type"))
			} else {
				fmt.Fprintf(&b, "ERROR: %s\n", err)
			}
			b.WriteTo(os.Stdout)

			return res, err
		}
	}
}

func main() {
	var wg sync.WaitGroup
	var count atomic.Uint64

	es, _ := elasticsearch.NewTyped(
		elasticsearch.WithTransportOptions(elastictransport.WithInterceptors(
			CountingInterceptor(&count),
		)),
	)

	ctx := context.Background()
	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			es.Info().Do(ctx)
		}()
	}
	wg.Wait()

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%80s\n", fmt.Sprintf("Total Requests: %d", count.Load()))
}
