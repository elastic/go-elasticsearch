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
)

func newSynonymsGetSynonymsSetsFunc(t Transport) SynonymsGetSynonymsSets {
	return func(o ...func(*SynonymsGetSynonymsSetsRequest)) (*Response, error) {
		var r = SynonymsGetSynonymsSetsRequest{}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// SynonymsGetSynonymsSets retrieves a summary of all defined synonym sets
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/list-synonyms-sets.html.
type SynonymsGetSynonymsSets func(o ...func(*SynonymsGetSynonymsSetsRequest)) (*Response, error)

// SynonymsGetSynonymsSetsRequest configures the Synonyms Get Synonyms Sets API request.
type SynonymsGetSynonymsSetsRequest struct {
	From *int
	Size *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r SynonymsGetSynonymsSetsRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(7 + len("/_synonyms"))
	path.WriteString("http://")
	path.WriteString("/_synonyms")

	params = make(map[string]string)

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
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
func (f SynonymsGetSynonymsSets) WithContext(v context.Context) func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		r.ctx = v
	}
}

// WithFrom - starting offset.
func (f SynonymsGetSynonymsSets) WithFrom(v int) func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		r.From = &v
	}
}

// WithSize - specifies a max number of results to get.
func (f SynonymsGetSynonymsSets) WithSize(v int) func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		r.Size = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f SynonymsGetSynonymsSets) WithPretty() func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f SynonymsGetSynonymsSets) WithHuman() func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f SynonymsGetSynonymsSets) WithErrorTrace() func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f SynonymsGetSynonymsSets) WithFilterPath(v ...string) func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f SynonymsGetSynonymsSets) WithHeader(h map[string]string) func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f SynonymsGetSynonymsSets) WithOpaqueID(s string) func(*SynonymsGetSynonymsSetsRequest) {
	return func(r *SynonymsGetSynonymsSetsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
