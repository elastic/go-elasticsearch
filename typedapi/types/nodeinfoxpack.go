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

// NodeInfoXpack type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L228-L232
type NodeInfoXpack struct {
	License      *NodeInfoXpackLicense  `json:"license,omitempty"`
	Notification map[string]interface{} `json:"notification,omitempty"`
	Security     NodeInfoXpackSecurity  `json:"security"`
}

// NodeInfoXpackBuilder holds NodeInfoXpack struct and provides a builder API.
type NodeInfoXpackBuilder struct {
	v *NodeInfoXpack
}

// NewNodeInfoXpack provides a builder for the NodeInfoXpack struct.
func NewNodeInfoXpackBuilder() *NodeInfoXpackBuilder {
	r := NodeInfoXpackBuilder{
		&NodeInfoXpack{
			Notification: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoXpack struct
func (rb *NodeInfoXpackBuilder) Build() NodeInfoXpack {
	return *rb.v
}

func (rb *NodeInfoXpackBuilder) License(license *NodeInfoXpackLicenseBuilder) *NodeInfoXpackBuilder {
	v := license.Build()
	rb.v.License = &v
	return rb
}

func (rb *NodeInfoXpackBuilder) Notification(value map[string]interface{}) *NodeInfoXpackBuilder {
	rb.v.Notification = value
	return rb
}

func (rb *NodeInfoXpackBuilder) Security(security *NodeInfoXpackSecurityBuilder) *NodeInfoXpackBuilder {
	v := security.Build()
	rb.v.Security = v
	return rb
}
