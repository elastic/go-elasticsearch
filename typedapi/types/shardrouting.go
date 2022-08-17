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

// ShardRouting type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/stats/types.ts#L150-L155
type ShardRouting struct {
	Node           string                              `json:"node"`
	Primary        bool                                `json:"primary"`
	RelocatingNode string                              `json:"relocating_node,omitempty"`
	State          shardroutingstate.ShardRoutingState `json:"state"`
}

// ShardRoutingBuilder holds ShardRouting struct and provides a builder API.
type ShardRoutingBuilder struct {
	v *ShardRouting
}

// NewShardRouting provides a builder for the ShardRouting struct.
func NewShardRoutingBuilder() *ShardRoutingBuilder {
	r := ShardRoutingBuilder{
		&ShardRouting{},
	}

	return &r
}

// Build finalize the chain and returns the ShardRouting struct
func (rb *ShardRoutingBuilder) Build() ShardRouting {
	return *rb.v
}

func (rb *ShardRoutingBuilder) Node(node string) *ShardRoutingBuilder {
	rb.v.Node = node
	return rb
}

func (rb *ShardRoutingBuilder) Primary(primary bool) *ShardRoutingBuilder {
	rb.v.Primary = primary
	return rb
}

func (rb *ShardRoutingBuilder) RelocatingNode(relocatingnode string) *ShardRoutingBuilder {
	rb.v.RelocatingNode = relocatingnode
	return rb
}

func (rb *ShardRoutingBuilder) State(state shardroutingstate.ShardRoutingState) *ShardRoutingBuilder {
	rb.v.State = state
	return rb
}
