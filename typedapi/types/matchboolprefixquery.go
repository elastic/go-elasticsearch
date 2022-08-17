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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

// MatchBoolPrefixQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L160-L171
type MatchBoolPrefixQuery struct {
	Analyzer            *string                `json:"analyzer,omitempty"`
	Boost               *float32               `json:"boost,omitempty"`
	Fuzziness           *Fuzziness             `json:"fuzziness,omitempty"`
	FuzzyRewrite        *MultiTermQueryRewrite `json:"fuzzy_rewrite,omitempty"`
	FuzzyTranspositions *bool                  `json:"fuzzy_transpositions,omitempty"`
	MaxExpansions       *int                   `json:"max_expansions,omitempty"`
	MinimumShouldMatch  *MinimumShouldMatch    `json:"minimum_should_match,omitempty"`
	Operator            *operator.Operator     `json:"operator,omitempty"`
	PrefixLength        *int                   `json:"prefix_length,omitempty"`
	Query               string                 `json:"query"`
	QueryName_          *string                `json:"_name,omitempty"`
}

// MatchBoolPrefixQueryBuilder holds MatchBoolPrefixQuery struct and provides a builder API.
type MatchBoolPrefixQueryBuilder struct {
	v *MatchBoolPrefixQuery
}

// NewMatchBoolPrefixQuery provides a builder for the MatchBoolPrefixQuery struct.
func NewMatchBoolPrefixQueryBuilder() *MatchBoolPrefixQueryBuilder {
	r := MatchBoolPrefixQueryBuilder{
		&MatchBoolPrefixQuery{},
	}

	return &r
}

// Build finalize the chain and returns the MatchBoolPrefixQuery struct
func (rb *MatchBoolPrefixQueryBuilder) Build() MatchBoolPrefixQuery {
	return *rb.v
}

func (rb *MatchBoolPrefixQueryBuilder) Analyzer(analyzer string) *MatchBoolPrefixQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) Boost(boost float32) *MatchBoolPrefixQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) Fuzziness(fuzziness *FuzzinessBuilder) *MatchBoolPrefixQueryBuilder {
	v := fuzziness.Build()
	rb.v.Fuzziness = &v
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) FuzzyRewrite(fuzzyrewrite MultiTermQueryRewrite) *MatchBoolPrefixQueryBuilder {
	rb.v.FuzzyRewrite = &fuzzyrewrite
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) FuzzyTranspositions(fuzzytranspositions bool) *MatchBoolPrefixQueryBuilder {
	rb.v.FuzzyTranspositions = &fuzzytranspositions
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) MaxExpansions(maxexpansions int) *MatchBoolPrefixQueryBuilder {
	rb.v.MaxExpansions = &maxexpansions
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) MinimumShouldMatch(minimumshouldmatch *MinimumShouldMatchBuilder) *MatchBoolPrefixQueryBuilder {
	v := minimumshouldmatch.Build()
	rb.v.MinimumShouldMatch = &v
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) Operator(operator operator.Operator) *MatchBoolPrefixQueryBuilder {
	rb.v.Operator = &operator
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) PrefixLength(prefixlength int) *MatchBoolPrefixQueryBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) Query(query string) *MatchBoolPrefixQueryBuilder {
	rb.v.Query = query
	return rb
}

func (rb *MatchBoolPrefixQueryBuilder) QueryName_(queryname_ string) *MatchBoolPrefixQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}
