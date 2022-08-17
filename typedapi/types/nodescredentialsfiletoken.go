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

// NodesCredentialsFileToken type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/get_service_credentials/types.ts#L30-L32
type NodesCredentialsFileToken struct {
	Nodes []string `json:"nodes"`
}

// NodesCredentialsFileTokenBuilder holds NodesCredentialsFileToken struct and provides a builder API.
type NodesCredentialsFileTokenBuilder struct {
	v *NodesCredentialsFileToken
}

// NewNodesCredentialsFileToken provides a builder for the NodesCredentialsFileToken struct.
func NewNodesCredentialsFileTokenBuilder() *NodesCredentialsFileTokenBuilder {
	r := NodesCredentialsFileTokenBuilder{
		&NodesCredentialsFileToken{},
	}

	return &r
}

// Build finalize the chain and returns the NodesCredentialsFileToken struct
func (rb *NodesCredentialsFileTokenBuilder) Build() NodesCredentialsFileToken {
	return *rb.v
}

func (rb *NodesCredentialsFileTokenBuilder) Nodes(nodes ...string) *NodesCredentialsFileTokenBuilder {
	rb.v.Nodes = nodes
	return rb
}
