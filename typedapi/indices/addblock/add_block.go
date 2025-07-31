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

// Add an index block.
//
// Add an index block to an index.
// Index blocks limit the operations allowed on an index by blocking specific
// operation types.
package addblock

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

	blockMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type AddBlock struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	index string
	block string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewAddBlock type alias for index.
type NewAddBlock func(index, block string) *AddBlock

// NewAddBlockFunc returns a new instance of AddBlock with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewAddBlockFunc(tp elastictransport.Interface) NewAddBlock {
	return func(index, block string) *AddBlock {
		n := New(tp)

		n._index(index)

		n._block(block)

		return n
	}
}

// Add an index block.
//
// Add an index block to an index.
// Index blocks limit the operations allowed on an index by blocking specific
// operation types.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/index-modules-blocks.html#add-index-block
func New(tp elastictransport.Interface) *AddBlock {
	r := &AddBlock{
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
func (r *AddBlock) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|blockMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_block")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "block", r.block)
		}
		path.WriteString(r.block)

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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r AddBlock) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.add_block")
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
		instrument.BeforeRequest(req, "indices.add_block")
		if reader := instrument.RecordRequestBody(ctx, "indices.add_block", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.add_block")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the AddBlock query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a addblock.Response
func (r AddBlock) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.add_block")
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
func (r AddBlock) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.add_block")
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
		err := fmt.Errorf("an error happened during the AddBlock query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the AddBlock headers map.
func (r *AddBlock) Header(key, value string) *AddBlock {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list or wildcard expression of index names used to limit
// the request.
// By default, you must explicitly name the indices you are adding blocks to.
// To allow the adding of blocks to indices with `_all`, `*`, or other wildcard
// expressions, change the `action.destructive_requires_name` setting to
// `false`.
// You can update this setting in the `elasticsearch.yml` file or by using the
// cluster update settings API.
// API Name: index
func (r *AddBlock) _index(index string) *AddBlock {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Block The block type to add to the index.
// API Name: block
func (r *AddBlock) _block(block string) *AddBlock {
	r.paramSet |= blockMask
	r.block = block

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// For example, a request targeting `foo*,bar*` returns an error if an index
// starts with `foo` but no index starts with `bar`.
// API name: allow_no_indices
func (r *AddBlock) AllowNoIndices(allownoindices bool) *AddBlock {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards The type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// It supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *AddBlock) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *AddBlock {
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
func (r *AddBlock) IgnoreUnavailable(ignoreunavailable bool) *AddBlock {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MasterTimeout The period to wait for the master node.
// If the master node is not available before the timeout expires, the request
// fails and returns an error.
// It can also be set to `-1` to indicate that the request should never timeout.
// API name: master_timeout
func (r *AddBlock) MasterTimeout(duration string) *AddBlock {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout The period to wait for a response from all relevant nodes in the cluster
// after updating the cluster metadata.
// If no response is received before the timeout expires, the cluster metadata
// update still applies but the response will indicate that it was not
// completely acknowledged.
// It can also be set to `-1` to indicate that the request should never timeout.
// API name: timeout
func (r *AddBlock) Timeout(duration string) *AddBlock {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *AddBlock) ErrorTrace(errortrace bool) *AddBlock {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *AddBlock) FilterPath(filterpaths ...string) *AddBlock {
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
func (r *AddBlock) Human(human bool) *AddBlock {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *AddBlock) Pretty(pretty bool) *AddBlock {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
