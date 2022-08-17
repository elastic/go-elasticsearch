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

// DatafeedAuthorization type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/Authorization.ts#L31-L43
type DatafeedAuthorization struct {
	// ApiKey If an API key was used for the most recent update to the datafeed, its name
	// and identifier are listed in the response.
	ApiKey *ApiKeyAuthorization `json:"api_key,omitempty"`
	// Roles If a user ID was used for the most recent update to the datafeed, its roles
	// at the time of the update are listed in the response.
	Roles []string `json:"roles,omitempty"`
	// ServiceAccount If a service account was used for the most recent update to the datafeed, the
	// account name is listed in the response.
	ServiceAccount *string `json:"service_account,omitempty"`
}

// DatafeedAuthorizationBuilder holds DatafeedAuthorization struct and provides a builder API.
type DatafeedAuthorizationBuilder struct {
	v *DatafeedAuthorization
}

// NewDatafeedAuthorization provides a builder for the DatafeedAuthorization struct.
func NewDatafeedAuthorizationBuilder() *DatafeedAuthorizationBuilder {
	r := DatafeedAuthorizationBuilder{
		&DatafeedAuthorization{},
	}

	return &r
}

// Build finalize the chain and returns the DatafeedAuthorization struct
func (rb *DatafeedAuthorizationBuilder) Build() DatafeedAuthorization {
	return *rb.v
}

// ApiKey If an API key was used for the most recent update to the datafeed, its name
// and identifier are listed in the response.

func (rb *DatafeedAuthorizationBuilder) ApiKey(apikey *ApiKeyAuthorizationBuilder) *DatafeedAuthorizationBuilder {
	v := apikey.Build()
	rb.v.ApiKey = &v
	return rb
}

// Roles If a user ID was used for the most recent update to the datafeed, its roles
// at the time of the update are listed in the response.

func (rb *DatafeedAuthorizationBuilder) Roles(roles ...string) *DatafeedAuthorizationBuilder {
	rb.v.Roles = roles
	return rb
}

// ServiceAccount If a service account was used for the most recent update to the datafeed, the
// account name is listed in the response.

func (rb *DatafeedAuthorizationBuilder) ServiceAccount(serviceaccount string) *DatafeedAuthorizationBuilder {
	rb.v.ServiceAccount = &serviceaccount
	return rb
}
