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
// Code generated from specification version 8.12.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newConnectorLastSyncFunc(t Transport) ConnectorLastSync {
	return func(body io.Reader, connector_id string, o ...func(*ConnectorLastSyncRequest)) (*Response, error) {
		var r = ConnectorLastSyncRequest{Body: body, ConnectorID: connector_id}
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

// ConnectorLastSync updates the stats of last sync in the connector document.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/update-connector-last-sync-api.html.
type ConnectorLastSync func(body io.Reader, connector_id string, o ...func(*ConnectorLastSyncRequest)) (*Response, error)

// ConnectorLastSyncRequest configures the Connector Last Sync API request.
type ConnectorLastSyncRequest struct {
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
func (r ConnectorLastSyncRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "connector.last_sync")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_connector") + 1 + len(r.ConnectorID) + 1 + len("_last_sync"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_connector")
	path.WriteString("/")
	path.WriteString(r.ConnectorID)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "connector_id", r.ConnectorID)
	}
	path.WriteString("/")
	path.WriteString("_last_sync")

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
		instrument.BeforeRequest(req, "connector.last_sync")
		if reader := instrument.RecordRequestBody(ctx, "connector.last_sync", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "connector.last_sync")
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
func (f ConnectorLastSync) WithContext(v context.Context) func(*ConnectorLastSyncRequest) {
	return func(r *ConnectorLastSyncRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ConnectorLastSync) WithPretty() func(*ConnectorLastSyncRequest) {
	return func(r *ConnectorLastSyncRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ConnectorLastSync) WithHuman() func(*ConnectorLastSyncRequest) {
	return func(r *ConnectorLastSyncRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ConnectorLastSync) WithErrorTrace() func(*ConnectorLastSyncRequest) {
	return func(r *ConnectorLastSyncRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ConnectorLastSync) WithFilterPath(v ...string) func(*ConnectorLastSyncRequest) {
	return func(r *ConnectorLastSyncRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ConnectorLastSync) WithHeader(h map[string]string) func(*ConnectorLastSyncRequest) {
	return func(r *ConnectorLastSyncRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ConnectorLastSync) WithOpaqueID(s string) func(*ConnectorLastSyncRequest) {
	return func(r *ConnectorLastSyncRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
