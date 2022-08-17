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

// RegexpQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/term.ts#L102-L114
type RegexpQuery struct {
	Boost                 *float32               `json:"boost,omitempty"`
	CaseInsensitive       *bool                  `json:"case_insensitive,omitempty"`
	Flags                 *string                `json:"flags,omitempty"`
	MaxDeterminizedStates *int                   `json:"max_determinized_states,omitempty"`
	QueryName_            *string                `json:"_name,omitempty"`
	Rewrite               *MultiTermQueryRewrite `json:"rewrite,omitempty"`
	Value                 string                 `json:"value"`
}

// RegexpQueryBuilder holds RegexpQuery struct and provides a builder API.
type RegexpQueryBuilder struct {
	v *RegexpQuery
}

// NewRegexpQuery provides a builder for the RegexpQuery struct.
func NewRegexpQueryBuilder() *RegexpQueryBuilder {
	r := RegexpQueryBuilder{
		&RegexpQuery{},
	}

	return &r
}

// Build finalize the chain and returns the RegexpQuery struct
func (rb *RegexpQueryBuilder) Build() RegexpQuery {
	return *rb.v
}

func (rb *RegexpQueryBuilder) Boost(boost float32) *RegexpQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *RegexpQueryBuilder) CaseInsensitive(caseinsensitive bool) *RegexpQueryBuilder {
	rb.v.CaseInsensitive = &caseinsensitive
	return rb
}

func (rb *RegexpQueryBuilder) Flags(flags string) *RegexpQueryBuilder {
	rb.v.Flags = &flags
	return rb
}

func (rb *RegexpQueryBuilder) MaxDeterminizedStates(maxdeterminizedstates int) *RegexpQueryBuilder {
	rb.v.MaxDeterminizedStates = &maxdeterminizedstates
	return rb
}

func (rb *RegexpQueryBuilder) QueryName_(queryname_ string) *RegexpQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *RegexpQueryBuilder) Rewrite(rewrite MultiTermQueryRewrite) *RegexpQueryBuilder {
	rb.v.Rewrite = &rewrite
	return rb
}

func (rb *RegexpQueryBuilder) Value(value string) *RegexpQueryBuilder {
	rb.v.Value = value
	return rb
}
