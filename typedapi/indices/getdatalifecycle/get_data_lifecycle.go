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

// Returns the data stream lifecycle of the selected data streams.
package getdatalifecycle

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/expandwildcard"
)

const (
	nameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetDataLifecycle struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	name string
}

// NewGetDataLifecycle type alias for index.
type NewGetDataLifecycle func(name string) *GetDataLifecycle

// NewGetDataLifecycleFunc returns a new instance of GetDataLifecycle with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetDataLifecycleFunc(tp elastictransport.Interface) NewGetDataLifecycle {
	return func(name string) *GetDataLifecycle {
		n := New(tp)

		n._name(name)

		return n
	}
}

// Returns the data stream lifecycle of the selected data streams.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/data-streams-get-lifecycle.html
func New(tp elastictransport.Interface) *GetDataLifecycle {
	r := &GetDataLifecycle{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetDataLifecycle) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == nameMask:
		path.WriteString("/")
		path.WriteString("_data_stream")
		path.WriteString("/")

		path.WriteString(r.name)
		path.WriteString("/")
		path.WriteString("_lifecycle")

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

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r GetDataLifecycle) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetDataLifecycle query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getdatalifecycle.Response
func (r GetDataLifecycle) Do(ctx context.Context) (*Response, error) {

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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetDataLifecycle) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetDataLifecycle headers map.
func (r *GetDataLifecycle) Header(key, value string) *GetDataLifecycle {
	r.headers.Set(key, value)

	return r
}

// Name Comma-separated list of data streams to limit the request.
// Supports wildcards (`*`).
// To target all data streams, omit this parameter or use `*` or `_all`.
// API Name: name
func (r *GetDataLifecycle) _name(name string) *GetDataLifecycle {
	r.paramSet |= nameMask
	r.name = name

	return r
}

// ExpandWildcards Type of data stream that wildcard patterns can match.
// Supports comma-separated values, such as `open,hidden`.
// Valid values are: `all`, `open`, `closed`, `hidden`, `none`.
// API name: expand_wildcards
func (r *GetDataLifecycle) ExpandWildcards(expandwildcards ...expandwildcard.ExpandWildcard) *GetDataLifecycle {
	tmp := []string{}
	for _, item := range expandwildcards {
		tmp = append(tmp, item.String())
	}
	r.values.Set("expand_wildcards", strings.Join(tmp, ","))

	return r
}

// IncludeDefaults If `true`, return all default settings in the response.
// API name: include_defaults
func (r *GetDataLifecycle) IncludeDefaults(includedefaults bool) *GetDataLifecycle {
	r.values.Set("include_defaults", strconv.FormatBool(includedefaults))

	return r
}
