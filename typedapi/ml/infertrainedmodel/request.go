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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package infertrainedmodel

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package infertrainedmodel
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/infer_trained_model/MlInferTrainedModelRequest.ts#L27-L59
type Request struct {

	// Docs An array of objects to pass to the model for inference. The objects should
	// contain a fields matching your
	// configured trained model input. Typically, for NLP models, the field name is
	// `text_field`.
	// Currently, for NLP models, only a single value is allowed.
	Docs []map[string]interface{} `json:"docs"`

	// InferenceConfig The inference configuration updates to apply on the API call
	InferenceConfig *types.InferenceConfigUpdateContainer `json:"inference_config,omitempty"`
}

// RequestBuilder is the builder API for the infertrainedmodel.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Infertrainedmodel request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Docs(value ...map[string]interface{}) *RequestBuilder {
	rb.v.Docs = value
	return rb
}

func (rb *RequestBuilder) InferenceConfig(inferenceconfig *types.InferenceConfigUpdateContainerBuilder) *RequestBuilder {
	v := inferenceconfig.Build()
	rb.v.InferenceConfig = &v
	return rb
}
