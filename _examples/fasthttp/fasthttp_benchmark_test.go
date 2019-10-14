// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

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
