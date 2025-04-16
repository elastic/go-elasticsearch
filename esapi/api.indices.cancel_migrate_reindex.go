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
	"strings"
)

func newIndicesCancelMigrateReindexFunc(t Transport) IndicesCancelMigrateReindex {
	return func(index string, o ...func(*IndicesCancelMigrateReindexRequest)) (*Response, error) {
		var r = IndicesCancelMigrateReindexRequest{Index: index}
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

// IndicesCancelMigrateReindex this API returns the status of a migration reindex attempt for a data stream or index
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/data-stream-reindex-cancel-api.html.
type IndicesCancelMigrateReindex func(index string, o ...func(*IndicesCancelMigrateReindexRequest)) (*Response, error)

// IndicesCancelMigrateReindexRequest configures the Indices Cancel Migrate Reindex API request.
type IndicesCancelMigrateReindexRequest struct {
	Index string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IndicesCancelMigrateReindexRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.cancel_migrate_reindex")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len("_migration") + 1 + len("reindex") + 1 + len(r.Index) + 1 + len("_cancel"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_migration")
	path.WriteString("/")
	path.WriteString("reindex")
	path.WriteString("/")
	path.WriteString(r.Index)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "index", r.Index)
	}
	path.WriteString("/")
	path.WriteString("_cancel")

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
		instrument.BeforeRequest(req, "indices.cancel_migrate_reindex")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.cancel_migrate_reindex")
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
func (f IndicesCancelMigrateReindex) WithContext(v context.Context) func(*IndicesCancelMigrateReindexRequest) {
	return func(r *IndicesCancelMigrateReindexRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IndicesCancelMigrateReindex) WithPretty() func(*IndicesCancelMigrateReindexRequest) {
	return func(r *IndicesCancelMigrateReindexRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IndicesCancelMigrateReindex) WithHuman() func(*IndicesCancelMigrateReindexRequest) {
	return func(r *IndicesCancelMigrateReindexRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IndicesCancelMigrateReindex) WithErrorTrace() func(*IndicesCancelMigrateReindexRequest) {
	return func(r *IndicesCancelMigrateReindexRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IndicesCancelMigrateReindex) WithFilterPath(v ...string) func(*IndicesCancelMigrateReindexRequest) {
	return func(r *IndicesCancelMigrateReindexRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IndicesCancelMigrateReindex) WithHeader(h map[string]string) func(*IndicesCancelMigrateReindexRequest) {
	return func(r *IndicesCancelMigrateReindexRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IndicesCancelMigrateReindex) WithOpaqueID(s string) func(*IndicesCancelMigrateReindexRequest) {
	return func(r *IndicesCancelMigrateReindexRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
