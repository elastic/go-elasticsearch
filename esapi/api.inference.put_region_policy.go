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
// Code generated from specification version 9.5.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newInferencePutRegionPolicyFunc(t Transport) InferencePutRegionPolicy {
	return func(body io.Reader, o ...func(*InferencePutRegionPolicyRequest)) (*Response, error) {
		var r = InferencePutRegionPolicyRequest{Body: body}
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

// InferencePutRegionPolicy create or update the inference region policy
//
// See full documentation at https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-put-region-policy.
type InferencePutRegionPolicy func(body io.Reader, o ...func(*InferencePutRegionPolicyRequest)) (*Response, error)

// InferencePutRegionPolicyRequest configures the Inference Put Region Policy API request.
type InferencePutRegionPolicyRequest struct {
	Body io.Reader

	Force *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r InferencePutRegionPolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_region_policy")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + len("/_inference/_region_policy"))
	path.WriteString("http://")
	path.WriteString("/_inference/_region_policy")

	params = make(map[string]string)

	if r.Force != nil {
		params["force"] = strconv.FormatBool(*r.Force)
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

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "inference.put_region_policy")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_region_policy", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_region_policy")
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
func (f InferencePutRegionPolicy) WithContext(v context.Context) func(*InferencePutRegionPolicyRequest) {
	return func(r *InferencePutRegionPolicyRequest) {
		r.ctx = v
	}
}

// WithForce - if true, the region policy will be applied even if it would deny access to inference endpoints that are currently in use by ingest pipelines or semantic text fields..
func (f InferencePutRegionPolicy) WithForce(v bool) func(*InferencePutRegionPolicyRequest) {
	return func(r *InferencePutRegionPolicyRequest) {
		r.Force = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferencePutRegionPolicy) WithPretty() func(*InferencePutRegionPolicyRequest) {
	return func(r *InferencePutRegionPolicyRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferencePutRegionPolicy) WithHuman() func(*InferencePutRegionPolicyRequest) {
	return func(r *InferencePutRegionPolicyRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferencePutRegionPolicy) WithErrorTrace() func(*InferencePutRegionPolicyRequest) {
	return func(r *InferencePutRegionPolicyRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferencePutRegionPolicy) WithFilterPath(v ...string) func(*InferencePutRegionPolicyRequest) {
	return func(r *InferencePutRegionPolicyRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferencePutRegionPolicy) WithHeader(h map[string]string) func(*InferencePutRegionPolicyRequest) {
	return func(r *InferencePutRegionPolicyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferencePutRegionPolicy) WithOpaqueID(s string) func(*InferencePutRegionPolicyRequest) {
	return func(r *InferencePutRegionPolicyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
