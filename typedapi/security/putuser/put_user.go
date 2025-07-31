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

// Create or update users.
//
// Add and update users in the native realm.
// A password is required for adding a new user but is optional when updating an
// existing user.
// To change a user's password without updating any other fields, use the change
// password API.
package putuser

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

const (
	usernameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type PutUser struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	raw io.Reader

	req      *Request
	deferred []func(request *Request) error
	buf      *gobytes.Buffer

	paramSet int

	username string

	spanStarted bool

	instrument elastictransport.Instrumentation
}

// NewPutUser type alias for index.
type NewPutUser func(username string) *PutUser

// NewPutUserFunc returns a new instance of PutUser with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewPutUserFunc(tp elastictransport.Interface) NewPutUser {
	return func(username string) *PutUser {
		n := New(tp)

		n._username(username)

		return n
	}
}

// Create or update users.
//
// Add and update users in the native realm.
// A password is required for adding a new user but is optional when updating an
// existing user.
// To change a user's password without updating any other fields, use the change
// password API.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-put-user.html
func New(tp elastictransport.Interface) *PutUser {
	r := &PutUser{
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
func (r *PutUser) Raw(raw io.Reader) *PutUser {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *PutUser) Request(req *Request) *PutUser {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *PutUser) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for PutUser: %w", err)
		}

		r.buf.Write(data)

	}

	if r.buf.Len() > 0 {
		r.raw = r.buf
	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == usernameMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("user")
		path.WriteString("/")

		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordPathPart(ctx, "username", r.username)
		}
		path.WriteString(r.username)

		method = http.MethodPut
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
func (r PutUser) Perform(providedCtx context.Context) (*http.Response, error) {
	var ctx context.Context
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		if r.spanStarted == false {
			ctx := instrument.Start(providedCtx, "security.put_user")
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
		instrument.BeforeRequest(req, "security.put_user")
		if reader := instrument.RecordRequestBody(ctx, "security.put_user", r.raw); reader != nil {
			req.Body = reader
		}
	}
	res, err := r.transport.Perform(req)
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.put_user")
	}
	if err != nil {
		localErr := fmt.Errorf("an error happened during the PutUser query execution: %w", err)
		if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
			instrument.RecordError(ctx, localErr)
		}
		return nil, localErr
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a putuser.Response
func (r PutUser) Do(providedCtx context.Context) (*Response, error) {
	var ctx context.Context
	r.spanStarted = true
	if instrument, ok := r.instrument.(elastictransport.Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.put_user")
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

// Header set a key, value pair in the PutUser headers map.
func (r *PutUser) Header(key, value string) *PutUser {
	r.headers.Set(key, value)

	return r
}

// Username An identifier for the user.
//
// NOTE: Usernames must be at least 1 and no more than 507 characters.
// They can contain alphanumeric characters (a-z, A-Z, 0-9), spaces,
// punctuation, and printable symbols in the Basic Latin (ASCII) block.
// Leading or trailing whitespace is not allowed.
// API Name: username
func (r *PutUser) _username(username string) *PutUser {
	r.paramSet |= usernameMask
	r.username = username

	return r
}

// Refresh Valid values are `true`, `false`, and `wait_for`.
// These values have the same meaning as in the index API, but the default value
// for this API is true.
// API name: refresh
func (r *PutUser) Refresh(refresh refresh.Refresh) *PutUser {
	r.values.Set("refresh", refresh.String())

	return r
}

// ErrorTrace When set to `true` Elasticsearch will include the full stack trace of errors
// when they occur.
// API name: error_trace
func (r *PutUser) ErrorTrace(errortrace bool) *PutUser {
	r.values.Set("error_trace", strconv.FormatBool(errortrace))

	return r
}

// FilterPath Comma-separated list of filters in dot notation which reduce the response
// returned by Elasticsearch.
// API name: filter_path
func (r *PutUser) FilterPath(filterpaths ...string) *PutUser {
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
func (r *PutUser) Human(human bool) *PutUser {
	r.values.Set("human", strconv.FormatBool(human))

	return r
}

// Pretty If set to `true` the returned JSON will be "pretty-formatted". Only use
// this option for debugging only.
// API name: pretty
func (r *PutUser) Pretty(pretty bool) *PutUser {
	r.values.Set("pretty", strconv.FormatBool(pretty))

	return r
}

// Email The email of the user.
// API name: email
func (r *PutUser) Email(email string) *PutUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Email = &email

	return r
}

// Enabled Specifies whether the user is enabled.
// API name: enabled
func (r *PutUser) Enabled(enabled bool) *PutUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Enabled = &enabled

	return r
}

// FullName The full name of the user.
// API name: full_name
func (r *PutUser) FullName(fullname string) *PutUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.FullName = &fullname

	return r
}

// Metadata Arbitrary metadata that you want to associate with the user.
// API name: metadata
func (r *PutUser) Metadata(metadata types.Metadata) *PutUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Metadata = metadata

	return r
}

// Password The user's password.
// Passwords must be at least 6 characters long.
// When adding a user, one of `password` or `password_hash` is required.
// When updating an existing user, the password is optional, so that other
// fields on the user (such as their roles) may be updated without modifying the
// user's password
// API name: password
func (r *PutUser) Password(password string) *PutUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Password = &password

	return r
}

// PasswordHash A hash of the user's password.
// This must be produced using the same hashing algorithm as has been configured
// for password storage.
// For more details, see the explanation of the
// `xpack.security.authc.password_hashing.algorithm` setting in the user cache
// and password hash algorithm documentation.
// Using this parameter allows the client to pre-hash the password for
// performance and/or confidentiality reasons.
// The `password` parameter and the `password_hash` parameter cannot be used in
// the same request.
// API name: password_hash
func (r *PutUser) PasswordHash(passwordhash string) *PutUser {
	if r.req == nil {
		r.req = NewRequest()
	}

	r.req.PasswordHash = &passwordhash

	return r
}

// Roles A set of roles the user has.
// The roles determine the user's access permissions.
// To create a user without any roles, specify an empty list (`[]`).
// API name: roles
func (r *PutUser) Roles(roles ...string) *PutUser {
	if r.req == nil {
		r.req = NewRequest()
	}
	r.req.Roles = roles

	return r
}
