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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/deploymentassignmentstate"

	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// TrainedModelAssignment type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/_types/TrainedModel.ts#L387-L402
type TrainedModelAssignment struct {
	// AssignmentState The overall assignment state.
	AssignmentState        deploymentassignmentstate.DeploymentAssignmentState `json:"assignment_state"`
	MaxAssignedAllocations *int                                                `json:"max_assigned_allocations,omitempty"`
	// RoutingTable The allocation state for each node.
	RoutingTable map[string]TrainedModelAssignmentRoutingTable `json:"routing_table"`
	// StartTime The timestamp when the deployment started.
	StartTime      DateTime                             `json:"start_time"`
	TaskParameters TrainedModelAssignmentTaskParameters `json:"task_parameters"`
}

func (s *TrainedModelAssignment) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "assignment_state":
			if err := dec.Decode(&s.AssignmentState); err != nil {
				return err
			}

		case "max_assigned_allocations":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxAssignedAllocations = &value
			case float64:
				f := int(v)
				s.MaxAssignedAllocations = &f
			}

		case "routing_table":
			if s.RoutingTable == nil {
				s.RoutingTable = make(map[string]TrainedModelAssignmentRoutingTable, 0)
			}
			if err := dec.Decode(&s.RoutingTable); err != nil {
				return err
			}

		case "start_time":
			if err := dec.Decode(&s.StartTime); err != nil {
				return err
			}

		case "task_parameters":
			if err := dec.Decode(&s.TaskParameters); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewTrainedModelAssignment returns a TrainedModelAssignment.
func NewTrainedModelAssignment() *TrainedModelAssignment {
	r := &TrainedModelAssignment{
		RoutingTable: make(map[string]TrainedModelAssignmentRoutingTable, 0),
	}

	return r
}
