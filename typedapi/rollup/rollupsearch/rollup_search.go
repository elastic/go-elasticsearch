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

// Search rolled-up data.
// The rollup search endpoint is needed because, internally, rolled-up documents
// utilize a different document structure than the original data.
// It rewrites standard Query DSL into a format that matches the rollup
// documents then takes the response and rewrites it back to what a client would
// expect given the original query.
//
// The request body supports a subset of features from the regular search API.
// The following functionality is not available:
//
// `size`: Because rollups work on pre-aggregated data, no search hits can be
// returned and so size must be set to zero or omitted entirely.
// `highlighter`, `suggestors`, `post_filter`, `profile`, `explain`: These are
// similarly disallowed.
//
// **Searching both historical rollup and non-rollup data**
//
// The rollup search API has the capability to search across both "live"
// non-rollup data and the aggregated rollup data.
// This is done by simply adding the live indices to the URI. For example:
//
// ```
// GET sensor-1,sensor_rollup/_rollup_search
//
//	{
//	  "size": 0,
//	  "aggregations": {
//	     "max_temperature": {
//	      "max": {
//	        "field": "temperature"
//	      }
//	    }
//	  }
//	}
//
// ```
//
// The rollup search endpoint does two things when the search runs:
//
// * The original request is sent to the non-rollup index unaltered.
// * A rewritten version of the original request is sent to the rollup index.
//
// When the two responses are received, the endpoint rewrites the rollup
// response and merges the two together.
// During the merging process, if there is any overlap in buckets between the
// two responses, the buckets from the non-rollup index are used.
package rollupsearch

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
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type RollupSearch struct {
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

// NewRollupSearch type alias for index.
type NewRollupSearch func(index string) *RollupSearch

// NewRollupSearchFunc returns a new instance of RollupSearch with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRollupSearchFunc(tp elastictransport.Interface) NewRollupSearch {
	return func(index string) *RollupSearch {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Search rolled-up data.
// The rollup search endpoint is needed because, internally, rolled-up documents
// utilize a different document structure than the original data.
// It rewrites standard Query DSL into a format that matches the rollup
// documents then takes the response and rewrites it back to what a client would
// expect given the original query.
//
// The request body supports a subset of features from the regular search API.
// The following functionality is not available:
//
// `size`: Because rollups work on pre-aggregated data, no search hits can be
// returned and so size must be set to zero or omitted entirely.
// `highlighter`, `suggestors`, `post_filter`, `profile`, `explain`: These are
// similarly disallowed.
//
// **Searching both historical rollup and non-rollup data**
//
// The rollup search API has the capability to search across both "live"
// non-rollup data and the aggregated rollup data.
// This is done by simply adding the live indices to the URI. For example:
//
// ```
// GET sensor-1,sensor_rollup/_rollup_search
//
//	{
//	  "size": 0,
//	  "aggregations": {
//	     "max_temperature": {
//	      "max": {
//	        "field": "temperature"
//	      }
//	    }
//	  }
//	}
//
// ```
//
// The rollup search endpoint does two things when the search runs:
//
// * The original request is sent to the non-rollup index unaltered.
// * A rewritten version of the original request is sent to the rollup index.
//
// When the two responses are received, the endpoint rewrites the rollup
// response and merges the two together.
// During the merging process, if there is any overlap in buckets between the
// two responses, the buckets from the non-rollup index are used.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/rollup-search.html
func New(tp elastictransport.Interface) *RollupSearch {
	r := &RollupSearch{
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
func (r *RollupSearch) Raw(raw io.Reader) *RollupSearch {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *RollupSearch) Request(req *Request) *RollupSearch {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *RollupSearch) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for RollupSearch: %w", err)
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
		path.WriteString("_rollup_search")

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
func (r RollupSearch) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "rollup.rollup_search")
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
		instrument.BeforeRequest(req, "rollup.rollup_search")
		if reader := instrument.RecordRequestBody(ctx, "rollup.rollup_search", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "rollup.rollup_search")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the RollupSearch query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a rollupsearch.Response
func (r RollupSearch) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "rollup.rollup_search")
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

// Header set a key, value pair in the RollupSearch headers map.
func (r *RollupSearch) Header(key, value string) *RollupSearch {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of data streams and indices used to limit the request.
// This parameter has the following rules:
//
// * At least one data stream, index, or wildcard expression must be specified.
// This target can include a rollup or non-rollup index. For data streams, the
// stream's backing indices can only serve as non-rollup indices. Omitting the
// parameter or using `_all` are not permitted.
// * Multiple non-rollup indices may be specified.
// * Only one rollup index may be specified. If more than one are supplied, an
// exception occurs.
// * Wildcard expressions (`*`) may be used. If they match more than one rollup
// index, an exception occurs. However, you can use an expression to match
// multiple non-rollup indices or data streams.
// API Name: index
func (r *RollupSearch) _index(index string) *RollupSearch {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// RestTotalHitsAsInt Indicates whether hits.total should be rendered as an integer or an object in
// the rest search response
// API name: rest_total_hits_as_int
func (r *RollupSearch) RestTotalHitsAsInt(resttotalhitsasint bool) *RollupSearch {
	r.values.Set("rest_total_hits_as_int", strconv.FormatBool(resttotalhitsasint))

	return r
}

// TypedKeys Specify whether aggregation and suggester names should be prefixed by their
// respective types in the response
// API name: typed_keys
func (r *RollupSearch) TypedKeys(typedkeys bool) *RollupSearch {
	r.values.Set("typed_keys", strconv.FormatBool(typedkeys))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *RollupSearch) ErrorTrace(errortrace bool) *RollupSearch {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *RollupSearch) FilterPath(filterpaths ...string) *RollupSearch {
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
func (r *RollupSearch) Human(human bool) *RollupSearch {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *RollupSearch) Pretty(pretty bool) *RollupSearch {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Aggregations Specifies aggregations.
// API name: aggregations
func (r *RollupSearch) Aggregations(aggregations map[string]types.Aggregations) *RollupSearch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Aggregations = aggregations

	return r
}

// Query Specifies a DSL query that is subject to some limitations.
// API name: query
func (r *RollupSearch) Query(query *types.Query) *RollupSearch {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Query = query

	return r
}

// Size Must be zero if set, as rollups work on pre-aggregated data.
// API name: size
func (r *RollupSearch) Size(size int) *RollupSearch {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Size = &size

	return r
}
