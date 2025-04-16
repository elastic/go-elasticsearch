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
	"time"
)

func newSlmStopFunc(t Transport) SlmStop {
	return func(o ...func(*SlmStopRequest)) (*Response, error) {
		var r = SlmStopRequest{}
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

// SlmStop - Turns off snapshot lifecycle management (SLM).
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/slm-api-stop.html.
type SlmStop func(o ...func(*SlmStopRequest)) (*Response, error)

// SlmStopRequest configures the Slm Stop API request.
type SlmStopRequest struct {
	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SlmStopRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "slm.stop")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_slm/stop"))
	path.WriteString("http://")
	path.WriteString("/_slm/stop")

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
		instrument.BeforeRequest(req, "slm.stop")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "slm.stop")
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
func (f SlmStop) WithContext(v context.Context) func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - timeout for processing on master node.
func (f SlmStop) WithMasterTimeout(v time.Duration) func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - timeout for acknowledgement of update from all nodes in cluster.
func (f SlmStop) WithTimeout(v time.Duration) func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SlmStop) WithPretty() func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SlmStop) WithHuman() func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SlmStop) WithErrorTrace() func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SlmStop) WithFilterPath(v ...string) func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SlmStop) WithHeader(h map[string]string) func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SlmStop) WithOpaqueID(s string) func(*SlmStopRequest) {
	return func(r *SlmStopRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
