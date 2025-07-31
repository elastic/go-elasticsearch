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

// Create or update a component template.
// Component templates are building blocks for constructing index templates that
// specify index mappings, settings, and aliases.
//
// An index template can be composed of multiple component templates.
// To use a component template, specify it in an index template’s `composed_of`
// list.
// Component templates are only applied to new data streams and indices as part
// of a matching index template.
//
// Settings and mappings specified directly in the index template or the create
// index request override any settings or mappings specified in a component
// template.
//
// Component templates are only used during index creation.
// For data streams, this includes data stream creation and the creation of a
// stream’s backing indices.
// Changes to component templates do not affect existing indices, including a
// stream’s backing indices.
//
// You can use C-style `/* *\/` block comments in component templates.
// You can include comments anywhere in the request body except before the
// opening curly bracket.
//
// **Applying component templates**
//
// You cannot directly apply a component template to a data stream or index.
// To be applied, a component template must be included in an index template's
// `composed_of` list.
package putcomponenttemplate

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

type PutComponentTemplate struct {
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

// NewPutComponentTemplate type alias for index.
type NewPutComponentTemplate func(name string) *PutComponentTemplate

// NewPutComponentTemplateFunc returns a new instance of PutComponentTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutComponentTemplateFunc(tp elastictransport.Interface) NewPutComponentTemplate {
	return func(name string) *PutComponentTemplate {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Create or update a component template.
// Component templates are building blocks for constructing index templates that
// specify index mappings, settings, and aliases.
//
// An index template can be composed of multiple component templates.
// To use a component template, specify it in an index template’s `composed_of`
// list.
// Component templates are only applied to new data streams and indices as part
// of a matching index template.
//
// Settings and mappings specified directly in the index template or the create
// index request override any settings or mappings specified in a component
// template.
//
// Component templates are only used during index creation.
// For data streams, this includes data stream creation and the creation of a
// stream’s backing indices.
// Changes to component templates do not affect existing indices, including a
// stream’s backing indices.
//
// You can use C-style `/* *\/` block comments in component templates.
// You can include comments anywhere in the request body except before the
// opening curly bracket.
//
// **Applying component templates**
//
// You cannot directly apply a component template to a data stream or index.
// To be applied, a component template must be included in an index template's
// `composed_of` list.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-component-template.html
func New(tp elastictransport.Interface) *PutComponentTemplate {
	r := &PutComponentTemplate{
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
func (r *PutComponentTemplate) Raw(raw io.Reader) *PutComponentTemplate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutComponentTemplate) Request(req *Request) *PutComponentTemplate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutComponentTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutComponentTemplate: %w", err)
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
		path.WriteString("_component_template")
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
func (r PutComponentTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "cluster.put_component_template")
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
		instrument.BeforeRequest(req, "cluster.put_component_template")
		if reader := instrument.RecordRequestBody(ctx, "cluster.put_component_template", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.put_component_template")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutComponentTemplate query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putcomponenttemplate.Response
func (r PutComponentTemplate) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.put_component_template")
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

// Header set a key, value pair in the PutComponentTemplate headers map.
func (r *PutComponentTemplate) Header(key, value string) *PutComponentTemplate {
	r.headers.Set(key, value)

	return r
}

// Name Name of the component template to create.
// Elasticsearch includes the following built-in component templates:
// `logs-mappings`; `logs-settings`; `metrics-mappings`;
// `metrics-settings`;`synthetics-mapping`; `synthetics-settings`.
// Elastic Agent uses these templates to configure backing indices for its data
// streams.
// If you use Elastic Agent and want to overwrite one of these templates, set
// the `version` for your replacement template higher than the current version.
// If you don’t use Elastic Agent and want to disable all built-in component and
// index templates, set `stack.templates.enabled` to `false` using the cluster
// update settings API.
// API Name: name
func (r *PutComponentTemplate) _name(name string) *PutComponentTemplate {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Create If `true`, this request cannot replace or update existing component
// templates.
// API name: create
func (r *PutComponentTemplate) Create(create bool) *PutComponentTemplate {
	r.values.Set("create", strconv.FormatBool(create))

	return r
}

// Cause User defined reason for create the component template.
// API name: cause
func (r *PutComponentTemplate) Cause(cause string) *PutComponentTemplate {
	r.values.Set("cause", cause)

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *PutComponentTemplate) MasterTimeout(duration string) *PutComponentTemplate {
	r.values.Set("master_timeout", duration)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutComponentTemplate) ErrorTrace(errortrace bool) *PutComponentTemplate {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutComponentTemplate) FilterPath(filterpaths ...string) *PutComponentTemplate {
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
func (r *PutComponentTemplate) Human(human bool) *PutComponentTemplate {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutComponentTemplate) Pretty(pretty bool) *PutComponentTemplate {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Deprecated Marks this index template as deprecated. When creating or updating a
// non-deprecated index template
// that uses deprecated components, Elasticsearch will emit a deprecation
// warning.
// API name: deprecated
func (r *PutComponentTemplate) Deprecated(deprecated bool) *PutComponentTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Deprecated = &deprecated

	return r
}

// Meta_ Optional user metadata about the component template.
// It may have any contents. This map is not automatically generated by
// Elasticsearch.
// This information is stored in the cluster state, so keeping it short is
// preferable.
// To unset `_meta`, replace the template without specifying this information.
// API name: _meta
func (r *PutComponentTemplate) Meta_(metadata types.Metadata) *PutComponentTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Meta_ = metadata

	return r
}

// Template The template to be applied which includes mappings, settings, or aliases
// configuration.
// API name: template
func (r *PutComponentTemplate) Template(template *types.IndexState) *PutComponentTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Template = *template

	return r
}

// Version Version number used to manage component templates externally.
// This number isn't automatically generated or incremented by Elasticsearch.
// To unset a version, replace the template without specifying a version.
// API name: version
func (r *PutComponentTemplate) Version(versionnumber int64) *PutComponentTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Version = &versionnumber

	return r
}
