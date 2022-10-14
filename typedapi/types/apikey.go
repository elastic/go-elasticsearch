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

// ApiKey type.
//
// https://github.com/elastic/elasticsearch-specification/blob/93ed2b29c9e75f49cd340f06286d6ead5965f900/specification/security/_types/ApiKey.ts#L27-L41
type ApiKey struct {
	Creation        *int64                      `json:"creation,omitempty"`
	Expiration      *int64                      `json:"expiration,omitempty"`
	Id              Id                          `json:"id"`
	Invalidated     *bool                       `json:"invalidated,omitempty"`
	LimitedBy       []map[string]RoleDescriptor `json:"limited_by,omitempty"`
	Metadata        *Metadata                   `json:"metadata,omitempty"`
	Name            Name                        `json:"name"`
	Realm           *string                     `json:"realm,omitempty"`
	RoleDescriptors map[string]RoleDescriptor   `json:"role_descriptors,omitempty"`
	Sort_           *SortResults                `json:"_sort,omitempty"`
	Username        *Username                   `json:"username,omitempty"`
}

// ApiKeyBuilder holds ApiKey struct and provides a builder API.
type ApiKeyBuilder struct {
	v *ApiKey
}

// NewApiKey provides a builder for the ApiKey struct.
func NewApiKeyBuilder() *ApiKeyBuilder {
	r := ApiKeyBuilder{
		&ApiKey{
			RoleDescriptors: make(map[string]RoleDescriptor, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the ApiKey struct
func (rb *ApiKeyBuilder) Build() ApiKey {
	return *rb.v
}

func (rb *ApiKeyBuilder) Creation(creation int64) *ApiKeyBuilder {
	rb.v.Creation = &creation
	return rb
}

func (rb *ApiKeyBuilder) Expiration(expiration int64) *ApiKeyBuilder {
	rb.v.Expiration = &expiration
	return rb
}

func (rb *ApiKeyBuilder) Id(id Id) *ApiKeyBuilder {
	rb.v.Id = id
	return rb
}

func (rb *ApiKeyBuilder) Invalidated(invalidated bool) *ApiKeyBuilder {
	rb.v.Invalidated = &invalidated
	return rb
}

func (rb *ApiKeyBuilder) LimitedBy(value ...map[string]RoleDescriptor) *ApiKeyBuilder {
	rb.v.LimitedBy = value
	return rb
}

func (rb *ApiKeyBuilder) Metadata(metadata *MetadataBuilder) *ApiKeyBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *ApiKeyBuilder) Name(name Name) *ApiKeyBuilder {
	rb.v.Name = name
	return rb
}

func (rb *ApiKeyBuilder) Realm(realm string) *ApiKeyBuilder {
	rb.v.Realm = &realm
	return rb
}

func (rb *ApiKeyBuilder) RoleDescriptors(values map[string]*RoleDescriptorBuilder) *ApiKeyBuilder {
	tmp := make(map[string]RoleDescriptor, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.RoleDescriptors = tmp
	return rb
}

func (rb *ApiKeyBuilder) Sort_(sort_ *SortResultsBuilder) *ApiKeyBuilder {
	v := sort_.Build()
	rb.v.Sort_ = &v
	return rb
}

func (rb *ApiKeyBuilder) Username(username Username) *ApiKeyBuilder {
	rb.v.Username = &username
	return rb
}
