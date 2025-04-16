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

func newQueryRulesListRulesetsFunc(t Transport) QueryRulesListRulesets {
	return func(o ...func(*QueryRulesListRulesetsRequest)) (*Response, error) {
		var r = QueryRulesListRulesetsRequest{}
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

// QueryRulesListRulesets lists query rulesets.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/list-query-rulesets.html.
type QueryRulesListRulesets func(o ...func(*QueryRulesListRulesetsRequest)) (*Response, error)

// QueryRulesListRulesetsRequest configures the Query Rules List Rulesets API request.
type QueryRulesListRulesetsRequest struct {
	From *int
	Size *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r QueryRulesListRulesetsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "query_rules.list_rulesets")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_query_rules"))
	path.WriteString("http://")
	path.WriteString("/_query_rules")

	params = make(map[string]string)

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
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
		instrument.BeforeRequest(req, "query_rules.list_rulesets")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "query_rules.list_rulesets")
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
func (f QueryRulesListRulesets) WithContext(v context.Context) func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		r.ctx = v
	}
}

// WithFrom - starting offset (default: 0).
func (f QueryRulesListRulesets) WithFrom(v int) func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		r.From = &v
	}
}

// WithSize - specifies a max number of results to get (default: 100).
func (f QueryRulesListRulesets) WithSize(v int) func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		r.Size = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f QueryRulesListRulesets) WithPretty() func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f QueryRulesListRulesets) WithHuman() func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f QueryRulesListRulesets) WithErrorTrace() func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f QueryRulesListRulesets) WithFilterPath(v ...string) func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f QueryRulesListRulesets) WithHeader(h map[string]string) func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f QueryRulesListRulesets) WithOpaqueID(s string) func(*QueryRulesListRulesetsRequest) {
	return func(r *QueryRulesListRulesetsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
