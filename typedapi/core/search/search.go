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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Returns results matching a query.
package search

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"

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

	req *Request
	raw json.RawMessage

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
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *Search) Raw(raw json.RawMessage) *Search {
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

	if r.raw != nil {
		r.buf.Write(r.raw)
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

// Do runs the http.Request through the provided transport.
func (r Search) Do(ctx context.Context) (*http.Response, error) {
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

// Header set a key, value pair in the Search headers map.
func (r *Search) Header(key, value string) *Search {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to search; use `_all` or empty string
// to perform the operation on all indices
// API Name: index
func (r *Search) Index(v string) *Search {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *Search) AllowNoIndices(b bool) *Search {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// AllowPartialSearchResults Indicate if an error should be returned if there is a partial search failure
// or timeout
// API name: allow_partial_search_results
func (r *Search) AllowPartialSearchResults(b bool) *Search {
	r.values.Set("allow_partial_search_results", strconv.FormatBool(b))

	return r
}

// Analyzer The analyzer to use for the query string
// API name: analyzer
func (r *Search) Analyzer(value string) *Search {
	r.values.Set("analyzer", value)

	return r
}

// AnalyzeWildcard Specify whether wildcard and prefix queries should be analyzed (default:
// false)
// API name: analyze_wildcard
func (r *Search) AnalyzeWildcard(b bool) *Search {
	r.values.Set("analyze_wildcard", strconv.FormatBool(b))

	return r
}

// BatchedReduceSize The number of shard results that should be reduced at once on the
// coordinating node. This value should be used as a protection mechanism to
// reduce the memory overhead per search request if the potential number of
// shards in the request can be large.
// API name: batched_reduce_size
func (r *Search) BatchedReduceSize(value string) *Search {
	r.values.Set("batched_reduce_size", value)

	return r
}

// CcsMinimizeRoundtrips Indicates whether network round-trips should be minimized as part of
// cross-cluster search requests execution
// API name: ccs_minimize_roundtrips
func (r *Search) CcsMinimizeRoundtrips(b bool) *Search {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(b))

	return r
}

// DefaultOperator The default operator for query string query (AND or OR)
// API name: default_operator
func (r *Search) DefaultOperator(enum operator.Operator) *Search {
	r.values.Set("default_operator", enum.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string
// API name: df
func (r *Search) Df(value string) *Search {
	r.values.Set("df", value)

	return r
}

// DocvalueFields A comma-separated list of fields to return as the docvalue representation of
// a field for each hit
// API name: docvalue_fields
func (r *Search) DocvalueFields(value string) *Search {
	r.values.Set("docvalue_fields", value)

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *Search) ExpandWildcards(value string) *Search {
	r.values.Set("expand_wildcards", value)

	return r
}

// Explain Specify whether to return detailed information about score computation as
// part of a hit
// API name: explain
func (r *Search) Explain(b bool) *Search {
	r.values.Set("explain", strconv.FormatBool(b))

	return r
}

// IgnoreThrottled Whether specified concrete, expanded or aliased indices should be ignored
// when throttled
// API name: ignore_throttled
func (r *Search) IgnoreThrottled(b bool) *Search {
	r.values.Set("ignore_throttled", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *Search) IgnoreUnavailable(b bool) *Search {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// Lenient Specify whether format-based query failures (such as providing text to a
// numeric field) should be ignored
// API name: lenient
func (r *Search) Lenient(b bool) *Search {
	r.values.Set("lenient", strconv.FormatBool(b))

	return r
}

// MaxConcurrentShardRequests The number of concurrent shard requests per node this search executes
// concurrently. This value should be used to limit the impact of the search on
// the cluster in order to limit the number of concurrent shard requests
// API name: max_concurrent_shard_requests
func (r *Search) MaxConcurrentShardRequests(value string) *Search {
	r.values.Set("max_concurrent_shard_requests", value)

	return r
}

// MinCompatibleShardNode The minimum compatible version that all shards involved in search should have
// for this request to be successful
// API name: min_compatible_shard_node
func (r *Search) MinCompatibleShardNode(value string) *Search {
	r.values.Set("min_compatible_shard_node", value)

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *Search) Preference(value string) *Search {
	r.values.Set("preference", value)

	return r
}

// PreFilterShardSize A threshold that enforces a pre-filter roundtrip to prefilter search shards
// based on query rewriting if the number of shards the search request expands
// to exceeds the threshold. This filter roundtrip can limit the number of
// shards significantly if for instance a shard can not match any documents
// based on its rewrite method ie. if date filters are mandatory to match but
// the shard bounds and the query are disjoint.
// API name: pre_filter_shard_size
func (r *Search) PreFilterShardSize(value string) *Search {
	r.values.Set("pre_filter_shard_size", value)

	return r
}

// RequestCache Specify if request cache should be used for this request or not, defaults to
// index level setting
// API name: request_cache
func (r *Search) RequestCache(b bool) *Search {
	r.values.Set("request_cache", strconv.FormatBool(b))

	return r
}

// Routing A comma-separated list of specific routing values
// API name: routing
func (r *Search) Routing(value string) *Search {
	r.values.Set("routing", value)

	return r
}

// Scroll Specify how long a consistent view of the index should be maintained for
// scrolled search
// API name: scroll
func (r *Search) Scroll(value string) *Search {
	r.values.Set("scroll", value)

	return r
}

// SearchType Search operation type
// API name: search_type
func (r *Search) SearchType(enum searchtype.SearchType) *Search {
	r.values.Set("search_type", enum.String())

	return r
}

// Stats Specific 'tag' of the request for logging and statistical purposes
// API name: stats
func (r *Search) Stats(value string) *Search {
	r.values.Set("stats", value)

	return r
}

// StoredFields A comma-separated list of stored fields to return as part of a hit
// API name: stored_fields
func (r *Search) StoredFields(value string) *Search {
	r.values.Set("stored_fields", value)

	return r
}

// SuggestField Specifies which field to use for suggestions.
// API name: suggest_field
func (r *Search) SuggestField(value string) *Search {
	r.values.Set("suggest_field", value)

	return r
}

// SuggestMode Specify suggest mode
// API name: suggest_mode
func (r *Search) SuggestMode(enum suggestmode.SuggestMode) *Search {
	r.values.Set("suggest_mode", enum.String())

	return r
}

// SuggestSize How many suggestions to return in response
// API name: suggest_size
func (r *Search) SuggestSize(value string) *Search {
	r.values.Set("suggest_size", value)

	return r
}

// SuggestText The source text for which the suggestions should be returned.
// API name: suggest_text
func (r *Search) SuggestText(value string) *Search {
	r.values.Set("suggest_text", value)

	return r
}

// TerminateAfter The maximum number of documents to collect for each shard, upon reaching
// which the query execution will terminate early.
// API name: terminate_after
func (r *Search) TerminateAfter(value string) *Search {
	r.values.Set("terminate_after", value)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Search) Timeout(value string) *Search {
	r.values.Set("timeout", value)

	return r
}

// TrackTotalHits Indicate if the number of documents that match the query should be tracked. A
// number can also be specified, to accurately track the total hit count up to
// the number.
// API name: track_total_hits
func (r *Search) TrackTotalHits(value string) *Search {
	r.values.Set("track_total_hits", value)

	return r
}

// TrackScores Whether to calculate and return scores even if they are not used for sorting
// API name: track_scores
func (r *Search) TrackScores(b bool) *Search {
	r.values.Set("track_scores", strconv.FormatBool(b))

	return r
}

// TypedKeys Specify whether aggregation and suggester names should be prefixed by their
// respective types in the response
// API name: typed_keys
func (r *Search) TypedKeys(b bool) *Search {
	r.values.Set("typed_keys", strconv.FormatBool(b))

	return r
}

// RestTotalHitsAsInt Indicates whether hits.total should be rendered as an integer or an object in
// the rest search response
// API name: rest_total_hits_as_int
func (r *Search) RestTotalHitsAsInt(b bool) *Search {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(b))

	return r
}

// Version Specify whether to return document version as part of a hit
// API name: version
func (r *Search) Version(b bool) *Search {
	r.values.Set("version", strconv.FormatBool(b))

	return r
}

// Source_ True or false to return the _source field or not, or a list of fields to
// return
// API name: _source
func (r *Search) Source_(value string) *Search {
	r.values.Set("_source", value)

	return r
}

// SourceExcludes_ A list of fields to exclude from the returned _source field
// API name: _source_excludes
func (r *Search) SourceExcludes_(value string) *Search {
	r.values.Set("_source_excludes", value)

	return r
}

// SourceIncludes_ A list of fields to extract and return from the _source field
// API name: _source_includes
func (r *Search) SourceIncludes_(value string) *Search {
	r.values.Set("_source_includes", value)

	return r
}

// SeqNoPrimaryTerm Specify whether to return sequence number and primary term of the last
// modification of each hit
// API name: seq_no_primary_term
func (r *Search) SeqNoPrimaryTerm(b bool) *Search {
	r.values.Set("seq_no_primary_term", strconv.FormatBool(b))

	return r
}

// Q Query in the Lucene query string syntax
// API name: q
func (r *Search) Q(value string) *Search {
	r.values.Set("q", value)

	return r
}

// Size Number of hits to return (default: 10)
// API name: size
func (r *Search) Size(i int) *Search {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// From Starting offset (default: 0)
// API name: from
func (r *Search) From(i int) *Search {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// Sort A comma-separated list of <field>:<direction> pairs
// API name: sort
func (r *Search) Sort(value string) *Search {
	r.values.Set("sort", value)

	return r
}
