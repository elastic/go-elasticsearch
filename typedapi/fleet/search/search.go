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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Run a Fleet search.
// The purpose of the Fleet search API is to provide an API where the search
// will be run only
// after the provided checkpoint has been processed and is visible for searches
// inside of Elasticsearch.
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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestmode"
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

// Run a Fleet search.
// The purpose of the Fleet search API is to provide an API where the search
// will be run only
// after the provided checkpoint has been processed and is visible for searches
// inside of Elasticsearch.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-fleet-search
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
		path.WriteString("_fleet")
		path.WriteString("/")
		path.WriteString("_fleet_search")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=9")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
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
			ctx := instrument.Start(providedCtx, "fleet.search")
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
		instrument.BeforeRequest(req, "fleet.search")
		if reader := instrument.RecordRequestBody(ctx, "fleet.search", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "fleet.search")
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
		ctx = instrument.Start(providedCtx, "fleet.search")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	r.TypedKeys(true)

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

// Index A single target to search. If the target is an index alias, it must resolve
// to a single index.
// API Name: index
func (r *Search) _index(index string) *Search {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// API name: allow_no_indices
func (r *Search) AllowNoIndices(allownoindices bool) *Search {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// API name: analyzer
func (r *Search) Analyzer(analyzer string) *Search {
	r.values.Set("analyzer", analyzer)

	return r
}

// API name: analyze_wildcard
func (r *Search) AnalyzeWildcard(analyzewildcard bool) *Search {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// API name: batched_reduce_size
func (r *Search) BatchedReduceSize(batchedreducesize string) *Search {
	r.values.Set("batched_reduce_size", batchedreducesize)

	return r
}

// API name: ccs_minimize_roundtrips
func (r *Search) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Search {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// API name: default_operator
func (r *Search) DefaultOperator(defaultoperator operator.Operator) *Search {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// API name: df
func (r *Search) Df(df string) *Search {
	r.values.Set("df", df)

	return r
}

// API name: expand_wildcards
func (r *Search) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Search {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// API name: ignore_throttled
func (r *Search) IgnoreThrottled(ignorethrottled bool) *Search {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// API name: ignore_unavailable
func (r *Search) IgnoreUnavailable(ignoreunavailable bool) *Search {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// API name: lenient
func (r *Search) Lenient(lenient bool) *Search {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// API name: max_concurrent_shard_requests
func (r *Search) MaxConcurrentShardRequests(maxconcurrentshardrequests int) *Search {
	r.values.Set("max_concurrent_shard_requests", strconv.Itoa(maxconcurrentshardrequests))

	return r
}

// API name: preference
func (r *Search) Preference(preference string) *Search {
	r.values.Set("preference", preference)

	return r
}

// API name: pre_filter_shard_size
func (r *Search) PreFilterShardSize(prefiltershardsize string) *Search {
	r.values.Set("pre_filter_shard_size", prefiltershardsize)

	return r
}

// API name: request_cache
func (r *Search) RequestCache(requestcache bool) *Search {
	r.values.Set("request_cache", strconv.FormatBool(requestcache))

	return r
}

// API name: routing
func (r *Search) Routing(routing string) *Search {
	r.values.Set("routing", routing)

	return r
}

// API name: scroll
func (r *Search) Scroll(duration string) *Search {
	r.values.Set("scroll", duration)

	return r
}

// API name: search_type
func (r *Search) SearchType(searchtype searchtype.SearchType) *Search {
	r.values.Set("search_type", searchtype.String())

	return r
}

// SuggestField Specifies which field to use for suggestions.
// API name: suggest_field
func (r *Search) SuggestField(field string) *Search {
	r.values.Set("suggest_field", field)

	return r
}

// API name: suggest_mode
func (r *Search) SuggestMode(suggestmode suggestmode.SuggestMode) *Search {
	r.values.Set("suggest_mode", suggestmode.String())

	return r
}

// API name: suggest_size
func (r *Search) SuggestSize(suggestsize string) *Search {
	r.values.Set("suggest_size", suggestsize)

	return r
}

// SuggestText The source text for which the suggestions should be returned.
// API name: suggest_text
func (r *Search) SuggestText(suggesttext string) *Search {
	r.values.Set("suggest_text", suggesttext)

	return r
}

// API name: typed_keys
func (r *Search) TypedKeys(typedkeys bool) *Search {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// API name: rest_total_hits_as_int
func (r *Search) RestTotalHitsAsInt(resttotalhitsasint bool) *Search {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// API name: _source_excludes
func (r *Search) SourceExcludes_(fields ...string) *Search {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// API name: _source_includes
func (r *Search) SourceIncludes_(fields ...string) *Search {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// API name: q
func (r *Search) Q(q string) *Search {
	r.values.Set("q", q)

	return r
}

// WaitForCheckpoints A comma separated list of checkpoints. When configured, the search API will
// only be executed on a shard
// after the relevant checkpoint has become visible for search. Defaults to an
// empty list which will cause
// Elasticsearch to immediately execute the search.
// API name: wait_for_checkpoints
func (r *Search) WaitForCheckpoints(waitforcheckpoints ...int64) *Search {
	tmp := []string{}
	for _, item := range waitforcheckpoints {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("wait_for_checkpoints", strings.Join(tmp, ","))

	return r
}

// AllowPartialSearchResults If true, returns partial results if there are shard request timeouts or shard
// failures.
// If false, returns an error with no partial results.
// Defaults to the configured cluster setting
// `search.default_allow_partial_results`, which is true by default.
// API name: allow_partial_search_results
func (r *Search) AllowPartialSearchResults(allowpartialsearchresults bool) *Search {
	r.values.Set("allow_partial_search_results", strconv.FormatBool(allowpartialsearchresults))

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
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
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

// API name: aggregations
func (r *Search) Aggregations(aggregations map[string]types.Aggregations) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Aggregations = aggregations
	return r
}

func (r *Search) AddAggregation(key string, value types.AggregationsVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]types.Aggregations
	if r.req.Aggregations == nil {
		r.req.Aggregations = make(map[string]types.Aggregations)
	} else {
		tmp = r.req.Aggregations
	}

	tmp[key] = *value.AggregationsCaster()

	r.req.Aggregations = tmp
	return r
}

// API name: collapse
func (r *Search) Collapse(collapse types.FieldCollapseVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Collapse = collapse.FieldCollapseCaster()

	return r
}

// Array of wildcard (*) patterns. The request returns doc values for field
// names matching these patterns in the hits.fields property of the response.
// API name: docvalue_fields
func (r *Search) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range docvaluefields {

		r.req.DocvalueFields = append(r.req.DocvalueFields, *v.FieldAndFormatCaster())

	}
	return r
}

// If true, returns detailed information about score computation as part of a
// hit.
// API name: explain
func (r *Search) Explain(explain bool) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Explain = &explain

	return r
}

// Configuration of search extensions defined by Elasticsearch plugins.
// API name: ext
func (r *Search) Ext(ext map[string]json.RawMessage) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Ext = ext
	return r
}

func (r *Search) AddExt(key string, value json.RawMessage) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]json.RawMessage
	if r.req.Ext == nil {
		r.req.Ext = make(map[string]json.RawMessage)
	} else {
		tmp = r.req.Ext
	}

	tmp[key] = value

	r.req.Ext = tmp
	return r
}

// Array of wildcard (*) patterns. The request returns values for field names
// matching these patterns in the hits.fields property of the response.
// API name: fields
func (r *Search) Fields(fields ...types.FieldAndFormatVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range fields {

		r.req.Fields = append(r.req.Fields, *v.FieldAndFormatCaster())

	}
	return r
}

// Starting document offset. By default, you cannot page through more than
// 10,000
// hits using the from and size parameters. To page through more hits, use the
// search_after parameter.
// API name: from
func (r *Search) From(from int) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.From = &from

	return r
}

// API name: highlight
func (r *Search) Highlight(highlight types.HighlightVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Highlight = highlight.HighlightCaster()

	return r
}

// Boosts the _score of documents from specified indices.
// API name: indices_boost
func (r *Search) IndicesBoost(indicesboost []map[string]types.Float64) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndicesBoost = indicesboost

	return r
}

// Minimum _score for matching documents. Documents with a lower _score are
// not included in search results and results collected by aggregations.
// API name: min_score
func (r *Search) MinScore(minscore types.Float64) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MinScore = &minscore

	return r
}

// Limits the search to a point in time (PIT). If you provide a PIT, you
// cannot specify an <index> in the request path.
// API name: pit
func (r *Search) Pit(pit types.PointInTimeReferenceVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Pit = pit.PointInTimeReferenceCaster()

	return r
}

// API name: post_filter
func (r *Search) PostFilter(postfilter types.QueryVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PostFilter = postfilter.QueryCaster()

	return r
}

// API name: profile
func (r *Search) Profile(profile bool) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Profile = &profile

	return r
}

// Defines the search definition using the Query DSL.
// API name: query
func (r *Search) Query(query types.QueryVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query.QueryCaster()

	return r
}

// API name: rescore
func (r *Search) Rescore(rescores ...types.RescoreVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Rescore = make([]types.Rescore, len(rescores))
	for i, v := range rescores {
		r.req.Rescore[i] = *v.RescoreCaster()
	}

	return r
}

// Defines one or more runtime fields in the search request. These fields take
// precedence over mapped fields with the same name.
// API name: runtime_mappings
func (r *Search) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return r
}

// Retrieve a script evaluation (based on different fields) for each hit.
// API name: script_fields
func (r *Search) ScriptFields(scriptfields map[string]types.ScriptField) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ScriptFields = scriptfields
	return r
}

func (r *Search) AddScriptField(key string, value types.ScriptFieldVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]types.ScriptField
	if r.req.ScriptFields == nil {
		r.req.ScriptFields = make(map[string]types.ScriptField)
	} else {
		tmp = r.req.ScriptFields
	}

	tmp[key] = *value.ScriptFieldCaster()

	r.req.ScriptFields = tmp
	return r
}

// API name: search_after
func (r *Search) SearchAfter(sortresults ...types.FieldValueVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	for _, v := range sortresults {
		r.req.SearchAfter = append(r.req.SearchAfter, *v.FieldValueCaster())
	}

	return r
}

// If true, returns sequence number and primary term of the last modification
// of each hit. See Optimistic concurrency control.
// API name: seq_no_primary_term
func (r *Search) SeqNoPrimaryTerm(seqnoprimaryterm bool) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.SeqNoPrimaryTerm = &seqnoprimaryterm

	return r
}

