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

// +build ignore

package main

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/elastic/go-elasticsearch/v7"
)

// This example demonstrates how to provide a custom transport implementation to the client
// in order to read or manipulate the requests and responses, eg. for custom logging
// implementation, passing custom headers to the request, and so on.
//
// CountingTransport adds a custom header to the request, logs the information about
// the request and response, and counts the number of requests to the client.
//
// Since it implements the `http.RoundTripper` interface, it can be passed
// to the client as a custom HTTP transport implementation.
//
type CountingTransport struct {
	count uint64
}

// RoundTrip executes a request and returns a response.
//
func (t *CountingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var b bytes.Buffer

	atomic.AddUint64(&t.count, 1)

	req.Header.Set("Accept", "application/yaml")
	req.Header.Set("X-Request-ID", "foo-123")

	res, err := http.DefaultTransport.RoundTrip(req)

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

func main() {
	var wg sync.WaitGroup

	// Create the custom transport.
	//
	tp := CountingTransport{}

	// Pass the custom transport to the client.
	//
	es, _ := elasticsearch.NewClient(
		elasticsearch.Config{Transport: &tp},
	)

	for i := 0; i < 25; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			es.Info()
		}()
	}
	wg.Wait()

	fmt.Println(strings.Repeat("=", 80))
	fmt.Printf("%80s\n", fmt.Sprintf("Total Requests: %d", atomic.LoadUint64(&tp.count)))
}
