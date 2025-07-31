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

// Open a closed index.
// For data streams, the API opens any closed backing indices.
//
// A closed index is blocked for read/write operations and does not allow all
// operations that opened indices allow.
// It is not possible to index documents or to search for documents in a closed
// index.
// This allows closed indices to not have to maintain internal data structures
// for indexing or searching documents, resulting in a smaller overhead on the
// cluster.
//
// When opening or closing an index, the master is responsible for restarting
// the index shards to reflect the new state of the index.
// The shards will then go through the normal recovery process.
// The data of opened or closed indices is automatically replicated by the
// cluster to ensure that enough shard copies are safely kept around at all
// times.
//
// You can open and close multiple indices.
// An error is thrown if the request explicitly refers to a missing index.
// This behavior can be turned off by using the `ignore_unavailable=true`
// parameter.
//
// By default, you must explicitly name the indices you are opening or closing.
// To open or close indices with `_all`, `*`, or other wildcard expressions,
// change the `action.destructive_requires_name` setting to `false`.
// This setting can also be changed with the cluster update settings API.
//
// Closed indices consume a significant amount of disk-space which can cause
// problems in managed environments.
// Closing indices can be turned off with the cluster settings API by setting
// `cluster.indices.close.enable` to `false`.
//
// Because opening or closing an index allocates its shards, the
// `wait_for_active_shards` setting on index creation applies to the `_open` and
// `_close` index actions as well.
package open

import (
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

type Open struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewOpen type alias for index.
type NewOpen func(index string) *Open

// NewOpenFunc returns a new instance of Open with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewOpenFunc(tp elastictransport.Interface) NewOpen {
	return func(index string) *Open {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Open a closed index.
// For data streams, the API opens any closed backing indices.
//
// A closed index is blocked for read/write operations and does not allow all
// operations that opened indices allow.
// It is not possible to index documents or to search for documents in a closed
// index.
// This allows closed indices to not have to maintain internal data structures
// for indexing or searching documents, resulting in a smaller overhead on the
// cluster.
//
// When opening or closing an index, the master is responsible for restarting
// the index shards to reflect the new state of the index.
// The shards will then go through the normal recovery process.
// The data of opened or closed indices is automatically replicated by the
// cluster to ensure that enough shard copies are safely kept around at all
// times.
//
// You can open and close multiple indices.
// An error is thrown if the request explicitly refers to a missing index.
// This behavior can be turned off by using the `ignore_unavailable=true`
// parameter.
//
// By default, you must explicitly name the indices you are opening or closing.
// To open or close indices with `_all`, `*`, or other wildcard expressions,
// change the `action.destructive_requires_name` setting to `false`.
// This setting can also be changed with the cluster update settings API.
//
// Closed indices consume a significant amount of disk-space which can cause
// problems in managed environments.
// Closing indices can be turned off with the cluster settings API by setting
// `cluster.indices.close.enable` to `false`.
//
// Because opening or closing an index allocates its shards, the
// `wait_for_active_shards` setting on index creation applies to the `_open` and
// `_close` index actions as well.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-open-close.html
func New(tp elastictransport.Interface) *Open {
	r := &Open{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Open) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_open")

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r Open) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.open")
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
		instrument.BeforeRequest(req, "indices.open")
		if reader := instrument.RecordRequestBody(ctx, "indices.open", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.open")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Open query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a open.Response
func (r Open) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.open")
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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r Open) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.open")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	res, err := r.Perform(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(io.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	if res.StatusCode != 404 {
		err := fmt.Errorf("an error happened during the Open query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the Open headers map.
func (r *Open) Header(key, value string) *Open {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams, indices, and aliases used to limit the
// request.
// Supports wildcards (`*`).
// By default, you must explicitly name the indices you using to limit the
// request.
// To limit a request using `_all`, `*`, or other wildcard expressions, change
// the `action.destructive_requires_name` setting to false.
// You can update this setting in the `elasticsearch.yml` file or using the
// cluster update settings API.
// API Name: index
func (r *Open) _index(index string) *Open {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// API name: allow_no_indices
func (r *Open) AllowNoIndices(allownoindices bool) *Open {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *Open) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *Open {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreUnavailable If `false`, the request returns an error if it targets a missing or closed
// index.
// API name: ignore_unavailable
func (r *Open) IgnoreUnavailable(ignoreunavailable bool) *Open {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *Open) MasterTimeout(duration string) *Open {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *Open) Timeout(duration string) *Open {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to `all` or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *Open) WaitForActiveShards(waitforactiveshards string) *Open {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Open) ErrorTrace(errortrace bool) *Open {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Open) FilterPath(filterpaths ...string) *Open {
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
func (r *Open) Human(human bool) *Open {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Open) Pretty(pretty bool) *Open {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
