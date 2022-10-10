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
	"net/http"
	"strings"
)

func newMLClearTrainedModelDeploymentCacheFunc(t Transport) MLClearTrainedModelDeploymentCache {
	return func(model_id string, o ...func(*MLClearTrainedModelDeploymentCacheRequest)) (*Response, error) {
		var r = MLClearTrainedModelDeploymentCacheRequest{ModelID: model_id}
		for _, f := range o {
			f(&r)
		}
		return r.Do(r.ctx, t)
	}
}

// ----- API Definition -------------------------------------------------------

// MLClearTrainedModelDeploymentCache - Clear the cached results from a trained model deployment
//
// This API is beta.
//
// See full documentation at https://www.elastic.co/guide/en/elasticsearch/reference/master/clear-trained-model-deployment-cache.html.
type MLClearTrainedModelDeploymentCache func(model_id string, o ...func(*MLClearTrainedModelDeploymentCacheRequest)) (*Response, error)

// MLClearTrainedModelDeploymentCacheRequest configures the ML Clear Trained Model Deployment Cache API request.
type MLClearTrainedModelDeploymentCacheRequest struct {
	ModelID string

	Pretty     bool
	Human      bool
	ErrorTrace bool
	FilterPath []string

	Header http.Header

	ctx context.Context
}

// Do executes the request and returns response or error.
func (r MLClearTrainedModelDeploymentCacheRequest) Do(ctx context.Context, transport Transport) (*Response, error) {
	var (
		method string
		path   strings.Builder
		params map[string]string
	)

	method = "POST"

	path.Grow(7 + 1 + len("_ml") + 1 + len("trained_models") + 1 + len(r.ModelID) + 1 + len("deployment") + 1 + len("cache") + 1 + len("_clear"))
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
	path.WriteString("cache")
	path.WriteString("/")
	path.WriteString("_clear")

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
func (f MLClearTrainedModelDeploymentCache) WithContext(v context.Context) func(*MLClearTrainedModelDeploymentCacheRequest) {
	return func(r *MLClearTrainedModelDeploymentCacheRequest) {
		r.ctx = v
	}
}

// WithPretty makes the response body pretty-printed.
func (f MLClearTrainedModelDeploymentCache) WithPretty() func(*MLClearTrainedModelDeploymentCacheRequest) {
	return func(r *MLClearTrainedModelDeploymentCacheRequest) {
		r.Pretty = true
	}
}

// WithHuman makes statistical values human-readable.
func (f MLClearTrainedModelDeploymentCache) WithHuman() func(*MLClearTrainedModelDeploymentCacheRequest) {
	return func(r *MLClearTrainedModelDeploymentCacheRequest) {
		r.Human = true
	}
}

// WithErrorTrace includes the stack trace for errors in the response body.
func (f MLClearTrainedModelDeploymentCache) WithErrorTrace() func(*MLClearTrainedModelDeploymentCacheRequest) {
	return func(r *MLClearTrainedModelDeploymentCacheRequest) {
		r.ErrorTrace = true
	}
}

// WithFilterPath filters the properties of the response body.
func (f MLClearTrainedModelDeploymentCache) WithFilterPath(v ...string) func(*MLClearTrainedModelDeploymentCacheRequest) {
	return func(r *MLClearTrainedModelDeploymentCacheRequest) {
		r.FilterPath = v
	}
}

// WithHeader adds the headers to the HTTP request.
func (f MLClearTrainedModelDeploymentCache) WithHeader(h map[string]string) func(*MLClearTrainedModelDeploymentCacheRequest) {
	return func(r *MLClearTrainedModelDeploymentCacheRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		for k, v := range h {
			r.Header.Add(k, v)
		}
	}
}

// WithOpaqueID adds the X-Opaque-Id header to the HTTP request.
func (f MLClearTrainedModelDeploymentCache) WithOpaqueID(s string) func(*MLClearTrainedModelDeploymentCacheRequest) {
	return func(r *MLClearTrainedModelDeploymentCacheRequest) {
		if r.Header == nil {
			r.Header = make(http.Header)
		}
		r.Header.Set("X-Opaque-Id", s)
	}
}
