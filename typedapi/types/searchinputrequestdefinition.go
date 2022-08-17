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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/searchtype"
)

// SearchInputRequestDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/watcher/_types/Input.ts#L120-L127
type SearchInputRequestDefinition struct {
	Body               *SearchInputRequestBody    `json:"body,omitempty"`
	Indices            []IndexName                `json:"indices,omitempty"`
	IndicesOptions     *IndicesOptions            `json:"indices_options,omitempty"`
	RestTotalHitsAsInt *bool                      `json:"rest_total_hits_as_int,omitempty"`
	SearchType         *searchtype.SearchType     `json:"search_type,omitempty"`
	Template           *SearchTemplateRequestBody `json:"template,omitempty"`
}

// SearchInputRequestDefinitionBuilder holds SearchInputRequestDefinition struct and provides a builder API.
type SearchInputRequestDefinitionBuilder struct {
	v *SearchInputRequestDefinition
}

// NewSearchInputRequestDefinition provides a builder for the SearchInputRequestDefinition struct.
func NewSearchInputRequestDefinitionBuilder() *SearchInputRequestDefinitionBuilder {
	r := SearchInputRequestDefinitionBuilder{
		&SearchInputRequestDefinition{},
	}

	return &r
}

// Build finalize the chain and returns the SearchInputRequestDefinition struct
func (rb *SearchInputRequestDefinitionBuilder) Build() SearchInputRequestDefinition {
	return *rb.v
}

func (rb *SearchInputRequestDefinitionBuilder) Body(body *SearchInputRequestBodyBuilder) *SearchInputRequestDefinitionBuilder {
	v := body.Build()
	rb.v.Body = &v
	return rb
}

func (rb *SearchInputRequestDefinitionBuilder) Indices(indices ...IndexName) *SearchInputRequestDefinitionBuilder {
	rb.v.Indices = indices
	return rb
}

func (rb *SearchInputRequestDefinitionBuilder) IndicesOptions(indicesoptions *IndicesOptionsBuilder) *SearchInputRequestDefinitionBuilder {
	v := indicesoptions.Build()
	rb.v.IndicesOptions = &v
	return rb
}

func (rb *SearchInputRequestDefinitionBuilder) RestTotalHitsAsInt(resttotalhitsasint bool) *SearchInputRequestDefinitionBuilder {
	rb.v.RestTotalHitsAsInt = &resttotalhitsasint
	return rb
}

func (rb *SearchInputRequestDefinitionBuilder) SearchType(searchtype searchtype.SearchType) *SearchInputRequestDefinitionBuilder {
	rb.v.SearchType = &searchtype
	return rb
}

func (rb *SearchInputRequestDefinitionBuilder) Template(template *SearchTemplateRequestBodyBuilder) *SearchInputRequestDefinitionBuilder {
	v := template.Build()
	rb.v.Template = &v
	return rb
}
