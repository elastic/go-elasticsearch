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

func newInferenceChatCompletionUnifiedFunc(t Transport) InferenceChatCompletionUnified {
	return func(inference_id string, o ...func(*InferenceChatCompletionUnifiedRequest)) (*Response, error) {
		var r = InferenceChatCompletionUnifiedRequest{InferenceID: inference_id}
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

// InferenceChatCompletionUnified perform chat completion inference
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/chat-completion-inference.html.
type InferenceChatCompletionUnified func(inference_id string, o ...func(*InferenceChatCompletionUnifiedRequest)) (*Response, error)

// InferenceChatCompletionUnifiedRequest configures the Inference Chat Completion Unified API request.
type InferenceChatCompletionUnifiedRequest struct {
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
func (r InferenceChatCompletionUnifiedRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.chat_completion_unified")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_inference") + 1 + len("chat_completion") + 1 + len(r.InferenceID) + 1 + len("_stream"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_inference")
	path.WriteString("/")
	path.WriteString("chat_completion")
	path.WriteString("/")
	path.WriteString(r.InferenceID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "inference_id", r.InferenceID)
	}
	path.WriteString("/")
	path.WriteString("_stream")

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
		instrument.BeforeRequest(req, "inference.chat_completion_unified")
		if reader := instrument.RecordRequestBody(ctx, "inference.chat_completion_unified", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.chat_completion_unified")
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
func (f InferenceChatCompletionUnified) WithContext(v context.Context) func(*InferenceChatCompletionUnifiedRequest) {
	return func(r *InferenceChatCompletionUnifiedRequest) {
		r.ctx = v
	}
}

// WithBody - The inference payload.
func (f InferenceChatCompletionUnified) WithBody(v io.Reader) func(*InferenceChatCompletionUnifiedRequest) {
	return func(r *InferenceChatCompletionUnifiedRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferenceChatCompletionUnified) WithPretty() func(*InferenceChatCompletionUnifiedRequest) {
	return func(r *InferenceChatCompletionUnifiedRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferenceChatCompletionUnified) WithHuman() func(*InferenceChatCompletionUnifiedRequest) {
	return func(r *InferenceChatCompletionUnifiedRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferenceChatCompletionUnified) WithErrorTrace() func(*InferenceChatCompletionUnifiedRequest) {
	return func(r *InferenceChatCompletionUnifiedRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferenceChatCompletionUnified) WithFilterPath(v ...string) func(*InferenceChatCompletionUnifiedRequest) {
	return func(r *InferenceChatCompletionUnifiedRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferenceChatCompletionUnified) WithHeader(h map[string]string) func(*InferenceChatCompletionUnifiedRequest) {
	return func(r *InferenceChatCompletionUnifiedRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferenceChatCompletionUnified) WithOpaqueID(s string) func(*InferenceChatCompletionUnifiedRequest) {
	return func(r *InferenceChatCompletionUnifiedRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
