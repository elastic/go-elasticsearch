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

func newInferencePutVoyageaiFunc(t Transport) InferencePutVoyageai {
	return func(task_type string, voyageai_inference_id string, o ...func(*InferencePutVoyageaiRequest)) (*Response, error) {
		var r = InferencePutVoyageaiRequest{TaskType: task_type, VoyageaiInferenceID: voyageai_inference_id}
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

// InferencePutVoyageai configure a VoyageAI inference endpoint
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/inference-apis.html.
type InferencePutVoyageai func(task_type string, voyageai_inference_id string, o ...func(*InferencePutVoyageaiRequest)) (*Response, error)

// InferencePutVoyageaiRequest configures the Inference Put Voyageai API request.
type InferencePutVoyageaiRequest struct {
	Body io.Reader

	TaskType            string
	VoyageaiInferenceID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r InferencePutVoyageaiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_voyageai")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_inference") + 1 + len(r.TaskType) + 1 + len(r.VoyageaiInferenceID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_inference")
	path.WriteString("/")
	path.WriteString(r.TaskType)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "task_type", r.TaskType)
	}
	path.WriteString("/")
	path.WriteString(r.VoyageaiInferenceID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "voyageai_inference_id", r.VoyageaiInferenceID)
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
		instrument.BeforeRequest(req, "inference.put_voyageai")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_voyageai", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_voyageai")
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
func (f InferencePutVoyageai) WithContext(v context.Context) func(*InferencePutVoyageaiRequest) {
	return func(r *InferencePutVoyageaiRequest) {
		r.ctx = v
	}
}

// WithBody - The inference endpoint's task and service settings.
func (f InferencePutVoyageai) WithBody(v io.Reader) func(*InferencePutVoyageaiRequest) {
	return func(r *InferencePutVoyageaiRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferencePutVoyageai) WithPretty() func(*InferencePutVoyageaiRequest) {
	return func(r *InferencePutVoyageaiRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferencePutVoyageai) WithHuman() func(*InferencePutVoyageaiRequest) {
	return func(r *InferencePutVoyageaiRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferencePutVoyageai) WithErrorTrace() func(*InferencePutVoyageaiRequest) {
	return func(r *InferencePutVoyageaiRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferencePutVoyageai) WithFilterPath(v ...string) func(*InferencePutVoyageaiRequest) {
	return func(r *InferencePutVoyageaiRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferencePutVoyageai) WithHeader(h map[string]string) func(*InferencePutVoyageaiRequest) {
	return func(r *InferencePutVoyageaiRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferencePutVoyageai) WithOpaqueID(s string) func(*InferencePutVoyageaiRequest) {
	return func(r *InferencePutVoyageaiRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
