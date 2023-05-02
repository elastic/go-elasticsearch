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
// Code generated from specification version 8.7.1: DO NOT EDIT

package esapi

import (
	"context"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func newMLStartTrainedModelDeploymentFunc(t Transport) MLStartTrainedModelDeployment {
	return func(model_id string, o ...func(*MLStartTrainedModelDeploymentRequest)) (*Response, error) {
		var r = MLStartTrainedModelDeploymentRequest{ModelID: model_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// MLStartTrainedModelDeployment - Start a trained model deployment.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/start-trained-model-deployment.html.
type MLStartTrainedModelDeployment func(model_id string, o ...func(*MLStartTrainedModelDeploymentRequest)) (*Response, error)

// MLStartTrainedModelDeploymentRequest configures the ML Start Trained Model Deployment API request.
type MLStartTrainedModelDeploymentRequest struct {
	ModelID string

	CacheSize            string
	NumberOfAllocations  *int
	Priority             string
	QueueCapacity        *int
	ThreadsPerAllocation *int
	Timeout              time.Duration
	WaitFor              string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r MLStartTrainedModelDeploymentRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(7 + 1 + len("_ml") + 1 + len("trained_models") + 1 + len(r.ModelID) + 1 + len("deployment") + 1 + len("_start"))
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
	path.WriteString("_start")

	params = make(map[string]string)

	if r.CacheSize != "" {
		params["cache_size"] = r.CacheSize
	}

	if r.NumberOfAllocations != nil {
		params["number_of_allocations"] = strconv.FormatInt(int64(*r.NumberOfAllocations), 10)
	}

	if r.Priority != "" {
		params["priority"] = r.Priority
	}

	if r.QueueCapacity != nil {
		params["queue_capacity"] = strconv.FormatInt(int64(*r.QueueCapacity), 10)
	}

	if r.ThreadsPerAllocation != nil {
		params["threads_per_allocation"] = strconv.FormatInt(int64(*r.ThreadsPerAllocation), 10)
	}

	if r.Timeout != 0 {
		params["timeout"] = formatDuration(r.Timeout)
	}

	if r.WaitFor != "" {
		params["wait_for"] = r.WaitFor
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
func (f MLStartTrainedModelDeployment) WithContext(v context.Context) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.ctx = v
	}
}

// WithCacheSize - a byte-size value for configuring the inference cache size. for example, 20mb..
func (f MLStartTrainedModelDeployment) WithCacheSize(v string) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.CacheSize = v
	}
}

// WithNumberOfAllocations - the total number of allocations this model is assigned across machine learning nodes..
func (f MLStartTrainedModelDeployment) WithNumberOfAllocations(v int) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.NumberOfAllocations = &v
	}
}

// WithPriority - the deployment priority..
func (f MLStartTrainedModelDeployment) WithPriority(v string) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.Priority = v
	}
}

// WithQueueCapacity - controls how many inference requests are allowed in the queue at a time..
func (f MLStartTrainedModelDeployment) WithQueueCapacity(v int) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.QueueCapacity = &v
	}
}

// WithThreadsPerAllocation - the number of threads used by each model allocation during inference..
func (f MLStartTrainedModelDeployment) WithThreadsPerAllocation(v int) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.ThreadsPerAllocation = &v
	}
}

// WithTimeout - controls the amount of time to wait for the model to deploy..
func (f MLStartTrainedModelDeployment) WithTimeout(v time.Duration) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.Timeout = v
	}
}

// WithWaitFor - the allocation status for which to wait.
func (f MLStartTrainedModelDeployment) WithWaitFor(v string) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.WaitFor = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLStartTrainedModelDeployment) WithPretty() func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLStartTrainedModelDeployment) WithHuman() func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLStartTrainedModelDeployment) WithErrorTrace() func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLStartTrainedModelDeployment) WithFilterPath(v ...string) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLStartTrainedModelDeployment) WithHeader(h map[string]string) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLStartTrainedModelDeployment) WithOpaqueID(s string) func(*MLStartTrainedModelDeploymentRequest) {
	return func(r *MLStartTrainedModelDeploymentRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
