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
// Code generated from specification version 9.6.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newEsqlTestDataSourceConnectionFunc(t Transport) EsqlTestDataSourceConnection {
	return func(body io.Reader, o ...func(*EsqlTestDataSourceConnectionRequest)) (*Response, error) {
		var r = EsqlTestDataSourceConnectionRequest{Body: body}
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

// EsqlTestDataSourceConnection - Tests whether the supplied data source configuration can establish a live connection. The data source does not need to exist in cluster state — this endpoint is intended for validating a new configuration before saving it.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-esql-data-source-test-connection.
type EsqlTestDataSourceConnection func(body io.Reader, o ...func(*EsqlTestDataSourceConnectionRequest)) (*Response, error)

// EsqlTestDataSourceConnectionRequest configures the Esql Test Data Source Connection API request.
type EsqlTestDataSourceConnectionRequest struct {
	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r EsqlTestDataSourceConnectionRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "esql.test_data_source_connection")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_query/data_source/_test"))
	path.WriteString("http://")
	path.WriteString("/_query/data_source/_test")

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
		instrument.BeforeRequest(req, "esql.test_data_source_connection")
		if reader := instrument.RecordRequestBody(ctx, "esql.test_data_source_connection", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "esql.test_data_source_connection")
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
func (f EsqlTestDataSourceConnection) WithContext(v context.Context) func(*EsqlTestDataSourceConnectionRequest) {
	return func(r *EsqlTestDataSourceConnectionRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f EsqlTestDataSourceConnection) WithPretty() func(*EsqlTestDataSourceConnectionRequest) {
	return func(r *EsqlTestDataSourceConnectionRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f EsqlTestDataSourceConnection) WithHuman() func(*EsqlTestDataSourceConnectionRequest) {
	return func(r *EsqlTestDataSourceConnectionRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f EsqlTestDataSourceConnection) WithErrorTrace() func(*EsqlTestDataSourceConnectionRequest) {
	return func(r *EsqlTestDataSourceConnectionRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f EsqlTestDataSourceConnection) WithFilterPath(v ...string) func(*EsqlTestDataSourceConnectionRequest) {
	return func(r *EsqlTestDataSourceConnectionRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f EsqlTestDataSourceConnection) WithHeader(h map[string]string) func(*EsqlTestDataSourceConnectionRequest) {
	return func(r *EsqlTestDataSourceConnectionRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f EsqlTestDataSourceConnection) WithOpaqueID(s string) func(*EsqlTestDataSourceConnectionRequest) {
	return func(r *EsqlTestDataSourceConnectionRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
