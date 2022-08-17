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

// Alias type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/_types/Alias.ts#L23-L30
type Alias struct {
	Filter        *QueryContainer `json:"filter,omitempty"`
	IndexRouting  *Routing        `json:"index_routing,omitempty"`
	IsHidden      *bool           `json:"is_hidden,omitempty"`
	IsWriteIndex  *bool           `json:"is_write_index,omitempty"`
	Routing       *Routing        `json:"routing,omitempty"`
	SearchRouting *Routing        `json:"search_routing,omitempty"`
}

// AliasBuilder holds Alias struct and provides a builder API.
type AliasBuilder struct {
	v *Alias
}

// NewAlias provides a builder for the Alias struct.
func NewAliasBuilder() *AliasBuilder {
	r := AliasBuilder{
		&Alias{},
	}

	return &r
}

// Build finalize the chain and returns the Alias struct
func (rb *AliasBuilder) Build() Alias {
	return *rb.v
}

func (rb *AliasBuilder) Filter(filter *QueryContainerBuilder) *AliasBuilder {
	v := filter.Build()
	rb.v.Filter = &v
	return rb
}

func (rb *AliasBuilder) IndexRouting(indexrouting Routing) *AliasBuilder {
	rb.v.IndexRouting = &indexrouting
	return rb
}

func (rb *AliasBuilder) IsHidden(ishidden bool) *AliasBuilder {
	rb.v.IsHidden = &ishidden
	return rb
}

func (rb *AliasBuilder) IsWriteIndex(iswriteindex bool) *AliasBuilder {
	rb.v.IsWriteIndex = &iswriteindex
	return rb
}

func (rb *AliasBuilder) Routing(routing Routing) *AliasBuilder {
	rb.v.Routing = &routing
	return rb
}

func (rb *AliasBuilder) SearchRouting(searchrouting Routing) *AliasBuilder {
	rb.v.SearchRouting = &searchrouting
	return rb
}
