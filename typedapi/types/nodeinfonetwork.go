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

// NodeInfoNetwork type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L320-L323
type NodeInfoNetwork struct {
	PrimaryInterface NodeInfoNetworkInterface `json:"primary_interface"`
	RefreshInterval  int                      `json:"refresh_interval"`
}

// NodeInfoNetworkBuilder holds NodeInfoNetwork struct and provides a builder API.
type NodeInfoNetworkBuilder struct {
	v *NodeInfoNetwork
}

// NewNodeInfoNetwork provides a builder for the NodeInfoNetwork struct.
func NewNodeInfoNetworkBuilder() *NodeInfoNetworkBuilder {
	r := NodeInfoNetworkBuilder{
		&NodeInfoNetwork{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoNetwork struct
func (rb *NodeInfoNetworkBuilder) Build() NodeInfoNetwork {
	return *rb.v
}

func (rb *NodeInfoNetworkBuilder) PrimaryInterface(primaryinterface *NodeInfoNetworkInterfaceBuilder) *NodeInfoNetworkBuilder {
	v := primaryinterface.Build()
	rb.v.PrimaryInterface = v
	return rb
}

func (rb *NodeInfoNetworkBuilder) RefreshInterval(refreshinterval int) *NodeInfoNetworkBuilder {
	rb.v.RefreshInterval = refreshinterval
	return rb
}
