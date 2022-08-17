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


package suggestuserprofiles

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package suggestuserprofiles
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/suggest_user_profiles/Request.ts#L24-L66
type Request struct {

	// Data List of filters for the `data` field of the profile document.
	// To return all content use `data=*`. To return a subset of content
	// use `data=<key>` to retrieve content nested under the specified `<key>`.
	// By default returns no `data` content.
	Data []string `json:"data,omitempty"`

	// Hint Extra search criteria to improve relevance of the suggestion result.
	// Profiles matching the spcified hint are ranked higher in the response.
	// Profiles not matching the hint don't exclude the profile from the response
	// as long as the profile matches the `name` field query.
	Hint *types.Hint `json:"hint,omitempty"`

	// Name Query string used to match name-related fields in user profile documents.
	// Name-related fields are the user's `username`, `full_name`, and `email`.
	Name *string `json:"name,omitempty"`

	// Size Number of profiles to return.
	Size *int64 `json:"size,omitempty"`
}

// RequestBuilder is the builder API for the suggestuserprofiles.Request
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
		return nil, fmt.Errorf("could not deserialise json into Suggestuserprofiles request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Data(arg []string) *RequestBuilder {
	rb.v.Data = arg
	return rb
}

func (rb *RequestBuilder) Hint(hint *types.HintBuilder) *RequestBuilder {
	v := hint.Build()
	rb.v.Hint = &v
	return rb
}

func (rb *RequestBuilder) Name(name string) *RequestBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *RequestBuilder) Size(size int64) *RequestBuilder {
	rb.v.Size = &size
	return rb
}
