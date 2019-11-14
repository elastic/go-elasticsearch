// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package estransport

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	_ = fmt.Print
	_ = os.Stdout
)

func TestTransportLogger(t *testing.T) {
	newRoundTripper := func() http.RoundTripper {
		return &mockTransp{
			RoundTripFunc: func(req *http.Request) (*http.Response, error) {
				return &http.Response{
					Status:        "200 OK",
					StatusCode:    200,
					ContentLength: 13,
					Header:        http.Header(map[string][]string{"Content-Type": {"application/json"}}),
					Body:          ioutil.NopCloser(strings.NewReader(`{"foo":"bar"}`)),
				}, nil
			},
		}
	}

	t.Run("Defaults", func(t *testing.T) {
		var wg sync.WaitGroup

		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			// Logger: ioutil.Discard,
		})

		for i := 0; i < 100; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				req, _ := http.NewRequest("GET", "/abc", nil)
				_, err := tp.Perform(req)
				if err != nil {
					t.Fatalf("Unexpected error: %s", err)
				}
			}()
		}
		wg.Wait()
	})

	t.Run("Nil", func(t *testing.T) {
		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			Logger:    nil,
		})

		req, _ := http.NewRequest("GET", "/abc", nil)
		_, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
	})

	t.Run("No HTTP response", func(t *testing.T) {
		tp := New(Config{
			URLs: []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("Mock error")
				},
			},
			Logger: &TextLogger{Output: ioutil.Discard},
		})

		req, _ := http.NewRequest("GET", "/abc", nil)
		res, err := tp.Perform(req)
		if err == nil {
			t.Errorf("Expected error: %v", err)
		}
		if res != nil {
			t.Errorf("Expected nil response, got: %v", err)
		}
	})

	t.Run("Keep response body", func(t *testing.T) {
		var dst strings.Builder

		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			Logger:    &TextLogger{Output: &dst, EnableRequestBody: true, EnableResponseBody: true},
		})

		req, _ := http.NewRequest("GET", "/abc?q=a,b", nil)
		req.Body = ioutil.NopCloser(strings.NewReader(`{"query":"42"}`))

		res, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %s", err)
		}

		if len(dst.String()) < 1 {
			t.Errorf("Log is empty: %#v", dst.String())
		}

		if len(body) < 1 {
			t.Fatalf("Body is empty: %#v", body)
		}
	})

	t.Run("Text with body", func(t *testing.T) {
		var dst strings.Builder

		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			Logger:    &TextLogger{Output: &dst, EnableRequestBody: true, EnableResponseBody: true},
		})

		req, _ := http.NewRequest("GET", "/abc?q=a,b", nil)
		req.Body = ioutil.NopCloser(strings.NewReader(`{"query":"42"}`))

		res, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %s", err)
		}

		output := dst.String()
		output = strings.TrimSuffix(output, "\n")
		// fmt.Println(output)

		lines := strings.Split(output, "\n")

		if len(lines) != 3 {
			t.Fatalf("Expected 3 lines, got %d", len(lines))
		}

		if !strings.Contains(lines[0], "GET http://foo/abc?q=a,b") {
			t.Errorf("Unexpected output: %s", lines[0])
		}

		if lines[1] != `> {"query":"42"}` {
			t.Errorf("Unexpected output: %s", lines[1])
		}

		if lines[2] != `< {"foo":"bar"}` {
			t.Errorf("Unexpected output: %s", lines[1])
		}
	})

	t.Run("Color with body", func(t *testing.T) {
		var dst strings.Builder

		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			Logger:    &ColorLogger{Output: &dst, EnableRequestBody: true, EnableResponseBody: true},
		})

		req, _ := http.NewRequest("GET", "/abc?q=a,b", nil)
		req.Body = ioutil.NopCloser(strings.NewReader(`{"query":"42"}`))

		res, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %s", err)
		}

		var output string
		stripANSI := regexp.MustCompile("(?sm)\x1b\\[.+?m([^\x1b]+?)|\x1b\\[0m")
		for _, v := range strings.Split(dst.String(), "\n") {
			if v != "" {
				output += stripANSI.ReplaceAllString(v, "$1")
				if !strings.HasSuffix(output, "\n") {
					output += "\n"
				}
			}
		}
		output = strings.TrimSuffix(output, "\n")
		// fmt.Println(output)

		lines := strings.Split(output, "\n")

		if len(lines) != 4 {
			t.Fatalf("Expected 4 lines, got %d", len(lines))
		}

		if !strings.Contains(lines[0], "GET http://foo/abc?q=a,b") {
			t.Errorf("Unexpected output: %s", lines[0])
		}

		if !strings.Contains(lines[1], `» {"query":"42"}`) {
			t.Errorf("Unexpected output: %s", lines[1])
		}

		if !strings.Contains(lines[2], `« {"foo":"bar"}`) {
			t.Errorf("Unexpected output: %s", lines[2])
		}
	})

	t.Run("Curl", func(t *testing.T) {
		var dst strings.Builder

		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			Logger:    &CurlLogger{Output: &dst, EnableRequestBody: true, EnableResponseBody: true},
		})

		req, _ := http.NewRequest("GET", "/abc?q=a,b", nil)
		req.Body = ioutil.NopCloser(strings.NewReader(`{"query":"42"}`))

		res, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %s", err)
		}

		output := dst.String()
		output = strings.TrimSuffix(output, "\n")
		// fmt.Println(output)

		lines := strings.Split(output, "\n")

		if len(lines) != 9 {
			t.Fatalf("Expected 9 lines, got %d", len(lines))
		}

		if !strings.Contains(lines[0], "curl -X GET 'http://localhost:9200/abc?pretty&q=a%2Cb'") {
			t.Errorf("Unexpected output: %s", lines[0])
		}
	})

	t.Run("JSON", func(t *testing.T) {
		var dst strings.Builder

		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			Logger:    &JSONLogger{Output: &dst},
		})

		req, _ := http.NewRequest("GET", "/abc?q=a,b", nil)
		req.Body = ioutil.NopCloser(strings.NewReader(`{"query":"42"}`))
		_, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		output := dst.String()
		output = strings.TrimSuffix(output, "\n")
		// fmt.Println(output)

		lines := strings.Split(output, "\n")

		if len(lines) != 1 {
			t.Fatalf("Expected 1 line, got %d", len(lines))
		}

		var j map[string]interface{}
		if err := json.Unmarshal([]byte(output), &j); err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}

		domain := j["url"].(map[string]interface{})["domain"]
		if domain != "foo" {
			t.Errorf("Unexpected JSON output: %s", domain)
		}
	})

	t.Run("JSON with request body", func(t *testing.T) {
		var dst strings.Builder

		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			Logger:    &JSONLogger{Output: &dst, EnableRequestBody: true},
		})

		req, _ := http.NewRequest("GET", "/abc?q=a,b", nil)
		req.Body = ioutil.NopCloser(strings.NewReader(`{"query":"42"}`))

		res, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		_, err = ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatalf("Error reading response body: %s", err)
		}

		output := dst.String()
		output = strings.TrimSuffix(output, "\n")
		// fmt.Println(output)

		lines := strings.Split(output, "\n")

		if len(lines) != 1 {
			t.Fatalf("Expected 1 line, got %d", len(lines))
		}

		var j map[string]interface{}
		if err := json.Unmarshal([]byte(output), &j); err != nil {
			t.Errorf("Error decoding JSON: %s", err)
		}

		body := j["http"].(map[string]interface{})["request"].(map[string]interface{})["body"].(string)
		if !strings.Contains(body, "query") {
			t.Errorf("Unexpected JSON output: %s", body)
		}
	})

	t.Run("Custom", func(t *testing.T) {
		var dst strings.Builder

		tp := New(Config{
			URLs:      []*url.URL{{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			Logger:    &CustomLogger{Output: &dst},
		})

		req, _ := http.NewRequest("GET", "/abc?q=a,b", nil)
		req.Body = ioutil.NopCloser(strings.NewReader(`{"query":"42"}`))

		_, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if !strings.HasPrefix(dst.String(), "GET http://foo/abc?q=a,b") {
			t.Errorf("Unexpected output: %s", dst.String())
		}
	})

	t.Run("Duplicate body", func(t *testing.T) {
		input := ResponseBody{content: strings.NewReader("FOOBAR")}

		b1, b2, err := duplicateBody(&input)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if !input.closed {
			t.Errorf("Expected input to be closed: %#v", input)
		}

		read, _ := ioutil.ReadAll(&input)
		if len(read) > 0 {
			t.Errorf("Expected input to be drained: %#v", input.content)
		}

		b1r, _ := ioutil.ReadAll(b1)
		b2r, _ := ioutil.ReadAll(b2)
		if len(b1r) != 6 || len(b2r) != 6 {
			t.Errorf(
				"Unexpected duplicate content, b1=%q (%db), b2=%q (%db)",
				string(b1r), len(b1r), string(b2r), len(b2r),
			)
		}
	})

	t.Run("Duplicate body with error", func(t *testing.T) {
		input := ResponseBody{content: &ErrorReader{r: strings.NewReader("FOOBAR")}}

		b1, b2, err := duplicateBody(&input)
		if err == nil {
			t.Errorf("Expected error, got: %v", err)
		}
		if err.Error() != "MOCK ERROR" {
			t.Errorf("Unexpected error value, expected [ERROR MOCK], got [%s]", err.Error())
		}

		read, _ := ioutil.ReadAll(&input)
		if string(read) != "BAR" {
			t.Errorf("Unexpected undrained part: %q", read)
		}

		b2r, _ := ioutil.ReadAll(b2)
		if string(b2r) != "FOO" {
			t.Errorf("Unexpected value, b2=%q", string(b2r))
		}

		b1c, err := ioutil.ReadAll(b1)
		if string(b1c) != "FOO" {
			t.Errorf("Unexpected value, b1=%q", string(b1c))
		}
		if err == nil {
			t.Errorf("Expected error when reading b1, got: %v", err)
		}
		if err.Error() != "MOCK ERROR" {
			t.Errorf("Unexpected error value, expected [ERROR MOCK], got [%s]", err.Error())
		}
	})
}

func TestDebuggingLogger(t *testing.T) {
	logger := &debuggingLogger{Output: ioutil.Discard}

	t.Run("Log", func(t *testing.T) {
		if err := logger.Log("Foo"); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})
	t.Run("Logf", func(t *testing.T) {
		if err := logger.Logf("Foo %d", 1); err != nil {
			t.Errorf("Unexpected error: %s", err)
		}
	})
}

type CustomLogger struct {
	Output io.Writer
}

func (l *CustomLogger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) error {
	fmt.Fprintln(l.Output, req.Method, req.URL, "->", res.Status)
	return nil
}

func (l *CustomLogger) RequestBodyEnabled() bool  { return false }
func (l *CustomLogger) ResponseBodyEnabled() bool { return false }

type ResponseBody struct {
	content io.Reader
	closed  bool
}

func (r *ResponseBody) Read(p []byte) (int, error) {
	return r.content.Read(p)
}

func (r *ResponseBody) Close() error {
	r.closed = true
	return nil
}

type ErrorReader struct {
	r io.Reader
}

func (r *ErrorReader) Read(p []byte) (int, error) {
	lr := io.LimitReader(r.r, 3)
	c, _ := lr.Read(p)
	return c, errors.New("MOCK ERROR")
}
