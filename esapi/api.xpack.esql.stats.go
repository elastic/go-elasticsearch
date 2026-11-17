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
// Code generated from specification version 9.6.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newEsqlStatsFunc(t Transport) EsqlStats {
	return func(o ...func(*EsqlStatsRequest)) (*Response, error) {
		var r = EsqlStatsRequest{}
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

// EsqlStats - Get live ES|QL datasource admission gauges (used, limit, untracked, waiters, permits, rescues). Sample during load; after-query occupancy is idle. Core GET /_nodes/stats is not pluggable for this plugin.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-esql-stats.
type EsqlStats func(o ...func(*EsqlStatsRequest)) (*Response, error)

// EsqlStatsRequest configures the Esql Stats API request.
type EsqlStatsRequest struct {
	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r EsqlStatsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "esql.stats")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_esql/stats"))
	path.WriteString("http://")
	path.WriteString("/_esql/stats")

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
		instrument.BeforeRequest(req, "esql.stats")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "esql.stats")
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
func (f EsqlStats) WithContext(v context.Context) func(*EsqlStatsRequest) {
	return func(r *EsqlStatsRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f EsqlStats) WithPretty() func(*EsqlStatsRequest) {
	return func(r *EsqlStatsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f EsqlStats) WithHuman() func(*EsqlStatsRequest) {
	return func(r *EsqlStatsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f EsqlStats) WithErrorTrace() func(*EsqlStatsRequest) {
	return func(r *EsqlStatsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f EsqlStats) WithFilterPath(v ...string) func(*EsqlStatsRequest) {
	return func(r *EsqlStatsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f EsqlStats) WithHeader(h map[string]string) func(*EsqlStatsRequest) {
	return func(r *EsqlStatsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f EsqlStats) WithOpaqueID(s string) func(*EsqlStatsRequest) {
	return func(r *EsqlStatsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
