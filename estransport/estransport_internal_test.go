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

package estransport

import (
	"bytes"
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"
)

var (
	_ = fmt.Print
)

func init() {
	rand.Seed(time.Now().Unix())
}

type mockTransp struct {
	RoundTripFunc func(req *http.Request) (*http.Response, error)
}

func (t *mockTransp) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.RoundTripFunc(req)
}

type mockNetError struct{ error }

func (e *mockNetError) Timeout() bool   { return false }
func (e *mockNetError) Temporary() bool { return false }

func TestTransport(t *testing.T) {
	t.Run("Interface", func(t *testing.T) {
		tp, _ := New(Config{})
		var _ Interface = tp
		var _ http.RoundTripper = tp.transport
	})

	t.Run("Default", func(t *testing.T) {
		tp, _ := New(Config{})
		if tp.transport == nil {
			t.Error("Expected the transport to not be nil")
		}
		if tp.transport != http.DefaultTransport {
			t.Errorf("Expected the transport to be http.DefaultTransport, got: %T", tp.transport)
		}
	})

	t.Run("Custom", func(t *testing.T) {
		tp, _ := New(Config{
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

func TestTransportConfig(t *testing.T) {
	t.Run("Defaults", func(t *testing.T) {
		tp, _ := New(Config{})

		if !reflect.DeepEqual(tp.retryOnStatus, []int{502, 503, 504}) {
			t.Errorf("Unexpected retryOnStatus: %v", tp.retryOnStatus)
		}

		if tp.disableRetry {
			t.Errorf("Unexpected disableRetry: %v", tp.disableRetry)
		}

		if tp.enableRetryOnTimeout {
			t.Errorf("Unexpected enableRetryOnTimeout: %v", tp.enableRetryOnTimeout)
		}

		if tp.maxRetries != 3 {
			t.Errorf("Unexpected maxRetries: %v", tp.maxRetries)
		}

		if tp.compressRequestBody {
			t.Errorf("Unexpected compressRequestBody: %v", tp.compressRequestBody)
		}
	})

	t.Run("Custom", func(t *testing.T) {
		tp, _ := New(Config{
			RetryOnStatus:        []int{404, 408},
			DisableRetry:         true,
			EnableRetryOnTimeout: true,
			MaxRetries:           5,
			CompressRequestBody:  true,
		})

		if !reflect.DeepEqual(tp.retryOnStatus, []int{404, 408}) {
			t.Errorf("Unexpected retryOnStatus: %v", tp.retryOnStatus)
		}

		if !tp.disableRetry {
			t.Errorf("Unexpected disableRetry: %v", tp.disableRetry)
		}

		if !tp.enableRetryOnTimeout {
			t.Errorf("Unexpected enableRetryOnTimeout: %v", tp.enableRetryOnTimeout)
		}

		if tp.maxRetries != 5 {
			t.Errorf("Unexpected maxRetries: %v", tp.maxRetries)
		}

		if !tp.compressRequestBody {
			t.Errorf("Unexpected compressRequestBody: %v", tp.compressRequestBody)
		}
	})
}

func TestTransportConnectionPool(t *testing.T) {
	t.Run("Single URL", func(t *testing.T) {
		tp, _ := New(Config{URLs: []*url.URL{{Scheme: "http", Host: "foo1"}}})

		if _, ok := tp.pool.(*singleConnectionPool); !ok {
			t.Errorf("Expected connection to be singleConnectionPool, got: %T", tp)
		}

		conn, err := tp.pool.Next()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}

		if conn.URL.String() != "http://foo1" {
			t.Errorf("Unexpected URL, want=http://foo1, got=%s", conn.URL)
		}
	})

	t.Run("Two URLs", func(t *testing.T) {
		var (
			conn *Connection
			err  error
		)

		tp, _ := New(Config{URLs: []*url.URL{
			{Scheme: "http", Host: "foo1"},
			{Scheme: "http", Host: "foo2"},
		}})

		if _, ok := tp.pool.(*statusConnectionPool); !ok {
			t.Errorf("Expected connection to be statusConnectionPool, got: %T", tp)
		}

		conn, err = tp.pool.Next()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if conn.URL.String() != "http://foo1" {
			t.Errorf("Unexpected URL, want=foo1, got=%s", conn.URL)
		}

		conn, err = tp.pool.Next()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if conn.URL.String() != "http://foo2" {
			t.Errorf("Unexpected URL, want=http://foo2, got=%s", conn.URL)
		}

		conn, err = tp.pool.Next()
		if err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
		if conn.URL.String() != "http://foo1" {
			t.Errorf("Unexpected URL, want=http://foo1, got=%s", conn.URL)
		}
	})
}

type CustomConnectionPool struct {
	urls []*url.URL
}

// Next returns a random connection.
func (cp *CustomConnectionPool) Next() (*Connection, error) {
	u := cp.urls[rand.Intn(len(cp.urls))]
	return &Connection{URL: u}, nil
}

func (cp *CustomConnectionPool) OnFailure(c *Connection) error {
	var index = -1
	for i, u := range cp.urls {
		if u == c.URL {
			index = i
		}
	}
	if index > -1 {
		cp.urls = append(cp.urls[:index], cp.urls[index+1:]...)
		return nil
	}
	return fmt.Errorf("connection not found")
}
func (cp *CustomConnectionPool) OnSuccess(c *Connection) error { return nil }
func (cp *CustomConnectionPool) URLs() []*url.URL              { return cp.urls }

func TestTransportCustomConnectionPool(t *testing.T) {
	t.Run("Run", func(t *testing.T) {
		tp, _ := New(Config{
			ConnectionPoolFunc: func(conns []*Connection, selector Selector) ConnectionPool {
				return &CustomConnectionPool{
					urls: []*url.URL{
						{Scheme: "http", Host: "custom1"},
						{Scheme: "http", Host: "custom2"},
					},
				}
			},
		})

		if _, ok := tp.pool.(*CustomConnectionPool); !ok {
			t.Fatalf("Unexpected connection pool, want=CustomConnectionPool, got=%T", tp.pool)
		}

		conn, err := tp.pool.Next()
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if conn.URL == nil {
			t.Errorf("Empty connection URL: %+v", conn)
		}
		if err := tp.pool.OnFailure(conn); err != nil {
			t.Errorf("Error removing the %q connection: %s", conn.URL, err)
		}
		if len(tp.pool.URLs()) != 1 {
			t.Errorf("Unexpected number of connections in pool: %q", tp.pool)
		}
	})
}

func TestTransportPerform(t *testing.T) {
	t.Run("Executes", func(t *testing.T) {
		u, _ := url.Parse("https://foo.com/bar")
		tp, _ := New(Config{
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
		tp, _ := New(Config{URLs: []*url.URL{u}})

		req, _ := http.NewRequest("GET", "/abc", nil)
		tp.setReqURL(u, req)

		expected := "https://foo.com/bar/abc"

		if req.URL.String() != expected {
			t.Errorf("req.URL: got=%s, want=%s", req.URL, expected)
		}
	})

	t.Run("Sets HTTP Basic Auth from URL", func(t *testing.T) {
		u, _ := url.Parse("https://foo:bar@example.com")
		tp, _ := New(Config{URLs: []*url.URL{u}})

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setReqAuth(u, req)

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
		tp, _ := New(Config{URLs: []*url.URL{u}, Username: "foo", Password: "bar"})

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setReqAuth(u, req)

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
		tp, _ := New(Config{URLs: []*url.URL{u}, APIKey: "Zm9vYmFy"}) // foobar

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setReqAuth(u, req)

		value := req.Header.Get("Authorization")
		if value == "" {
			t.Errorf("Expected the request to have the Authorization header set")
		}

		if value != "APIKey Zm9vYmFy" {
			t.Errorf(`Unexpected value for Authorization: want="APIKey Zm9vYmFy", got="%s"`, value)
		}
	})

	t.Run("Sets APIKey Authentication over ServiceToken", func(t *testing.T) {
		u, _ := url.Parse("http://example.com")
		tp, _ := New(Config{URLs: []*url.URL{u}, APIKey: "Zm9vYmFy", ServiceToken: "AAEAAWVs"}) // foobar

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setReqAuth(u, req)

		value := req.Header.Get("Authorization")
		if value == "" {
			t.Errorf("Expected the request to have the Authorization header set")
		}

		if value != "APIKey Zm9vYmFy" {
			t.Errorf(`Unexpected value for Authorization: want="APIKey Zm9vYmFy", got="%s"`, value)
		}
	})

	t.Run("Sets ServiceToken Authentication from configuration", func(t *testing.T) {
		u, _ := url.Parse("http://example.com")
		tp, _ := New(Config{URLs: []*url.URL{u}, ServiceToken: "AAEAAWVs"})

		req, _ := http.NewRequest("GET", "/", nil)
		tp.setReqAuth(u, req)

		value := req.Header.Get("Authorization")
		if value == "" {
			t.Errorf("Expected the request to have the Authorization header set")
		}

		if value != "Bearer AAEAAWVs" {
			t.Errorf(`Unexpected value for Authorization: want="Bearer AAEAAWVs", got="%s"`, value)
		}
	})

	t.Run("Sets UserAgent", func(t *testing.T) {
		u, _ := url.Parse("http://example.com")
		tp, _ := New(Config{URLs: []*url.URL{u}})

		req, _ := http.NewRequest("GET", "/abc", nil)
		tp.setReqUserAgent(req)

		if !strings.HasPrefix(req.UserAgent(), "go-elasticsearch") {
			t.Errorf("Unexpected user agent: %s", req.UserAgent())
		}
	})

	t.Run("Overwrites UserAgent", func(t *testing.T) {
		u, _ := url.Parse("http://example.com")

		tp, _ := New(Config{URLs: []*url.URL{u}, Header: http.Header{
			userAgentHeader: []string{"Elastic-Fleet-Server/7.11.1 (darwin; amd64; Go 1.16.6)"},
		}})

		req, _ := http.NewRequest("GET", "/abc", nil)
		tp.setReqUserAgent(req)

		if !strings.HasPrefix(req.UserAgent(), "Elastic-Fleet-Server") {
			t.Errorf("Unexpected user agent: %s", req.UserAgent())
		}
	})

	t.Run("Sets global HTTP request headers", func(t *testing.T) {
		hdr := http.Header{}
		hdr.Set("X-Foo", "bar")

		tp, _ := New(Config{Header: hdr})

		{
			// Set the global HTTP header
			req, _ := http.NewRequest("GET", "/abc", nil)
			tp.setReqGlobalHeader(req)

			if req.Header.Get("X-Foo") != "bar" {
				t.Errorf("Unexpected global HTTP request header value: %s", req.Header.Get("X-Foo"))
			}
		}

		{
			// Do NOT overwrite an existing request header
			req, _ := http.NewRequest("GET", "/abc", nil)
			req.Header.Set("X-Foo", "baz")
			tp.setReqGlobalHeader(req)

			if req.Header.Get("X-Foo") != "baz" {
				t.Errorf("Unexpected global HTTP request header value: %s", req.Header.Get("X-Foo"))
			}
		}
	})

	t.Run("Error No URL", func(t *testing.T) {
		tp, _ := New(Config{
			URLs: []*url.URL{},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) { return &http.Response{Status: "MOCK"}, nil },
			}})

		req, _ := http.NewRequest("GET", "/abc", nil)

		_, err := tp.Perform(req)
		if err.Error() != `cannot get connection: no connection available` {
			t.Fatalf("Expected error `cannot get URL`: but got error %q", err)
		}
	})
}

func TestTransportPerformRetries(t *testing.T) {
	t.Run("Retry request on network error and return the response", func(t *testing.T) {
		var (
			i       int
			numReqs = 2
		)

		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
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
					return nil, &mockNetError{error: fmt.Errorf("Mock network error (%d)", i)}
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

	t.Run("Retry request on EOF error and return the response", func(t *testing.T) {
		var (
			i       int
			numReqs = 2
		)

		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
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
					return nil, io.EOF
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

	t.Run("Retry request on 5xx response and return new response", func(t *testing.T) {
		var (
			i       int
			numReqs = 2
		)

		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
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

	t.Run("Close response body for a 5xx response", func(t *testing.T) {
		var (
			i       int
			numReqs = 5
		)

		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
			URLs:       []*url.URL{u, u, u},
			MaxRetries: numReqs,
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					fmt.Print(": 502\n")
					body := ioutil.NopCloser(strings.NewReader(`MOCK`))
					return &http.Response{StatusCode: 502, Body: body}, nil
				},
			}})

		req, _ := http.NewRequest("GET", "/", nil)

		res, err := tp.Perform(req)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if i != numReqs+1 {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", numReqs, i)
		}

		if res.StatusCode != 502 {
			t.Errorf("Unexpected response: %+v", res)
		}

		resBody, _ := ioutil.ReadAll(res.Body)
		res.Body.Close()

		if string(resBody) != "MOCK" {
			t.Errorf("Unexpected body, want=MOCK, got=%s", resBody)
		}
	})

	t.Run("Retry request and return error when max retries exhausted", func(t *testing.T) {
		var (
			i       int
			numReqs = 3
		)

		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
			URLs: []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					fmt.Print(": ERR\n")
					return nil, &mockNetError{error: fmt.Errorf("Mock network error (%d)", i)}
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

		// Should be initial HTTP request + 3 retries
		if i != numReqs+1 {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", numReqs, i)
		}
	})

	t.Run("Reset request body during retry", func(t *testing.T) {
		var bodies []string
		u, _ := url.Parse("https://foo.com/bar")
		tp, _ := New(Config{
			URLs: []*url.URL{u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					body, err := ioutil.ReadAll(req.Body)
					if err != nil {
						panic(err)
					}
					bodies = append(bodies, string(body))
					return &http.Response{Status: "MOCK", StatusCode: 502}, nil
				},
			}},
		)

		req, _ := http.NewRequest("POST", "/abc", strings.NewReader("FOOBAR"))
		res, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		_ = res

		if n := len(bodies); n != 4 {
			t.Fatalf("expected 4 requests, got %d", n)
		}
		for i, body := range bodies {
			if body != "FOOBAR" {
				t.Fatalf("request %d body: expected %q, got %q", i, "FOOBAR", body)
			}
		}
	})

	t.Run("Don't retry request on regular error", func(t *testing.T) {
		var i int

		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
			URLs: []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					fmt.Print(": ERR\n")
					return nil, fmt.Errorf("Mock regular error (%d)", i)
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

		if i != 1 {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", 1, i)
		}
	})

	t.Run("Don't retry request when retries are disabled", func(t *testing.T) {
		var i int

		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
			URLs: []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					fmt.Print(": ERR\n")
					return nil, &mockNetError{error: fmt.Errorf("Mock network error (%d)", i)}
				},
			},
			DisableRetry: true,
		})

		req, _ := http.NewRequest("GET", "/abc", nil)
		tp.Perform(req)

		if i != 1 {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", 1, i)
		}
	})

	t.Run("Delay the retry with a backoff function", func(t *testing.T) {
		var (
			i                int
			numReqs          = 4
			start            = time.Now()
			expectedDuration = time.Duration((numReqs-1)*100) * time.Millisecond
		)

		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
			MaxRetries: numReqs,
			URLs:       []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					fmt.Printf("Request #%d", i)
					if i == numReqs {
						fmt.Print(": OK\n")
						return &http.Response{Status: "OK"}, nil
					}
					fmt.Print(": ERR\n")
					return nil, &mockNetError{error: fmt.Errorf("Mock network error (%d)", i)}
				},
			},

			// A simple incremental backoff function
			//
			RetryBackoff: func(i int) time.Duration {
				d := time.Duration(i) * 100 * time.Millisecond
				fmt.Printf("Attempt: %d | Sleeping for %s...\n", i, d)
				return d
			},
		})

		req, _ := http.NewRequest("GET", "/abc", nil)

		res, err := tp.Perform(req)
		end := time.Since(start)

		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if res.Status != "OK" {
			t.Errorf("Unexpected response: %+v", res)
		}

		if i != numReqs {
			t.Errorf("Unexpected number of requests, want=%d, got=%d", numReqs, i)
		}

		if end < expectedDuration {
			t.Errorf("Unexpected duration, want=>%s, got=%s", expectedDuration, end)
		}
	})

	t.Run("Delay the retry with retry on timeout and context deadline", func(t *testing.T) {
		var i int
		u, _ := url.Parse("http://foo.bar")
		tp, _ := New(Config{
			EnableRetryOnTimeout: true,
			MaxRetries:           100,
			RetryBackoff:         func(i int) time.Duration { return time.Hour },
			URLs:                 []*url.URL{u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					i++
					<-req.Context().Done()
					return nil, req.Context().Err()
				},
			},
		})

		req, _ := http.NewRequest("GET", "/abc", nil)
		ctx, cancel := context.WithTimeout(req.Context(), 50*time.Millisecond)
		defer cancel()
		req = req.WithContext(ctx)

		_, err := tp.Perform(req)
		if err != context.DeadlineExceeded {
			t.Fatalf("expected context.DeadlineExceeded, got %s", err)
		}
		if i != 1 {
			t.Fatalf("unexpected number of requests: expected 1, got got %d", i)
		}
	})
}

