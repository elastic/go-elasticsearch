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

// PrefixQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L57-L66
type PrefixQuery struct {
	Boost           *float32               `json:"boost,omitempty"`
	CaseInsensitive *bool                  `json:"case_insensitive,omitempty"`
	QueryName_      *string                `json:"_name,omitempty"`
	Rewrite         *MultiTermQueryRewrite `json:"rewrite,omitempty"`
	Value           string                 `json:"value"`
}

// PrefixQueryBuilder holds PrefixQuery struct and provides a builder API.
type PrefixQueryBuilder struct {
	v *PrefixQuery
}

// NewPrefixQuery provides a builder for the PrefixQuery struct.
func NewPrefixQueryBuilder() *PrefixQueryBuilder {
	r := PrefixQueryBuilder{
		&PrefixQuery{},
	}

	return &r
}

// Build finalize the chain and returns the PrefixQuery struct
func (rb *PrefixQueryBuilder) Build() PrefixQuery {
	return *rb.v
}

func (rb *PrefixQueryBuilder) Boost(boost float32) *PrefixQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *PrefixQueryBuilder) CaseInsensitive(caseinsensitive bool) *PrefixQueryBuilder {
	rb.v.CaseInsensitive = &caseinsensitive
	return rb
}

func (rb *PrefixQueryBuilder) QueryName_(queryname_ string) *PrefixQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *PrefixQueryBuilder) Rewrite(rewrite MultiTermQueryRewrite) *PrefixQueryBuilder {
	rb.v.Rewrite = &rewrite
	return rb
}

func (rb *PrefixQueryBuilder) Value(value string) *PrefixQueryBuilder {
	rb.v.Value = value
	return rb
}
