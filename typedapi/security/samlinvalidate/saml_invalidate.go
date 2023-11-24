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

// Consumes a SAML LogoutRequest
package samlinvalidate

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

type SamlInvalidate struct {
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

// NewSamlInvalidate type alias for index.
type NewSamlInvalidate func() *SamlInvalidate

// NewSamlInvalidateFunc returns a new instance of SamlInvalidate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSamlInvalidateFunc(tp elastictransport.Interface) NewSamlInvalidate {
	return func() *SamlInvalidate {
		n := New(tp)

		return n
	}
}

// Consumes a SAML LogoutRequest
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-invalidate.html
func New(tp elastictransport.Interface) *SamlInvalidate {
	r := &SamlInvalidate{
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
func (r *SamlInvalidate) Raw(raw io.Reader) *SamlInvalidate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SamlInvalidate) Request(req *Request) *SamlInvalidate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SamlInvalidate) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for SamlInvalidate: %w", err)
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
		path.WriteString("invalidate")

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
func (r SamlInvalidate) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the SamlInvalidate query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a samlinvalidate.Response
func (r SamlInvalidate) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the SamlInvalidate headers map.
func (r *SamlInvalidate) Header(key, value string) *SamlInvalidate {
	r.headers.Set(key, value)

	return r
}

// Acs The Assertion Consumer Service URL that matches the one of the SAML realm in
// Elasticsearch that should be used. You must specify either this parameter or
// the realm parameter.
// API name: acs
func (r *SamlInvalidate) Acs(acs string) *SamlInvalidate {

	r.req.Acs = &acs

	return r
}

// QueryString The query part of the URL that the user was redirected to by the SAML IdP to
// initiate the Single Logout.
// This query should include a single parameter named SAMLRequest that contains
// a SAML logout request that is deflated and Base64 encoded.
// If the SAML IdP has signed the logout request, the URL should include two
// extra parameters named SigAlg and Signature that contain the algorithm used
// for the signature and the signature value itself.
// In order for Elasticsearch to be able to verify the IdP’s signature, the
// value of the query_string field must be an exact match to the string provided
// by the browser.
// The client application must not attempt to parse or process the string in any
// way.
// API name: query_string
func (r *SamlInvalidate) QueryString(querystring string) *SamlInvalidate {

	r.req.QueryString = querystring

	return r
}

// Realm The name of the SAML realm in Elasticsearch the configuration. You must
// specify either this parameter or the acs parameter.
// API name: realm
func (r *SamlInvalidate) Realm(realm string) *SamlInvalidate {

	r.req.Realm = &realm

	return r
}
