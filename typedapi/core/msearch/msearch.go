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

// Run multiple searches.
//
// The format of the request is similar to the bulk API format and makes use of
// the newline delimited JSON (NDJSON) format.
// The structure is as follows:
//
// ```
// header\n
// body\n
// header\n
// body\n
// ```
//
// This structure is specifically optimized to reduce parsing if a specific
// search ends up redirected to another node.
//
// IMPORTANT: The final line of data must end with a newline character `\n`.
// Each newline character may be preceded by a carriage return `\r`.
// When sending requests to this endpoint the `Content-Type` header should be
// set to `application/x-ndjson`.
package msearch

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/searchtype"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Msearch struct {
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

// NewMsearch type alias for index.
type NewMsearch func() *Msearch

// NewMsearchFunc returns a new instance of Msearch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewMsearchFunc(tp elastictransport.Interface) NewMsearch {
	return func() *Msearch {
		n := New(tp)

		return n
	}
}

// Run multiple searches.
//
// The format of the request is similar to the bulk API format and makes use of
// the newline delimited JSON (NDJSON) format.
// The structure is as follows:
//
// ```
// header\n
// body\n
// header\n
// body\n
// ```
//
// This structure is specifically optimized to reduce parsing if a specific
// search ends up redirected to another node.
//
// IMPORTANT: The final line of data must end with a newline character `\n`.
// Each newline character may be preceded by a carriage return `\r`.
// When sending requests to this endpoint the `Content-Type` header should be
// set to `application/x-ndjson`.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-msearch
func New(tp elastictransport.Interface) *Msearch {
	r := &Msearch{
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
func (r *Msearch) Raw(raw io.Reader) *Msearch {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Msearch) Request(req *Request) *Msearch {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Msearch) HttpRequest(ctx context.Context) (*http.Request, error) {
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

		for _, elem := range *r.req {
			data, err := json.Marshal(elem)
			if err != nil {
				return nil, err
			}
			r.buf.Write(data)
			r.buf.Write([]byte("\n"))
		}

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for Msearch: %w", err)
		}

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_msearch")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_msearch")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+x-ndjson;compatible-with=9")
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
func (r Msearch) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "msearch")
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
		instrument.BeforeRequest(req, "msearch")
		if reader := instrument.RecordRequestBody(ctx, "msearch", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "msearch")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Msearch query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a msearch.Response
func (r Msearch) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "msearch")
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

// Header set a key, value pair in the Msearch headers map.
func (r *Msearch) Header(key, value string) *Msearch {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and index aliases to search.
// API Name: index
func (r *Msearch) Index(index string) *Msearch {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If false, the request returns an error if any wildcard expression, index
// alias, or _all value targets only missing or closed indices. This behavior
// applies even if the request targets other open indices. For example, a
// request targeting foo*,bar* returns an error if an index starts with foo but
// no index starts with bar.
// API name: allow_no_indices
func (r *Msearch) AllowNoIndices(allownoindices bool) *Msearch {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// CcsMinimizeRoundtrips If true, network roundtrips between the coordinating node and remote clusters
// are minimized for cross-cluster search requests.
// API name: ccs_minimize_roundtrips
func (r *Msearch) CcsMinimizeRoundtrips(ccsminimizeroundtrips bool) *Msearch {
	r.values.Set("ccs_minimize_roundtrips", strconv.FormatBool(ccsminimizeroundtrips))

	return r
}

// ExpandWildcards Type of index that wildcard expressions can match. If the request can target
// data streams, this argument determines whether wildcard expressions match
// hidden data streams.
// API name: expand_wildcards
func (r *Msearch) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Msearch {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreThrottled If true, concrete, expanded or aliased indices are ignored when frozen.
// API name: ignore_throttled
func (r *Msearch) IgnoreThrottled(ignorethrottled bool) *Msearch {
	r.values.Set("ignore_throttled", strconv.FormatBool(ignorethrottled))

	return r
}

// IgnoreUnavailable If true, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *Msearch) IgnoreUnavailable(ignoreunavailable bool) *Msearch {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// IncludeNamedQueriesScore Indicates whether hit.matched_queries should be rendered as a map that
// includes
// the name of the matched query associated with its score (true)
// or as an array containing the name of the matched queries (false)
// This functionality reruns each named query on every hit in a search response.
// Typically, this adds a small overhead to a request.
// However, using computationally expensive named queries on a large number of
// hits may add significant overhead.
// API name: include_named_queries_score
func (r *Msearch) IncludeNamedQueriesScore(includenamedqueriesscore bool) *Msearch {
	r.values.Set("include_named_queries_score", strconv.FormatBool(includenamedqueriesscore))

	return r
}

// MaxConcurrentSearches Maximum number of concurrent searches the multi search API can execute.
// Defaults to `max(1, (# of data nodes * min(search thread pool size, 10)))`.
// API name: max_concurrent_searches
func (r *Msearch) MaxConcurrentSearches(maxconcurrentsearches int) *Msearch {
	r.values.Set("max_concurrent_searches", strconv.Itoa(maxconcurrentsearches))

	return r
}

// MaxConcurrentShardRequests Maximum number of concurrent shard requests that each sub-search request
// executes per node.
// API name: max_concurrent_shard_requests
func (r *Msearch) MaxConcurrentShardRequests(maxconcurrentshardrequests int) *Msearch {
	r.values.Set("max_concurrent_shard_requests", strconv.Itoa(maxconcurrentshardrequests))

	return r
}

// PreFilterShardSize Defines a threshold that enforces a pre-filter roundtrip to prefilter search
// shards based on query rewriting if the number of shards the search request
// expands to exceeds the threshold. This filter roundtrip can limit the number
// of shards significantly if for instance a shard can not match any documents
// based on its rewrite method i.e., if date filters are mandatory to match but
// the shard bounds and the query are disjoint.
// API name: pre_filter_shard_size
func (r *Msearch) PreFilterShardSize(prefiltershardsize string) *Msearch {
	r.values.Set("pre_filter_shard_size", prefiltershardsize)

	return r
}

// RestTotalHitsAsInt If true, hits.total are returned as an integer in the response. Defaults to
// false, which returns an object.
// API name: rest_total_hits_as_int
func (r *Msearch) RestTotalHitsAsInt(resttotalhitsasint bool) *Msearch {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// Routing Custom routing value used to route search operations to a specific shard.
// API name: routing
func (r *Msearch) Routing(routing string) *Msearch {
	r.values.Set("routing", routing)

	return r
}

// SearchType Indicates whether global term and document frequencies should be used when
// scoring returned documents.
// API name: search_type
func (r *Msearch) SearchType(searchtype searchtype.SearchType) *Msearch {
	r.values.Set("search_type", searchtype.String())

	return r
}

// TypedKeys Specifies whether aggregation and suggester names should be prefixed by their
// respective types in the response.
// API name: typed_keys
func (r *Msearch) TypedKeys(typedkeys bool) *Msearch {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Msearch) ErrorTrace(errortrace bool) *Msearch {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Msearch) FilterPath(filterpaths ...string) *Msearch {
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
func (r *Msearch) Human(human bool) *Msearch {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Msearch) Pretty(pretty bool) *Msearch {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
