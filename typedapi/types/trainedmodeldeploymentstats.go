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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentstate"
)

// TrainedModelDeploymentStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L62-L96
type TrainedModelDeploymentStats struct {
	// AllocationStatus The detailed allocation status for the deployment.
	AllocationStatus TrainedModelDeploymentAllocationStatus `json:"allocation_status"`
	// ErrorCount The sum of `error_count` for all nodes in the deployment.
	ErrorCount int `json:"error_count"`
	// InferenceCount The sum of `inference_count` for all nodes in the deployment.
	InferenceCount int `json:"inference_count"`
	// ModelId The unique identifier for the trained model.
	ModelId Id `json:"model_id"`
	// Nodes The deployent stats for each node that currently has the model allocated.
	Nodes TrainedModelDeploymentNodesStats `json:"nodes"`
	// NumberOfAllocations The number of allocations requested.
	NumberOfAllocations int `json:"number_of_allocations"`
	// QueueCapacity The number of inference requests that can be queued before new requests are
	// rejected.
	QueueCapacity int `json:"queue_capacity"`
	// Reason The reason for the current deployment state. Usually only populated when
	// the model is not deployed to a node.
	Reason string `json:"reason"`
	// RejectedExecutionCount The sum of `rejected_execution_count` for all nodes in the deployment.
	// Individual nodes reject an inference request if the inference queue is full.
	// The queue size is controlled by the `queue_capacity` setting in the start
	// trained model deployment API.
	RejectedExecutionCount int `json:"rejected_execution_count"`
	// StartTime The epoch timestamp when the deployment started.
	StartTime EpochTimeUnitMillis `json:"start_time"`
	// State The overall state of the deployment.
	State deploymentstate.DeploymentState `json:"state"`
	// ThreadsPerAllocation The number of threads used be each allocation during inference.
	ThreadsPerAllocation int `json:"threads_per_allocation"`
	// TimeoutCount The sum of `timeout_count` for all nodes in the deployment.
	TimeoutCount int `json:"timeout_count"`
}

// TrainedModelDeploymentStatsBuilder holds TrainedModelDeploymentStats struct and provides a builder API.
type TrainedModelDeploymentStatsBuilder struct {
	v *TrainedModelDeploymentStats
}

// NewTrainedModelDeploymentStats provides a builder for the TrainedModelDeploymentStats struct.
func NewTrainedModelDeploymentStatsBuilder() *TrainedModelDeploymentStatsBuilder {
	r := TrainedModelDeploymentStatsBuilder{
		&TrainedModelDeploymentStats{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelDeploymentStats struct
func (rb *TrainedModelDeploymentStatsBuilder) Build() TrainedModelDeploymentStats {
	return *rb.v
}

// AllocationStatus The detailed allocation status for the deployment.

func (rb *TrainedModelDeploymentStatsBuilder) AllocationStatus(allocationstatus *TrainedModelDeploymentAllocationStatusBuilder) *TrainedModelDeploymentStatsBuilder {
	v := allocationstatus.Build()
	rb.v.AllocationStatus = v
	return rb
}

// ErrorCount The sum of `error_count` for all nodes in the deployment.

func (rb *TrainedModelDeploymentStatsBuilder) ErrorCount(errorcount int) *TrainedModelDeploymentStatsBuilder {
	rb.v.ErrorCount = errorcount
	return rb
}

// InferenceCount The sum of `inference_count` for all nodes in the deployment.

func (rb *TrainedModelDeploymentStatsBuilder) InferenceCount(inferencecount int) *TrainedModelDeploymentStatsBuilder {
	rb.v.InferenceCount = inferencecount
	return rb
}

// ModelId The unique identifier for the trained model.

func (rb *TrainedModelDeploymentStatsBuilder) ModelId(modelid Id) *TrainedModelDeploymentStatsBuilder {
	rb.v.ModelId = modelid
	return rb
}

// Nodes The deployent stats for each node that currently has the model allocated.

func (rb *TrainedModelDeploymentStatsBuilder) Nodes(nodes *TrainedModelDeploymentNodesStatsBuilder) *TrainedModelDeploymentStatsBuilder {
	v := nodes.Build()
	rb.v.Nodes = v
	return rb
}

// NumberOfAllocations The number of allocations requested.

func (rb *TrainedModelDeploymentStatsBuilder) NumberOfAllocations(numberofallocations int) *TrainedModelDeploymentStatsBuilder {
	rb.v.NumberOfAllocations = numberofallocations
	return rb
}

// QueueCapacity The number of inference requests that can be queued before new requests are
// rejected.

func (rb *TrainedModelDeploymentStatsBuilder) QueueCapacity(queuecapacity int) *TrainedModelDeploymentStatsBuilder {
	rb.v.QueueCapacity = queuecapacity
	return rb
}

// Reason The reason for the current deployment state. Usually only populated when
// the model is not deployed to a node.

func (rb *TrainedModelDeploymentStatsBuilder) Reason(reason string) *TrainedModelDeploymentStatsBuilder {
	rb.v.Reason = reason
	return rb
}

// RejectedExecutionCount The sum of `rejected_execution_count` for all nodes in the deployment.
// Individual nodes reject an inference request if the inference queue is full.
// The queue size is controlled by the `queue_capacity` setting in the start
// trained model deployment API.

func (rb *TrainedModelDeploymentStatsBuilder) RejectedExecutionCount(rejectedexecutioncount int) *TrainedModelDeploymentStatsBuilder {
	rb.v.RejectedExecutionCount = rejectedexecutioncount
	return rb
}

// StartTime The epoch timestamp when the deployment started.

func (rb *TrainedModelDeploymentStatsBuilder) StartTime(starttime *EpochTimeUnitMillisBuilder) *TrainedModelDeploymentStatsBuilder {
	v := starttime.Build()
	rb.v.StartTime = v
	return rb
}

// State The overall state of the deployment.

func (rb *TrainedModelDeploymentStatsBuilder) State(state deploymentstate.DeploymentState) *TrainedModelDeploymentStatsBuilder {
	rb.v.State = state
	return rb
}

// ThreadsPerAllocation The number of threads used be each allocation during inference.

func (rb *TrainedModelDeploymentStatsBuilder) ThreadsPerAllocation(threadsperallocation int) *TrainedModelDeploymentStatsBuilder {
	rb.v.ThreadsPerAllocation = threadsperallocation
	return rb
}

// TimeoutCount The sum of `timeout_count` for all nodes in the deployment.

func (rb *TrainedModelDeploymentStatsBuilder) TimeoutCount(timeoutcount int) *TrainedModelDeploymentStatsBuilder {
	rb.v.TimeoutCount = timeoutcount
	return rb
}
