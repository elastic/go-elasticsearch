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
)

func newCapabilitiesFunc(t Transport) Capabilities {
	return func(o ...func(*CapabilitiesRequest)) (*Response, error) {
		var r = CapabilitiesRequest{}
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

// Capabilities checks if the specified combination of method, API, parameters, and arbitrary capabilities are supported
//
// This API is experimental.
//
// See full documentation at https://github.com/elastic/elasticsearch/blob/main/rest-api-spec/src/yamlRestTest/resources/rest-api-spec/test/README.asciidoc#require-or-skip-api-capabilities.
type Capabilities func(o ...func(*CapabilitiesRequest)) (*Response, error)

// CapabilitiesRequest configures the Capabilities API request.
type CapabilitiesRequest struct {
	Capabilities string
	LocalOnly    *bool
	Method       string
	Parameters   string
	Path         string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r CapabilitiesRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "capabilities")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_capabilities"))
	path.WriteString("http://")
	path.WriteString("/_capabilities")

	params = make(map[string]string)

	if r.Capabilities != "" {
		params["capabilities"] = r.Capabilities
	}

	if r.LocalOnly != nil {
		params["local_only"] = strconv.FormatBool(*r.LocalOnly)
	}

	if r.Method != "" {
		params["method"] = r.Method
	}

	if r.Parameters != "" {
		params["parameters"] = r.Parameters
	}

	if r.Path != "" {
		params["path"] = r.Path
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
		instrument.BeforeRequest(req, "capabilities")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "capabilities")
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
func (f Capabilities) WithContext(v context.Context) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.ctx = v
	}
}

// WithCapabilities - comma-separated list of arbitrary api capabilities to check.
func (f Capabilities) WithCapabilities(v string) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.Capabilities = v
	}
}

// WithLocalOnly - true if only the node being called should be considered.
func (f Capabilities) WithLocalOnly(v bool) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.LocalOnly = &v
	}
}

// WithMethod - rest method to check.
func (f Capabilities) WithMethod(v string) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.Method = v
	}
}

// WithParameters - comma-separated list of api parameters to check.
func (f Capabilities) WithParameters(v string) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.Parameters = v
	}
}

// WithPath - api path to check.
func (f Capabilities) WithPath(v string) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.Path = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f Capabilities) WithPretty() func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f Capabilities) WithHuman() func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f Capabilities) WithErrorTrace() func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f Capabilities) WithFilterPath(v ...string) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f Capabilities) WithHeader(h map[string]string) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f Capabilities) WithOpaqueID(s string) func(*CapabilitiesRequest) {
	return func(r *CapabilitiesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
