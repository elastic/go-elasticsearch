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

func newSearchApplicationRenderQueryFunc(t Transport) SearchApplicationRenderQuery {
	return func(name string, o ...func(*SearchApplicationRenderQueryRequest)) (*Response, error) {
		var r = SearchApplicationRenderQueryRequest{Name: name}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SearchApplicationRenderQuery renders a query for given search application search parameters
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/search-application-render-query.html.
type SearchApplicationRenderQuery func(name string, o ...func(*SearchApplicationRenderQueryRequest)) (*Response, error)

// SearchApplicationRenderQueryRequest configures the Search Application Render Query API request.
type SearchApplicationRenderQueryRequest struct {
	Body io.Reader

	Name string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r SearchApplicationRenderQueryRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(7 + 1 + len("_application") + 1 + len("search_application") + 1 + len(r.Name) + 1 + len("_render_query"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_application")
	path.WriteString("/")
	path.WriteString("search_application")
	path.WriteString("/")
	path.WriteString(r.Name)
	path.WriteString("/")
	path.WriteString("_render_query")

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
func (f SearchApplicationRenderQuery) WithContext(v context.Context) func(*SearchApplicationRenderQueryRequest) {
	return func(r *SearchApplicationRenderQueryRequest) {
		r.ctx = v
	}
}

// WithBody - Search parameters, which will override any default search parameters defined in the search application template.
func (f SearchApplicationRenderQuery) WithBody(v io.Reader) func(*SearchApplicationRenderQueryRequest) {
	return func(r *SearchApplicationRenderQueryRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SearchApplicationRenderQuery) WithPretty() func(*SearchApplicationRenderQueryRequest) {
	return func(r *SearchApplicationRenderQueryRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SearchApplicationRenderQuery) WithHuman() func(*SearchApplicationRenderQueryRequest) {
	return func(r *SearchApplicationRenderQueryRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SearchApplicationRenderQuery) WithErrorTrace() func(*SearchApplicationRenderQueryRequest) {
	return func(r *SearchApplicationRenderQueryRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SearchApplicationRenderQuery) WithFilterPath(v ...string) func(*SearchApplicationRenderQueryRequest) {
	return func(r *SearchApplicationRenderQueryRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SearchApplicationRenderQuery) WithHeader(h map[string]string) func(*SearchApplicationRenderQueryRequest) {
	return func(r *SearchApplicationRenderQueryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SearchApplicationRenderQuery) WithOpaqueID(s string) func(*SearchApplicationRenderQueryRequest) {
	return func(r *SearchApplicationRenderQueryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
