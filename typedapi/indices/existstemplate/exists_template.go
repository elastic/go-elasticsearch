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

// Returns information about whether a particular index template exists.
package existstemplate

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

type ExistsTemplate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	name string
}

// NewExistsTemplate type alias for index.
type NewExistsTemplate func(name string) *ExistsTemplate

// NewExistsTemplateFunc returns a new instance of ExistsTemplate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewExistsTemplateFunc(tp elastictransport.Interface) NewExistsTemplate {
	return func(name string) *ExistsTemplate {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Returns information about whether a particular index template exists.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/indices-templates.html
func New(tp elastictransport.Interface) *ExistsTemplate {
	r := &ExistsTemplate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ExistsTemplate) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_template")
		path.WriteString("/")

		path.WriteString(r.name)

		method = http.MethodHead
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

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r ExistsTemplate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ExistsTemplate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a existstemplate.Response
func (r ExistsTemplate) Do(ctx context.Context) (bool, error) {
	return r.IsSuccess(ctx)
}

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r ExistsTemplate) IsSuccess(ctx context.Context) (bool, error) {
	res, err := r.Perform(ctx)

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

// Header set a key, value pair in the ExistsTemplate headers map.
func (r *ExistsTemplate) Header(key, value string) *ExistsTemplate {
	r.headers.Set(key, value)

	return r
}

// Name The comma separated names of the index templates
// API Name: name
func (r *ExistsTemplate) _name(name string) *ExistsTemplate {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// FlatSettings Return settings in flat format (default: false)
// API name: flat_settings
func (r *ExistsTemplate) FlatSettings(flatsettings bool) *ExistsTemplate {
	r.values.Set("flat_settings", strconv.FormatBool(flatsettings))

	return r
}

// Local Return local information, do not retrieve the state from master node
// (default: false)
// API name: local
func (r *ExistsTemplate) Local(local bool) *ExistsTemplate {
	r.values.Set("local", strconv.FormatBool(local))

	return r
}

// MasterTimeout Explicit operation timeout for connection to master node
// API name: master_timeout
func (r *ExistsTemplate) MasterTimeout(duration string) *ExistsTemplate {
	r.values.Set("master_timeout", duration)

	return r
}
