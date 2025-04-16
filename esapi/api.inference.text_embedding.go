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

func newInferenceTextEmbeddingFunc(t Transport) InferenceTextEmbedding {
	return func(inference_id string, o ...func(*InferenceTextEmbeddingRequest)) (*Response, error) {
		var r = InferenceTextEmbeddingRequest{InferenceID: inference_id}
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

// InferenceTextEmbedding perform text embedding inference
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/post-inference-api.html.
type InferenceTextEmbedding func(inference_id string, o ...func(*InferenceTextEmbeddingRequest)) (*Response, error)

// InferenceTextEmbeddingRequest configures the Inference Text Embedding API request.
type InferenceTextEmbeddingRequest struct {
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
func (r InferenceTextEmbeddingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "inference.text_embedding")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_inference") + 1 + len("text_embedding") + 1 + len(r.InferenceID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_inference")
	path.WriteString("/")
	path.WriteString("text_embedding")
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
		instrument.BeforeRequest(req, "inference.text_embedding")
		if reader := instrument.RecordRequestBody(ctx, "inference.text_embedding", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "inference.text_embedding")
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
func (f InferenceTextEmbedding) WithContext(v context.Context) func(*InferenceTextEmbeddingRequest) {
	return func(r *InferenceTextEmbeddingRequest) {
		r.ctx = v
	}
}

// WithBody - The inference payload.
func (f InferenceTextEmbedding) WithBody(v io.Reader) func(*InferenceTextEmbeddingRequest) {
	return func(r *InferenceTextEmbeddingRequest) {
		r.Body = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f InferenceTextEmbedding) WithPretty() func(*InferenceTextEmbeddingRequest) {
	return func(r *InferenceTextEmbeddingRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f InferenceTextEmbedding) WithHuman() func(*InferenceTextEmbeddingRequest) {
	return func(r *InferenceTextEmbeddingRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f InferenceTextEmbedding) WithErrorTrace() func(*InferenceTextEmbeddingRequest) {
	return func(r *InferenceTextEmbeddingRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f InferenceTextEmbedding) WithFilterPath(v ...string) func(*InferenceTextEmbeddingRequest) {
	return func(r *InferenceTextEmbeddingRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f InferenceTextEmbedding) WithHeader(h map[string]string) func(*InferenceTextEmbeddingRequest) {
	return func(r *InferenceTextEmbeddingRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f InferenceTextEmbedding) WithOpaqueID(s string) func(*InferenceTextEmbeddingRequest) {
	return func(r *InferenceTextEmbeddingRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