// The number of hits to return. By default, you cannot page through more
// than 10,000 hits using the from and size parameters. To page through more
// hits, use the search_after parameter.
// API name: size
func (r *Search) Size(size int) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Size = &size

	return r
}

// API name: slice
func (r *Search) Slice(slice types.SlicedScrollVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Slice = slice.SlicedScrollCaster()

	return r
}

// API name: sort
func (r *Search) Sort(sorts ...types.SortCombinationsVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	for _, v := range sorts {
		r.req.Sort = append(r.req.Sort, *v.SortCombinationsCaster())
	}

	return r
}

// Indicates which source fields are returned for matching documents. These
// fields are returned in the hits._source property of the search response.
// API name: _source
func (r *Search) Source_(sourceconfig types.SourceConfigVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Source_ = *sourceconfig.SourceConfigCaster()

	return r
}

// Stats groups to associate with the search. Each group maintains a statistics
// aggregation for its associated searches. You can retrieve these stats using
// the indices stats API.
// API name: stats
func (r *Search) Stats(stats ...string) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range stats {

		r.req.Stats = append(r.req.Stats, v)

	}
	return r
}

// List of stored fields to return as part of a hit. If no fields are specified,
// no stored fields are included in the response. If this field is specified,
// the _source
// parameter defaults to false. You can pass _source: true to return both source
// fields
// and stored fields in the search response.
// API name: stored_fields
func (r *Search) StoredFields(fields ...string) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.StoredFields = fields

	return r
}

