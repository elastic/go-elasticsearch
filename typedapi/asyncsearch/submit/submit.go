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

// Run an async search.
//
// When the primary sort of the results is an indexed field, shards get sorted
// based on minimum and maximum value that they hold for that field. Partial
// results become available following the sort criteria that was requested.
//
// Warning: Asynchronous search does not support scroll or search requests that
// include only the suggest section.
//
// By default, Elasticsearch does not allow you to store an async search
// response larger than 10Mb and an attempt to do this results in an error.
// The maximum allowed size for a stored async search response can be set by
// changing the `search.max_async_search_response_size` cluster level setting.
package submit

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

type Submit struct {
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

// NewSubmit type alias for index.
type NewSubmit func() *Submit

// NewSubmitFunc returns a new instance of Submit with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSubmitFunc(tp elastictransport.Interface) NewSubmit {
	return func() *Submit {
		n := New(tp)

		return n
	}
}

// Run an async search.
//
// When the primary sort of the results is an indexed field, shards get sorted
// based on minimum and maximum value that they hold for that field. Partial
// results become available following the sort criteria that was requested.
//
// Warning: Asynchronous search does not support scroll or search requests that
// include only the suggest section.
//
// By default, Elasticsearch does not allow you to store an async search
// response larger than 10Mb and an attempt to do this results in an error.
// The maximum allowed size for a stored async search response can be set by
// changing the `search.max_async_search_response_size` cluster level setting.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-async-search-submit
func New(tp elastictransport.Interface) *Submit {
	r := &Submit{
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
func (r *Submit) Raw(raw io.Reader) *Submit {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Submit) Request(req *Request) *Submit {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Submit) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Submit: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_async_search")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_async_search")

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
func (r Submit) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "async_search.submit")
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
		instrument.BeforeRequest(req, "async_search.submit")
		if reader := instrument.RecordRequestBody(ctx, "async_search.submit", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "async_search.submit")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Submit query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a submit.Response
func (r Submit) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "async_search.submit")
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

