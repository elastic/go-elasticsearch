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

// Grant an API key.
//
// Create an API key on behalf of another user.
// This API is similar to the create API keys API, however it creates the API
// key for a user that is different than the user that runs the API.
// The caller must have authentication credentials for the user on whose behalf
// the API key will be created.
// It is not possible to use this API to create an API key without that user's
// credentials.
// The supported user authentication credential types are:
//
// * username and password
// * Elasticsearch access tokens
// * JWTs
//
// The user, for whom the authentication credentials is provided, can optionally
// "run as" (impersonate) another user.
// In this case, the API key will be created on behalf of the impersonated user.
//
// This API is intended be used by applications that need to create and manage
// API keys for end users, but cannot guarantee that those users have permission
// to create API keys on their own behalf.
// The API keys are created by the Elasticsearch API key service, which is
// automatically enabled.
//
// A successful grant API key API call returns a JSON structure that contains
// the API key, its unique id, and its name.
// If applicable, it also returns expiration information for the API key in
// milliseconds.
//
// By default, API keys never expire. You can specify expiration information
// when you create the API keys.
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
	"strconv"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/apikeygranttype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type GrantApiKey struct {
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

// Grant an API key.
//
// Create an API key on behalf of another user.
// This API is similar to the create API keys API, however it creates the API
// key for a user that is different than the user that runs the API.
// The caller must have authentication credentials for the user on whose behalf
// the API key will be created.
// It is not possible to use this API to create an API key without that user's
// credentials.
// The supported user authentication credential types are:
//
// * username and password
// * Elasticsearch access tokens
// * JWTs
//
// The user, for whom the authentication credentials is provided, can optionally
// "run as" (impersonate) another user.
// In this case, the API key will be created on behalf of the impersonated user.
//
// This API is intended be used by applications that need to create and manage
// API keys for end users, but cannot guarantee that those users have permission
// to create API keys on their own behalf.
// The API keys are created by the Elasticsearch API key service, which is
// automatically enabled.
//
// A successful grant API key API call returns a JSON structure that contains
// the API key, its unique id, and its name.
// If applicable, it also returns expiration information for the API key in
// milliseconds.
//
// By default, API keys never expire. You can specify expiration information
// when you create the API keys.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-grant-api-key.html
func New(tp elastictransport.Interface) *GrantApiKey {
	r := &GrantApiKey{
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

	if r.raw == nil && r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for GrantApiKey: %w", err)
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
func (r GrantApiKey) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.grant_api_key")
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
		instrument.BeforeRequest(req, "security.grant_api_key")
		if reader := instrument.RecordRequestBody(ctx, "security.grant_api_key", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.grant_api_key")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the GrantApiKey query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a grantapikey.Response
func (r GrantApiKey) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.grant_api_key")
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

// Header set a key, value pair in the GrantApiKey headers map.
func (r *GrantApiKey) Header(key, value string) *GrantApiKey {
	r.headers.Set(key, value)

	return r
}

// Refresh If 'true', Elasticsearch refreshes the affected shards to make this operation
// visible to search.
// If 'wait_for', it waits for a refresh to make this operation visible to
// search.
// If 'false', nothing is done with refreshes.
// API name: refresh
func (r *GrantApiKey) Refresh(refresh refresh.Refresh) *GrantApiKey {
	r.values.Set("refresh", refresh.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *GrantApiKey) ErrorTrace(errortrace bool) *GrantApiKey {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *GrantApiKey) FilterPath(filterpaths ...string) *GrantApiKey {
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
func (r *GrantApiKey) Human(human bool) *GrantApiKey {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *GrantApiKey) Pretty(pretty bool) *GrantApiKey {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AccessToken The user's access token.
// If you specify the `access_token` grant type, this parameter is required.
// It is not valid with other grant types.
// API name: access_token
func (r *GrantApiKey) AccessToken(accesstoken string) *GrantApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AccessToken = &accesstoken

	return r
}

// ApiKey The API key.
// API name: api_key
func (r *GrantApiKey) ApiKey(apikey *types.GrantApiKey) *GrantApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.ApiKey = *apikey

	return r
}

// GrantType The type of grant. Supported grant types are: `access_token`, `password`.
// API name: grant_type
func (r *GrantApiKey) GrantType(granttype apikeygranttype.ApiKeyGrantType) *GrantApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.GrantType = granttype

	return r
}

// Password The user's password.
// If you specify the `password` grant type, this parameter is required.
// It is not valid with other grant types.
// API name: password
func (r *GrantApiKey) Password(password string) *GrantApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Password = &password

	return r
}

// RunAs The name of the user to be impersonated.
// API name: run_as
func (r *GrantApiKey) RunAs(username string) *GrantApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.RunAs = &username

	return r
}

// Username The user name that identifies the user.
// If you specify the `password` grant type, this parameter is required.
// It is not valid with other grant types.
// API name: username
func (r *GrantApiKey) Username(username string) *GrantApiKey {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Username = &username

	return r
}
