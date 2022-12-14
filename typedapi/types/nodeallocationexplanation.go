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
// https://github.com/elastic/elasticsearch-specification/tree/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/decision"
)

// NodeAllocationExplanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/66fc1fdaeee07b44c6d4ddcab3bd6934e3625e33/specification/cluster/allocation_explain/types.ts#L97-L106
type NodeAllocationExplanation struct {
	Deciders         []AllocationDecision `json:"deciders"`
	NodeAttributes   map[string]string    `json:"node_attributes"`
	NodeDecision     decision.Decision    `json:"node_decision"`
	NodeId           string               `json:"node_id"`
	NodeName         string               `json:"node_name"`
	Store            *AllocationStore     `json:"store,omitempty"`
	TransportAddress string               `json:"transport_address"`
	WeightRanking    int                  `json:"weight_ranking"`
}

// NewNodeAllocationExplanation returns a NodeAllocationExplanation.
func NewNodeAllocationExplanation() *NodeAllocationExplanation {
	r := &NodeAllocationExplanation{
		NodeAttributes: make(map[string]string, 0),
	}

	return r
}
