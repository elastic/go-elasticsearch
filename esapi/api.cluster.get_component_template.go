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
// Code generated from specification version 9.1.0: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newClusterGetComponentTemplateFunc(t Transport) ClusterGetComponentTemplate {
	return func(o ...func(*ClusterGetComponentTemplateRequest)) (*Response, error) {
		var r = ClusterGetComponentTemplateRequest{}
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

// ClusterGetComponentTemplate returns one or more component templates
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/indices-component-template.html.
type ClusterGetComponentTemplate func(o ...func(*ClusterGetComponentTemplateRequest)) (*Response, error)

// ClusterGetComponentTemplateRequest configures the Cluster Get Component Template API request.
type ClusterGetComponentTemplateRequest struct {
	Name []string

	FlatSettings    *bool
	IncludeDefaults *bool
	Local           *bool
	MasterTimeout   time.Duration
	SettingsFilter  string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r ClusterGetComponentTemplateRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "cluster.get_component_template")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_component_template") + 1 + len(strings.Join(r.Name, ",")))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_component_template")
	if len(r.Name) > 0 {
		path.WriteString("/")
		path.WriteString(strings.Join(r.Name, ","))
		if instrument, ok := r.Instrument.(Instrumentation); ok {
			instrument.RecordPathPart(ctx, "name", strings.Join(r.Name, ","))
		}
	}

	params = make(map[string]string)

	if r.FlatSettings != nil {
		params["flat_settings"] = strconv.FormatBool(*r.FlatSettings)
	}

	if r.IncludeDefaults != nil {
		params["include_defaults"] = strconv.FormatBool(*r.IncludeDefaults)
	}

	if r.Local != nil {
		params["local"] = strconv.FormatBool(*r.Local)
	}

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
	}

	if r.SettingsFilter != "" {
		params["settings_filter"] = r.SettingsFilter
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
		instrument.BeforeRequest(req, "cluster.get_component_template")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "cluster.get_component_template")
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
func (f ClusterGetComponentTemplate) WithContext(v context.Context) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.ctx = v
	}
}

// WithName - the comma separated names of the component templates.
func (f ClusterGetComponentTemplate) WithName(v ...string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.Name = v
	}
}

// WithFlatSettings - return settings in flat format (default: false).
func (f ClusterGetComponentTemplate) WithFlatSettings(v bool) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.FlatSettings = &v
	}
}

// WithIncludeDefaults - return all default configurations for the component template (default: false).
func (f ClusterGetComponentTemplate) WithIncludeDefaults(v bool) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.IncludeDefaults = &v
	}
}

// WithLocal - return local information, do not retrieve the state from master node (default: false).
func (f ClusterGetComponentTemplate) WithLocal(v bool) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.Local = &v
	}
}

// WithMasterTimeout - timeout for waiting for new cluster state in case it is blocked.
func (f ClusterGetComponentTemplate) WithMasterTimeout(v time.Duration) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.MasterTimeout = v
	}
}

// WithSettingsFilter - filter out results, for example to filter out sensitive information. supports wildcards or full settings keys.
func (f ClusterGetComponentTemplate) WithSettingsFilter(v string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.SettingsFilter = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f ClusterGetComponentTemplate) WithPretty() func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f ClusterGetComponentTemplate) WithHuman() func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f ClusterGetComponentTemplate) WithErrorTrace() func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f ClusterGetComponentTemplate) WithFilterPath(v ...string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f ClusterGetComponentTemplate) WithHeader(h map[string]string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f ClusterGetComponentTemplate) WithOpaqueID(s string) func(*ClusterGetComponentTemplateRequest) {
	return func(r *ClusterGetComponentTemplateRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
