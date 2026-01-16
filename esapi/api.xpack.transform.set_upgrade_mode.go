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
// Code generated from specification version 9.3.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newTransformSetUpgradeModeFunc(t Transport) TransformSetUpgradeMode {
	return func(o ...func(*TransformSetUpgradeModeRequest)) (*Response, error) {
		var r = TransformSetUpgradeModeRequest{}
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

// TransformSetUpgradeMode - Set upgrade_mode for transform indices
//
// See full documentation at https://www.elastic.co/docs/api/doc/elasticsearch/operation/operation-transform-set-upgrade-mode.
type TransformSetUpgradeMode func(o ...func(*TransformSetUpgradeModeRequest)) (*Response, error)

// TransformSetUpgradeModeRequest configures the Transform Set Upgrade Mode API request.
type TransformSetUpgradeModeRequest struct {
	Enabled *bool
	Timeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r TransformSetUpgradeModeRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "transform.set_upgrade_mode")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "POST"

	path.Grow(7 + len("/_transform/set_upgrade_mode"))
	path.WriteString("http://")
	path.WriteString("/_transform/set_upgrade_mode")

	params = make(map[string]string)

	if r.Enabled != nil {
		params["enabled"] = strconv.FormatBool(*r.Enabled)
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
		instrument.BeforeRequest(req, "transform.set_upgrade_mode")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "transform.set_upgrade_mode")
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
func (f TransformSetUpgradeMode) WithContext(v context.Context) func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		r.ctx = v
	}
}

// WithEnabled - whether to enable upgrade_mode transform setting or not. defaults to false..
func (f TransformSetUpgradeMode) WithEnabled(v bool) func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		r.Enabled = &v
	}
}

// WithTimeout - controls the time to wait before action times out. defaults to 30 seconds.
func (f TransformSetUpgradeMode) WithTimeout(v time.Duration) func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f TransformSetUpgradeMode) WithPretty() func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f TransformSetUpgradeMode) WithHuman() func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f TransformSetUpgradeMode) WithErrorTrace() func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f TransformSetUpgradeMode) WithFilterPath(v ...string) func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f TransformSetUpgradeMode) WithHeader(h map[string]string) func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f TransformSetUpgradeMode) WithOpaqueID(s string) func(*TransformSetUpgradeModeRequest) {
	return func(r *TransformSetUpgradeModeRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
