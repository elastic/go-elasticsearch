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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

// Simulate an index.
//
// Get the index configuration that would be applied to the specified index from
// an existing index template.
package simulateindextemplate

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
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SimulateIndexTemplate struct {
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

// NewSimulateIndexTemplate type alias for index.
type NewSimulateIndexTemplate func(name string) *SimulateIndexTemplate

// NewSimulateIndexTemplateFunc returns a new instance of SimulateIndexTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSimulateIndexTemplateFunc(tp elastictransport.Interface) NewSimulateIndexTemplate {
	return func(name string) *SimulateIndexTemplate {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Simulate an index.
//
// Get the index configuration that would be applied to the specified index from
// an existing index template.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-indices-simulate-index-template
func New(tp elastictransport.Interface) *SimulateIndexTemplate {
	r := &SimulateIndexTemplate{
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
func (r *SimulateIndexTemplate) Raw(raw io.Reader) *SimulateIndexTemplate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SimulateIndexTemplate) Request(req *Request) *SimulateIndexTemplate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SimulateIndexTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for SimulateIndexTemplate: %w", err)
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
		path.WriteString("_index_template")
		path.WriteString("/")
		path.WriteString("_simulate_index")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", r.name)
		}
		path.WriteString(r.name)

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
func (r SimulateIndexTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.simulate_index_template")
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
		instrument.BeforeRequest(req, "indices.simulate_index_template")
		if reader := instrument.RecordRequestBody(ctx, "indices.simulate_index_template", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.simulate_index_template")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the SimulateIndexTemplate query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a simulateindextemplate.Response
func (r SimulateIndexTemplate) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.simulate_index_template")
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

// Header set a key, value pair in the SimulateIndexTemplate headers map.
func (r *SimulateIndexTemplate) Header(key, value string) *SimulateIndexTemplate {
	r.headers.Set(key, value)

	return r
}

// Name Name of the index to simulate
// API Name: name
func (r *SimulateIndexTemplate) _name(name string) *SimulateIndexTemplate {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Create Whether the index template we optionally defined in the body should only be
// dry-run added if new or can also replace an existing one
// API name: create
func (r *SimulateIndexTemplate) Create(create bool) *SimulateIndexTemplate {
	r.values.Set("create", strconv.FormatBool(create))

	return r
}

// Cause User defined reason for dry-run creating the new template for simulation
// purposes
// API name: cause
func (r *SimulateIndexTemplate) Cause(cause string) *SimulateIndexTemplate {
	r.values.Set("cause", cause)

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *SimulateIndexTemplate) MasterTimeout(duration string) *SimulateIndexTemplate {
	r.values.Set("master_timeout", duration)

	return r
}

// IncludeDefaults If true, returns all relevant default configurations for the index template.
// API name: include_defaults
func (r *SimulateIndexTemplate) IncludeDefaults(includedefaults bool) *SimulateIndexTemplate {
	r.values.Set("include_defaults", strconv.FormatBool(includedefaults))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *SimulateIndexTemplate) ErrorTrace(errortrace bool) *SimulateIndexTemplate {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *SimulateIndexTemplate) FilterPath(filterpaths ...string) *SimulateIndexTemplate {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and `"exists_time_in_millis":
// 3600000` for computers. When disabled the human readable values will be
// omitted. This makes sense for responses being consumed only by machines.
// API name: human
func (r *SimulateIndexTemplate) Human(human bool) *SimulateIndexTemplate {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use this
// option for debugging only.
// API name: pretty
func (r *SimulateIndexTemplate) Pretty(pretty bool) *SimulateIndexTemplate {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// API name: allow_auto_create
func (r *SimulateIndexTemplate) AllowAutoCreate(allowautocreate bool) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AllowAutoCreate = &allowautocreate

	return r
}

// An ordered list of component template names. Component templates are merged
// in the order specified, meaning that the last component template specified
// has the highest precedence.
// API name: composed_of
func (r *SimulateIndexTemplate) ComposedOf(composedofs ...string) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}
	for _, v := range composedofs {

		r.req.ComposedOf = append(r.req.ComposedOf, v)

	}
	return r
}

// Date and time when the index template was created. Only returned if the
// `human` query parameter is `true`.
// API name: created_date
func (r *SimulateIndexTemplate) CreatedDate(datetime types.DateTimeVariant) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.CreatedDate = *datetime.DateTimeCaster()

	return r
}

// Date and time when the index template was created, in milliseconds since the
// epoch.
// API name: created_date_millis
func (r *SimulateIndexTemplate) CreatedDateMillis(epochtimeunitmillis int64) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.CreatedDateMillis = &epochtimeunitmillis

	return r
}

// If this object is included, the template is used to create data streams and
// their backing indices. Supports an empty object. Data streams require a
// matching index template with a `data_stream` object.
// API name: data_stream
func (r *SimulateIndexTemplate) DataStream(datastream types.IndexTemplateDataStreamConfigurationVariant) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DataStream = datastream.IndexTemplateDataStreamConfigurationCaster()

	return r
}

// Marks this index template as deprecated. When creating or updating a
// non-deprecated index template that uses deprecated components, Elasticsearch
// will emit a deprecation warning.
// API name: deprecated
func (r *SimulateIndexTemplate) Deprecated(deprecated bool) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Deprecated = &deprecated

	return r
}

// A list of component template names that are allowed to be absent.
// API name: ignore_missing_component_templates
func (r *SimulateIndexTemplate) IgnoreMissingComponentTemplates(names ...string) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IgnoreMissingComponentTemplates = names

	return r
}

// Array of wildcard (`*`) expressions used to match the names of data streams
// and indices during creation.
// API name: index_patterns
func (r *SimulateIndexTemplate) IndexPatterns(names ...string) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.IndexPatterns = names

	return r
}

// Optional user metadata about the index template. May have any contents. This
// map is not automatically generated by Elasticsearch.
// API name: _meta
func (r *SimulateIndexTemplate) Meta_(metadata types.MetadataVariant) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Meta_ = *metadata.MetadataCaster()

	return r
}

// Date and time when the index template was last modified. Only returned if the
// `human` query parameter is `true`.
// API name: modified_date
func (r *SimulateIndexTemplate) ModifiedDate(datetime types.DateTimeVariant) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ModifiedDate = *datetime.DateTimeCaster()

	return r
}

// Date and time when the index template was last modified, in milliseconds
// since the epoch.
// API name: modified_date_millis
func (r *SimulateIndexTemplate) ModifiedDateMillis(epochtimeunitmillis int64) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ModifiedDateMillis = &epochtimeunitmillis

	return r
}

// Priority to determine index template precedence when a new data stream or
// index is created. The index template with the highest priority is chosen. If
// no priority is specified the template is treated as though it is of priority
// 0 (lowest priority). This number is not automatically generated by
// Elasticsearch.
// API name: priority
func (r *SimulateIndexTemplate) Priority(priority int64) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Priority = &priority

	return r
}

// Template to be applied. It may optionally include an `aliases`, `mappings`,
// or `settings` configuration.
// API name: template
func (r *SimulateIndexTemplate) Template(template types.IndexTemplateSummaryVariant) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Template = template.IndexTemplateSummaryCaster()

	return r
}

// Version number used to manage index templates externally. This number is not
// automatically generated by Elasticsearch.
// API name: version
func (r *SimulateIndexTemplate) Version(versionnumber int64) *SimulateIndexTemplate {
	// Initialize the request if it is not already initialized
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Version = &versionnumber

	return r
}
