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
// Code generated from specification version 8.18.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func newAsyncSearchStatusFunc(t Transport) AsyncSearchStatus {
	return func(id string, o ...func(*AsyncSearchStatusRequest)) (*Response, error) {
		var r = AsyncSearchStatusRequest{DocumentID: id}
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

// AsyncSearchStatus - Retrieves the status of a previously submitted async search request given its ID.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/async-search.html.
type AsyncSearchStatus func(id string, o ...func(*AsyncSearchStatusRequest)) (*Response, error)

// AsyncSearchStatusRequest configures the Async Search Status API request.
type AsyncSearchStatusRequest struct {
	DocumentID string

	KeepAlive time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r AsyncSearchStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "async_search.status")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_async_search") + 1 + len("status") + 1 + len(r.DocumentID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_async_search")
	path.WriteString("/")
	path.WriteString("status")
	path.WriteString("/")
	path.WriteString(r.DocumentID)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "id", r.DocumentID)
	}

	params = make(map[string]string)

	if r.KeepAlive != 0 {
		params["keep_alive"] = formatDuration(r.KeepAlive)
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

	req, err := newRequest(method, path.String(), nil)
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

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "async_search.status")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "async_search.status")
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
func (f AsyncSearchStatus) WithContext(v context.Context) func(*AsyncSearchStatusRequest) {
	return func(r *AsyncSearchStatusRequest) {
		r.ctx = v
	}
}

// WithKeepAlive - specify the time interval in which the results (partial or final) for this search will be available.
func (f AsyncSearchStatus) WithKeepAlive(v time.Duration) func(*AsyncSearchStatusRequest) {
	return func(r *AsyncSearchStatusRequest) {
		r.KeepAlive = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f AsyncSearchStatus) WithPretty() func(*AsyncSearchStatusRequest) {
	return func(r *AsyncSearchStatusRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f AsyncSearchStatus) WithHuman() func(*AsyncSearchStatusRequest) {
	return func(r *AsyncSearchStatusRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f AsyncSearchStatus) WithErrorTrace() func(*AsyncSearchStatusRequest) {
	return func(r *AsyncSearchStatusRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f AsyncSearchStatus) WithFilterPath(v ...string) func(*AsyncSearchStatusRequest) {
	return func(r *AsyncSearchStatusRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f AsyncSearchStatus) WithHeader(h map[string]string) func(*AsyncSearchStatusRequest) {
	return func(r *AsyncSearchStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f AsyncSearchStatus) WithOpaqueID(s string) func(*AsyncSearchStatusRequest) {
	return func(r *AsyncSearchStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
