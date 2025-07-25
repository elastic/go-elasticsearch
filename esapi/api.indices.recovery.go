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
	"net/http"
	"strconv"
	"strings"
)

func newIndicesRecoveryFunc(t Transport) IndicesRecovery {
	return func(o ...func(*IndicesRecoveryRequest)) (*Response, error) {
		var r = IndicesRecoveryRequest{}
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

// IndicesRecovery returns information about ongoing index shard recoveries.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-recovery.html.
type IndicesRecovery func(o ...func(*IndicesRecoveryRequest)) (*Response, error)

// IndicesRecoveryRequest configures the Indices Recovery API request.
type IndicesRecoveryRequest struct {
	Index []string

	ActiveOnly        *bool
	AllowNoIndices    *bool
	Detailed          *bool
	ExpandWildcards   string
	IgnoreUnavailable *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IndicesRecoveryRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.recovery")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len(strings.Join(r.Index, ",")) + 1 + len("_recovery"))
	path.WriteString("http://")
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Index, ","))
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", strings.Join(r.Index, ","))
		}
	}
	path.WriteString("/")
	path.WriteString("_recovery")

	params = make(map[string]string)

	if r.ActiveOnly != nil {
		params["active_only"] = strconv.FormatBool(*r.ActiveOnly)
	}

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.Detailed != nil {
		params["detailed"] = strconv.FormatBool(*r.Detailed)
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
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
		instrument.BeforeRequest(req, "indices.recovery")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.recovery")
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
func (f IndicesRecovery) WithContext(v context.Context) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.ctx = v
	}
}

// WithIndex - a list of index names; use _all to perform the operation on all indices.
func (f IndicesRecovery) WithIndex(v ...string) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.Index = v
	}
}

// WithActiveOnly - display only those recoveries that are currently on-going.
func (f IndicesRecovery) WithActiveOnly(v bool) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.ActiveOnly = &v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
func (f IndicesRecovery) WithAllowNoIndices(v bool) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.AllowNoIndices = &v
	}
}

// WithDetailed - whether to display detailed information about shard recovery.
func (f IndicesRecovery) WithDetailed(v bool) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.Detailed = &v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
func (f IndicesRecovery) WithExpandWildcards(v string) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.ExpandWildcards = v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func (f IndicesRecovery) WithIgnoreUnavailable(v bool) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IndicesRecovery) WithPretty() func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IndicesRecovery) WithHuman() func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IndicesRecovery) WithErrorTrace() func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IndicesRecovery) WithFilterPath(v ...string) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IndicesRecovery) WithHeader(h map[string]string) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IndicesRecovery) WithOpaqueID(s string) func(*IndicesRecoveryRequest) {
	return func(r *IndicesRecoveryRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
