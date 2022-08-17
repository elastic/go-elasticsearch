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

// NodeUsage type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/usage/types.ts#L25-L30
type NodeUsage struct {
	Aggregations map[string]interface{} `json:"aggregations"`
	RestActions  map[string]int         `json:"rest_actions"`
	Since        EpochTimeUnitMillis    `json:"since"`
	Timestamp    EpochTimeUnitMillis    `json:"timestamp"`
}

// NodeUsageBuilder holds NodeUsage struct and provides a builder API.
type NodeUsageBuilder struct {
	v *NodeUsage
}

// NewNodeUsage provides a builder for the NodeUsage struct.
func NewNodeUsageBuilder() *NodeUsageBuilder {
	r := NodeUsageBuilder{
		&NodeUsage{
			Aggregations: make(map[string]interface{}, 0),
			RestActions:  make(map[string]int, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeUsage struct
func (rb *NodeUsageBuilder) Build() NodeUsage {
	return *rb.v
}

func (rb *NodeUsageBuilder) Aggregations(value map[string]interface{}) *NodeUsageBuilder {
	rb.v.Aggregations = value
	return rb
}

func (rb *NodeUsageBuilder) RestActions(value map[string]int) *NodeUsageBuilder {
	rb.v.RestActions = value
	return rb
}

func (rb *NodeUsageBuilder) Since(since *EpochTimeUnitMillisBuilder) *NodeUsageBuilder {
	v := since.Build()
	rb.v.Since = v
	return rb
}

func (rb *NodeUsageBuilder) Timestamp(timestamp *EpochTimeUnitMillisBuilder) *NodeUsageBuilder {
	v := timestamp.Build()
	rb.v.Timestamp = v
	return rb
}
