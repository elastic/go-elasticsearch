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
// Code generated from specification version 8.18.0: DO NOT EDIT

package esapi

import (
	"context"
	"errors"
	"net/http"
	"strconv"
	"strings"
)

func newIndicesGetFieldMappingFunc(t Transport) IndicesGetFieldMapping {
	return func(fields []string, o ...func(*IndicesGetFieldMappingRequest)) (*Response, error) {
		var r = IndicesGetFieldMappingRequest{Fields: fields}
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

// IndicesGetFieldMapping returns mapping for one or more fields.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-get-field-mapping.html.
type IndicesGetFieldMapping func(fields []string, o ...func(*IndicesGetFieldMappingRequest)) (*Response, error)

// IndicesGetFieldMappingRequest configures the Indices Get Field Mapping API request.
type IndicesGetFieldMappingRequest struct {
	Index []string

	Fields []string

	AllowNoIndices    *bool
	ExpandWildcards   string
	IgnoreUnavailable *bool
	IncludeDefaults   *bool
	Local             *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IndicesGetFieldMappingRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.get_field_mapping")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	if len(r.Fields) == 0 {
		return nil, errors.New("fields is required and cannot be nil or empty")
	}

	path.Grow(7 + 1 + len(strings.Join(r.Index, ",")) + 1 + len("_mapping") + 1 + len("field") + 1 + len(strings.Join(r.Fields, ",")))
	path.WriteString("http://")
	if len(r.Index) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Index, ","))
		if instrument, ok := r.instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "index", strings.Join(r.Index, ","))
		}
	}
	path.WriteString("/")
	path.WriteString("_mapping")
	path.WriteString("/")
	path.WriteString("field")
	path.WriteString("/")
	path.WriteString(strings.Join(r.Fields, ","))
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "fields", strings.Join(r.Fields, ","))
	}

	params = make(map[string]string)

	if r.AllowNoIndices != nil {
		params["allow_no_indices"] = strconv.FormatBool(*r.AllowNoIndices)
	}

	if r.ExpandWildcards != "" {
		params["expand_wildcards"] = r.ExpandWildcards
	}

	if r.IgnoreUnavailable != nil {
		params["ignore_unavailable"] = strconv.FormatBool(*r.IgnoreUnavailable)
	}

	if r.IncludeDefaults != nil {
		params["include_defaults"] = strconv.FormatBool(*r.IncludeDefaults)
	}

	if r.Local != nil {
		params["local"] = strconv.FormatBool(*r.Local)
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
		instrument.BeforeRequest(req, "indices.get_field_mapping")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.get_field_mapping")
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
func (f IndicesGetFieldMapping) WithContext(v context.Context) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.ctx = v
	}
}

// WithIndex - a list of index names.
func (f IndicesGetFieldMapping) WithIndex(v ...string) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.Index = v
	}
}

// WithAllowNoIndices - whether to ignore if a wildcard indices expression resolves into no concrete indices. (this includes `_all` string or when no indices have been specified).
func (f IndicesGetFieldMapping) WithAllowNoIndices(v bool) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.AllowNoIndices = &v
	}
}

// WithExpandWildcards - whether to expand wildcard expression to concrete indices that are open, closed or both..
func (f IndicesGetFieldMapping) WithExpandWildcards(v string) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.ExpandWildcards = v
	}
}

// WithIgnoreUnavailable - whether specified concrete indices should be ignored when unavailable (missing or closed).
func (f IndicesGetFieldMapping) WithIgnoreUnavailable(v bool) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.IgnoreUnavailable = &v
	}
}

// WithIncludeDefaults - whether the default mapping values should be returned as well.
func (f IndicesGetFieldMapping) WithIncludeDefaults(v bool) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.IncludeDefaults = &v
	}
}

// WithLocal - return local information, do not retrieve the state from master node (default: false).
func (f IndicesGetFieldMapping) WithLocal(v bool) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.Local = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IndicesGetFieldMapping) WithPretty() func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IndicesGetFieldMapping) WithHuman() func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IndicesGetFieldMapping) WithErrorTrace() func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IndicesGetFieldMapping) WithFilterPath(v ...string) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IndicesGetFieldMapping) WithHeader(h map[string]string) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IndicesGetFieldMapping) WithOpaqueID(s string) func(*IndicesGetFieldMappingRequest) {
	return func(r *IndicesGetFieldMappingRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
