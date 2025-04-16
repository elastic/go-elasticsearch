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
// Code generated from specification version 9.1.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newInferenceCompletionFunc(t Transport) InferenceCompletion {
	return func(inference_id string, o ...func(*InferenceCompletionRequest)) (*Response, error) {
		var r = InferenceCompletionRequest{InferenceID: inference_id}
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

// InferenceCompletion perform completion inference
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/post-inference-api.html.
type InferenceCompletion func(inference_id string, o ...func(*InferenceCompletionRequest)) (*Response, error)

// InferenceCompletionRequest configures the Inference Completion API request.
type InferenceCompletionRequest struct {
	Body io.Reader

	InferenceID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r InferenceCompletionRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.completion")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_inference") + 1 + len("completion") + 1 + len(r.InferenceID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_inference")
	path.WriteString("/")
	path.WriteString("completion")
	path.WriteString("/")
	path.WriteString(r.InferenceID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "inference_id", r.InferenceID)
	}

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
		instrument.BeforeRequest(req, "inference.completion")
		if reader := instrument.RecordRequestBody(ctx, "inference.completion", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.completion")
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
func (f InferenceCompletion) WithContext(v context.Context) func(*InferenceCompletionRequest) {
	return func(r *InferenceCompletionRequest) {
		r.ctx = v
	}
}

// WithBody - The inference payload.
func (f InferenceCompletion) WithBody(v io.Reader) func(*InferenceCompletionRequest) {
	return func(r *InferenceCompletionRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferenceCompletion) WithPretty() func(*InferenceCompletionRequest) {
	return func(r *InferenceCompletionRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferenceCompletion) WithHuman() func(*InferenceCompletionRequest) {
	return func(r *InferenceCompletionRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferenceCompletion) WithErrorTrace() func(*InferenceCompletionRequest) {
	return func(r *InferenceCompletionRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferenceCompletion) WithFilterPath(v ...string) func(*InferenceCompletionRequest) {
	return func(r *InferenceCompletionRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferenceCompletion) WithHeader(h map[string]string) func(*InferenceCompletionRequest) {
	return func(r *InferenceCompletionRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferenceCompletion) WithOpaqueID(s string) func(*InferenceCompletionRequest) {
	return func(r *InferenceCompletionRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
