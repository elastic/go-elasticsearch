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

func newInferenceDeleteFunc(t Transport) InferenceDelete {
	return func(inference_id string, o ...func(*InferenceDeleteRequest)) (*Response, error) {
		var r = InferenceDeleteRequest{InferenceID: inference_id}
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

// InferenceDelete delete an inference endpoint
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/delete-inference-api.html.
type InferenceDelete func(inference_id string, o ...func(*InferenceDeleteRequest)) (*Response, error)

// InferenceDeleteRequest configures the Inference Delete API request.
type InferenceDeleteRequest struct {
	InferenceID string
	TaskType    string

	DryRun *bool
	Force  *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r InferenceDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.delete")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "DELETE"

	path.Grow(7 + 1 + len("_inference") + 1 + len(r.TaskType) + 1 + len(r.InferenceID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_inference")
	if r.TaskType != "" {
		path.WriteString("/")
		path.WriteString(r.TaskType)
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "task_type", r.TaskType)
		}
	}
	path.WriteString("/")
	path.WriteString(r.InferenceID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "inference_id", r.InferenceID)
	}

	params = make(map[string]string)

	if r.DryRun != nil {
		params["dry_run"] = strconv.FormatBool(*r.DryRun)
	}

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
		instrument.BeforeRequest(req, "inference.delete")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.delete")
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
func (f InferenceDelete) WithContext(v context.Context) func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		r.ctx = v
	}
}

// WithTaskType - the task type.
func (f InferenceDelete) WithTaskType(v string) func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		r.TaskType = v
	}
}

// WithDryRun - if true the endpoint will not be deleted and a list of ingest processors which reference this endpoint will be returned..
func (f InferenceDelete) WithDryRun(v bool) func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		r.DryRun = &v
	}
}

// WithForce - if true the endpoint will be forcefully stopped (regardless of whether or not it is referenced by any ingest processors or semantic text fields)..
func (f InferenceDelete) WithForce(v bool) func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		r.Force = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferenceDelete) WithPretty() func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferenceDelete) WithHuman() func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferenceDelete) WithErrorTrace() func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferenceDelete) WithFilterPath(v ...string) func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferenceDelete) WithHeader(h map[string]string) func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferenceDelete) WithOpaqueID(s string) func(*InferenceDeleteRequest) {
	return func(r *InferenceDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
