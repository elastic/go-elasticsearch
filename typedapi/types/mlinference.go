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

// MlInference type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L189-L194
type MlInference struct {
	Deployments      *MlInferenceDeployments               `json:"deployments,omitempty"`
	IngestProcessors map[string]MlInferenceIngestProcessor `json:"ingest_processors"`
	TrainedModels    MlInferenceTrainedModels              `json:"trained_models"`
}

// MlInferenceBuilder holds MlInference struct and provides a builder API.
type MlInferenceBuilder struct {
	v *MlInference
}

// NewMlInference provides a builder for the MlInference struct.
func NewMlInferenceBuilder() *MlInferenceBuilder {
	r := MlInferenceBuilder{
		&MlInference{
			IngestProcessors: make(map[string]MlInferenceIngestProcessor, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the MlInference struct
func (rb *MlInferenceBuilder) Build() MlInference {
	return *rb.v
}

func (rb *MlInferenceBuilder) Deployments(deployments *MlInferenceDeploymentsBuilder) *MlInferenceBuilder {
	v := deployments.Build()
	rb.v.Deployments = &v
	return rb
}

func (rb *MlInferenceBuilder) IngestProcessors(values map[string]*MlInferenceIngestProcessorBuilder) *MlInferenceBuilder {
	tmp := make(map[string]MlInferenceIngestProcessor, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.IngestProcessors = tmp
	return rb
}

func (rb *MlInferenceBuilder) TrainedModels(trainedmodels *MlInferenceTrainedModelsBuilder) *MlInferenceBuilder {
	v := trainedmodels.Build()
	rb.v.TrainedModels = v
	return rb
}
