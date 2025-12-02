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
// Code generated from specification version 9.3.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

func newInferencePutOpenshiftAiFunc(t Transport) InferencePutOpenshiftAi {
	return func(body io.Reader, openshiftai_inference_id string, task_type string, o ...func(*InferencePutOpenshiftAiRequest)) (*Response, error) {
		var r = InferencePutOpenshiftAiRequest{Body: body, OpenshiftaiInferenceID: openshiftai_inference_id, TaskType: task_type}
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

// InferencePutOpenshiftAi create an OpenShift AI inference endpoint
//
// See full documentation at https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-inference-put-openshift-ai.
type InferencePutOpenshiftAi func(body io.Reader, openshiftai_inference_id string, task_type string, o ...func(*InferencePutOpenshiftAiRequest)) (*Response, error)

// InferencePutOpenshiftAiRequest configures the Inference Put Openshift Ai API request.
type InferencePutOpenshiftAiRequest struct {
	Body io.Reader

	OpenshiftaiInferenceID string
	TaskType               string

	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r InferencePutOpenshiftAiRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.put_openshift_ai")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_inference") + 1 + len(r.TaskType) + 1 + len(r.OpenshiftaiInferenceID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_inference")
	path.WriteString("/")
	path.WriteString(r.TaskType)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "task_type", r.TaskType)
	}
	path.WriteString("/")
	path.WriteString(r.OpenshiftaiInferenceID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "openshiftai_inference_id", r.OpenshiftaiInferenceID)
	}

	params = make(map[string]string)

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
		instrument.BeforeRequest(req, "inference.put_openshift_ai")
		if reader := instrument.RecordRequestBody(ctx, "inference.put_openshift_ai", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.put_openshift_ai")
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
func (f InferencePutOpenshiftAi) WithContext(v context.Context) func(*InferencePutOpenshiftAiRequest) {
	return func(r *InferencePutOpenshiftAiRequest) {
		r.ctx = v
	}
}

// WithTimeout - specifies the amount of time to wait for the inference endpoint to be created..
func (f InferencePutOpenshiftAi) WithTimeout(v time.Duration) func(*InferencePutOpenshiftAiRequest) {
	return func(r *InferencePutOpenshiftAiRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferencePutOpenshiftAi) WithPretty() func(*InferencePutOpenshiftAiRequest) {
	return func(r *InferencePutOpenshiftAiRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferencePutOpenshiftAi) WithHuman() func(*InferencePutOpenshiftAiRequest) {
	return func(r *InferencePutOpenshiftAiRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferencePutOpenshiftAi) WithErrorTrace() func(*InferencePutOpenshiftAiRequest) {
	return func(r *InferencePutOpenshiftAiRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferencePutOpenshiftAi) WithFilterPath(v ...string) func(*InferencePutOpenshiftAiRequest) {
	return func(r *InferencePutOpenshiftAiRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferencePutOpenshiftAi) WithHeader(h map[string]string) func(*InferencePutOpenshiftAiRequest) {
	return func(r *InferencePutOpenshiftAiRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferencePutOpenshiftAi) WithOpaqueID(s string) func(*InferencePutOpenshiftAiRequest) {
	return func(r *InferencePutOpenshiftAiRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
