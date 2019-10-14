// Licensed to Elasticsearch B.V. under one or more agreements.
// Elasticsearch B.V. licenses this file to you under the Apache 2.0 License.
// See the LICENSE file in the project root for more information.

// +build !integration

package esapi

import (
	"errors"
	"io/ioutil"
	"strings"
	"testing"
)

type errReader struct{}

func (errReader) Read(p []byte) (n int, err error) { return 1, errors.New("MOCK ERROR") }

func TestAPIResponse(t *testing.T) {
	var (
		body string
		res  *Response
	)

	t.Run("String", func(t *testing.T) {
		body = `{"foo":"bar"}`

		res = &Response{StatusCode: 200, Body: ioutil.NopCloser(strings.NewReader(body))}

		expected := `[200 OK]` + ` ` + body
		if res.String() != expected {
			t.Errorf("Unexpected response: %s, want: %s", res.String(), expected)
		}
	})

	t.Run("String with empty response", func(t *testing.T) {
		res = &Response{}

		if res.String() != "[0 ]" {
			t.Errorf("Unexpected response: %s", res.String())
		}
	})

	t.Run("String with nil response", func(t *testing.T) {
		res = nil

		if res.String() != "[0 <nil>]" {
			t.Errorf("Unexpected response: %s", res.String())
		}
	})

	t.Run("String Error", func(t *testing.T) {
		res = &Response{StatusCode: 200, Body: ioutil.NopCloser(errReader{})}

		t.Log(res.String())
		t.Log(res.String())

		if !strings.Contains(res.String(), `error reading response`) {
			t.Errorf("Expected response string to contain 'error reading response', got: %s", res.String())
		}
	})

	t.Run("Status", func(t *testing.T) {
		res = &Response{StatusCode: 404}

		if res.Status() != `404 Not Found` {
			t.Errorf("Unexpected response status text: %s, want: 404 Not Found", res.Status())
		}
	})

	t.Run("IsError", func(t *testing.T) {
		res = &Response{StatusCode: 201}

		if res.IsError() {
			t.Errorf("Unexpected error for response: %s", res.Status())
		}

		res = &Response{StatusCode: 403}

		if !res.IsError() {
			t.Errorf("Expected error for response: %s", res.Status())
		}
	})
}
