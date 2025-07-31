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

// Prepare SAML authentication.
//
// Create a SAML authentication request (`<AuthnRequest>`) as a URL string based
// on the configuration of the respective SAML realm in Elasticsearch.
//
// NOTE: This API is intended for use by custom web applications other than
// Kibana.
// If you are using Kibana, refer to the documentation for configuring SAML
// single-sign-on on the Elastic Stack.
//
// This API returns a URL pointing to the SAML Identity Provider.
// You can use the URL to redirect the browser of the user in order to continue
// the authentication process.
// The URL includes a single parameter named `SAMLRequest`, which contains a
// SAML Authentication request that is deflated and Base64 encoded.
// If the configuration dictates that SAML authentication requests should be
// signed, the URL has two extra parameters named `SigAlg` and `Signature`.
// These parameters contain the algorithm used for the signature and the
// signature value itself.
// It also returns a random string that uniquely identifies this SAML
// Authentication request.
// The caller of this API needs to store this identifier as it needs to be used
// in a following step of the authentication process.
package samlprepareauthentication

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

type SamlPrepareAuthentication struct {
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

// NewSamlPrepareAuthentication type alias for index.
type NewSamlPrepareAuthentication func() *SamlPrepareAuthentication

// NewSamlPrepareAuthenticationFunc returns a new instance of SamlPrepareAuthentication with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewSamlPrepareAuthenticationFunc(tp elastictransport.Interface) NewSamlPrepareAuthentication {
	return func() *SamlPrepareAuthentication {
		n := New(tp)

		return n
	}
}

// Prepare SAML authentication.
//
// Create a SAML authentication request (`<AuthnRequest>`) as a URL string based
// on the configuration of the respective SAML realm in Elasticsearch.
//
// NOTE: This API is intended for use by custom web applications other than
// Kibana.
// If you are using Kibana, refer to the documentation for configuring SAML
// single-sign-on on the Elastic Stack.
//
// This API returns a URL pointing to the SAML Identity Provider.
// You can use the URL to redirect the browser of the user in order to continue
// the authentication process.
// The URL includes a single parameter named `SAMLRequest`, which contains a
// SAML Authentication request that is deflated and Base64 encoded.
// If the configuration dictates that SAML authentication requests should be
// signed, the URL has two extra parameters named `SigAlg` and `Signature`.
// These parameters contain the algorithm used for the signature and the
// signature value itself.
// It also returns a random string that uniquely identifies this SAML
// Authentication request.
// The caller of this API needs to store this identifier as it needs to be used
// in a following step of the authentication process.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-saml-prepare-authentication.html
func New(tp elastictransport.Interface) *SamlPrepareAuthentication {
	r := &SamlPrepareAuthentication{
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
func (r *SamlPrepareAuthentication) Raw(raw io.Reader) *SamlPrepareAuthentication {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *SamlPrepareAuthentication) Request(req *Request) *SamlPrepareAuthentication {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *SamlPrepareAuthentication) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for SamlPrepareAuthentication: %w", err)
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
		path.WriteString("saml")
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
func (r SamlPrepareAuthentication) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.saml_prepare_authentication")
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
		instrument.BeforeRequest(req, "security.saml_prepare_authentication")
		if reader := instrument.RecordRequestBody(ctx, "security.saml_prepare_authentication", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.saml_prepare_authentication")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the SamlPrepareAuthentication query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a samlprepareauthentication.Response
func (r SamlPrepareAuthentication) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.saml_prepare_authentication")
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

// Header set a key, value pair in the SamlPrepareAuthentication headers map.
func (r *SamlPrepareAuthentication) Header(key, value string) *SamlPrepareAuthentication {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *SamlPrepareAuthentication) ErrorTrace(errortrace bool) *SamlPrepareAuthentication {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *SamlPrepareAuthentication) FilterPath(filterpaths ...string) *SamlPrepareAuthentication {
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
func (r *SamlPrepareAuthentication) Human(human bool) *SamlPrepareAuthentication {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *SamlPrepareAuthentication) Pretty(pretty bool) *SamlPrepareAuthentication {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Acs The Assertion Consumer Service URL that matches the one of the SAML realms in
// Elasticsearch.
// The realm is used to generate the authentication request. You must specify
// either this parameter or the `realm` parameter.
// API name: acs
func (r *SamlPrepareAuthentication) Acs(acs string) *SamlPrepareAuthentication {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Acs = &acs

	return r
}

// Realm The name of the SAML realm in Elasticsearch for which the configuration is
// used to generate the authentication request.
// You must specify either this parameter or the `acs` parameter.
// API name: realm
func (r *SamlPrepareAuthentication) Realm(realm string) *SamlPrepareAuthentication {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Realm = &realm

	return r
}

// RelayState A string that will be included in the redirect URL that this API returns as
// the `RelayState` query parameter.
// If the Authentication Request is signed, this value is used as part of the
// signature computation.
// API name: relay_state
func (r *SamlPrepareAuthentication) RelayState(relaystate string) *SamlPrepareAuthentication {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RelayState = &relaystate

	return r
}
