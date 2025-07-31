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

// Roll over to a new index.
// TIP: It is recommended to use the index lifecycle rollover action to automate
// rollovers.
//
// The rollover API creates a new index for a data stream or index alias.
// The API behavior depends on the rollover target.
//
// **Roll over a data stream**
//
// If you roll over a data stream, the API creates a new write index for the
// stream.
// The stream's previous write index becomes a regular backing index.
// A rollover also increments the data stream's generation.
//
// **Roll over an index alias with a write index**
//
// TIP: Prior to Elasticsearch 7.9, you'd typically use an index alias with a
// write index to manage time series data.
// Data streams replace this functionality, require less maintenance, and
// automatically integrate with data tiers.
//
// If an index alias points to multiple indices, one of the indices must be a
// write index.
// The rollover API creates a new write index for the alias with
// `is_write_index` set to `true`.
// The API also `sets is_write_index` to `false` for the previous write index.
//
// **Roll over an index alias with one index**
//
// If you roll over an index alias that points to only one index, the API
// creates a new index for the alias and removes the original index from the
// alias.
//
// NOTE: A rollover creates a new index and is subject to the
// `wait_for_active_shards` setting.
//
// **Increment index names for an alias**
//
// When you roll over an index alias, you can specify a name for the new index.
// If you don't specify a name and the current index ends with `-` and a number,
// such as `my-index-000001` or `my-index-3`, the new index name increments that
// number.
// For example, if you roll over an alias with a current index of
// `my-index-000001`, the rollover creates a new index named `my-index-000002`.
// This number is always six characters and zero-padded, regardless of the
// previous index's name.
//
// If you use an index alias for time series data, you can use date math in the
// index name to track the rollover date.
// For example, you can create an alias that points to an index named
// `<my-index-{now/d}-000001>`.
// If you create the index on May 6, 2099, the index's name is
// `my-index-2099.05.06-000001`.
// If you roll over the alias on May 7, 2099, the new index's name is
// `my-index-2099.05.07-000002`.
package rollover

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
	aliasMask = iota + 1

	newindexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type Rollover struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	alias    string
	newindex string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewRollover type alias for index.
type NewRollover func(alias string) *Rollover

// NewRolloverFunc returns a new instance of Rollover with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewRolloverFunc(tp elastictransport.Interface) NewRollover {
	return func(alias string) *Rollover {
		n := New(tp)

		n._alias(alias)

		return n
	}
}

