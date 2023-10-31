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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Simulate resolving the given template name or body
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	name string
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

// Simulate resolving the given template name or body
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-templates.html
func New(tp elastictransport.Interface) *SimulateTemplate {
	r := &SimulateTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for SimulateTemplate: %w", err)
		}

		r.buf.Write(data)

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

		path.WriteString(r.name)

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
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
func (r SimulateTemplate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the SimulateTemplate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a simulatetemplate.Response
func (r SimulateTemplate) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
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

// API name: allow_auto_create
func (r *SimulateTemplate) AllowAutoCreate(allowautocreate bool) *SimulateTemplate {
	r.req.AllowAutoCreate = &allowautocreate

	return r
}

// ComposedOf An ordered list of component template names.
// Component templates are merged in the order specified, meaning that the last
// component template specified has the highest precedence.
// API name: composed_of
func (r *SimulateTemplate) ComposedOf(composedofs ...string) *SimulateTemplate {
	r.req.ComposedOf = composedofs

	return r
}

// DataStream If this object is included, the template is used to create data streams and
// their backing indices.
// Supports an empty object.
// Data streams require a matching index template with a `data_stream` object.
// API name: data_stream
func (r *SimulateTemplate) DataStream(datastream *types.IndexTemplateDataStreamConfiguration) *SimulateTemplate {

	r.req.DataStream = datastream

	return r
}

// IndexPatterns Name of the index template.
// API name: index_patterns
func (r *SimulateTemplate) IndexPatterns(names ...string) *SimulateTemplate {
	r.req.IndexPatterns = names

	return r
}

// Meta_ Optional user metadata about the index template. May have any contents.
// This map is not automatically generated by Elasticsearch.
// API name: _meta
func (r *SimulateTemplate) Meta_(metadata types.Metadata) *SimulateTemplate {
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

	r.req.Priority = &priority

	return r
}

// Template Template to be applied.
// It may optionally include an `aliases`, `mappings`, or `settings`
// configuration.
// API name: template
func (r *SimulateTemplate) Template(template *types.IndexTemplateSummary) *SimulateTemplate {

	r.req.Template = template

	return r
}

// Version Version number used to manage index templates externally.
// This number is not automatically generated by Elasticsearch.
// API name: version
func (r *SimulateTemplate) Version(versionnumber int64) *SimulateTemplate {
	r.req.Version = &versionnumber

	return r
}
