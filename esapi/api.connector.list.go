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
	"net/http"
	"strconv"
	"strings"
)

func newConnectorListFunc(t Transport) ConnectorList {
	return func(o ...func(*ConnectorListRequest)) (*Response, error) {
		var r = ConnectorListRequest{}
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

// ConnectorList lists all connectors.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/list-connector-api.html.
type ConnectorList func(o ...func(*ConnectorListRequest)) (*Response, error)

// ConnectorListRequest configures the Connector List API request.
type ConnectorListRequest struct {
	ConnectorName []string
	From          *int
	IndexName     []string
	Query         string
	ServiceType   []string
	Size          *int

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ConnectorListRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "connector.list")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_connector"))
	path.WriteString("http://")
	path.WriteString("/_connector")

	params = make(map[string]string)

	if len(r.ConnectorName) > 0 {
		params["connector_name"] = strings.Join(r.ConnectorName, ",")
	}

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if len(r.IndexName) > 0 {
		params["index_name"] = strings.Join(r.IndexName, ",")
	}

	if r.Query != "" {
		params["query"] = r.Query
	}

	if len(r.ServiceType) > 0 {
		params["service_type"] = strings.Join(r.ServiceType, ",")
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
		instrument.BeforeRequest(req, "connector.list")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "connector.list")
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
func (f ConnectorList) WithContext(v context.Context) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.ctx = v
	}
}

// WithConnectorName - a list of connector names to fetch connector documents for.
func (f ConnectorList) WithConnectorName(v ...string) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.ConnectorName = v
	}
}

// WithFrom - starting offset (default: 0).
func (f ConnectorList) WithFrom(v int) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.From = &v
	}
}

// WithIndexName - a list of connector index names to fetch connector documents for.
func (f ConnectorList) WithIndexName(v ...string) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.IndexName = v
	}
}

// WithQuery - a search string for querying connectors, filtering results by matching against connector names, descriptions, and index names.
func (f ConnectorList) WithQuery(v string) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.Query = v
	}
}

// WithServiceType - a list of connector service types to fetch connector documents for.
func (f ConnectorList) WithServiceType(v ...string) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.ServiceType = v
	}
}

// WithSize - specifies a max number of results to get (default: 100).
func (f ConnectorList) WithSize(v int) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.Size = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ConnectorList) WithPretty() func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ConnectorList) WithHuman() func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ConnectorList) WithErrorTrace() func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ConnectorList) WithFilterPath(v ...string) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ConnectorList) WithHeader(h map[string]string) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ConnectorList) WithOpaqueID(s string) func(*ConnectorListRequest) {
	return func(r *ConnectorListRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
