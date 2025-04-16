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

func newConnectorDeleteFunc(t Transport) ConnectorDelete {
	return func(connector_id string, o ...func(*ConnectorDeleteRequest)) (*Response, error) {
		var r = ConnectorDeleteRequest{ConnectorID: connector_id}
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

// ConnectorDelete deletes a connector.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/delete-connector-api.html.
type ConnectorDelete func(connector_id string, o ...func(*ConnectorDeleteRequest)) (*Response, error)

// ConnectorDeleteRequest configures the Connector Delete API request.
type ConnectorDeleteRequest struct {
	ConnectorID string

	DeleteSyncJobs *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ConnectorDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "connector.delete")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "DELETE"

	path.Grow(7 + 1 + len("_connector") + 1 + len(r.ConnectorID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_connector")
	path.WriteString("/")
	path.WriteString(r.ConnectorID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "connector_id", r.ConnectorID)
	}

	params = make(map[string]string)

	if r.DeleteSyncJobs != nil {
		params["delete_sync_jobs"] = strconv.FormatBool(*r.DeleteSyncJobs)
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
		instrument.BeforeRequest(req, "connector.delete")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "connector.delete")
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
func (f ConnectorDelete) WithContext(v context.Context) func(*ConnectorDeleteRequest) {
	return func(r *ConnectorDeleteRequest) {
		r.ctx = v
	}
}

// WithDeleteSyncJobs - determines whether associated sync jobs are also deleted..
func (f ConnectorDelete) WithDeleteSyncJobs(v bool) func(*ConnectorDeleteRequest) {
	return func(r *ConnectorDeleteRequest) {
		r.DeleteSyncJobs = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ConnectorDelete) WithPretty() func(*ConnectorDeleteRequest) {
	return func(r *ConnectorDeleteRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ConnectorDelete) WithHuman() func(*ConnectorDeleteRequest) {
	return func(r *ConnectorDeleteRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ConnectorDelete) WithErrorTrace() func(*ConnectorDeleteRequest) {
	return func(r *ConnectorDeleteRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ConnectorDelete) WithFilterPath(v ...string) func(*ConnectorDeleteRequest) {
	return func(r *ConnectorDeleteRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ConnectorDelete) WithHeader(h map[string]string) func(*ConnectorDeleteRequest) {
	return func(r *ConnectorDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ConnectorDelete) WithOpaqueID(s string) func(*ConnectorDeleteRequest) {
	return func(r *ConnectorDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
