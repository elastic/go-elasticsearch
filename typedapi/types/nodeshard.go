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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/shardroutingstate"
)

// NodeShard type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/Node.ts#L55-L65
type NodeShard struct {
	AllocationId   map[string]Id                       `json:"allocation_id,omitempty"`
	Index          IndexName                           `json:"index"`
	Node           *NodeName                           `json:"node,omitempty"`
	Primary        bool                                `json:"primary"`
	RecoverySource map[string]Id                       `json:"recovery_source,omitempty"`
	RelocatingNode NodeId                              `json:"relocating_node,omitempty"`
	Shard          int                                 `json:"shard"`
	State          shardroutingstate.ShardRoutingState `json:"state"`
	UnassignedInfo *UnassignedInformation              `json:"unassigned_info,omitempty"`
}

// NodeShardBuilder holds NodeShard struct and provides a builder API.
type NodeShardBuilder struct {
	v *NodeShard
}

// NewNodeShard provides a builder for the NodeShard struct.
func NewNodeShardBuilder() *NodeShardBuilder {
	r := NodeShardBuilder{
		&NodeShard{
			AllocationId:   make(map[string]Id, 0),
			RecoverySource: make(map[string]Id, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeShard struct
func (rb *NodeShardBuilder) Build() NodeShard {
	return *rb.v
}

func (rb *NodeShardBuilder) AllocationId(value map[string]Id) *NodeShardBuilder {
	rb.v.AllocationId = value
	return rb
}

func (rb *NodeShardBuilder) Index(index IndexName) *NodeShardBuilder {
	rb.v.Index = index
	return rb
}

func (rb *NodeShardBuilder) Node(node NodeName) *NodeShardBuilder {
	rb.v.Node = &node
	return rb
}

func (rb *NodeShardBuilder) Primary(primary bool) *NodeShardBuilder {
	rb.v.Primary = primary
	return rb
}

func (rb *NodeShardBuilder) RecoverySource(value map[string]Id) *NodeShardBuilder {
	rb.v.RecoverySource = value
	return rb
}

func (rb *NodeShardBuilder) RelocatingNode(relocatingnode NodeId) *NodeShardBuilder {
	rb.v.RelocatingNode = relocatingnode
	return rb
}

func (rb *NodeShardBuilder) Shard(shard int) *NodeShardBuilder {
	rb.v.Shard = shard
	return rb
}

func (rb *NodeShardBuilder) State(state shardroutingstate.ShardRoutingState) *NodeShardBuilder {
	rb.v.State = state
	return rb
}

func (rb *NodeShardBuilder) UnassignedInfo(unassignedinfo *UnassignedInformationBuilder) *NodeShardBuilder {
	v := unassignedinfo.Build()
	rb.v.UnassignedInfo = &v
	return rb
}
