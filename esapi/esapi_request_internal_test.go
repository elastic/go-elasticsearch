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

		req, err = newRequest("GET", "http:////_aliases/test", nil)
		if err != nil {
			t.Fatalf("Unexpected error: %s", err)
		}
		if req.URL.String() != "http:////_aliases/test" {
			t.Errorf("Unexpected url for request: %s", req.URL.String())
		}
		if req.URL.Host != "" {
			t.Errorf("Unexpected host, should be empty, got: %s", req.URL.Host)
		}
	})
}
