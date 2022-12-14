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
//
// Code generated from specification version 8.7.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newSecurityDisableUserProfileFunc(t Transport) SecurityDisableUserProfile {
	return func(uid string, o ...func(*SecurityDisableUserProfileRequest)) (*Response, error) {
		var r = SecurityDisableUserProfileRequest{UID: uid}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SecurityDisableUserProfile - Disables a user profile so it's not visible in user profile searches.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/security-api-disable-user-profile.html.
type SecurityDisableUserProfile func(uid string, o ...func(*SecurityDisableUserProfileRequest)) (*Response, error)

// SecurityDisableUserProfileRequest configures the Security Disable User Profile API request.
type SecurityDisableUserProfileRequest struct {
	UID string

	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r SecurityDisableUserProfileRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(7 + 1 + len("_security") + 1 + len("profile") + 1 + len(r.UID) + 1 + len("_disable"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("profile")
	path.WriteString("/")
	path.WriteString(r.UID)
	path.WriteString("/")
	path.WriteString("_disable")

	params = make(map[string]string)

	if r.Refresh != "" {
		params["refresh"] = r.Refresh
	}

	if r.Pretty {
		params["pretty"] = "true"
	}

	if r.Human {
		params["human"] = "true"
	}

	if r.ErrorTrace {
		params["error_trace"] = "true"
	}

	if len(r.FilterPath) > 0 {
		params["filter_path"] = strings.Join(r.FilterPath, ",")
	}

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		return nil, err
	}

	if len(params) > 0 {
		q := req.URL.Query()
		for k, v := range params {
			q.Set(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	if len(r.Header) > 0 {
		if len(req.Header) == 0 {
			req.Header = r.Header
		} else {
			for k, vv := range r.Header {
				for _, v := range vv {
					req.Header.Add(k, v)
				}
			}
		}
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	res, err := transport.Perform(req)
	if err != nil {
		return nil, err
	}

	response := Response{
		StatusCode: res.StatusCode,
		Body:       res.Body,
		Header:     res.Header,
	}

	return &response, nil
}

// WithContext sets the request context.
func (f SecurityDisableUserProfile) WithContext(v context.Context) func(*SecurityDisableUserProfileRequest) {
	return func(r *SecurityDisableUserProfileRequest) {
		r.ctx = v
	}
}

// WithRefresh - if `true` then refresh the affected shards to make this operation visible to search, if `wait_for` (the default) then wait for a refresh to make this operation visible to search, if `false` then do nothing with refreshes..
func (f SecurityDisableUserProfile) WithRefresh(v string) func(*SecurityDisableUserProfileRequest) {
	return func(r *SecurityDisableUserProfileRequest) {
		r.Refresh = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityDisableUserProfile) WithPretty() func(*SecurityDisableUserProfileRequest) {
	return func(r *SecurityDisableUserProfileRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityDisableUserProfile) WithHuman() func(*SecurityDisableUserProfileRequest) {
	return func(r *SecurityDisableUserProfileRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityDisableUserProfile) WithErrorTrace() func(*SecurityDisableUserProfileRequest) {
	return func(r *SecurityDisableUserProfileRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityDisableUserProfile) WithFilterPath(v ...string) func(*SecurityDisableUserProfileRequest) {
	return func(r *SecurityDisableUserProfileRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityDisableUserProfile) WithHeader(h map[string]string) func(*SecurityDisableUserProfileRequest) {
	return func(r *SecurityDisableUserProfileRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityDisableUserProfile) WithOpaqueID(s string) func(*SecurityDisableUserProfileRequest) {
	return func(r *SecurityDisableUserProfileRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
