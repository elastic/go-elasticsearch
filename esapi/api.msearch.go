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
)

func newMsearchFunc(t Transport) Msearch {
	return func(body io.Reader, o ...func(*MsearchRequest)) (*Response, error) {
		var r = MsearchRequest{Body: body}
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

// Msearch allows to execute several search operations in one request.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/search-multi-search.html.
type Msearch func(body io.Reader, o ...func(*MsearchRequest)) (*Response, error)

// MsearchRequest configures the Msearch API request.
type MsearchRequest struct {
	Index []string

	Body io.Reader

	AllowNoIndices             *bool
	CcsMinimizeRoundtrips      *bool
	ExpandWildcards            string
	IgnoreThrottled            *bool
	IgnoreUnavailable          *bool
	IncludeNamedQueriesScore   *bool
	MaxConcurrentSearches      *int
	MaxConcurrentShardRequests *int
	PreFilterShardSize         *int
	RestTotalHitsAsInt         *bool
	Routing                    []string
	SearchType                 string
	TypedKeys                  *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r MsearchRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "msearch")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + 1 + len(strings.Join(r.Index, ",")) + 1 + len("_msearch"))
	path.WriteString("http://")
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Index, ","))
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", strings.Join(r.Index, ","))
		}
	}
	path.WriteString("/")
	path.WriteString("_msearch")

	params = make(map[string]string)

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.CcsMinimizeRoundtrips != nil {
		params["ccs_minimize_roundtrips"] = strconv.FormatBool(*r.CcsMinimizeRoundtrips)
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.IgnoreThrottled != nil {
		params["ignore_throttled"] = strconv.FormatBool(*r.IgnoreThrottled)
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
	}

	if r.IncludeNamedQueriesScore != nil {
		params["include_named_queries_score"] = strconv.FormatBool(*r.IncludeNamedQueriesScore)
	}

	if len(r.Index) > 0 {
		params["index"] = strings.Join(r.Index, ",")
	}

	if r.MaxConcurrentSearches != nil {
		params["max_concurrent_searches"] = strconv.FormatInt(int64(*r.MaxConcurrentSearches), 10)
	}

	if r.MaxConcurrentShardRequests != nil {
		params["max_concurrent_shard_requests"] = strconv.FormatInt(int64(*r.MaxConcurrentShardRequests), 10)
	}

	if r.PreFilterShardSize != nil {
		params["pre_filter_shard_size"] = strconv.FormatInt(int64(*r.PreFilterShardSize), 10)
	}

	if r.RestTotalHitsAsInt != nil {
		params["rest_total_hits_as_int"] = strconv.FormatBool(*r.RestTotalHitsAsInt)
	}

	if len(r.Routing) > 0 {
		params["routing"] = strings.Join(r.Routing, ",")
	}

	if r.SearchType != "" {
		params["search_type"] = r.SearchType
	}

	if r.TypedKeys != nil {
		params["typed_keys"] = strconv.FormatBool(*r.TypedKeys)
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
		instrument.BeforeRequest(req, "msearch")
		if reader := instrument.RecordRequestBody(ctx, "msearch", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "msearch")
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
func (f Msearch) WithContext(v context.Context) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.ctx = v
	}
}

// WithIndex - a list of index names to use as default.
func (f Msearch) WithIndex(v ...string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.Index = v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
func (f Msearch) WithAllowNoIndices(v bool) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.AllowNoIndices = &v
	}
}

// WithCcsMinimizeRoundtrips - indicates whether network round-trips should be minimized as part of cross-cluster search requests execution.
func (f Msearch) WithCcsMinimizeRoundtrips(v bool) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.CcsMinimizeRoundtrips = &v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
func (f Msearch) WithExpandWildcards(v string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.ExpandWildcards = v
	}
}

// WithIgnoreThrottled - whether specified concrete, expanded or aliased indices should be ignored when throttled.
func (f Msearch) WithIgnoreThrottled(v bool) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.IgnoreThrottled = &v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func (f Msearch) WithIgnoreUnavailable(v bool) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithIncludeNamedQueriesScore - indicates whether hit.matched_queries should be rendered as a map that includes the name of the matched query associated with its score (true) or as an array containing the name of the matched queries (false).
func (f Msearch) WithIncludeNamedQueriesScore(v bool) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.IncludeNamedQueriesScore = &v
	}
}

// WithMaxConcurrentSearches - controls the maximum number of concurrent searches the multi search api will execute.
func (f Msearch) WithMaxConcurrentSearches(v int) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.MaxConcurrentSearches = &v
	}
}

// WithMaxConcurrentShardRequests - the number of concurrent shard requests each sub search executes concurrently per node. this value should be used to limit the impact of the search on the cluster in order to limit the number of concurrent shard requests.
func (f Msearch) WithMaxConcurrentShardRequests(v int) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.MaxConcurrentShardRequests = &v
	}
}

// WithPreFilterShardSize - a threshold that enforces a pre-filter roundtrip to prefilter search shards based on query rewriting if the number of shards the search request expands to exceeds the threshold. this filter roundtrip can limit the number of shards significantly if for instance a shard can not match any documents based on its rewrite method ie. if date filters are mandatory to match but the shard bounds and the query are disjoint..
func (f Msearch) WithPreFilterShardSize(v int) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.PreFilterShardSize = &v
	}
}

// WithRestTotalHitsAsInt - indicates whether hits.total should be rendered as an integer or an object in the rest search response.
func (f Msearch) WithRestTotalHitsAsInt(v bool) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.RestTotalHitsAsInt = &v
	}
}

// WithRouting - a list of specific routing values.
func (f Msearch) WithRouting(v ...string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.Routing = v
	}
}

// WithSearchType - search operation type.
func (f Msearch) WithSearchType(v string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.SearchType = v
	}
}

// WithTypedKeys - specify whether aggregation and suggester names should be prefixed by their respective types in the response.
func (f Msearch) WithTypedKeys(v bool) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.TypedKeys = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f Msearch) WithPretty() func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f Msearch) WithHuman() func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f Msearch) WithErrorTrace() func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f Msearch) WithFilterPath(v ...string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f Msearch) WithHeader(h map[string]string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f Msearch) WithOpaqueID(s string) func(*MsearchRequest) {
	return func(r *MsearchRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
