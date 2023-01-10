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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


// Returns an index template.
package getindextemplate

import (
	gobytes "bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetIndexTemplate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	name string
}

// NewGetIndexTemplate type alias for index.
type NewGetIndexTemplate func() *GetIndexTemplate

// NewGetIndexTemplateFunc returns a new instance of GetIndexTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetIndexTemplateFunc(tp elastictransport.Interface) NewGetIndexTemplate {
	return func() *GetIndexTemplate {
		n := New(tp)

		return n
	}
}

// Returns an index template.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-templates.html
func New(tp elastictransport.Interface) *GetIndexTemplate {
	r := &GetIndexTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetIndexTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_index_template")

		method = http.MethodGet
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_index_template")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodGet
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

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Do runs the http.Request through the provided transport.
func (r GetIndexTemplate) Do(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetIndexTemplate query execution: %w", err)
	}

	return res, nil
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetIndexTemplate) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Do(ctx)

	if err != nil {
		return false, err
	}
	io.Copy(ioutil.Discard, res.Body)
	err = res.Body.Close()
	if err != nil {
		return false, err
	}

	if res.StatusCode >= 200 && res.StatusCode < 300 {
		return true, nil
	}

	return false, nil
}

// Header set a key, value pair in the GetIndexTemplate headers map.
func (r *GetIndexTemplate) Header(key, value string) *GetIndexTemplate {
	r.headers.Set(key, value)

	return r
}

// Name Comma-separated list of index template names used to limit the request.
// Wildcard (*) expressions are supported.
// API Name: name
func (r *GetIndexTemplate) Name(v string) *GetIndexTemplate {
	r.paramSet |= nameMask
	r.name = v

	return r
}

// Local If true, the request retrieves information from the local node only. Defaults
// to false, which means information is retrieved from the master node.
// API name: local
func (r *GetIndexTemplate) Local(b bool) *GetIndexTemplate {
	r.values.Set("local", strconv.FormatBool(b))

	return r
}

// FlatSettings If true, returns settings in flat format.
// API name: flat_settings
func (r *GetIndexTemplate) FlatSettings(b bool) *GetIndexTemplate {
	r.values.Set("flat_settings", strconv.FormatBool(b))

	return r
}

// MasterTimeout Period to wait for a connection to the master node. If no response is
// received before the timeout expires, the request fails and returns an error.
// API name: master_timeout
func (r *GetIndexTemplate) MasterTimeout(value string) *GetIndexTemplate {
	r.values.Set("master_timeout", value)

	return r
}
