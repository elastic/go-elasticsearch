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
// Code generated from specification version 8.19.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newSecurityQueryAPIKeysFunc(t Transport) SecurityQueryAPIKeys {
	return func(o ...func(*SecurityQueryAPIKeysRequest)) (*Response, error) {
		var r = SecurityQueryAPIKeysRequest{}
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

// SecurityQueryAPIKeys - Retrieves information for API keys using a subset of query DSL
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/security-api-query-api-key.html.
type SecurityQueryAPIKeys func(o ...func(*SecurityQueryAPIKeysRequest)) (*Response, error)

// SecurityQueryAPIKeysRequest configures the Security QueryAPI Keys API request.
type SecurityQueryAPIKeysRequest struct {
	Body io.Reader

	TypedKeys      *bool
	WithLimitedBy  *bool
	WithProfileUID *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r SecurityQueryAPIKeysRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "security.query_api_keys")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_security/_query/api_key"))
	path.WriteString("http://")
	path.WriteString("/_security/_query/api_key")

	params = make(map[string]string)

	if r.TypedKeys != nil {
		params["typed_keys"] = strconv.FormatBool(*r.TypedKeys)
	}

	if r.WithLimitedBy != nil {
		params["with_limited_by"] = strconv.FormatBool(*r.WithLimitedBy)
	}

	if r.WithProfileUID != nil {
		params["with_profile_uid"] = strconv.FormatBool(*r.WithProfileUID)
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
		instrument.BeforeRequest(req, "security.query_api_keys")
		if reader := instrument.RecordRequestBody(ctx, "security.query_api_keys", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "security.query_api_keys")
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
func (f SecurityQueryAPIKeys) WithContext(v context.Context) func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.ctx = v
	}
}

// WithBody - From, size, query, sort and search_after.
func (f SecurityQueryAPIKeys) WithBody(v io.Reader) func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.Body = v
	}
}

// WithTypedKeys - flag to prefix aggregation names by their respective types in the response.
func (f SecurityQueryAPIKeys) WithTypedKeys(v bool) func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.TypedKeys = &v
	}
}

// WithWithLimitedBy - flag to show the limited-by role descriptors of api keys.
func (f SecurityQueryAPIKeys) WithWithLimitedBy(v bool) func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.WithLimitedBy = &v
	}
}

// WithWithProfileUID - flag to also retrieve the api key's owner profile uid, if it exists.
func (f SecurityQueryAPIKeys) WithWithProfileUID(v bool) func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.WithProfileUID = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SecurityQueryAPIKeys) WithPretty() func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SecurityQueryAPIKeys) WithHuman() func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SecurityQueryAPIKeys) WithErrorTrace() func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SecurityQueryAPIKeys) WithFilterPath(v ...string) func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SecurityQueryAPIKeys) WithHeader(h map[string]string) func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SecurityQueryAPIKeys) WithOpaqueID(s string) func(*SecurityQueryAPIKeysRequest) {
	return func(r *SecurityQueryAPIKeysRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
