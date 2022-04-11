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

//go:build !integration
// +build !integration

package esapi

import (
	"errors"
	"io/ioutil"
	"net/http"
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

	t.Run("Warnings", func(t *testing.T) {
		hdr := http.Header{}
		hdr.Add("Warning", "Foo 1")
		hdr.Add("Warning", "Foo 2")
		res = &Response{StatusCode: 201, Header: hdr}

		if !res.HasWarnings() {
			t.Errorf("Expected response to have warnings")
		}

		if len(res.Warnings()) != 2 {
			t.Errorf("Expected [2] warnings, got: %d", len(res.Warnings()))
		}
	})
}
