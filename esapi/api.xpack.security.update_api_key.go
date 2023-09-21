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
//
// Code generated from specification version 8.11.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newSecurityUpdateAPIKeyFunc(t Transport) SecurityUpdateAPIKey {
	return func(id string, o ...func(*SecurityUpdateAPIKeyRequest)) (*Response, error) {
		var r = SecurityUpdateAPIKeyRequest{DocumentID: id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SecurityUpdateAPIKey - Updates attributes of an existing API key.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-api-key.html.
type SecurityUpdateAPIKey func(id string, o ...func(*SecurityUpdateAPIKeyRequest)) (*Response, error)

// SecurityUpdateAPIKeyRequest configures the Security UpdateAPI Key API request.
type SecurityUpdateAPIKeyRequest struct {
	DocumentID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r SecurityUpdateAPIKeyRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(7 + 1 + len("_security") + 1 + len("api_key") + 1 + len(r.DocumentID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("api_key")
	path.WriteString("/")
	path.WriteString(r.DocumentID)

	params = make(map[string]string)

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), r.Body)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f SecurityUpdateAPIKey) WithContext(v context.Context) func(*SecurityUpdateAPIKeyRequest) {
	return func(r *SecurityUpdateAPIKeyRequest) {
		r.ctx = v
	}
}

// WithBody - The API key request to update attributes of an API key..
func (f SecurityUpdateAPIKey) WithBody(v io.Reader) func(*SecurityUpdateAPIKeyRequest) {
	return func(r *SecurityUpdateAPIKeyRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityUpdateAPIKey) WithPretty() func(*SecurityUpdateAPIKeyRequest) {
	return func(r *SecurityUpdateAPIKeyRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityUpdateAPIKey) WithHuman() func(*SecurityUpdateAPIKeyRequest) {
	return func(r *SecurityUpdateAPIKeyRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityUpdateAPIKey) WithErrorTrace() func(*SecurityUpdateAPIKeyRequest) {
	return func(r *SecurityUpdateAPIKeyRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityUpdateAPIKey) WithFilterPath(v ...string) func(*SecurityUpdateAPIKeyRequest) {
	return func(r *SecurityUpdateAPIKeyRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityUpdateAPIKey) WithHeader(h map[string]string) func(*SecurityUpdateAPIKeyRequest) {
	return func(r *SecurityUpdateAPIKeyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityUpdateAPIKey) WithOpaqueID(s string) func(*SecurityUpdateAPIKeyRequest) {
	return func(r *SecurityUpdateAPIKeyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
