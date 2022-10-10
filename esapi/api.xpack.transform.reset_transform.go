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
// Code generated from specification version 8.6.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newTransformResetTransformFunc(t Transport) TransformResetTransform {
	return func(transform_id string, o ...func(*TransformResetTransformRequest)) (*Response, error) {
		var r = TransformResetTransformRequest{TransformID: transform_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// TransformResetTransform - Resets an existing transform.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/reset-transform.html.
type TransformResetTransform func(transform_id string, o ...func(*TransformResetTransformRequest)) (*Response, error)

// TransformResetTransformRequest configures the Transform Reset Transform API request.
type TransformResetTransformRequest struct {
	TransformID string

	Force   *bool
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r TransformResetTransformRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(7 + 1 + len("_transform") + 1 + len(r.TransformID) + 1 + len("_reset"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_transform")
	path.WriteString("/")
	path.WriteString(r.TransformID)
	path.WriteString("/")
	path.WriteString("_reset")

	params = make(map[string]string)

	if r.Force != nil {
		params["force"] = strconv.FormatBool(*r.Force)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
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
func (f TransformResetTransform) WithContext(v context.Context) func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		r.ctx = v
	}
}

// WithForce - when `true`, the transform is reset regardless of its current state. the default value is `false`, meaning that the transform must be `stopped` before it can be reset..
func (f TransformResetTransform) WithForce(v bool) func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		r.Force = &v
	}
}

// WithTimeout - controls the time to wait for the transform to reset.
func (f TransformResetTransform) WithTimeout(v time.Duration) func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f TransformResetTransform) WithPretty() func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f TransformResetTransform) WithHuman() func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f TransformResetTransform) WithErrorTrace() func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f TransformResetTransform) WithFilterPath(v ...string) func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f TransformResetTransform) WithHeader(h map[string]string) func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f TransformResetTransform) WithOpaqueID(s string) func(*TransformResetTransformRequest) {
	return func(r *TransformResetTransformRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
