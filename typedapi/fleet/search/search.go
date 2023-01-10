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


// Search API where the search will only be executed after specified checkpoints
// are available due to a refresh. This API is designed for internal use by the
// fleet server project.
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
type NewSearch func(index string) *Search

// NewSearchFunc returns a new instance of Search with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSearchFunc(tp elastictransport.Interface) NewSearch {
	return func(index string) *Search {
		n := New(tp)

		n.Index(index)

		return n
	}
}

// Search API where the search will only be executed after specified checkpoints
// are available due to a refresh. This API is designed for internal use by the
// fleet server project.
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
	case r.paramSet == indexMask:
		path.WriteString("/")

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

// Index A single target to search. If the target is an index alias, it must resolve
// to a single index.
// API Name: index
func (r *Search) Index(v string) *Search {
	r.paramSet |= indexMask
	r.index = v

	return r
}

// API name: allow_no_indices
func (r *Search) AllowNoIndices(b bool) *Search {
	r.values.Set("allow_no_indices", strconv.FormatBool(b))

	return r
}

// API name: analyzer
func (r *Search) Analyzer(value string) *Search {
	r.values.Set("analyzer", value)

	return r
}

// API name: analyze_wildcard
func (r *Search) AnalyzeWildcard(b bool) *Search {
	r.values.Set("analyze_wildcard", strconv.FormatBool(b))

	return r
}

// API name: batched_reduce_size
func (r *Search) BatchedReduceSize(value string) *Search {
	r.values.Set("batched_reduce_size", value)

	return r
}

// API name: ccs_minimize_roundtrips
func (r *Search) CcsMinimizeRoundtrips(b bool) *Search {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(b))

	return r
}

// API name: default_operator
func (r *Search) DefaultOperator(enum operator.Operator) *Search {
	r.values.Set("default_operator", enum.String())

	return r
}

// API name: df
func (r *Search) Df(value string) *Search {
	r.values.Set("df", value)

	return r
}

// API name: docvalue_fields
func (r *Search) DocvalueFields(value string) *Search {
	r.values.Set("docvalue_fields", value)

	return r
}

// API name: expand_wildcards
func (r *Search) ExpandWildcards(value string) *Search {
	r.values.Set("expand_wildcards", value)

	return r
}

// API name: explain
func (r *Search) Explain(b bool) *Search {
	r.values.Set("explain", strconv.FormatBool(b))

	return r
}

// API name: ignore_throttled
func (r *Search) IgnoreThrottled(b bool) *Search {
	r.values.Set("ignore_throttled", strconv.FormatBool(b))

	return r
}

// API name: ignore_unavailable
func (r *Search) IgnoreUnavailable(b bool) *Search {
	r.values.Set("ignore_unavailable", strconv.FormatBool(b))

	return r
}

// API name: lenient
func (r *Search) Lenient(b bool) *Search {
	r.values.Set("lenient", strconv.FormatBool(b))

	return r
}

// API name: max_concurrent_shard_requests
func (r *Search) MaxConcurrentShardRequests(value string) *Search {
	r.values.Set("max_concurrent_shard_requests", value)

	return r
}

// API name: min_compatible_shard_node
func (r *Search) MinCompatibleShardNode(value string) *Search {
	r.values.Set("min_compatible_shard_node", value)

	return r
}

// API name: preference
func (r *Search) Preference(value string) *Search {
	r.values.Set("preference", value)

	return r
}

// API name: pre_filter_shard_size
func (r *Search) PreFilterShardSize(value string) *Search {
	r.values.Set("pre_filter_shard_size", value)

	return r
}

// API name: request_cache
func (r *Search) RequestCache(b bool) *Search {
	r.values.Set("request_cache", strconv.FormatBool(b))

	return r
}

// API name: routing
func (r *Search) Routing(value string) *Search {
	r.values.Set("routing", value)

	return r
}

// API name: scroll
func (r *Search) Scroll(value string) *Search {
	r.values.Set("scroll", value)

	return r
}

// API name: search_type
func (r *Search) SearchType(enum searchtype.SearchType) *Search {
	r.values.Set("search_type", enum.String())

	return r
}

// API name: stats
func (r *Search) Stats(value string) *Search {
	r.values.Set("stats", value)

	return r
}

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

// API name: suggest_mode
func (r *Search) SuggestMode(enum suggestmode.SuggestMode) *Search {
	r.values.Set("suggest_mode", enum.String())

	return r
}

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

// API name: terminate_after
func (r *Search) TerminateAfter(value string) *Search {
	r.values.Set("terminate_after", value)

	return r
}

// API name: timeout
func (r *Search) Timeout(value string) *Search {
	r.values.Set("timeout", value)

	return r
}

// API name: track_total_hits
func (r *Search) TrackTotalHits(value string) *Search {
	r.values.Set("track_total_hits", value)

	return r
}

// API name: track_scores
func (r *Search) TrackScores(b bool) *Search {
	r.values.Set("track_scores", strconv.FormatBool(b))

	return r
}

// API name: typed_keys
func (r *Search) TypedKeys(b bool) *Search {
	r.values.Set("typed_keys", strconv.FormatBool(b))

	return r
}

// API name: rest_total_hits_as_int
func (r *Search) RestTotalHitsAsInt(b bool) *Search {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(b))

	return r
}

// API name: version
func (r *Search) Version(b bool) *Search {
	r.values.Set("version", strconv.FormatBool(b))

	return r
}

// API name: _source
func (r *Search) Source_(value string) *Search {
	r.values.Set("_source", value)

	return r
}

// API name: _source_excludes
func (r *Search) SourceExcludes_(value string) *Search {
	r.values.Set("_source_excludes", value)

	return r
}

// API name: _source_includes
func (r *Search) SourceIncludes_(value string) *Search {
	r.values.Set("_source_includes", value)

	return r
}

// API name: seq_no_primary_term
func (r *Search) SeqNoPrimaryTerm(b bool) *Search {
	r.values.Set("seq_no_primary_term", strconv.FormatBool(b))

	return r
}

// API name: q
func (r *Search) Q(value string) *Search {
	r.values.Set("q", value)

	return r
}

// API name: size
func (r *Search) Size(i int) *Search {
	r.values.Set("size", strconv.Itoa(i))

	return r
}

// API name: from
func (r *Search) From(i int) *Search {
	r.values.Set("from", strconv.Itoa(i))

	return r
}

// API name: sort
func (r *Search) Sort(value string) *Search {
	r.values.Set("sort", value)

	return r
}

// WaitForCheckpoints A comma separated list of checkpoints. When configured, the search API will
// only be executed on a shard
// after the relevant checkpoint has become visible for search. Defaults to an
// empty list which will cause
// Elasticsearch to immediately execute the search.
// API name: wait_for_checkpoints
func (r *Search) WaitForCheckpoints(value string) *Search {
	r.values.Set("wait_for_checkpoints", value)

	return r
}

// AllowPartialSearchResults If true, returns partial results if there are shard request timeouts or
// [shard
// failures](https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-replication.html#shard-failures).
// If false, returns
// an error with no partial results. Defaults to the configured cluster setting
// `search.default_allow_partial_results`
// which is true by default.
// API name: allow_partial_search_results
func (r *Search) AllowPartialSearchResults(b bool) *Search {
	r.values.Set("allow_partial_search_results", strconv.FormatBool(b))

	return r
}
