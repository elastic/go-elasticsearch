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

// Prepare OpenID connect authentication.
//
// Create an oAuth 2.0 authentication request as a URL string based on the
// configuration of the OpenID Connect authentication realm in Elasticsearch.
//
// The response of this API is a URL pointing to the Authorization Endpoint of
// the configured OpenID Connect Provider, which can be used to redirect the
// browser of the user in order to continue the authentication process.
//
// Elasticsearch exposes all the necessary OpenID Connect related functionality
// with the OpenID Connect APIs.
// These APIs are used internally by Kibana in order to provide OpenID Connect
// based authentication, but can also be used by other, custom web applications
// or other clients.
package oidcprepareauthentication

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

type OidcPrepareAuthentication struct {
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

// NewOidcPrepareAuthentication type alias for index.
type NewOidcPrepareAuthentication func() *OidcPrepareAuthentication

// NewOidcPrepareAuthenticationFunc returns a new instance of OidcPrepareAuthentication with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewOidcPrepareAuthenticationFunc(tp elastictransport.Interface) NewOidcPrepareAuthentication {
	return func() *OidcPrepareAuthentication {
		n := New(tp)

		return n
	}
}

// Prepare OpenID connect authentication.
//
// Create an oAuth 2.0 authentication request as a URL string based on the
// configuration of the OpenID Connect authentication realm in Elasticsearch.
//
// The response of this API is a URL pointing to the Authorization Endpoint of
// the configured OpenID Connect Provider, which can be used to redirect the
// browser of the user in order to continue the authentication process.
//
// Elasticsearch exposes all the necessary OpenID Connect related functionality
// with the OpenID Connect APIs.
// These APIs are used internally by Kibana in order to provide OpenID Connect
// based authentication, but can also be used by other, custom web applications
// or other clients.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-oidc-prepare-authentication.html
func New(tp elastictransport.Interface) *OidcPrepareAuthentication {
	r := &OidcPrepareAuthentication{
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
func (r *OidcPrepareAuthentication) Raw(raw io.Reader) *OidcPrepareAuthentication {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *OidcPrepareAuthentication) Request(req *Request) *OidcPrepareAuthentication {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *OidcPrepareAuthentication) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for OidcPrepareAuthentication: %w", err)
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
		path.WriteString("prepare")

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
func (r OidcPrepareAuthentication) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.oidc_prepare_authentication")
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
		instrument.BeforeRequest(req, "security.oidc_prepare_authentication")
		if reader := instrument.RecordRequestBody(ctx, "security.oidc_prepare_authentication", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.oidc_prepare_authentication")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the OidcPrepareAuthentication query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a oidcprepareauthentication.Response
func (r OidcPrepareAuthentication) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.oidc_prepare_authentication")
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

// Header set a key, value pair in the OidcPrepareAuthentication headers map.
func (r *OidcPrepareAuthentication) Header(key, value string) *OidcPrepareAuthentication {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *OidcPrepareAuthentication) ErrorTrace(errortrace bool) *OidcPrepareAuthentication {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *OidcPrepareAuthentication) FilterPath(filterpaths ...string) *OidcPrepareAuthentication {
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
func (r *OidcPrepareAuthentication) Human(human bool) *OidcPrepareAuthentication {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *OidcPrepareAuthentication) Pretty(pretty bool) *OidcPrepareAuthentication {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Iss In the case of a third party initiated single sign on, this is the issuer
// identifier for the OP that the RP is to send the authentication request to.
// It cannot be specified when *realm* is specified.
// One of *realm* or *iss* is required.
// API name: iss
func (r *OidcPrepareAuthentication) Iss(iss string) *OidcPrepareAuthentication {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Iss = &iss

	return r
}

// LoginHint In the case of a third party initiated single sign on, it is a string value
// that is included in the authentication request as the *login_hint* parameter.
// This parameter is not valid when *realm* is specified.
// API name: login_hint
func (r *OidcPrepareAuthentication) LoginHint(loginhint string) *OidcPrepareAuthentication {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.LoginHint = &loginhint

	return r
}

// Nonce The value used to associate a client session with an ID token and to mitigate
// replay attacks.
// If the caller of the API does not provide a value, Elasticsearch will
// generate one with sufficient entropy and return it in the response.
// API name: nonce
func (r *OidcPrepareAuthentication) Nonce(nonce string) *OidcPrepareAuthentication {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Nonce = &nonce

	return r
}

// Realm The name of the OpenID Connect realm in Elasticsearch the configuration of
// which should be used in order to generate the authentication request.
// It cannot be specified when *iss* is specified.
// One of *realm* or *iss* is required.
// API name: realm
func (r *OidcPrepareAuthentication) Realm(realm string) *OidcPrepareAuthentication {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Realm = &realm

	return r
}

// State The value used to maintain state between the authentication request and the
// response, typically used as a Cross-Site Request Forgery mitigation.
// If the caller of the API does not provide a value, Elasticsearch will
// generate one with sufficient entropy and return it in the response.
// API name: state
func (r *OidcPrepareAuthentication) State(state string) *OidcPrepareAuthentication {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.State = &state

	return r
}
