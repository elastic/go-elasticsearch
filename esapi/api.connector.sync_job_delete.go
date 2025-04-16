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
	"strings"
)

func newConnectorSyncJobDeleteFunc(t Transport) ConnectorSyncJobDelete {
	return func(connector_sync_job_id string, o ...func(*ConnectorSyncJobDeleteRequest)) (*Response, error) {
		var r = ConnectorSyncJobDeleteRequest{ConnectorSyncJobID: connector_sync_job_id}
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

// ConnectorSyncJobDelete deletes a connector sync job.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/delete-connector-sync-job-api.html.
type ConnectorSyncJobDelete func(connector_sync_job_id string, o ...func(*ConnectorSyncJobDeleteRequest)) (*Response, error)

// ConnectorSyncJobDeleteRequest configures the Connector Sync Job Delete API request.
type ConnectorSyncJobDeleteRequest struct {
	ConnectorSyncJobID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ConnectorSyncJobDeleteRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "connector.sync_job_delete")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "DELETE"

	path.Grow(7 + 1 + len("_connector") + 1 + len("_sync_job") + 1 + len(r.ConnectorSyncJobID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_connector")
	path.WriteString("/")
	path.WriteString("_sync_job")
	path.WriteString("/")
	path.WriteString(r.ConnectorSyncJobID)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "connector_sync_job_id", r.ConnectorSyncJobID)
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
		instrument.BeforeRequest(req, "connector.sync_job_delete")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "connector.sync_job_delete")
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
func (f ConnectorSyncJobDelete) WithContext(v context.Context) func(*ConnectorSyncJobDeleteRequest) {
	return func(r *ConnectorSyncJobDeleteRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ConnectorSyncJobDelete) WithPretty() func(*ConnectorSyncJobDeleteRequest) {
	return func(r *ConnectorSyncJobDeleteRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ConnectorSyncJobDelete) WithHuman() func(*ConnectorSyncJobDeleteRequest) {
	return func(r *ConnectorSyncJobDeleteRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ConnectorSyncJobDelete) WithErrorTrace() func(*ConnectorSyncJobDeleteRequest) {
	return func(r *ConnectorSyncJobDeleteRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ConnectorSyncJobDelete) WithFilterPath(v ...string) func(*ConnectorSyncJobDeleteRequest) {
	return func(r *ConnectorSyncJobDeleteRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ConnectorSyncJobDelete) WithHeader(h map[string]string) func(*ConnectorSyncJobDeleteRequest) {
	return func(r *ConnectorSyncJobDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ConnectorSyncJobDelete) WithOpaqueID(s string) func(*ConnectorSyncJobDeleteRequest) {
	return func(r *ConnectorSyncJobDeleteRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
