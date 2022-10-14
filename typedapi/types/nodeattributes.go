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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// NodeAttributes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_types/Node.ts#L41-L57
type NodeAttributes struct {
	// Attributes Lists node attributes.
	Attributes map[string]string `json:"attributes"`
	// EphemeralId The ephemeral ID of the node.
	EphemeralId Id     `json:"ephemeral_id"`
	ExternalId  string `json:"external_id"`
	// Id The unique identifier of the node.
	Id *NodeId `json:"id,omitempty"`
	// Name The unique identifier of the node.
	Name  NodeName   `json:"name"`
	Roles *NodeRoles `json:"roles,omitempty"`
	// TransportAddress The host and port where transport HTTP connections are accepted.
	TransportAddress TransportAddress `json:"transport_address"`
}

// NodeAttributesBuilder holds NodeAttributes struct and provides a builder API.
type NodeAttributesBuilder struct {
	v *NodeAttributes
}

// NewNodeAttributes provides a builder for the NodeAttributes struct.
func NewNodeAttributesBuilder() *NodeAttributesBuilder {
	r := NodeAttributesBuilder{
		&NodeAttributes{
			Attributes: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeAttributes struct
func (rb *NodeAttributesBuilder) Build() NodeAttributes {
	return *rb.v
}

// Attributes Lists node attributes.

func (rb *NodeAttributesBuilder) Attributes(value map[string]string) *NodeAttributesBuilder {
	rb.v.Attributes = value
	return rb
}

// EphemeralId The ephemeral ID of the node.

func (rb *NodeAttributesBuilder) EphemeralId(ephemeralid Id) *NodeAttributesBuilder {
	rb.v.EphemeralId = ephemeralid
	return rb
}

func (rb *NodeAttributesBuilder) ExternalId(externalid string) *NodeAttributesBuilder {
	rb.v.ExternalId = externalid
	return rb
}

// Id The unique identifier of the node.

func (rb *NodeAttributesBuilder) Id(id NodeId) *NodeAttributesBuilder {
	rb.v.Id = &id
	return rb
}

// Name The unique identifier of the node.

func (rb *NodeAttributesBuilder) Name(name NodeName) *NodeAttributesBuilder {
	rb.v.Name = name
	return rb
}

func (rb *NodeAttributesBuilder) Roles(roles *NodeRolesBuilder) *NodeAttributesBuilder {
	v := roles.Build()
	rb.v.Roles = &v
	return rb
}

// TransportAddress The host and port where transport HTTP connections are accepted.

func (rb *NodeAttributesBuilder) TransportAddress(transportaddress TransportAddress) *NodeAttributesBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
