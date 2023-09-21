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
	"net/http"
	"strings"
)

func newSearchApplicationGetBehavioralAnalyticsFunc(t Transport) SearchApplicationGetBehavioralAnalytics {
	return func(o ...func(*SearchApplicationGetBehavioralAnalyticsRequest)) (*Response, error) {
		var r = SearchApplicationGetBehavioralAnalyticsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SearchApplicationGetBehavioralAnalytics returns the existing behavioral analytics collections.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/list-analytics-collection.html.
type SearchApplicationGetBehavioralAnalytics func(o ...func(*SearchApplicationGetBehavioralAnalyticsRequest)) (*Response, error)

// SearchApplicationGetBehavioralAnalyticsRequest configures the Search Application Get Behavioral Analytics API request.
type SearchApplicationGetBehavioralAnalyticsRequest struct {
	Name []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r SearchApplicationGetBehavioralAnalyticsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(7 + 1 + len("_application") + 1 + len("analytics") + 1 + len(strings.Join(r.Name, ",")))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_application")
	path.WriteString("/")
	path.WriteString("analytics")
	if len(r.Name) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Name, ","))
	}

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
func (f SearchApplicationGetBehavioralAnalytics) WithContext(v context.Context) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	return func(r *SearchApplicationGetBehavioralAnalyticsRequest) {
		r.ctx = v
	}
}

// WithName - a list of analytics collections to limit the returned information.
func (f SearchApplicationGetBehavioralAnalytics) WithName(v ...string) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	return func(r *SearchApplicationGetBehavioralAnalyticsRequest) {
		r.Name = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SearchApplicationGetBehavioralAnalytics) WithPretty() func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	return func(r *SearchApplicationGetBehavioralAnalyticsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SearchApplicationGetBehavioralAnalytics) WithHuman() func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	return func(r *SearchApplicationGetBehavioralAnalyticsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SearchApplicationGetBehavioralAnalytics) WithErrorTrace() func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	return func(r *SearchApplicationGetBehavioralAnalyticsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SearchApplicationGetBehavioralAnalytics) WithFilterPath(v ...string) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	return func(r *SearchApplicationGetBehavioralAnalyticsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SearchApplicationGetBehavioralAnalytics) WithHeader(h map[string]string) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	return func(r *SearchApplicationGetBehavioralAnalyticsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SearchApplicationGetBehavioralAnalytics) WithOpaqueID(s string) func(*SearchApplicationGetBehavioralAnalyticsRequest) {
	return func(r *SearchApplicationGetBehavioralAnalyticsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
