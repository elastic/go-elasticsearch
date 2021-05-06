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

package estransport_test

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/estransport"
)

func BenchmarkTransportLogger(b *testing.B) {
	b.ReportAllocs()

	b.Run("Text", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp, _ := estransport.New(estransport.Config{
				URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
				Logger:    &estransport.TextLogger{Output: ioutil.Discard},
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			_, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
		}
	})

	b.Run("Text-Body", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp, _ := estransport.New(estransport.Config{
				URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
				Logger:    &estransport.TextLogger{Output: ioutil.Discard, EnableRequestBody: true, EnableResponseBody: true},
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			res, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}

			body, err := ioutil.ReadAll(res.Body)
			if err != nil {
				b.Fatalf("Error reading response body: %s", err)
			}
			res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
			if len(body) < 13 {
				b.Errorf("Error reading response body bytes, want=13, got=%d", len(body))
			}
		}
	})

	b.Run("JSON", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp, _ := estransport.New(estransport.Config{
				URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
				Logger:    &estransport.JSONLogger{Output: ioutil.Discard},
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			_, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
		}
	})

	b.Run("JSON-Body", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp, _ := estransport.New(estransport.Config{
				URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
				Logger:    &estransport.JSONLogger{Output: ioutil.Discard, EnableRequestBody: true, EnableResponseBody: true},
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			_, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
		}
	})
}
