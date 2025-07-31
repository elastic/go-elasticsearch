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

// Update field mappings.
// Add new fields to an existing data stream or index.
// You can also use this API to change the search settings of existing fields
// and add new properties to existing object fields.
// For data streams, these changes are applied to all backing indices by
// default.
//
// **Add multi-fields to an existing field**
//
// Multi-fields let you index the same field in different ways.
// You can use this API to update the fields mapping parameter and enable
// multi-fields for an existing field.
// WARNING: If an index (or data stream) contains documents when you add a
// multi-field, those documents will not have values for the new multi-field.
// You can populate the new multi-field with the update by query API.
//
// **Change supported mapping parameters for an existing field**
//
// The documentation for each mapping parameter indicates whether you can update
// it for an existing field using this API.
// For example, you can use the update mapping API to update the `ignore_above`
// parameter.
//
// **Change the mapping of an existing field**
//
// Except for supported mapping parameters, you can't change the mapping or
// field type of an existing field.
// Changing an existing field could invalidate data that's already indexed.
//
// If you need to change the mapping of a field in a data stream's backing
// indices, refer to documentation about modifying data streams.
// If you need to change the mapping of a field in other indices, create a new
// index with the correct mapping and reindex your data into that index.
//
// **Rename a field**
//
// Renaming a field would invalidate data already indexed under the old field
// name.
// Instead, add an alias field to create an alternate field name.
package putmapping

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	indexMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutMapping struct {
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

// NewPutMapping type alias for index.
type NewPutMapping func(index string) *PutMapping

// NewPutMappingFunc returns a new instance of PutMapping with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutMappingFunc(tp elastictransport.Interface) NewPutMapping {
	return func(index string) *PutMapping {
		n := New(tp)

		n._index(index)

		return n
	}
}

