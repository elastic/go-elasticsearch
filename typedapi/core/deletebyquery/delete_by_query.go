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

// Delete documents.
//
// Deletes documents that match the specified query.
//
// If the Elasticsearch security features are enabled, you must have the
// following index privileges for the target data stream, index, or alias:
//
// * `read`
// * `delete` or `write`
//
// You can specify the query criteria in the request URI or the request body
// using the same syntax as the search API.
// When you submit a delete by query request, Elasticsearch gets a snapshot of
// the data stream or index when it begins processing the request and deletes
// matching documents using internal versioning.
// If a document changes between the time that the snapshot is taken and the
// delete operation is processed, it results in a version conflict and the
// delete operation fails.
//
// NOTE: Documents with a version equal to 0 cannot be deleted using delete by
// query because internal versioning does not support 0 as a valid version
// number.
//
// While processing a delete by query request, Elasticsearch performs multiple
// search requests sequentially to find all of the matching documents to delete.
// A bulk delete request is performed for each batch of matching documents.
// If a search or bulk request is rejected, the requests are retried up to 10
// times, with exponential back off.
// If the maximum retry limit is reached, processing halts and all failed
// requests are returned in the response.
// Any delete requests that completed successfully still stick, they are not
// rolled back.
//
// You can opt to count version conflicts instead of halting and returning by
// setting `conflicts` to `proceed`.
// Note that if you opt to count version conflicts the operation could attempt
// to delete more documents from the source than `max_docs` until it has
// successfully deleted `max_docs documents`, or it has gone through every
// document in the source query.
//
// **Throttling delete requests**
//
// To control the rate at which delete by query issues batches of delete
// operations, you can set `requests_per_second` to any positive decimal number.
// This pads each batch with a wait time to throttle the rate.
// Set `requests_per_second` to `-1` to disable throttling.
//
// Throttling uses a wait time between batches so that the internal scroll
// requests can be given a timeout that takes the request padding into account.
// The padding time is the difference between the batch size divided by the
// `requests_per_second` and the time spent writing.
// By default the batch size is `1000`, so if `requests_per_second` is set to
// `500`:
//
// ```
// target_time = 1000 / 500 per second = 2 seconds
// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
// ```
//
// Since the batch is issued as a single `_bulk` request, large batch sizes
// cause Elasticsearch to create many requests and wait before starting the next
// set.
// This is "bursty" instead of "smooth".
//
// **Slicing**
//
// Delete by query supports sliced scroll to parallelize the delete process.
// This can improve efficiency and provide a convenient way to break the request
// down into smaller parts.
//
// Setting `slices` to `auto` lets Elasticsearch choose the number of slices to
// use.
// This setting will use one slice per shard, up to a certain limit.
// If there are multiple source data streams or indices, it will choose the
// number of slices based on the index or backing index with the smallest number
// of shards.
// Adding slices to the delete by query operation creates sub-requests which
// means it has some quirks:
//
// * You can see these requests in the tasks APIs. These sub-requests are
// "child" tasks of the task for the request with slices.
// * Fetching the status of the task for the request with slices only contains
// the status of completed slices.
// * These sub-requests are individually addressable for things like
// cancellation and rethrottling.
// * Rethrottling the request with `slices` will rethrottle the unfinished
// sub-request proportionally.
// * Canceling the request with `slices` will cancel each sub-request.
// * Due to the nature of `slices` each sub-request won't get a perfectly even
// portion of the documents. All documents will be addressed, but some slices
// may be larger than others. Expect larger slices to have a more even
// distribution.
// * Parameters like `requests_per_second` and `max_docs` on a request with
// `slices` are distributed proportionally to each sub-request. Combine that
// with the earlier point about distribution being uneven and you should
// conclude that using `max_docs` with `slices` might not result in exactly
// `max_docs` documents being deleted.
// * Each sub-request gets a slightly different snapshot of the source data
// stream or index though these are all taken at approximately the same time.
//
// If you're slicing manually or otherwise tuning automatic slicing, keep in
// mind that:
//
// * Query performance is most efficient when the number of slices is equal to
// the number of shards in the index or backing index. If that number is large
// (for example, 500), choose a lower number as too many `slices` hurts
// performance. Setting `slices` higher than the number of shards generally does
// not improve efficiency and adds overhead.
// * Delete performance scales linearly across available resources with the
// number of slices.
//
// Whether query or delete performance dominates the runtime depends on the
// documents being reindexed and cluster resources.
//
// **Cancel a delete by query operation**
//
// Any delete by query can be canceled using the task cancel API. For example:
//
// ```
// POST _tasks/r1A2WoRbTwKZ516z6NEs5A:36619/_cancel
// ```
//
// The task ID can be found by using the get tasks API.
//
// Cancellation should happen quickly but might take a few seconds.
// The get task status API will continue to list the delete by query task until
// this task checks that it has been cancelled and terminates itself.
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

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
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