// Roll over to a new index.
// TIP: It is recommended to use the index lifecycle rollover action to automate
// rollovers.
//
// The rollover API creates a new index for a data stream or index alias.
// The API behavior depends on the rollover target.
//
// **Roll over a data stream**
//
// If you roll over a data stream, the API creates a new write index for the
// stream.
// The stream's previous write index becomes a regular backing index.
// A rollover also increments the data stream's generation.
//
// **Roll over an index alias with a write index**
//
// TIP: Prior to Elasticsearch 7.9, you'd typically use an index alias with a
// write index to manage time series data.
// Data streams replace this functionality, require less maintenance, and
// automatically integrate with data tiers.
//
// If an index alias points to multiple indices, one of the indices must be a
// write index.
// The rollover API creates a new write index for the alias with
// `is_write_index` set to `true`.
// The API also `sets is_write_index` to `false` for the previous write index.
//
// **Roll over an index alias with one index**
//
// If you roll over an index alias that points to only one index, the API
// creates a new index for the alias and removes the original index from the
// alias.
//
// NOTE: A rollover creates a new index and is subject to the
// `wait_for_active_shards` setting.
//
// **Increment index names for an alias**
//
// When you roll over an index alias, you can specify a name for the new index.
// If you don't specify a name and the current index ends with `-` and a number,
// such as `my-index-000001` or `my-index-3`, the new index name increments that
// number.
// For example, if you roll over an alias with a current index of
// `my-index-000001`, the rollover creates a new index named `my-index-000002`.
// This number is always six characters and zero-padded, regardless of the
// previous index's name.
//
// If you use an index alias for time series data, you can use date math in the
// index name to track the rollover date.
// For example, you can create an alias that points to an index named
// `<my-index-{now/d}-000001>`.
// If you create the index on May 6, 2099, the index's name is
// `my-index-2099.05.06-000001`.
// If you roll over the alias on May 7, 2099, the new index's name is
// `my-index-2099.05.07-000002`.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-rollover-index.html
func New(tp elastictransport.Interface) *Rollover {
	r := &Rollover{
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
func (r *Rollover) Raw(raw io.Reader) *Rollover {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Rollover) Request(req *Request) *Rollover {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Rollover) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Rollover: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == aliasMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "alias", r.alias)
		}
		path.WriteString(r.alias)
		path.WriteString("/")
		path.WriteString("_rollover")

		method = http.MethodPost
	case r.paramSet == aliasMask|newindexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "alias", r.alias)
		}
		path.WriteString(r.alias)
		path.WriteString("/")
		path.WriteString("_rollover")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "newindex", r.newindex)
		}
		path.WriteString(r.newindex)

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
func (r Rollover) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.rollover")
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
		instrument.BeforeRequest(req, "indices.rollover")
		if reader := instrument.RecordRequestBody(ctx, "indices.rollover", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.rollover")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Rollover query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a rollover.Response
func (r Rollover) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.rollover")
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

// Header set a key, value pair in the Rollover headers map.
func (r *Rollover) Header(key, value string) *Rollover {
	r.headers.Set(key, value)

	return r
}

// Alias Name of the data stream or index alias to roll over.
// API Name: alias
func (r *Rollover) _alias(alias string) *Rollover {
	r.paramSet |= aliasMask
	r.alias = alias

	return r
}

// NewIndex Name of the index to create.
// Supports date math.
// Data streams do not support this parameter.
// API Name: newindex
func (r *Rollover) NewIndex(newindex string) *Rollover {
	r.paramSet |= newindexMask
	r.newindex = newindex

	return r
}

// DryRun If `true`, checks whether the current index satisfies the specified
// conditions but does not perform a rollover.
// API name: dry_run
func (r *Rollover) DryRun(dryrun bool) *Rollover {
	r.values.Set("dry_run", strconv.FormatBool(dryrun))

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *Rollover) MasterTimeout(duration string) *Rollover {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *Rollover) Timeout(duration string) *Rollover {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to all or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *Rollover) WaitForActiveShards(waitforactiveshards string) *Rollover {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// Lazy If set to true, the rollover action will only mark a data stream to signal
// that it needs to be rolled over at the next write.
// Only allowed on data streams.
// API name: lazy
func (r *Rollover) Lazy(lazy bool) *Rollover {
	r.values.Set("lazy", strconv.FormatBool(lazy))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Rollover) ErrorTrace(errortrace bool) *Rollover {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Rollover) FilterPath(filterpaths ...string) *Rollover {
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
func (r *Rollover) Human(human bool) *Rollover {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Rollover) Pretty(pretty bool) *Rollover {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Aliases Aliases for the target index.
// Data streams do not support this parameter.
// API name: aliases
func (r *Rollover) Aliases(aliases map[string]types.Alias) *Rollover {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Aliases = aliases

	return r
}

// Conditions Conditions for the rollover.
// If specified, Elasticsearch only performs the rollover if the current index
// satisfies these conditions.
// If this parameter is not specified, Elasticsearch performs the rollover
// unconditionally.
// If conditions are specified, at least one of them must be a `max_*`
// condition.
// The index will rollover if any `max_*` condition is satisfied and all `min_*`
// conditions are satisfied.
// API name: conditions
func (r *Rollover) Conditions(conditions *types.RolloverConditions) *Rollover {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Conditions = conditions

	return r
}

// Mappings Mapping for fields in the index.
// If specified, this mapping can include field names, field data types, and
// mapping paramaters.
// API name: mappings
func (r *Rollover) Mappings(mappings *types.TypeMapping) *Rollover {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Mappings = mappings

	return r
}

// Settings Configuration options for the index.
// Data streams do not support this parameter.
// API name: settings
func (r *Rollover) Settings(settings map[string]json.RawMessage) *Rollover {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}
