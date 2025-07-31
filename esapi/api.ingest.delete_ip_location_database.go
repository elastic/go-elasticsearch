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
	"errors"
	"net/http"
	"strings"
	"time"
)

func newIngestDeleteIPLocationDatabaseFunc(t Transport) IngestDeleteIPLocationDatabase {
	return func(id []string, o ...func(*IngestDeleteIPLocationDatabaseRequest)) (*Response, error) {
		var r = IngestDeleteIPLocationDatabaseRequest{DocumentID: id}
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

// IngestDeleteIPLocationDatabase deletes an ip location database configuration
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/delete-ip-location-database-api.html.
type IngestDeleteIPLocationDatabase func(id []string, o ...func(*IngestDeleteIPLocationDatabaseRequest)) (*Response, error)

// IngestDeleteIPLocationDatabaseRequest configures the Ingest DeleteIP Location Database API request.
type IngestDeleteIPLocationDatabaseRequest struct {
	DocumentID []string

	MasterTimeout time.Duration
	Timeout       time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IngestDeleteIPLocationDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ingest.delete_ip_location_database")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "DELETE"

	if len(r.DocumentID) == 0 {
		return nil, errors.New("id is required and cannot be nil or empty")
	}

	path.Grow(7 + 1 + len("_ingest") + 1 + len("ip_location") + 1 + len("database") + 1 + len(strings.Join(r.DocumentID, ",")))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ingest")
	path.WriteString("/")
	path.WriteString("ip_location")
	path.WriteString("/")
	path.WriteString("database")
	path.WriteString("/")
	path.WriteString(strings.Join(r.DocumentID, ","))
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "id", strings.Join(r.DocumentID, ","))
	}

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
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
		instrument.BeforeRequest(req, "ingest.delete_ip_location_database")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ingest.delete_ip_location_database")
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
func (f IngestDeleteIPLocationDatabase) WithContext(v context.Context) func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - explicit operation timeout for connection to master node.
func (f IngestDeleteIPLocationDatabase) WithMasterTimeout(v time.Duration) func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - explicit operation timeout.
func (f IngestDeleteIPLocationDatabase) WithTimeout(v time.Duration) func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IngestDeleteIPLocationDatabase) WithPretty() func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IngestDeleteIPLocationDatabase) WithHuman() func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IngestDeleteIPLocationDatabase) WithErrorTrace() func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IngestDeleteIPLocationDatabase) WithFilterPath(v ...string) func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IngestDeleteIPLocationDatabase) WithHeader(h map[string]string) func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IngestDeleteIPLocationDatabase) WithOpaqueID(s string) func(*IngestDeleteIPLocationDatabaseRequest) {
	return func(r *IngestDeleteIPLocationDatabaseRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
