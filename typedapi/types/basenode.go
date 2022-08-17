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

// BaseNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_spec_utils/BaseNode.ts#L25-L32
type BaseNode struct {
	Attributes       map[string]string `json:"attributes"`
	Host             Host              `json:"host"`
	Ip               Ip                `json:"ip"`
	Name             Name              `json:"name"`
	Roles            *NodeRoles        `json:"roles,omitempty"`
	TransportAddress TransportAddress  `json:"transport_address"`
}

// BaseNodeBuilder holds BaseNode struct and provides a builder API.
type BaseNodeBuilder struct {
	v *BaseNode
}

// NewBaseNode provides a builder for the BaseNode struct.
func NewBaseNodeBuilder() *BaseNodeBuilder {
	r := BaseNodeBuilder{
		&BaseNode{
			Attributes: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the BaseNode struct
func (rb *BaseNodeBuilder) Build() BaseNode {
	return *rb.v
}

func (rb *BaseNodeBuilder) Attributes(value map[string]string) *BaseNodeBuilder {
	rb.v.Attributes = value
	return rb
}

func (rb *BaseNodeBuilder) Host(host Host) *BaseNodeBuilder {
	rb.v.Host = host
	return rb
}

func (rb *BaseNodeBuilder) Ip(ip Ip) *BaseNodeBuilder {
	rb.v.Ip = ip
	return rb
}

func (rb *BaseNodeBuilder) Name(name Name) *BaseNodeBuilder {
	rb.v.Name = name
	return rb
}

func (rb *BaseNodeBuilder) Roles(roles *NodeRolesBuilder) *BaseNodeBuilder {
	v := roles.Build()
	rb.v.Roles = &v
	return rb
}

func (rb *BaseNodeBuilder) TransportAddress(transportaddress TransportAddress) *BaseNodeBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
