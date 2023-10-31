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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package allocationexplain

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/decision"
)

// Response holds the response body struct for the package allocationexplain
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/cluster/allocation_explain/ClusterAllocationExplainResponse.ts#L32-L64

type Response struct {
	AllocateExplanation          *string                           `json:"allocate_explanation,omitempty"`
	AllocationDelay              types.Duration                    `json:"allocation_delay,omitempty"`
	AllocationDelayInMillis      *int64                            `json:"allocation_delay_in_millis,omitempty"`
	CanAllocate                  *decision.Decision                `json:"can_allocate,omitempty"`
	CanMoveToOtherNode           *decision.Decision                `json:"can_move_to_other_node,omitempty"`
	CanRebalanceCluster          *decision.Decision                `json:"can_rebalance_cluster,omitempty"`
	CanRebalanceClusterDecisions []types.AllocationDecision        `json:"can_rebalance_cluster_decisions,omitempty"`
	CanRebalanceToOtherNode      *decision.Decision                `json:"can_rebalance_to_other_node,omitempty"`
	CanRemainDecisions           []types.AllocationDecision        `json:"can_remain_decisions,omitempty"`
	CanRemainOnCurrentNode       *decision.Decision                `json:"can_remain_on_current_node,omitempty"`
	ClusterInfo                  *types.ClusterInfo                `json:"cluster_info,omitempty"`
	ConfiguredDelay              types.Duration                    `json:"configured_delay,omitempty"`
	ConfiguredDelayInMillis      *int64                            `json:"configured_delay_in_millis,omitempty"`
	CurrentNode                  *types.CurrentNode                `json:"current_node,omitempty"`
	CurrentState                 string                            `json:"current_state"`
	Index                        string                            `json:"index"`
	MoveExplanation              *string                           `json:"move_explanation,omitempty"`
	NodeAllocationDecisions      []types.NodeAllocationExplanation `json:"node_allocation_decisions,omitempty"`
	Note                         *string                           `json:"note,omitempty"`
	Primary                      bool                              `json:"primary"`
	RebalanceExplanation         *string                           `json:"rebalance_explanation,omitempty"`
	RemainingDelay               types.Duration                    `json:"remaining_delay,omitempty"`
	RemainingDelayInMillis       *int64                            `json:"remaining_delay_in_millis,omitempty"`
	Shard                        int                               `json:"shard"`
	UnassignedInfo               *types.UnassignedInformation      `json:"unassigned_info,omitempty"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
