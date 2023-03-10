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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/routingstate"
)

// TrainedModelAssignmentRoutingTable type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/TrainedModel.ts#L358-L376
type TrainedModelAssignmentRoutingTable struct {
	// CurrentAllocations Current number of allocations.
	CurrentAllocations int `json:"current_allocations"`
	// Reason The reason for the current state. It is usually populated only when the
	// `routing_state` is `failed`.
	Reason string `json:"reason"`
	// RoutingState The current routing state.
	RoutingState routingstate.RoutingState `json:"routing_state"`
	// TargetAllocations Target number of allocations.
	TargetAllocations int `json:"target_allocations"`
}

// NewTrainedModelAssignmentRoutingTable returns a TrainedModelAssignmentRoutingTable.
func NewTrainedModelAssignmentRoutingTable() *TrainedModelAssignmentRoutingTable {
	r := &TrainedModelAssignmentRoutingTable{}

	return r
}
