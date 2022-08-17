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

// ApiKeyAuthorization type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Authorization.ts#L20-L29
type ApiKeyAuthorization struct {
	// Id The identifier for the API key.
	Id string `json:"id"`
	// Name The name of the API key.
	Name string `json:"name"`
}

// ApiKeyAuthorizationBuilder holds ApiKeyAuthorization struct and provides a builder API.
type ApiKeyAuthorizationBuilder struct {
	v *ApiKeyAuthorization
}

// NewApiKeyAuthorization provides a builder for the ApiKeyAuthorization struct.
func NewApiKeyAuthorizationBuilder() *ApiKeyAuthorizationBuilder {
	r := ApiKeyAuthorizationBuilder{
		&ApiKeyAuthorization{},
	}

	return &r
}

// Build finalize the chain and returns the ApiKeyAuthorization struct
func (rb *ApiKeyAuthorizationBuilder) Build() ApiKeyAuthorization {
	return *rb.v
}

// Id The identifier for the API key.

func (rb *ApiKeyAuthorizationBuilder) Id(id string) *ApiKeyAuthorizationBuilder {
	rb.v.Id = id
	return rb
}

// Name The name of the API key.

func (rb *ApiKeyAuthorizationBuilder) Name(name string) *ApiKeyAuthorizationBuilder {
	rb.v.Name = name
	return rb
}
