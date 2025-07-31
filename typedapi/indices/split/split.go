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

// Split an index.
// Split an index into a new index with more primary shards.
// * Before you can split an index:
//
// * The index must be read-only.
// * The cluster health status must be green.
//
// You can do make an index read-only with the following request using the add
// index block API:
//
// ```
// PUT /my_source_index/_block/write
// ```
//
// The current write index on a data stream cannot be split.
// In order to split the current write index, the data stream must first be
// rolled over so that a new write index is created and then the previous write
// index can be split.
//
// The number of times the index can be split (and the number of shards that
// each original shard can be split into) is determined by the
// `index.number_of_routing_shards` setting.
// The number of routing shards specifies the hashing space that is used
// internally to distribute documents across shards with consistent hashing.
// For instance, a 5 shard index with `number_of_routing_shards` set to 30 (5 x
// 2 x 3) could be split by a factor of 2 or 3.
//
// A split operation:
//
// * Creates a new target index with the same definition as the source index,
// but with a larger number of primary shards.
// * Hard-links segments from the source index into the target index. If the
// file system doesn't support hard-linking, all segments are copied into the
// new index, which is a much more time consuming process.
// * Hashes all documents again, after low level files are created, to delete
// documents that belong to a different shard.
// * Recovers the target index as though it were a closed index which had just
// been re-opened.
//
// IMPORTANT: Indices can only be split if they satisfy the following
// requirements:
//
// * The target index must not exist.
// * The source index must have fewer primary shards than the target index.
// * The number of primary shards in the target index must be a multiple of the
// number of primary shards in the source index.
// * The node handling the split process must have sufficient free disk space to
// accommodate a second copy of the existing index.
package split

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

	targetMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Split struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	index  string
	target string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewSplit type alias for index.
type NewSplit func(index, target string) *Split

// NewSplitFunc returns a new instance of Split with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSplitFunc(tp elastictransport.Interface) NewSplit {
	return func(index, target string) *Split {
		n := New(tp)

		n._index(index)

		n._target(target)

		return n
	}
}

// Split an index.
// Split an index into a new index with more primary shards.
// * Before you can split an index:
//
// * The index must be read-only.
// * The cluster health status must be green.
//
// You can do make an index read-only with the following request using the add
// index block API:
//
// ```
// PUT /my_source_index/_block/write
// ```
//
// The current write index on a data stream cannot be split.
// In order to split the current write index, the data stream must first be
// rolled over so that a new write index is created and then the previous write
// index can be split.
//
// The number of times the index can be split (and the number of shards that
// each original shard can be split into) is determined by the
// `index.number_of_routing_shards` setting.
// The number of routing shards specifies the hashing space that is used
// internally to distribute documents across shards with consistent hashing.
// For instance, a 5 shard index with `number_of_routing_shards` set to 30 (5 x
// 2 x 3) could be split by a factor of 2 or 3.
//
// A split operation:
//
// * Creates a new target index with the same definition as the source index,
// but with a larger number of primary shards.
// * Hard-links segments from the source index into the target index. If the
// file system doesn't support hard-linking, all segments are copied into the
// new index, which is a much more time consuming process.
// * Hashes all documents again, after low level files are created, to delete
// documents that belong to a different shard.
// * Recovers the target index as though it were a closed index which had just
// been re-opened.
//
// IMPORTANT: Indices can only be split if they satisfy the following
// requirements:
//
// * The target index must not exist.
// * The source index must have fewer primary shards than the target index.
// * The number of primary shards in the target index must be a multiple of the
// number of primary shards in the source index.
// * The node handling the split process must have sufficient free disk space to
// accommodate a second copy of the existing index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-split-index.html
func New(tp elastictransport.Interface) *Split {
	r := &Split{
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
func (r *Split) Raw(raw io.Reader) *Split {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Split) Request(req *Request) *Split {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Split) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Split: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|targetMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_split")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "target", r.target)
		}
		path.WriteString(r.target)

		method = http.MethodPut
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
func (r Split) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.split")
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
		instrument.BeforeRequest(req, "indices.split")
		if reader := instrument.RecordRequestBody(ctx, "indices.split", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.split")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Split query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a split.Response
func (r Split) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.split")
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

// Header set a key, value pair in the Split headers map.
func (r *Split) Header(key, value string) *Split {
	r.headers.Set(key, value)

	return r
}

// Index Name of the source index to split.
// API Name: index
func (r *Split) _index(index string) *Split {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Target Name of the target index to create.
// API Name: target
func (r *Split) _target(target string) *Split {
	r.paramSet |= targetMask
	r.target = target

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *Split) MasterTimeout(duration string) *Split {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *Split) Timeout(duration string) *Split {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to `all` or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *Split) WaitForActiveShards(waitforactiveshards string) *Split {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Split) ErrorTrace(errortrace bool) *Split {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Split) FilterPath(filterpaths ...string) *Split {
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
func (r *Split) Human(human bool) *Split {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Split) Pretty(pretty bool) *Split {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Aliases Aliases for the resulting index.
// API name: aliases
func (r *Split) Aliases(aliases map[string]types.Alias) *Split {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Aliases = aliases

	return r
}

// Settings Configuration options for the target index.
// API name: settings
func (r *Split) Settings(settings map[string]json.RawMessage) *Split {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}
