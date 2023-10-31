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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// TrainedModelStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/TrainedModel.ts#L42-L60
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

func (s *TrainedModelStats) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "deployment_stats":
			if err := dec.Decode(&s.DeploymentStats); err != nil {
				return err
			}

		case "inference_stats":
			if err := dec.Decode(&s.InferenceStats); err != nil {
				return err
			}

		case "ingest":
			if s.Ingest == nil {
				s.Ingest = make(map[string]json.RawMessage, 0)
			}
			if err := dec.Decode(&s.Ingest); err != nil {
				return err
			}

		case "model_id":
			if err := dec.Decode(&s.ModelId); err != nil {
				return err
			}

		case "model_size_stats":
			if err := dec.Decode(&s.ModelSizeStats); err != nil {
				return err
			}

		case "pipeline_count":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PipelineCount = value
			case float64:
				f := int(v)
				s.PipelineCount = f
			}

		}
	}
	return nil
}

// NewTrainedModelStats returns a TrainedModelStats.
func NewTrainedModelStats() *TrainedModelStats {
	r := &TrainedModelStats{
		Ingest: make(map[string]json.RawMessage, 0),
	}

	return r
}
