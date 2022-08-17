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

// Realm type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/xpack/usage/types.ts#L406-L415
type Realm struct {
	Available                 bool         `json:"available"`
	Cache                     []RealmCache `json:"cache,omitempty"`
	Enabled                   bool         `json:"enabled"`
	HasAuthorizationRealms    []bool       `json:"has_authorization_realms,omitempty"`
	HasDefaultUsernamePattern []bool       `json:"has_default_username_pattern,omitempty"`
	HasTruststore             []bool       `json:"has_truststore,omitempty"`
	IsAuthenticationDelegated []bool       `json:"is_authentication_delegated,omitempty"`
	Name                      []string     `json:"name,omitempty"`
	Order                     []int64      `json:"order,omitempty"`
	Size                      []int64      `json:"size,omitempty"`
}

// RealmBuilder holds Realm struct and provides a builder API.
type RealmBuilder struct {
	v *Realm
}

// NewRealm provides a builder for the Realm struct.
func NewRealmBuilder() *RealmBuilder {
	r := RealmBuilder{
		&Realm{},
	}

	return &r
}

// Build finalize the chain and returns the Realm struct
func (rb *RealmBuilder) Build() Realm {
	return *rb.v
}

func (rb *RealmBuilder) Available(available bool) *RealmBuilder {
	rb.v.Available = available
	return rb
}

func (rb *RealmBuilder) Cache(cache []RealmCacheBuilder) *RealmBuilder {
	tmp := make([]RealmCache, len(cache))
	for _, value := range cache {
		tmp = append(tmp, value.Build())
	}
	rb.v.Cache = tmp
	return rb
}

func (rb *RealmBuilder) Enabled(enabled bool) *RealmBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *RealmBuilder) HasAuthorizationRealms(has_authorization_realms ...bool) *RealmBuilder {
	rb.v.HasAuthorizationRealms = has_authorization_realms
	return rb
}

func (rb *RealmBuilder) HasDefaultUsernamePattern(has_default_username_pattern ...bool) *RealmBuilder {
	rb.v.HasDefaultUsernamePattern = has_default_username_pattern
	return rb
}

func (rb *RealmBuilder) HasTruststore(has_truststore ...bool) *RealmBuilder {
	rb.v.HasTruststore = has_truststore
	return rb
}

func (rb *RealmBuilder) IsAuthenticationDelegated(is_authentication_delegated ...bool) *RealmBuilder {
	rb.v.IsAuthenticationDelegated = is_authentication_delegated
	return rb
}

func (rb *RealmBuilder) Name(name ...string) *RealmBuilder {
	rb.v.Name = name
	return rb
}

func (rb *RealmBuilder) Order(order ...int64) *RealmBuilder {
	rb.v.Order = order
	return rb
}

func (rb *RealmBuilder) Size(size ...int64) *RealmBuilder {
	rb.v.Size = size
	return rb
}
