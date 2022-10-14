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
// https://github.com/elastic/elasticsearch-specification/tree/93ed2b29c9e75f49cd340f06286d6ead5965f900


package types

// GrantApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/security/grant_api_key/types.ts#L25-L32
type GrantApiKey struct {
	Expiration      *DurationLarge              `json:"expiration,omitempty"`
	Metadata        *Metadata                   `json:"metadata,omitempty"`
	Name            Name                        `json:"name"`
	RoleDescriptors []map[string]RoleDescriptor `json:"role_descriptors,omitempty"`
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

func (rb *GrantApiKeyBuilder) Expiration(expiration DurationLarge) *GrantApiKeyBuilder {
	rb.v.Expiration = &expiration
	return rb
}

func (rb *GrantApiKeyBuilder) Metadata(metadata *MetadataBuilder) *GrantApiKeyBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *GrantApiKeyBuilder) Name(name Name) *GrantApiKeyBuilder {
	rb.v.Name = name
	return rb
}

func (rb *GrantApiKeyBuilder) RoleDescriptors(arg []map[string]RoleDescriptor) *GrantApiKeyBuilder {
	rb.v.RoleDescriptors = arg
	return rb
}
