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

// Invalidate a token.
//
// The access tokens returned by the get token API have a finite period of time
// for which they are valid.
// After that time period, they can no longer be used.
// The time period is defined by the `xpack.security.authc.token.timeout`
// setting.
//
// The refresh tokens returned by the get token API are only valid for 24 hours.
// They can also be used exactly once.
// If you want to invalidate one or more access or refresh tokens immediately,
// use this invalidate token API.
//
// NOTE: While all parameters are optional, at least one of them is required.
// More specifically, either one of `token` or `refresh_token` parameters is
// required.
// If none of these two are specified, then `realm_name` and/or `username` need
// to be specified.
package invalidatetoken

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

type InvalidateToken struct {
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

// NewInvalidateToken type alias for index.
type NewInvalidateToken func() *InvalidateToken

// NewInvalidateTokenFunc returns a new instance of InvalidateToken with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewInvalidateTokenFunc(tp elastictransport.Interface) NewInvalidateToken {
	return func() *InvalidateToken {
		n := New(tp)

		return n
	}
}

// Invalidate a token.
//
// The access tokens returned by the get token API have a finite period of time
// for which they are valid.
// After that time period, they can no longer be used.
// The time period is defined by the `xpack.security.authc.token.timeout`
// setting.
//
// The refresh tokens returned by the get token API are only valid for 24 hours.
// They can also be used exactly once.
// If you want to invalidate one or more access or refresh tokens immediately,
// use this invalidate token API.
//
// NOTE: While all parameters are optional, at least one of them is required.
// More specifically, either one of `token` or `refresh_token` parameters is
// required.
// If none of these two are specified, then `realm_name` and/or `username` need
// to be specified.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-invalidate-token.html
func New(tp elastictransport.Interface) *InvalidateToken {
	r := &InvalidateToken{
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
func (r *InvalidateToken) Raw(raw io.Reader) *InvalidateToken {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *InvalidateToken) Request(req *Request) *InvalidateToken {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *InvalidateToken) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for InvalidateToken: %w", err)
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

		method = http.MethodDelete
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
func (r InvalidateToken) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.invalidate_token")
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
		instrument.BeforeRequest(req, "security.invalidate_token")
		if reader := instrument.RecordRequestBody(ctx, "security.invalidate_token", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.invalidate_token")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the InvalidateToken query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a invalidatetoken.Response
func (r InvalidateToken) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.invalidate_token")
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

// Header set a key, value pair in the InvalidateToken headers map.
func (r *InvalidateToken) Header(key, value string) *InvalidateToken {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *InvalidateToken) ErrorTrace(errortrace bool) *InvalidateToken {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *InvalidateToken) FilterPath(filterpaths ...string) *InvalidateToken {
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
func (r *InvalidateToken) Human(human bool) *InvalidateToken {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *InvalidateToken) Pretty(pretty bool) *InvalidateToken {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// RealmName The name of an authentication realm.
// This parameter cannot be used with either `refresh_token` or `token`.
// API name: realm_name
func (r *InvalidateToken) RealmName(name string) *InvalidateToken {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RealmName = &name

	return r
}

// RefreshToken A refresh token.
// This parameter cannot be used if any of `refresh_token`, `realm_name`, or
// `username` are used.
// API name: refresh_token
func (r *InvalidateToken) RefreshToken(refreshtoken string) *InvalidateToken {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.RefreshToken = &refreshtoken

	return r
}

// Token An access token.
// This parameter cannot be used if any of `refresh_token`, `realm_name`, or
// `username` are used.
// API name: token
func (r *InvalidateToken) Token(token string) *InvalidateToken {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Token = &token

	return r
}

// Username The username of a user.
// This parameter cannot be used with either `refresh_token` or `token`.
// API name: username
func (r *InvalidateToken) Username(username string) *InvalidateToken {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Username = &username

	return r
}
