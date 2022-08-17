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

// MlInferenceDeployments type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L212-L217
type MlInferenceDeployments struct {
	Count           int                          `json:"count"`
	InferenceCounts JobStatistics                `json:"inference_counts"`
	ModelSizesBytes JobStatistics                `json:"model_sizes_bytes"`
	TimeMs          MlInferenceDeploymentsTimeMs `json:"time_ms"`
}

// MlInferenceDeploymentsBuilder holds MlInferenceDeployments struct and provides a builder API.
type MlInferenceDeploymentsBuilder struct {
	v *MlInferenceDeployments
}

// NewMlInferenceDeployments provides a builder for the MlInferenceDeployments struct.
func NewMlInferenceDeploymentsBuilder() *MlInferenceDeploymentsBuilder {
	r := MlInferenceDeploymentsBuilder{
		&MlInferenceDeployments{},
	}

	return &r
}

// Build finalize the chain and returns the MlInferenceDeployments struct
func (rb *MlInferenceDeploymentsBuilder) Build() MlInferenceDeployments {
	return *rb.v
}

func (rb *MlInferenceDeploymentsBuilder) Count(count int) *MlInferenceDeploymentsBuilder {
	rb.v.Count = count
	return rb
}

func (rb *MlInferenceDeploymentsBuilder) InferenceCounts(inferencecounts *JobStatisticsBuilder) *MlInferenceDeploymentsBuilder {
	v := inferencecounts.Build()
	rb.v.InferenceCounts = v
	return rb
}

func (rb *MlInferenceDeploymentsBuilder) ModelSizesBytes(modelsizesbytes *JobStatisticsBuilder) *MlInferenceDeploymentsBuilder {
	v := modelsizesbytes.Build()
	rb.v.ModelSizesBytes = v
	return rb
}

func (rb *MlInferenceDeploymentsBuilder) TimeMs(timems *MlInferenceDeploymentsTimeMsBuilder) *MlInferenceDeploymentsBuilder {
	v := timems.Build()
	rb.v.TimeMs = v
	return rb
}
