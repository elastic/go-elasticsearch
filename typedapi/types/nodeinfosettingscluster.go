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

// NodeInfoSettingsCluster type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L131-L138
type NodeInfoSettingsCluster struct {
	DeprecationIndexing *DeprecationIndexing            `json:"deprecation_indexing,omitempty"`
	Election            NodeInfoSettingsClusterElection `json:"election"`
	InitialMasterNodes  *string                         `json:"initial_master_nodes,omitempty"`
	Name                Name                            `json:"name"`
	Routing             *IndexRouting                   `json:"routing,omitempty"`
}

// NodeInfoSettingsClusterBuilder holds NodeInfoSettingsCluster struct and provides a builder API.
type NodeInfoSettingsClusterBuilder struct {
	v *NodeInfoSettingsCluster
}

// NewNodeInfoSettingsCluster provides a builder for the NodeInfoSettingsCluster struct.
func NewNodeInfoSettingsClusterBuilder() *NodeInfoSettingsClusterBuilder {
	r := NodeInfoSettingsClusterBuilder{
		&NodeInfoSettingsCluster{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsCluster struct
func (rb *NodeInfoSettingsClusterBuilder) Build() NodeInfoSettingsCluster {
	return *rb.v
}

func (rb *NodeInfoSettingsClusterBuilder) DeprecationIndexing(deprecationindexing *DeprecationIndexingBuilder) *NodeInfoSettingsClusterBuilder {
	v := deprecationindexing.Build()
	rb.v.DeprecationIndexing = &v
	return rb
}

func (rb *NodeInfoSettingsClusterBuilder) Election(election *NodeInfoSettingsClusterElectionBuilder) *NodeInfoSettingsClusterBuilder {
	v := election.Build()
	rb.v.Election = v
	return rb
}

func (rb *NodeInfoSettingsClusterBuilder) InitialMasterNodes(initialmasternodes string) *NodeInfoSettingsClusterBuilder {
	rb.v.InitialMasterNodes = &initialmasternodes
	return rb
}

func (rb *NodeInfoSettingsClusterBuilder) Name(name Name) *NodeInfoSettingsClusterBuilder {
	rb.v.Name = name
	return rb
}

func (rb *NodeInfoSettingsClusterBuilder) Routing(routing *IndexRoutingBuilder) *NodeInfoSettingsClusterBuilder {
	v := routing.Build()
	rb.v.Routing = &v
	return rb
}
