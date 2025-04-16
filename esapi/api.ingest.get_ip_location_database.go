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

func newIngestGetIPLocationDatabaseFunc(t Transport) IngestGetIPLocationDatabase {
	return func(o ...func(*IngestGetIPLocationDatabaseRequest)) (*Response, error) {
		var r = IngestGetIPLocationDatabaseRequest{}
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

// IngestGetIPLocationDatabase returns the specified ip location database configuration
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/get-ip-location-database-api.html.
type IngestGetIPLocationDatabase func(o ...func(*IngestGetIPLocationDatabaseRequest)) (*Response, error)

// IngestGetIPLocationDatabaseRequest configures the Ingest GetIP Location Database API request.
type IngestGetIPLocationDatabaseRequest struct {
	DocumentID []string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IngestGetIPLocationDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ingest.get_ip_location_database")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_ingest") + 1 + len("ip_location") + 1 + len("database") + 1 + len(strings.Join(r.DocumentID, ",")))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ingest")
	path.WriteString("/")
	path.WriteString("ip_location")
	path.WriteString("/")
	path.WriteString("database")
	if len(r.DocumentID) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.DocumentID, ","))
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "id", strings.Join(r.DocumentID, ","))
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
		instrument.BeforeRequest(req, "ingest.get_ip_location_database")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ingest.get_ip_location_database")
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
func (f IngestGetIPLocationDatabase) WithContext(v context.Context) func(*IngestGetIPLocationDatabaseRequest) {
	return func(r *IngestGetIPLocationDatabaseRequest) {
		r.ctx = v
	}
}

// WithDocumentID - a list of ip location database configurations to get; use `*` to get all ip location database configurations.
func (f IngestGetIPLocationDatabase) WithDocumentID(v ...string) func(*IngestGetIPLocationDatabaseRequest) {
	return func(r *IngestGetIPLocationDatabaseRequest) {
		r.DocumentID = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IngestGetIPLocationDatabase) WithPretty() func(*IngestGetIPLocationDatabaseRequest) {
	return func(r *IngestGetIPLocationDatabaseRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IngestGetIPLocationDatabase) WithHuman() func(*IngestGetIPLocationDatabaseRequest) {
	return func(r *IngestGetIPLocationDatabaseRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IngestGetIPLocationDatabase) WithErrorTrace() func(*IngestGetIPLocationDatabaseRequest) {
	return func(r *IngestGetIPLocationDatabaseRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IngestGetIPLocationDatabase) WithFilterPath(v ...string) func(*IngestGetIPLocationDatabaseRequest) {
	return func(r *IngestGetIPLocationDatabaseRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IngestGetIPLocationDatabase) WithHeader(h map[string]string) func(*IngestGetIPLocationDatabaseRequest) {
	return func(r *IngestGetIPLocationDatabaseRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IngestGetIPLocationDatabase) WithOpaqueID(s string) func(*IngestGetIPLocationDatabaseRequest) {
	return func(r *IngestGetIPLocationDatabaseRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
