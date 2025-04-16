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
	"strings"
	"time"
)

func newAutoscalingGetAutoscalingPolicyFunc(t Transport) AutoscalingGetAutoscalingPolicy {
	return func(name string, o ...func(*AutoscalingGetAutoscalingPolicyRequest)) (*Response, error) {
		var r = AutoscalingGetAutoscalingPolicyRequest{Name: name}
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

// AutoscalingGetAutoscalingPolicy - Retrieves an autoscaling policy. Designed for indirect use by ECE/ESS and ECK. Direct use is not supported.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-get-autoscaling-policy.html.
type AutoscalingGetAutoscalingPolicy func(name string, o ...func(*AutoscalingGetAutoscalingPolicyRequest)) (*Response, error)

// AutoscalingGetAutoscalingPolicyRequest configures the Autoscaling Get Autoscaling Policy API request.
type AutoscalingGetAutoscalingPolicyRequest struct {
	Name string

	MasterTimeout time.Duration

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context

	Instrument Instrumentation
}

// Do executes the request and returns response or error.
func (r AutoscalingGetAutoscalingPolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "autoscaling.get_autoscaling_policy")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "GET"

	path.Grow(7 + 1 + len("_autoscaling") + 1 + len("policy") + 1 + len(r.Name))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_autoscaling")
	path.WriteString("/")
	path.WriteString("policy")
	path.WriteString("/")
	path.WriteString(r.Name)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.RecordPathPart(ctx, "name", r.Name)
	}

	params = make(map[string]string)

	if r.MasterTimeout != 0 {
		params["master_timeout"] = formatDuration(r.MasterTimeout)
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
		instrument.BeforeRequest(req, "autoscaling.get_autoscaling_policy")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "autoscaling.get_autoscaling_policy")
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
func (f AutoscalingGetAutoscalingPolicy) WithContext(v context.Context) func(*AutoscalingGetAutoscalingPolicyRequest) {
	return func(r *AutoscalingGetAutoscalingPolicyRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - timeout for processing on master node.
func (f AutoscalingGetAutoscalingPolicy) WithMasterTimeout(v time.Duration) func(*AutoscalingGetAutoscalingPolicyRequest) {
	return func(r *AutoscalingGetAutoscalingPolicyRequest) {
		r.MasterTimeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f AutoscalingGetAutoscalingPolicy) WithPretty() func(*AutoscalingGetAutoscalingPolicyRequest) {
	return func(r *AutoscalingGetAutoscalingPolicyRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f AutoscalingGetAutoscalingPolicy) WithHuman() func(*AutoscalingGetAutoscalingPolicyRequest) {
	return func(r *AutoscalingGetAutoscalingPolicyRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f AutoscalingGetAutoscalingPolicy) WithErrorTrace() func(*AutoscalingGetAutoscalingPolicyRequest) {
	return func(r *AutoscalingGetAutoscalingPolicyRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f AutoscalingGetAutoscalingPolicy) WithFilterPath(v ...string) func(*AutoscalingGetAutoscalingPolicyRequest) {
	return func(r *AutoscalingGetAutoscalingPolicyRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f AutoscalingGetAutoscalingPolicy) WithHeader(h map[string]string) func(*AutoscalingGetAutoscalingPolicyRequest) {
	return func(r *AutoscalingGetAutoscalingPolicyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f AutoscalingGetAutoscalingPolicy) WithOpaqueID(s string) func(*AutoscalingGetAutoscalingPolicyRequest) {
	return func(r *AutoscalingGetAutoscalingPolicyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
