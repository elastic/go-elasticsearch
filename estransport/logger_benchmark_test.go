// +build !integration

package estransport_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/elastic/go-elasticsearch/estransport"
)

func BenchmarkTransportLogger(b *testing.B) {
	b.ReportAllocs()

	b.Run("Text", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp := estransport.New(estransport.Config{
				URLs:      []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
				LogOutput: ioutil.Discard,
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
			tp := estransport.New(estransport.Config{
				URLs:            []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
				Transport:       newFakeTransport(b),
				LogOutput:       ioutil.Discard,
				LogRequestBody:  true,
				LogResponseBody: true,
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			_, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
		}
	})

	b.Run("JSON", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp := estransport.New(estransport.Config{
				URLs:      []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
				LogOutput: ioutil.Discard,
				LogFormat: estransport.LogFormatJSON,
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
			tp := estransport.New(estransport.Config{
				URLs:            []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
				Transport:       newFakeTransport(b),
				LogOutput:       ioutil.Discard,
				LogRequestBody:  true,
				LogResponseBody: true,
				LogFormat:       estransport.LogFormatJSON,
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			_, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
		}
	})
}
