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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

// Authenticate OpenID Connect.
//
// Exchange an OpenID Connect authentication response message for an
// Elasticsearch internal access token and refresh token that can be
// subsequently used for authentication.
//
// Elasticsearch exposes all the necessary OpenID Connect related functionality
// with the OpenID Connect APIs.
// These APIs are used internally by Kibana in order to provide OpenID Connect
// based authentication, but can also be used by other, custom web applications
// or other clients.
package oidcauthenticate

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

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type OidcAuthenticate struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewOidcAuthenticate type alias for index.
type NewOidcAuthenticate func() *OidcAuthenticate

// NewOidcAuthenticateFunc returns a new instance of OidcAuthenticate with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewOidcAuthenticateFunc(tp elastictransport.Interface) NewOidcAuthenticate {
	return func() *OidcAuthenticate {
		n := New(tp)

		return n
	}
}

// Authenticate OpenID Connect.
//
// Exchange an OpenID Connect authentication response message for an
// Elasticsearch internal access token and refresh token that can be
// subsequently used for authentication.
//
// Elasticsearch exposes all the necessary OpenID Connect related functionality
// with the OpenID Connect APIs.
// These APIs are used internally by Kibana in order to provide OpenID Connect
// based authentication, but can also be used by other, custom web applications
// or other clients.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-oidc-authenticate.html
func New(tp elastictransport.Interface) *OidcAuthenticate {
	r := &OidcAuthenticate{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),

		buf: gobytes.NewBuffer(nil),
	}

	if instrumented, ok := r.transport.(elastictransport.Instrumented); ok {
		if instrument := instrumented.InstrumentationEnabled(); instrument != nil {
			r.instrument = instrument
		}
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *OidcAuthenticate) Raw(raw io.Reader) *OidcAuthenticate {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *OidcAuthenticate) Request(req *Request) *OidcAuthenticate {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *OidcAuthenticate) HttpRequest(ctx context.Context) (*http.Request, error) {
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for OidcAuthenticate: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("oidc")
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
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.raw)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.raw)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.raw != nil {
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
func (r OidcAuthenticate) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.oidc_authenticate")
			defer instrument.Close(ctx)
		}
	}
	if ctx == nil {
		ctx = providedCtx
	}

	req, err := r.HttpRequest(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.BeforeRequest(req, "security.oidc_authenticate")
		if reader := instrument.RecordRequestBody(ctx, "security.oidc_authenticate", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.oidc_authenticate")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the OidcAuthenticate query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a oidcauthenticate.Response
func (r OidcAuthenticate) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.oidc_authenticate")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
				instrument.RecordError(ctx, err)
			}
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.RecordError(ctx, errorResponse)
	}
	return nil, errorResponse
}

// Header set a key, value pair in the OidcAuthenticate headers map.
func (r *OidcAuthenticate) Header(key, value string) *OidcAuthenticate {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *OidcAuthenticate) ErrorTrace(errortrace bool) *OidcAuthenticate {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *OidcAuthenticate) FilterPath(filterpaths ...string) *OidcAuthenticate {
	tmp := []string{}
	for _, item := range filterpaths {
		tmp = append(tmp, fmt.Sprintf("%v", item))
	}
	r.values.Set("filter_path", strings.Join(tmp, ","))

	return r
}

// Human When set to `true` will return statistics in a format suitable for humans.
// For example `"exists_time": "1h"` for humans and
// `"eixsts_time_in_millis": 3600000` for computers. When disabled the human
// readable values will be omitted. This makes sense for responses being
// consumed
// only by machines.
// API name: human
func (r *OidcAuthenticate) Human(human bool) *OidcAuthenticate {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *OidcAuthenticate) Pretty(pretty bool) *OidcAuthenticate {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Nonce Associate a client session with an ID token and mitigate replay attacks.
// This value needs to be the same as the one that was provided to the
// `/_security/oidc/prepare` API or the one that was generated by Elasticsearch
// and included in the response to that call.
// API name: nonce
func (r *OidcAuthenticate) Nonce(nonce string) *OidcAuthenticate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Nonce = nonce

	return r
}

// Realm The name of the OpenID Connect realm.
// This property is useful in cases where multiple realms are defined.
// API name: realm
func (r *OidcAuthenticate) Realm(realm string) *OidcAuthenticate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Realm = &realm

	return r
}

// RedirectUri The URL to which the OpenID Connect Provider redirected the User Agent in
// response to an authentication request after a successful authentication.
// This URL must be provided as-is (URL encoded), taken from the body of the
// response or as the value of a location header in the response from the OpenID
// Connect Provider.
// API name: redirect_uri
func (r *OidcAuthenticate) RedirectUri(redirecturi string) *OidcAuthenticate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RedirectUri = redirecturi

	return r
}

// State Maintain state between the authentication request and the response.
// This value needs to be the same as the one that was provided to the
// `/_security/oidc/prepare` API or the one that was generated by Elasticsearch
// and included in the response to that call.
// API name: state
func (r *OidcAuthenticate) State(state string) *OidcAuthenticate {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.State = state

	return r
}
