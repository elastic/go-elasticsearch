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
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newEsqlAsyncQueryGetFunc(t Transport) EsqlAsyncQueryGet {
	return func(id string, o ...func(*EsqlAsyncQueryGetRequest)) (*Response, error) {
		var r = EsqlAsyncQueryGetRequest{DocumentID: id}
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

// EsqlAsyncQueryGet - Retrieves the results of a previously submitted async query request given its ID.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/esql-async-query-get-api.html.
type EsqlAsyncQueryGet func(id string, o ...func(*EsqlAsyncQueryGetRequest)) (*Response, error)

// EsqlAsyncQueryGetRequest configures the Esql Async Query Get API request.
type EsqlAsyncQueryGetRequest struct {
	DocumentID string

	DropNullColumns          *bool
	Format                   string
	KeepAlive                time.Duration
	WaitForCompletionTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r EsqlAsyncQueryGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "esql.async_query_get")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_query") + 1 + len("async") + 1 + len(r.DocumentID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_query")
	path.WriteString("/")
	path.WriteString("async")
	path.WriteString("/")
	path.WriteString(r.DocumentID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "id", r.DocumentID)
	}

	params = make(map[string]string)

	if r.DropNullColumns != nil {
		params["drop_null_columns"] = strconv.FormatBool(*r.DropNullColumns)
	}

	if r.Format != "" {
		params["format"] = r.Format
	}

	if r.KeepAlive != 0 {
		params["keep_alive"] = formatDuration(r.KeepAlive)
	}

	if r.WaitForCompletionTimeout != 0 {
		params["wait_for_completion_timeout"] = formatDuration(r.WaitForCompletionTimeout)
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
		instrument.BeforeRequest(req, "esql.async_query_get")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "esql.async_query_get")
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
func (f EsqlAsyncQueryGet) WithContext(v context.Context) func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.ctx = v
	}
}

// WithDropNullColumns - should entirely null columns be removed from the results? their name and type will be returning in a new `all_columns` section..
func (f EsqlAsyncQueryGet) WithDropNullColumns(v bool) func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.DropNullColumns = &v
	}
}

// WithFormat - a short version of the accept header, e.g. json, yaml.
func (f EsqlAsyncQueryGet) WithFormat(v string) func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.Format = v
	}
}

// WithKeepAlive - specify the time interval in which the results (partial or final) for this search will be available.
func (f EsqlAsyncQueryGet) WithKeepAlive(v time.Duration) func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.KeepAlive = v
	}
}

// WithWaitForCompletionTimeout - specify the time that the request should block waiting for the final response.
func (f EsqlAsyncQueryGet) WithWaitForCompletionTimeout(v time.Duration) func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.WaitForCompletionTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f EsqlAsyncQueryGet) WithPretty() func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f EsqlAsyncQueryGet) WithHuman() func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f EsqlAsyncQueryGet) WithErrorTrace() func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f EsqlAsyncQueryGet) WithFilterPath(v ...string) func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f EsqlAsyncQueryGet) WithHeader(h map[string]string) func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f EsqlAsyncQueryGet) WithOpaqueID(s string) func(*EsqlAsyncQueryGetRequest) {
	return func(r *EsqlAsyncQueryGetRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
