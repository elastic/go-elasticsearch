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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

// Retrieves information for one or more API keys.
package getapikey

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
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	paramSet int
}

// NewGetApiKey type alias for index.
type NewGetApiKey func() *GetApiKey

// NewGetApiKeyFunc returns a new instance of GetApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetApiKeyFunc(tp elastictransport.Interface) NewGetApiKey {
	return func() *GetApiKey {
		n := New(tp)

		return n
	}
}

// Retrieves information for one or more API keys.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-api-key.html
func New(tp elastictransport.Interface) *GetApiKey {
	r := &GetApiKey{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),
	}

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("api_key")

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
func (r GetApiKey) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GetApiKey query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a getapikey.Response
func (r GetApiKey) Do(ctx context.Context) (*Response, error) {

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
func (r GetApiKey) IsSuccess(ctx context.Context) (bool, error) {
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

// Header set a key, value pair in the GetApiKey headers map.
func (r *GetApiKey) Header(key, value string) *GetApiKey {
	r.headers.Set(key, value)

	return r
}

// Id API key id of the API key to be retrieved
// API name: id
func (r *GetApiKey) Id(v string) *GetApiKey {
	r.values.Set("id", v)

	return r
}

// Name API key name of the API key to be retrieved
// API name: name
func (r *GetApiKey) Name(v string) *GetApiKey {
	r.values.Set("name", v)

	return r
}

// Owner flag to query API keys owned by the currently authenticated user
// API name: owner
func (r *GetApiKey) Owner(b bool) *GetApiKey {
	r.values.Set("owner", strconv.FormatBool(b))

	return r
}

// RealmName realm name of the user who created this API key to be retrieved
// API name: realm_name
func (r *GetApiKey) RealmName(v string) *GetApiKey {
	r.values.Set("realm_name", v)

	return r
}

// Username user name of the user who created this API key to be retrieved
// API name: username
func (r *GetApiKey) Username(v string) *GetApiKey {
	r.values.Set("username", v)

	return r
}

// WithLimitedBy Return the snapshot of the owner user's role descriptors
// associated with the API key. An API key's actual
// permission is the intersection of its assigned role
// descriptors and the owner user's role descriptors.
// API name: with_limited_by
func (r *GetApiKey) WithLimitedBy(b bool) *GetApiKey {
	r.values.Set("with_limited_by", strconv.FormatBool(b))

	return r
}