func TestURLs(t *testing.T) {
	t.Run("Returns URLs", func(t *testing.T) {
		tp, _ := New(Config{URLs: []*url.URL{
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

func TestMetaHeader(t *testing.T) {
	t.Run("X-Elastic-Client-Meta header should be present by default", func(t *testing.T) {
		u := &url.URL{Scheme: "http", Host: "foo:9200"}
		tp, _ := New(Config{
			URLs: []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{Status: "OK"}, nil
				},
			},
		})

		req, _ := http.NewRequest("GET", "/", nil)

		tp.Perform(req)

		headerValue := req.Header.Get("X-Elastic-Client-Meta")
		fmt.Println(headerValue)
		if headerValue != initMetaHeader() {
			t.Errorf("Meta header should be present, want: %s, got : %s",
				initMetaHeader(),
				headerValue,
			)
		}
	})
	t.Run("X-Elastic-Client-Meta header should be disabled by config", func(t *testing.T) {
		u := &url.URL{Scheme: "http", Host: "foo:9200"}
		tp, _ := New(Config{
			URLs: []*url.URL{u, u, u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					return &http.Response{Status: "OK"}, nil
				},
			},
			DisableMetaHeader: true,
		})

		req, _ := http.NewRequest("GET", "/", nil)

		_, _ = tp.Perform(req)

		headerValue := req.Header.Get("X-Elastic-Client-Meta")
		if headerValue != "" {
			t.Errorf("Meta header should be empty, got: %s", headerValue)
		}
	})
}

func TestMaxRetries(t *testing.T) {
	tests := []struct {
		name              string
		maxRetries        int
		disableRetry      bool
		expectedCallCount int
	}{
		{
			name:              "MaxRetries Active set to default",
			disableRetry:      false,
			expectedCallCount: 4,
		},
		{
			name:              "MaxRetries Active set to 1",
			maxRetries:        1,
			disableRetry:      false,
			expectedCallCount: 2,
		},
		{
			name:              "Max Retries Active set to 2",
			maxRetries:        2,
			disableRetry:      false,
			expectedCallCount: 3,
		},
		{
			name:              "Max Retries Active set to 3",
			maxRetries:        3,
			disableRetry:      false,
			expectedCallCount: 4,
		},
		{
			name:              "MaxRetries Inactive set to 0",
			maxRetries:        0,
			disableRetry:      true,
			expectedCallCount: 1,
		},
		{
			name:              "MaxRetries Inactive set to 3",
			maxRetries:        3,
			disableRetry:      true,
			expectedCallCount: 1,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var callCount int
			c, _ := New(Config{
				URLs: []*url.URL{{}},
				Transport: &mockTransp{
					RoundTripFunc: func(req *http.Request) (*http.Response, error) {
						callCount++
						return &http.Response{
							StatusCode: http.StatusBadGateway,
							Status:     "MOCK",
						}, nil
					},
				},
				MaxRetries:   test.maxRetries,
				DisableRetry: test.disableRetry,
			})

			c.Perform(&http.Request{URL: &url.URL{}, Header: make(http.Header)}) // errcheck ignore

			if test.expectedCallCount != callCount {
				t.Errorf("Bad retry call count, got : %d, want : %d", callCount, test.expectedCallCount)
			}
		})
	}
}

func TestCompatibilityHeader(t *testing.T) {
	tests := []struct {
		name                string
		compatibilityHeader bool
		bodyPresent         bool
		expectsHeader       []string
	}{
		{
			name:                "Compatibility header disabled",
			compatibilityHeader: false,
			bodyPresent:         false,
			expectsHeader:       []string{"application/json"},
		},
		{
			name:                "Compatibility header enabled",
			compatibilityHeader: true,
			bodyPresent:         false,
			expectsHeader:       []string{"application/vnd.elasticsearch+json;compatible-with=7"},
		},
		{
			name:                "Compatibility header enabled with body",
			compatibilityHeader: true,
			bodyPresent:         true,
			expectsHeader:       []string{"application/vnd.elasticsearch+json;compatible-with=7"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compatibilityHeader = test.compatibilityHeader

			c, _ := New(Config{
				URLs: []*url.URL{{}},
				Transport: &mockTransp{
					RoundTripFunc: func(req *http.Request) (*http.Response, error) {
						if test.compatibilityHeader {
							if !reflect.DeepEqual(req.Header["Accept"], test.expectsHeader) {
								t.Errorf("Compatibility header enabled but header is, not in request headers, got: %s, want: %s", req.Header["Accept"], test.expectsHeader)
							}
						}
						if test.bodyPresent {
							if !reflect.DeepEqual(req.Header["Content-Type"], test.expectsHeader) {
								t.Errorf("Compatibility header with Body enabled, not in request headers, got: %s, want: %s", req.Header["Content-Type"], test.expectsHeader)
							}
						}

						return &http.Response{
							StatusCode: http.StatusOK,
							Status:     "MOCK",
						}, nil
					},
				},
			})

			req := &http.Request{URL: &url.URL{}, Header: make(http.Header)}
			if test.bodyPresent {
				req.Body = ioutil.NopCloser(strings.NewReader("{}"))
			}

			_, _ = c.Perform(req)

			compatibilityHeader = false
		})
	}
}

func TestRequestCompression(t *testing.T) {

	tests := []struct {
		name            string
		compressionFlag bool
		inputBody       string
	}{
		{
			name:            "Uncompressed",
			compressionFlag: false,
			inputBody:       "elasticsearch",
		},
		{
			name:            "Compressed",
			compressionFlag: true,
			inputBody:       "elasticsearch",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tp, _ := New(Config{
				URLs:                []*url.URL{{}},
				CompressRequestBody: test.compressionFlag,
				Transport: &mockTransp{
					RoundTripFunc: func(req *http.Request) (*http.Response, error) {
						if req.Body == nil || req.Body == http.NoBody {
							return nil, fmt.Errorf("unexpected body: %v", req.Body)
						}

						var buf bytes.Buffer
						buf.ReadFrom(req.Body)

						if req.ContentLength != int64(buf.Len()) {
							return nil, fmt.Errorf("mismatched Content-Length: %d vs actual %d", req.ContentLength, buf.Len())
						}

						if test.compressionFlag {
							var unBuf bytes.Buffer
							zr, err := gzip.NewReader(&buf)
							if err != nil {
								return nil, fmt.Errorf("decompression error: %v", err)
							}
							unBuf.ReadFrom(zr)
							buf = unBuf
						}

						if buf.String() != test.inputBody {
							return nil, fmt.Errorf("unexpected body: %s", buf.String())
						}

						return &http.Response{Status: "MOCK"}, nil
					},
				},
			})

			req, _ := http.NewRequest("POST", "/abc", bytes.NewBufferString(test.inputBody))

			res, err := tp.Perform(req)
			if err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}

			if res.Status != "MOCK" {
				t.Errorf("Unexpected response: %+v", res)
			}
		})
	}
}

func TestTransportPerformAndReadMetricsResponses(t *testing.T) {
	t.Run("Read Metrics.Responses", func(t *testing.T) {
		u, _ := url.Parse("https://foo.com/bar")
		tp, _ := New(Config{
			EnableMetrics: true,
			URLs:          []*url.URL{u},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) { return &http.Response{Status: "MOCK"}, nil },
			}})

		ch := make(chan struct{})
		go func() {
			for {
				select {
				case <-ch:
					break
				default:
					metrics, _ := tp.Metrics()
					for range metrics.Responses {
					}
				}
			}
		}()

		for i := 0; i < 100000; i++ {
			req, _ := http.NewRequest("GET", "/abc", nil)
			_, _ = tp.Perform(req)
		}

		ch <- struct{}{}
		close(ch)
	})
}