// Update field mappings.
// Add new fields to an existing data stream or index.
// You can also use this API to change the search settings of existing fields
// and add new properties to existing object fields.
// For data streams, these changes are applied to all backing indices by
// default.
//
// **Add multi-fields to an existing field**
//
// Multi-fields let you index the same field in different ways.
// You can use this API to update the fields mapping parameter and enable
// multi-fields for an existing field.
// WARNING: If an index (or data stream) contains documents when you add a
// multi-field, those documents will not have values for the new multi-field.
// You can populate the new multi-field with the update by query API.
//
// **Change supported mapping parameters for an existing field**
//
// The documentation for each mapping parameter indicates whether you can update
// it for an existing field using this API.
// For example, you can use the update mapping API to update the `ignore_above`
// parameter.
//
// **Change the mapping of an existing field**
//
// Except for supported mapping parameters, you can't change the mapping or
// field type of an existing field.
// Changing an existing field could invalidate data that's already indexed.
//
// If you need to change the mapping of a field in a data stream's backing
// indices, refer to documentation about modifying data streams.
// If you need to change the mapping of a field in other indices, create a new
// index with the correct mapping and reindex your data into that index.
//
// **Rename a field**
//
// Renaming a field would invalidate data already indexed under the old field
// name.
// Instead, add an alias field to create an alternate field name.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-mapping.html
func New(tp elastictransport.Interface) *PutMapping {
	r := &PutMapping{
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
func (r *PutMapping) Raw(raw io.Reader) *PutMapping {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutMapping) Request(req *Request) *PutMapping {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutMapping) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutMapping: %w", err)
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
		path.WriteString("/")
		path.WriteString("_mapping")

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
func (r PutMapping) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.put_mapping")
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
		instrument.BeforeRequest(req, "indices.put_mapping")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_mapping", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_mapping")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutMapping query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putmapping.Response
func (r PutMapping) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_mapping")
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

// Header set a key, value pair in the PutMapping headers map.
func (r *PutMapping) Header(key, value string) *PutMapping {
	r.headers.Set(key, value)

	return r
}

// Index A comma-separated list of index names the mapping should be added to
// (supports wildcards); use `_all` or omit to add the mapping on all indices.
// API Name: index
func (r *PutMapping) _index(index string) *PutMapping {
	r.paramSet |= indexMask
	r.index = index

	return r
}

// AllowNoIndices If `false`, the request returns an error if any wildcard expression, index
// alias, or `_all` value targets only missing or closed indices.
// This behavior applies even if the request targets other open indices.
// API name: allow_no_indices
func (r *PutMapping) AllowNoIndices(allownoindices bool) *PutMapping {
	r.values.Set("allow_no_indices", strconv.FormatBool(allownoindices))

	return r
}

// ExpandWildcards Type of index that wildcard patterns can match.
// If the request can target data streams, this argument determines whether
// wildcard expressions match hidden data streams.
// Supports comma-separated values, such as `open,hidden`.
// API name: expand_wildcards
func (r *PutMapping) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *PutMapping {
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
func (r *PutMapping) IgnoreUnavailable(ignoreunavailable bool) *PutMapping {
	r.values.Set("ignore_unavailable", strconv.FormatBool(ignoreunavailable))

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *PutMapping) MasterTimeout(duration string) *PutMapping {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *PutMapping) Timeout(duration string) *PutMapping {
	r.values.Set("timeout", duration)

	return r
}

// WriteIndexOnly If `true`, the mappings are applied only to the current write index for the
// target.
// API name: write_index_only
func (r *PutMapping) WriteIndexOnly(writeindexonly bool) *PutMapping {
	r.values.Set("write_index_only", strconv.FormatBool(writeindexonly))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutMapping) ErrorTrace(errortrace bool) *PutMapping {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutMapping) FilterPath(filterpaths ...string) *PutMapping {
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
func (r *PutMapping) Human(human bool) *PutMapping {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutMapping) Pretty(pretty bool) *PutMapping {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// DateDetection Controls whether dynamic date detection is enabled.
// API name: date_detection
func (r *PutMapping) DateDetection(datedetection bool) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DateDetection = &datedetection

	return r
}

// Dynamic Controls whether new fields are added dynamically.
// API name: dynamic
func (r *PutMapping) Dynamic(dynamic dynamicmapping.DynamicMapping) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Dynamic = &dynamic

	return r
}

// DynamicDateFormats If date detection is enabled then new string fields are checked
// against 'dynamic_date_formats' and if the value matches then
// a new date field is added instead of string.
// API name: dynamic_date_formats
func (r *PutMapping) DynamicDateFormats(dynamicdateformats ...string) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DynamicDateFormats = dynamicdateformats

	return r
}

// DynamicTemplates Specify dynamic templates for the mapping.
// API name: dynamic_templates
func (r *PutMapping) DynamicTemplates(dynamictemplates ...map[string]types.DynamicTemplate) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.DynamicTemplates = dynamictemplates

	return r
}

// FieldNames_ Control whether field names are enabled for the index.
// API name: _field_names
func (r *PutMapping) FieldNames_(fieldnames_ *types.FieldNamesField) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.FieldNames_ = fieldnames_

	return r
}

// Meta_ A mapping type can have custom meta data associated with it. These are
// not used at all by Elasticsearch, but can be used to store
// application-specific metadata.
// API name: _meta
func (r *PutMapping) Meta_(metadata types.Metadata) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Meta_ = metadata

	return r
}

// NumericDetection Automatically map strings into numeric data types for all fields.
// API name: numeric_detection
func (r *PutMapping) NumericDetection(numericdetection bool) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.NumericDetection = &numericdetection

	return r
}

// Properties Mapping for a field. For new fields, this mapping can include:
//
// - Field name
// - Field data type
// - Mapping parameters
// API name: properties
func (r *PutMapping) Properties(properties map[string]types.Property) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Properties = properties

	return r
}

// Routing_ Enable making a routing value required on indexed documents.
// API name: _routing
func (r *PutMapping) Routing_(routing_ *types.RoutingField) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Routing_ = routing_

	return r
}

// Runtime Mapping of runtime fields for the index.
// API name: runtime
func (r *PutMapping) Runtime(runtimefields types.RuntimeFields) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Runtime = runtimefields

	return r
}

// Source_ Control whether the _source field is enabled on the index.
// API name: _source
func (r *PutMapping) Source_(source_ *types.SourceField) *PutMapping {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Source_ = source_

	return r
}
