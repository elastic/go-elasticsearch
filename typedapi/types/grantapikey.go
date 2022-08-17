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

// GrantApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/grant_api_key/types.ts#L25-L29
type GrantApiKey struct {
	Expiration      *Duration                `json:"expiration,omitempty"`
	Name            Name                     `json:"name"`
	RoleDescriptors []map[string]interface{} `json:"role_descriptors,omitempty"`
}

// GrantApiKeyBuilder holds GrantApiKey struct and provides a builder API.
type GrantApiKeyBuilder struct {
	v *GrantApiKey
}

// NewGrantApiKey provides a builder for the GrantApiKey struct.
func NewGrantApiKeyBuilder() *GrantApiKeyBuilder {
	r := GrantApiKeyBuilder{
		&GrantApiKey{},
	}

	return &r
}

// Build finalize the chain and returns the GrantApiKey struct
func (rb *GrantApiKeyBuilder) Build() GrantApiKey {
	return *rb.v
}

func (rb *GrantApiKeyBuilder) Expiration(expiration *DurationBuilder) *GrantApiKeyBuilder {
	v := expiration.Build()
	rb.v.Expiration = &v
	return rb
}

func (rb *GrantApiKeyBuilder) Name(name Name) *GrantApiKeyBuilder {
	rb.v.Name = name
	return rb
}

func (rb *GrantApiKeyBuilder) RoleDescriptors(value ...map[string]interface{}) *GrantApiKeyBuilder {
	rb.v.RoleDescriptors = value
	return rb
}
