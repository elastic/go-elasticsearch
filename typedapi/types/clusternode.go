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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

// ClusterNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/security/_types/ClusterNode.ts#L22-L24
type ClusterNode struct {
	Name Name `json:"name"`
}

// ClusterNodeBuilder holds ClusterNode struct and provides a builder API.
type ClusterNodeBuilder struct {
	v *ClusterNode
}

// NewClusterNode provides a builder for the ClusterNode struct.
func NewClusterNodeBuilder() *ClusterNodeBuilder {
	r := ClusterNodeBuilder{
		&ClusterNode{},
	}

	return &r
}

// Build finalize the chain and returns the ClusterNode struct
func (rb *ClusterNodeBuilder) Build() ClusterNode {
	return *rb.v
}

func (rb *ClusterNodeBuilder) Name(name Name) *ClusterNodeBuilder {
	rb.v.Name = name
	return rb
}
