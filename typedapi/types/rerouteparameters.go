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

// RerouteParameters type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/reroute/types.ts#L98-L105
type RerouteParameters struct {
	AllowPrimary bool      `json:"allow_primary"`
	FromNode     *NodeName `json:"from_node,omitempty"`
	Index        IndexName `json:"index"`
	Node         NodeName  `json:"node"`
	Shard        int       `json:"shard"`
	ToNode       *NodeName `json:"to_node,omitempty"`
}

// RerouteParametersBuilder holds RerouteParameters struct and provides a builder API.
type RerouteParametersBuilder struct {
	v *RerouteParameters
}

// NewRerouteParameters provides a builder for the RerouteParameters struct.
func NewRerouteParametersBuilder() *RerouteParametersBuilder {
	r := RerouteParametersBuilder{
		&RerouteParameters{},
	}

	return &r
}

// Build finalize the chain and returns the RerouteParameters struct
func (rb *RerouteParametersBuilder) Build() RerouteParameters {
	return *rb.v
}

func (rb *RerouteParametersBuilder) AllowPrimary(allowprimary bool) *RerouteParametersBuilder {
	rb.v.AllowPrimary = allowprimary
	return rb
}

func (rb *RerouteParametersBuilder) FromNode(fromnode NodeName) *RerouteParametersBuilder {
	rb.v.FromNode = &fromnode
	return rb
}

func (rb *RerouteParametersBuilder) Index(index IndexName) *RerouteParametersBuilder {
	rb.v.Index = index
	return rb
}

func (rb *RerouteParametersBuilder) Node(node NodeName) *RerouteParametersBuilder {
	rb.v.Node = node
	return rb
}

func (rb *RerouteParametersBuilder) Shard(shard int) *RerouteParametersBuilder {
	rb.v.Shard = shard
	return rb
}

func (rb *RerouteParametersBuilder) ToNode(tonode NodeName) *RerouteParametersBuilder {
	rb.v.ToNode = &tonode
	return rb
}
