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

// AliasesRecord type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cat/aliases/types.ts#L22-L53
type AliasesRecord struct {
	// Alias alias name
	Alias *string `json:"alias,omitempty"`
	// Filter filter
	Filter *string `json:"filter,omitempty"`
	// Index index alias points to
	Index *IndexName `json:"index,omitempty"`
	// IsWriteIndex write index
	IsWriteIndex *string `json:"is_write_index,omitempty"`
	// RoutingIndex index routing
	RoutingIndex *string `json:"routing.index,omitempty"`
	// RoutingSearch search routing
	RoutingSearch *string `json:"routing.search,omitempty"`
}

// AliasesRecordBuilder holds AliasesRecord struct and provides a builder API.
type AliasesRecordBuilder struct {
	v *AliasesRecord
}

// NewAliasesRecord provides a builder for the AliasesRecord struct.
func NewAliasesRecordBuilder() *AliasesRecordBuilder {
	r := AliasesRecordBuilder{
		&AliasesRecord{},
	}

	return &r
}

// Build finalize the chain and returns the AliasesRecord struct
func (rb *AliasesRecordBuilder) Build() AliasesRecord {
	return *rb.v
}

// Alias alias name

func (rb *AliasesRecordBuilder) Alias(alias string) *AliasesRecordBuilder {
	rb.v.Alias = &alias
	return rb
}

// Filter filter

func (rb *AliasesRecordBuilder) Filter(filter string) *AliasesRecordBuilder {
	rb.v.Filter = &filter
	return rb
}

// Index index alias points to

func (rb *AliasesRecordBuilder) Index(index IndexName) *AliasesRecordBuilder {
	rb.v.Index = &index
	return rb
}

// IsWriteIndex write index

func (rb *AliasesRecordBuilder) IsWriteIndex(iswriteindex string) *AliasesRecordBuilder {
	rb.v.IsWriteIndex = &iswriteindex
	return rb
}

// RoutingIndex index routing

func (rb *AliasesRecordBuilder) RoutingIndex(routingindex string) *AliasesRecordBuilder {
	rb.v.RoutingIndex = &routingindex
	return rb
}

// RoutingSearch search routing

func (rb *AliasesRecordBuilder) RoutingSearch(routingsearch string) *AliasesRecordBuilder {
	rb.v.RoutingSearch = &routingsearch
	return rb
}
