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
	"strconv"
	"strings"
	"time"
)

func newEqlSearchFunc(t Transport) EqlSearch {
	return func(index string, body io.Reader, o ...func(*EqlSearchRequest)) (*Response, error) {
		var r = EqlSearchRequest{Index: index, Body: body}
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

// EqlSearch - Returns results matching a query expressed in Event Query Language (EQL)
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html.
type EqlSearch func(index string, body io.Reader, o ...func(*EqlSearchRequest)) (*Response, error)

// EqlSearchRequest configures the Eql Search API request.
type EqlSearchRequest struct {
	Index string

	Body io.Reader

	AllowNoIndices              *bool
	AllowPartialSearchResults   *bool
	AllowPartialSequenceResults *bool
	CcsMinimizeRoundtrips       *bool
	ExpandWildcards             string
	IgnoreUnavailable           *bool
	KeepAlive                   time.Duration
	KeepOnCompletion            *bool
	WaitForCompletionTimeout    time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r EqlSearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "eql.search")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len(r.Index) + 1 + len("_eql") + 1 + len("search"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString(r.Index)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "index", r.Index)
	}
	path.WriteString("/")
	path.WriteString("_eql")
	path.WriteString("/")
	path.WriteString("search")

	params = make(map[string]string)

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.AllowPartialSearchResults != nil {
		params["allow_partial_search_results"] = strconv.FormatBool(*r.AllowPartialSearchResults)
	}

	if r.AllowPartialSequenceResults != nil {
		params["allow_partial_sequence_results"] = strconv.FormatBool(*r.AllowPartialSequenceResults)
	}

	if r.CcsMinimizeRoundtrips != nil {
		params["ccs_minimize_roundtrips"] = strconv.FormatBool(*r.CcsMinimizeRoundtrips)
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
	}

	if r.KeepAlive != 0 {
		params["keep_alive"] = formatDuration(r.KeepAlive)
	}

	if r.KeepOnCompletion != nil {
		params["keep_on_completion"] = strconv.FormatBool(*r.KeepOnCompletion)
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
		instrument.BeforeRequest(req, "eql.search")
		if reader := instrument.RecordRequestBody(ctx, "eql.search", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "eql.search")
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
func (f EqlSearch) WithContext(v context.Context) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.ctx = v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
func (f EqlSearch) WithAllowNoIndices(v bool) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.AllowNoIndices = &v
	}
}

// WithAllowPartialSearchResults - control whether the query should keep running in case of shard failures, and return partial results.
func (f EqlSearch) WithAllowPartialSearchResults(v bool) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.AllowPartialSearchResults = &v
	}
}

// WithAllowPartialSequenceResults - control whether a sequence query should return partial results or no results at all in case of shard failures. this option has effect only if [allow_partial_search_results] is true..
func (f EqlSearch) WithAllowPartialSequenceResults(v bool) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.AllowPartialSequenceResults = &v
	}
}

// WithCcsMinimizeRoundtrips - indicates whether network round-trips should be minimized as part of cross-cluster search requests execution.
func (f EqlSearch) WithCcsMinimizeRoundtrips(v bool) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.CcsMinimizeRoundtrips = &v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
func (f EqlSearch) WithExpandWildcards(v string) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.ExpandWildcards = v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func (f EqlSearch) WithIgnoreUnavailable(v bool) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithKeepAlive - update the time interval in which the results (partial or final) for this search will be available.
func (f EqlSearch) WithKeepAlive(v time.Duration) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.KeepAlive = v
	}
}

// WithKeepOnCompletion - control whether the response should be stored in the cluster if it completed within the provided [wait_for_completion] time (default: false).
func (f EqlSearch) WithKeepOnCompletion(v bool) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.KeepOnCompletion = &v
	}
}

// WithWaitForCompletionTimeout - specify the time that the request should block waiting for the final response.
func (f EqlSearch) WithWaitForCompletionTimeout(v time.Duration) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.WaitForCompletionTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f EqlSearch) WithPretty() func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f EqlSearch) WithHuman() func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f EqlSearch) WithErrorTrace() func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f EqlSearch) WithFilterPath(v ...string) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f EqlSearch) WithHeader(h map[string]string) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f EqlSearch) WithOpaqueID(s string) func(*EqlSearchRequest) {
	return func(r *EqlSearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
