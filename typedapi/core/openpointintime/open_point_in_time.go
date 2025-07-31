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

// Open a point in time.
//
// A search request by default runs against the most recent visible data of the
// target indices,
// which is called point in time. Elasticsearch pit (point in time) is a
// lightweight view into the
// state of the data as it existed when initiated. In some cases, it’s preferred
// to perform multiple
// search requests using the same point in time. For example, if refreshes
// happen between
// `search_after` requests, then the results of those requests might not be
// consistent as changes happening
// between searches are only visible to the more recent point in time.
//
// A point in time must be opened explicitly before being used in search
// requests.
//
// A subsequent search request with the `pit` parameter must not specify
// `index`, `routing`, or `preference` values as these parameters are copied
// from the point in time.
//
// Just like regular searches, you can use `from` and `size` to page through
// point in time search results, up to the first 10,000 hits.
// If you want to retrieve more hits, use PIT with `search_after`.
//
// IMPORTANT: The open point in time request and each subsequent search request
// can return different identifiers; always use the most recently received ID
// for the next search request.
//
// When a PIT that contains shard failures is used in a search request, the
// missing are always reported in the search response as a
// `NoShardAvailableActionException` exception.
// To get rid of these exceptions, a new PIT needs to be created so that shards
// missing from the previous PIT can be handled, assuming they become available
// in the meantime.
//
// **Keeping point in time alive**
//
// The `keep_alive` parameter, which is passed to a open point in time request
// and search request, extends the time to live of the corresponding point in
// time.
// The value does not need to be long enough to process all data — it just needs
// to be long enough for the next request.
//
// Normally, the background merge process optimizes the index by merging
// together smaller segments to create new, bigger segments.
// Once the smaller segments are no longer needed they are deleted.
// However, open point-in-times prevent the old segments from being deleted
// since they are still in use.
//
// TIP: Keeping older segments alive means that more disk space and file handles
// are needed.
// Ensure that you have configured your nodes to have ample free file handles.
//
// Additionally, if a segment contains deleted or updated documents then the
// point in time must keep track of whether each document in the segment was
// live at the time of the initial search request.
// Ensure that your nodes have sufficient heap space if you have many open
// point-in-times on an index that is subject to ongoing deletes or updates.
// Note that a point-in-time doesn't prevent its associated indices from being
// deleted.
// You can check how many point-in-times (that is, search contexts) are open
// with the nodes stats API.
package openpointintime

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
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type OpenPointInTime struct {
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

// NewOpenPointInTime type alias for index.
type NewOpenPointInTime func(index string) *OpenPointInTime

// NewOpenPointInTimeFunc returns a new instance of OpenPointInTime with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewOpenPointInTimeFunc(tp elastictransport.Interface) NewOpenPointInTime {
	return func(index string) *OpenPointInTime {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Open a point in time.
//
// A search request by default runs against the most recent visible data of the
// target indices,
// which is called point in time. Elasticsearch pit (point in time) is a
// lightweight view into the
// state of the data as it existed when initiated. In some cases, it’s preferred
// to perform multiple
// search requests using the same point in time. For example, if refreshes
// happen between
// `search_after` requests, then the results of those requests might not be
// consistent as changes happening
// between searches are only visible to the more recent point in time.
//
// A point in time must be opened explicitly before being used in search
// requests.
//
// A subsequent search request with the `pit` parameter must not specify
// `index`, `routing`, or `preference` values as these parameters are copied
// from the point in time.
//
// Just like regular searches, you can use `from` and `size` to page through
// point in time search results, up to the first 10,000 hits.
// If you want to retrieve more hits, use PIT with `search_after`.
//
// IMPORTANT: The open point in time request and each subsequent search request
// can return different identifiers; always use the most recently received ID
// for the next search request.
//
// When a PIT that contains shard failures is used in a search request, the
// missing are always reported in the search response as a
// `NoShardAvailableActionException` exception.
// To get rid of these exceptions, a new PIT needs to be created so that shards
// missing from the previous PIT can be handled, assuming they become available
// in the meantime.
//
// **Keeping point in time alive**
//
// The `keep_alive` parameter, which is passed to a open point in time request
// and search request, extends the time to live of the corresponding point in
// time.
// The value does not need to be long enough to process all data — it just needs
// to be long enough for the next request.
//
// Normally, the background merge process optimizes the index by merging
// together smaller segments to create new, bigger segments.
// Once the smaller segments are no longer needed they are deleted.
// However, open point-in-times prevent the old segments from being deleted
// since they are still in use.
//
// TIP: Keeping older segments alive means that more disk space and file handles
// are needed.
// Ensure that you have configured your nodes to have ample free file handles.
//
// Additionally, if a segment contains deleted or updated documents then the
// point in time must keep track of whether each document in the segment was
// live at the time of the initial search request.
// Ensure that your nodes have sufficient heap space if you have many open
// point-in-times on an index that is subject to ongoing deletes or updates.
// Note that a point-in-time doesn't prevent its associated indices from being
// deleted.
// You can check how many point-in-times (that is, search contexts) are open
// with the nodes stats API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/point-in-time-api.html
func New(tp elastictransport.Interface) *OpenPointInTime {
	r := &OpenPointInTime{
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
func (r *OpenPointInTime) Raw(raw io.Reader) *OpenPointInTime {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *OpenPointInTime) Request(req *Request) *OpenPointInTime {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *OpenPointInTime) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for OpenPointInTime: %w", err)
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
		path.WriteString("_pit")

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
func (r OpenPointInTime) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "open_point_in_time")
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
		instrument.BeforeRequest(req, "open_point_in_time")
		if reader := instrument.RecordRequestBody(ctx, "open_point_in_time", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "open_point_in_time")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the OpenPointInTime query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a openpointintime.Response
func (r OpenPointInTime) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "open_point_in_time")
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

// Header set a key, value pair in the OpenPointInTime headers map.
func (r *OpenPointInTime) Header(key, value string) *OpenPointInTime {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names to open point in time; use `_all` or
// empty string to perform the operation on all indices
// API Name: index
func (r *OpenPointInTime) _index(index string) *OpenPointInTime {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// KeepAlive Extend the length of time that the point in time persists.
// API name: keep_alive
func (r *OpenPointInTime) KeepAlive(duration string) *OpenPointInTime {
	r.values.Set("keep_alive", duration)

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *OpenPointInTime) IgnoreUnavailable(ignoreunavailable bool) *OpenPointInTime {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// Preference The node or shard the operation should be performed on.
// By default, it is random.
// API name: preference
func (r *OpenPointInTime) Preference(preference string) *OpenPointInTime {
	r.values.Set("preference", preference)

	return r
}

// Routing A custom value that is used to route operations to a specific shard.
// API name: routing
func (r *OpenPointInTime) Routing(routing string) *OpenPointInTime {
	r.values.Set("routing", routing)

	return r
}

// ExpandWildcards The type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// It supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *OpenPointInTime) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *OpenPointInTime {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// AllowPartialSearchResults Indicates whether the point in time tolerates unavailable shards or shard
// failures when initially creating the PIT.
// If `false`, creating a point in time request when a shard is missing or
// unavailable will throw an exception.
// If `true`, the point in time will contain all the shards that are available
// at the time of the request.
// API name: allow_partial_search_results
func (r *OpenPointInTime) AllowPartialSearchResults(allowpartialsearchresults bool) *OpenPointInTime {
	r.values.Set("allow_partial_search_results", strconv.FormatBool(allowpartialsearchresults))

	return r
}

// MaxConcurrentShardRequests Maximum number of concurrent shard requests that each sub-search request
// executes per node.
// API name: max_concurrent_shard_requests
func (r *OpenPointInTime) MaxConcurrentShardRequests(maxconcurrentshardrequests int) *OpenPointInTime {
	r.values.Set("max_concurrent_shard_requests", strconv.Itoa(maxconcurrentshardrequests))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *OpenPointInTime) ErrorTrace(errortrace bool) *OpenPointInTime {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *OpenPointInTime) FilterPath(filterpaths ...string) *OpenPointInTime {
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
func (r *OpenPointInTime) Human(human bool) *OpenPointInTime {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *OpenPointInTime) Pretty(pretty bool) *OpenPointInTime {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// IndexFilter Filter indices if the provided query rewrites to `match_none` on every shard.
// API name: index_filter
func (r *OpenPointInTime) IndexFilter(indexfilter *types.Query) *OpenPointInTime {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexFilter = indexfilter

	return r
}
