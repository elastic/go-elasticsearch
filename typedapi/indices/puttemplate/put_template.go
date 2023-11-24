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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

// Creates or updates an index template.
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

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	name string
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

// Creates or updates an index template.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
func New(tp elastictransport.Interface) *PutTemplate {
	r := &PutTemplate{
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for PutTemplate: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_template")
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
func (r PutTemplate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the PutTemplate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a puttemplate.Response
func (r PutTemplate) Do(ctx context.Context) (*Response, error) {

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

// FlatSettings If `true`, returns settings in flat format.
// API name: flat_settings
func (r *PutTemplate) FlatSettings(flatsettings bool) *PutTemplate {
	r.values.Set("flat_settings", strconv.FormatBool(flatsettings))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *PutTemplate) MasterTimeout(duration string) *PutTemplate {
	r.values.Set("master_timeout", duration)

	return r
}

// Timeout Period to wait for a response.
// If no response is received before the timeout expires, the request fails and
// returns an error.
// API name: timeout
func (r *PutTemplate) Timeout(duration string) *PutTemplate {
	r.values.Set("timeout", duration)

	return r
}

// Aliases Aliases for the index.
// API name: aliases
func (r *PutTemplate) Aliases(aliases map[string]types.Alias) *PutTemplate {

	r.req.Aliases = aliases

	return r
}

// IndexPatterns Array of wildcard expressions used to match the names
// of indices during creation.
// API name: index_patterns
func (r *PutTemplate) IndexPatterns(indexpatterns ...string) *PutTemplate {
	r.req.IndexPatterns = indexpatterns

	return r
}

// Mappings Mapping for fields in the index.
// API name: mappings
func (r *PutTemplate) Mappings(mappings *types.TypeMapping) *PutTemplate {

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
	r.req.Order = &order

	return r
}

// Settings Configuration options for the index.
// API name: settings
func (r *PutTemplate) Settings(settings map[string]json.RawMessage) *PutTemplate {

	r.req.Settings = settings

	return r
}

// Version Version number used to manage index templates externally. This number
// is not automatically generated by Elasticsearch.
// API name: version
func (r *PutTemplate) Version(versionnumber int64) *PutTemplate {
	r.req.Version = &versionnumber

	return r
}
