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

// Clone an index.
// Clone an existing index into a new index.
// Each original primary shard is cloned into a new primary shard in the new
// index.
//
// IMPORTANT: Elasticsearch does not apply index templates to the resulting
// index.
// The API also does not copy index metadata from the original index.
// Index metadata includes aliases, index lifecycle management phase
// definitions, and cross-cluster replication (CCR) follower information.
// For example, if you clone a CCR follower index, the resulting clone will not
// be a follower index.
//
// The clone API copies most index settings from the source index to the
// resulting index, with the exception of `index.number_of_replicas` and
// `index.auto_expand_replicas`.
// To set the number of replicas in the resulting index, configure these
// settings in the clone request.
//
// Cloning works as follows:
//
// * First, it creates a new target index with the same definition as the source
// index.
// * Then it hard-links segments from the source index into the target index. If
// the file system does not support hard-linking, all segments are copied into
// the new index, which is a much more time consuming process.
// * Finally, it recovers the target index as though it were a closed index
// which had just been re-opened.
//
// IMPORTANT: Indices can only be cloned if they meet the following
// requirements:
//
// * The index must be marked as read-only and have a cluster health status of
// green.
// * The target index must not exist.
// * The source index must have the same number of primary shards as the target
// index.
// * The node handling the clone process must have sufficient free disk space to
// accommodate a second copy of the existing index.
//
// The current write index on a data stream cannot be cloned.
// In order to clone the current write index, the data stream must first be
// rolled over so that a new write index is created and then the previous write
// index can be cloned.
//
// NOTE: Mappings cannot be specified in the `_clone` request. The mappings of
// the source index will be used for the target index.
//
// **Monitor the cloning process**
//
// The cloning process can be monitored with the cat recovery API or the cluster
// health API can be used to wait until all primary shards have been allocated
// by setting the `wait_for_status` parameter to `yellow`.
//
// The `_clone` API returns as soon as the target index has been added to the
// cluster state, before any shards have been allocated.
// At this point, all shards are in the state unassigned.
// If, for any reason, the target index can't be allocated, its primary shard
// will remain unassigned until it can be allocated on that node.
//
// Once the primary shard is allocated, it moves to state initializing, and the
// clone process begins.
// When the clone operation completes, the shard will become active.
// At that point, Elasticsearch will try to allocate any replicas and may decide
// to relocate the primary shard to another node.
//
// **Wait for active shards**
//
// Because the clone operation creates a new index to clone the shards to, the
// wait for active shards setting on index creation applies to the clone index
// action as well.
package clone

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

type Clone struct {
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

// NewClone type alias for index.
type NewClone func(index, target string) *Clone

// NewCloneFunc returns a new instance of Clone with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCloneFunc(tp elastictransport.Interface) NewClone {
	return func(index, target string) *Clone {
		n := New(tp)

		n._index(index)

		n._target(target)

		return n
	}
}

// Clone an index.
// Clone an existing index into a new index.
// Each original primary shard is cloned into a new primary shard in the new
// index.
//
// IMPORTANT: Elasticsearch does not apply index templates to the resulting
// index.
// The API also does not copy index metadata from the original index.
// Index metadata includes aliases, index lifecycle management phase
// definitions, and cross-cluster replication (CCR) follower information.
// For example, if you clone a CCR follower index, the resulting clone will not
// be a follower index.
//
// The clone API copies most index settings from the source index to the
// resulting index, with the exception of `index.number_of_replicas` and
// `index.auto_expand_replicas`.
// To set the number of replicas in the resulting index, configure these
// settings in the clone request.
//
// Cloning works as follows:
//
// * First, it creates a new target index with the same definition as the source
// index.
// * Then it hard-links segments from the source index into the target index. If
// the file system does not support hard-linking, all segments are copied into
// the new index, which is a much more time consuming process.
// * Finally, it recovers the target index as though it were a closed index
// which had just been re-opened.
//
// IMPORTANT: Indices can only be cloned if they meet the following
// requirements:
//
// * The index must be marked as read-only and have a cluster health status of
// green.
// * The target index must not exist.
// * The source index must have the same number of primary shards as the target
// index.
// * The node handling the clone process must have sufficient free disk space to
// accommodate a second copy of the existing index.
//
// The current write index on a data stream cannot be cloned.
// In order to clone the current write index, the data stream must first be
// rolled over so that a new write index is created and then the previous write
// index can be cloned.
//
// NOTE: Mappings cannot be specified in the `_clone` request. The mappings of
// the source index will be used for the target index.
//
// **Monitor the cloning process**
//
// The cloning process can be monitored with the cat recovery API or the cluster
// health API can be used to wait until all primary shards have been allocated
// by setting the `wait_for_status` parameter to `yellow`.
//
// The `_clone` API returns as soon as the target index has been added to the
// cluster state, before any shards have been allocated.
// At this point, all shards are in the state unassigned.
// If, for any reason, the target index can't be allocated, its primary shard
// will remain unassigned until it can be allocated on that node.
//
// Once the primary shard is allocated, it moves to state initializing, and the
// clone process begins.
// When the clone operation completes, the shard will become active.
// At that point, Elasticsearch will try to allocate any replicas and may decide
// to relocate the primary shard to another node.
//
// **Wait for active shards**
//
// Because the clone operation creates a new index to clone the shards to, the
// wait for active shards setting on index creation applies to the clone index
// action as well.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-clone-index.html
func New(tp elastictransport.Interface) *Clone {
	r := &Clone{
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
func (r *Clone) Raw(raw io.Reader) *Clone {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Clone) Request(req *Request) *Clone {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Clone) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Clone: %w", err)
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
		path.WriteString("_clone")
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
func (r Clone) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.clone")
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
		instrument.BeforeRequest(req, "indices.clone")
		if reader := instrument.RecordRequestBody(ctx, "indices.clone", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.clone")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Clone query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a clone.Response
func (r Clone) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.clone")
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

// Header set a key, value pair in the Clone headers map.
func (r *Clone) Header(key, value string) *Clone {
	r.headers.Set(key, value)

	return r
}

// Index Name of the source index to clone.
// API Name: index
func (r *Clone) _index(index string) *Clone {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Target Name of the target index to create.
// API Name: target
func (r *Clone) _target(target string) *Clone {
	r.paramSet |= targetMask
	r.target = target

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *Clone) MasterTimeout(duration string) *Clone {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *Clone) Timeout(duration string) *Clone {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to `all` or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *Clone) WaitForActiveShards(waitforactiveshards string) *Clone {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Clone) ErrorTrace(errortrace bool) *Clone {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Clone) FilterPath(filterpaths ...string) *Clone {
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
func (r *Clone) Human(human bool) *Clone {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Clone) Pretty(pretty bool) *Clone {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Aliases Aliases for the resulting index.
// API name: aliases
func (r *Clone) Aliases(aliases map[string]types.Alias) *Clone {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Aliases = aliases

	return r
}

// Settings Configuration options for the target index.
// API name: settings
func (r *Clone) Settings(settings map[string]json.RawMessage) *Clone {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}
