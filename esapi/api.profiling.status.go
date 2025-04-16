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
	"strconv"
	"strings"
	"time"
)

func newProfilingStatusFunc(t Transport) ProfilingStatus {
	return func(o ...func(*ProfilingStatusRequest)) (*Response, error) {
		var r = ProfilingStatusRequest{}
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

// ProfilingStatus returns basic information about the status of Universal Profiling.
//
// See full documentation at https://www.elastic.co/guide/en/observability/current/universal-profiling.html.
type ProfilingStatus func(o ...func(*ProfilingStatusRequest)) (*Response, error)

// ProfilingStatusRequest configures the Profiling Status API request.
type ProfilingStatusRequest struct {
	MasterTimeout           time.Duration
	Timeout                 time.Duration
	WaitForResourcesCreated *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ProfilingStatusRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "profiling.status")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_profiling/status"))
	path.WriteString("http://")
	path.WriteString("/_profiling/status")

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.WaitForResourcesCreated != nil {
		params["wait_for_resources_created"] = strconv.FormatBool(*r.WaitForResourcesCreated)
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
		instrument.BeforeRequest(req, "profiling.status")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "profiling.status")
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
func (f ProfilingStatus) WithContext(v context.Context) func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func (f ProfilingStatus) WithMasterTimeout(v time.Duration) func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
func (f ProfilingStatus) WithTimeout(v time.Duration) func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		r.Timeout = v
	}
}

// WithWaitForResourcesCreated - whether to return immediately or wait until resources have been created.
func (f ProfilingStatus) WithWaitForResourcesCreated(v bool) func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		r.WaitForResourcesCreated = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ProfilingStatus) WithPretty() func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ProfilingStatus) WithHuman() func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ProfilingStatus) WithErrorTrace() func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ProfilingStatus) WithFilterPath(v ...string) func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ProfilingStatus) WithHeader(h map[string]string) func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ProfilingStatus) WithOpaqueID(s string) func(*ProfilingStatusRequest) {
	return func(r *ProfilingStatusRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
