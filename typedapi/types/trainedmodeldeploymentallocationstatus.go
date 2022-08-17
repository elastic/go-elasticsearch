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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentallocationstate"
)

// TrainedModelDeploymentAllocationStatus type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L370-L377
type TrainedModelDeploymentAllocationStatus struct {
	// AllocationCount The current number of nodes where the model is allocated.
	AllocationCount int `json:"allocation_count"`
	// State The detailed allocation state related to the nodes.
	State deploymentallocationstate.DeploymentAllocationState `json:"state"`
	// TargetAllocationCount The desired number of nodes for model allocation.
	TargetAllocationCount int `json:"target_allocation_count"`
}

// TrainedModelDeploymentAllocationStatusBuilder holds TrainedModelDeploymentAllocationStatus struct and provides a builder API.
type TrainedModelDeploymentAllocationStatusBuilder struct {
	v *TrainedModelDeploymentAllocationStatus
}

// NewTrainedModelDeploymentAllocationStatus provides a builder for the TrainedModelDeploymentAllocationStatus struct.
func NewTrainedModelDeploymentAllocationStatusBuilder() *TrainedModelDeploymentAllocationStatusBuilder {
	r := TrainedModelDeploymentAllocationStatusBuilder{
		&TrainedModelDeploymentAllocationStatus{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelDeploymentAllocationStatus struct
func (rb *TrainedModelDeploymentAllocationStatusBuilder) Build() TrainedModelDeploymentAllocationStatus {
	return *rb.v
}

// AllocationCount The current number of nodes where the model is allocated.

func (rb *TrainedModelDeploymentAllocationStatusBuilder) AllocationCount(allocationcount int) *TrainedModelDeploymentAllocationStatusBuilder {
	rb.v.AllocationCount = allocationcount
	return rb
}

// State The detailed allocation state related to the nodes.

func (rb *TrainedModelDeploymentAllocationStatusBuilder) State(state deploymentallocationstate.DeploymentAllocationState) *TrainedModelDeploymentAllocationStatusBuilder {
	rb.v.State = state
	return rb
}

// TargetAllocationCount The desired number of nodes for model allocation.

func (rb *TrainedModelDeploymentAllocationStatusBuilder) TargetAllocationCount(targetallocationcount int) *TrainedModelDeploymentAllocationStatusBuilder {
	rb.v.TargetAllocationCount = targetallocationcount
	return rb
}
