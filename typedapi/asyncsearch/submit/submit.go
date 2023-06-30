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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

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

	req *Request
	raw io.Reader

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
func (r *Submit) Index(v string) *Submit {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// WaitForCompletionTimeout Specify the time that the request should block waiting for the final response
// API name: wait_for_completion_timeout
func (r *Submit) WaitForCompletionTimeout(v string) *Submit {
	r.values.Set("wait_for_completion_timeout", v)

	return r
}

// KeepOnCompletion Control whether the response should be stored in the cluster if it completed
// within the provided [wait_for_completion] time (default: false)
// API name: keep_on_completion
func (r *Submit) KeepOnCompletion(b bool) *Submit {
	r.values.Set("keep_on_completion", strconv.FormatBool(b))

	return r
}

// KeepAlive Update the time interval in which the results (partial or final) for this
// search will be available
// API name: keep_alive
func (r *Submit) KeepAlive(v string) *Submit {
	r.values.Set("keep_alive", v)

	return r
}

// AllowNoIndices Whether to ignore if a wildcard indices expression resolves into no concrete
// indices. (This includes `_all` string or when no indices have been specified)
// API name: allow_no_indices
func (r *Submit) AllowNoIndices(b bool) *Submit {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// AllowPartialSearchResults Indicate if an error should be returned if there is a partial search failure
// or timeout
// API name: allow_partial_search_results
func (r *Submit) AllowPartialSearchResults(b bool) *Submit {
	r.values.Set("allow_partial_search_results", strconv.FormatBool(b))

	return r
}

// Analyzer The analyzer to use for the query string
// API name: analyzer
func (r *Submit) Analyzer(v string) *Submit {
	r.values.Set("analyzer", v)

	return r
}

// AnalyzeWildcard Specify whether wildcard and prefix queries should be analyzed (default:
// false)
// API name: analyze_wildcard
func (r *Submit) AnalyzeWildcard(b bool) *Submit {
	r.values.Set("analyze_wildcard", strconv.FormatBool(b))

	return r
}

// BatchedReduceSize The number of shard results that should be reduced at once on the
// coordinating node. This value should be used as the granularity at which
// progress results will be made available.
// API name: batched_reduce_size
func (r *Submit) BatchedReduceSize(v string) *Submit {
	r.values.Set("batched_reduce_size", v)

	return r
}

// API name: ccs_minimize_roundtrips
func (r *Submit) CcsMinimizeRoundtrips(b bool) *Submit {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(b))

	return r
}

// DefaultOperator The default operator for query string query (AND or OR)
// API name: default_operator
func (r *Submit) DefaultOperator(enum operator.Operator) *Submit {
	r.values.Set("default_operator", enum.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string
// API name: df
func (r *Submit) Df(v string) *Submit {
	r.values.Set("df", v)

	return r
}

// DocvalueFields A comma-separated list of fields to return as the docvalue representation of
// a field for each hit
// API name: docvalue_fields
func (r *Submit) DocvalueFields(v string) *Submit {
	r.values.Set("docvalue_fields", v)

	return r
}

// ExpandWildcards Whether to expand wildcard expression to concrete indices that are open,
// closed or both.
// API name: expand_wildcards
func (r *Submit) ExpandWildcards(v string) *Submit {
	r.values.Set("expand_wildcards", v)

	return r
}

// Explain Specify whether to return detailed information about score computation as
// part of a hit
// API name: explain
func (r *Submit) Explain(b bool) *Submit {
	r.values.Set("explain", strconv.FormatBool(b))

	return r
}

// IgnoreThrottled Whether specified concrete, expanded or aliased indices should be ignored
// when throttled
// API name: ignore_throttled
func (r *Submit) IgnoreThrottled(b bool) *Submit {
	r.values.Set("ignore_throttled", strconv.FormatBool(b))

	return r
}

// IgnoreUnavailable Whether specified concrete indices should be ignored when unavailable
// (missing or closed)
// API name: ignore_unavailable
func (r *Submit) IgnoreUnavailable(b bool) *Submit {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// Lenient Specify whether format-based query failures (such as providing text to a
// numeric field) should be ignored
// API name: lenient
func (r *Submit) Lenient(b bool) *Submit {
	r.values.Set("lenient", strconv.FormatBool(b))

	return r
}

// MaxConcurrentShardRequests The number of concurrent shard requests per node this search executes
// concurrently. This value should be used to limit the impact of the search on
// the cluster in order to limit the number of concurrent shard requests
// API name: max_concurrent_shard_requests
func (r *Submit) MaxConcurrentShardRequests(v string) *Submit {
	r.values.Set("max_concurrent_shard_requests", v)

	return r
}

// API name: min_compatible_shard_node
func (r *Submit) MinCompatibleShardNode(v string) *Submit {
	r.values.Set("min_compatible_shard_node", v)

	return r
}

// Preference Specify the node or shard the operation should be performed on (default:
// random)
// API name: preference
func (r *Submit) Preference(v string) *Submit {
	r.values.Set("preference", v)

	return r
}

// API name: pre_filter_shard_size
func (r *Submit) PreFilterShardSize(v string) *Submit {
	r.values.Set("pre_filter_shard_size", v)

	return r
}

// RequestCache Specify if request cache should be used for this request or not, defaults to
// true
// API name: request_cache
func (r *Submit) RequestCache(b bool) *Submit {
	r.values.Set("request_cache", strconv.FormatBool(b))

	return r
}

// Routing A comma-separated list of specific routing values
// API name: routing
func (r *Submit) Routing(v string) *Submit {
	r.values.Set("routing", v)

	return r
}

// API name: scroll
func (r *Submit) Scroll(v string) *Submit {
	r.values.Set("scroll", v)

	return r
}

// SearchType Search operation type
// API name: search_type
func (r *Submit) SearchType(enum searchtype.SearchType) *Submit {
	r.values.Set("search_type", enum.String())

	return r
}

// Stats Specific 'tag' of the request for logging and statistical purposes
// API name: stats
func (r *Submit) Stats(v string) *Submit {
	r.values.Set("stats", v)

	return r
}

// StoredFields A comma-separated list of stored fields to return as part of a hit
// API name: stored_fields
func (r *Submit) StoredFields(v string) *Submit {
	r.values.Set("stored_fields", v)

	return r
}

// SuggestField Specifies which field to use for suggestions.
// API name: suggest_field
func (r *Submit) SuggestField(v string) *Submit {
	r.values.Set("suggest_field", v)

	return r
}

// SuggestMode Specify suggest mode
// API name: suggest_mode
func (r *Submit) SuggestMode(enum suggestmode.SuggestMode) *Submit {
	r.values.Set("suggest_mode", enum.String())

	return r
}

// SuggestSize How many suggestions to return in response
// API name: suggest_size
func (r *Submit) SuggestSize(v string) *Submit {
	r.values.Set("suggest_size", v)

	return r
}

// SuggestText The source text for which the suggestions should be returned.
// API name: suggest_text
func (r *Submit) SuggestText(v string) *Submit {
	r.values.Set("suggest_text", v)

	return r
}

// TerminateAfter The maximum number of documents to collect for each shard, upon reaching
// which the query execution will terminate early.
// API name: terminate_after
func (r *Submit) TerminateAfter(v string) *Submit {
	r.values.Set("terminate_after", v)

	return r
}

// Timeout Explicit operation timeout
// API name: timeout
func (r *Submit) Timeout(v string) *Submit {
	r.values.Set("timeout", v)

	return r
}

// TrackTotalHits Indicate if the number of documents that match the query should be tracked. A
// number can also be specified, to accurately track the total hit count up to
// the number.
// API name: track_total_hits
func (r *Submit) TrackTotalHits(v string) *Submit {
	r.values.Set("track_total_hits", v)

	return r
}

// TrackScores Whether to calculate and return scores even if they are not used for sorting
// API name: track_scores
func (r *Submit) TrackScores(b bool) *Submit {
	r.values.Set("track_scores", strconv.FormatBool(b))

	return r
}

// TypedKeys Specify whether aggregation and suggester names should be prefixed by their
// respective types in the response
// API name: typed_keys
func (r *Submit) TypedKeys(b bool) *Submit {
	r.values.Set("typed_keys", strconv.FormatBool(b))

	return r
}

// API name: rest_total_hits_as_int
func (r *Submit) RestTotalHitsAsInt(b bool) *Submit {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(b))

	return r
}

// Version Specify whether to return document version as part of a hit
// API name: version
func (r *Submit) Version(b bool) *Submit {
	r.values.Set("version", strconv.FormatBool(b))

	return r
}

// Source_ True or false to return the _source field or not, or a list of fields to
// return
// API name: _source
func (r *Submit) Source_(v string) *Submit {
	r.values.Set("_source", v)

	return r
}

// SourceExcludes_ A list of fields to exclude from the returned _source field
// API name: _source_excludes
func (r *Submit) SourceExcludes_(v string) *Submit {
	r.values.Set("_source_excludes", v)

	return r
}

// SourceIncludes_ A list of fields to extract and return from the _source field
// API name: _source_includes
func (r *Submit) SourceIncludes_(v string) *Submit {
	r.values.Set("_source_includes", v)

	return r
}

// SeqNoPrimaryTerm Specify whether to return sequence number and primary term of the last
// modification of each hit
// API name: seq_no_primary_term
func (r *Submit) SeqNoPrimaryTerm(b bool) *Submit {
	r.values.Set("seq_no_primary_term", strconv.FormatBool(b))

	return r
}

// Q Query in the Lucene query string syntax
// API name: q
func (r *Submit) Q(v string) *Submit {
	r.values.Set("q", v)

	return r
}

// Size Number of hits to return (default: 10)
// API name: size
func (r *Submit) Size(i int) *Submit {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// From Starting offset (default: 0)
// API name: from
func (r *Submit) From(i int) *Submit {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// Sort A comma-separated list of <field>:<direction> pairs
// API name: sort
func (r *Submit) Sort(v string) *Submit {
	r.values.Set("sort", v)

	return r
}
