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


package searchtemplate

import (
	"encoding/json"
	"fmt"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

// Request holds the request body struct for the package searchtemplate
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search_template/SearchTemplateRequest.ts#L32-L96
type Request struct {
	Explain *bool `json:"explain,omitempty"`

	// Id ID of the search template to use. If no source is specified,
	// this parameter is required.
	Id *types.Id `json:"id,omitempty"`

	Params map[string]interface{} `json:"params,omitempty"`

	Profile *bool `json:"profile,omitempty"`

	// Source An inline search template. Supports the same parameters as the search API's
	// request body. Also supports Mustache variables. If no id is specified, this
	// parameter is required.
	Source *string `json:"source,omitempty"`
}

// RequestBuilder is the builder API for the searchtemplate.Request
type RequestBuilder struct {
	v *Request
}

// NewRequest returns a RequestBuilder which can be chained and built to retrieve a RequestBuilder
func NewRequestBuilder() *RequestBuilder {
	r := RequestBuilder{
		&Request{
			Params: make(map[string]interface{}, 0),
		},
	}
	return &r
}

// FromJSON allows to load an arbitrary json into the request structure
func (rb *RequestBuilder) FromJSON(data string) (*Request, error) {
	var req Request
	err := json.Unmarshal([]byte(data), &req)

	if err != nil {
		return nil, fmt.Errorf("could not deserialise json into Searchtemplate request: %w", err)
	}

	return &req, nil
}

// Build finalize the chain and returns the Request struct.
func (rb *RequestBuilder) Build() *Request {
	return rb.v
}

func (rb *RequestBuilder) Explain(explain bool) *RequestBuilder {
	rb.v.Explain = &explain
	return rb
}

func (rb *RequestBuilder) Id(id types.Id) *RequestBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *RequestBuilder) Params(value map[string]interface{}) *RequestBuilder {
	rb.v.Params = value
	return rb
}

func (rb *RequestBuilder) Profile(profile bool) *RequestBuilder {
	rb.v.Profile = &profile
	return rb
}

func (rb *RequestBuilder) Source(source string) *RequestBuilder {
	rb.v.Source = &source
	return rb
}
