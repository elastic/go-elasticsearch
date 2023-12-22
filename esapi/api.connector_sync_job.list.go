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
	"net/http"
	"strconv"
	"strings"
)

func newConnectorSyncJobListFunc(t Transport) ConnectorSyncJobList {
	return func(o ...func(*ConnectorSyncJobListRequest)) (*Response, error) {
		var r = ConnectorSyncJobListRequest{}
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

// ConnectorSyncJobList lists all connector sync jobs.
//
// This API is experimental.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/list-connector-sync-jobs-api.html.
type ConnectorSyncJobList func(o ...func(*ConnectorSyncJobListRequest)) (*Response, error)

// ConnectorSyncJobListRequest configures the Connector Sync Job List API request.
type ConnectorSyncJobListRequest struct {
	ConnectorID string
	From        *int
	Size        *int
	Status      string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ConnectorSyncJobListRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "connector_sync_job.list")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + len("/_connector/_sync_job"))
	path.WriteString("http://")
	path.WriteString("/_connector/_sync_job")

	params = make(map[string]string)

	if r.ConnectorID != "" {
		params["connector_id"] = r.ConnectorID
	}

	if r.From != nil {
		params["from"] = strconv.FormatInt(int64(*r.From), 10)
	}

	if r.Size != nil {
		params["size"] = strconv.FormatInt(int64(*r.Size), 10)
	}

	if r.Status != "" {
		params["status"] = r.Status
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
		instrument.BeforeRequest(req, "connector_sync_job.list")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "connector_sync_job.list")
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
func (f ConnectorSyncJobList) WithContext(v context.Context) func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.ctx = v
	}
}

// WithConnectorID - ID of the connector to fetch the sync jobs for.
func (f ConnectorSyncJobList) WithConnectorID(v string) func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.ConnectorID = v
	}
}

// WithFrom - starting offset (default: 0).
func (f ConnectorSyncJobList) WithFrom(v int) func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.From = &v
	}
}

// WithSize - specifies a max number of results to get (default: 100).
func (f ConnectorSyncJobList) WithSize(v int) func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.Size = &v
	}
}

// WithStatus - sync job status, which sync jobs are fetched for.
func (f ConnectorSyncJobList) WithStatus(v string) func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.Status = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ConnectorSyncJobList) WithPretty() func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ConnectorSyncJobList) WithHuman() func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ConnectorSyncJobList) WithErrorTrace() func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ConnectorSyncJobList) WithFilterPath(v ...string) func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ConnectorSyncJobList) WithHeader(h map[string]string) func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ConnectorSyncJobList) WithOpaqueID(s string) func(*ConnectorSyncJobListRequest) {
	return func(r *ConnectorSyncJobListRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
