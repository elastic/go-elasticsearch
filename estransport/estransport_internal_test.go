// +build !integration

package estransport

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"
)

var (
	_ = fmt.Print
)

type mockTransp struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (t *mockTransp) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.RoundTripFunc(req)
}

func TestTransport(t *testing.T) {
	t.Run("Interface", func(t *testing.T) {
		var _ Interface = New(Config{})
		var _ http.RoundTripper = New(Config{}).transport
	})

	t.Run("Default", func(t *testing.T) {
		tp := New(Config{})
		if tp.transport == nil {
			t.Errorf("Expected the transport to not be nil")
		}
		if tp.transport != http.DefaultTransport {
			t.Errorf("Expected the transport to be http.DefaultTransport, got: %T", tp.transport)
		}
	})

	t.Run("Custom", func(t *testing.T) {
		tp := New(Config{
			URLs: []*url.URL{&url.URL{}},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) { return &http.Response{Status: "MOCK"}, nil },
			},
		})

		res, err := tp.transport.RoundTrip(&http.Request{URL: &url.URL{}})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if res.Status != "MOCK" {
			t.Errorf("Unexpected response from transport: %+v", res)
		}
	})
}

func TestTransportPerform(t *testing.T) {
	t.Run("Executes", func(t *testing.T) {
		u, _ := url.Parse("https://foo.com/bar")
		tp := New(Config{
			URLs: []*url.URL{u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) { return &http.Response{Status: "MOCK"}, nil },
			}})

		req, _ := http.NewRequest("GET", "/abc", nil)

		res, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if res.Status != "MOCK" {
			t.Errorf("Unexpected response: %+v", res)
		}
	})

	t.Run("Sets URL", func(t *testing.T) {
		u, _ := url.Parse("https://foo.com/bar")
		tp := New(Config{URLs: []*url.URL{u}})

		req, _ := http.NewRequest("GET", "/abc", nil)
		tp.setURL(u, req)

		expected := "https://foo.com/bar/abc"

		if req.URL.String() != expected {
			t.Errorf("req.URL: got=%s, want=%s", req.URL, expected)
		}
	})

	t.Run("Sets HTTP Basic Auth", func(t *testing.T) {
		u, _ := url.Parse("https://foo:bar@example.com")
		tp := New(Config{URLs: []*url.URL{u}})

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setBasicAuth(u, req)

		username, password, ok := req.BasicAuth()
		if !ok {
			t.Errorf("Expected the request to have Basic Auth set")
		}

		if username != "foo" || password != "bar" {
			t.Errorf("Unexpected values for username and password: %s:%s", username, password)
		}
	})
}

func TestTransportSelector(t *testing.T) {
	t.Run("Empty URLs", func(t *testing.T) {
		tp := New(Config{})

		u, err := tp.selector.Select()
		if err == nil {
			t.Errorf("Expected error, but got %s", u)
		}
	})

	t.Run("Nil value", func(t *testing.T) {
		tp := New(Config{URLs: []*url.URL{nil}})

		u, err := tp.selector.Select()
		if err == nil {
			t.Errorf("Expected error, but got %s", u)
		}
	})

	t.Run("Single URL", func(t *testing.T) {
		tp := New(Config{URLs: []*url.URL{&url.URL{Scheme: "http", Host: "localhost:9200"}}})

		for i := 0; i < 7; i++ {
			u, err := tp.selector.Select()
			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			if u.String() != "http://localhost:9200" {
				t.Errorf("Unexpected URL, want=http://localhost:9200, got=%s", u)
			}
		}
	})

	t.Run("Two URLs", func(t *testing.T) {
		var u *url.URL

		tp := New(Config{URLs: []*url.URL{
			&url.URL{Scheme: "http", Host: "localhost:9200"},
			&url.URL{Scheme: "http", Host: "localhost:9201"},
		}})

		u, _ = tp.selector.Select()
		if u.String() != "http://localhost:9200" {
			t.Errorf("Unexpected URL, want=http://localhost:9200, got=%s", u)
		}

		u, _ = tp.selector.Select()
		if u.String() != "http://localhost:9201" {
			t.Errorf("Unexpected URL, want=http://localhost:9201, got=%s", u)
		}

		u, _ = tp.selector.Select()
		if u.String() != "http://localhost:9200" {
			t.Errorf("Unexpected URL, want=http://localhost:9200, got=%s", u)
		}
	})

	t.Run("Three URLs", func(t *testing.T) {
		tp := New(Config{URLs: []*url.URL{
			&url.URL{Scheme: "http", Host: "localhost:9200"},
			&url.URL{Scheme: "http", Host: "localhost:9201"},
			&url.URL{Scheme: "http", Host: "localhost:9202"},
		}})

		var expected string
		for i := 0; i < 11; i++ {
			u, err := tp.selector.Select()
			// fmt.Printf("> %s\n", u)

			if err != nil {
				t.Errorf("Unexpected error: %s", err)
			}

			switch i % 3 {
			case 0:
				expected = "http://localhost:9200"
			case 1:
				expected = "http://localhost:9201"
			case 2:
				expected = "http://localhost:9202"
			default:
				t.Fatalf("Unexpected i %% 3: %d", i%3)
			}

			if u.String() != expected {
				t.Errorf("Unexpected URL, want=%s, got=%s", expected, u)
			}
		}
	})
}
