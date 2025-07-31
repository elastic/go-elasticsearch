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

// Get mapping definitions.
// Retrieves mapping definitions for one or more fields.
// For data streams, the API retrieves field mappings for the stream’s backing
// indices.
//
// This API is useful if you don't need a complete mapping or if an index
// mapping contains a large number of fields.
package getfieldmapping

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
	fieldsMask = iota + 1

	indexMask
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetFieldMapping struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	paramSet int

	fields string
	index  string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewGetFieldMapping type alias for index.
type NewGetFieldMapping func(fields string) *GetFieldMapping

// NewGetFieldMappingFunc returns a new instance of GetFieldMapping with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetFieldMappingFunc(tp elastictransport.Interface) NewGetFieldMapping {
	return func(fields string) *GetFieldMapping {
		n := New(tp)

		n._fields(fields)

		return n
	}
}

// Get mapping definitions.
// Retrieves mapping definitions for one or more fields.
// For data streams, the API retrieves field mappings for the stream’s backing
// indices.
//
// This API is useful if you don't need a complete mapping or if an index
// mapping contains a large number of fields.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-get-field-mapping.html
func New(tp elastictransport.Interface) *GetFieldMapping {
	r := &GetFieldMapping{
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
func (r *GetFieldMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == fieldsMask:
		path.WriteString("/")
		path.WriteString("_mapping")
		path.WriteString("/")
		path.WriteString("field")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "fields", r.fields)
		}
		path.WriteString(r.fields)

		method = http.MethodGet
	case r.paramSet == indexMask|fieldsMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_mapping")
		path.WriteString("/")
		path.WriteString("field")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "fields", r.fields)
		}
		path.WriteString(r.fields)

		method = http.MethodGet
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
func (r GetFieldMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.get_field_mapping")
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
		instrument.BeforeRequest(req, "indices.get_field_mapping")
		if reader := instrument.RecordRequestBody(ctx, "indices.get_field_mapping", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.get_field_mapping")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetFieldMapping query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getfieldmapping.Response
func (r GetFieldMapping) Do(providedCtx context.Context) (Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.get_field_mapping")
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
		err = json.NewDecoder(res.Body).Decode(&response)
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
func (r GetFieldMapping) IsSuccess(providedCtx context.Context) (bool, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.get_field_mapping")
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
		err := fmt.Errorf("an error happened during the GetFieldMapping query execution, status code: %d", res.StatusCode)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return false, err
	}

	return false, nil
}

// Header set a key, value pair in the GetFieldMapping headers map.
func (r *GetFieldMapping) Header(key, value string) *GetFieldMapping {
	r.headers.Set(key, value)

	return r
}

// Fields Comma-separated list or wildcard expression of fields used to limit returned
// information.
// Supports wildcards (`*`).
// API Name: fields
func (r *GetFieldMapping) _fields(fields string) *GetFieldMapping {
	r.paramSet |= fieldsMask
	r.fields = fields

	return r
}

// Index Comma-separated list of data streams, indices, and aliases used to limit the
// request.
// Supports wildcards (`*`).
// To target all data streams and indices, omit this parameter or use `*` or
// `_all`.
// API Name: index
func (r *GetFieldMapping) Index(index string) *GetFieldMapping {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// API name: allow_no_indices
func (r *GetFieldMapping) AllowNoIndices(allownoindices bool) *GetFieldMapping {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *GetFieldMapping) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetFieldMapping {
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
func (r *GetFieldMapping) IgnoreUnavailable(ignoreunavailable bool) *GetFieldMapping {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// IncludeDefaults If `true`, return all default settings in the response.
// API name: include_defaults
func (r *GetFieldMapping) IncludeDefaults(includedefaults bool) *GetFieldMapping {
	r.values.Set("include_defaults", strconv.FormatBool(includedefaults))

	return r
}

// Local If `true`, the request retrieves information from the local node only.
// API name: local
func (r *GetFieldMapping) Local(local bool) *GetFieldMapping {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetFieldMapping) ErrorTrace(errortrace bool) *GetFieldMapping {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetFieldMapping) FilterPath(filterpaths ...string) *GetFieldMapping {
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
func (r *GetFieldMapping) Human(human bool) *GetFieldMapping {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetFieldMapping) Pretty(pretty bool) *GetFieldMapping {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}
