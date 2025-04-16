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

func newSimulateIngestFunc(t Transport) SimulateIngest {
	return func(body io.Reader, o ...func(*SimulateIngestRequest)) (*Response, error) {
		var r = SimulateIngestRequest{Body: body}
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

// SimulateIngest simulates running ingest with example documents.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/simulate-ingest-api.html.
type SimulateIngest func(body io.Reader, o ...func(*SimulateIngestRequest)) (*Response, error)

// SimulateIngestRequest configures the Simulate Ingest API request.
type SimulateIngestRequest struct {
	Index string

	Body io.Reader

	Pipeline string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SimulateIngestRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "simulate.ingest")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_ingest") + 1 + len(r.Index) + 1 + len("_simulate"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ingest")
	if r.Index != "" {
		path.WriteString("/")
		path.WriteString(r.Index)
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.Index)
		}
	}
	path.WriteString("/")
	path.WriteString("_simulate")

	params = make(map[string]string)

	if r.Pipeline != "" {
		params["pipeline"] = r.Pipeline
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
		instrument.BeforeRequest(req, "simulate.ingest")
		if reader := instrument.RecordRequestBody(ctx, "simulate.ingest", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "simulate.ingest")
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
func (f SimulateIngest) WithContext(v context.Context) func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		r.ctx = v
	}
}

// WithIndex - default index for docs which don't provide one.
func (f SimulateIngest) WithIndex(v string) func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		r.Index = v
	}
}

// WithPipeline - the pipeline ID to preprocess incoming documents with if no pipeline is given for a particular document.
func (f SimulateIngest) WithPipeline(v string) func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		r.Pipeline = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SimulateIngest) WithPretty() func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SimulateIngest) WithHuman() func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SimulateIngest) WithErrorTrace() func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SimulateIngest) WithFilterPath(v ...string) func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SimulateIngest) WithHeader(h map[string]string) func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SimulateIngest) WithOpaqueID(s string) func(*SimulateIngestRequest) {
	return func(r *SimulateIngestRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
