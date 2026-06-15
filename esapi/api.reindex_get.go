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
// Code generated from specification version 9.4.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newReindexGetFunc(t Transport) ReindexGet {
	return func(task_id string, o ...func(*ReindexGetRequest)) (*Response, error) {
		var r = ReindexGetRequest{TaskID: task_id}
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

// ReindexGet get information for a reindex operation
//
// This API is beta.
//
// See full documentation at https://www.elastic.co/docs/api/doc/elasticsearch#TODO.
type ReindexGet func(task_id string, o ...func(*ReindexGetRequest)) (*Response, error)

// ReindexGetRequest configures the Reindex Get API request.
type ReindexGetRequest struct {
	TaskID string

	Timeout           time.Duration
	WaitForCompletion *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ReindexGetRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "reindex_get")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_reindex") + 1 + len(r.TaskID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_reindex")
	path.WriteString("/")
	path.WriteString(r.TaskID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "task_id", r.TaskID)
	}

	params = make(map[string]string)

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.WaitForCompletion != nil {
		params["wait_for_completion"] = strconv.FormatBool(*r.WaitForCompletion)
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
		instrument.BeforeRequest(req, "reindex_get")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "reindex_get")
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
func (f ReindexGet) WithContext(v context.Context) func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		r.ctx = v
	}
}

// WithTimeout - explicit operation timeout, only used when wait_for_completion is true.
func (f ReindexGet) WithTimeout(v time.Duration) func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		r.Timeout = v
	}
}

// WithWaitForCompletion - wait for the reindex operation to complete (default: false).
func (f ReindexGet) WithWaitForCompletion(v bool) func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		r.WaitForCompletion = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ReindexGet) WithPretty() func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ReindexGet) WithHuman() func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ReindexGet) WithErrorTrace() func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ReindexGet) WithFilterPath(v ...string) func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ReindexGet) WithHeader(h map[string]string) func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ReindexGet) WithOpaqueID(s string) func(*ReindexGetRequest) {
	return func(r *ReindexGetRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
