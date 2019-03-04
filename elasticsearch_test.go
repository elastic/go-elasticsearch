// +build !integration

package elasticsearch

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"testing"

	"github.com/elastic/go-elasticsearch/estransport"
)

func TestClientConfiguration(t *testing.T) {
	t.Parallel()

	t.Run("With empty", func(t *testing.T) {
		_, err := NewDefaultClient()

		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})

	t.Run("With address", func(t *testing.T) {
		c, err := NewClient(Config{Addresses: []string{"http://localhost:8080//"}})
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		u := c.Transport.(*estransport.Client).URLs()[0].String()

		if u != "http://localhost:8080" {
			t.Errorf("Unexpected URL, want=http://localhost:8080, got=%s", u)
		}
	})

	t.Run("With URL from environment", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", "http://example.com")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewDefaultClient()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		u := c.Transport.(*estransport.Client).URLs()[0].String()

		if u != "http://example.com" {
			t.Errorf("Unexpected URL, want=http://example.com, got=%s", u)
		}
	})

	t.Run("With URL from environment and cfg.Addresses", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", "http://example.com")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		_, err := NewClient(Config{Addresses: []string{"http://localhost:8080//"}})
		if err.Error() != "cannot create client: both ELASTICSEARCH_URL and Addresses are set" {
			t.Errorf("Expected error: %v", errors.New("cannot create client: both ELASTICSEARCH_URL and Addresses are set"))
		}
	})

	t.Run("With invalid URL", func(t *testing.T) {
		u := ":foo"
		_, err := NewClient(Config{Addresses: []string{u}})

		if err == nil {
			t.Errorf("Expected error for URL %q, got %v", u, err)
		}
	})

	t.Run("With invalid URL from environment", func(t *testing.T) {
		os.Setenv("ELASTICSEARCH_URL", ":foobar")
		defer func() { os.Setenv("ELASTICSEARCH_URL", "") }()

		c, err := NewDefaultClient()
		if err == nil {
			t.Errorf("Expected error, got: %+v", c)
		}
	})
}

var called bool

type mockTransp struct{}

func (t *mockTransp) RoundTrip(req *http.Request) (*http.Response, error) {
	called = true
	return &http.Response{}, nil
}

func TestClientInterface(t *testing.T) {
	t.Run("Transport", func(t *testing.T) {
		c, err := NewClient(Config{Transport: &mockTransp{}})

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if called != false { // megacheck ignore
			t.Errorf("Unexpected call to transport by client")
		}

		c.Perform(&http.Request{URL: &url.URL{}}) // errcheck ignore

		if called != true { // megacheck ignore
			t.Errorf("Expected client to call transport")
		}
	})
}

func TestAddrsToURLs(t *testing.T) {
	tt := []struct {
		name          string
		urls          []string
		uStructs      []*url.URL
		expectedError error
	}{
		{
			name: "all ok",
			urls: []string{"http://example.com", "http://192.168.255.255", "https://www.elastic.co/"},
			uStructs: []*url.URL{
				{
					Scheme: "http",
					Host:   "example.com",
				},
				{
					Scheme: "http",
					Host:   "192.168.255.255",
				},
				{
					Scheme: "https",
					Host:   "www.elastic.co",
				},
			},
			expectedError: nil,
		},
		{
			name:          "parse error: invalid url",
			urls:          []string{"://эк?:%;.com"},
			uStructs:      nil,
			expectedError: fmt.Errorf("cannot parse url: %v", errors.New("parse ://эк?:%;.com: missing protocol scheme")),
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			res, err := addrsToURLs(tc.urls)
			for i := range tc.uStructs {
				if res[i].Scheme != tc.uStructs[i].Scheme {
					t.Errorf("test case name: %s\nexpected Scheme %s\nactual %s", tc.name, tc.uStructs[i].Scheme, res[i].Scheme)
				}
			}
			for i := range tc.uStructs {
				if res[i].Host != tc.uStructs[i].Host {
					t.Errorf("test case name: %s\nexpected Host %s\nactual %s", tc.name, tc.uStructs[i].Host, res[i].Host)
				}
			}
			if err != tc.expectedError {
				if err == nil || err.Error() != tc.expectedError.Error() {
					t.Errorf("test case name: %s\nexpected error: %v\nactual error: %v", tc.name, tc.expectedError, err)
				}
			}
		})
	}
}
