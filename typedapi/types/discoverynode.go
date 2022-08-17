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

// DiscoveryNode type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DiscoveryNode.ts#L24-L30
type DiscoveryNode struct {
	Attributes       map[string]string `json:"attributes"`
	EphemeralId      Id                `json:"ephemeral_id"`
	Id               Id                `json:"id"`
	Name             Name              `json:"name"`
	TransportAddress TransportAddress  `json:"transport_address"`
}

// DiscoveryNodeBuilder holds DiscoveryNode struct and provides a builder API.
type DiscoveryNodeBuilder struct {
	v *DiscoveryNode
}

// NewDiscoveryNode provides a builder for the DiscoveryNode struct.
func NewDiscoveryNodeBuilder() *DiscoveryNodeBuilder {
	r := DiscoveryNodeBuilder{
		&DiscoveryNode{
			Attributes: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DiscoveryNode struct
func (rb *DiscoveryNodeBuilder) Build() DiscoveryNode {
	return *rb.v
}

func (rb *DiscoveryNodeBuilder) Attributes(value map[string]string) *DiscoveryNodeBuilder {
	rb.v.Attributes = value
	return rb
}

func (rb *DiscoveryNodeBuilder) EphemeralId(ephemeralid Id) *DiscoveryNodeBuilder {
	rb.v.EphemeralId = ephemeralid
	return rb
}

func (rb *DiscoveryNodeBuilder) Id(id Id) *DiscoveryNodeBuilder {
	rb.v.Id = id
	return rb
}

func (rb *DiscoveryNodeBuilder) Name(name Name) *DiscoveryNodeBuilder {
	rb.v.Name = name
	return rb
}

func (rb *DiscoveryNodeBuilder) TransportAddress(transportaddress TransportAddress) *DiscoveryNodeBuilder {
	rb.v.TransportAddress = transportaddress
	return rb
}
