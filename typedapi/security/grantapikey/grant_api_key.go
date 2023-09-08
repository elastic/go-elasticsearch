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
// https://github.com/elastic/elasticsearch-specification/tree/b89646a75dd9e8001caf92d22bd8b3704c59dfdf

// Creates an API key on behalf of another user.
package grantapikey

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/apikeygranttype"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GrantApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int
}

// NewGrantApiKey type alias for index.
type NewGrantApiKey func() *GrantApiKey

// NewGrantApiKeyFunc returns a new instance of GrantApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGrantApiKeyFunc(tp elastictransport.Interface) NewGrantApiKey {
	return func() *GrantApiKey {
		n := New(tp)

		return n
	}
}

// Creates an API key on behalf of another user.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-grant-api-key.html
func New(tp elastictransport.Interface) *GrantApiKey {
	r := &GrantApiKey{
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
func (r *GrantApiKey) Raw(raw io.Reader) *GrantApiKey {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GrantApiKey) Request(req *Request) *GrantApiKey {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GrantApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for GrantApiKey: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("api_key")
		path.WriteString("/")
		path.WriteString("grant")

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
func (r GrantApiKey) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the GrantApiKey query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a grantapikey.Response
func (r GrantApiKey) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the GrantApiKey headers map.
func (r *GrantApiKey) Header(key, value string) *GrantApiKey {
	r.headers.Set(key, value)

	return r
}

// AccessToken The user’s access token.
// If you specify the `access_token` grant type, this parameter is required.
// It is not valid with other grant types.
// API name: access_token
func (r *GrantApiKey) AccessToken(accesstoken string) *GrantApiKey {

	r.req.AccessToken = &accesstoken

	return r
}

// ApiKey Defines the API key.
// API name: api_key
func (r *GrantApiKey) ApiKey(apikey *types.GrantApiKey) *GrantApiKey {

	r.req.ApiKey = *apikey

	return r
}

// GrantType The type of grant. Supported grant types are: `access_token`, `password`.
// API name: grant_type
func (r *GrantApiKey) GrantType(granttype apikeygranttype.ApiKeyGrantType) *GrantApiKey {
	r.req.GrantType = granttype

	return r
}

// Password The user’s password. If you specify the `password` grant type, this parameter
// is required.
// It is not valid with other grant types.
// API name: password
func (r *GrantApiKey) Password(password string) *GrantApiKey {
	r.req.Password = &password

	return r
}

// RunAs The name of the user to be impersonated.
// API name: run_as
func (r *GrantApiKey) RunAs(username string) *GrantApiKey {
	r.req.RunAs = &username

	return r
}

// Username The user name that identifies the user.
// If you specify the `password` grant type, this parameter is required.
// It is not valid with other grant types.
// API name: username
func (r *GrantApiKey) Username(username string) *GrantApiKey {
	r.req.Username = &username

	return r
}
