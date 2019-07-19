// +build !integration

package estransport

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
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
			t.Error("Expected the transport to not be nil")
		}
		if tp.transport != http.DefaultTransport {
			t.Errorf("Expected the transport to be http.DefaultTransport, got: %T", tp.transport)
		}
	})

	t.Run("Custom", func(t *testing.T) {
		tp := New(Config{
			URLs: []*url.URL{{}},
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

	t.Run("Sets HTTP Basic Auth from URL", func(t *testing.T) {
		u, _ := url.Parse("https://foo:bar@example.com")
		tp := New(Config{URLs: []*url.URL{u}})

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setAuthorization(u, req)

		username, password, ok := req.BasicAuth()
		if !ok {
			t.Error("Expected the request to have Basic Auth set")
		}

		if username != "foo" || password != "bar" {
			t.Errorf("Unexpected values for username and password: %s:%s", username, password)
		}
	})

	t.Run("Sets HTTP Basic Auth from configuration", func(t *testing.T) {
		u, _ := url.Parse("http://example.com")
		tp := New(Config{URLs: []*url.URL{u}, Username: "foo", Password: "bar"})

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setAuthorization(u, req)

		username, password, ok := req.BasicAuth()
		if !ok {
			t.Errorf("Expected the request to have Basic Auth set")
		}

		if username != "foo" || password != "bar" {
			t.Errorf("Unexpected values for username and password: %s:%s", username, password)
		}
	})

	t.Run("Sets APIKey Authentication from configuration", func(t *testing.T) {
		u, _ := url.Parse("http://example.com")
		tp := New(Config{URLs: []*url.URL{u}, APIKey: "Zm9vYmFy"}) // foobar

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setAuthorization(u, req)

		value := req.Header.Get("Authorization")
		if value == "" {
			t.Errorf("Expected the request to have the Authorization header set")
		}

		if value != "APIKey Zm9vYmFy" {
			t.Errorf(`Unexpected value for Authorization: want="APIKey Zm9vYmFy", got="%s"`, value)
		}
	})

	t.Run("Sets UserAgent", func(t *testing.T) {
		u, _ := url.Parse("http://example.com")
		tp := New(Config{URLs: []*url.URL{u}})

		req, _ := http.NewRequest("GET", "/abc", nil)
		tp.setUserAgent(req)

		if !strings.HasPrefix(req.UserAgent(), "go-elasticsearch") {
			t.Errorf("Unexpected user agent: %s", req.UserAgent())
		}
	})

	t.Run("Error No URL", func(t *testing.T) {
		tp := New(Config{
			URLs: []*url.URL{},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) { return &http.Response{Status: "MOCK"}, nil },
			}})

		req, _ := http.NewRequest("GET", "/abc", nil)

		res, err := tp.Perform(req)
		if err.Error() != "cannot get URL: No URL available" {
			t.Fatalf("Expected error `cannot get URL`: but got error %v, response %v", err, res)
		}
	})
}

func TestTransportPerformRetries(t *testing.T) {
	t.Run("Retry request on error and return response", func(t *testing.T) {
		var (
			i       int
			numReqs = 2
		)

		u, _ := url.Parse("http://foo.bar")
		tp := New(Config{
			URLs: []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					if i == numReqs {
						fmt.Print(": OK\n")
						return &http.Response{Status: "OK"}, nil
					}
					fmt.Print(": ERR\n")
					return nil, fmt.Errorf("Mock error (%d)", i)
				},
			}})

		req, _ := http.NewRequest("GET", "/abc", nil)

		res, err := tp.Perform(req)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if res.Status != "OK" {
			t.Errorf("Unexpected response: %+v", res)
		}

		if i != numReqs {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", numReqs, i)
		}
	})

	t.Run("Retry request on 5xx response and return response", func(t *testing.T) {
		var (
			i       int
			numReqs = 2
		)

		u, _ := url.Parse("http://foo.bar")
		tp := New(Config{
			URLs: []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					if i == numReqs {
						fmt.Print(": 200\n")
						return &http.Response{StatusCode: 200}, nil
					}
					fmt.Print(": 502\n")
					return &http.Response{StatusCode: 502}, nil
				},
			}})

		req, _ := http.NewRequest("GET", "/abc", nil)

		res, err := tp.Perform(req)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if res.StatusCode != 200 {
			t.Errorf("Unexpected response: %+v", res)
		}

		if i != numReqs {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", numReqs, i)
		}
	})

	t.Run("Retry request and return error when max retries exhausted", func(t *testing.T) {
		var (
			i       int
			numReqs = 3
		)

		u, _ := url.Parse("http://foo.bar")
		tp := New(Config{
			URLs: []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					fmt.Print(": ERR\n")
					return nil, fmt.Errorf("Mock error (%d)", i)
				},
			}})

		req, _ := http.NewRequest("GET", "/abc", nil)

		res, err := tp.Perform(req)

		if err == nil {
			t.Fatalf("Expected error, got: %v", err)
		}

		if res != nil {
			t.Errorf("Unexpected response: %+v", res)
		}

		if i != numReqs {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", numReqs, i)
		}
	})

	t.Run("Don't retry for single URL", func(t *testing.T) {
		var (
			i       int
			numReqs = 1
		)

		u, _ := url.Parse("http://foo.bar")
		tp := New(Config{
			URLs: []*url.URL{u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					fmt.Print(": ERR\n")
					return nil, fmt.Errorf("Mock error (%d)", i)
				},
			}})

		req, _ := http.NewRequest("GET", "/abc", nil)

		res, err := tp.Perform(req)

		if err == nil {
			t.Fatalf("Expected error, got: %v", err)
		}

		if res != nil {
			t.Errorf("Unexpected response: %+v", res)
		}

		if i != numReqs {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", numReqs, i)
		}
	})
}

func TestTransportSelector(t *testing.T) {
	t.Run("Nil value", func(t *testing.T) {
		tp := New(Config{URLs: []*url.URL{nil}})

		u, err := tp.selector.Select()
		if err == nil {
			t.Errorf("Expected error, but got %s", u)
		}
	})

	t.Run("No URL", func(t *testing.T) {
		tp := New(Config{})

		u, err := tp.selector.Select()
		if err == nil {
			t.Errorf("Expected error, but got %s", u)
		}
	})

	t.Run("Single URL", func(t *testing.T) {
		tp := New(Config{URLs: []*url.URL{{Scheme: "http", Host: "localhost:9200"}}})

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
			{Scheme: "http", Host: "localhost:9200"},
			{Scheme: "http", Host: "localhost:9201"},
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
			{Scheme: "http", Host: "localhost:9200"},
			{Scheme: "http", Host: "localhost:9201"},
			{Scheme: "http", Host: "localhost:9202"},
		}})

		var expected string
		for i := 0; i < 11; i++ {
			u, err := tp.selector.Select()

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

func TestURLs(t *testing.T) {
	t.Run("Returns URLs", func(t *testing.T) {
		tp := New(Config{URLs: []*url.URL{
			{Scheme: "http", Host: "localhost:9200"},
			{Scheme: "http", Host: "localhost:9201"},
		}})
		urls := tp.URLs()
		if len(urls) != 2 {
			t.Errorf("Expected get 2 urls, but got: %d", len(urls))
		}
		if urls[0].Host != "localhost:9200" {
			t.Errorf("Unexpected URL, want=localhost:9200, got=%s", urls[0].Host)
		}
	})
}

func TestVersion(t *testing.T) {
	if Version == "" {
		t.Error("Version is empty")
	}
}
