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

// AuthenticationProvider type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/get_token/types.ts#L35-L38
type AuthenticationProvider struct {
	Name Name   `json:"name"`
	Type string `json:"type"`
}

// AuthenticationProviderBuilder holds AuthenticationProvider struct and provides a builder API.
type AuthenticationProviderBuilder struct {
	v *AuthenticationProvider
}

// NewAuthenticationProvider provides a builder for the AuthenticationProvider struct.
func NewAuthenticationProviderBuilder() *AuthenticationProviderBuilder {
	r := AuthenticationProviderBuilder{
		&AuthenticationProvider{},
	}

	return &r
}

// Build finalize the chain and returns the AuthenticationProvider struct
func (rb *AuthenticationProviderBuilder) Build() AuthenticationProvider {
	return *rb.v
}

func (rb *AuthenticationProviderBuilder) Name(name Name) *AuthenticationProviderBuilder {
	rb.v.Name = name
	return rb
}

func (rb *AuthenticationProviderBuilder) Type_(type_ string) *AuthenticationProviderBuilder {
	rb.v.Type = type_
	return rb
}