// Header set a key, value pair in the Submit headers map.
func (r *Submit) Header(key, value string) *Submit {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to search; use `_all` or empty string
// to perform the operation on all indices
// API Name: index
func (r *Submit) Index(index string) *Submit {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// WaitForCompletionTimeout Blocks and waits until the search is completed up to a certain timeout.
// When the async search completes within the timeout, the response won’t
// include the ID as the results are not stored in the cluster.
// API name: wait_for_completion_timeout
func (r *Submit) WaitForCompletionTimeout(duration string) *Submit {
	r.values.Set("wait_for_completion_timeout", duration)

	return r
}

// KeepAlive Specifies how long the async search needs to be available.
// Ongoing async searches and any saved search results are deleted after this
// period.
// API name: keep_alive
func (r *Submit) KeepAlive(duration string) *Submit {
	r.values.Set("keep_alive", duration)

	return r
}

// KeepOnCompletion If `true`, results are stored for later retrieval when the search completes
// within the `wait_for_completion_timeout`.
// API name: keep_on_completion
func (r *Submit) KeepOnCompletion(keeponcompletion bool) *Submit {
	r.values.Set("keep_on_completion", strconv.FormatBool(keeponcompletion))

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *Submit) AllowNoIndices(allownoindices bool) *Submit {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// AllowPartialSearchResults Indicate if an error should be returned if there is a partial search failure
// or timeout
// API name: allow_partial_search_results
func (r *Submit) AllowPartialSearchResults(allowpartialsearchresults bool) *Submit {
	r.values.Set("allow_partial_search_results", strconv.FormatBool(allowpartialsearchresults))

	return r
}

// Analyzer The analyzer to use for the query string
// API name: analyzer
func (r *Submit) Analyzer(analyzer string) *Submit {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard Specify whether wildcard and prefix queries should be analyzed (default:
// false)
// API name: analyze_wildcard
func (r *Submit) AnalyzeWildcard(analyzewildcard bool) *Submit {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// BatchedReduceSize Affects how often partial results become available, which happens whenever
// shard results are reduced.
// A partial reduction is performed every time the coordinating node has
// received a certain number of new shard responses (5 by default).
// API name: batched_reduce_size
func (r *Submit) BatchedReduceSize(batchedreducesize string) *Submit {
	r.values.Set("batched_reduce_size", batchedreducesize)

	return r
}

// CcsMinimizeRoundtrips The default value is the only supported value.
// API name: ccs_minimize_roundtrips
func (r *Submit) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Submit {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// DefaultOperator The default operator for query string query (AND or OR)
// API name: default_operator
func (r *Submit) DefaultOperator(defaultoperator operator.Operator) *Submit {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string
// API name: df
func (r *Submit) Df(df string) *Submit {
	r.values.Set("df", df)

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *Submit) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Submit {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled Whether specified concrete, expanded or aliased indices should be ignored
// when throttled
// API name: ignore_throttled
func (r *Submit) IgnoreThrottled(ignorethrottled bool) *Submit {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *Submit) IgnoreUnavailable(ignoreunavailable bool) *Submit {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Lenient Specify whether format-based query failures (such as providing text to a
// numeric field) should be ignored
// API name: lenient
func (r *Submit) Lenient(lenient bool) *Submit {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// MaxConcurrentShardRequests The number of concurrent shard requests per node this search executes
// concurrently. This value should be used to limit the impact of the search on
// the cluster in order to limit the number of concurrent shard requests
// API name: max_concurrent_shard_requests
func (r *Submit) MaxConcurrentShardRequests(maxconcurrentshardrequests int) *Submit {
	r.values.Set("max_concurrent_shard_requests", strconv.Itoa(maxconcurrentshardrequests))

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *Submit) Preference(preference string) *Submit {
	r.values.Set("preference", preference)

	return r
}

// RequestCache Specify if request cache should be used for this request or not, defaults to
// true
// API name: request_cache
func (r *Submit) RequestCache(requestcache bool) *Submit {
	r.values.Set("request_cache", strconv.FormatBool(requestcache))

	return r
}

// Routing A comma-separated list of specific routing values
// API name: routing
func (r *Submit) Routing(routing string) *Submit {
	r.values.Set("routing", routing)

	return r
}

// SearchType Search operation type
// API name: search_type
func (r *Submit) SearchType(searchtype searchtype.SearchType) *Submit {
	r.values.Set("search_type", searchtype.String())

	return r
}

// SuggestField Specifies which field to use for suggestions.
// API name: suggest_field
func (r *Submit) SuggestField(field string) *Submit {
	r.values.Set("suggest_field", field)

	return r
}

// SuggestMode Specify suggest mode
// API name: suggest_mode
func (r *Submit) SuggestMode(suggestmode suggestmode.SuggestMode) *Submit {
	r.values.Set("suggest_mode", suggestmode.String())

	return r
}

// SuggestSize How many suggestions to return in response
// API name: suggest_size
func (r *Submit) SuggestSize(suggestsize string) *Submit {
	r.values.Set("suggest_size", suggestsize)

	return r
}

// SuggestText The source text for which the suggestions should be returned.
// API name: suggest_text
func (r *Submit) SuggestText(suggesttext string) *Submit {
	r.values.Set("suggest_text", suggesttext)

	return r
}

// TypedKeys Specify whether aggregation and suggester names should be prefixed by their
// respective types in the response
// API name: typed_keys
func (r *Submit) TypedKeys(typedkeys bool) *Submit {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// RestTotalHitsAsInt Indicates whether hits.total should be rendered as an integer or an object in
// the rest search response
// API name: rest_total_hits_as_int
func (r *Submit) RestTotalHitsAsInt(resttotalhitsasint bool) *Submit {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// SourceExcludes_ A list of fields to exclude from the returned _source field
// API name: _source_excludes
func (r *Submit) SourceExcludes_(fields ...string) *Submit {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A list of fields to extract and return from the _source field
// API name: _source_includes
func (r *Submit) SourceIncludes_(fields ...string) *Submit {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// Q Query in the Lucene query string syntax
// API name: q
func (r *Submit) Q(q string) *Submit {
	r.values.Set("q", q)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Submit) ErrorTrace(errortrace bool) *Submit {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Submit) FilterPath(filterpaths ...string) *Submit {
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
func (r *Submit) Human(human bool) *Submit {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Submit) Pretty(pretty bool) *Submit {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: aggregations
func (r *Submit) Aggregations(aggregations map[string]types.Aggregations) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Aggregations = aggregations
	return r
}

func (r *Submit) AddAggregation(key string, value types.AggregationsVariant) *Submit {
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
func (r *Submit) Collapse(collapse types.FieldCollapseVariant) *Submit {
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
func (r *Submit) DocvalueFields(docvaluefields ...types.FieldAndFormatVariant) *Submit {
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
func (r *Submit) Explain(explain bool) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Explain = &explain

	return r
}

// Configuration of search extensions defined by Elasticsearch plugins.
// API name: ext
func (r *Submit) Ext(ext map[string]json.RawMessage) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Ext = ext
	return r
}

func (r *Submit) AddExt(key string, value json.RawMessage) *Submit {
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
func (r *Submit) Fields(fields ...types.FieldAndFormatVariant) *Submit {
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
func (r *Submit) From(from int) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.From = &from

	return r
}

// API name: highlight
func (r *Submit) Highlight(highlight types.HighlightVariant) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Highlight = highlight.HighlightCaster()

	return r
}

// Boosts the _score of documents from specified indices.
// API name: indices_boost
func (r *Submit) IndicesBoost(indicesboost []map[string]types.Float64) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndicesBoost = indicesboost

	return r
}

// Defines the approximate kNN search to run.
// API name: knn
func (r *Submit) Knn(knns ...types.KnnSearchVariant) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Knn = make([]types.KnnSearch, len(knns))
	for i, v := range knns {
		r.req.Knn[i] = *v.KnnSearchCaster()
	}

	return r
}

// Minimum _score for matching documents. Documents with a lower _score are
// not included in search results and results collected by aggregations.
// API name: min_score
func (r *Submit) MinScore(minscore types.Float64) *Submit {
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
func (r *Submit) Pit(pit types.PointInTimeReferenceVariant) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Pit = pit.PointInTimeReferenceCaster()

	return r
}

// API name: post_filter
func (r *Submit) PostFilter(postfilter types.QueryVariant) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PostFilter = postfilter.QueryCaster()

	return r
}

// API name: profile
func (r *Submit) Profile(profile bool) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Profile = &profile

	return r
}

// Defines the search definition using the Query DSL.
// API name: query
func (r *Submit) Query(query types.QueryVariant) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query.QueryCaster()

	return r
}

// API name: rescore
func (r *Submit) Rescore(rescores ...types.RescoreVariant) *Submit {
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
func (r *Submit) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return r
}

// Retrieve a script evaluation (based on different fields) for each hit.
// API name: script_fields
func (r *Submit) ScriptFields(scriptfields map[string]types.ScriptField) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ScriptFields = scriptfields
	return r
}

func (r *Submit) AddScriptField(key string, value types.ScriptFieldVariant) *Submit {
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
func (r *Submit) SearchAfter(sortresults ...types.FieldValueVariant) *Submit {
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
func (r *Submit) SeqNoPrimaryTerm(seqnoprimaryterm bool) *Submit {
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
func (r *Submit) Size(size int) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Size = &size

	return r
}

// API name: slice
func (r *Submit) Slice(slice types.SlicedScrollVariant) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Slice = slice.SlicedScrollCaster()

	return r
}

// API name: sort
func (r *Submit) Sort(sorts ...types.SortCombinationsVariant) *Submit {
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
func (r *Submit) Source_(sourceconfig types.SourceConfigVariant) *Submit {
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
func (r *Submit) Stats(stats ...string) *Submit {
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
func (r *Submit) StoredFields(fields ...string) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.StoredFields = fields

	return r
}

// API name: suggest
func (r *Submit) Suggest(suggest types.SuggesterVariant) *Submit {
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
func (r *Submit) TerminateAfter(terminateafter int64) *Submit {
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
func (r *Submit) Timeout(timeout string) *Submit {
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
func (r *Submit) TrackScores(trackscores bool) *Submit {
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
func (r *Submit) TrackTotalHits(trackhits types.TrackHitsVariant) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.TrackTotalHits = *trackhits.TrackHitsCaster()

	return r
}

// If true, returns document version as part of a hit.
// API name: version
func (r *Submit) Version(version bool) *Submit {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Version = &version

	return r
}
