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

package fasthttp_test

import (
	"testing"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/_examples/fasthttp"
)

func BenchmarkHTTPClient(b *testing.B) {
	b.ReportAllocs()

	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		b.Fatalf("ERROR: %s", err)
	}

	b.Run("Info()", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if res, err := client.Info(); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			} else {
				res.Body.Close()
			}
		}
	})
}

func BenchmarkFastHTTPClient(b *testing.B) {
	b.ReportAllocs()

	client, err := elasticsearch.NewClient(
		elasticsearch.Config{Transport: &fasthttp.Transport{}},
	)
	if err != nil {
		b.Fatalf("ERROR: %s", err)
	}

	b.Run("Info()", func(b *testing.B) {
		b.ResetTimer()

		for i := 0; i < b.N; i++ {
			if res, err := client.Info(); err != nil {
				b.Errorf("Unexpected error when getting a response: %s", err)
			} else {
				res.Body.Close()
			}
		}
	})
}
