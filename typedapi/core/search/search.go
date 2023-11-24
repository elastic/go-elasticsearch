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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

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
// https://www.elastic.co/guide/en/elasticsearch/reference/current/search-search.html
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

// Index Comma-separated list of data streams, indices, and aliases to search.
// Supports wildcards (`*`).
// To search all data streams and indices, omit this parameter or use `*` or
// `_all`.
// API Name: index
func (r *Search) Index(index string) *Search {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// For example, a request targeting `foo*,bar*` returns an error if an index
// starts with `foo` but no index starts with `bar`.
// API name: allow_no_indices
func (r *Search) AllowNoIndices(allownoindices bool) *Search {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// AllowPartialSearchResults If true, returns partial results if there are shard request timeouts or shard
// failures. If false, returns an error with no partial results.
// API name: allow_partial_search_results
func (r *Search) AllowPartialSearchResults(allowpartialsearchresults bool) *Search {
	r.values.Set("allow_partial_search_results", strconv.FormatBool(allowpartialsearchresults))

	return r
}

// Analyzer Analyzer to use for the query string.
// This parameter can only be used when the q query string parameter is
// specified.
// API name: analyzer
func (r *Search) Analyzer(analyzer string) *Search {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard If true, wildcard and prefix queries are analyzed.
// This parameter can only be used when the q query string parameter is
// specified.
// API name: analyze_wildcard
func (r *Search) AnalyzeWildcard(analyzewildcard bool) *Search {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// BatchedReduceSize The number of shard results that should be reduced at once on the
// coordinating node.
// This value should be used as a protection mechanism to reduce the memory
// overhead per search request if the potential number of shards in the request
// can be large.
// API name: batched_reduce_size
func (r *Search) BatchedReduceSize(batchedreducesize string) *Search {
	r.values.Set("batched_reduce_size", batchedreducesize)

	return r
}

// CcsMinimizeRoundtrips If true, network round-trips between the coordinating node and the remote
// clusters are minimized when executing cross-cluster search (CCS) requests.
// API name: ccs_minimize_roundtrips
func (r *Search) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Search {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// DefaultOperator The default operator for query string query: AND or OR.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: default_operator
func (r *Search) DefaultOperator(defaultoperator operator.Operator) *Search {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df Field to use as default where no field prefix is given in the query string.
// This parameter can only be used when the q query string parameter is
// specified.
// API name: df
func (r *Search) Df(df string) *Search {
	r.values.Set("df", df)

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *Search) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Search {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled If `true`, concrete, expanded or aliased indices will be ignored when frozen.
// API name: ignore_throttled
func (r *Search) IgnoreThrottled(ignorethrottled bool) *Search {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *Search) IgnoreUnavailable(ignoreunavailable bool) *Search {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Lenient If `true`, format-based query failures (such as providing text to a numeric
// field) in the query string will be ignored.
// This parameter can only be used when the `q` query string parameter is
// specified.
// API name: lenient
func (r *Search) Lenient(lenient bool) *Search {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// MaxConcurrentShardRequests Defines the number of concurrent shard requests per node this search executes
// concurrently.
// This value should be used to limit the impact of the search on the cluster in
// order to limit the number of concurrent shard requests.
// API name: max_concurrent_shard_requests
func (r *Search) MaxConcurrentShardRequests(maxconcurrentshardrequests string) *Search {
	r.values.Set("max_concurrent_shard_requests", maxconcurrentshardrequests)

	return r
}

// MinCompatibleShardNode The minimum version of the node that can handle the request
// Any handling node with a lower version will fail the request.
// API name: min_compatible_shard_node
func (r *Search) MinCompatibleShardNode(versionstring string) *Search {
	r.values.Set("min_compatible_shard_node", versionstring)

	return r
}

// Preference Nodes and shards used for the search.
// By default, Elasticsearch selects from eligible nodes and shards using
// adaptive replica selection, accounting for allocation awareness. Valid values
// are:
// `_only_local` to run the search only on shards on the local node;
// `_local` to, if possible, run the search on shards on the local node, or if
// not, select shards using the default method;
// `_only_nodes:<node-id>,<node-id>` to run the search on only the specified
// nodes IDs, where, if suitable shards exist on more than one selected node,
// use shards on those nodes using the default method, or if none of the
// specified nodes are available, select shards from any available node using
// the default method;
// `_prefer_nodes:<node-id>,<node-id>` to if possible, run the search on the
// specified nodes IDs, or if not, select shards using the default method;
// `_shards:<shard>,<shard>` to run the search only on the specified shards;
// `<custom-string>` (any string that does not start with `_`) to route searches
// with the same `<custom-string>` to the same shards in the same order.
// API name: preference
func (r *Search) Preference(preference string) *Search {
	r.values.Set("preference", preference)

	return r
}

// PreFilterShardSize Defines a threshold that enforces a pre-filter roundtrip to prefilter search
// shards based on query rewriting if the number of shards the search request
// expands to exceeds the threshold.
// This filter roundtrip can limit the number of shards significantly if for
// instance a shard can not match any documents based on its rewrite method (if
// date filters are mandatory to match but the shard bounds and the query are
// disjoint).
// When unspecified, the pre-filter phase is executed if any of these conditions
// is met:
// the request targets more than 128 shards;
// the request targets one or more read-only index;
// the primary sort of the query targets an indexed field.
// API name: pre_filter_shard_size
func (r *Search) PreFilterShardSize(prefiltershardsize string) *Search {
	r.values.Set("pre_filter_shard_size", prefiltershardsize)

	return r
}

// RequestCache If `true`, the caching of search results is enabled for requests where `size`
// is `0`.
// Defaults to index level settings.
// API name: request_cache
func (r *Search) RequestCache(requestcache bool) *Search {
	r.values.Set("request_cache", strconv.FormatBool(requestcache))

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *Search) Routing(routing string) *Search {
	r.values.Set("routing", routing)

	return r
}

// Scroll Period to retain the search context for scrolling. See Scroll search results.
// By default, this value cannot exceed `1d` (24 hours).
// You can change this limit using the `search.max_keep_alive` cluster-level
// setting.
// API name: scroll
func (r *Search) Scroll(duration string) *Search {
	r.values.Set("scroll", duration)

	return r
}

// SearchType How distributed term frequencies are calculated for relevance scoring.
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

// SuggestMode Specifies the suggest mode.
// This parameter can only be used when the `suggest_field` and `suggest_text`
// query string parameters are specified.
// API name: suggest_mode
func (r *Search) SuggestMode(suggestmode suggestmode.SuggestMode) *Search {
	r.values.Set("suggest_mode", suggestmode.String())

	return r
}

// SuggestSize Number of suggestions to return.
// This parameter can only be used when the `suggest_field` and `suggest_text`
// query string parameters are specified.
// API name: suggest_size
func (r *Search) SuggestSize(suggestsize string) *Search {
	r.values.Set("suggest_size", suggestsize)

	return r
}

// SuggestText The source text for which the suggestions should be returned.
// This parameter can only be used when the `suggest_field` and `suggest_text`
// query string parameters are specified.
// API name: suggest_text
func (r *Search) SuggestText(suggesttext string) *Search {
	r.values.Set("suggest_text", suggesttext)

	return r
}

// TypedKeys If `true`, aggregation and suggester names are be prefixed by their
// respective types in the response.
// API name: typed_keys
func (r *Search) TypedKeys(typedkeys bool) *Search {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// RestTotalHitsAsInt Indicates whether `hits.total` should be rendered as an integer or an object
// in the rest search response.
// API name: rest_total_hits_as_int
func (r *Search) RestTotalHitsAsInt(resttotalhitsasint bool) *Search {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// SourceExcludes_ A comma-separated list of source fields to exclude from the response.
// You can also use this parameter to exclude fields from the subset specified
// in `_source_includes` query parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_excludes
func (r *Search) SourceExcludes_(fields ...string) *Search {
	r.values.Set("_source_excludes", strings.Join(fields, ","))

	return r
}

// SourceIncludes_ A comma-separated list of source fields to include in the response.
// If this parameter is specified, only these source fields are returned.
// You can exclude fields from this subset using the `_source_excludes` query
// parameter.
// If the `_source` parameter is `false`, this parameter is ignored.
// API name: _source_includes
func (r *Search) SourceIncludes_(fields ...string) *Search {
	r.values.Set("_source_includes", strings.Join(fields, ","))

	return r
}

// Q Query in the Lucene query string syntax using query parameter search.
// Query parameter searches do not support the full Elasticsearch Query DSL but
// are handy for testing.
// API name: q
func (r *Search) Q(q string) *Search {
	r.values.Set("q", q)

	return r
}

// Aggregations Defines the aggregations that are run as part of the search request.
// API name: aggregations
func (r *Search) Aggregations(aggregations map[string]types.Aggregations) *Search {

	r.req.Aggregations = aggregations

	return r
}

// Collapse Collapses search results the values of the specified field.
// API name: collapse
func (r *Search) Collapse(collapse *types.FieldCollapse) *Search {

	r.req.Collapse = collapse

	return r
}

// DocvalueFields Array of wildcard (`*`) patterns.
// The request returns doc values for field names matching these patterns in the
// `hits.fields` property of the response.
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

// Fields Array of wildcard (`*`) patterns.
// The request returns values for field names matching these patterns in the
// `hits.fields` property of the response.
// API name: fields
func (r *Search) Fields(fields ...types.FieldAndFormat) *Search {
	r.req.Fields = fields

	return r
}

// From Starting document offset.
// Needs to be non-negative.
// By default, you cannot page through more than 10,000 hits using the `from`
// and `size` parameters.
// To page through more hits, use the `search_after` parameter.
// API name: from
func (r *Search) From(from int) *Search {
	r.req.From = &from

	return r
}

// Highlight Specifies the highlighter to use for retrieving highlighted snippets from one
// or more fields in your search results.
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

// MinScore Minimum `_score` for matching documents.
// Documents with a lower `_score` are not included in the search results.
// API name: min_score
func (r *Search) MinScore(minscore types.Float64) *Search {

	r.req.MinScore = &minscore

	return r
}

// Pit Limits the search to a point in time (PIT).
// If you provide a PIT, you cannot specify an `<index>` in the request path.
// API name: pit
func (r *Search) Pit(pit *types.PointInTimeReference) *Search {

	r.req.Pit = pit

	return r
}

// PostFilter Use the `post_filter` parameter to filter search results.
// The search hits are filtered after the aggregations are calculated.
// A post filter has no impact on the aggregation results.
// API name: post_filter
func (r *Search) PostFilter(postfilter *types.Query) *Search {

	r.req.PostFilter = postfilter

	return r
}

// Profile Set to `true` to return detailed timing information about the execution of
// individual components in a search request.
// NOTE: This is a debugging tool and adds significant overhead to search
// execution.
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

// Rank Defines the Reciprocal Rank Fusion (RRF) to use.
// API name: rank
func (r *Search) Rank(rank *types.RankContainer) *Search {

	r.req.Rank = rank

	return r
}

// Rescore Can be used to improve precision by reordering just the top (for example 100
// - 500) documents returned by the `query` and `post_filter` phases.
// API name: rescore
func (r *Search) Rescore(rescores ...types.Rescore) *Search {
	r.req.Rescore = rescores

	return r
}

// RuntimeMappings Defines one or more runtime fields in the search request.
// These fields take precedence over mapped fields with the same name.
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

// SearchAfter Used to retrieve the next page of hits using a set of sort values from the
// previous page.
// API name: search_after
func (r *Search) SearchAfter(sortresults ...types.FieldValue) *Search {
	r.req.SearchAfter = sortresults

	return r
}

// SeqNoPrimaryTerm If `true`, returns sequence number and primary term of the last modification
// of each hit.
// API name: seq_no_primary_term
func (r *Search) SeqNoPrimaryTerm(seqnoprimaryterm bool) *Search {
	r.req.SeqNoPrimaryTerm = &seqnoprimaryterm

	return r
}

// Size The number of hits to return.
// By default, you cannot page through more than 10,000 hits using the `from`
// and `size` parameters.
// To page through more hits, use the `search_after` parameter.
// API name: size
func (r *Search) Size(size int) *Search {
	r.req.Size = &size

	return r
}

// Slice Can be used to split a scrolled search into multiple slices that can be
// consumed independently.
// API name: slice
func (r *Search) Slice(slice *types.SlicedScroll) *Search {

	r.req.Slice = slice

	return r
}

// Sort A comma-separated list of <field>:<direction> pairs.
// API name: sort
func (r *Search) Sort(sorts ...types.SortCombinations) *Search {
	r.req.Sort = sorts

	return r
}

// Source_ Indicates which source fields are returned for matching documents.
// These fields are returned in the hits._source property of the search
// response.
// API name: _source
func (r *Search) Source_(sourceconfig types.SourceConfig) *Search {
	r.req.Source_ = sourceconfig

	return r
}

// Stats Stats groups to associate with the search.
// Each group maintains a statistics aggregation for its associated searches.
// You can retrieve these stats using the indices stats API.
// API name: stats
func (r *Search) Stats(stats ...string) *Search {
	r.req.Stats = stats

	return r
}

// StoredFields List of stored fields to return as part of a hit.
// If no fields are specified, no stored fields are included in the response.
// If this field is specified, the `_source` parameter defaults to `false`.
// You can pass `_source: true` to return both source fields and stored fields
// in the search response.
// API name: stored_fields
func (r *Search) StoredFields(fields ...string) *Search {
	r.req.StoredFields = fields

	return r
}

// Suggest Defines a suggester that provides similar looking terms based on a provided
// text.
// API name: suggest
func (r *Search) Suggest(suggest *types.Suggester) *Search {

	r.req.Suggest = suggest

	return r
}

// TerminateAfter Maximum number of documents to collect for each shard.
// If a query reaches this limit, Elasticsearch terminates the query early.
// Elasticsearch collects documents before sorting.
// Use with caution.
// Elasticsearch applies this parameter to each shard handling the request.
// When possible, let Elasticsearch perform early termination automatically.
// Avoid specifying this parameter for requests that target data streams with
// backing indices across multiple data tiers.
// If set to `0` (default), the query does not terminate early.
// API name: terminate_after
func (r *Search) TerminateAfter(terminateafter int64) *Search {

	r.req.TerminateAfter = &terminateafter

	return r
}

// Timeout Specifies the period of time to wait for a response from each shard.
// If no response is received before the timeout expires, the request fails and
// returns an error.
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

// TrackTotalHits Number of hits matching the query to count accurately.
// If `true`, the exact number of hits is returned at the cost of some
// performance.
// If `false`, the  response does not include the total number of hits matching
// the query.
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
