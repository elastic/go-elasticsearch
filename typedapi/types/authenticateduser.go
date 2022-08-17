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

// AuthenticatedUser type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/get_token/types.ts#L40-L45
type AuthenticatedUser struct {
	AuthenticationProvider *AuthenticationProvider `json:"authentication_provider,omitempty"`
	AuthenticationRealm    UserRealm               `json:"authentication_realm"`
	AuthenticationType     string                  `json:"authentication_type"`
	Email                  string                  `json:"email,omitempty"`
	Enabled                bool                    `json:"enabled"`
	FullName               Name                    `json:"full_name,omitempty"`
	LookupRealm            UserRealm               `json:"lookup_realm"`
	Metadata               Metadata                `json:"metadata"`
	Roles                  []string                `json:"roles"`
	Username               Username                `json:"username"`
}

// AuthenticatedUserBuilder holds AuthenticatedUser struct and provides a builder API.
type AuthenticatedUserBuilder struct {
	v *AuthenticatedUser
}

// NewAuthenticatedUser provides a builder for the AuthenticatedUser struct.
func NewAuthenticatedUserBuilder() *AuthenticatedUserBuilder {
	r := AuthenticatedUserBuilder{
		&AuthenticatedUser{},
	}

	return &r
}

// Build finalize the chain and returns the AuthenticatedUser struct
func (rb *AuthenticatedUserBuilder) Build() AuthenticatedUser {
	return *rb.v
}

func (rb *AuthenticatedUserBuilder) AuthenticationProvider(authenticationprovider *AuthenticationProviderBuilder) *AuthenticatedUserBuilder {
	v := authenticationprovider.Build()
	rb.v.AuthenticationProvider = &v
	return rb
}

func (rb *AuthenticatedUserBuilder) AuthenticationRealm(authenticationrealm *UserRealmBuilder) *AuthenticatedUserBuilder {
	v := authenticationrealm.Build()
	rb.v.AuthenticationRealm = v
	return rb
}

func (rb *AuthenticatedUserBuilder) AuthenticationType(authenticationtype string) *AuthenticatedUserBuilder {
	rb.v.AuthenticationType = authenticationtype
	return rb
}

func (rb *AuthenticatedUserBuilder) Email(email string) *AuthenticatedUserBuilder {
	rb.v.Email = email
	return rb
}

func (rb *AuthenticatedUserBuilder) Enabled(enabled bool) *AuthenticatedUserBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *AuthenticatedUserBuilder) FullName(fullname Name) *AuthenticatedUserBuilder {
	rb.v.FullName = fullname
	return rb
}

func (rb *AuthenticatedUserBuilder) LookupRealm(lookuprealm *UserRealmBuilder) *AuthenticatedUserBuilder {
	v := lookuprealm.Build()
	rb.v.LookupRealm = v
	return rb
}

func (rb *AuthenticatedUserBuilder) Metadata(metadata *MetadataBuilder) *AuthenticatedUserBuilder {
	v := metadata.Build()
	rb.v.Metadata = v
	return rb
}

func (rb *AuthenticatedUserBuilder) Roles(roles ...string) *AuthenticatedUserBuilder {
	rb.v.Roles = roles
	return rb
}

func (rb *AuthenticatedUserBuilder) Username(username Username) *AuthenticatedUserBuilder {
	rb.v.Username = username
	return rb
}
