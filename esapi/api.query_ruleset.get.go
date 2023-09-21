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

func newQueryRulesetGetFunc(t Transport) QueryRulesetGet {
	return func(ruleset_id string, o ...func(*QueryRulesetGetRequest)) (*Response, error) {
		var r = QueryRulesetGetRequest{RulesetID: ruleset_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// QueryRulesetGet returns the details about a query ruleset.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/get-query-ruleset.html.
type QueryRulesetGet func(ruleset_id string, o ...func(*QueryRulesetGetRequest)) (*Response, error)

// QueryRulesetGetRequest configures the Query Ruleset Get API request.
type QueryRulesetGetRequest struct {
	RulesetID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r QueryRulesetGetRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "GET"

	path.Grow(7 + 1 + len("_query_rules") + 1 + len(r.RulesetID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_query_rules")
	path.WriteString("/")
	path.WriteString(r.RulesetID)

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
func (f QueryRulesetGet) WithContext(v context.Context) func(*QueryRulesetGetRequest) {
	return func(r *QueryRulesetGetRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f QueryRulesetGet) WithPretty() func(*QueryRulesetGetRequest) {
	return func(r *QueryRulesetGetRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f QueryRulesetGet) WithHuman() func(*QueryRulesetGetRequest) {
	return func(r *QueryRulesetGetRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f QueryRulesetGet) WithErrorTrace() func(*QueryRulesetGetRequest) {
	return func(r *QueryRulesetGetRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f QueryRulesetGet) WithFilterPath(v ...string) func(*QueryRulesetGetRequest) {
	return func(r *QueryRulesetGetRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f QueryRulesetGet) WithHeader(h map[string]string) func(*QueryRulesetGetRequest) {
	return func(r *QueryRulesetGetRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f QueryRulesetGet) WithOpaqueID(s string) func(*QueryRulesetGetRequest) {
	return func(r *QueryRulesetGetRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
