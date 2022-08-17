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

// NodeInfoSettingsTransportFeatures type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/nodes/info/types.ts#L208-L210
type NodeInfoSettingsTransportFeatures struct {
	XPack string `json:"x-pack"`
}

// NodeInfoSettingsTransportFeaturesBuilder holds NodeInfoSettingsTransportFeatures struct and provides a builder API.
type NodeInfoSettingsTransportFeaturesBuilder struct {
	v *NodeInfoSettingsTransportFeatures
}

// NewNodeInfoSettingsTransportFeatures provides a builder for the NodeInfoSettingsTransportFeatures struct.
func NewNodeInfoSettingsTransportFeaturesBuilder() *NodeInfoSettingsTransportFeaturesBuilder {
	r := NodeInfoSettingsTransportFeaturesBuilder{
		&NodeInfoSettingsTransportFeatures{},
	}

	return &r
}

// Build finalize the chain and returns the NodeInfoSettingsTransportFeatures struct
func (rb *NodeInfoSettingsTransportFeaturesBuilder) Build() NodeInfoSettingsTransportFeatures {
	return *rb.v
}

func (rb *NodeInfoSettingsTransportFeaturesBuilder) XPack(xpack string) *NodeInfoSettingsTransportFeaturesBuilder {
	rb.v.XPack = xpack
	return rb
}
