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

// NodeAttributesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/nodeattrs/types.ts#L20-L55
type NodeAttributesRecord struct {
	// Attr attribute description
	Attr *string `json:"attr,omitempty"`
	// Host host name
	Host *string `json:"host,omitempty"`
	// Id unique node id
	Id *string `json:"id,omitempty"`
	// Ip ip address
	Ip *string `json:"ip,omitempty"`
	// Node node name
	Node *string `json:"node,omitempty"`
	// Pid process id
	Pid *string `json:"pid,omitempty"`
	// Port bound transport port
	Port *string `json:"port,omitempty"`
	// Value attribute value
	Value *string `json:"value,omitempty"`
}

// NodeAttributesRecordBuilder holds NodeAttributesRecord struct and provides a builder API.
type NodeAttributesRecordBuilder struct {
	v *NodeAttributesRecord
}

// NewNodeAttributesRecord provides a builder for the NodeAttributesRecord struct.
func NewNodeAttributesRecordBuilder() *NodeAttributesRecordBuilder {
	r := NodeAttributesRecordBuilder{
		&NodeAttributesRecord{},
	}

	return &r
}

// Build finalize the chain and returns the NodeAttributesRecord struct
func (rb *NodeAttributesRecordBuilder) Build() NodeAttributesRecord {
	return *rb.v
}

// Attr attribute description

func (rb *NodeAttributesRecordBuilder) Attr(attr string) *NodeAttributesRecordBuilder {
	rb.v.Attr = &attr
	return rb
}

// Host host name

func (rb *NodeAttributesRecordBuilder) Host(host string) *NodeAttributesRecordBuilder {
	rb.v.Host = &host
	return rb
}

// Id unique node id

func (rb *NodeAttributesRecordBuilder) Id(id string) *NodeAttributesRecordBuilder {
	rb.v.Id = &id
	return rb
}

// Ip ip address

func (rb *NodeAttributesRecordBuilder) Ip(ip string) *NodeAttributesRecordBuilder {
	rb.v.Ip = &ip
	return rb
}

// Node node name

func (rb *NodeAttributesRecordBuilder) Node(node string) *NodeAttributesRecordBuilder {
	rb.v.Node = &node
	return rb
}

// Pid process id

func (rb *NodeAttributesRecordBuilder) Pid(pid string) *NodeAttributesRecordBuilder {
	rb.v.Pid = &pid
	return rb
}

// Port bound transport port

func (rb *NodeAttributesRecordBuilder) Port(port string) *NodeAttributesRecordBuilder {
	rb.v.Port = &port
	return rb
}

// Value attribute value

func (rb *NodeAttributesRecordBuilder) Value(value string) *NodeAttributesRecordBuilder {
	rb.v.Value = &value
	return rb
}
