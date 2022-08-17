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

// TransformAuthorization type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Authorization.ts#L59-L71
type TransformAuthorization struct {
	// ApiKey If an API key was used for the most recent update to the transform, its name
	// and identifier are listed in the response.
	ApiKey *ApiKeyAuthorization `json:"api_key,omitempty"`
	// Roles If a user ID was used for the most recent update to the transform, its roles
	// at the time of the update are listed in the response.
	Roles []string `json:"roles,omitempty"`
	// ServiceAccount If a service account was used for the most recent update to the transform,
	// the account name is listed in the response.
	ServiceAccount *string `json:"service_account,omitempty"`
}

// TransformAuthorizationBuilder holds TransformAuthorization struct and provides a builder API.
type TransformAuthorizationBuilder struct {
	v *TransformAuthorization
}

// NewTransformAuthorization provides a builder for the TransformAuthorization struct.
func NewTransformAuthorizationBuilder() *TransformAuthorizationBuilder {
	r := TransformAuthorizationBuilder{
		&TransformAuthorization{},
	}

	return &r
}

// Build finalize the chain and returns the TransformAuthorization struct
func (rb *TransformAuthorizationBuilder) Build() TransformAuthorization {
	return *rb.v
}

// ApiKey If an API key was used for the most recent update to the transform, its name
// and identifier are listed in the response.

func (rb *TransformAuthorizationBuilder) ApiKey(apikey *ApiKeyAuthorizationBuilder) *TransformAuthorizationBuilder {
	v := apikey.Build()
	rb.v.ApiKey = &v
	return rb
}

// Roles If a user ID was used for the most recent update to the transform, its roles
// at the time of the update are listed in the response.

func (rb *TransformAuthorizationBuilder) Roles(roles ...string) *TransformAuthorizationBuilder {
	rb.v.Roles = roles
	return rb
}

// ServiceAccount If a service account was used for the most recent update to the transform,
// the account name is listed in the response.

func (rb *TransformAuthorizationBuilder) ServiceAccount(serviceaccount string) *TransformAuthorizationBuilder {
	rb.v.ServiceAccount = &serviceaccount
	return rb
}
