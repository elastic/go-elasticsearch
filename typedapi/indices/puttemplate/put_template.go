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

// Create or update a legacy index template.
// Index templates define settings, mappings, and aliases that can be applied
// automatically to new indices.
// Elasticsearch applies templates to new indices based on an index pattern that
// matches the index name.
//
// IMPORTANT: This documentation is about legacy index templates, which are
// deprecated and will be replaced by the composable templates introduced in
// Elasticsearch 7.8.
//
// Composable templates always take precedence over legacy templates.
// If no composable template matches a new index, matching legacy templates are
// applied according to their order.
//
// Index templates are only applied during index creation.
// Changes to index templates do not affect existing indices.
// Settings and mappings specified in create index API requests override any
// settings or mappings specified in an index template.
//
// You can use C-style `/* *\/` block comments in index templates.
// You can include comments anywhere in the request body, except before the
// opening curly bracket.
//
// **Indices matching multiple templates**
//
// Multiple index templates can potentially match an index, in this case, both
// the settings and mappings are merged into the final configuration of the
// index.
// The order of the merging can be controlled using the order parameter, with
// lower order being applied first, and higher orders overriding them.
// NOTE: Multiple matching templates with the same order value will result in a
// non-deterministic merging order.
package puttemplate

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

type PutTemplate struct {
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

// NewPutTemplate type alias for index.
type NewPutTemplate func(name string) *PutTemplate

// NewPutTemplateFunc returns a new instance of PutTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutTemplateFunc(tp elastictransport.Interface) NewPutTemplate {
	return func(name string) *PutTemplate {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Create or update a legacy index template.
// Index templates define settings, mappings, and aliases that can be applied
// automatically to new indices.
// Elasticsearch applies templates to new indices based on an index pattern that
// matches the index name.
//
// IMPORTANT: This documentation is about legacy index templates, which are
// deprecated and will be replaced by the composable templates introduced in
// Elasticsearch 7.8.
//
// Composable templates always take precedence over legacy templates.
// If no composable template matches a new index, matching legacy templates are
// applied according to their order.
//
// Index templates are only applied during index creation.
// Changes to index templates do not affect existing indices.
// Settings and mappings specified in create index API requests override any
// settings or mappings specified in an index template.
//
// You can use C-style `/* *\/` block comments in index templates.
// You can include comments anywhere in the request body, except before the
// opening curly bracket.
//
// **Indices matching multiple templates**
//
// Multiple index templates can potentially match an index, in this case, both
// the settings and mappings are merged into the final configuration of the
// index.
// The order of the merging can be controlled using the order parameter, with
// lower order being applied first, and higher orders overriding them.
// NOTE: Multiple matching templates with the same order value will result in a
// non-deterministic merging order.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates-v1.html
func New(tp elastictransport.Interface) *PutTemplate {
	r := &PutTemplate{
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
func (r *PutTemplate) Raw(raw io.Reader) *PutTemplate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutTemplate) Request(req *Request) *PutTemplate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutTemplate: %w", err)
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
		path.WriteString("_template")
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
func (r PutTemplate) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "indices.put_template")
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
		instrument.BeforeRequest(req, "indices.put_template")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_template", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_template")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutTemplate query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttemplate.Response
func (r PutTemplate) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_template")
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

// Header set a key, value pair in the PutTemplate headers map.
func (r *PutTemplate) Header(key, value string) *PutTemplate {
	r.headers.Set(key, value)

	return r
}

// Name The name of the template
// API Name: name
func (r *PutTemplate) _name(name string) *PutTemplate {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// Create If true, this request cannot replace or update existing index templates.
// API name: create
func (r *PutTemplate) Create(create bool) *PutTemplate {
	r.values.Set("create", strconv.FormatBool(create))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *PutTemplate) MasterTimeout(duration string) *PutTemplate {
	r.values.Set("master_timeout", duration)

	return r
}

// Cause User defined reason for creating/updating the index template
// API name: cause
func (r *PutTemplate) Cause(cause string) *PutTemplate {
	r.values.Set("cause", cause)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutTemplate) ErrorTrace(errortrace bool) *PutTemplate {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutTemplate) FilterPath(filterpaths ...string) *PutTemplate {
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
func (r *PutTemplate) Human(human bool) *PutTemplate {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutTemplate) Pretty(pretty bool) *PutTemplate {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Aliases Aliases for the index.
// API name: aliases
func (r *PutTemplate) Aliases(aliases map[string]types.Alias) *PutTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Aliases = aliases

	return r
}

// IndexPatterns Array of wildcard expressions used to match the names
// of indices during creation.
// API name: index_patterns
func (r *PutTemplate) IndexPatterns(indexpatterns ...string) *PutTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.IndexPatterns = indexpatterns

	return r
}

// Mappings Mapping for fields in the index.
// API name: mappings
func (r *PutTemplate) Mappings(mappings *types.TypeMapping) *PutTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Mappings = mappings

	return r
}

// Order Order in which Elasticsearch applies this template if index
// matches multiple templates.
//
// Templates with lower 'order' values are merged first. Templates with higher
// 'order' values are merged later, overriding templates with lower values.
// API name: order
func (r *PutTemplate) Order(order int) *PutTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Order = &order

	return r
}

// Settings Configuration options for the index.
// API name: settings
func (r *PutTemplate) Settings(settings *types.IndexSettings) *PutTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Settings = settings

	return r
}

// Version Version number used to manage index templates externally. This number
// is not automatically generated by Elasticsearch.
// To unset a version, replace the template without specifying one.
// API name: version
func (r *PutTemplate) Version(versionnumber int64) *PutTemplate {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Version = &versionnumber

	return r
}
