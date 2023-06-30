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
// Code generated from specification version 8.8.2: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newCatComponentTemplatesFunc(t Transport) CatComponentTemplates {
	return func(o ...func(*CatComponentTemplatesRequest)) (*Response, error) {
		var r = CatComponentTemplatesRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// CatComponentTemplates returns information about existing component_templates templates.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/cat-component-templates.html.
type CatComponentTemplates func(o ...func(*CatComponentTemplatesRequest)) (*Response, error)

// CatComponentTemplatesRequest configures the Cat Component Templates API request.
type CatComponentTemplatesRequest struct {
	Name string

	Format        string
	H             []string
	Help          *bool
	Local         *bool
	MasterTimeout time.Duration
	S             []string
	V             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r CatComponentTemplatesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(7 + 1 + len("_cat") + 1 + len("component_templates") + 1 + len(r.Name))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_cat")
	path.WriteString("/")
	path.WriteString("component_templates")
	if r.Name != "" {
		path.WriteString("/")
		path.WriteString(r.Name)
	}

	params = make(map[string]string)

	if r.Format != "" {
		params["format"] = r.Format
	}

	if len(r.H) > 0 {
		params["h"] = strings.Join(r.H, ",")
	}

	if r.Help != nil {
		params["help"] = strconv.FormatBool(*r.Help)
	}

	if r.Local != nil {
		params["local"] = strconv.FormatBool(*r.Local)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if len(r.S) > 0 {
		params["s"] = strings.Join(r.S, ",")
	}

	if r.V != nil {
		params["v"] = strconv.FormatBool(*r.V)
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
func (f CatComponentTemplates) WithContext(v context.Context) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.ctx = v
	}
}

// WithName - a pattern that returned component template names must match.
func (f CatComponentTemplates) WithName(v string) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.Name = v
	}
}

// WithFormat - a short version of the accept header, e.g. json, yaml.
func (f CatComponentTemplates) WithFormat(v string) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.Format = v
	}
}

// WithH - comma-separated list of column names to display.
func (f CatComponentTemplates) WithH(v ...string) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.H = v
	}
}

// WithHelp - return help information.
func (f CatComponentTemplates) WithHelp(v bool) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.Help = &v
	}
}

// WithLocal - return local information, do not retrieve the state from master node (default: false).
func (f CatComponentTemplates) WithLocal(v bool) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.Local = &v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func (f CatComponentTemplates) WithMasterTimeout(v time.Duration) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.MasterTimeout = v
	}
}

// WithS - comma-separated list of column names or column aliases to sort by.
func (f CatComponentTemplates) WithS(v ...string) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.S = v
	}
}

// WithV - verbose mode. display column headers.
func (f CatComponentTemplates) WithV(v bool) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.V = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f CatComponentTemplates) WithPretty() func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f CatComponentTemplates) WithHuman() func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f CatComponentTemplates) WithErrorTrace() func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f CatComponentTemplates) WithFilterPath(v ...string) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f CatComponentTemplates) WithHeader(h map[string]string) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f CatComponentTemplates) WithOpaqueID(s string) func(*CatComponentTemplatesRequest) {
	return func(r *CatComponentTemplatesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
