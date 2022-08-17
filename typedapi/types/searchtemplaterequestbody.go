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

// SearchTemplateRequestBody type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Input.ts#L130-L147
type SearchTemplateRequestBody struct {
	Explain *bool `json:"explain,omitempty"`
	// Id ID of the search template to use. If no source is specified,
	// this parameter is required.
	Id      *Id                    `json:"id,omitempty"`
	Params  map[string]interface{} `json:"params,omitempty"`
	Profile *bool                  `json:"profile,omitempty"`
	// Source An inline search template. Supports the same parameters as the search API's
	// request body. Also supports Mustache variables. If no id is specified, this
	// parameter is required.
	Source *string `json:"source,omitempty"`
}

// SearchTemplateRequestBodyBuilder holds SearchTemplateRequestBody struct and provides a builder API.
type SearchTemplateRequestBodyBuilder struct {
	v *SearchTemplateRequestBody
}

// NewSearchTemplateRequestBody provides a builder for the SearchTemplateRequestBody struct.
func NewSearchTemplateRequestBodyBuilder() *SearchTemplateRequestBodyBuilder {
	r := SearchTemplateRequestBodyBuilder{
		&SearchTemplateRequestBody{
			Params: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the SearchTemplateRequestBody struct
func (rb *SearchTemplateRequestBodyBuilder) Build() SearchTemplateRequestBody {
	return *rb.v
}

func (rb *SearchTemplateRequestBodyBuilder) Explain(explain bool) *SearchTemplateRequestBodyBuilder {
	rb.v.Explain = &explain
	return rb
}

// Id ID of the search template to use. If no source is specified,
// this parameter is required.

func (rb *SearchTemplateRequestBodyBuilder) Id(id Id) *SearchTemplateRequestBodyBuilder {
	rb.v.Id = &id
	return rb
}

func (rb *SearchTemplateRequestBodyBuilder) Params(value map[string]interface{}) *SearchTemplateRequestBodyBuilder {
	rb.v.Params = value
	return rb
}

func (rb *SearchTemplateRequestBodyBuilder) Profile(profile bool) *SearchTemplateRequestBodyBuilder {
	rb.v.Profile = &profile
	return rb
}

// Source An inline search template. Supports the same parameters as the search API's
// request body. Also supports Mustache variables. If no id is specified, this
// parameter is required.

func (rb *SearchTemplateRequestBodyBuilder) Source(source string) *SearchTemplateRequestBodyBuilder {
	rb.v.Source = &source
	return rb
}
