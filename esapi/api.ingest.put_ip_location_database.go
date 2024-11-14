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
// Code generated from specification version 8.16.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strings"
)

func newIngestPutIPLocationDatabaseFunc(t Transport) IngestPutIPLocationDatabase {
	return func(id string, body io.Reader, o ...func(*IngestPutIPLocationDatabaseRequest)) (*Response, error) {
		var r = IngestPutIPLocationDatabaseRequest{DocumentID: id, Body: body}
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

// IngestPutIPLocationDatabase puts the configuration for a ip location database to be downloaded
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/put-ip-location-database-api.html.
type IngestPutIPLocationDatabase func(id string, body io.Reader, o ...func(*IngestPutIPLocationDatabaseRequest)) (*Response, error)

// IngestPutIPLocationDatabaseRequest configures the Ingest PutIP Location Database API request.
type IngestPutIPLocationDatabaseRequest struct {
	DocumentID string

	Body io.Reader

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r IngestPutIPLocationDatabaseRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "ingest.put_ip_location_database")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_ingest") + 1 + len("ip_location") + 1 + len("database") + 1 + len(r.DocumentID))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ingest")
	path.WriteString("/")
	path.WriteString("ip_location")
	path.WriteString("/")
	path.WriteString("database")
	path.WriteString("/")
	path.WriteString(r.DocumentID)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "id", r.DocumentID)
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
		instrument.BeforeRequest(req, "ingest.put_ip_location_database")
		if reader := instrument.RecordRequestBody(ctx, "ingest.put_ip_location_database", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "ingest.put_ip_location_database")
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
func (f IngestPutIPLocationDatabase) WithContext(v context.Context) func(*IngestPutIPLocationDatabaseRequest) {
	return func(r *IngestPutIPLocationDatabaseRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IngestPutIPLocationDatabase) WithPretty() func(*IngestPutIPLocationDatabaseRequest) {
	return func(r *IngestPutIPLocationDatabaseRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IngestPutIPLocationDatabase) WithHuman() func(*IngestPutIPLocationDatabaseRequest) {
	return func(r *IngestPutIPLocationDatabaseRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IngestPutIPLocationDatabase) WithErrorTrace() func(*IngestPutIPLocationDatabaseRequest) {
	return func(r *IngestPutIPLocationDatabaseRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IngestPutIPLocationDatabase) WithFilterPath(v ...string) func(*IngestPutIPLocationDatabaseRequest) {
	return func(r *IngestPutIPLocationDatabaseRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IngestPutIPLocationDatabase) WithHeader(h map[string]string) func(*IngestPutIPLocationDatabaseRequest) {
	return func(r *IngestPutIPLocationDatabaseRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IngestPutIPLocationDatabase) WithOpaqueID(s string) func(*IngestPutIPLocationDatabaseRequest) {
	return func(r *IngestPutIPLocationDatabaseRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
