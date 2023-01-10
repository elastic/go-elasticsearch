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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// TrainedModelAssignmentTaskParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/TrainedModel.ts#L299-L325
type TrainedModelAssignmentTaskParameters struct {
	// CacheSize The size of the trained model cache.
	CacheSize ByteSize `json:"cache_size"`
	// ModelBytes The size of the trained model in bytes.
	ModelBytes int `json:"model_bytes"`
	// ModelId The unique identifier for the trained model.
	ModelId string `json:"model_id"`
	// NumberOfAllocations The total number of allocations this model is assigned across ML nodes.
	NumberOfAllocations int `json:"number_of_allocations"`
	// QueueCapacity Number of inference requests are allowed in the queue at a time.
	QueueCapacity int `json:"queue_capacity"`
	// ThreadsPerAllocation Number of threads per allocation.
	ThreadsPerAllocation int `json:"threads_per_allocation"`
}

// NewTrainedModelAssignmentTaskParameters returns a TrainedModelAssignmentTaskParameters.
func NewTrainedModelAssignmentTaskParameters() *TrainedModelAssignmentTaskParameters {
	r := &TrainedModelAssignmentTaskParameters{}

	return r
}
