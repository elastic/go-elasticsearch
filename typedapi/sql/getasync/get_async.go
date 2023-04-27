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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

// Returns the current status and available results for an async SQL search or
// stored synchronous SQL search
package getasync

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
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetAsync struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int

	id string
}

// NewGetAsync type alias for index.
type NewGetAsync func(id string) *GetAsync

// NewGetAsyncFunc returns a new instance of GetAsync with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetAsyncFunc(tp elastictransport.Interface) NewGetAsync {
	return func(id string) *GetAsync {
		n := New(tp)

		n.Id(id)

		return n
	}
}

// Returns the current status and available results for an async SQL search or
// stored synchronous SQL search
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/get-async-sql-search-api.html
func New(tp elastictransport.Interface) *GetAsync {
	r := &GetAsync{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetAsync) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_sql")
		path.WriteString("/")
		path.WriteString("async")
		path.WriteString("/")

		path.WriteString(r.id)

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
func (r GetAsync) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetAsync query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getasync.Response
func (r GetAsync) Do(ctx context.Context) (*Response, error) {

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

// IsSuccess allows to run a query with a context and retrieve the result as a boolean.
// This only exists for endpoints without a request payload and allows for quick control flow.
func (r GetAsync) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetAsync headers map.
func (r *GetAsync) Header(key, value string) *GetAsync {
	r.headers.Set(key, value)

	return r
}

// Id The async search ID
// API Name: id
func (r *GetAsync) Id(v string) *GetAsync {
	r.paramSet |= idMask
	r.id = v

	return r
}

// Delimiter Separator for CSV results. The API only supports this parameter for CSV
// responses.
// API name: delimiter
func (r *GetAsync) Delimiter(v string) *GetAsync {
	r.values.Set("delimiter", v)

	return r
}

// Format Format for the response. You must specify a format using this parameter or
// the
// Accept HTTP header. If you specify both, the API uses this parameter.
// API name: format
func (r *GetAsync) Format(v string) *GetAsync {
	r.values.Set("format", v)

	return r
}

// KeepAlive Retention period for the search and its results. Defaults
// to the `keep_alive` period for the original SQL search.
// API name: keep_alive
func (r *GetAsync) KeepAlive(v string) *GetAsync {
	r.values.Set("keep_alive", v)

	return r
}

// WaitForCompletionTimeout Period to wait for complete results. Defaults to no timeout,
// meaning the request waits for complete search results.
// API name: wait_for_completion_timeout
func (r *GetAsync) WaitForCompletionTimeout(v string) *GetAsync {
	r.values.Set("wait_for_completion_timeout", v)

	return r
}
