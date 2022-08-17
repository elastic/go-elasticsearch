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

// MlInferenceTrainedModelsCount type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L229-L236
type MlInferenceTrainedModelsCount struct {
	Classification *int64 `json:"classification,omitempty"`
	Ner            *int64 `json:"ner,omitempty"`
	Other          int64  `json:"other"`
	Prepackaged    int64  `json:"prepackaged"`
	Regression     *int64 `json:"regression,omitempty"`
	Total          int64  `json:"total"`
}

// MlInferenceTrainedModelsCountBuilder holds MlInferenceTrainedModelsCount struct and provides a builder API.
type MlInferenceTrainedModelsCountBuilder struct {
	v *MlInferenceTrainedModelsCount
}

// NewMlInferenceTrainedModelsCount provides a builder for the MlInferenceTrainedModelsCount struct.
func NewMlInferenceTrainedModelsCountBuilder() *MlInferenceTrainedModelsCountBuilder {
	r := MlInferenceTrainedModelsCountBuilder{
		&MlInferenceTrainedModelsCount{},
	}

	return &r
}

// Build finalize the chain and returns the MlInferenceTrainedModelsCount struct
func (rb *MlInferenceTrainedModelsCountBuilder) Build() MlInferenceTrainedModelsCount {
	return *rb.v
}

func (rb *MlInferenceTrainedModelsCountBuilder) Classification(classification int64) *MlInferenceTrainedModelsCountBuilder {
	rb.v.Classification = &classification
	return rb
}

func (rb *MlInferenceTrainedModelsCountBuilder) Ner(ner int64) *MlInferenceTrainedModelsCountBuilder {
	rb.v.Ner = &ner
	return rb
}

func (rb *MlInferenceTrainedModelsCountBuilder) Other(other int64) *MlInferenceTrainedModelsCountBuilder {
	rb.v.Other = other
	return rb
}

func (rb *MlInferenceTrainedModelsCountBuilder) Prepackaged(prepackaged int64) *MlInferenceTrainedModelsCountBuilder {
	rb.v.Prepackaged = prepackaged
	return rb
}

func (rb *MlInferenceTrainedModelsCountBuilder) Regression(regression int64) *MlInferenceTrainedModelsCountBuilder {
	rb.v.Regression = &regression
	return rb
}

func (rb *MlInferenceTrainedModelsCountBuilder) Total(total int64) *MlInferenceTrainedModelsCountBuilder {
	rb.v.Total = total
	return rb
}
