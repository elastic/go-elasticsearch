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
// https://github.com/elastic/elasticsearch-specification/tree/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c

// Returns results matching a query.
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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	index string
}

// NewSearch type alias for index.
type NewSearch func() *Search

// NewSearchFunc returns a new instance of Search with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSearchFunc(tp elastictransport.Interface) NewSearch {
	return func() *Search {
		n := New(tp)

		return n
	}
}

// Returns results matching a query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/search-search.html
func New(tp elastictransport.Interface) *Search {
	r := &Search{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Search: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_search")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_search")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
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
func (r Search) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Search query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a search.Response
func (r Search) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	r.TypedKeys(true)

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the Search headers map.
func (r *Search) Header(key, value string) *Search {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to search; use `_all` or empty string
// to perform the operation on all indices
// API Name: index
func (r *Search) Index(index string) *Search {
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

// AllowPartialSearchResults Indicate if an error should be returned if there is a partial search failure
// or timeout
// API name: allow_partial_search_results
func (r *Search) AllowPartialSearchResults(allowpartialsearchresults bool) *Search {
	r.values.Set("allow_partial_search_results", strconv.FormatBool(allowpartialsearchresults))

	return r
}

// Analyzer The analyzer to use for the query string
// API name: analyzer
func (r *Search) Analyzer(analyzer string) *Search {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard Specify whether wildcard and prefix queries should be analyzed (default:
// false)
// API name: analyze_wildcard
func (r *Search) AnalyzeWildcard(analyzewildcard bool) *Search {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// BatchedReduceSize The number of shard results that should be reduced at once on the
// coordinating node. This value should be used as a protection mechanism to
// reduce the memory overhead per search request if the potential number of
// shards in the request can be large.
// API name: batched_reduce_size
func (r *Search) BatchedReduceSize(batchedreducesize string) *Search {
	r.values.Set("batched_reduce_size", batchedreducesize)

	return r
}

// CcsMinimizeRoundtrips Indicates whether network round-trips should be minimized as part of
// cross-cluster search requests execution
// API name: ccs_minimize_roundtrips
func (r *Search) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Search {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// DefaultOperator The default operator for query string query (AND or OR)
// API name: default_operator
func (r *Search) DefaultOperator(defaultoperator operator.Operator) *Search {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string
// API name: df
func (r *Search) Df(df string) *Search {
	r.values.Set("df", df)

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

// IgnoreThrottled Whether specified concrete, expanded or aliased indices should be ignored
// when throttled
// API name: ignore_throttled
func (r *Search) IgnoreThrottled(ignorethrottled bool) *Search {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *Search) IgnoreUnavailable(ignoreunavailable bool) *Search {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Lenient Specify whether format-based query failures (such as providing text to a
// numeric field) should be ignored
// API name: lenient
func (r *Search) Lenient(lenient bool) *Search {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// MaxConcurrentShardRequests The number of concurrent shard requests per node this search executes
// concurrently. This value should be used to limit the impact of the search on
// the cluster in order to limit the number of concurrent shard requests
// API name: max_concurrent_shard_requests
func (r *Search) MaxConcurrentShardRequests(maxconcurrentshardrequests string) *Search {
	r.values.Set("max_concurrent_shard_requests", maxconcurrentshardrequests)

	return r
}

// MinCompatibleShardNode The minimum compatible version that all shards involved in search should have
// for this request to be successful
// API name: min_compatible_shard_node
func (r *Search) MinCompatibleShardNode(versionstring string) *Search {
	r.values.Set("min_compatible_shard_node", versionstring)

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *Search) Preference(preference string) *Search {
	r.values.Set("preference", preference)

	return r
}

// PreFilterShardSize A threshold that enforces a pre-filter roundtrip to prefilter search shards
// based on query rewriting if the number of shards the search request expands
// to exceeds the threshold. This filter roundtrip can limit the number of
// shards significantly if for instance a shard can not match any documents
// based on its rewrite method ie. if date filters are mandatory to match but
// the shard bounds and the query are disjoint.
// API name: pre_filter_shard_size
func (r *Search) PreFilterShardSize(prefiltershardsize string) *Search {
	r.values.Set("pre_filter_shard_size", prefiltershardsize)

	return r
}

// RequestCache Specify if request cache should be used for this request or not, defaults to
// index level setting
// API name: request_cache
func (r *Search) RequestCache(requestcache bool) *Search {
	r.values.Set("request_cache", strconv.FormatBool(requestcache))

	return r
}

// Routing A comma-separated list of specific routing values
// API name: routing
func (r *Search) Routing(routing string) *Search {
	r.values.Set("routing", routing)

	return r
}

// Scroll Specify how long a consistent view of the index should be maintained for
// scrolled search
// API name: scroll
func (r *Search) Scroll(duration string) *Search {
	r.values.Set("scroll", duration)

	return r
}

// SearchType Search operation type
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

// SuggestMode Specify suggest mode
// API name: suggest_mode
func (r *Search) SuggestMode(suggestmode suggestmode.SuggestMode) *Search {
	r.values.Set("suggest_mode", suggestmode.String())

	return r
}

// SuggestSize How many suggestions to return in response
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

// TypedKeys Specify whether aggregation and suggester names should be prefixed by their
// respective types in the response
// API name: typed_keys
func (r *Search) TypedKeys(typedkeys bool) *Search {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// RestTotalHitsAsInt Indicates whether hits.total should be rendered as an integer or an object in
// the rest search response
// API name: rest_total_hits_as_int
func (r *Search) RestTotalHitsAsInt(resttotalhitsasint bool) *Search {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// SourceExcludes_ A list of fields to exclude from the returned _source field
// API name: _source_excludes
func (r *Search) SourceExcludes_(fields ...string) *Search {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A list of fields to extract and return from the _source field
// API name: _source_includes
func (r *Search) SourceIncludes_(fields ...string) *Search {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// Q Query in the Lucene query string syntax
// API name: q
func (r *Search) Q(q string) *Search {
	r.values.Set("q", q)

	return r
}

// API name: aggregations
func (r *Search) Aggregations(aggregations map[string]types.Aggregations) *Search {

	r.req.Aggregations = aggregations

	return r
}

// API name: collapse
func (r *Search) Collapse(collapse *types.FieldCollapse) *Search {

	r.req.Collapse = collapse

	return r
}

// DocvalueFields Array of wildcard (*) patterns. The request returns doc values for field
// names matching these patterns in the hits.fields property of the response.
// API name: docvalue_fields
func (r *Search) DocvalueFields(docvaluefields ...types.FieldAndFormat) *Search {
	r.req.DocvalueFields = docvaluefields

	return r
}

// Explain If true, returns detailed information about score computation as part of a
// hit.
// API name: explain
func (r *Search) Explain(explain bool) *Search {
	r.req.Explain = &explain

	return r
}

// Ext Configuration of search extensions defined by Elasticsearch plugins.
// API name: ext
func (r *Search) Ext(ext map[string]json.RawMessage) *Search {

	r.req.Ext = ext

	return r
}

// Fields Array of wildcard (*) patterns. The request returns values for field names
// matching these patterns in the hits.fields property of the response.
// API name: fields
func (r *Search) Fields(fields ...types.FieldAndFormat) *Search {
	r.req.Fields = fields

	return r
}

// From Starting document offset. By default, you cannot page through more than
// 10,000
// hits using the from and size parameters. To page through more hits, use the
// search_after parameter.
// API name: from
func (r *Search) From(from int) *Search {
	r.req.From = &from

	return r
}

// API name: highlight
func (r *Search) Highlight(highlight *types.Highlight) *Search {

	r.req.Highlight = highlight

	return r
}

// IndicesBoost Boosts the _score of documents from specified indices.
// API name: indices_boost
func (r *Search) IndicesBoost(indicesboosts ...map[string]types.Float64) *Search {
	r.req.IndicesBoost = indicesboosts

	return r
}

// Knn Defines the approximate kNN search to run.
// API name: knn
func (r *Search) Knn(knns ...types.KnnQuery) *Search {
	r.req.Knn = knns

	return r
}

// MinScore Minimum _score for matching documents. Documents with a lower _score are
// not included in the search results.
// API name: min_score
func (r *Search) MinScore(minscore types.Float64) *Search {

	r.req.MinScore = &minscore

	return r
}

// Pit Limits the search to a point in time (PIT). If you provide a PIT, you
// cannot specify an <index> in the request path.
// API name: pit
func (r *Search) Pit(pit *types.PointInTimeReference) *Search {

	r.req.Pit = pit

	return r
}

// API name: post_filter
func (r *Search) PostFilter(postfilter *types.Query) *Search {

	r.req.PostFilter = postfilter

	return r
}

// API name: profile
func (r *Search) Profile(profile bool) *Search {
	r.req.Profile = &profile

	return r
}

// Query Defines the search definition using the Query DSL.
// API name: query
func (r *Search) Query(query *types.Query) *Search {

	r.req.Query = query

	return r
}

// Rank Defines the Reciprocal Rank Fusion (RRF) to use
// API name: rank
func (r *Search) Rank(rank *types.RankContainer) *Search {

	r.req.Rank = rank

	return r
}

// API name: rescore
func (r *Search) Rescore(rescores ...types.Rescore) *Search {
	r.req.Rescore = rescores

	return r
}

// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
// precedence over mapped fields with the same name.
// API name: runtime_mappings
func (r *Search) RuntimeMappings(runtimefields types.RuntimeFields) *Search {
	r.req.RuntimeMappings = runtimefields

	return r
}

// ScriptFields Retrieve a script evaluation (based on different fields) for each hit.
// API name: script_fields
func (r *Search) ScriptFields(scriptfields map[string]types.ScriptField) *Search {

	r.req.ScriptFields = scriptfields

	return r
}

// API name: search_after
func (r *Search) SearchAfter(sortresults ...types.FieldValue) *Search {
	r.req.SearchAfter = sortresults

	return r
}

// SeqNoPrimaryTerm If true, returns sequence number and primary term of the last modification
// of each hit. See Optimistic concurrency control.
// API name: seq_no_primary_term
func (r *Search) SeqNoPrimaryTerm(seqnoprimaryterm bool) *Search {
	r.req.SeqNoPrimaryTerm = &seqnoprimaryterm

	return r
}

// Size The number of hits to return. By default, you cannot page through more
// than 10,000 hits using the from and size parameters. To page through more
// hits, use the search_after parameter.
// API name: size
func (r *Search) Size(size int) *Search {
	r.req.Size = &size

	return r
}

// API name: slice
func (r *Search) Slice(slice *types.SlicedScroll) *Search {

	r.req.Slice = slice

	return r
}

// API name: sort
func (r *Search) Sort(sorts ...types.SortCombinations) *Search {
	r.req.Sort = sorts

	return r
}

// Source_ Indicates which source fields are returned for matching documents. These
// fields are returned in the hits._source property of the search response.
// API name: _source
func (r *Search) Source_(sourceconfig types.SourceConfig) *Search {
	r.req.Source_ = sourceconfig

	return r
}

// Stats Stats groups to associate with the search. Each group maintains a statistics
// aggregation for its associated searches. You can retrieve these stats using
// the indices stats API.
// API name: stats
func (r *Search) Stats(stats ...string) *Search {
	r.req.Stats = stats

	return r
}

// StoredFields List of stored fields to return as part of a hit. If no fields are specified,
// no stored fields are included in the response. If this field is specified,
// the _source
// parameter defaults to false. You can pass _source: true to return both source
// fields
// and stored fields in the search response.
// API name: stored_fields
func (r *Search) StoredFields(fields ...string) *Search {
	r.req.StoredFields = fields

	return r
}

// API name: suggest
func (r *Search) Suggest(suggest *types.Suggester) *Search {

	r.req.Suggest = suggest

	return r
}

// TerminateAfter Maximum number of documents to collect for each shard. If a query reaches
// this
// limit, Elasticsearch terminates the query early. Elasticsearch collects
// documents
// before sorting. Defaults to 0, which does not terminate query execution
// early.
// API name: terminate_after
func (r *Search) TerminateAfter(terminateafter int64) *Search {

	r.req.TerminateAfter = &terminateafter

	return r
}

// Timeout Specifies the period of time to wait for a response from each shard. If no
// response
// is received before the timeout expires, the request fails and returns an
// error.
// Defaults to no timeout.
// API name: timeout
func (r *Search) Timeout(timeout string) *Search {

	r.req.Timeout = &timeout

	return r
}

// TrackScores If true, calculate and return document scores, even if the scores are not
// used for sorting.
// API name: track_scores
func (r *Search) TrackScores(trackscores bool) *Search {
	r.req.TrackScores = &trackscores

	return r
}

// TrackTotalHits Number of hits matching the query to count accurately. If true, the exact
// number of hits is returned at the cost of some performance. If false, the
// response does not include the total number of hits matching the query.
// Defaults to 10,000 hits.
// API name: track_total_hits
func (r *Search) TrackTotalHits(trackhits types.TrackHits) *Search {
	r.req.TrackTotalHits = trackhits

	return r
}

// Version If true, returns document version as part of a hit.
// API name: version
func (r *Search) Version(version bool) *Search {
	r.req.Version = &version

	return r
}
