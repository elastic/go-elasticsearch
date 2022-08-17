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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentassignmentstate"
)

// TrainedModelAssignment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/TrainedModel.ts#L379-L393
type TrainedModelAssignment struct {
	// AssignmentState The overall assignment state.
	AssignmentState deploymentassignmentstate.DeploymentAssignmentState `json:"assignment_state"`
	// RoutingTable The allocation state for each node.
	RoutingTable map[string]TrainedModelAssignmentRoutingTable `json:"routing_table"`
	// StartTime The timestamp when the deployment started.
	StartTime      DateTime                             `json:"start_time"`
	TaskParameters TrainedModelAssignmentTaskParameters `json:"task_parameters"`
}

// TrainedModelAssignmentBuilder holds TrainedModelAssignment struct and provides a builder API.
type TrainedModelAssignmentBuilder struct {
	v *TrainedModelAssignment
}

// NewTrainedModelAssignment provides a builder for the TrainedModelAssignment struct.
func NewTrainedModelAssignmentBuilder() *TrainedModelAssignmentBuilder {
	r := TrainedModelAssignmentBuilder{
		&TrainedModelAssignment{
			RoutingTable: make(map[string]TrainedModelAssignmentRoutingTable, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelAssignment struct
func (rb *TrainedModelAssignmentBuilder) Build() TrainedModelAssignment {
	return *rb.v
}

// AssignmentState The overall assignment state.

func (rb *TrainedModelAssignmentBuilder) AssignmentState(assignmentstate deploymentassignmentstate.DeploymentAssignmentState) *TrainedModelAssignmentBuilder {
	rb.v.AssignmentState = assignmentstate
	return rb
}

// RoutingTable The allocation state for each node.

func (rb *TrainedModelAssignmentBuilder) RoutingTable(values map[string]*TrainedModelAssignmentRoutingTableBuilder) *TrainedModelAssignmentBuilder {
	tmp := make(map[string]TrainedModelAssignmentRoutingTable, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.RoutingTable = tmp
	return rb
}

// StartTime The timestamp when the deployment started.

func (rb *TrainedModelAssignmentBuilder) StartTime(starttime *DateTimeBuilder) *TrainedModelAssignmentBuilder {
	v := starttime.Build()
	rb.v.StartTime = v
	return rb
}

func (rb *TrainedModelAssignmentBuilder) TaskParameters(taskparameters *TrainedModelAssignmentTaskParametersBuilder) *TrainedModelAssignmentBuilder {
	v := taskparameters.Build()
	rb.v.TaskParameters = v
	return rb
}
