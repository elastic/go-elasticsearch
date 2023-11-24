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

// Updates attributes of an existing API key.
package updateapikey

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
)

const (
	idMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type UpdateApiKey struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	id string
}

// NewUpdateApiKey type alias for index.
type NewUpdateApiKey func(id string) *UpdateApiKey

// NewUpdateApiKeyFunc returns a new instance of UpdateApiKey with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewUpdateApiKeyFunc(tp elastictransport.Interface) NewUpdateApiKey {
	return func(id string) *UpdateApiKey {
		n := New(tp)

		n._id(id)

		return n
	}
}

// Updates attributes of an existing API key.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-api-key.html
func New(tp elastictransport.Interface) *UpdateApiKey {
	r := &UpdateApiKey{
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
func (r *UpdateApiKey) Raw(raw io.Reader) *UpdateApiKey {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *UpdateApiKey) Request(req *Request) *UpdateApiKey {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *UpdateApiKey) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for UpdateApiKey: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == idMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("api_key")
		path.WriteString("/")

		path.WriteString(r.id)

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
func (r UpdateApiKey) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the UpdateApiKey query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a updateapikey.Response
func (r UpdateApiKey) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the UpdateApiKey headers map.
func (r *UpdateApiKey) Header(key, value string) *UpdateApiKey {
	r.headers.Set(key, value)

	return r
}

// Id The ID of the API key to update.
// API Name: id
func (r *UpdateApiKey) _id(id string) *UpdateApiKey {
	r.paramSet |= idMask
	r.id = id

	return r
}

// Metadata Arbitrary metadata that you want to associate with the API key. It supports
// nested data structure. Within the metadata object, keys beginning with _ are
// reserved for system usage.
// API name: metadata
func (r *UpdateApiKey) Metadata(metadata types.Metadata) *UpdateApiKey {
	r.req.Metadata = metadata

	return r
}

// RoleDescriptors An array of role descriptors for this API key. This parameter is optional.
// When it is not specified or is an empty array, then the API key will have a
// point in time snapshot of permissions of the authenticated user. If you
// supply role descriptors then the resultant permissions would be an
// intersection of API keys permissions and authenticated user’s permissions
// thereby limiting the access scope for API keys. The structure of role
// descriptor is the same as the request for create role API. For more details,
// see create or update roles API.
// API name: role_descriptors
func (r *UpdateApiKey) RoleDescriptors(roledescriptors map[string]types.RoleDescriptor) *UpdateApiKey {

	r.req.RoleDescriptors = roledescriptors

	return r
}
