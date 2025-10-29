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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

// Update data stream mappings.
//
// This API can be used to override mappings on specific data streams. These
// overrides will take precedence over what
// is specified in the template that the data stream matches. The mapping change
// is only applied to new write indices
// that are created during rollover after this API is called. No indices are
// changed by this API.
package putdatastreammappings

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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/dynamicmapping"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/subobjects"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutDataStreamMappings struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	name string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutDataStreamMappings type alias for index.
type NewPutDataStreamMappings func(name string) *PutDataStreamMappings

// NewPutDataStreamMappingsFunc returns a new instance of PutDataStreamMappings with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutDataStreamMappingsFunc(tp elastictransport.Interface) NewPutDataStreamMappings {
	return func(name string) *PutDataStreamMappings {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Update data stream mappings.
//
// This API can be used to override mappings on specific data streams. These
// overrides will take precedence over what
// is specified in the template that the data stream matches. The mapping change
// is only applied to new write indices
// that are created during rollover after this API is called. No indices are
// changed by this API.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-indices-put-data-stream-mappings
func New(tp elastictransport.Interface) *PutDataStreamMappings {
	r := &PutDataStreamMappings{
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
func (r *PutDataStreamMappings) Raw(raw io.Reader) *PutDataStreamMappings {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutDataStreamMappings) Request(req *Request) *PutDataStreamMappings {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutDataStreamMappings) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutDataStreamMappings: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_data_stream")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)
		path.WriteString("/")
		path.WriteString("_mappings")

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
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=9")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r PutDataStreamMappings) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.put_data_stream_mappings")
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
		instrument.BeforeRequest(req, "indices.put_data_stream_mappings")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_data_stream_mappings", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_data_stream_mappings")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutDataStreamMappings query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putdatastreammappings.Response
func (r PutDataStreamMappings) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_data_stream_mappings")
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

// Header set a key, value pair in the PutDataStreamMappings headers map.
func (r *PutDataStreamMappings) Header(key, value string) *PutDataStreamMappings {
	r.headers.Set(key, value)

	return r
}

// Name A comma-separated list of data streams or data stream patterns.
// API Name: name
func (r *PutDataStreamMappings) _name(name string) *PutDataStreamMappings {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// DryRun If `true`, the request does not actually change the mappings on any data
// streams. Instead, it
// simulates changing the settings and reports back to the user what would have
// happened had these settings
// actually been applied.
// API name: dry_run
func (r *PutDataStreamMappings) DryRun(dryrun bool) *PutDataStreamMappings {
	r.values.Set("dry_run", strconv.FormatBool(dryrun))

	return r
}

// MasterTimeout The period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an
// error.
// API name: master_timeout
func (r *PutDataStreamMappings) MasterTimeout(duration string) *PutDataStreamMappings {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout The period to wait for a response. If no response is received before the
//
//	timeout expires, the request fails and returns an error.
//
// API name: timeout
func (r *PutDataStreamMappings) Timeout(duration string) *PutDataStreamMappings {
	r.values.Set("timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutDataStreamMappings) ErrorTrace(errortrace bool) *PutDataStreamMappings {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutDataStreamMappings) FilterPath(filterpaths ...string) *PutDataStreamMappings {
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
func (r *PutDataStreamMappings) Human(human bool) *PutDataStreamMappings {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutDataStreamMappings) Pretty(pretty bool) *PutDataStreamMappings {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: all_field
func (r *PutDataStreamMappings) AllField(allfield types.AllFieldVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AllField = allfield.AllFieldCaster()

	return r
}

// API name: _data_stream_timestamp
func (r *PutDataStreamMappings) DataStreamTimestamp_(datastreamtimestamp_ types.DataStreamTimestampVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DataStreamTimestamp_ = datastreamtimestamp_.DataStreamTimestampCaster()

	return r
}

// API name: date_detection
func (r *PutDataStreamMappings) DateDetection(datedetection bool) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DateDetection = &datedetection

	return r
}

// API name: dynamic
func (r *PutDataStreamMappings) Dynamic(dynamic dynamicmapping.DynamicMapping) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Dynamic = &dynamic
	return r
}

// API name: dynamic_date_formats
func (r *PutDataStreamMappings) DynamicDateFormats(dynamicdateformats ...string) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range dynamicdateformats {

		r.req.DynamicDateFormats = append(r.req.DynamicDateFormats, v)

	}
	return r
}

// API name: dynamic_templates
func (r *PutDataStreamMappings) DynamicTemplates(dynamictemplates []map[string]types.DynamicTemplate) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DynamicTemplates = dynamictemplates

	return r
}

// API name: enabled
func (r *PutDataStreamMappings) Enabled(enabled bool) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Enabled = &enabled

	return r
}

// API name: _field_names
func (r *PutDataStreamMappings) FieldNames_(fieldnames_ types.FieldNamesFieldVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.FieldNames_ = fieldnames_.FieldNamesFieldCaster()

	return r
}

// API name: index_field
func (r *PutDataStreamMappings) IndexField(indexfield types.IndexFieldVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexField = indexfield.IndexFieldCaster()

	return r
}

// API name: _meta
func (r *PutDataStreamMappings) Meta_(metadata types.MetadataVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Meta_ = *metadata.MetadataCaster()

	return r
}

// API name: numeric_detection
func (r *PutDataStreamMappings) NumericDetection(numericdetection bool) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.NumericDetection = &numericdetection

	return r
}

// API name: properties
func (r *PutDataStreamMappings) Properties(properties map[string]types.Property) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Properties = properties
	return r
}

func (r *PutDataStreamMappings) AddProperty(key string, value types.PropertyVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]types.Property
	if r.req.Properties == nil {
		r.req.Properties = make(map[string]types.Property)
	} else {
		tmp = r.req.Properties
	}

	tmp[key] = *value.PropertyCaster()

	r.req.Properties = tmp
	return r
}

// API name: _routing
func (r *PutDataStreamMappings) Routing_(routing_ types.RoutingFieldVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Routing_ = routing_.RoutingFieldCaster()

	return r
}

// API name: runtime
func (r *PutDataStreamMappings) Runtime(runtime map[string]types.RuntimeField) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Runtime = runtime
	return r
}

func (r *PutDataStreamMappings) AddRuntime(key string, value types.RuntimeFieldVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	var tmp map[string]types.RuntimeField
	if r.req.Runtime == nil {
		r.req.Runtime = make(map[string]types.RuntimeField)
	} else {
		tmp = r.req.Runtime
	}

	tmp[key] = *value.RuntimeFieldCaster()

	r.req.Runtime = tmp
	return r
}

// API name: _size
func (r *PutDataStreamMappings) Size_(size_ types.SizeFieldVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Size_ = size_.SizeFieldCaster()

	return r
}

// API name: _source
func (r *PutDataStreamMappings) Source_(source_ types.SourceFieldVariant) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Source_ = source_.SourceFieldCaster()

	return r
}

// API name: subobjects
func (r *PutDataStreamMappings) Subobjects(subobjects subobjects.Subobjects) *PutDataStreamMappings {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Subobjects = &subobjects
	return r
}
