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

// TrainedModelDeploymentNodesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L127-L154
type TrainedModelDeploymentNodesStats struct {
	// AverageInferenceTimeMs The average time for each inference call to complete on this node.
	AverageInferenceTimeMs DurationValueUnitFloatMillis `json:"average_inference_time_ms"`
	// ErrorCount The number of errors when evaluating the trained model.
	ErrorCount int `json:"error_count"`
	// InferenceCount The total number of inference calls made against this node for this model.
	InferenceCount int `json:"inference_count"`
	// LastAccess The epoch time stamp of the last inference call for the model on this node.
	LastAccess int64 `json:"last_access"`
	// Node Information pertaining to the node.
	Node DiscoveryNode `json:"node"`
	// NumberOfAllocations The number of allocations assigned to this node.
	NumberOfAllocations int `json:"number_of_allocations"`
	// NumberOfPendingRequests The number of inference requests queued to be processed.
	NumberOfPendingRequests int `json:"number_of_pending_requests"`
	// RejectionExecutionCount The number of inference requests that were not processed because the queue
	// was full.
	RejectionExecutionCount int `json:"rejection_execution_count"`
	// RoutingState The current routing state and reason for the current routing state for this
	// allocation.
	RoutingState TrainedModelAssignmentRoutingTable `json:"routing_state"`
	// StartTime The epoch timestamp when the allocation started.
	StartTime EpochTimeUnitMillis `json:"start_time"`
	// ThreadsPerAllocation The number of threads used by each allocation during inference.
	ThreadsPerAllocation int `json:"threads_per_allocation"`
	// TimeoutCount The number of inference requests that timed out before being processed.
	TimeoutCount int `json:"timeout_count"`
}

// TrainedModelDeploymentNodesStatsBuilder holds TrainedModelDeploymentNodesStats struct and provides a builder API.
type TrainedModelDeploymentNodesStatsBuilder struct {
	v *TrainedModelDeploymentNodesStats
}

// NewTrainedModelDeploymentNodesStats provides a builder for the TrainedModelDeploymentNodesStats struct.
func NewTrainedModelDeploymentNodesStatsBuilder() *TrainedModelDeploymentNodesStatsBuilder {
	r := TrainedModelDeploymentNodesStatsBuilder{
		&TrainedModelDeploymentNodesStats{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelDeploymentNodesStats struct
func (rb *TrainedModelDeploymentNodesStatsBuilder) Build() TrainedModelDeploymentNodesStats {
	return *rb.v
}

// AverageInferenceTimeMs The average time for each inference call to complete on this node.

func (rb *TrainedModelDeploymentNodesStatsBuilder) AverageInferenceTimeMs(averageinferencetimems *DurationValueUnitFloatMillisBuilder) *TrainedModelDeploymentNodesStatsBuilder {
	v := averageinferencetimems.Build()
	rb.v.AverageInferenceTimeMs = v
	return rb
}

// ErrorCount The number of errors when evaluating the trained model.

func (rb *TrainedModelDeploymentNodesStatsBuilder) ErrorCount(errorcount int) *TrainedModelDeploymentNodesStatsBuilder {
	rb.v.ErrorCount = errorcount
	return rb
}

// InferenceCount The total number of inference calls made against this node for this model.

func (rb *TrainedModelDeploymentNodesStatsBuilder) InferenceCount(inferencecount int) *TrainedModelDeploymentNodesStatsBuilder {
	rb.v.InferenceCount = inferencecount
	return rb
}

// LastAccess The epoch time stamp of the last inference call for the model on this node.

func (rb *TrainedModelDeploymentNodesStatsBuilder) LastAccess(lastaccess int64) *TrainedModelDeploymentNodesStatsBuilder {
	rb.v.LastAccess = lastaccess
	return rb
}

// Node Information pertaining to the node.

func (rb *TrainedModelDeploymentNodesStatsBuilder) Node(node *DiscoveryNodeBuilder) *TrainedModelDeploymentNodesStatsBuilder {
	v := node.Build()
	rb.v.Node = v
	return rb
}

// NumberOfAllocations The number of allocations assigned to this node.

func (rb *TrainedModelDeploymentNodesStatsBuilder) NumberOfAllocations(numberofallocations int) *TrainedModelDeploymentNodesStatsBuilder {
	rb.v.NumberOfAllocations = numberofallocations
	return rb
}

// NumberOfPendingRequests The number of inference requests queued to be processed.

func (rb *TrainedModelDeploymentNodesStatsBuilder) NumberOfPendingRequests(numberofpendingrequests int) *TrainedModelDeploymentNodesStatsBuilder {
	rb.v.NumberOfPendingRequests = numberofpendingrequests
	return rb
}

// RejectionExecutionCount The number of inference requests that were not processed because the queue
// was full.

func (rb *TrainedModelDeploymentNodesStatsBuilder) RejectionExecutionCount(rejectionexecutioncount int) *TrainedModelDeploymentNodesStatsBuilder {
	rb.v.RejectionExecutionCount = rejectionexecutioncount
	return rb
}

// RoutingState The current routing state and reason for the current routing state for this
// allocation.

func (rb *TrainedModelDeploymentNodesStatsBuilder) RoutingState(routingstate *TrainedModelAssignmentRoutingTableBuilder) *TrainedModelDeploymentNodesStatsBuilder {
	v := routingstate.Build()
	rb.v.RoutingState = v
	return rb
}

// StartTime The epoch timestamp when the allocation started.

func (rb *TrainedModelDeploymentNodesStatsBuilder) StartTime(starttime *EpochTimeUnitMillisBuilder) *TrainedModelDeploymentNodesStatsBuilder {
	v := starttime.Build()
	rb.v.StartTime = v
	return rb
}

// ThreadsPerAllocation The number of threads used by each allocation during inference.

func (rb *TrainedModelDeploymentNodesStatsBuilder) ThreadsPerAllocation(threadsperallocation int) *TrainedModelDeploymentNodesStatsBuilder {
	rb.v.ThreadsPerAllocation = threadsperallocation
	return rb
}

// TimeoutCount The number of inference requests that timed out before being processed.

func (rb *TrainedModelDeploymentNodesStatsBuilder) TimeoutCount(timeoutcount int) *TrainedModelDeploymentNodesStatsBuilder {
	rb.v.TimeoutCount = timeoutcount
	return rb
}
