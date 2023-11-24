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

// Exchanges a SAML Response message for an Elasticsearch access token and
// refresh token pair
package samlauthenticate

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type SamlAuthenticate struct {
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

// NewSamlAuthenticate type alias for index.
type NewSamlAuthenticate func() *SamlAuthenticate

// NewSamlAuthenticateFunc returns a new instance of SamlAuthenticate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSamlAuthenticateFunc(tp elastictransport.Interface) NewSamlAuthenticate {
	return func() *SamlAuthenticate {
		n := New(tp)

		return n
	}
}

// Exchanges a SAML Response message for an Elasticsearch access token and
// refresh token pair
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-authenticate.html
func New(tp elastictransport.Interface) *SamlAuthenticate {
	r := &SamlAuthenticate{
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
func (r *SamlAuthenticate) Raw(raw io.Reader) *SamlAuthenticate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SamlAuthenticate) Request(req *Request) *SamlAuthenticate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SamlAuthenticate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for SamlAuthenticate: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("saml")
		path.WriteString("/")
		path.WriteString("authenticate")

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
func (r SamlAuthenticate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the SamlAuthenticate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a samlauthenticate.Response
func (r SamlAuthenticate) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the SamlAuthenticate headers map.
func (r *SamlAuthenticate) Header(key, value string) *SamlAuthenticate {
	r.headers.Set(key, value)

	return r
}

// Content The SAML response as it was sent by the user’s browser, usually a Base64
// encoded XML document.
// API name: content
func (r *SamlAuthenticate) Content(content string) *SamlAuthenticate {

	r.req.Content = content

	return r
}

// Ids A json array with all the valid SAML Request Ids that the caller of the API
// has for the current user.
// API name: ids
func (r *SamlAuthenticate) Ids(ids ...string) *SamlAuthenticate {
	r.req.Ids = ids

	return r
}

// Realm The name of the realm that should authenticate the SAML response. Useful in
// cases where many SAML realms are defined.
// API name: realm
func (r *SamlAuthenticate) Realm(realm string) *SamlAuthenticate {

	r.req.Realm = &realm

	return r
}
