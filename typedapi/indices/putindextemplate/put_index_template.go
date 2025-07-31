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

// Create or update an index template.
// Index templates define settings, mappings, and aliases that can be applied
// automatically to new indices.
//
// Elasticsearch applies templates to new indices based on an wildcard pattern
// that matches the index name.
// Index templates are applied during data stream or index creation.
// For data streams, these settings and mappings are applied when the stream's
// backing indices are created.
// Settings and mappings specified in a create index API request override any
// settings or mappings specified in an index template.
// Changes to index templates do not affect existing indices, including the
// existing backing indices of a data stream.
//
// You can use C-style `/* *\/` block comments in index templates.
// You can include comments anywhere in the request body, except before the
// opening curly bracket.
//
// **Multiple matching templates**
//
// If multiple index templates match the name of a new index or data stream, the
// template with the highest priority is used.
//
// Multiple templates with overlapping index patterns at the same priority are
// not allowed and an error will be thrown when attempting to create a template
// matching an existing index template at identical priorities.
//
// **Composing aliases, mappings, and settings**
//
// When multiple component templates are specified in the `composed_of` field
// for an index template, they are merged in the order specified, meaning that
// later component templates override earlier component templates.
// Any mappings, settings, or aliases from the parent index template are merged
// in next.
// Finally, any configuration on the index request itself is merged.
// Mapping definitions are merged recursively, which means that later mapping
// components can introduce new field mappings and update the mapping
// configuration.
// If a field mapping is already contained in an earlier component, its
// definition will be completely overwritten by the later one.
// This recursive merging strategy applies not only to field mappings, but also
// root options like `dynamic_templates` and `meta`.
// If an earlier component contains a `dynamic_templates` block, then by default
// new `dynamic_templates` entries are appended onto the end.
// If an entry already exists with the same key, then it is overwritten by the
// new definition.
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
// Elasticsearch applies templates to new indices based on an wildcard pattern
// that matches the index name.
// Index templates are applied during data stream or index creation.
// For data streams, these settings and mappings are applied when the stream's
// backing indices are created.
// Settings and mappings specified in a create index API request override any
// settings or mappings specified in an index template.
// Changes to index templates do not affect existing indices, including the
// existing backing indices of a data stream.
//
// You can use C-style `/* *\/` block comments in index templates.
// You can include comments anywhere in the request body, except before the
// opening curly bracket.
//
// **Multiple matching templates**
//
// If multiple index templates match the name of a new index or data stream, the
// template with the highest priority is used.
//
// Multiple templates with overlapping index patterns at the same priority are
// not allowed and an error will be thrown when attempting to create a template
// matching an existing index template at identical priorities.
//
// **Composing aliases, mappings, and settings**
//
// When multiple component templates are specified in the `composed_of` field
// for an index template, they are merged in the order specified, meaning that
// later component templates override earlier component templates.
// Any mappings, settings, or aliases from the parent index template are merged
// in next.
// Finally, any configuration on the index request itself is merged.
// Mapping definitions are merged recursively, which means that later mapping
// components can introduce new field mappings and update the mapping
// configuration.
// If a field mapping is already contained in an earlier component, its
// definition will be completely overwritten by the later one.
// This recursive merging strategy applies not only to field mappings, but also
// root options like `dynamic_templates` and `meta`.
// If an earlier component contains a `dynamic_templates` block, then by default
// new `dynamic_templates` entries are appended onto the end.
// If an entry already exists with the same key, then it is overwritten by the
// new definition.
//
// https://www.elastic.co/docs/api/doc/elasticsearch/v8/operation/operation-indices-put-index-template
func New(tp elastictransport.Interface) *PutIndexTemplate {
	r := &PutIndexTemplate{
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
func (r *PutIndexTemplate) ComposedOf(composedofs ...string) *PutIndexTemplate {
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
func (r *PutIndexTemplate) DataStream(datastream *types.DataStreamVisibility) *PutIndexTemplate {
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
func (r *PutIndexTemplate) Deprecated(deprecated bool) *PutIndexTemplate {
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
func (r *PutIndexTemplate) IgnoreMissingComponentTemplates(ignoremissingcomponenttemplates ...string) *PutIndexTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IgnoreMissingComponentTemplates = ignoremissingcomponenttemplates

	return r
}

// IndexPatterns Name of the index template to create.
// API name: index_patterns
func (r *PutIndexTemplate) IndexPatterns(indices ...string) *PutIndexTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IndexPatterns = indices

	return r
}

// Meta_ Optional user metadata about the index template.
// It may have any contents.
// It is not automatically generated or used by Elasticsearch.
// This user-defined object is stored in the cluster state, so keeping it short
// is preferable
// To unset the metadata, replace the template without specifying it.
// API name: _meta
func (r *PutIndexTemplate) Meta_(metadata types.Metadata) *PutIndexTemplate {
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
func (r *PutIndexTemplate) Priority(priority int64) *PutIndexTemplate {
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
func (r *PutIndexTemplate) Template(template *types.IndexTemplateMapping) *PutIndexTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Template = template

	return r
}

// Version Version number used to manage index templates externally.
// This number is not automatically generated by Elasticsearch.
// External systems can use these version numbers to simplify template
// management.
// To unset a version, replace the template without specifying one.
// API name: version
func (r *PutIndexTemplate) Version(versionnumber int64) *PutIndexTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Version = &versionnumber

	return r
}
