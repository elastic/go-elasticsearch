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

// Shrink an index.
// Shrink an index into a new index with fewer primary shards.
//
// Before you can shrink an index:
//
// * The index must be read-only.
// * A copy of every shard in the index must reside on the same node.
// * The index must have a green health status.
//
// To make shard allocation easier, we recommend you also remove the index's
// replica shards.
// You can later re-add replica shards as part of the shrink operation.
//
// The requested number of primary shards in the target index must be a factor
// of the number of shards in the source index.
// For example an index with 8 primary shards can be shrunk into 4, 2 or 1
// primary shards or an index with 15 primary shards can be shrunk into 5, 3 or
// 1.
// If the number of shards in the index is a prime number it can only be shrunk
// into a single primary shard
//
//	Before shrinking, a (primary or replica) copy of every shard in the index
//
// must be present on the same node.
//
// The current write index on a data stream cannot be shrunk. In order to shrink
// the current write index, the data stream must first be rolled over so that a
// new write index is created and then the previous write index can be shrunk.
//
// A shrink operation:
//
// * Creates a new target index with the same definition as the source index,
// but with a smaller number of primary shards.
// * Hard-links segments from the source index into the target index. If the
// file system does not support hard-linking, then all segments are copied into
// the new index, which is a much more time consuming process. Also if using
// multiple data paths, shards on different data paths require a full copy of
// segment files if they are not on the same disk since hardlinks do not work
// across disks.
// * Recovers the target index as though it were a closed index which had just
// been re-opened. Recovers shards to the
// `.routing.allocation.initial_recovery._id` index setting.
//
// IMPORTANT: Indices can only be shrunk if they satisfy the following
// requirements:
//
// * The target index must not exist.
// * The source index must have more primary shards than the target index.
// * The number of primary shards in the target index must be a factor of the
// number of primary shards in the source index. The source index must have more
// primary shards than the target index.
// * The index must not contain more than 2,147,483,519 documents in total
// across all shards that will be shrunk into a single shard on the target index
// as this is the maximum number of docs that can fit into a single shard.
// * The node handling the shrink process must have sufficient free disk space
// to accommodate a second copy of the existing index.
package shrink

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

type Shrink struct {
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

// NewShrink type alias for index.
type NewShrink func(index, target string) *Shrink

// NewShrinkFunc returns a new instance of Shrink with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewShrinkFunc(tp elastictransport.Interface) NewShrink {
	return func(index, target string) *Shrink {
		n := New(tp)

		n._index(index)

		n._target(target)

		return n
	}
}

// Shrink an index.
// Shrink an index into a new index with fewer primary shards.
//
// Before you can shrink an index:
//
// * The index must be read-only.
// * A copy of every shard in the index must reside on the same node.
// * The index must have a green health status.
//
// To make shard allocation easier, we recommend you also remove the index's
// replica shards.
// You can later re-add replica shards as part of the shrink operation.
//
// The requested number of primary shards in the target index must be a factor
// of the number of shards in the source index.
// For example an index with 8 primary shards can be shrunk into 4, 2 or 1
// primary shards or an index with 15 primary shards can be shrunk into 5, 3 or
// 1.
// If the number of shards in the index is a prime number it can only be shrunk
// into a single primary shard
//
//	Before shrinking, a (primary or replica) copy of every shard in the index
//
// must be present on the same node.
//
// The current write index on a data stream cannot be shrunk. In order to shrink
// the current write index, the data stream must first be rolled over so that a
// new write index is created and then the previous write index can be shrunk.
//
// A shrink operation:
//
// * Creates a new target index with the same definition as the source index,
// but with a smaller number of primary shards.
// * Hard-links segments from the source index into the target index. If the
// file system does not support hard-linking, then all segments are copied into
// the new index, which is a much more time consuming process. Also if using
// multiple data paths, shards on different data paths require a full copy of
// segment files if they are not on the same disk since hardlinks do not work
// across disks.
// * Recovers the target index as though it were a closed index which had just
// been re-opened. Recovers shards to the
// `.routing.allocation.initial_recovery._id` index setting.
//
// IMPORTANT: Indices can only be shrunk if they satisfy the following
// requirements:
//
// * The target index must not exist.
// * The source index must have more primary shards than the target index.
// * The number of primary shards in the target index must be a factor of the
// number of primary shards in the source index. The source index must have more
// primary shards than the target index.
// * The index must not contain more than 2,147,483,519 documents in total
// across all shards that will be shrunk into a single shard on the target index
// as this is the maximum number of docs that can fit into a single shard.
// * The node handling the shrink process must have sufficient free disk space
// to accommodate a second copy of the existing index.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-shrink-index.html
func New(tp elastictransport.Interface) *Shrink {
	r := &Shrink{
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
func (r *Shrink) Raw(raw io.Reader) *Shrink {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Shrink) Request(req *Request) *Shrink {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Shrink) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Shrink: %w", err)
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
		path.WriteString("_shrink")
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
func (r Shrink) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.shrink")
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
		instrument.BeforeRequest(req, "indices.shrink")
		if reader := instrument.RecordRequestBody(ctx, "indices.shrink", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.shrink")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Shrink query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a shrink.Response
func (r Shrink) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.shrink")
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

// Header set a key, value pair in the Shrink headers map.
func (r *Shrink) Header(key, value string) *Shrink {
	r.headers.Set(key, value)

	return r
}

// Index Name of the source index to shrink.
// API Name: index
func (r *Shrink) _index(index string) *Shrink {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Target Name of the target index to create.
// API Name: target
func (r *Shrink) _target(target string) *Shrink {
	r.paramSet |= targetMask
	r.target = target

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *Shrink) MasterTimeout(duration string) *Shrink {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *Shrink) Timeout(duration string) *Shrink {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to `all` or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *Shrink) WaitForActiveShards(waitforactiveshards string) *Shrink {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Shrink) ErrorTrace(errortrace bool) *Shrink {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Shrink) FilterPath(filterpaths ...string) *Shrink {
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
func (r *Shrink) Human(human bool) *Shrink {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Shrink) Pretty(pretty bool) *Shrink {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Aliases The key is the alias name.
// Index alias names support date math.
// API name: aliases
func (r *Shrink) Aliases(aliases map[string]types.Alias) *Shrink {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Aliases = aliases

	return r
}

// Settings Configuration options for the target index.
// API name: settings
func (r *Shrink) Settings(settings map[string]json.RawMessage) *Shrink {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}
