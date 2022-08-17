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

// TrainedModelAssignmentTaskParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L299-L325
type TrainedModelAssignmentTaskParameters struct {
	// CacheSize The size of the trained model cache.
	CacheSize ByteSize `json:"cache_size"`
	// ModelBytes The size of the trained model in bytes.
	ModelBytes int `json:"model_bytes"`
	// ModelId The unique identifier for the trained model.
	ModelId Id `json:"model_id"`
	// NumberOfAllocations The total number of allocations this model is assigned across ML nodes.
	NumberOfAllocations int `json:"number_of_allocations"`
	// QueueCapacity Number of inference requests are allowed in the queue at a time.
	QueueCapacity int `json:"queue_capacity"`
	// ThreadsPerAllocation Number of threads per allocation.
	ThreadsPerAllocation int `json:"threads_per_allocation"`
}

// TrainedModelAssignmentTaskParametersBuilder holds TrainedModelAssignmentTaskParameters struct and provides a builder API.
type TrainedModelAssignmentTaskParametersBuilder struct {
	v *TrainedModelAssignmentTaskParameters
}

// NewTrainedModelAssignmentTaskParameters provides a builder for the TrainedModelAssignmentTaskParameters struct.
func NewTrainedModelAssignmentTaskParametersBuilder() *TrainedModelAssignmentTaskParametersBuilder {
	r := TrainedModelAssignmentTaskParametersBuilder{
		&TrainedModelAssignmentTaskParameters{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelAssignmentTaskParameters struct
func (rb *TrainedModelAssignmentTaskParametersBuilder) Build() TrainedModelAssignmentTaskParameters {
	return *rb.v
}

// CacheSize The size of the trained model cache.

func (rb *TrainedModelAssignmentTaskParametersBuilder) CacheSize(cachesize *ByteSizeBuilder) *TrainedModelAssignmentTaskParametersBuilder {
	v := cachesize.Build()
	rb.v.CacheSize = v
	return rb
}

// ModelBytes The size of the trained model in bytes.

func (rb *TrainedModelAssignmentTaskParametersBuilder) ModelBytes(modelbytes int) *TrainedModelAssignmentTaskParametersBuilder {
	rb.v.ModelBytes = modelbytes
	return rb
}

// ModelId The unique identifier for the trained model.

func (rb *TrainedModelAssignmentTaskParametersBuilder) ModelId(modelid Id) *TrainedModelAssignmentTaskParametersBuilder {
	rb.v.ModelId = modelid
	return rb
}

// NumberOfAllocations The total number of allocations this model is assigned across ML nodes.

func (rb *TrainedModelAssignmentTaskParametersBuilder) NumberOfAllocations(numberofallocations int) *TrainedModelAssignmentTaskParametersBuilder {
	rb.v.NumberOfAllocations = numberofallocations
	return rb
}

// QueueCapacity Number of inference requests are allowed in the queue at a time.

func (rb *TrainedModelAssignmentTaskParametersBuilder) QueueCapacity(queuecapacity int) *TrainedModelAssignmentTaskParametersBuilder {
	rb.v.QueueCapacity = queuecapacity
	return rb
}

// ThreadsPerAllocation Number of threads per allocation.

func (rb *TrainedModelAssignmentTaskParametersBuilder) ThreadsPerAllocation(threadsperallocation int) *TrainedModelAssignmentTaskParametersBuilder {
	rb.v.ThreadsPerAllocation = threadsperallocation
	return rb
}
