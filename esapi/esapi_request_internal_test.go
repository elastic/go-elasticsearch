// +build !integration

package esapi

import (
	"bytes"
	"io"
	"net/http"
	"strings"
	"testing"
)

func TestAPIRequest(t *testing.T) {
	var (
		body string
		req  *http.Request
		err  error
	)

	t.Run("newRequest", func(t *testing.T) {
		req, err = newRequest("GET", "/foo", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if req.Method != "GET" {
			t.Errorf("Unexpected method %s, want GET", req.Method)
		}
		if req.URL.String() != "/foo" {
			t.Errorf("Unexpected URL %s, want /foo", req.URL)
		}

		body = `{"foo":"bar"}`
		req, err = newRequest("GET", "/foo", strings.NewReader(body))
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}

		if _, ok := req.Body.(io.ReadCloser); !ok {
			t.Errorf("Unexpected type for req.Body: %T", req.Body)
		}
		if int(req.ContentLength) != len(body) {
			t.Errorf("Unexpected length of req.Body, got=%d, want=%d", req.ContentLength, len(body))
		}

		req, err = newRequest("GET", "/foo", bytes.NewBuffer([]byte(body)))
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if _, ok := req.Body.(io.ReadCloser); !ok {
			t.Errorf("Unexpected type for req.Body: %T", req.Body)
		}
		req, err = newRequest("GET", "/foo", bytes.NewReader([]byte(body)))
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if _, ok := req.Body.(io.ReadCloser); !ok {
			t.Errorf("Unexpected type for req.Body: %T", req.Body)
		}
	})
}
