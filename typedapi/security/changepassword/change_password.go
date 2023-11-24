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

// Changes the passwords of users in the native realm and built-in users.
package changepassword

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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/refresh"
)

const (
	usernameMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type ChangePassword struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	username string
}

// NewChangePassword type alias for index.
type NewChangePassword func() *ChangePassword

// NewChangePasswordFunc returns a new instance of ChangePassword with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewChangePasswordFunc(tp elastictransport.Interface) NewChangePassword {
	return func() *ChangePassword {
		n := New(tp)

		return n
	}
}

// Changes the passwords of users in the native realm and built-in users.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-change-password.html
func New(tp elastictransport.Interface) *ChangePassword {
	r := &ChangePassword{
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
func (r *ChangePassword) Raw(raw io.Reader) *ChangePassword {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *ChangePassword) Request(req *Request) *ChangePassword {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *ChangePassword) HttpRequest(ctx context.Context) (*http.Request, error) {
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
			return nil, fmt.Errorf("could not serialise request for ChangePassword: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == usernameMask:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("user")
		path.WriteString("/")

		path.WriteString(r.username)
		path.WriteString("/")
		path.WriteString("_password")

		method = http.MethodPut
	case r.paramSet == 0:
		path.WriteString("/")
		path.WriteString("_security")
		path.WriteString("/")
		path.WriteString("user")
		path.WriteString("/")
		path.WriteString("_password")

		method = http.MethodPut
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
func (r ChangePassword) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the ChangePassword query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a changepassword.Response
func (r ChangePassword) Do(ctx context.Context) (*Response, error) {

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

// Header set a key, value pair in the ChangePassword headers map.
func (r *ChangePassword) Header(key, value string) *ChangePassword {
	r.headers.Set(key, value)

	return r
}

// Username The user whose password you want to change. If you do not specify this
// parameter, the password is changed for the current user.
// API Name: username
func (r *ChangePassword) Username(username string) *ChangePassword {
	r.paramSet |= usernameMask
	r.username = username

	return r
}

// Refresh If `true` (the default) then refresh the affected shards to make this
// operation visible to search, if `wait_for` then wait for a refresh to make
// this operation visible to search, if `false` then do nothing with refreshes.
// API name: refresh
func (r *ChangePassword) Refresh(refresh refresh.Refresh) *ChangePassword {
	r.values.Set("refresh", refresh.String())

	return r
}

// Password The new password value. Passwords must be at least 6 characters long.
// API name: password
func (r *ChangePassword) Password(password string) *ChangePassword {
	r.req.Password = &password

	return r
}

// PasswordHash A hash of the new password value. This must be produced using the same
// hashing algorithm as has been configured for password storage. For more
// details,
// see the explanation of the `xpack.security.authc.password_hashing.algorithm`
// setting.
// API name: password_hash
func (r *ChangePassword) PasswordHash(passwordhash string) *ChangePassword {

	r.req.PasswordHash = &passwordhash

	return r
}
