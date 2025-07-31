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

// Create an index from a source index.
//
// Copy the mappings and settings from the source index to a destination index
// while allowing request settings and mappings to override the source values.
package createfrom

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
	sourceMask = iota + 1

	destMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type CreateFrom struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	source string
	dest   string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewCreateFrom type alias for index.
type NewCreateFrom func(source, dest string) *CreateFrom

// NewCreateFromFunc returns a new instance of CreateFrom with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewCreateFromFunc(tp elastictransport.Interface) NewCreateFrom {
	return func(source, dest string) *CreateFrom {
		n := New(tp)

		n._source(source)

		n._dest(dest)

		return n
	}
}

// Create an index from a source index.
//
// Copy the mappings and settings from the source index to a destination index
// while allowing request settings and mappings to override the source values.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/migrate-data-stream.html
func New(tp elastictransport.Interface) *CreateFrom {
	r := &CreateFrom{
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
func (r *CreateFrom) Raw(raw io.Reader) *CreateFrom {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *CreateFrom) Request(req *Request) *CreateFrom {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *CreateFrom) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for CreateFrom: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == sourceMask|destMask:
		path.WriteString("/")
		path.WriteString("_create_from")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "source", r.source)
		}
		path.WriteString(r.source)
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "dest", r.dest)
		}
		path.WriteString(r.dest)

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
func (r CreateFrom) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.create_from")
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
		instrument.BeforeRequest(req, "indices.create_from")
		if reader := instrument.RecordRequestBody(ctx, "indices.create_from", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.create_from")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the CreateFrom query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a createfrom.Response
func (r CreateFrom) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.create_from")
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

// Header set a key, value pair in the CreateFrom headers map.
func (r *CreateFrom) Header(key, value string) *CreateFrom {
	r.headers.Set(key, value)

	return r
}

// Source The source index or data stream name
// API Name: source
func (r *CreateFrom) _source(source string) *CreateFrom {
	r.paramSet |= sourceMask
	r.source = source

	return r
}

// Dest The destination index or data stream name
// API Name: dest
func (r *CreateFrom) _dest(dest string) *CreateFrom {
	r.paramSet |= destMask
	r.dest = dest

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *CreateFrom) ErrorTrace(errortrace bool) *CreateFrom {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *CreateFrom) FilterPath(filterpaths ...string) *CreateFrom {
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
func (r *CreateFrom) Human(human bool) *CreateFrom {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *CreateFrom) Pretty(pretty bool) *CreateFrom {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// MappingsOverride Mappings overrides to be applied to the destination index (optional)
// API name: mappings_override
func (r *CreateFrom) MappingsOverride(mappingsoverride *types.TypeMapping) *CreateFrom {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.MappingsOverride = mappingsoverride

	return r
}

// RemoveIndexBlocks If index blocks should be removed when creating destination index (optional)
// API name: remove_index_blocks
func (r *CreateFrom) RemoveIndexBlocks(removeindexblocks bool) *CreateFrom {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RemoveIndexBlocks = &removeindexblocks

	return r
}

// SettingsOverride Settings overrides to be applied to the destination index (optional)
// API name: settings_override
func (r *CreateFrom) SettingsOverride(settingsoverride *types.IndexSettings) *CreateFrom {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.SettingsOverride = settingsoverride

	return r
}
