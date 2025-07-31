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
// Code generated from specification version 8.19.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newStreamsStatusFunc(t Transport) StreamsStatus {
	return func(o ...func(*StreamsStatusRequest)) (*Response, error) {
		var r = StreamsStatusRequest{}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.Instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// StreamsStatus return the current status of the streams feature for each streams type
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/streams-status.html.
type StreamsStatus func(o ...func(*StreamsStatusRequest)) (*Response, error)

// StreamsStatusRequest configures the Streams Status API request.
type StreamsStatusRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r StreamsStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "streams.status")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_streams/status"))
	path.WriteString("http://")
	path.WriteString("/_streams/status")

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

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
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

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "streams.status")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "streams.status")
	}
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
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
func (f StreamsStatus) WithContext(v context.Context) func(*StreamsStatusRequest) {
	return func(r *StreamsStatusRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f StreamsStatus) WithPretty() func(*StreamsStatusRequest) {
	return func(r *StreamsStatusRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f StreamsStatus) WithHuman() func(*StreamsStatusRequest) {
	return func(r *StreamsStatusRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f StreamsStatus) WithErrorTrace() func(*StreamsStatusRequest) {
	return func(r *StreamsStatusRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f StreamsStatus) WithFilterPath(v ...string) func(*StreamsStatusRequest) {
	return func(r *StreamsStatusRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f StreamsStatus) WithHeader(h map[string]string) func(*StreamsStatusRequest) {
	return func(r *StreamsStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f StreamsStatus) WithOpaqueID(s string) func(*StreamsStatusRequest) {
	return func(r *StreamsStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
