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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

import (
	"encoding/json"
)

// TrainedModelStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/TrainedModel.ts#L42-L60
type TrainedModelStats struct {
	// DeploymentStats A collection of deployment stats, which is present when the models are
	// deployed.
	DeploymentStats *TrainedModelDeploymentStats `json:"deployment_stats,omitempty"`
	// InferenceStats A collection of inference stats fields.
	InferenceStats *TrainedModelInferenceStats `json:"inference_stats,omitempty"`
	// Ingest A collection of ingest stats for the model across all nodes.
	// The values are summations of the individual node statistics.
	// The format matches the ingest section in the nodes stats API.
	Ingest map[string]json.RawMessage `json:"ingest,omitempty"`
	// ModelId The unique identifier of the trained model.
	ModelId string `json:"model_id"`
	// ModelSizeStats A collection of model size stats.
	ModelSizeStats TrainedModelSizeStats `json:"model_size_stats"`
	// PipelineCount The number of ingest pipelines that currently refer to the model.
	PipelineCount int `json:"pipeline_count"`
}

// NewTrainedModelStats returns a TrainedModelStats.
func NewTrainedModelStats() *TrainedModelStats {
	r := &TrainedModelStats{
		Ingest: make(map[string]json.RawMessage, 0),
	}

	return r
}
