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
// Code generated from specification version 8.12.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

func newIndicesUpdateAliasesFunc(t Transport) IndicesUpdateAliases {
	return func(body io.Reader, o ...func(*IndicesUpdateAliasesRequest)) (*Response, error) {
		var r = IndicesUpdateAliasesRequest{Body: body}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IndicesUpdateAliases updates index aliases.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-aliases.html.
type IndicesUpdateAliases func(body io.Reader, o ...func(*IndicesUpdateAliasesRequest)) (*Response, error)

// IndicesUpdateAliasesRequest configures the Indices Update Aliases API request.
type IndicesUpdateAliasesRequest struct {
	Body io.Reader

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IndicesUpdateAliasesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.update_aliases")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_aliases"))
	path.WriteString("http://")
	path.WriteString("/_aliases")

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
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
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
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

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "indices.update_aliases")
		if reader := instrument.RecordRequestBody(ctx, "indices.update_aliases", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.update_aliases")
	}
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
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
func (f IndicesUpdateAliases) WithContext(v context.Context) func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
func (f IndicesUpdateAliases) WithMasterTimeout(v time.Duration) func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - request timeout.
func (f IndicesUpdateAliases) WithTimeout(v time.Duration) func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IndicesUpdateAliases) WithPretty() func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IndicesUpdateAliases) WithHuman() func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IndicesUpdateAliases) WithErrorTrace() func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IndicesUpdateAliases) WithFilterPath(v ...string) func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IndicesUpdateAliases) WithHeader(h map[string]string) func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IndicesUpdateAliases) WithOpaqueID(s string) func(*IndicesUpdateAliasesRequest) {
	return func(r *IndicesUpdateAliasesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
