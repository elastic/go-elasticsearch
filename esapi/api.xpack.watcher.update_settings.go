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
	"strings"
	"time"
)

func newWatcherUpdateSettingsFunc(t Transport) WatcherUpdateSettings {
	return func(body io.Reader, o ...func(*WatcherUpdateSettingsRequest)) (*Response, error) {
		var r = WatcherUpdateSettingsRequest{Body: body}
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

// WatcherUpdateSettings - Update settings for the watcher system index
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/watcher-api-update-settings.html.
type WatcherUpdateSettings func(body io.Reader, o ...func(*WatcherUpdateSettingsRequest)) (*Response, error)

// WatcherUpdateSettingsRequest configures the Watcher Update Settings API request.
type WatcherUpdateSettingsRequest struct {
	Body io.Reader

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
func (r WatcherUpdateSettingsRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "watcher.update_settings")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "PUT"

	path.Grow(7 + len("/_watcher/settings"))
	path.WriteString("http://")
	path.WriteString("/_watcher/settings")

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
		instrument.BeforeRequest(req, "watcher.update_settings")
		if reader := instrument.RecordRequestBody(ctx, "watcher.update_settings", r.Body); reader != nil {
			req.Body = reader
		}
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "watcher.update_settings")
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
func (f WatcherUpdateSettings) WithContext(v context.Context) func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - specify timeout for connection to master.
func (f WatcherUpdateSettings) WithMasterTimeout(v time.Duration) func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - specify timeout for waiting for acknowledgement from all nodes.
func (f WatcherUpdateSettings) WithTimeout(v time.Duration) func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f WatcherUpdateSettings) WithPretty() func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f WatcherUpdateSettings) WithHuman() func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f WatcherUpdateSettings) WithErrorTrace() func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f WatcherUpdateSettings) WithFilterPath(v ...string) func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f WatcherUpdateSettings) WithHeader(h map[string]string) func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f WatcherUpdateSettings) WithOpaqueID(s string) func(*WatcherUpdateSettingsRequest) {
	return func(r *WatcherUpdateSettingsRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
