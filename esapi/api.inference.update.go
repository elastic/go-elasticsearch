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

func newInferenceUpdateFunc(t Transport) InferenceUpdate {
	return func(inference_id string, o ...func(*InferenceUpdateRequest)) (*Response, error) {
		var r = InferenceUpdateRequest{InferenceID: inference_id}
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

// InferenceUpdate update inference
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/update-inference-api.html.
type InferenceUpdate func(inference_id string, o ...func(*InferenceUpdateRequest)) (*Response, error)

// InferenceUpdateRequest configures the Inference Update API request.
type InferenceUpdateRequest struct {
	Body io.Reader

	InferenceID string
	TaskType    string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r InferenceUpdateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.update")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_inference") + 1 + len(r.TaskType) + 1 + len(r.InferenceID) + 1 + len("_update"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_inference")
	if r.TaskType != "" {
		path.WriteString("/")
		path.WriteString(r.TaskType)
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "task_type", r.TaskType)
		}
	}
	path.WriteString("/")
	path.WriteString(r.InferenceID)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "inference_id", r.InferenceID)
	}
	path.WriteString("/")
	path.WriteString("_update")

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
		instrument.BeforeRequest(req, "inference.update")
		if reader := instrument.RecordRequestBody(ctx, "inference.update", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.update")
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
func (f InferenceUpdate) WithContext(v context.Context) func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		r.ctx = v
	}
}

// WithBody - The inference endpoint's task and service settings.
func (f InferenceUpdate) WithBody(v io.Reader) func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		r.Body = v
	}
}

// WithTaskType - the task type.
func (f InferenceUpdate) WithTaskType(v string) func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		r.TaskType = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferenceUpdate) WithPretty() func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferenceUpdate) WithHuman() func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferenceUpdate) WithErrorTrace() func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferenceUpdate) WithFilterPath(v ...string) func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferenceUpdate) WithHeader(h map[string]string) func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferenceUpdate) WithOpaqueID(s string) func(*InferenceUpdateRequest) {
	return func(r *InferenceUpdateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
