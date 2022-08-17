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

// NodeInfoSettingsHttp type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L181-L186
type NodeInfoSettingsHttp struct {
	Compression string                   `json:"compression,omitempty"`
	Port        string                   `json:"port,omitempty"`
	Type        NodeInfoSettingsHttpType `json:"type"`
	TypeDefault *string                  `json:"type.default,omitempty"`
}

// NodeInfoSettingsHttpBuilder holds NodeInfoSettingsHttp struct and provides a builder API.
type NodeInfoSettingsHttpBuilder struct {
	v *NodeInfoSettingsHttp
}

// NewNodeInfoSettingsHttp provides a builder for the NodeInfoSettingsHttp struct.
func NewNodeInfoSettingsHttpBuilder() *NodeInfoSettingsHttpBuilder {
	r := NodeInfoSettingsHttpBuilder{
		&NodeInfoSettingsHttp{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsHttp struct
func (rb *NodeInfoSettingsHttpBuilder) Build() NodeInfoSettingsHttp {
	return *rb.v
}

func (rb *NodeInfoSettingsHttpBuilder) Compression(arg string) *NodeInfoSettingsHttpBuilder {
	rb.v.Compression = arg
	return rb
}

func (rb *NodeInfoSettingsHttpBuilder) Port(arg string) *NodeInfoSettingsHttpBuilder {
	rb.v.Port = arg
	return rb
}

func (rb *NodeInfoSettingsHttpBuilder) Type_(type_ *NodeInfoSettingsHttpTypeBuilder) *NodeInfoSettingsHttpBuilder {
	v := type_.Build()
	rb.v.Type = v
	return rb
}

func (rb *NodeInfoSettingsHttpBuilder) TypeDefault(typedefault string) *NodeInfoSettingsHttpBuilder {
	rb.v.TypeDefault = &typedefault
	return rb
}
