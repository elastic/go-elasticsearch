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

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

// Evaluate a trained model.
package infertrainedmodel

import (
	gobytes "bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/elastic/elastic-transport-go/v8/elastictransport"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

const (
	modelidMask = iota + 1
)

// ErrBuildPath is returned in case of missing parameters within the build of the request.
var ErrBuildPath = errors.New("cannot build path, check for missing path parameters")

type InferTrainedModel struct {
	transport elastictransport.Interface

	headers http.Header
	values  url.Values
	path    url.URL

	buf *gobytes.Buffer

	req      *Request
	deferred []func(request *Request) error
	raw      io.Reader

	paramSet int

	modelid string
}

// NewInferTrainedModel type alias for index.
type NewInferTrainedModel func(modelid string) *InferTrainedModel

// NewInferTrainedModelFunc returns a new instance of InferTrainedModel with the provided transport.
// Used in the index of the library this allows to retrieve every apis in once place.
func NewInferTrainedModelFunc(tp elastictransport.Interface) NewInferTrainedModel {
	return func(modelid string) *InferTrainedModel {
		n := New(tp)

		n._modelid(modelid)

		return n
	}
}

// Evaluate a trained model.
//
// https://www.elastic.co/guide/en/elasticsearch/reference/master/infer-trained-model.html
func New(tp elastictransport.Interface) *InferTrainedModel {
	r := &InferTrainedModel{
		transport: tp,
		values:    make(url.Values),
		headers:   make(http.Header),
		buf:       gobytes.NewBuffer(nil),

		req: NewRequest(),
	}

	return r
}

// Raw takes a json payload as input which is then passed to the http.Request
// If specified Raw takes precedence on Request method.
func (r *InferTrainedModel) Raw(raw io.Reader) *InferTrainedModel {
	r.raw = raw

	return r
}

// Request allows to set the request property with the appropriate payload.
func (r *InferTrainedModel) Request(req *Request) *InferTrainedModel {
	r.req = req

	return r
}

// HttpRequest returns the http.Request object built from the
// given parameters.
func (r *InferTrainedModel) HttpRequest(ctx context.Context) (*http.Request, error) {
	var path strings.Builder
	var method string
	var req *http.Request

	var err error

	if len(r.deferred) > 0 {
		for _, f := range r.deferred {
			deferredErr := f(r.req)
			if deferredErr != nil {
				return nil, deferredErr
			}
		}
	}

	if r.raw != nil {
		r.buf.ReadFrom(r.raw)
	} else if r.req != nil {

		data, err := json.Marshal(r.req)

		if err != nil {
			return nil, fmt.Errorf("could not serialise request for InferTrainedModel: %w", err)
		}

		r.buf.Write(data)

	}

	r.path.Scheme = "http"

	switch {
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		path.WriteString(r.modelid)
		path.WriteString("/")
		path.WriteString("_infer")

		method = http.MethodPost
	case r.paramSet == modelidMask:
		path.WriteString("/")
		path.WriteString("_ml")
		path.WriteString("/")
		path.WriteString("trained_models")
		path.WriteString("/")

		path.WriteString(r.modelid)
		path.WriteString("/")
		path.WriteString("deployment")
		path.WriteString("/")
		path.WriteString("_infer")

		method = http.MethodPost
	}

	r.path.Path = path.String()
	r.path.RawQuery = r.values.Encode()

	if r.path.Path == "" {
		return nil, ErrBuildPath
	}

	if ctx != nil {
		req, err = http.NewRequestWithContext(ctx, method, r.path.String(), r.buf)
	} else {
		req, err = http.NewRequest(method, r.path.String(), r.buf)
	}

	req.Header = r.headers.Clone()

	if req.Header.Get("Content-Type") == "" {
		if r.buf.Len() > 0 {
			req.Header.Set("Content-Type", "application/vnd.elasticsearch+json;compatible-with=8")
		}
	}

	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/vnd.elasticsearch+json;compatible-with=8")
	}

	if err != nil {
		return req, fmt.Errorf("could not build http.Request: %w", err)
	}

	return req, nil
}

// Perform runs the http.Request through the provided transport and returns an http.Response.
func (r InferTrainedModel) Perform(ctx context.Context) (*http.Response, error) {
	req, err := r.HttpRequest(ctx)
	if err != nil {
		return nil, err
	}

	res, err := r.transport.Perform(req)
	if err != nil {
		return nil, fmt.Errorf("an error happened during the InferTrainedModel query execution: %w", err)
	}

	return res, nil
}

// Do runs the request through the transport, handle the response and returns a infertrainedmodel.Response
func (r InferTrainedModel) Do(ctx context.Context) (*Response, error) {

	response := NewResponse()

	res, err := r.Perform(ctx)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode < 299 {
		err = json.NewDecoder(res.Body).Decode(response)
		if err != nil {
			return nil, err
		}

		return response, nil
	}

	errorResponse := types.NewElasticsearchError()
	err = json.NewDecoder(res.Body).Decode(errorResponse)
	if err != nil {
		return nil, err
	}

	if errorResponse.Status == 0 {
		errorResponse.Status = res.StatusCode
	}

	return nil, errorResponse
}

// Header set a key, value pair in the InferTrainedModel headers map.
func (r *InferTrainedModel) Header(key, value string) *InferTrainedModel {
	r.headers.Set(key, value)

	return r
}

// ModelId The unique identifier of the trained model.
// API Name: modelid
func (r *InferTrainedModel) _modelid(modelid string) *InferTrainedModel {
	r.paramSet |= modelidMask
	r.modelid = modelid

	return r
}

// Timeout Controls the amount of time to wait for inference results.
// API name: timeout
func (r *InferTrainedModel) Timeout(duration string) *InferTrainedModel {
	r.values.Set("timeout", duration)

	return r
}

// Docs An array of objects to pass to the model for inference. The objects should
// contain a fields matching your
// configured trained model input. Typically, for NLP models, the field name is
// `text_field`.
// Currently, for NLP models, only a single value is allowed.
// API name: docs
func (r *InferTrainedModel) Docs(docs ...map[string]json.RawMessage) *InferTrainedModel {
	r.req.Docs = docs

	return r
}

// InferenceConfig The inference configuration updates to apply on the API call
// API name: inference_config
func (r *InferTrainedModel) InferenceConfig(inferenceconfig *types.InferenceConfigUpdateContainer) *InferTrainedModel {

	r.req.InferenceConfig = inferenceconfig

	return r
}
