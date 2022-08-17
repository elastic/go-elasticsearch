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

// SerializedClusterState type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/_types/Stats.ts#L95-L98
type SerializedClusterState struct {
	Diffs      *SerializedClusterStateDetail `json:"diffs,omitempty"`
	FullStates *SerializedClusterStateDetail `json:"full_states,omitempty"`
}

// SerializedClusterStateBuilder holds SerializedClusterState struct and provides a builder API.
type SerializedClusterStateBuilder struct {
	v *SerializedClusterState
}

// NewSerializedClusterState provides a builder for the SerializedClusterState struct.
func NewSerializedClusterStateBuilder() *SerializedClusterStateBuilder {
	r := SerializedClusterStateBuilder{
		&SerializedClusterState{},
	}

	return &r
}

// Build finalize the chain and returns the SerializedClusterState struct
func (rb *SerializedClusterStateBuilder) Build() SerializedClusterState {
	return *rb.v
}

func (rb *SerializedClusterStateBuilder) Diffs(diffs *SerializedClusterStateDetailBuilder) *SerializedClusterStateBuilder {
	v := diffs.Build()
	rb.v.Diffs = &v
	return rb
}

func (rb *SerializedClusterStateBuilder) FullStates(fullstates *SerializedClusterStateDetailBuilder) *SerializedClusterStateBuilder {
	v := fullstates.Build()
	rb.v.FullStates = &v
	return rb
}
