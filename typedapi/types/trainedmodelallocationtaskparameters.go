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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// TrainedModelAllocationTaskParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ml/_types/TrainedModel.ts#L286-L295
type TrainedModelAllocationTaskParameters struct {
	// ModelBytes The size of the trained model in bytes.
	ModelBytes int `json:"model_bytes"`
	// ModelId The unique identifier for the trained model.
	ModelId Id `json:"model_id"`
}

// TrainedModelAllocationTaskParametersBuilder holds TrainedModelAllocationTaskParameters struct and provides a builder API.
type TrainedModelAllocationTaskParametersBuilder struct {
	v *TrainedModelAllocationTaskParameters
}

// NewTrainedModelAllocationTaskParameters provides a builder for the TrainedModelAllocationTaskParameters struct.
func NewTrainedModelAllocationTaskParametersBuilder() *TrainedModelAllocationTaskParametersBuilder {
	r := TrainedModelAllocationTaskParametersBuilder{
		&TrainedModelAllocationTaskParameters{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelAllocationTaskParameters struct
func (rb *TrainedModelAllocationTaskParametersBuilder) Build() TrainedModelAllocationTaskParameters {
	return *rb.v
}

// ModelBytes The size of the trained model in bytes.

func (rb *TrainedModelAllocationTaskParametersBuilder) ModelBytes(modelbytes int) *TrainedModelAllocationTaskParametersBuilder {
	rb.v.ModelBytes = modelbytes
	return rb
}

// ModelId The unique identifier for the trained model.

func (rb *TrainedModelAllocationTaskParametersBuilder) ModelId(modelid Id) *TrainedModelAllocationTaskParametersBuilder {
	rb.v.ModelId = modelid
	return rb
}
