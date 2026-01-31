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
// Code generated from specification version 9.4.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
	"time"
)

func newIndicesDeleteSampleConfigurationFunc(t Transport) IndicesDeleteSampleConfiguration {
	return func(index string, o ...func(*IndicesDeleteSampleConfigurationRequest)) (*Response, error) {
		var r = IndicesDeleteSampleConfigurationRequest{Index: index}
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

// IndicesDeleteSampleConfiguration delete sampling configuration for an index or data stream
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/docs/api/doc/elasticsearch#TODO.
type IndicesDeleteSampleConfiguration func(index string, o ...func(*IndicesDeleteSampleConfigurationRequest)) (*Response, error)

// IndicesDeleteSampleConfigurationRequest configures the Indices Delete Sample Configuration API request.
type IndicesDeleteSampleConfigurationRequest struct {
	Index string

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
func (r IndicesDeleteSampleConfigurationRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.delete_sample_configuration")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "DELETE"

	path.Grow(7 + 1 + len(r.Index) + 1 + len("_sample") + 1 + len("config"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString(r.Index)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "index", r.Index)
	}
	path.WriteString("/")
	path.WriteString("_sample")
	path.WriteString("/")
	path.WriteString("config")

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
		instrument.BeforeRequest(req, "indices.delete_sample_configuration")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.delete_sample_configuration")
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
func (f IndicesDeleteSampleConfiguration) WithContext(v context.Context) func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - timeout for connection to master node.
func (f IndicesDeleteSampleConfiguration) WithMasterTimeout(v time.Duration) func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - timeout for the request.
func (f IndicesDeleteSampleConfiguration) WithTimeout(v time.Duration) func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IndicesDeleteSampleConfiguration) WithPretty() func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IndicesDeleteSampleConfiguration) WithHuman() func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IndicesDeleteSampleConfiguration) WithErrorTrace() func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IndicesDeleteSampleConfiguration) WithFilterPath(v ...string) func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IndicesDeleteSampleConfiguration) WithHeader(h map[string]string) func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IndicesDeleteSampleConfiguration) WithOpaqueID(s string) func(*IndicesDeleteSampleConfigurationRequest) {
	return func(r *IndicesDeleteSampleConfigurationRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
