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

// NodeInfoTransport type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L342-L346
type NodeInfoTransport struct {
	BoundAddress   []string          `json:"bound_address"`
	Profiles       map[string]string `json:"profiles"`
	PublishAddress string            `json:"publish_address"`
}

// NodeInfoTransportBuilder holds NodeInfoTransport struct and provides a builder API.
type NodeInfoTransportBuilder struct {
	v *NodeInfoTransport
}

// NewNodeInfoTransport provides a builder for the NodeInfoTransport struct.
func NewNodeInfoTransportBuilder() *NodeInfoTransportBuilder {
	r := NodeInfoTransportBuilder{
		&NodeInfoTransport{
			Profiles: make(map[string]string, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoTransport struct
func (rb *NodeInfoTransportBuilder) Build() NodeInfoTransport {
	return *rb.v
}

func (rb *NodeInfoTransportBuilder) BoundAddress(bound_address ...string) *NodeInfoTransportBuilder {
	rb.v.BoundAddress = bound_address
	return rb
}

func (rb *NodeInfoTransportBuilder) Profiles(value map[string]string) *NodeInfoTransportBuilder {
	rb.v.Profiles = value
	return rb
}

func (rb *NodeInfoTransportBuilder) PublishAddress(publishaddress string) *NodeInfoTransportBuilder {
	rb.v.PublishAddress = publishaddress
	return rb
}
