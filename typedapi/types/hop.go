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

// Hop type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/graph/_types/Hop.ts#L23-L27
type Hop struct {
	Connections *Hop               `json:"connections,omitempty"`
	Query       QueryContainer     `json:"query"`
	Vertices    []VertexDefinition `json:"vertices"`
}

// HopBuilder holds Hop struct and provides a builder API.
type HopBuilder struct {
	v *Hop
}

// NewHop provides a builder for the Hop struct.
func NewHopBuilder() *HopBuilder {
	r := HopBuilder{
		&Hop{},
	}

	return &r
}

// Build finalize the chain and returns the Hop struct
func (rb *HopBuilder) Build() Hop {
	return *rb.v
}

func (rb *HopBuilder) Connections(connections *HopBuilder) *HopBuilder {
	v := connections.Build()
	rb.v.Connections = &v
	return rb
}

func (rb *HopBuilder) Query(query *QueryContainerBuilder) *HopBuilder {
	v := query.Build()
	rb.v.Query = v
	return rb
}

func (rb *HopBuilder) Vertices(vertices []VertexDefinitionBuilder) *HopBuilder {
	tmp := make([]VertexDefinition, len(vertices))
	for _, value := range vertices {
		tmp = append(tmp, value.Build())
	}
	rb.v.Vertices = tmp
	return rb
}