// Delete documents.
//
// Deletes documents that match the specified query.
//
// If the Elasticsearch security features are enabled, you must have the
// following index privileges for the target data stream, index, or alias:
//
// * `read`
// * `delete` or `write`
//
// You can specify the query criteria in the request URI or the request body
// using the same syntax as the search API.
// When you submit a delete by query request, Elasticsearch gets a snapshot of
// the data stream or index when it begins processing the request and deletes
// matching documents using internal versioning.
// If a document changes between the time that the snapshot is taken and the
// delete operation is processed, it results in a version conflict and the
// delete operation fails.
//
// NOTE: Documents with a version equal to 0 cannot be deleted using delete by
// query because internal versioning does not support 0 as a valid version
// number.
//
// While processing a delete by query request, Elasticsearch performs multiple
// search requests sequentially to find all of the matching documents to delete.
// A bulk delete request is performed for each batch of matching documents.
// If a search or bulk request is rejected, the requests are retried up to 10
// times, with exponential back off.
// If the maximum retry limit is reached, processing halts and all failed
// requests are returned in the response.
// Any delete requests that completed successfully still stick, they are not
// rolled back.
//
// You can opt to count version conflicts instead of halting and returning by
// setting `conflicts` to `proceed`.
// Note that if you opt to count version conflicts the operation could attempt
// to delete more documents from the source than `max_docs` until it has
// successfully deleted `max_docs documents`, or it has gone through every
// document in the source query.
//
// **Throttling delete requests**
//
// To control the rate at which delete by query issues batches of delete
// operations, you can set `requests_per_second` to any positive decimal number.
// This pads each batch with a wait time to throttle the rate.
// Set `requests_per_second` to `-1` to disable throttling.
//
// Throttling uses a wait time between batches so that the internal scroll
// requests can be given a timeout that takes the request padding into account.
// The padding time is the difference between the batch size divided by the
// `requests_per_second` and the time spent writing.
// By default the batch size is `1000`, so if `requests_per_second` is set to
// `500`:
//
// ```
// target_time = 1000 / 500 per second = 2 seconds
// wait_time = target_time - write_time = 2 seconds - .5 seconds = 1.5 seconds
// ```
//
// Since the batch is issued as a single `_bulk` request, large batch sizes
// cause Elasticsearch to create many requests and wait before starting the next
// set.
// This is "bursty" instead of "smooth".
//
// **Slicing**
//
// Delete by query supports sliced scroll to parallelize the delete process.
// This can improve efficiency and provide a convenient way to break the request
// down into smaller parts.
//
// Setting `slices` to `auto` lets Elasticsearch choose the number of slices to
// use.
// This setting will use one slice per shard, up to a certain limit.
// If there are multiple source data streams or indices, it will choose the
// number of slices based on the index or backing index with the smallest number
// of shards.
// Adding slices to the delete by query operation creates sub-requests which
// means it has some quirks:
//
// * You can see these requests in the tasks APIs. These sub-requests are
// "child" tasks of the task for the request with slices.
// * Fetching the status of the task for the request with slices only contains
// the status of completed slices.
// * These sub-requests are individually addressable for things like
// cancellation and rethrottling.
// * Rethrottling the request with `slices` will rethrottle the unfinished
// sub-request proportionally.
// * Canceling the request with `slices` will cancel each sub-request.
// * Due to the nature of `slices` each sub-request won't get a perfectly even
// portion of the documents. All documents will be addressed, but some slices
// may be larger than others. Expect larger slices to have a more even
// distribution.
// * Parameters like `requests_per_second` and `max_docs` on a request with
// `slices` are distributed proportionally to each sub-request. Combine that
// with the earlier point about distribution being uneven and you should
// conclude that using `max_docs` with `slices` might not result in exactly
// `max_docs` documents being deleted.
// * Each sub-request gets a slightly different snapshot of the source data
// stream or index though these are all taken at approximately the same time.
//
// If you're slicing manually or otherwise tuning automatic slicing, keep in
// mind that:
//
// * Query performance is most efficient when the number of slices is equal to
// the number of shards in the index or backing index. If that number is large
// (for example, 500), choose a lower number as too many `slices` hurts
// performance. Setting `slices` higher than the number of shards generally does
// not improve efficiency and adds overhead.
// * Delete performance scales linearly across available resources with the
// number of slices.
//
// Whether query or delete performance dominates the runtime depends on the
// documents being reindexed and cluster resources.
//
// **Cancel a delete by query operation**
//
// Any delete by query can be canceled using the task cancel API. For example:
//
// ```
// POST _tasks/r1A2WoRbTwKZ516z6NEs5A:36619/_cancel
// ```
//
// The task ID can be found by using the get tasks API.
//
// Cancellation should happen quickly but might take a few seconds.
// The get task status API will continue to list the delete by query task until
// this task checks that it has been cancelled and terminates itself.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/docs-delete-by-query.html
func New(tp elastictransport.Interface) *DeleteByQuery {
	r := &DeleteByQuery{
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for DeleteByQuery: %w", err)
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
		path.WriteString("_delete_by_query")

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
func (r DeleteByQuery) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "delete_by_query")
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
		instrument.BeforeRequest(req, "delete_by_query")
		if reader := instrument.RecordRequestBody(ctx, "delete_by_query", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "delete_by_query")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the DeleteByQuery query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a deletebyquery.Response
func (r DeleteByQuery) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "delete_by_query")
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

