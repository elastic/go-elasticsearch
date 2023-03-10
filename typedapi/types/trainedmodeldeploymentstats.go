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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentstate"
)

// TrainedModelDeploymentStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/TrainedModel.ts#L62-L97
type TrainedModelDeploymentStats struct {
	// AllocationStatus The detailed allocation status for the deployment.
	AllocationStatus TrainedModelDeploymentAllocationStatus `json:"allocation_status"`
	CacheSize        ByteSize                               `json:"cache_size,omitempty"`
	// ErrorCount The sum of `error_count` for all nodes in the deployment.
	ErrorCount int `json:"error_count"`
	// InferenceCount The sum of `inference_count` for all nodes in the deployment.
	InferenceCount int `json:"inference_count"`
	// ModelId The unique identifier for the trained model.
	ModelId string `json:"model_id"`
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
	StartTime int64 `json:"start_time"`
	// State The overall state of the deployment.
	State deploymentstate.DeploymentState `json:"state"`
	// ThreadsPerAllocation The number of threads used be each allocation during inference.
	ThreadsPerAllocation int `json:"threads_per_allocation"`
	// TimeoutCount The sum of `timeout_count` for all nodes in the deployment.
	TimeoutCount int `json:"timeout_count"`
}

// NewTrainedModelDeploymentStats returns a TrainedModelDeploymentStats.
func NewTrainedModelDeploymentStats() *TrainedModelDeploymentStats {
	r := &TrainedModelDeploymentStats{}

	return r
}
