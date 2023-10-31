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

// Creates or updates a component template
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	name string
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

// Creates or updates a component template
//
// https://www.elastic.co/guide/en/elasticsearch/reference/{branch}/indices-component-template.html
func New(tp elastictransport.Interface) *PutComponentTemplate {
	r := &PutComponentTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutComponentTemplate: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_component_template")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodPut
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
func (r PutComponentTemplate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutComponentTemplate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putcomponenttemplate.Response
func (r PutComponentTemplate) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the PutComponentTemplate headers map.
func (r *PutComponentTemplate) Header(key, value string) *PutComponentTemplate {
	r.headers.Set(key, value)

	return r
}

// Name Name of the component template to create.
// Elasticsearch includes the following built-in component templates:
// `logs-mappings`; 'logs-settings`; `metrics-mappings`;
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

// MasterTimeout Period to wait for a connection to the master node.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: master_timeout
func (r *PutComponentTemplate) MasterTimeout(duration string) *PutComponentTemplate {
	r.values.Set("master_timeout", duration)

	return r
}

// AllowAutoCreate This setting overrides the value of the `action.auto_create_index` cluster
// setting.
// If set to `true` in a template, then indices can be automatically created
// using that
// template even if auto-creation of indices is disabled via
// `actions.auto_create_index`.
// If set to `false` then data streams matching the template must always be
// explicitly created.
// API name: allow_auto_create
func (r *PutComponentTemplate) AllowAutoCreate(allowautocreate bool) *PutComponentTemplate {
	r.req.AllowAutoCreate = &allowautocreate

	return r
}

// Meta_ Optional user metadata about the component template.
// May have any contents. This map is not automatically generated by
// Elasticsearch.
// This information is stored in the cluster state, so keeping it short is
// preferable.
// To unset `_meta`, replace the template without specifying this information.
// API name: _meta
func (r *PutComponentTemplate) Meta_(metadata types.Metadata) *PutComponentTemplate {
	r.req.Meta_ = metadata

	return r
}

// Template The template to be applied which includes mappings, settings, or aliases
// configuration.
// API name: template
func (r *PutComponentTemplate) Template(template *types.IndexState) *PutComponentTemplate {

	r.req.Template = *template

	return r
}

// Version Version number used to manage component templates externally.
// This number isn't automatically generated or incremented by Elasticsearch.
// To unset a version, replace the template without specifying a version.
// API name: version
func (r *PutComponentTemplate) Version(versionnumber int64) *PutComponentTemplate {
	r.req.Version = &versionnumber

	return r
}
