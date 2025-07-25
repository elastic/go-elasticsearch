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
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newIndicesPutDataStreamSettingsFunc(t Transport) IndicesPutDataStreamSettings {
	return func(name string, body io.Reader, o ...func(*IndicesPutDataStreamSettingsRequest)) (*Response, error) {
		var r = IndicesPutDataStreamSettingsRequest{Name: name, Body: body}
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

// IndicesPutDataStreamSettings updates a data stream's settings
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/data-streams.html.
type IndicesPutDataStreamSettings func(name string, body io.Reader, o ...func(*IndicesPutDataStreamSettingsRequest)) (*Response, error)

// IndicesPutDataStreamSettingsRequest configures the Indices Put Data Stream Settings API request.
type IndicesPutDataStreamSettingsRequest struct {
	Body io.Reader

	Name string

	DryRun        *bool
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
func (r IndicesPutDataStreamSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "indices.put_data_stream_settings")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + 1 + len("_data_stream") + 1 + len(r.Name) + 1 + len("_settings"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_data_stream")
	path.WriteString("/")
	path.WriteString(r.Name)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "name", r.Name)
	}
	path.WriteString("/")
	path.WriteString("_settings")

	params = make(map[string]string)

	if r.DryRun != nil {
		params["dry_run"] = strconv.FormatBool(*r.DryRun)
	}

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
		instrument.BeforeRequest(req, "indices.put_data_stream_settings")
		if reader := instrument.RecordRequestBody(ctx, "indices.put_data_stream_settings", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "indices.put_data_stream_settings")
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
func (f IndicesPutDataStreamSettings) WithContext(v context.Context) func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		r.ctx = v
	}
}

// WithDryRun - whether this request should only be a dry run rather than actually applying settings.
func (f IndicesPutDataStreamSettings) WithDryRun(v bool) func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		r.DryRun = &v
	}
}

// WithMasterTimeout - period to wait for a connection to the master node.
func (f IndicesPutDataStreamSettings) WithMasterTimeout(v time.Duration) func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - period to wait for a response.
func (f IndicesPutDataStreamSettings) WithTimeout(v time.Duration) func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f IndicesPutDataStreamSettings) WithPretty() func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f IndicesPutDataStreamSettings) WithHuman() func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f IndicesPutDataStreamSettings) WithErrorTrace() func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f IndicesPutDataStreamSettings) WithFilterPath(v ...string) func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f IndicesPutDataStreamSettings) WithHeader(h map[string]string) func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f IndicesPutDataStreamSettings) WithOpaqueID(s string) func(*IndicesPutDataStreamSettingsRequest) {
	return func(r *IndicesPutDataStreamSettingsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
