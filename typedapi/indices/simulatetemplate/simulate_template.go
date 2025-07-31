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

// Simulate an index template.
// Get the index configuration that would be applied by a particular index
// template.
package simulatetemplate

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
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SimulateTemplate struct {
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

// NewSimulateTemplate type alias for index.
type NewSimulateTemplate func() *SimulateTemplate

// NewSimulateTemplateFunc returns a new instance of SimulateTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSimulateTemplateFunc(tp elastictransport.Interface) NewSimulateTemplate {
	return func() *SimulateTemplate {
		n := New(tp)

		return n
	}
}

// Simulate an index template.
// Get the index configuration that would be applied by a particular index
// template.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-simulate-template.html
func New(tp elastictransport.Interface) *SimulateTemplate {
	r := &SimulateTemplate{
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
func (r *SimulateTemplate) Raw(raw io.Reader) *SimulateTemplate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SimulateTemplate) Request(req *Request) *SimulateTemplate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SimulateTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for SimulateTemplate: %w", err)
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
		path.WriteString("_index_template")
		path.WriteString("/")
		path.WriteString("_simulate")

		method = http.MethodPost
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_index_template")
		path.WriteString("/")
		path.WriteString("_simulate")
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
func (r SimulateTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.simulate_template")
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
		instrument.BeforeRequest(req, "indices.simulate_template")
		if reader := instrument.RecordRequestBody(ctx, "indices.simulate_template", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.simulate_template")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the SimulateTemplate query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a simulatetemplate.Response
func (r SimulateTemplate) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.simulate_template")
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

// Header set a key, value pair in the SimulateTemplate headers map.
func (r *SimulateTemplate) Header(key, value string) *SimulateTemplate {
	r.headers.Set(key, value)

	return r
}

// Name Name of the index template to simulate. To test a template configuration
// before you add it to the cluster, omit
// this parameter and specify the template configuration in the request body.
// API Name: name
func (r *SimulateTemplate) Name(name string) *SimulateTemplate {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Create If true, the template passed in the body is only used if no existing
// templates match the same index patterns. If false, the simulation uses the
// template with the highest priority. Note that the template is not permanently
// added or updated in either case; it is only used for the simulation.
// API name: create
func (r *SimulateTemplate) Create(create bool) *SimulateTemplate {
	r.values.Set("create", strconv.FormatBool(create))

	return r
}

// Cause User defined reason for dry-run creating the new template for simulation
// purposes
// API name: cause
func (r *SimulateTemplate) Cause(cause string) *SimulateTemplate {
	r.values.Set("cause", cause)

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *SimulateTemplate) MasterTimeout(duration string) *SimulateTemplate {
	r.values.Set("master_timeout", duration)

	return r
}

// IncludeDefaults If true, returns all relevant default configurations for the index template.
// API name: include_defaults
func (r *SimulateTemplate) IncludeDefaults(includedefaults bool) *SimulateTemplate {
	r.values.Set("include_defaults", strconv.FormatBool(includedefaults))

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *SimulateTemplate) ErrorTrace(errortrace bool) *SimulateTemplate {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *SimulateTemplate) FilterPath(filterpaths ...string) *SimulateTemplate {
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
func (r *SimulateTemplate) Human(human bool) *SimulateTemplate {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *SimulateTemplate) Pretty(pretty bool) *SimulateTemplate {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AllowAutoCreate This setting overrides the value of the `action.auto_create_index` cluster
// setting.
// If set to `true` in a template, then indices can be automatically created
// using that template even if auto-creation of indices is disabled via
// `actions.auto_create_index`.
// If set to `false`, then indices or data streams matching the template must
// always be explicitly created, and may never be automatically created.
// API name: allow_auto_create
func (r *SimulateTemplate) AllowAutoCreate(allowautocreate bool) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.AllowAutoCreate = &allowautocreate

	return r
}

// ComposedOf An ordered list of component template names.
// Component templates are merged in the order specified, meaning that the last
// component template specified has the highest precedence.
// API name: composed_of
func (r *SimulateTemplate) ComposedOf(composedofs ...string) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.ComposedOf = composedofs

	return r
}

// DataStream If this object is included, the template is used to create data streams and
// their backing indices.
// Supports an empty object.
// Data streams require a matching index template with a `data_stream` object.
// API name: data_stream
func (r *SimulateTemplate) DataStream(datastream *types.DataStreamVisibility) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.DataStream = datastream

	return r
}

// Deprecated Marks this index template as deprecated. When creating or updating a
// non-deprecated index template
// that uses deprecated components, Elasticsearch will emit a deprecation
// warning.
// API name: deprecated
func (r *SimulateTemplate) Deprecated(deprecated bool) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Deprecated = &deprecated

	return r
}

// IgnoreMissingComponentTemplates The configuration option ignore_missing_component_templates can be used when
// an index template
// references a component template that might not exist
// API name: ignore_missing_component_templates
func (r *SimulateTemplate) IgnoreMissingComponentTemplates(ignoremissingcomponenttemplates ...string) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IgnoreMissingComponentTemplates = ignoremissingcomponenttemplates

	return r
}

// IndexPatterns Array of wildcard (`*`) expressions used to match the names of data streams
// and indices during creation.
// API name: index_patterns
func (r *SimulateTemplate) IndexPatterns(indices ...string) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IndexPatterns = indices

	return r
}

// Meta_ Optional user metadata about the index template.
// May have any contents.
// This map is not automatically generated by Elasticsearch.
// API name: _meta
func (r *SimulateTemplate) Meta_(metadata types.Metadata) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Meta_ = metadata

	return r
}

// Priority Priority to determine index template precedence when a new data stream or
// index is created.
// The index template with the highest priority is chosen.
// If no priority is specified the template is treated as though it is of
// priority 0 (lowest priority).
// This number is not automatically generated by Elasticsearch.
// API name: priority
func (r *SimulateTemplate) Priority(priority int64) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Priority = &priority

	return r
}

// Template Template to be applied.
// It may optionally include an `aliases`, `mappings`, or `settings`
// configuration.
// API name: template
func (r *SimulateTemplate) Template(template *types.IndexTemplateMapping) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Template = template

	return r
}

// Version Version number used to manage index templates externally.
// This number is not automatically generated by Elasticsearch.
// API name: version
func (r *SimulateTemplate) Version(versionnumber int64) *SimulateTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Version = &versionnumber

	return r
}
