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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/decision"
)

// NodeAllocationExplanation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/allocation_explain/types.ts#L97-L106
type NodeAllocationExplanation struct {
	Deciders         []AllocationDecision `json:"deciders"`
	NodeAttributes   map[string]string    `json:"node_attributes"`
	NodeDecision     decision.Decision    `json:"node_decision"`
	NodeId           Id                   `json:"node_id"`
	NodeName         Name                 `json:"node_name"`
	Store            *AllocationStore     `json:"store,omitempty"`
	TransportAddress TransportAddress     `json:"transport_address"`
	WeightRanking    int                  `json:"weight_ranking"`
}

// NodeAllocationExplanationBuilder holds NodeAllocationExplanation struct and provides a builder API.
type NodeAllocationExplanationBuilder struct {
	v *NodeAllocationExplanation
}

// NewNodeAllocationExplanation provides a builder for the NodeAllocationExplanation struct.
func NewNodeAllocationExplanationBuilder() *NodeAllocationExplanationBuilder {
	r := NodeAllocationExplanationBuilder{
		&NodeAllocationExplanation{
			NodeAttributes: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeAllocationExplanation struct
func (rb *NodeAllocationExplanationBuilder) Build() NodeAllocationExplanation {
	return *rb.v
}

func (rb *NodeAllocationExplanationBuilder) Deciders(deciders []AllocationDecisionBuilder) *NodeAllocationExplanationBuilder {
	tmp := make([]AllocationDecision, len(deciders))
	for _, value := range deciders {
		tmp = append(tmp, value.Build())
	}
	rb.v.Deciders = tmp
	return rb
}

func (rb *NodeAllocationExplanationBuilder) NodeAttributes(value map[string]string) *NodeAllocationExplanationBuilder {
	rb.v.NodeAttributes = value
	return rb
}

func (rb *NodeAllocationExplanationBuilder) NodeDecision(nodedecision decision.Decision) *NodeAllocationExplanationBuilder {
	rb.v.NodeDecision = nodedecision
	return rb
}

func (rb *NodeAllocationExplanationBuilder) NodeId(nodeid Id) *NodeAllocationExplanationBuilder {
	rb.v.NodeId = nodeid
	return rb
}

func (rb *NodeAllocationExplanationBuilder) NodeName(nodename Name) *NodeAllocationExplanationBuilder {
	rb.v.NodeName = nodename
	return rb
}

func (rb *NodeAllocationExplanationBuilder) Store(store *AllocationStoreBuilder) *NodeAllocationExplanationBuilder {
	v := store.Build()
	rb.v.Store = &v
	return rb
}

func (rb *NodeAllocationExplanationBuilder) TransportAddress(transportaddress TransportAddress) *NodeAllocationExplanationBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}

func (rb *NodeAllocationExplanationBuilder) WeightRanking(weightranking int) *NodeAllocationExplanationBuilder {
	rb.v.WeightRanking = weightranking
	return rb
}
