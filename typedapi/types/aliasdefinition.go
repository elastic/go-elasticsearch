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

// AliasDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/AliasDefinition.ts#L22-L30
type AliasDefinition struct {
	Filter        *QueryContainer `json:"filter,omitempty"`
	IndexRouting  *string         `json:"index_routing,omitempty"`
	IsHidden      *bool           `json:"is_hidden,omitempty"`
	IsWriteIndex  *bool           `json:"is_write_index,omitempty"`
	Routing       *string         `json:"routing,omitempty"`
	SearchRouting *string         `json:"search_routing,omitempty"`
}

// AliasDefinitionBuilder holds AliasDefinition struct and provides a builder API.
type AliasDefinitionBuilder struct {
	v *AliasDefinition
}

// NewAliasDefinition provides a builder for the AliasDefinition struct.
func NewAliasDefinitionBuilder() *AliasDefinitionBuilder {
	r := AliasDefinitionBuilder{
		&AliasDefinition{},
	}

	return &r
}

// Build finalize the chain and returns the AliasDefinition struct
func (rb *AliasDefinitionBuilder) Build() AliasDefinition {
	return *rb.v
}

func (rb *AliasDefinitionBuilder) Filter(filter *QueryContainerBuilder) *AliasDefinitionBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *AliasDefinitionBuilder) IndexRouting(indexrouting string) *AliasDefinitionBuilder {
	rb.v.IndexRouting = &indexrouting
	return rb
}

func (rb *AliasDefinitionBuilder) IsHidden(ishidden bool) *AliasDefinitionBuilder {
	rb.v.IsHidden = &ishidden
	return rb
}

func (rb *AliasDefinitionBuilder) IsWriteIndex(iswriteindex bool) *AliasDefinitionBuilder {
	rb.v.IsWriteIndex = &iswriteindex
	return rb
}

func (rb *AliasDefinitionBuilder) Routing(routing string) *AliasDefinitionBuilder {
	rb.v.Routing = &routing
	return rb
}

func (rb *AliasDefinitionBuilder) SearchRouting(searchrouting string) *AliasDefinitionBuilder {
	rb.v.SearchRouting = &searchrouting
	return rb
}
