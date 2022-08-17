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


package createapikey

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package createapikey
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/create_api_key/SecurityCreateApiKeyRequest.ts#L26-L51
type Request struct {

	// Expiration Expiration time for the API key. By default, API keys never expire.
	Expiration *types.Duration `json:"expiration,omitempty"`

	// Metadata Arbitrary metadata that you want to associate with the API key. It supports
	// nested data structure. Within the metadata object, keys beginning with _ are
	// reserved for system usage.
	Metadata *types.Metadata `json:"metadata,omitempty"`

	// Name Specifies the name for this API key.
	Name *types.Name `json:"name,omitempty"`

	// RoleDescriptors An array of role descriptors for this API key. This parameter is optional.
	// When it is not specified or is an empty array, then the API key will have a
	// point in time snapshot of permissions of the authenticated user. If you
	// supply role descriptors then the resultant permissions would be an
	// intersection of API keys permissions and authenticated user’s permissions
	// thereby limiting the access scope for API keys. The structure of role
	// descriptor is the same as the request for create role API. For more details,
	// see create or update roles API.
	RoleDescriptors map[string]types.RoleDescriptor `json:"role_descriptors,omitempty"`
}

// RequestBuilder is the builder API for the createapikey.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			RoleDescriptors: make(map[string]types.RoleDescriptor, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Createapikey request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Expiration(expiration *types.DurationBuilder) *RequestBuilder {
	v := expiration.Build()
	rb.v.Expiration = &v
	return rb
}

func (rb *RequestBuilder) Metadata(metadata *types.MetadataBuilder) *RequestBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *RequestBuilder) Name(name types.Name) *RequestBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *RequestBuilder) RoleDescriptors(values map[string]*types.RoleDescriptorBuilder) *RequestBuilder {
	tmp := make(map[string]types.RoleDescriptor, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.RoleDescriptors = tmp
	return rb
}
