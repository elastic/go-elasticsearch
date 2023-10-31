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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Executes a search request asynchronously.
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

type Submit struct {
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

// Executes a search request asynchronously.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/async-search.html
func New(tp elastictransport.Interface) *Submit {
	r := &Submit{
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Submit: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_async_search")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

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
func (r Submit) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the Submit query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a submit.Response
func (r Submit) Do(ctx context.Context) (*Response, error) {

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

// KeepOnCompletion If `true`, results are stored for later retrieval when the search completes
// within the `wait_for_completion_timeout`.
// API name: keep_on_completion
func (r *Submit) KeepOnCompletion(keeponcompletion bool) *Submit {
	r.values.Set("keep_on_completion", strconv.FormatBool(keeponcompletion))

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
func (r *Submit) MaxConcurrentShardRequests(maxconcurrentshardrequests string) *Submit {
	r.values.Set("max_concurrent_shard_requests", maxconcurrentshardrequests)

	return r
}

// API name: min_compatible_shard_node
func (r *Submit) MinCompatibleShardNode(versionstring string) *Submit {
	r.values.Set("min_compatible_shard_node", versionstring)

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *Submit) Preference(preference string) *Submit {
	r.values.Set("preference", preference)

	return r
}

// PreFilterShardSize The default value cannot be changed, which enforces the execution of a
// pre-filter roundtrip to retrieve statistics from each shard so that the ones
// that surely don’t hold any document matching the query get skipped.
// API name: pre_filter_shard_size
func (r *Submit) PreFilterShardSize(prefiltershardsize string) *Submit {
	r.values.Set("pre_filter_shard_size", prefiltershardsize)

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

// API name: scroll
func (r *Submit) Scroll(duration string) *Submit {
	r.values.Set("scroll", duration)

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

// API name: aggregations
func (r *Submit) Aggregations(aggregations map[string]types.Aggregations) *Submit {

	r.req.Aggregations = aggregations

	return r
}

// API name: collapse
func (r *Submit) Collapse(collapse *types.FieldCollapse) *Submit {

	r.req.Collapse = collapse

	return r
}

// DocvalueFields Array of wildcard (*) patterns. The request returns doc values for field
// names matching these patterns in the hits.fields property of the response.
// API name: docvalue_fields
func (r *Submit) DocvalueFields(docvaluefields ...types.FieldAndFormat) *Submit {
	r.req.DocvalueFields = docvaluefields

	return r
}

// Explain If true, returns detailed information about score computation as part of a
// hit.
// API name: explain
func (r *Submit) Explain(explain bool) *Submit {
	r.req.Explain = &explain

	return r
}

// Ext Configuration of search extensions defined by Elasticsearch plugins.
// API name: ext
func (r *Submit) Ext(ext map[string]json.RawMessage) *Submit {

	r.req.Ext = ext

	return r
}

// Fields Array of wildcard (*) patterns. The request returns values for field names
// matching these patterns in the hits.fields property of the response.
// API name: fields
func (r *Submit) Fields(fields ...types.FieldAndFormat) *Submit {
	r.req.Fields = fields

	return r
}

// From Starting document offset. By default, you cannot page through more than
// 10,000
// hits using the from and size parameters. To page through more hits, use the
// search_after parameter.
// API name: from
func (r *Submit) From(from int) *Submit {
	r.req.From = &from

	return r
}

// API name: highlight
func (r *Submit) Highlight(highlight *types.Highlight) *Submit {

	r.req.Highlight = highlight

	return r
}

// IndicesBoost Boosts the _score of documents from specified indices.
// API name: indices_boost
func (r *Submit) IndicesBoost(indicesboosts ...map[string]types.Float64) *Submit {
	r.req.IndicesBoost = indicesboosts

	return r
}

// Knn Defines the approximate kNN search to run.
// API name: knn
func (r *Submit) Knn(knns ...types.KnnQuery) *Submit {
	r.req.Knn = knns

	return r
}

// MinScore Minimum _score for matching documents. Documents with a lower _score are
// not included in the search results.
// API name: min_score
func (r *Submit) MinScore(minscore types.Float64) *Submit {

	r.req.MinScore = &minscore

	return r
}

// Pit Limits the search to a point in time (PIT). If you provide a PIT, you
// cannot specify an <index> in the request path.
// API name: pit
func (r *Submit) Pit(pit *types.PointInTimeReference) *Submit {

	r.req.Pit = pit

	return r
}

// API name: post_filter
func (r *Submit) PostFilter(postfilter *types.Query) *Submit {

	r.req.PostFilter = postfilter

	return r
}

// API name: profile
func (r *Submit) Profile(profile bool) *Submit {
	r.req.Profile = &profile

	return r
}

// Query Defines the search definition using the Query DSL.
// API name: query
func (r *Submit) Query(query *types.Query) *Submit {

	r.req.Query = query

	return r
}

// API name: rescore
func (r *Submit) Rescore(rescores ...types.Rescore) *Submit {
	r.req.Rescore = rescores

	return r
}

// RuntimeMappings Defines one or more runtime fields in the search request. These fields take
// precedence over mapped fields with the same name.
// API name: runtime_mappings
func (r *Submit) RuntimeMappings(runtimefields types.RuntimeFields) *Submit {
	r.req.RuntimeMappings = runtimefields

	return r
}

// ScriptFields Retrieve a script evaluation (based on different fields) for each hit.
// API name: script_fields
func (r *Submit) ScriptFields(scriptfields map[string]types.ScriptField) *Submit {

	r.req.ScriptFields = scriptfields

	return r
}

// API name: search_after
func (r *Submit) SearchAfter(sortresults ...types.FieldValue) *Submit {
	r.req.SearchAfter = sortresults

	return r
}

// SeqNoPrimaryTerm If true, returns sequence number and primary term of the last modification
// of each hit. See Optimistic concurrency control.
// API name: seq_no_primary_term
func (r *Submit) SeqNoPrimaryTerm(seqnoprimaryterm bool) *Submit {
	r.req.SeqNoPrimaryTerm = &seqnoprimaryterm

	return r
}

// Size The number of hits to return. By default, you cannot page through more
// than 10,000 hits using the from and size parameters. To page through more
// hits, use the search_after parameter.
// API name: size
func (r *Submit) Size(size int) *Submit {
	r.req.Size = &size

	return r
}

// API name: slice
func (r *Submit) Slice(slice *types.SlicedScroll) *Submit {

	r.req.Slice = slice

	return r
}

// API name: sort
func (r *Submit) Sort(sorts ...types.SortCombinations) *Submit {
	r.req.Sort = sorts

	return r
}

// Source_ Indicates which source fields are returned for matching documents. These
// fields are returned in the hits._source property of the search response.
// API name: _source
func (r *Submit) Source_(sourceconfig types.SourceConfig) *Submit {
	r.req.Source_ = sourceconfig

	return r
}

// Stats Stats groups to associate with the search. Each group maintains a statistics
// aggregation for its associated searches. You can retrieve these stats using
// the indices stats API.
// API name: stats
func (r *Submit) Stats(stats ...string) *Submit {
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
func (r *Submit) StoredFields(fields ...string) *Submit {
	r.req.StoredFields = fields

	return r
}

// API name: suggest
func (r *Submit) Suggest(suggest *types.Suggester) *Submit {

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
func (r *Submit) TerminateAfter(terminateafter int64) *Submit {

	r.req.TerminateAfter = &terminateafter

	return r
}

// Timeout Specifies the period of time to wait for a response from each shard. If no
// response
// is received before the timeout expires, the request fails and returns an
// error.
// Defaults to no timeout.
// API name: timeout
func (r *Submit) Timeout(timeout string) *Submit {

	r.req.Timeout = &timeout

	return r
}

// TrackScores If true, calculate and return document scores, even if the scores are not
// used for sorting.
// API name: track_scores
func (r *Submit) TrackScores(trackscores bool) *Submit {
	r.req.TrackScores = &trackscores

	return r
}

// TrackTotalHits Number of hits matching the query to count accurately. If true, the exact
// number of hits is returned at the cost of some performance. If false, the
// response does not include the total number of hits matching the query.
// Defaults to 10,000 hits.
// API name: track_total_hits
func (r *Submit) TrackTotalHits(trackhits types.TrackHits) *Submit {
	r.req.TrackTotalHits = trackhits

	return r
}

// Version If true, returns document version as part of a hit.
// API name: version
func (r *Submit) Version(version bool) *Submit {
	r.req.Version = &version

	return r
}
