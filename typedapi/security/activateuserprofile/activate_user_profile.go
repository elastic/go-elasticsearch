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

// Activate a user profile.
//
// Create or update a user profile on behalf of another user.
//
// NOTE: The user profile feature is designed only for use by Kibana and
// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
// Individual users and external applications should not call this API directly.
// The calling application must have either an `access_token` or a combination
// of `username` and `password` for the user that the profile document is
// intended for.
// Elastic reserves the right to change or remove this feature in future
// releases without prior notice.
//
// This API creates or updates a profile document for end users with information
// that is extracted from the user's authentication object including `username`,
// `full_name,` `roles`, and the authentication realm.
// For example, in the JWT `access_token` case, the profile user's `username` is
// extracted from the JWT token claim pointed to by the `claims.principal`
// setting of the JWT realm that authenticated the token.
//
// When updating a profile document, the API enables the document if it was
// disabled.
// Any updates do not change existing content for either the `labels` or `data`
// fields.
package activateuserprofile

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/granttype"
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ActivateUserProfile struct {
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

// NewActivateUserProfile type alias for index.
type NewActivateUserProfile func() *ActivateUserProfile

// NewActivateUserProfileFunc returns a new instance of ActivateUserProfile with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewActivateUserProfileFunc(tp elastictransport.Interface) NewActivateUserProfile {
	return func() *ActivateUserProfile {
		n := New(tp)

		return n
	}
}

// Activate a user profile.
//
// Create or update a user profile on behalf of another user.
//
// NOTE: The user profile feature is designed only for use by Kibana and
// Elastic's Observability, Enterprise Search, and Elastic Security solutions.
// Individual users and external applications should not call this API directly.
// The calling application must have either an `access_token` or a combination
// of `username` and `password` for the user that the profile document is
// intended for.
// Elastic reserves the right to change or remove this feature in future
// releases without prior notice.
//
// This API creates or updates a profile document for end users with information
// that is extracted from the user's authentication object including `username`,
// `full_name,` `roles`, and the authentication realm.
// For example, in the JWT `access_token` case, the profile user's `username` is
// extracted from the JWT token claim pointed to by the `claims.principal`
// setting of the JWT realm that authenticated the token.
//
// When updating a profile document, the API enables the document if it was
// disabled.
// Any updates do not change existing content for either the `labels` or `data`
// fields.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-activate-user-profile.html
func New(tp elastictransport.Interface) *ActivateUserProfile {
	r := &ActivateUserProfile{
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
func (r *ActivateUserProfile) Raw(raw io.Reader) *ActivateUserProfile {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ActivateUserProfile) Request(req *Request) *ActivateUserProfile {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ActivateUserProfile) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ActivateUserProfile: %w", err)
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
		path.WriteString("profile")
		path.WriteString("/")
		path.WriteString("_activate")

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
func (r ActivateUserProfile) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.activate_user_profile")
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
		instrument.BeforeRequest(req, "security.activate_user_profile")
		if reader := instrument.RecordRequestBody(ctx, "security.activate_user_profile", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.activate_user_profile")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the ActivateUserProfile query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a activateuserprofile.Response
func (r ActivateUserProfile) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.activate_user_profile")
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

// Header set a key, value pair in the ActivateUserProfile headers map.
func (r *ActivateUserProfile) Header(key, value string) *ActivateUserProfile {
	r.headers.Set(key, value)

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *ActivateUserProfile) ErrorTrace(errortrace bool) *ActivateUserProfile {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *ActivateUserProfile) FilterPath(filterpaths ...string) *ActivateUserProfile {
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
func (r *ActivateUserProfile) Human(human bool) *ActivateUserProfile {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *ActivateUserProfile) Pretty(pretty bool) *ActivateUserProfile {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// AccessToken The user's Elasticsearch access token or JWT.
// Both `access` and `id` JWT token types are supported and they depend on the
// underlying JWT realm configuration.
// If you specify the `access_token` grant type, this parameter is required.
// It is not valid with other grant types.
// API name: access_token
func (r *ActivateUserProfile) AccessToken(accesstoken string) *ActivateUserProfile {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.AccessToken = &accesstoken

	return r
}

// GrantType The type of grant.
// API name: grant_type
func (r *ActivateUserProfile) GrantType(granttype granttype.GrantType) *ActivateUserProfile {
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
func (r *ActivateUserProfile) Password(password string) *ActivateUserProfile {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Password = &password

	return r
}

// Username The username that identifies the user.
// If you specify the `password` grant type, this parameter is required.
// It is not valid with other grant types.
// API name: username
func (r *ActivateUserProfile) Username(username string) *ActivateUserProfile {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.Username = &username

	return r
}
