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

// Deletes documents matching the provided query.
package deletebyquery

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/conflicts"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type DeleteByQuery struct {
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

// NewDeleteByQuery type alias for index.
type NewDeleteByQuery func(index string) *DeleteByQuery

// NewDeleteByQueryFunc returns a new instance of DeleteByQuery with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewDeleteByQueryFunc(tp elastictransport.Interface) NewDeleteByQuery {
	return func(index string) *DeleteByQuery {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Deletes documents matching the provided query.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/docs-delete-by-query.html
func New(tp elastictransport.Interface) *DeleteByQuery {
	r := &DeleteByQuery{
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
func (r *DeleteByQuery) Raw(raw io.Reader) *DeleteByQuery {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *DeleteByQuery) Request(req *Request) *DeleteByQuery {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *DeleteByQuery) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for DeleteByQuery: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_delete_by_query")

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
func (r DeleteByQuery) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the DeleteByQuery query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deletebyquery.Response
func (r DeleteByQuery) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

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

// Header set a key, value pair in the DeleteByQuery headers map.
func (r *DeleteByQuery) Header(key, value string) *DeleteByQuery {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases to search.
// Supports wildcards (`*`).
// To search all data streams or indices, omit this parameter or use `*` or
// `_all`.
// API Name: index
func (r *DeleteByQuery) _index(index string) *DeleteByQuery {
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
func (r *DeleteByQuery) AllowNoIndices(allownoindices bool) *DeleteByQuery {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// Analyzer Analyzer to use for the query string.
// API name: analyzer
func (r *DeleteByQuery) Analyzer(analyzer string) *DeleteByQuery {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard If `true`, wildcard and prefix queries are analyzed.
// API name: analyze_wildcard
func (r *DeleteByQuery) AnalyzeWildcard(analyzewildcard bool) *DeleteByQuery {
	r.values.Set("analyze_wildcard", strconv.FormatBool(analyzewildcard))

	return r
}

// Conflicts What to do if delete by query hits version conflicts: `abort` or `proceed`.
// API name: conflicts
func (r *DeleteByQuery) Conflicts(conflicts conflicts.Conflicts) *DeleteByQuery {
	r.values.Set("conflicts", conflicts.String())

	return r
}

// DefaultOperator The default operator for query string query: `AND` or `OR`.
// API name: default_operator
func (r *DeleteByQuery) DefaultOperator(defaultoperator operator.Operator) *DeleteByQuery {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df Field to use as default where no field prefix is given in the query string.
// API name: df
func (r *DeleteByQuery) Df(df string) *DeleteByQuery {
	r.values.Set("df", df)

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`. Valid values are:
// `all`, `open`, `closed`, `hidden`, `none`.
// API name: expand_wildcards
func (r *DeleteByQuery) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DeleteByQuery {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// From Starting offset (default: 0)
// API name: from
func (r *DeleteByQuery) From(from string) *DeleteByQuery {
	r.values.Set("from", from)

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *DeleteByQuery) IgnoreUnavailable(ignoreunavailable bool) *DeleteByQuery {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Lenient If `true`, format-based query failures (such as providing text to a numeric
// field) in the query string will be ignored.
// API name: lenient
func (r *DeleteByQuery) Lenient(lenient bool) *DeleteByQuery {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// Preference Specifies the node or shard the operation should be performed on.
// Random by default.
// API name: preference
func (r *DeleteByQuery) Preference(preference string) *DeleteByQuery {
	r.values.Set("preference", preference)

	return r
}

// Refresh If `true`, Elasticsearch refreshes all shards involved in the delete by query
// after the request completes.
// API name: refresh
func (r *DeleteByQuery) Refresh(refresh bool) *DeleteByQuery {
	r.values.Set("refresh", strconv.FormatBool(refresh))

	return r
}

// RequestCache If `true`, the request cache is used for this request.
// Defaults to the index-level setting.
// API name: request_cache
func (r *DeleteByQuery) RequestCache(requestcache bool) *DeleteByQuery {
	r.values.Set("request_cache", strconv.FormatBool(requestcache))

	return r
}

// RequestsPerSecond The throttle for this request in sub-requests per second.
// API name: requests_per_second
func (r *DeleteByQuery) RequestsPerSecond(requestspersecond string) *DeleteByQuery {
	r.values.Set("requests_per_second", requestspersecond)

	return r
}

// Routing Custom value used to route operations to a specific shard.
// API name: routing
func (r *DeleteByQuery) Routing(routing string) *DeleteByQuery {
	r.values.Set("routing", routing)

	return r
}

// Q Query in the Lucene query string syntax.
// API name: q
func (r *DeleteByQuery) Q(q string) *DeleteByQuery {
	r.values.Set("q", q)

	return r
}

// Scroll Period to retain the search context for scrolling.
// API name: scroll
func (r *DeleteByQuery) Scroll(duration string) *DeleteByQuery {
	r.values.Set("scroll", duration)

	return r
}

// ScrollSize Size of the scroll request that powers the operation.
// API name: scroll_size
func (r *DeleteByQuery) ScrollSize(scrollsize string) *DeleteByQuery {
	r.values.Set("scroll_size", scrollsize)

	return r
}

// SearchTimeout Explicit timeout for each search request.
// Defaults to no timeout.
// API name: search_timeout
func (r *DeleteByQuery) SearchTimeout(duration string) *DeleteByQuery {
	r.values.Set("search_timeout", duration)

	return r
}

// SearchType The type of the search operation.
// Available options: `query_then_fetch`, `dfs_query_then_fetch`.
// API name: search_type
func (r *DeleteByQuery) SearchType(searchtype searchtype.SearchType) *DeleteByQuery {
	r.values.Set("search_type", searchtype.String())

	return r
}

// Slices The number of slices this task should be divided into.
// API name: slices
func (r *DeleteByQuery) Slices(slices string) *DeleteByQuery {
	r.values.Set("slices", slices)

	return r
}

// Sort A comma-separated list of <field>:<direction> pairs.
// API name: sort
func (r *DeleteByQuery) Sort(sorts ...string) *DeleteByQuery {
	tmp := []string{}
	for _, item := range sorts {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("sort", strings.Join(tmp, ","))

	return r
}

// Stats Specific `tag` of the request for logging and statistical purposes.
// API name: stats
func (r *DeleteByQuery) Stats(stats ...string) *DeleteByQuery {
	tmp := []string{}
	for _, item := range stats {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("stats", strings.Join(tmp, ","))

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
// API name: terminate_after
func (r *DeleteByQuery) TerminateAfter(terminateafter string) *DeleteByQuery {
	r.values.Set("terminate_after", terminateafter)

	return r
}

// Timeout Period each deletion request waits for active shards.
// API name: timeout
func (r *DeleteByQuery) Timeout(duration string) *DeleteByQuery {
	r.values.Set("timeout", duration)

	return r
}

// Version If `true`, returns the document version as part of a hit.
// API name: version
func (r *DeleteByQuery) Version(version bool) *DeleteByQuery {
	r.values.Set("version", strconv.FormatBool(version))

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to all or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *DeleteByQuery) WaitForActiveShards(waitforactiveshards string) *DeleteByQuery {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// WaitForCompletion If `true`, the request blocks until the operation is complete.
// API name: wait_for_completion
func (r *DeleteByQuery) WaitForCompletion(waitforcompletion bool) *DeleteByQuery {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// MaxDocs The maximum number of documents to delete.
// API name: max_docs
func (r *DeleteByQuery) MaxDocs(maxdocs int64) *DeleteByQuery {

	r.req.MaxDocs = &maxdocs

	return r
}

// Query Specifies the documents to delete using the Query DSL.
// API name: query
func (r *DeleteByQuery) Query(query *types.Query) *DeleteByQuery {

	r.req.Query = query

	return r
}

// Slice Slice the request manually using the provided slice ID and total number of
// slices.
// API name: slice
func (r *DeleteByQuery) Slice(slice *types.SlicedScroll) *DeleteByQuery {

	r.req.Slice = slice

	return r
}
