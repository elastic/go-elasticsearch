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

// TransformStats type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/get_transform_stats/types.ts#L31-L38
type TransformStats struct {
	Checkpointing Checkpointing         `json:"checkpointing"`
	Id            Id                    `json:"id"`
	Node          *NodeAttributes       `json:"node,omitempty"`
	Reason        *string               `json:"reason,omitempty"`
	State         string                `json:"state"`
	Stats         TransformIndexerStats `json:"stats"`
}

// TransformStatsBuilder holds TransformStats struct and provides a builder API.
type TransformStatsBuilder struct {
	v *TransformStats
}

// NewTransformStats provides a builder for the TransformStats struct.
func NewTransformStatsBuilder() *TransformStatsBuilder {
	r := TransformStatsBuilder{
		&TransformStats{},
	}

	return &r
}

// Build finalize the chain and returns the TransformStats struct
func (rb *TransformStatsBuilder) Build() TransformStats {
	return *rb.v
}

func (rb *TransformStatsBuilder) Checkpointing(checkpointing *CheckpointingBuilder) *TransformStatsBuilder {
	v := checkpointing.Build()
	rb.v.Checkpointing = v
	return rb
}

func (rb *TransformStatsBuilder) Id(id Id) *TransformStatsBuilder {
	rb.v.Id = id
	return rb
}

func (rb *TransformStatsBuilder) Node(node *NodeAttributesBuilder) *TransformStatsBuilder {
	v := node.Build()
	rb.v.Node = &v
	return rb
}

func (rb *TransformStatsBuilder) Reason(reason string) *TransformStatsBuilder {
	rb.v.Reason = &reason
	return rb
}

func (rb *TransformStatsBuilder) State(state string) *TransformStatsBuilder {
	rb.v.State = state
	return rb
}

func (rb *TransformStatsBuilder) Stats(stats *TransformIndexerStatsBuilder) *TransformStatsBuilder {
	v := stats.Build()
	rb.v.Stats = v
	return rb
}
