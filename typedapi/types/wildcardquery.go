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

// WildcardQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L149-L162
type WildcardQuery struct {
	Boost *float32 `json:"boost,omitempty"`
	// CaseInsensitive Allows case insensitive matching of the pattern with the indexed field values
	// when set to true. Default is false which means the case sensitivity of
	// matching depends on the underlying field’s mapping.
	CaseInsensitive *bool   `json:"case_insensitive,omitempty"`
	QueryName_      *string `json:"_name,omitempty"`
	// Rewrite Method used to rewrite the query
	Rewrite *MultiTermQueryRewrite `json:"rewrite,omitempty"`
	// Value Wildcard pattern for terms you wish to find in the provided field. Required,
	// when wildcard is not set.
	Value *string `json:"value,omitempty"`
	// Wildcard Wildcard pattern for terms you wish to find in the provided field. Required,
	// when value is not set.
	Wildcard *string `json:"wildcard,omitempty"`
}

// WildcardQueryBuilder holds WildcardQuery struct and provides a builder API.
type WildcardQueryBuilder struct {
	v *WildcardQuery
}

// NewWildcardQuery provides a builder for the WildcardQuery struct.
func NewWildcardQueryBuilder() *WildcardQueryBuilder {
	r := WildcardQueryBuilder{
		&WildcardQuery{},
	}

	return &r
}

// Build finalize the chain and returns the WildcardQuery struct
func (rb *WildcardQueryBuilder) Build() WildcardQuery {
	return *rb.v
}

func (rb *WildcardQueryBuilder) Boost(boost float32) *WildcardQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

// CaseInsensitive Allows case insensitive matching of the pattern with the indexed field values
// when set to true. Default is false which means the case sensitivity of
// matching depends on the underlying field’s mapping.

func (rb *WildcardQueryBuilder) CaseInsensitive(caseinsensitive bool) *WildcardQueryBuilder {
	rb.v.CaseInsensitive = &caseinsensitive
	return rb
}

func (rb *WildcardQueryBuilder) QueryName_(queryname_ string) *WildcardQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

// Rewrite Method used to rewrite the query

func (rb *WildcardQueryBuilder) Rewrite(rewrite MultiTermQueryRewrite) *WildcardQueryBuilder {
	rb.v.Rewrite = &rewrite
	return rb
}

// Value Wildcard pattern for terms you wish to find in the provided field. Required,
// when wildcard is not set.

func (rb *WildcardQueryBuilder) Value(value string) *WildcardQueryBuilder {
	rb.v.Value = &value
	return rb
}

// Wildcard Wildcard pattern for terms you wish to find in the provided field. Required,
// when value is not set.

func (rb *WildcardQueryBuilder) Wildcard(wildcard string) *WildcardQueryBuilder {
	rb.v.Wildcard = &wildcard
	return rb
}
