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

// Get a token.
//
// Create a bearer token for access without requiring basic authentication.
// The tokens are created by the Elasticsearch Token Service, which is
// automatically enabled when you configure TLS on the HTTP interface.
// Alternatively, you can explicitly enable the
// `xpack.security.authc.token.enabled` setting.
// When you are running in production mode, a bootstrap check prevents you from
// enabling the token service unless you also enable TLS on the HTTP interface.
//
// The get token API takes the same parameters as a typical OAuth 2.0 token API
// except for the use of a JSON request body.
//
// A successful get token API call returns a JSON structure that contains the
// access token, the amount of time (seconds) that the token expires in, the
// type, and the scope if available.
//
// The tokens returned by the get token API have a finite period of time for
// which they are valid and after that time period, they can no longer be used.
// That time period is defined by the `xpack.security.authc.token.timeout`
// setting.
// If you want to invalidate a token immediately, you can do so by using the
// invalidate token API.
package gettoken

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/accesstokengranttype"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GetToken struct {
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

// NewGetToken type alias for index.
type NewGetToken func() *GetToken

// NewGetTokenFunc returns a new instance of GetToken with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewGetTokenFunc(tp elastictransport.Interface) NewGetToken {
	return func() *GetToken {
		n := New(tp)

		return n
	}
}

// Get a token.
//
// Create a bearer token for access without requiring basic authentication.
// The tokens are created by the Elasticsearch Token Service, which is
// automatically enabled when you configure TLS on the HTTP interface.
// Alternatively, you can explicitly enable the
// `xpack.security.authc.token.enabled` setting.
// When you are running in production mode, a bootstrap check prevents you from
// enabling the token service unless you also enable TLS on the HTTP interface.
//
// The get token API takes the same parameters as a typical OAuth 2.0 token API
// except for the use of a JSON request body.
//
// A successful get token API call returns a JSON structure that contains the
// access token, the amount of time (seconds) that the token expires in, the
// type, and the scope if available.
//
// The tokens returned by the get token API have a finite period of time for
// which they are valid and after that time period, they can no longer be used.
// That time period is defined by the `xpack.security.authc.token.timeout`
// setting.
// If you want to invalidate a token immediately, you can do so by using the
// invalidate token API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-get-token.html
func New(tp elastictransport.Interface) *GetToken {
	r := &GetToken{
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
func (r *GetToken) Raw(raw io.Reader) *GetToken {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *GetToken) Request(req *Request) *GetToken {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *GetToken) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for GetToken: %w", err)
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
		path.WriteString("oauth2")
		path.WriteString("/")
		path.WriteString("token")

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
func (r GetToken) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.get_token")
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
		instrument.BeforeRequest(req, "security.get_token")
		if reader := instrument.RecordRequestBody(ctx, "security.get_token", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.get_token")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GetToken query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a gettoken.Response
func (r GetToken) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.get_token")
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

// Header set a key, value pair in the GetToken headers map.
func (r *GetToken) Header(key, value string) *GetToken {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GetToken) ErrorTrace(errortrace bool) *GetToken {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GetToken) FilterPath(filterpaths ...string) *GetToken {
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
func (r *GetToken) Human(human bool) *GetToken {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GetToken) Pretty(pretty bool) *GetToken {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// GrantType The type of grant.
// Supported grant types are: `password`, `_kerberos`, `client_credentials`, and
// `refresh_token`.
// API name: grant_type
func (r *GetToken) GrantType(granttype accesstokengranttype.AccessTokenGrantType) *GetToken {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.GrantType = &granttype

	return r
}

// KerberosTicket The base64 encoded kerberos ticket.
// If you specify the `_kerberos` grant type, this parameter is required.
// This parameter is not valid with any other supported grant type.
// API name: kerberos_ticket
func (r *GetToken) KerberosTicket(kerberosticket string) *GetToken {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.KerberosTicket = &kerberosticket

	return r
}

// Password The user's password.
// If you specify the `password` grant type, this parameter is required.
// This parameter is not valid with any other supported grant type.
// API name: password
func (r *GetToken) Password(password string) *GetToken {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Password = &password

	return r
}

// RefreshToken The string that was returned when you created the token, which enables you to
// extend its life.
// If you specify the `refresh_token` grant type, this parameter is required.
// This parameter is not valid with any other supported grant type.
// API name: refresh_token
func (r *GetToken) RefreshToken(refreshtoken string) *GetToken {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RefreshToken = &refreshtoken

	return r
}

// Scope The scope of the token.
// Currently tokens are only issued for a scope of FULL regardless of the value
// sent with the request.
// API name: scope
func (r *GetToken) Scope(scope string) *GetToken {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Scope = &scope

	return r
}

// Username The username that identifies the user.
// If you specify the `password` grant type, this parameter is required.
// This parameter is not valid with any other supported grant type.
// API name: username
func (r *GetToken) Username(username string) *GetToken {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Username = &username

	return r
}
