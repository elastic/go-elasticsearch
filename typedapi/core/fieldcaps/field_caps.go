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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

// Get the field capabilities.
//
// Get information about the capabilities of fields among multiple indices.
//
// For data streams, the API returns field capabilities among the stream’s
// backing indices.
// It returns runtime fields like any other field.
// For example, a runtime field with a type of keyword is returned the same as
// any other field that belongs to the `keyword` family.
package fieldcaps

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type FieldCaps struct {
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

// NewFieldCaps type alias for index.
type NewFieldCaps func() *FieldCaps

// NewFieldCapsFunc returns a new instance of FieldCaps with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewFieldCapsFunc(tp elastictransport.Interface) NewFieldCaps {
	return func() *FieldCaps {
		n := New(tp)

		return n
	}
}

// Get the field capabilities.
//
// Get information about the capabilities of fields among multiple indices.
//
// For data streams, the API returns field capabilities among the stream’s
// backing indices.
// It returns runtime fields like any other field.
// For example, a runtime field with a type of keyword is returned the same as
// any other field that belongs to the `keyword` family.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-field-caps
func New(tp elastictransport.Interface) *FieldCaps {
	r := &FieldCaps{
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
func (r *FieldCaps) Raw(raw io.Reader) *FieldCaps {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *FieldCaps) Request(req *Request) *FieldCaps {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *FieldCaps) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for FieldCaps: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_field_caps")

		method = http.MethodPost
	case r.paramSet == indexMask:
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", r.index)
		}
		path.WriteString(r.index)
		path.WriteString("/")
		path.WriteString("_field_caps")

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
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=9")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r FieldCaps) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "field_caps")
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
		instrument.BeforeRequest(req, "field_caps")
		if reader := instrument.RecordRequestBody(ctx, "field_caps", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "field_caps")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the FieldCaps query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a fieldcaps.Response
func (r FieldCaps) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "field_caps")
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

// Header set a key, value pair in the FieldCaps headers map.
func (r *FieldCaps) Header(key, value string) *FieldCaps {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of data streams, indices, and aliases used to limit
// the request. Supports wildcards (*). To target all data streams and indices,
// omit this parameter or use * or _all.
// API Name: index
func (r *FieldCaps) Index(index string) *FieldCaps {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If false, the request returns an error if any wildcard expression, index
// alias,
// or `_all` value targets only missing or closed indices. This behavior applies
// even if the request targets other open indices. For example, a request
// targeting `foo*,bar*` returns an error if an index starts with foo but no
// index starts with bar.
// API name: allow_no_indices
func (r *FieldCaps) AllowNoIndices(allownoindices bool) *FieldCaps {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards The type of index that wildcard patterns can match. If the request can target
// data streams, this argument determines whether wildcard expressions match
// hidden data streams. Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *FieldCaps) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *FieldCaps {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IgnoreUnavailable If `true`, missing or closed indices are not included in the response.
// API name: ignore_unavailable
func (r *FieldCaps) IgnoreUnavailable(ignoreunavailable bool) *FieldCaps {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// IncludeUnmapped If true, unmapped fields are included in the response.
// API name: include_unmapped
func (r *FieldCaps) IncludeUnmapped(includeunmapped bool) *FieldCaps {
	r.values.Set("include_unmapped", strconv.FormatBool(includeunmapped))

	return r
}

// Filters A comma-separated list of filters to apply to the response.
// API name: filters
func (r *FieldCaps) Filters(filters string) *FieldCaps {
	r.values.Set("filters", filters)

	return r
}

// Types A comma-separated list of field types to include.
// Any fields that do not match one of these types will be excluded from the
// results.
// It defaults to empty, meaning that all field types are returned.
// API name: types
func (r *FieldCaps) Types(types ...string) *FieldCaps {
	tmp := []string{}
	for _, item := range types {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("types", strings.Join(tmp, ","))

	return r
}

// IncludeEmptyFields If false, empty fields are not included in the response.
// API name: include_empty_fields
func (r *FieldCaps) IncludeEmptyFields(includeemptyfields bool) *FieldCaps {
	r.values.Set("include_empty_fields", strconv.FormatBool(includeemptyfields))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *FieldCaps) ErrorTrace(errortrace bool) *FieldCaps {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *FieldCaps) FilterPath(filterpaths ...string) *FieldCaps {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"exists_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *FieldCaps) Human(human bool) *FieldCaps {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *FieldCaps) Pretty(pretty bool) *FieldCaps {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// A list of fields to retrieve capabilities for. Wildcard (`*`) expressions are
// supported.
// API name: fields
func (r *FieldCaps) Fields(fields ...string) *FieldCaps {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Fields = fields

	return r
}

// Filter indices if the provided query rewrites to `match_none` on every shard.
//
// IMPORTANT: The filtering is done on a best-effort basis, it uses index
// statistics and mappings to rewrite queries to `match_none` instead of fully
// running the request.
// For instance a range query over a date field can rewrite to `match_none` if
// all documents within a shard (including deleted documents) are outside of the
// provided range.
// However, not all queries can rewrite to `match_none` so this API may return
// an index even if the provided filter matches no document.
// API name: index_filter
func (r *FieldCaps) IndexFilter(indexfilter types.QueryVariant) *FieldCaps {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexFilter = indexfilter.QueryCaster()

	return r
}

// Define ad-hoc runtime fields in the request similar to the way it is done in
// search requests.
// These fields exist only as part of the query and take precedence over fields
// defined with the same name in the index mappings.
// API name: runtime_mappings
func (r *FieldCaps) RuntimeMappings(runtimefields types.RuntimeFieldsVariant) *FieldCaps {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RuntimeMappings = *runtimefields.RuntimeFieldsCaster()

	return r
}