// API name: suggest
func (r *Search) Suggest(suggest types.SuggesterVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Suggest = suggest.SuggesterCaster()

	return r
}

// Maximum number of documents to collect for each shard. If a query reaches
// this
// limit, Elasticsearch terminates the query early. Elasticsearch collects
// documents
// before sorting. Defaults to 0, which does not terminate query execution
// early.
// API name: terminate_after
func (r *Search) TerminateAfter(terminateafter int64) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TerminateAfter = &terminateafter

	return r
}

// Specifies the period of time to wait for a response from each shard. If no
// response
// is received before the timeout expires, the request fails and returns an
// error.
// Defaults to no timeout.
// API name: timeout
func (r *Search) Timeout(timeout string) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Timeout = &timeout

	return r
}

// If true, calculate and return document scores, even if the scores are not
// used for sorting.
// API name: track_scores
func (r *Search) TrackScores(trackscores bool) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TrackScores = &trackscores

	return r
}

// Number of hits matching the query to count accurately. If true, the exact
// number of hits is returned at the cost of some performance. If false, the
// response does not include the total number of hits matching the query.
// Defaults to 10,000 hits.
// API name: track_total_hits
func (r *Search) TrackTotalHits(trackhits types.TrackHitsVariant) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TrackTotalHits = *trackhits.TrackHitsCaster()

	return r
}

// If true, returns document version as part of a hit.
// API name: version
func (r *Search) Version(version bool) *Search {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Version = &version

	return r
}
