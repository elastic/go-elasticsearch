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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

// Simulate matching the given index name against the index templates in the
// system
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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
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

	buf *gobytes.Buffer

	req *Request
	raw io.Reader

	paramSet int

	name string
}

// NewSimulateIndexTemplate type alias for index.
type NewSimulateIndexTemplate func(name string) *SimulateIndexTemplate

// NewSimulateIndexTemplateFunc returns a new instance of SimulateIndexTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSimulateIndexTemplateFunc(tp elastictransport.Interface) NewSimulateIndexTemplate {
	return func(name string) *SimulateIndexTemplate {
		n := New(tp)

		n.Name(name)

		return n
	}
}

// Simulate matching the given index name against the index templates in the
// system
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-templates.html
func New(tp elastictransport.Interface) *SimulateIndexTemplate {
	r := &SimulateIndexTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
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

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {
		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for SimulateIndexTemplate: %w", err)
		}

		r.buf.Write(data)
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_index_template")
		path.WriteString("/")
		path.WriteString("_simulate_index")
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
func (r SimulateIndexTemplate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the SimulateIndexTemplate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a simulateindextemplate.Response
func (r SimulateIndexTemplate) Do(ctx context.Context) (*Response, error) {

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

	return nil, errorResponse
}

// Header set a key, value pair in the SimulateIndexTemplate headers map.
func (r *SimulateIndexTemplate) Header(key, value string) *SimulateIndexTemplate {
	r.headers.Set(key, value)

	return r
}

// Name Index or template name to simulate
// API Name: name
func (r *SimulateIndexTemplate) Name(v string) *SimulateIndexTemplate {
	r.paramSet |= nameMask
	r.name = v

	return r
}

// Create If `true`, the template passed in the body is only used if no existing
// templates match the same index patterns. If `false`, the simulation uses
// the template with the highest priority. Note that the template is not
// permanently added or updated in either case; it is only used for the
// simulation.
// API name: create
func (r *SimulateIndexTemplate) Create(b bool) *SimulateIndexTemplate {
	r.values.Set("create", strconv.FormatBool(b))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received
// before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *SimulateIndexTemplate) MasterTimeout(v string) *SimulateIndexTemplate {
	r.values.Set("master_timeout", v)

	return r
}
