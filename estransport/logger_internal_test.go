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
			URLs:      []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			LogOutput: ioutil.Discard,
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
			URLs:      []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			LogOutput: nil,
		})

		req, _ := http.NewRequest("GET", "/abc", nil)
		_, err := tp.Perform(req)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
	})

	t.Run("No HTTP response", func(t *testing.T) {
		tp := New(Config{
			URLs: []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport: &mockTransp{
				RoundTripFunc: func(req *http.Request) (*http.Response, error) {
					return nil, errors.New("Mock error")
				},
			},
			LogOutput: ioutil.Discard,
		})

		req, _ := http.NewRequest("GET", "/abc", nil)
		_, err := tp.Perform(req)
		if err == nil {
			t.Fatalf("Expected error: %s", err)
		}
	})

	t.Run("Text with body", func(t *testing.T) {
		var dst strings.Builder

		tp := New(Config{
			URLs:            []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport:       newRoundTripper(),
			LogOutput:       &dst,
			LogFormat:       LogFormatText,
			LogRequestBody:  true,
			LogResponseBody: true,
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
			URLs:            []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport:       newRoundTripper(),
			LogOutput:       &dst,
			LogFormat:       LogFormatColor,
			LogRequestBody:  true,
			LogResponseBody: true,
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
			URLs:            []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport:       newRoundTripper(),
			LogOutput:       &dst,
			LogFormat:       LogFormatCurl,
			LogResponseBody: true,
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
			URLs:      []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
			LogOutput: &dst,
			LogFormat: LogFormatJSON,
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
			URLs:           []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport:      newRoundTripper(),
			LogOutput:      &dst,
			LogFormat:      LogFormatJSON,
			LogRequestBody: true,
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
			URLs:      []*url.URL{&url.URL{Scheme: "http", Host: "foo"}},
			Transport: newRoundTripper(),
		})
		tp.logger = &CustomLogger{output: &dst}

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
}

type CustomLogger struct {
	output io.Writer
}

func (l *CustomLogger) LogRoundTrip(
	req *http.Request,
	res *http.Response,
	err error,
	start time.Time,
	dur time.Duration,
) error {
	fmt.Fprintln(l.output, req.Method, req.URL, "->", res.Status)
	return nil
}

func (l *CustomLogger) RequestBodyEnabled() bool  { return false }
func (l *CustomLogger) ResponseBodyEnabled() bool { return false }
