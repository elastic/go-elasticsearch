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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentallocationstate"
)

// TrainedModelAllocation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/ml/_types/TrainedModel.ts#L341-L355
type TrainedModelAllocation struct {
	// AllocationState The overall allocation state.
	AllocationState deploymentallocationstate.DeploymentAllocationState `json:"allocation_state"`
	// RoutingTable The allocation state for each node.
	RoutingTable map[string]TrainedModelAllocationRoutingTable `json:"routing_table"`
	// StartTime The timestamp when the deployment started.
	StartTime      DateString                           `json:"start_time"`
	TaskParameters TrainedModelAllocationTaskParameters `json:"task_parameters"`
}

// TrainedModelAllocationBuilder holds TrainedModelAllocation struct and provides a builder API.
type TrainedModelAllocationBuilder struct {
	v *TrainedModelAllocation
}

// NewTrainedModelAllocation provides a builder for the TrainedModelAllocation struct.
func NewTrainedModelAllocationBuilder() *TrainedModelAllocationBuilder {
	r := TrainedModelAllocationBuilder{
		&TrainedModelAllocation{
			RoutingTable: make(map[string]TrainedModelAllocationRoutingTable, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelAllocation struct
func (rb *TrainedModelAllocationBuilder) Build() TrainedModelAllocation {
	return *rb.v
}

// AllocationState The overall allocation state.

func (rb *TrainedModelAllocationBuilder) AllocationState(allocationstate deploymentallocationstate.DeploymentAllocationState) *TrainedModelAllocationBuilder {
	rb.v.AllocationState = allocationstate
	return rb
}

// RoutingTable The allocation state for each node.

func (rb *TrainedModelAllocationBuilder) RoutingTable(values map[string]*TrainedModelAllocationRoutingTableBuilder) *TrainedModelAllocationBuilder {
	tmp := make(map[string]TrainedModelAllocationRoutingTable, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.RoutingTable = tmp
	return rb
}

// StartTime The timestamp when the deployment started.

func (rb *TrainedModelAllocationBuilder) StartTime(starttime DateString) *TrainedModelAllocationBuilder {
	rb.v.StartTime = starttime
	return rb
}

func (rb *TrainedModelAllocationBuilder) TaskParameters(taskparameters *TrainedModelAllocationTaskParametersBuilder) *TrainedModelAllocationBuilder {
	v := taskparameters.Build()
	rb.v.TaskParameters = v
	return rb
}
