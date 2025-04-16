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
	"time"
)

func newAutoscalingDeleteAutoscalingPolicyFunc(t Transport) AutoscalingDeleteAutoscalingPolicy {
	return func(name string, o ...func(*AutoscalingDeleteAutoscalingPolicyRequest)) (*Response, error) {
		var r = AutoscalingDeleteAutoscalingPolicyRequest{Name: name}
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

// AutoscalingDeleteAutoscalingPolicy - Deletes an autoscaling policy. Designed for indirect use by ECE/ESS and ECK. Direct use is not supported.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/current/autoscaling-delete-autoscaling-policy.html.
type AutoscalingDeleteAutoscalingPolicy func(name string, o ...func(*AutoscalingDeleteAutoscalingPolicyRequest)) (*Response, error)

// AutoscalingDeleteAutoscalingPolicyRequest configures the Autoscaling Delete Autoscaling Policy API request.
type AutoscalingDeleteAutoscalingPolicyRequest struct {
	Name string

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
func (r AutoscalingDeleteAutoscalingPolicyRequest) Do(providedCtx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
		ctx    context.Context
	)

	if instrument, ok := r.Instrument.(Instrumentation); ok {
		ctx = instrument.Start(providedCtx, "autoscaling.delete_autoscaling_policy")
		defer instrument.Close(ctx)
	}
	if ctx == nil {
		ctx = providedCtx
	}

	method = "DELETE"

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
		instrument.BeforeRequest(req, "autoscaling.delete_autoscaling_policy")
	}
	res, err := transport.Perform(req)
	if instrument, ok := r.Instrument.(Instrumentation); ok {
		instrument.AfterRequest(req, "elasticsearch", "autoscaling.delete_autoscaling_policy")
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
func (f AutoscalingDeleteAutoscalingPolicy) WithContext(v context.Context) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		r.ctx = v
	}
}

// WithMasterTimeout - timeout for processing on master node.
func (f AutoscalingDeleteAutoscalingPolicy) WithMasterTimeout(v time.Duration) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		r.MasterTimeout = v
	}
}

// WithTimeout - timeout for acknowledgement of update from all nodes in cluster.
func (f AutoscalingDeleteAutoscalingPolicy) WithTimeout(v time.Duration) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		r.Timeout = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f AutoscalingDeleteAutoscalingPolicy) WithPretty() func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f AutoscalingDeleteAutoscalingPolicy) WithHuman() func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f AutoscalingDeleteAutoscalingPolicy) WithErrorTrace() func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f AutoscalingDeleteAutoscalingPolicy) WithFilterPath(v ...string) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f AutoscalingDeleteAutoscalingPolicy) WithHeader(h map[string]string) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f AutoscalingDeleteAutoscalingPolicy) WithOpaqueID(s string) func(*AutoscalingDeleteAutoscalingPolicyRequest) {
	return func(r *AutoscalingDeleteAutoscalingPolicyRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