// Header set a key, value pair in the DeleteByQuery headers map.
func (r *DeleteByQuery) Header(key, value string) *DeleteByQuery {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of data streams, indices, and aliases to search.
// It supports wildcards (`*`).
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
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: analyzer
func (r *DeleteByQuery) Analyzer(analyzer string) *DeleteByQuery {
	r.values.Set("analyzer", analyzer)

	return r
}

// AnalyzeWildcard If `true`, wildcard and prefix queries are analyzed.
// This parameter can be used only when the `q` query string parameter is
// specified.
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
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: default_operator
func (r *DeleteByQuery) DefaultOperator(defaultoperator operator.Operator) *DeleteByQuery {
	r.values.Set("default_operator", defaultoperator.String())

	return r
}

// Df The field to use as default where no field prefix is given in the query
// string.
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: df
func (r *DeleteByQuery) Df(df string) *DeleteByQuery {
	r.values.Set("df", df)

	return r
}

// ExpandWildcards The type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// It supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *DeleteByQuery) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *DeleteByQuery {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// From Skips the specified number of documents.
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
// This parameter can be used only when the `q` query string parameter is
// specified.
// API name: lenient
func (r *DeleteByQuery) Lenient(lenient bool) *DeleteByQuery {
	r.values.Set("lenient", strconv.FormatBool(lenient))

	return r
}

// Preference The node or shard the operation should be performed on.
// It is random by default.
// API name: preference
func (r *DeleteByQuery) Preference(preference string) *DeleteByQuery {
	r.values.Set("preference", preference)

	return r
}

// Refresh If `true`, Elasticsearch refreshes all shards involved in the delete by query
// after the request completes.
// This is different than the delete API's `refresh` parameter, which causes
// just the shard that received the delete request to be refreshed.
// Unlike the delete API, it does not support `wait_for`.
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

// Routing A custom value used to route operations to a specific shard.
// API name: routing
func (r *DeleteByQuery) Routing(routing string) *DeleteByQuery {
	r.values.Set("routing", routing)

	return r
}

// Q A query in the Lucene query string syntax.
// API name: q
func (r *DeleteByQuery) Q(q string) *DeleteByQuery {
	r.values.Set("q", q)

	return r
}

// Scroll The period to retain the search context for scrolling.
// API name: scroll
func (r *DeleteByQuery) Scroll(duration string) *DeleteByQuery {
	r.values.Set("scroll", duration)

	return r
}

// ScrollSize The size of the scroll request that powers the operation.
// API name: scroll_size
func (r *DeleteByQuery) ScrollSize(scrollsize string) *DeleteByQuery {
	r.values.Set("scroll_size", scrollsize)

	return r
}

// SearchTimeout The explicit timeout for each search request.
// It defaults to no timeout.
// API name: search_timeout
func (r *DeleteByQuery) SearchTimeout(duration string) *DeleteByQuery {
	r.values.Set("search_timeout", duration)

	return r
}

// SearchType The type of the search operation.
// Available options include `query_then_fetch` and `dfs_query_then_fetch`.
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

// Sort A comma-separated list of `<field>:<direction>` pairs.
// API name: sort
func (r *DeleteByQuery) Sort(sorts ...string) *DeleteByQuery {
	tmp := []string{}
	for _, item := range sorts {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("sort", strings.Join(tmp, ","))

	return r
}

// Stats The specific `tag` of the request for logging and statistical purposes.
// API name: stats
func (r *DeleteByQuery) Stats(stats ...string) *DeleteByQuery {
	tmp := []string{}
	for _, item := range stats {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("stats", strings.Join(tmp, ","))

	return r
}

// TerminateAfter The maximum number of documents to collect for each shard.
// If a query reaches this limit, Elasticsearch terminates the query early.
// Elasticsearch collects documents before sorting.
//
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

// Timeout The period each deletion request waits for active shards.
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
// Set to `all` or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// The `timeout` value controls how long each write request waits for
// unavailable shards to become available.
// API name: wait_for_active_shards
func (r *DeleteByQuery) WaitForActiveShards(waitforactiveshards string) *DeleteByQuery {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// WaitForCompletion If `true`, the request blocks until the operation is complete.
// If `false`, Elasticsearch performs some preflight checks, launches the
// request, and returns a task you can use to cancel or get the status of the
// task. Elasticsearch creates a record of this task as a document at
// `.tasks/task/${taskId}`. When you are done with a task, you should delete the
// task document so Elasticsearch can reclaim the space.
// API name: wait_for_completion
func (r *DeleteByQuery) WaitForCompletion(waitforcompletion bool) *DeleteByQuery {
	r.values.Set("wait_for_completion", strconv.FormatBool(waitforcompletion))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *DeleteByQuery) ErrorTrace(errortrace bool) *DeleteByQuery {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *DeleteByQuery) FilterPath(filterpaths ...string) *DeleteByQuery {
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
func (r *DeleteByQuery) Human(human bool) *DeleteByQuery {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *DeleteByQuery) Pretty(pretty bool) *DeleteByQuery {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// MaxDocs The maximum number of documents to delete.
// API name: max_docs
func (r *DeleteByQuery) MaxDocs(maxdocs int64) *DeleteByQuery {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MaxDocs = &maxdocs

	return r
}

// Query The documents to delete specified with Query DSL.
// API name: query
func (r *DeleteByQuery) Query(query *types.Query) *DeleteByQuery {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}

// Slice Slice the request manually using the provided slice ID and total number of
// slices.
// API name: slice
func (r *DeleteByQuery) Slice(slice *types.SlicedScroll) *DeleteByQuery {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Slice = slice

	return r
}
