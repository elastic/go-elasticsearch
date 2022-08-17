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

// InferenceConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ingest/_types/Processors.ts#L244-L250
type InferenceConfig struct {
	Classification *InferenceConfigClassification `json:"classification,omitempty"`
	Regression     *InferenceConfigRegression     `json:"regression,omitempty"`
}

// InferenceConfigBuilder holds InferenceConfig struct and provides a builder API.
type InferenceConfigBuilder struct {
	v *InferenceConfig
}

// NewInferenceConfig provides a builder for the InferenceConfig struct.
func NewInferenceConfigBuilder() *InferenceConfigBuilder {
	r := InferenceConfigBuilder{
		&InferenceConfig{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceConfig struct
func (rb *InferenceConfigBuilder) Build() InferenceConfig {
	return *rb.v
}

func (rb *InferenceConfigBuilder) Classification(classification *InferenceConfigClassificationBuilder) *InferenceConfigBuilder {
	v := classification.Build()
	rb.v.Classification = &v
	return rb
}

func (rb *InferenceConfigBuilder) Regression(regression *InferenceConfigRegressionBuilder) *InferenceConfigBuilder {
	v := regression.Build()
	rb.v.Regression = &v
	return rb
}
