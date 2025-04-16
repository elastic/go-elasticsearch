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
	"io"
	"net/http"
	"strings"
)

func newInferencePutAzureaistudioFunc(t Transport) InferencePutAzureaistudio {
	return func(azureaistudio_inference_id string, task_type string, o ...func(*InferencePutAzureaistudioRequest)) (*Response, error) {
		var r = InferencePutAzureaistudioRequest{AzureaistudioInferenceID: azureaistudio_inference_id, TaskType: task_type}
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

// InferencePutAzureaistudio configure an Azure AI Studio inference endpoint
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/infer-service-azure-ai-studio.html.
type InferencePutAzureaistudio func(azureaistudio_inference_id string, task_type string, o ...func(*InferencePutAzureaistudioRequest)) (*Response, error)

// InferencePutAzureaistudioRequest configures the Inference Put Azureaistudio API request.
type InferencePutAzureaistudioRequest struct {
	Body io.Reader

	AzureaistudioInferenceID string
	TaskType                 string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r InferencePutAzureaistudioRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_azureaistudio")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_inference") + 1 + len(r.TaskType) + 1 + len(r.AzureaistudioInferenceID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_inference")
	path.WriteString("/")
	path.WriteString(r.TaskType)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "task_type", r.TaskType)
	}
	path.WriteString("/")
	path.WriteString(r.AzureaistudioInferenceID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "azureaistudio_inference_id", r.AzureaistudioInferenceID)
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
		instrument.BeforeRequest(req, "inference.put_azureaistudio")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_azureaistudio", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_azureaistudio")
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
func (f InferencePutAzureaistudio) WithContext(v context.Context) func(*InferencePutAzureaistudioRequest) {
	return func(r *InferencePutAzureaistudioRequest) {
		r.ctx = v
	}
}

// WithBody - The inference endpoint's task and service settings.
func (f InferencePutAzureaistudio) WithBody(v io.Reader) func(*InferencePutAzureaistudioRequest) {
	return func(r *InferencePutAzureaistudioRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferencePutAzureaistudio) WithPretty() func(*InferencePutAzureaistudioRequest) {
	return func(r *InferencePutAzureaistudioRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferencePutAzureaistudio) WithHuman() func(*InferencePutAzureaistudioRequest) {
	return func(r *InferencePutAzureaistudioRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferencePutAzureaistudio) WithErrorTrace() func(*InferencePutAzureaistudioRequest) {
	return func(r *InferencePutAzureaistudioRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferencePutAzureaistudio) WithFilterPath(v ...string) func(*InferencePutAzureaistudioRequest) {
	return func(r *InferencePutAzureaistudioRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferencePutAzureaistudio) WithHeader(h map[string]string) func(*InferencePutAzureaistudioRequest) {
	return func(r *InferencePutAzureaistudioRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferencePutAzureaistudio) WithOpaqueID(s string) func(*InferencePutAzureaistudioRequest) {
	return func(r *InferencePutAzureaistudioRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
