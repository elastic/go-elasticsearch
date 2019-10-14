// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package estransport_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/v7/estransport"
)

var defaultResponse = http.Response{
	Status:        "200 OK",
	StatusCode:    200,
	ContentLength: 13,
	Header:        http.Header(map[string][]string{"Content-Type": {"application/json"}}),
	Body:          ioutil.NopCloser(strings.NewReader(`{"foo":"bar"}`)),
}

type FakeTransport struct {
	FakeResponse *http.Response
}

func (t *FakeTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.FakeResponse, nil
}

func newFakeTransport(b *testing.B) *FakeTransport {
	return &FakeTransport{FakeResponse: &defaultResponse}
}

func BenchmarkTransport(b *testing.B) {
	b.ReportAllocs()

	b.Run("Defaults", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp := estransport.New(estransport.Config{
				URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			_, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
		}
	})
}
