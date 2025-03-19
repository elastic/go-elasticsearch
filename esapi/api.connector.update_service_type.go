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
// Code generated from specification version 9.0.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newConnectorUpdateServiceDocumentTypeFunc(t Transport) ConnectorUpdateServiceDocumentType {
	return func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateServiceDocumentTypeRequest)) (*Response, error) {
		var r = ConnectorUpdateServiceDocumentTypeRequest{Body: body, ConnectorID: connector_id}
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

// ConnectorUpdateServiceDocumentType updates the service type of the connector.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/update-connector-service-type-api.html.
type ConnectorUpdateServiceDocumentType func(body io.Reader, connector_id string, o ...func(*ConnectorUpdateServiceDocumentTypeRequest)) (*Response, error)

// ConnectorUpdateServiceDocumentTypeRequest configures the Connector Update Service Document Type API request.
type ConnectorUpdateServiceDocumentTypeRequest struct {
	Body io.Reader

	ConnectorID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ConnectorUpdateServiceDocumentTypeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "connector.update_service_type")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_connector") + 1 + len(r.ConnectorID) + 1 + len("_service_type"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_connector")
	path.WriteString("/")
	path.WriteString(r.ConnectorID)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "connector_id", r.ConnectorID)
	}
	path.WriteString("/")
	path.WriteString("_service_type")

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

	if r.Body != nil && req.Header.Get(headerContentType) == "" {
		req.Header[headerContentType] = headerContentTypeJSON
	}

	if ctx != nil {
		req = req.WithContext(ctx)
	}

	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.BeforeRequest(req, "connector.update_service_type")
		if reader := instrument.RecordRequestBody(ctx, "connector.update_service_type", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "connector.update_service_type")
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
func (f ConnectorUpdateServiceDocumentType) WithContext(v context.Context) func(*ConnectorUpdateServiceDocumentTypeRequest) {
	return func(r *ConnectorUpdateServiceDocumentTypeRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ConnectorUpdateServiceDocumentType) WithPretty() func(*ConnectorUpdateServiceDocumentTypeRequest) {
	return func(r *ConnectorUpdateServiceDocumentTypeRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ConnectorUpdateServiceDocumentType) WithHuman() func(*ConnectorUpdateServiceDocumentTypeRequest) {
	return func(r *ConnectorUpdateServiceDocumentTypeRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ConnectorUpdateServiceDocumentType) WithErrorTrace() func(*ConnectorUpdateServiceDocumentTypeRequest) {
	return func(r *ConnectorUpdateServiceDocumentTypeRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ConnectorUpdateServiceDocumentType) WithFilterPath(v ...string) func(*ConnectorUpdateServiceDocumentTypeRequest) {
	return func(r *ConnectorUpdateServiceDocumentTypeRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ConnectorUpdateServiceDocumentType) WithHeader(h map[string]string) func(*ConnectorUpdateServiceDocumentTypeRequest) {
	return func(r *ConnectorUpdateServiceDocumentTypeRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ConnectorUpdateServiceDocumentType) WithOpaqueID(s string) func(*ConnectorUpdateServiceDocumentTypeRequest) {
	return func(r *ConnectorUpdateServiceDocumentTypeRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
