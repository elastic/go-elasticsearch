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
// Code generated from specification version 9.5.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strings"
)

func newEsqlGetDataSourceFunc(t Transport) EsqlGetDataSource {
	return func(o ...func(*EsqlGetDataSourceRequest)) (*Response, error) {
		var r = EsqlGetDataSourceRequest{}
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

// EsqlGetDataSource - Get one or more ES|QL data sources. Secret setting values are returned masked.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-esql-data-source-get.
type EsqlGetDataSource func(o ...func(*EsqlGetDataSourceRequest)) (*Response, error)

// EsqlGetDataSourceRequest configures the Esql Get Data Source API request.
type EsqlGetDataSourceRequest struct {
	Name []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r EsqlGetDataSourceRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "esql.get_data_source")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_query") + 1 + len("data_source") + 1 + len(strings.Join(r.Name, ",")))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_query")
	path.WriteString("/")
	path.WriteString("data_source")
	if len(r.Name) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Name, ","))
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", strings.Join(r.Name, ","))
		}
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

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "esql.get_data_source")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "esql.get_data_source")
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
func (f EsqlGetDataSource) WithContext(v context.Context) func(*EsqlGetDataSourceRequest) {
	return func(r *EsqlGetDataSourceRequest) {
		r.ctx = v
	}
}

// WithName - a list of data source names or wildcard patterns.
func (f EsqlGetDataSource) WithName(v ...string) func(*EsqlGetDataSourceRequest) {
	return func(r *EsqlGetDataSourceRequest) {
		r.Name = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f EsqlGetDataSource) WithPretty() func(*EsqlGetDataSourceRequest) {
	return func(r *EsqlGetDataSourceRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f EsqlGetDataSource) WithHuman() func(*EsqlGetDataSourceRequest) {
	return func(r *EsqlGetDataSourceRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f EsqlGetDataSource) WithErrorTrace() func(*EsqlGetDataSourceRequest) {
	return func(r *EsqlGetDataSourceRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f EsqlGetDataSource) WithFilterPath(v ...string) func(*EsqlGetDataSourceRequest) {
	return func(r *EsqlGetDataSourceRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f EsqlGetDataSource) WithHeader(h map[string]string) func(*EsqlGetDataSourceRequest) {
	return func(r *EsqlGetDataSourceRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f EsqlGetDataSource) WithOpaqueID(s string) func(*EsqlGetDataSourceRequest) {
	return func(r *EsqlGetDataSourceRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
