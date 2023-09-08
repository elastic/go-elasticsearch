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
// Code generated from specification version 8.11.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newSecurityUpdateUserProfileDataFunc(t Transport) SecurityUpdateUserProfileData {
	return func(body io.Reader, uid string, o ...func(*SecurityUpdateUserProfileDataRequest)) (*Response, error) {
		var r = SecurityUpdateUserProfileDataRequest{Body: body, UID: uid}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SecurityUpdateUserProfileData - Update application specific data for the user profile of the given unique ID.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-update-user-profile-data.html.
type SecurityUpdateUserProfileData func(body io.Reader, uid string, o ...func(*SecurityUpdateUserProfileDataRequest)) (*Response, error)

// SecurityUpdateUserProfileDataRequest configures the Security Update User Profile Data API request.
type SecurityUpdateUserProfileDataRequest struct {
	Body io.Reader

	UID string

	IfPrimaryTerm *int
	IfSeqNo       *int
	Refresh       string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r SecurityUpdateUserProfileDataRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	path.Grow(7 + 1 + len("_security") + 1 + len("profile") + 1 + len(r.UID) + 1 + len("_data"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("profile")
	path.WriteString("/")
	path.WriteString(r.UID)
	path.WriteString("/")
	path.WriteString("_data")

	params = make(map[string]string)

	if r.IfPrimaryTerm != nil {
		params["if_primary_term"] = strconv.FormatInt(int64(*r.IfPrimaryTerm), 10)
	}

	if r.IfSeqNo != nil {
		params["if_seq_no"] = strconv.FormatInt(int64(*r.IfSeqNo), 10)
	}

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

	req, err := newRequest(method, path.String(), r.Body)
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

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
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
func (f SecurityUpdateUserProfileData) WithContext(v context.Context) func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		r.ctx = v
	}
}

// WithIfPrimaryTerm - only perform the update operation if the last operation that has changed the document has the specified primary term.
func (f SecurityUpdateUserProfileData) WithIfPrimaryTerm(v int) func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		r.IfPrimaryTerm = &v
	}
}

// WithIfSeqNo - only perform the update operation if the last operation that has changed the document has the specified sequence number.
func (f SecurityUpdateUserProfileData) WithIfSeqNo(v int) func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		r.IfSeqNo = &v
	}
}

// WithRefresh - if `true` (the default) then refresh the affected shards to make this operation visible to search, if `wait_for` then wait for a refresh to make this operation visible to search, if `false` then do nothing with refreshes..
func (f SecurityUpdateUserProfileData) WithRefresh(v string) func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		r.Refresh = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityUpdateUserProfileData) WithPretty() func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityUpdateUserProfileData) WithHuman() func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityUpdateUserProfileData) WithErrorTrace() func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityUpdateUserProfileData) WithFilterPath(v ...string) func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityUpdateUserProfileData) WithHeader(h map[string]string) func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityUpdateUserProfileData) WithOpaqueID(s string) func(*SecurityUpdateUserProfileDataRequest) {
	return func(r *SecurityUpdateUserProfileDataRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
