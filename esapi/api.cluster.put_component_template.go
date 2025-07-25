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
	"time"
)

func newClusterPutComponentTemplateFunc(t Transport) ClusterPutComponentTemplate {
	return func(name string, body io.Reader, o ...func(*ClusterPutComponentTemplateRequest)) (*Response, error) {
		var r = ClusterPutComponentTemplateRequest{Name: name, Body: body}
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

// ClusterPutComponentTemplate creates or updates a component template
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-component-template.html.
type ClusterPutComponentTemplate func(name string, body io.Reader, o ...func(*ClusterPutComponentTemplateRequest)) (*Response, error)

// ClusterPutComponentTemplateRequest configures the Cluster Put Component Template API request.
type ClusterPutComponentTemplateRequest struct {
	Body io.Reader

	Name string

	Cause         string
	Create        *bool
	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ClusterPutComponentTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.put_component_template")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_component_template") + 1 + len(r.Name))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_component_template")
	path.WriteString("/")
	path.WriteString(r.Name)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "name", r.Name)
	}

	params = make(map[string]string)

	if r.Cause != "" {
		params["cause"] = r.Cause
	}

	if r.Create != nil {
		params["create"] = strconv.FormatBool(*r.Create)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
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
		instrument.BeforeRequest(req, "cluster.put_component_template")
		if reader := instrument.RecordRequestBody(ctx, "cluster.put_component_template", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.put_component_template")
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
func (f ClusterPutComponentTemplate) WithContext(v context.Context) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.ctx = v
	}
}

// WithCause - user defined reason for create the component template.
func (f ClusterPutComponentTemplate) WithCause(v string) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.Cause = v
	}
}

// WithCreate - whether the index template should only be added if new or can also replace an existing one.
func (f ClusterPutComponentTemplate) WithCreate(v bool) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.Create = &v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
func (f ClusterPutComponentTemplate) WithMasterTimeout(v time.Duration) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ClusterPutComponentTemplate) WithPretty() func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ClusterPutComponentTemplate) WithHuman() func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ClusterPutComponentTemplate) WithErrorTrace() func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ClusterPutComponentTemplate) WithFilterPath(v ...string) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ClusterPutComponentTemplate) WithHeader(h map[string]string) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ClusterPutComponentTemplate) WithOpaqueID(s string) func(*ClusterPutComponentTemplateRequest) {
	return func(r *ClusterPutComponentTemplateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
