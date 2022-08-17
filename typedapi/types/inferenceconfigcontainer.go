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


package types

// InferenceConfigContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/pipeline.ts#L164-L170
type InferenceConfigContainer struct {
	// Classification Classification configuration for inference.
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`
	// Regression Regression configuration for inference.
	Regression *RegressionInferenceOptions `json:"regression,omitempty"`
}

// InferenceConfigContainerBuilder holds InferenceConfigContainer struct and provides a builder API.
type InferenceConfigContainerBuilder struct {
	v *InferenceConfigContainer
}

// NewInferenceConfigContainer provides a builder for the InferenceConfigContainer struct.
func NewInferenceConfigContainerBuilder() *InferenceConfigContainerBuilder {
	r := InferenceConfigContainerBuilder{
		&InferenceConfigContainer{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceConfigContainer struct
func (rb *InferenceConfigContainerBuilder) Build() InferenceConfigContainer {
	return *rb.v
}

// Classification Classification configuration for inference.

func (rb *InferenceConfigContainerBuilder) Classification(classification *ClassificationInferenceOptionsBuilder) *InferenceConfigContainerBuilder {
	v := classification.Build()
	rb.v.Classification = &v
	return rb
}

// Regression Regression configuration for inference.

func (rb *InferenceConfigContainerBuilder) Regression(regression *RegressionInferenceOptionsBuilder) *InferenceConfigContainerBuilder {
	v := regression.Build()
	rb.v.Regression = &v
	return rb
}
