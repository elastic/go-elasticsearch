// +build !integration

package estransport_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/elastic/go-elasticsearch/estransport"
)

var _ = fmt.Print

var defaultResponse = http.Response{
	Status:        "200 OK",
	StatusCode:    200,
	ContentLength: 2,
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

	b.Run("Defaults            ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp := estransport.New(estransport.Config{
				URLs:      []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			_, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
		}
	})

	b.Run("With Text Logger     ", func(b *testing.B) {
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

	b.Run("With Text Logger (body)", func(b *testing.B) {
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

	b.Run("With JSON Logger", func(b *testing.B) {
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

	b.Run("With JSON Logger (body)", func(b *testing.B) {
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

	b.Run("With LoggerFunc      ", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			tp := estransport.New(estransport.Config{
				URLs:      []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
				Transport: newFakeTransport(b),
				LoggerFunc: func(req http.Request, res http.Response) {
					var b bytes.Buffer
					b.ReadFrom(res.Body)
					fmt.Fprintf(ioutil.Discard, "> %s %s", req.Method, req.URL.String())
					fmt.Fprintf(ioutil.Discard, "< %s", res.Status)
					fmt.Fprintf(ioutil.Discard, "< %s", b.String())
				},
			})

			req, _ := http.NewRequest("GET", "/abc", nil)
			_, err := tp.Perform(req)
			if err != nil {
				b.Fatalf("Unexpected error: %s", err)
			}
		}
	})
}
