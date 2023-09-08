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
	"errors"
	"net/http"
	"strings"
)

func newClusterInfoFunc(t Transport) ClusterInfo {
	return func(target []string, o ...func(*ClusterInfoRequest)) (*Response, error) {
		var r = ClusterInfoRequest{Target: target}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// ClusterInfo returns different information about the cluster.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/cluster-info.html.
type ClusterInfo func(target []string, o ...func(*ClusterInfoRequest)) (*Response, error)

// ClusterInfoRequest configures the Cluster Info API request.
type ClusterInfoRequest struct {
	Target []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r ClusterInfoRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	if len(r.Target) == 0 {
		return nil, errors.New("target is required and cannot be nil or empty")
	}

	path.Grow(7 + 1 + len("_info") + 1 + len(strings.Join(r.Target, ",")))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_info")
	path.WriteString("/")
	path.WriteString(strings.Join(r.Target, ","))

	params = make(map[string]string)

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
func (f ClusterInfo) WithContext(v context.Context) func(*ClusterInfoRequest) {
	return func(r *ClusterInfoRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ClusterInfo) WithPretty() func(*ClusterInfoRequest) {
	return func(r *ClusterInfoRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ClusterInfo) WithHuman() func(*ClusterInfoRequest) {
	return func(r *ClusterInfoRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ClusterInfo) WithErrorTrace() func(*ClusterInfoRequest) {
	return func(r *ClusterInfoRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ClusterInfo) WithFilterPath(v ...string) func(*ClusterInfoRequest) {
	return func(r *ClusterInfoRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ClusterInfo) WithHeader(h map[string]string) func(*ClusterInfoRequest) {
	return func(r *ClusterInfoRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ClusterInfo) WithOpaqueID(s string) func(*ClusterInfoRequest) {
	return func(r *ClusterInfoRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
