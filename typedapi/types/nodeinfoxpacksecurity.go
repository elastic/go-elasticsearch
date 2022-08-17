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

// NodeInfoXpackSecurity type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L234-L239
type NodeInfoXpackSecurity struct {
	Authc     *NodeInfoXpackSecurityAuthc `json:"authc,omitempty"`
	Enabled   string                      `json:"enabled"`
	Http      NodeInfoXpackSecuritySsl    `json:"http"`
	Transport *NodeInfoXpackSecuritySsl   `json:"transport,omitempty"`
}

// NodeInfoXpackSecurityBuilder holds NodeInfoXpackSecurity struct and provides a builder API.
type NodeInfoXpackSecurityBuilder struct {
	v *NodeInfoXpackSecurity
}

// NewNodeInfoXpackSecurity provides a builder for the NodeInfoXpackSecurity struct.
func NewNodeInfoXpackSecurityBuilder() *NodeInfoXpackSecurityBuilder {
	r := NodeInfoXpackSecurityBuilder{
		&NodeInfoXpackSecurity{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoXpackSecurity struct
func (rb *NodeInfoXpackSecurityBuilder) Build() NodeInfoXpackSecurity {
	return *rb.v
}

func (rb *NodeInfoXpackSecurityBuilder) Authc(authc *NodeInfoXpackSecurityAuthcBuilder) *NodeInfoXpackSecurityBuilder {
	v := authc.Build()
	rb.v.Authc = &v
	return rb
}

func (rb *NodeInfoXpackSecurityBuilder) Enabled(enabled string) *NodeInfoXpackSecurityBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *NodeInfoXpackSecurityBuilder) Http(http *NodeInfoXpackSecuritySslBuilder) *NodeInfoXpackSecurityBuilder {
	v := http.Build()
	rb.v.Http = v
	return rb
}

func (rb *NodeInfoXpackSecurityBuilder) Transport(transport *NodeInfoXpackSecuritySslBuilder) *NodeInfoXpackSecurityBuilder {
	v := transport.Build()
	rb.v.Transport = &v
	return rb
}
