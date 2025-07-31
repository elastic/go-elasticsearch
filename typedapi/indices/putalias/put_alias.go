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

// Create or update an alias.
// Adds a data stream or index to an alias.
package putalias

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

	nameMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutAlias struct {
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
	name  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutAlias type alias for index.
type NewPutAlias func(index, name string) *PutAlias

// NewPutAliasFunc returns a new instance of PutAlias with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutAliasFunc(tp elastictransport.Interface) NewPutAlias {
	return func(index, name string) *PutAlias {
		n := New(tp)

		n._index(index)

		n._name(name)

		return n
	}
}

// Create or update an alias.
// Adds a data stream or index to an alias.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-put-alias
func New(tp elastictransport.Interface) *PutAlias {
	r := &PutAlias{
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
func (r *PutAlias) Raw(raw io.Reader) *PutAlias {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutAlias) Request(req *Request) *PutAlias {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutAlias) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutAlias: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == indexMask|nameMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_alias")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

		method = http.MethodPut
	case r.paramSet == indexMask|nameMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_aliases")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

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
func (r PutAlias) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.put_alias")
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
		instrument.BeforeRequest(req, "indices.put_alias")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_alias", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_alias")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutAlias query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putalias.Response
func (r PutAlias) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_alias")
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

// Header set a key, value pair in the PutAlias headers map.
func (r *PutAlias) Header(key, value string) *PutAlias {
	r.headers.Set(key, value)

	return r
}

// Index Comma-separated list of data streams or indices to add.
// Supports wildcards (`*`).
// Wildcard patterns that match both data streams and indices return an error.
// API Name: index
func (r *PutAlias) _index(index string) *PutAlias {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// Name Alias to update.
// If the alias doesn’t exist, the request creates it.
// Index alias names support date math.
// API Name: name
func (r *PutAlias) _name(name string) *PutAlias {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *PutAlias) MasterTimeout(duration string) *PutAlias {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *PutAlias) Timeout(duration string) *PutAlias {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutAlias) ErrorTrace(errortrace bool) *PutAlias {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutAlias) FilterPath(filterpaths ...string) *PutAlias {
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
func (r *PutAlias) Human(human bool) *PutAlias {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutAlias) Pretty(pretty bool) *PutAlias {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Filter Query used to limit documents the alias can access.
// API name: filter
func (r *PutAlias) Filter(filter *types.Query) *PutAlias {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Filter = filter

	return r
}

// IndexRouting Value used to route indexing operations to a specific shard.
// If specified, this overwrites the `routing` value for indexing operations.
// Data stream aliases don’t support this parameter.
// API name: index_routing
func (r *PutAlias) IndexRouting(routing string) *PutAlias {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IndexRouting = &routing

	return r
}

// IsWriteIndex If `true`, sets the write index or data stream for the alias.
// If an alias points to multiple indices or data streams and `is_write_index`
// isn’t set, the alias rejects write requests.
// If an index alias points to one index and `is_write_index` isn’t set, the
// index automatically acts as the write index.
// Data stream aliases don’t automatically set a write data stream, even if the
// alias points to one data stream.
// API name: is_write_index
func (r *PutAlias) IsWriteIndex(iswriteindex bool) *PutAlias {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IsWriteIndex = &iswriteindex

	return r
}

// Routing Value used to route indexing and search operations to a specific shard.
// Data stream aliases don’t support this parameter.
// API name: routing
func (r *PutAlias) Routing(routing string) *PutAlias {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Routing = &routing

	return r
}

// SearchRouting Value used to route search operations to a specific shard.
// If specified, this overwrites the `routing` value for search operations.
// Data stream aliases don’t support this parameter.
// API name: search_routing
func (r *PutAlias) SearchRouting(routing string) *PutAlias {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.SearchRouting = &routing

	return r
}
