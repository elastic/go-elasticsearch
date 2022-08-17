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

// ShardSegmentRouting type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/segments/types.ts#L41-L45
type ShardSegmentRouting struct {
	Node    string `json:"node"`
	Primary bool   `json:"primary"`
	State   string `json:"state"`
}

// ShardSegmentRoutingBuilder holds ShardSegmentRouting struct and provides a builder API.
type ShardSegmentRoutingBuilder struct {
	v *ShardSegmentRouting
}

// NewShardSegmentRouting provides a builder for the ShardSegmentRouting struct.
func NewShardSegmentRoutingBuilder() *ShardSegmentRoutingBuilder {
	r := ShardSegmentRoutingBuilder{
		&ShardSegmentRouting{},
	}

	return &r
}

// Build finalize the chain and returns the ShardSegmentRouting struct
func (rb *ShardSegmentRoutingBuilder) Build() ShardSegmentRouting {
	return *rb.v
}

func (rb *ShardSegmentRoutingBuilder) Node(node string) *ShardSegmentRoutingBuilder {
	rb.v.Node = node
	return rb
}

func (rb *ShardSegmentRoutingBuilder) Primary(primary bool) *ShardSegmentRoutingBuilder {
	rb.v.Primary = primary
	return rb
}

func (rb *ShardSegmentRoutingBuilder) State(state string) *ShardSegmentRoutingBuilder {
	rb.v.State = state
	return rb
}
