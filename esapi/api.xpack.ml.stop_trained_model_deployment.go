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
// Code generated from specification version 8.6.0: DO NOT EDIT

package esapi

import (
	"context"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func newMLStopTrainedModelDeploymentFunc(t Transport) MLStopTrainedModelDeployment {
	return func(model_id string, o ...func(*MLStopTrainedModelDeploymentRequest)) (*Response, error) {
		var r = MLStopTrainedModelDeploymentRequest{ModelID: model_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// MLStopTrainedModelDeployment - Stop a trained model deployment.
//
// This API is beta.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/stop-trained-model-deployment.html.
type MLStopTrainedModelDeployment func(model_id string, o ...func(*MLStopTrainedModelDeploymentRequest)) (*Response, error)

// MLStopTrainedModelDeploymentRequest configures the ML Stop Trained Model Deployment API request.
type MLStopTrainedModelDeploymentRequest struct {
	Body io.Reader

	ModelID string

	AllowNoMatch *bool
	Force        *bool

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r MLStopTrainedModelDeploymentRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(7 + 1 + len("_ml") + 1 + len("trained_models") + 1 + len(r.ModelID) + 1 + len("deployment") + 1 + len("_stop"))
	path.WriteString("http://")
	path.WriteString("/")
	path.WriteString("_ml")
	path.WriteString("/")
	path.WriteString("trained_models")
	path.WriteString("/")
	path.WriteString(r.ModelID)
	path.WriteString("/")
	path.WriteString("deployment")
	path.WriteString("/")
	path.WriteString("_stop")

	params = make(map[string]string)

	if r.AllowNoMatch != nil {
		params["allow_no_match"] = strconv.FormatBool(*r.AllowNoMatch)
	}

	if r.Force != nil {
		params["force"] = strconv.FormatBool(*r.Force)
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

	res, err := transport.Perform(req)
	if err != nil {
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
func (f MLStopTrainedModelDeployment) WithContext(v context.Context) func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		r.ctx = v
	}
}

// WithBody - The stop deployment parameters.
func (f MLStopTrainedModelDeployment) WithBody(v io.Reader) func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		r.Body = v
	}
}

// WithAllowNoMatch - whether to ignore if a wildcard expression matches no deployments. (this includes `_all` string or when no deployments have been specified).
func (f MLStopTrainedModelDeployment) WithAllowNoMatch(v bool) func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		r.AllowNoMatch = &v
	}
}

// WithForce - true if the deployment should be forcefully stopped.
func (f MLStopTrainedModelDeployment) WithForce(v bool) func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		r.Force = &v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLStopTrainedModelDeployment) WithPretty() func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLStopTrainedModelDeployment) WithHuman() func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLStopTrainedModelDeployment) WithErrorTrace() func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLStopTrainedModelDeployment) WithFilterPath(v ...string) func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLStopTrainedModelDeployment) WithHeader(h map[string]string) func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLStopTrainedModelDeployment) WithOpaqueID(s string) func(*MLStopTrainedModelDeploymentRequest) {
	return func(r *MLStopTrainedModelDeploymentRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
