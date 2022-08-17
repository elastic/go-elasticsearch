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

// TrainedModelStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L42-L60
type TrainedModelStats struct {
	// DeploymentStats A collection of deployment stats, which is present when the models are
	// deployed.
	DeploymentStats *TrainedModelDeploymentStats `json:"deployment_stats,omitempty"`
	// InferenceStats A collection of inference stats fields.
	InferenceStats *TrainedModelInferenceStats `json:"inference_stats,omitempty"`
	// Ingest A collection of ingest stats for the model across all nodes.
	// The values are summations of the individual node statistics.
	// The format matches the ingest section in the nodes stats API.
	Ingest map[string]interface{} `json:"ingest,omitempty"`
	// ModelId The unique identifier of the trained model.
	ModelId Id `json:"model_id"`
	// ModelSizeStats A collection of model size stats.
	ModelSizeStats TrainedModelSizeStats `json:"model_size_stats"`
	// PipelineCount The number of ingest pipelines that currently refer to the model.
	PipelineCount int `json:"pipeline_count"`
}

// TrainedModelStatsBuilder holds TrainedModelStats struct and provides a builder API.
type TrainedModelStatsBuilder struct {
	v *TrainedModelStats
}

// NewTrainedModelStats provides a builder for the TrainedModelStats struct.
func NewTrainedModelStatsBuilder() *TrainedModelStatsBuilder {
	r := TrainedModelStatsBuilder{
		&TrainedModelStats{
			Ingest: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelStats struct
func (rb *TrainedModelStatsBuilder) Build() TrainedModelStats {
	return *rb.v
}

// DeploymentStats A collection of deployment stats, which is present when the models are
// deployed.

func (rb *TrainedModelStatsBuilder) DeploymentStats(deploymentstats *TrainedModelDeploymentStatsBuilder) *TrainedModelStatsBuilder {
	v := deploymentstats.Build()
	rb.v.DeploymentStats = &v
	return rb
}

// InferenceStats A collection of inference stats fields.

func (rb *TrainedModelStatsBuilder) InferenceStats(inferencestats *TrainedModelInferenceStatsBuilder) *TrainedModelStatsBuilder {
	v := inferencestats.Build()
	rb.v.InferenceStats = &v
	return rb
}

// Ingest A collection of ingest stats for the model across all nodes.
// The values are summations of the individual node statistics.
// The format matches the ingest section in the nodes stats API.

func (rb *TrainedModelStatsBuilder) Ingest(value map[string]interface{}) *TrainedModelStatsBuilder {
	rb.v.Ingest = value
	return rb
}

// ModelId The unique identifier of the trained model.

func (rb *TrainedModelStatsBuilder) ModelId(modelid Id) *TrainedModelStatsBuilder {
	rb.v.ModelId = modelid
	return rb
}

// ModelSizeStats A collection of model size stats.

func (rb *TrainedModelStatsBuilder) ModelSizeStats(modelsizestats *TrainedModelSizeStatsBuilder) *TrainedModelStatsBuilder {
	v := modelsizestats.Build()
	rb.v.ModelSizeStats = v
	return rb
}

// PipelineCount The number of ingest pipelines that currently refer to the model.

func (rb *TrainedModelStatsBuilder) PipelineCount(pipelinecount int) *TrainedModelStatsBuilder {
	rb.v.PipelineCount = pipelinecount
	return rb
}
