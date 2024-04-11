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
// Code generated from specification version 8.13.2: DO NOT EDIT

package esapi

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func newIndicesResolveClusterFunc(t Transport) IndicesResolveCluster {
	return func(name []string, o ...func(*IndicesResolveClusterRequest)) (*Response, error) {
		var r = IndicesResolveClusterRequest{Name: name}
		for _, f := range o {
			f(&r)
		}

		if transport, ok := t.(Instrumented); ok {
			r.instrument = transport.InstrumentationEnabled()
		}

		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// IndicesResolveCluster resolves the specified index expressions to return information about each cluster, including the local cluster, if included.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-resolve-cluster-api.html.
type IndicesResolveCluster func(name []string, o ...func(*IndicesResolveClusterRequest)) (*Response, error)

// IndicesResolveClusterRequest configures the Indices Resolve Cluster API request.
type IndicesResolveClusterRequest struct {
	Name []string

	AllowNoIndices    *bool
	ExpandWildcards   string
	IgnoreThrottled   *bool
	IgnoreUnavailable *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IndicesResolveClusterRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.resolve_cluster")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	if len(r.Name) == 0 {
		return nil, errors.New("name is required and cannot be nil or empty")
	}

	path.Grow(7 + 1 + len("_resolve") + 1 + len("cluster") + 1 + len(strings.Join(r.Name, ",")))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_resolve")
	path.WriteString("/")
	path.WriteString("cluster")
	path.WriteString("/")
	path.WriteString(strings.Join(r.Name, ","))
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "name", strings.Join(r.Name, ","))
	}

	params = make(map[string]string)

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.IgnoreThrottled != nil {
		params["ignore_throttled"] = strconv.FormatBool(*r.IgnoreThrottled)
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
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
		if instrument, ok := r.instrument.(Instrumentation); ok {
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

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "indices.resolve_cluster")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.resolve_cluster")
	}
	if err != nil {
		if instrument, ok := r.instrument.(Instrumentation); ok {
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
func (f IndicesResolveCluster) WithContext(v context.Context) func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.ctx = v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
func (f IndicesResolveCluster) WithAllowNoIndices(v bool) func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.AllowNoIndices = &v
	}
}

// WithExpandWildcards - whether wildcard expressions should get expanded to open or closed indices (default: open).
func (f IndicesResolveCluster) WithExpandWildcards(v string) func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.ExpandWildcards = v
	}
}

// WithIgnoreThrottled - whether specified concrete, expanded or aliased indices should be ignored when throttled.
func (f IndicesResolveCluster) WithIgnoreThrottled(v bool) func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.IgnoreThrottled = &v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func (f IndicesResolveCluster) WithIgnoreUnavailable(v bool) func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IndicesResolveCluster) WithPretty() func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IndicesResolveCluster) WithHuman() func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IndicesResolveCluster) WithErrorTrace() func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IndicesResolveCluster) WithFilterPath(v ...string) func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IndicesResolveCluster) WithHeader(h map[string]string) func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IndicesResolveCluster) WithOpaqueID(s string) func(*IndicesResolveClusterRequest) {
	return func(r *IndicesResolveClusterRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
