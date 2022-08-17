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


package putrolemapping

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package putrolemapping
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/put_role_mapping/SecurityPutRoleMappingRequest.ts#L24-L43
type Request struct {
	Enabled *bool `json:"enabled,omitempty"`

	Metadata *types.Metadata `json:"metadata,omitempty"`

	Roles []string `json:"roles,omitempty"`

	Rules *types.RoleMappingRule `json:"rules,omitempty"`

	RunAs []string `json:"run_as,omitempty"`
}

// RequestBuilder is the builder API for the putrolemapping.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Putrolemapping request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Enabled(enabled bool) *RequestBuilder {
	rb.v.Enabled = &enabled
	return rb
}

func (rb *RequestBuilder) Metadata(metadata *types.MetadataBuilder) *RequestBuilder {
	v := metadata.Build()
	rb.v.Metadata = &v
	return rb
}

func (rb *RequestBuilder) Roles(roles ...string) *RequestBuilder {
	rb.v.Roles = roles
	return rb
}

func (rb *RequestBuilder) Rules(rules *types.RoleMappingRuleBuilder) *RequestBuilder {
	v := rules.Build()
	rb.v.Rules = &v
	return rb
}

func (rb *RequestBuilder) RunAs(run_as ...string) *RequestBuilder {
	rb.v.RunAs = run_as
	return rb
}
