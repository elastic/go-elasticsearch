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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

// Create or update an index template.
// Index templates define settings, mappings, and aliases that can be applied
// automatically to new indices.
package putindextemplate

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

type PutIndexTemplate struct {
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

// NewPutIndexTemplate type alias for index.
type NewPutIndexTemplate func(name string) *PutIndexTemplate

// NewPutIndexTemplateFunc returns a new instance of PutIndexTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutIndexTemplateFunc(tp elastictransport.Interface) NewPutIndexTemplate {
	return func(name string) *PutIndexTemplate {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Create or update an index template.
// Index templates define settings, mappings, and aliases that can be applied
// automatically to new indices.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-put-template.html
func New(tp elastictransport.Interface) *PutIndexTemplate {
	r := &PutIndexTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),

		req: NewRequest(),
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
func (r *PutIndexTemplate) Raw(raw io.Reader) *PutIndexTemplate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutIndexTemplate) Request(req *Request) *PutIndexTemplate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutIndexTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutIndexTemplate: %w", err)
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
func (r PutIndexTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.put_index_template")
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
		instrument.BeforeRequest(req, "indices.put_index_template")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_index_template", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_index_template")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutIndexTemplate query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putindextemplate.Response
func (r PutIndexTemplate) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_index_template")
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

// Header set a key, value pair in the PutIndexTemplate headers map.
func (r *PutIndexTemplate) Header(key, value string) *PutIndexTemplate {
	r.headers.Set(key, value)

	return r
}

// Name Index or template name
// API Name: name
func (r *PutIndexTemplate) _name(name string) *PutIndexTemplate {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Create If `true`, this request cannot replace or update existing index templates.
// API name: create
func (r *PutIndexTemplate) Create(create bool) *PutIndexTemplate {
	r.values.Set("create", strconv.FormatBool(create))

	return r
}

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *PutIndexTemplate) MasterTimeout(duration string) *PutIndexTemplate {
	r.values.Set("master_timeout", duration)

	return r
}

// Cause User defined reason for creating/updating the index template
// API name: cause
func (r *PutIndexTemplate) Cause(cause string) *PutIndexTemplate {
	r.values.Set("cause", cause)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutIndexTemplate) ErrorTrace(errortrace bool) *PutIndexTemplate {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutIndexTemplate) FilterPath(filterpaths ...string) *PutIndexTemplate {
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
func (r *PutIndexTemplate) Human(human bool) *PutIndexTemplate {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutIndexTemplate) Pretty(pretty bool) *PutIndexTemplate {
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
func (r *PutIndexTemplate) AllowAutoCreate(allowautocreate bool) *PutIndexTemplate {
	r.req.AllowAutoCreate = &allowautocreate

	return r
}

// ComposedOf An ordered list of component template names.
// Component templates are merged in the order specified, meaning that the last
// component template specified has the highest precedence.
// API name: composed_of
func (r *PutIndexTemplate) ComposedOf(composedofs ...string) *PutIndexTemplate {
	r.req.ComposedOf = composedofs

	return r
}

// DataStream If this object is included, the template is used to create data streams and
// their backing indices.
// Supports an empty object.
// Data streams require a matching index template with a `data_stream` object.
// API name: data_stream
func (r *PutIndexTemplate) DataStream(datastream *types.DataStreamVisibility) *PutIndexTemplate {

	r.req.DataStream = datastream

	return r
}

// Deprecated Marks this index template as deprecated. When creating or updating a
// non-deprecated index template
// that uses deprecated components, Elasticsearch will emit a deprecation
// warning.
// API name: deprecated
func (r *PutIndexTemplate) Deprecated(deprecated bool) *PutIndexTemplate {
	r.req.Deprecated = &deprecated

	return r
}

// IgnoreMissingComponentTemplates The configuration option ignore_missing_component_templates can be used when
// an index template
// references a component template that might not exist
// API name: ignore_missing_component_templates
func (r *PutIndexTemplate) IgnoreMissingComponentTemplates(ignoremissingcomponenttemplates ...string) *PutIndexTemplate {
	r.req.IgnoreMissingComponentTemplates = ignoremissingcomponenttemplates

	return r
}

// IndexPatterns Name of the index template to create.
// API name: index_patterns
func (r *PutIndexTemplate) IndexPatterns(indices ...string) *PutIndexTemplate {
	r.req.IndexPatterns = indices

	return r
}

// Meta_ Optional user metadata about the index template.
// May have any contents.
// This map is not automatically generated by Elasticsearch.
// API name: _meta
func (r *PutIndexTemplate) Meta_(metadata types.Metadata) *PutIndexTemplate {
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
func (r *PutIndexTemplate) Priority(priority int64) *PutIndexTemplate {

	r.req.Priority = &priority

	return r
}

// Template Template to be applied.
// It may optionally include an `aliases`, `mappings`, or `settings`
// configuration.
// API name: template
func (r *PutIndexTemplate) Template(template *types.IndexTemplateMapping) *PutIndexTemplate {

	r.req.Template = template

	return r
}

// Version Version number used to manage index templates externally.
// This number is not automatically generated by Elasticsearch.
// API name: version
func (r *PutIndexTemplate) Version(versionnumber int64) *PutIndexTemplate {
	r.req.Version = &versionnumber

	return r
}
