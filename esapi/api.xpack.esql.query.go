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

func newEsqlQueryFunc(t Transport) EsqlQuery {
	return func(body io.Reader, o ...func(*EsqlQueryRequest)) (*Response, error) {
		var r = EsqlQueryRequest{Body: body}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// EsqlQuery - Executes an ESQL request
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/esql-query-api.html.
type EsqlQuery func(body io.Reader, o ...func(*EsqlQueryRequest)) (*Response, error)

// EsqlQueryRequest configures the Esql Query API request.
type EsqlQueryRequest struct {
	Body io.Reader

	Delimiter string
	Format    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r EsqlQueryRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(7 + len("/_query"))
	path.WriteString("http://")
	path.WriteString("/_query")

	params = make(map[string]string)

	if r.Delimiter != "" {
		params["delimiter"] = r.Delimiter
	}

	if r.Format != "" {
		params["format"] = r.Format
	}

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
func (f EsqlQuery) WithContext(v context.Context) func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		r.ctx = v
	}
}

// WithDelimiter - the character to use between values within a csv row. only valid for the csv format..
func (f EsqlQuery) WithDelimiter(v string) func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		r.Delimiter = v
	}
}

// WithFormat - a short version of the accept header, e.g. json, yaml.
func (f EsqlQuery) WithFormat(v string) func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		r.Format = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f EsqlQuery) WithPretty() func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f EsqlQuery) WithHuman() func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f EsqlQuery) WithErrorTrace() func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f EsqlQuery) WithFilterPath(v ...string) func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f EsqlQuery) WithHeader(h map[string]string) func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f EsqlQuery) WithOpaqueID(s string) func(*EsqlQueryRequest) {
	return func(r *EsqlQueryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
