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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Get EQL search results.
// Returns search results for an Event Query Language (EQL) query.
// EQL assumes each document in a data stream or index corresponds to an event.
package search

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/resultposition"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Search struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewSearch type alias for index.
type NewSearch func(index string) *Search

// NewSearchFunc returns a new instance of Search with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSearchFunc(tp elastictransport.Interface) NewSearch {
	return func(index string) *Search {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Get EQL search results.
// Returns search results for an Event Query Language (EQL) query.
// EQL assumes each document in a data stream or index corresponds to an event.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/eql-search-api.html
func New(tp elastictransport.Interface) *Search {
	r := &Search{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Search) Raw(raw io.Reader) *Search {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Search) Request(req *Request) *Search {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Search) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Search: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_eql")
		path.WriteString("/")
		path.WriteString("search")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Search) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "eql.search")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "eql.search")
		if reader := instrument.RecordRequestBody(ctx, "eql.search", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "eql.search")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Search query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a search.Response
func (r Search) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "eql.search")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the Search headers map.
func (r *Search) Header(key, value string) *Search {
	r.headers.Set(key, value)

	return r
}

// Index The name of the index to scope the operation
// API Name: index
func (r *Search) _index(index string) *Search {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *Search) AllowNoIndices(allownoindices bool) *Search {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *Search) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Search {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// CcsMinimizeRoundtrips Indicates whether network round-trips should be minimized as part of
// cross-cluster search requests execution
// API name: ccs_minimize_roundtrips
func (r *Search) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Search {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// IgnoreUnavailable If true, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *Search) IgnoreUnavailable(ignoreunavailable bool) *Search {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Search) ErrorTrace(errortrace bool) *Search {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Search) FilterPath(filterpaths ...string) *Search {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *Search) Human(human bool) *Search {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Search) Pretty(pretty bool) *Search {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AllowPartialSearchResults Allow query execution also in case of shard failures.
// If true, the query will keep running and will return results based on the
// available shards.
// For sequences, the behavior can be further refined using
// allow_partial_sequence_results
// API name: allow_partial_search_results
func (r *Search) AllowPartialSearchResults(allowpartialsearchresults bool) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.AllowPartialSearchResults = &allowpartialsearchresults

	return r
}

// AllowPartialSequenceResults This flag applies only to sequences and has effect only if
// allow_partial_search_results=true.
// If true, the sequence query will return results based on the available
// shards, ignoring the others.
// If false, the sequence query will return successfully, but will always have
// empty results.
// API name: allow_partial_sequence_results
func (r *Search) AllowPartialSequenceResults(allowpartialsequenceresults bool) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.AllowPartialSequenceResults = &allowpartialsequenceresults

	return r
}

// API name: case_sensitive
func (r *Search) CaseSensitive(casesensitive bool) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.CaseSensitive = &casesensitive

	return r
}

// EventCategoryField Field containing the event classification, such as process, file, or network.
// API name: event_category_field
func (r *Search) EventCategoryField(field string) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.EventCategoryField = &field

	return r
}

// FetchSize Maximum number of events to search at a time for sequence queries.
// API name: fetch_size
func (r *Search) FetchSize(fetchsize uint) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.FetchSize = &fetchsize

	return r
}

// Fields Array of wildcard (*) patterns. The response returns values for field names
// matching these patterns in the fields property of each hit.
// API name: fields
func (r *Search) Fields(fields ...types.FieldAndFormat) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Fields = fields

	return r
}

// Filter Query, written in Query DSL, used to filter the events on which the EQL query
// runs.
// API name: filter
func (r *Search) Filter(filters ...types.Query) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Filter = filters

	return r
}

// API name: keep_alive
func (r *Search) KeepAlive(duration types.Duration) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.KeepAlive = duration

	return r
}

// API name: keep_on_completion
func (r *Search) KeepOnCompletion(keeponcompletion bool) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.KeepOnCompletion = &keeponcompletion

	return r
}

// MaxSamplesPerKey By default, the response of a sample query contains up to `10` samples, with
// one sample per unique set of join keys. Use the `size`
// parameter to get a smaller or larger set of samples. To retrieve more than
// one sample per set of join keys, use the
// `max_samples_per_key` parameter. Pipes are not supported for sample queries.
// API name: max_samples_per_key
func (r *Search) MaxSamplesPerKey(maxsamplesperkey int) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.MaxSamplesPerKey = &maxsamplesperkey

	return r
}

// Query EQL query you wish to run.
// API name: query
func (r *Search) Query(query string) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}

// API name: result_position
func (r *Search) ResultPosition(resultposition resultposition.ResultPosition) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ResultPosition = &resultposition

	return r
}

// API name: runtime_mappings
func (r *Search) RuntimeMappings(runtimefields types.RuntimeFields) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RuntimeMappings = runtimefields

	return r
}

// Size For basic queries, the maximum number of matching events to return. Defaults
// to 10
// API name: size
func (r *Search) Size(size uint) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Size = &size

	return r
}

// TiebreakerField Field used to sort hits with the same timestamp in ascending order
// API name: tiebreaker_field
func (r *Search) TiebreakerField(field string) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.TiebreakerField = &field

	return r
}

// TimestampField Field containing event timestamp. Default "@timestamp"
// API name: timestamp_field
func (r *Search) TimestampField(field string) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.TimestampField = &field

	return r
}

// API name: wait_for_completion_timeout
func (r *Search) WaitForCompletionTimeout(duration types.Duration) *Search {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.WaitForCompletionTimeout = duration

	return r
}
