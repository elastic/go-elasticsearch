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
	"strconv"
	"strings"
	"time"
)

func newHealthReportFunc(t Transport) HealthReport {
	return func(o ...func(*HealthReportRequest)) (*Response, error) {
		var r = HealthReportRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// HealthReport returns the health of the cluster.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/health-api.html.
type HealthReport func(o ...func(*HealthReportRequest)) (*Response, error)

// HealthReportRequest configures the Health Report API request.
type HealthReportRequest struct {
	Feature string

	Size    *int
	Timeout time.Duration
	Verbose *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r HealthReportRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(7 + 1 + len("_health_report") + 1 + len(r.Feature))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_health_report")
	if r.Feature != "" {
		path.WriteString("/")
		path.WriteString(r.Feature)
	}

	params = make(map[string]string)

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.Verbose != nil {
		params["verbose"] = strconv.FormatBool(*r.Verbose)
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
func (f HealthReport) WithContext(v context.Context) func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.ctx = v
	}
}

// WithFeature - a feature of the cluster, as returned by the top-level health api.
func (f HealthReport) WithFeature(v string) func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.Feature = v
	}
}

// WithSize - limit the number of affected resources the health api returns.
func (f HealthReport) WithSize(v int) func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.Size = &v
	}
}

// WithTimeout - explicit operation timeout.
func (f HealthReport) WithTimeout(v time.Duration) func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.Timeout = v
	}
}

// WithVerbose - opt in for more information about the health of the system.
func (f HealthReport) WithVerbose(v bool) func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.Verbose = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f HealthReport) WithPretty() func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f HealthReport) WithHuman() func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f HealthReport) WithErrorTrace() func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f HealthReport) WithFilterPath(v ...string) func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f HealthReport) WithHeader(h map[string]string) func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f HealthReport) WithOpaqueID(s string) func(*HealthReportRequest) {
	return func(r *HealthReportRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
