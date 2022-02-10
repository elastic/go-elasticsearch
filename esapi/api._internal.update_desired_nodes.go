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
// Code generated from specification version 8.1.0: DO NOT EDIT

package esapi

import (
	"context"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newInternalUpdateDesiredNodesFunc(t Transport) InternalUpdateDesiredNodes {
	return func(body io.Reader, history_id string, version *int, o ...func(*InternalUpdateDesiredNodesRequest)) (*Response, error) {
		var r = InternalUpdateDesiredNodesRequest{Body: body, HistoryID: history_id, Version: version}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// InternalUpdateDesiredNodes updates the desired nodes. Designed for indirect use by ECE/ESS and ECK. Direct use is not supported.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/update-desired-nodes.html.
//
type InternalUpdateDesiredNodes func(body io.Reader, history_id string, version *int, o ...func(*InternalUpdateDesiredNodesRequest)) (*Response, error)

// InternalUpdateDesiredNodesRequest configures the Internal Update Desired Nodes API request.
//
type InternalUpdateDesiredNodesRequest struct {
	Body io.Reader

	HistoryID string
	Version   *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
//
func (r InternalUpdateDesiredNodesRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "PUT"

	if r.Version == nil {
		return nil, errors.New("version is required and cannot be nil")
	}

	path.Grow(7 + 1 + len("_internal") + 1 + len("desired_nodes") + 1 + len(r.HistoryID) + 1 + len(strconv.Itoa(*r.Version)))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_internal")
	path.WriteString("/")
	path.WriteString("desired_nodes")
	path.WriteString("/")
	path.WriteString(r.HistoryID)
	path.WriteString("/")
	path.WriteString(strconv.Itoa(*r.Version))

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

	req, err := newRequest(method, path.String(), r.Body)
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

	if r.Body != nil {
		req.Header[headerContentType] = headerContentTypeJSON
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
//
func (f InternalUpdateDesiredNodes) WithContext(v context.Context) func(*InternalUpdateDesiredNodesRequest) {
	return func(r *InternalUpdateDesiredNodesRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
//
func (f InternalUpdateDesiredNodes) WithPretty() func(*InternalUpdateDesiredNodesRequest) {
	return func(r *InternalUpdateDesiredNodesRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
//
func (f InternalUpdateDesiredNodes) WithHuman() func(*InternalUpdateDesiredNodesRequest) {
	return func(r *InternalUpdateDesiredNodesRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
//
func (f InternalUpdateDesiredNodes) WithErrorTrace() func(*InternalUpdateDesiredNodesRequest) {
	return func(r *InternalUpdateDesiredNodesRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
//
func (f InternalUpdateDesiredNodes) WithFilterPath(v ...string) func(*InternalUpdateDesiredNodesRequest) {
	return func(r *InternalUpdateDesiredNodesRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
//
func (f InternalUpdateDesiredNodes) WithHeader(h map[string]string) func(*InternalUpdateDesiredNodesRequest) {
	return func(r *InternalUpdateDesiredNodesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
//
func (f InternalUpdateDesiredNodes) WithOpaqueID(s string) func(*InternalUpdateDesiredNodesRequest) {
	return func(r *InternalUpdateDesiredNodesRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
