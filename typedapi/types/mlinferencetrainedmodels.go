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

// MlInferenceTrainedModels type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L203-L210
type MlInferenceTrainedModels struct {
	All_                          MlCounter                      `json:"_all"`
	Count                         *MlInferenceTrainedModelsCount `json:"count,omitempty"`
	EstimatedHeapMemoryUsageBytes *JobStatistics                 `json:"estimated_heap_memory_usage_bytes,omitempty"`
	EstimatedOperations           *JobStatistics                 `json:"estimated_operations,omitempty"`
	ModelSizeBytes                *JobStatistics                 `json:"model_size_bytes,omitempty"`
}

// MlInferenceTrainedModelsBuilder holds MlInferenceTrainedModels struct and provides a builder API.
type MlInferenceTrainedModelsBuilder struct {
	v *MlInferenceTrainedModels
}

// NewMlInferenceTrainedModels provides a builder for the MlInferenceTrainedModels struct.
func NewMlInferenceTrainedModelsBuilder() *MlInferenceTrainedModelsBuilder {
	r := MlInferenceTrainedModelsBuilder{
		&MlInferenceTrainedModels{},
	}

	return &r
}

// Build finalize the chain and returns the MlInferenceTrainedModels struct
func (rb *MlInferenceTrainedModelsBuilder) Build() MlInferenceTrainedModels {
	return *rb.v
}

func (rb *MlInferenceTrainedModelsBuilder) All_(all_ *MlCounterBuilder) *MlInferenceTrainedModelsBuilder {
	v := all_.Build()
	rb.v.All_ = v
	return rb
}

func (rb *MlInferenceTrainedModelsBuilder) Count(count *MlInferenceTrainedModelsCountBuilder) *MlInferenceTrainedModelsBuilder {
	v := count.Build()
	rb.v.Count = &v
	return rb
}

func (rb *MlInferenceTrainedModelsBuilder) EstimatedHeapMemoryUsageBytes(estimatedheapmemoryusagebytes *JobStatisticsBuilder) *MlInferenceTrainedModelsBuilder {
	v := estimatedheapmemoryusagebytes.Build()
	rb.v.EstimatedHeapMemoryUsageBytes = &v
	return rb
}

func (rb *MlInferenceTrainedModelsBuilder) EstimatedOperations(estimatedoperations *JobStatisticsBuilder) *MlInferenceTrainedModelsBuilder {
	v := estimatedoperations.Build()
	rb.v.EstimatedOperations = &v
	return rb
}

func (rb *MlInferenceTrainedModelsBuilder) ModelSizeBytes(modelsizebytes *JobStatisticsBuilder) *MlInferenceTrainedModelsBuilder {
	v := modelsizebytes.Build()
	rb.v.ModelSizeBytes = &v
	return rb
}
