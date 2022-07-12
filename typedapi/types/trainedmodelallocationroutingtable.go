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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/routingstate"
)

// TrainedModelAllocationRoutingTable type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/ml/_types/TrainedModel.ts#L326-L336
type TrainedModelAllocationRoutingTable struct {
	// Reason The reason for the current state. It is usually populated only when the
	// `routing_state` is `failed`.
	Reason string `json:"reason"`
	// RoutingState The current routing state.
	RoutingState routingstate.RoutingState `json:"routing_state"`
}

// TrainedModelAllocationRoutingTableBuilder holds TrainedModelAllocationRoutingTable struct and provides a builder API.
type TrainedModelAllocationRoutingTableBuilder struct {
	v *TrainedModelAllocationRoutingTable
}

// NewTrainedModelAllocationRoutingTable provides a builder for the TrainedModelAllocationRoutingTable struct.
func NewTrainedModelAllocationRoutingTableBuilder() *TrainedModelAllocationRoutingTableBuilder {
	r := TrainedModelAllocationRoutingTableBuilder{
		&TrainedModelAllocationRoutingTable{},
	}

	return &r
}

// Build finalize the chain and returns the TrainedModelAllocationRoutingTable struct
func (rb *TrainedModelAllocationRoutingTableBuilder) Build() TrainedModelAllocationRoutingTable {
	return *rb.v
}

// Reason The reason for the current state. It is usually populated only when the
// `routing_state` is `failed`.

func (rb *TrainedModelAllocationRoutingTableBuilder) Reason(reason string) *TrainedModelAllocationRoutingTableBuilder {
	rb.v.Reason = reason
	return rb
}

// RoutingState The current routing state.

func (rb *TrainedModelAllocationRoutingTableBuilder) RoutingState(routingstate routingstate.RoutingState) *TrainedModelAllocationRoutingTableBuilder {
	rb.v.RoutingState = routingstate
	return rb
}
