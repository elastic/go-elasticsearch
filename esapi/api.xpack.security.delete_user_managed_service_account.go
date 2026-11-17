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
// Code generated from specification version 9.6.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
)

func newSecurityDeleteUserManagedServiceAccountFunc(t Transport) SecurityDeleteUserManagedServiceAccount {
	return func(namespace string, service string, o ...func(*SecurityDeleteUserManagedServiceAccountRequest)) (*Response, error) {
		var r = SecurityDeleteUserManagedServiceAccountRequest{Namespace: namespace, Service: service}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.Instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SecurityDeleteUserManagedServiceAccount - Delete a user-managed service account
//
// This API is experimental.
type SecurityDeleteUserManagedServiceAccount func(namespace string, service string, o ...func(*SecurityDeleteUserManagedServiceAccountRequest)) (*Response, error)

// SecurityDeleteUserManagedServiceAccountRequest configures the Security Delete User Managed Service Account API request.
type SecurityDeleteUserManagedServiceAccountRequest struct {
	Namespace string
	Service   string

	Force   *bool
	Refresh string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SecurityDeleteUserManagedServiceAccountRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.delete_user_managed_service_account")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "DELETE"

	path.Grow(7 + 1 + len("_security") + 1 + len("service") + 1 + len(r.Namespace) + 1 + len(r.Service))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_security")
	path.WriteString("/")
	path.WriteString("service")
	path.WriteString("/")
	path.WriteString(r.Namespace)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "namespace", r.Namespace)
	}
	path.WriteString("/")
	path.WriteString(r.Service)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "service", r.Service)
	}

	params = make(map[string]string)

	if r.Force != nil {
		params["force"] = strconv.FormatBool(*r.Force)
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

	req, err := newRequest(method, path.String(), nil)
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
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

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "security.delete_user_managed_service_account")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.delete_user_managed_service_account")
	}
	if err != nil {
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordError(ctx, err)
		}
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
func (f SecurityDeleteUserManagedServiceAccount) WithContext(v context.Context) func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		r.ctx = v
	}
}

// WithForce - if `false` (the default), deleting an account that still has service tokens is rejected. if `true`, the account is deleted and its tokens are left in place; they cannot authenticate while no account of that name exists, but creating one again later revives them..
func (f SecurityDeleteUserManagedServiceAccount) WithForce(v bool) func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		r.Force = &v
	}
}

// WithRefresh - if `true` then refresh the affected shards to make this operation visible to search, if `wait_for` (the default) then wait for a refresh to make this operation visible to search, if `false` then do nothing with refreshes..
func (f SecurityDeleteUserManagedServiceAccount) WithRefresh(v string) func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		r.Refresh = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityDeleteUserManagedServiceAccount) WithPretty() func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityDeleteUserManagedServiceAccount) WithHuman() func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityDeleteUserManagedServiceAccount) WithErrorTrace() func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityDeleteUserManagedServiceAccount) WithFilterPath(v ...string) func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityDeleteUserManagedServiceAccount) WithHeader(h map[string]string) func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityDeleteUserManagedServiceAccount) WithOpaqueID(s string) func(*SecurityDeleteUserManagedServiceAccountRequest) {
	return func(r *SecurityDeleteUserManagedServiceAccountRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
