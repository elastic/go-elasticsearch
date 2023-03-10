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

// TrainedModelDeploymentNodesStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/TrainedModel.ts#L128-L155
type TrainedModelDeploymentNodesStats struct {
	// AverageInferenceTimeMs The average time for each inference call to complete on this node.
	AverageInferenceTimeMs Float64 `json:"average_inference_time_ms"`
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
	StartTime int64 `json:"start_time"`
	// ThreadsPerAllocation The number of threads used by each allocation during inference.
	ThreadsPerAllocation int `json:"threads_per_allocation"`
	// TimeoutCount The number of inference requests that timed out before being processed.
	TimeoutCount int `json:"timeout_count"`
}

// NewTrainedModelDeploymentNodesStats returns a TrainedModelDeploymentNodesStats.
func NewTrainedModelDeploymentNodesStats() *TrainedModelDeploymentNodesStats {
	r := &TrainedModelDeploymentNodesStats{}

	return r
}
