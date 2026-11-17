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
	"io"
	"net/http"
	"strings"
)

func newSecurityPutUserManagedServiceAccountFunc(t Transport) SecurityPutUserManagedServiceAccount {
	return func(body io.Reader, namespace string, service string, o ...func(*SecurityPutUserManagedServiceAccountRequest)) (*Response, error) {
		var r = SecurityPutUserManagedServiceAccountRequest{Body: body, Namespace: namespace, Service: service}
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

// SecurityPutUserManagedServiceAccount - Create or update a user-managed service account
//
// This API is experimental.
type SecurityPutUserManagedServiceAccount func(body io.Reader, namespace string, service string, o ...func(*SecurityPutUserManagedServiceAccountRequest)) (*Response, error)

// SecurityPutUserManagedServiceAccountRequest configures the Security Put User Managed Service Account API request.
type SecurityPutUserManagedServiceAccountRequest struct {
	Body io.Reader

	Namespace string
	Service   string

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
func (r SecurityPutUserManagedServiceAccountRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.put_user_managed_service_account")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

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

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "security.put_user_managed_service_account")
		if reader := instrument.RecordRequestBody(ctx, "security.put_user_managed_service_account", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.put_user_managed_service_account")
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
func (f SecurityPutUserManagedServiceAccount) WithContext(v context.Context) func(*SecurityPutUserManagedServiceAccountRequest) {
	return func(r *SecurityPutUserManagedServiceAccountRequest) {
		r.ctx = v
	}
}

// WithRefresh - if `true` then refresh the affected shards to make this operation visible to search, if `wait_for` (the default) then wait for a refresh to make this operation visible to search, if `false` then do nothing with refreshes..
func (f SecurityPutUserManagedServiceAccount) WithRefresh(v string) func(*SecurityPutUserManagedServiceAccountRequest) {
	return func(r *SecurityPutUserManagedServiceAccountRequest) {
		r.Refresh = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityPutUserManagedServiceAccount) WithPretty() func(*SecurityPutUserManagedServiceAccountRequest) {
	return func(r *SecurityPutUserManagedServiceAccountRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityPutUserManagedServiceAccount) WithHuman() func(*SecurityPutUserManagedServiceAccountRequest) {
	return func(r *SecurityPutUserManagedServiceAccountRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityPutUserManagedServiceAccount) WithErrorTrace() func(*SecurityPutUserManagedServiceAccountRequest) {
	return func(r *SecurityPutUserManagedServiceAccountRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityPutUserManagedServiceAccount) WithFilterPath(v ...string) func(*SecurityPutUserManagedServiceAccountRequest) {
	return func(r *SecurityPutUserManagedServiceAccountRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityPutUserManagedServiceAccount) WithHeader(h map[string]string) func(*SecurityPutUserManagedServiceAccountRequest) {
	return func(r *SecurityPutUserManagedServiceAccountRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityPutUserManagedServiceAccount) WithOpaqueID(s string) func(*SecurityPutUserManagedServiceAccountRequest) {
	return func(r *SecurityPutUserManagedServiceAccountRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
