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

// Create an index.
// You can use the create index API to add a new index to an Elasticsearch
// cluster.
// When creating an index, you can specify the following:
//
// * Settings for the index.
// * Mappings for fields in the index.
// * Index aliases
//
// **Wait for active shards**
//
// By default, index creation will only return a response to the client when the
// primary copies of each shard have been started, or the request times out.
// The index creation response will indicate what happened.
// For example, `acknowledged` indicates whether the index was successfully
// created in the cluster, `while shards_acknowledged` indicates whether the
// requisite number of shard copies were started for each shard in the index
// before timing out.
// Note that it is still possible for either `acknowledged` or
// `shards_acknowledged` to be `false`, but for the index creation to be
// successful.
// These values simply indicate whether the operation completed before the
// timeout.
// If `acknowledged` is false, the request timed out before the cluster state
// was updated with the newly created index, but it probably will be created
// sometime soon.
// If `shards_acknowledged` is false, then the request timed out before the
// requisite number of shards were started (by default just the primaries), even
// if the cluster state was successfully updated to reflect the newly created
// index (that is to say, `acknowledged` is `true`).
//
// You can change the default of only waiting for the primary shards to start
// through the index setting `index.write.wait_for_active_shards`.
// Note that changing this setting will also affect the `wait_for_active_shards`
// value on all subsequent write operations.
package create

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

type Create struct {
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

// NewCreate type alias for index.
type NewCreate func(index string) *Create

// NewCreateFunc returns a new instance of Create with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateFunc(tp elastictransport.Interface) NewCreate {
	return func(index string) *Create {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Create an index.
// You can use the create index API to add a new index to an Elasticsearch
// cluster.
// When creating an index, you can specify the following:
//
// * Settings for the index.
// * Mappings for fields in the index.
// * Index aliases
//
// **Wait for active shards**
//
// By default, index creation will only return a response to the client when the
// primary copies of each shard have been started, or the request times out.
// The index creation response will indicate what happened.
// For example, `acknowledged` indicates whether the index was successfully
// created in the cluster, `while shards_acknowledged` indicates whether the
// requisite number of shard copies were started for each shard in the index
// before timing out.
// Note that it is still possible for either `acknowledged` or
// `shards_acknowledged` to be `false`, but for the index creation to be
// successful.
// These values simply indicate whether the operation completed before the
// timeout.
// If `acknowledged` is false, the request timed out before the cluster state
// was updated with the newly created index, but it probably will be created
// sometime soon.
// If `shards_acknowledged` is false, then the request timed out before the
// requisite number of shards were started (by default just the primaries), even
// if the cluster state was successfully updated to reflect the newly created
// index (that is to say, `acknowledged` is `true`).
//
// You can change the default of only waiting for the primary shards to start
// through the index setting `index.write.wait_for_active_shards`.
// Note that changing this setting will also affect the `wait_for_active_shards`
// value on all subsequent write operations.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-create-index.html
func New(tp elastictransport.Interface) *Create {
	r := &Create{
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
func (r *Create) Raw(raw io.Reader) *Create {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *Create) Request(req *Request) *Create {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *Create) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for Create: %w", err)
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
func (r Create) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.create")
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
		instrument.BeforeRequest(req, "indices.create")
		if reader := instrument.RecordRequestBody(ctx, "indices.create", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.create")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the Create query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a create.Response
func (r Create) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.create")
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

// Header set a key, value pair in the Create headers map.
func (r *Create) Header(key, value string) *Create {
	r.headers.Set(key, value)

	return r
}

// Index Name of the index you wish to create.
// API Name: index
func (r *Create) _index(index string) *Create {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *Create) MasterTimeout(duration string) *Create {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *Create) Timeout(duration string) *Create {
	r.values.Set("timeout", duration)

	return r
}

// WaitForActiveShards The number of shard copies that must be active before proceeding with the
// operation.
// Set to `all` or any positive integer up to the total number of shards in the
// index (`number_of_replicas+1`).
// API name: wait_for_active_shards
func (r *Create) WaitForActiveShards(waitforactiveshards string) *Create {
	r.values.Set("wait_for_active_shards", waitforactiveshards)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *Create) ErrorTrace(errortrace bool) *Create {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *Create) FilterPath(filterpaths ...string) *Create {
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
func (r *Create) Human(human bool) *Create {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *Create) Pretty(pretty bool) *Create {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Aliases Aliases for the index.
// API name: aliases
func (r *Create) Aliases(aliases map[string]types.Alias) *Create {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Aliases = aliases

	return r
}

// Mappings Mapping for fields in the index. If specified, this mapping can include:
// - Field names
// - Field data types
// - Mapping parameters
// API name: mappings
func (r *Create) Mappings(mappings *types.TypeMapping) *Create {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Mappings = mappings

	return r
}

// Settings Configuration options for the index.
// API name: settings
func (r *Create) Settings(settings *types.IndexSettings) *Create {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}
