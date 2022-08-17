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

// SimpleQueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L294-L312
type SimpleQueryStringQuery struct {
	AnalyzeWildcard                 *bool                   `json:"analyze_wildcard,omitempty"`
	Analyzer                        *string                 `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                   `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                `json:"boost,omitempty"`
	DefaultOperator                 *operator.Operator      `json:"default_operator,omitempty"`
	Fields                          []Field                 `json:"fields,omitempty"`
	Flags                           *SimpleQueryStringFlags `json:"flags,omitempty"`
	FuzzyMaxExpansions              *int                    `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength               *int                    `json:"fuzzy_prefix_length,omitempty"`
	FuzzyTranspositions             *bool                   `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool                   `json:"lenient,omitempty"`
	MinimumShouldMatch              *MinimumShouldMatch     `json:"minimum_should_match,omitempty"`
	Query                           string                  `json:"query"`
	QueryName_                      *string                 `json:"_name,omitempty"`
	QuoteFieldSuffix                *string                 `json:"quote_field_suffix,omitempty"`
}

// SimpleQueryStringQueryBuilder holds SimpleQueryStringQuery struct and provides a builder API.
type SimpleQueryStringQueryBuilder struct {
	v *SimpleQueryStringQuery
}

// NewSimpleQueryStringQuery provides a builder for the SimpleQueryStringQuery struct.
func NewSimpleQueryStringQueryBuilder() *SimpleQueryStringQueryBuilder {
	r := SimpleQueryStringQueryBuilder{
		&SimpleQueryStringQuery{},
	}

	return &r
}

// Build finalize the chain and returns the SimpleQueryStringQuery struct
func (rb *SimpleQueryStringQueryBuilder) Build() SimpleQueryStringQuery {
	return *rb.v
}

func (rb *SimpleQueryStringQueryBuilder) AnalyzeWildcard(analyzewildcard bool) *SimpleQueryStringQueryBuilder {
	rb.v.AnalyzeWildcard = &analyzewildcard
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) Analyzer(analyzer string) *SimpleQueryStringQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *SimpleQueryStringQueryBuilder {
	rb.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) Boost(boost float32) *SimpleQueryStringQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) DefaultOperator(defaultoperator operator.Operator) *SimpleQueryStringQueryBuilder {
	rb.v.DefaultOperator = &defaultoperator
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) Fields(fields ...Field) *SimpleQueryStringQueryBuilder {
	rb.v.Fields = fields
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) Flags(flags *SimpleQueryStringFlagsBuilder) *SimpleQueryStringQueryBuilder {
	v := flags.Build()
	rb.v.Flags = &v
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) FuzzyMaxExpansions(fuzzymaxexpansions int) *SimpleQueryStringQueryBuilder {
	rb.v.FuzzyMaxExpansions = &fuzzymaxexpansions
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) FuzzyPrefixLength(fuzzyprefixlength int) *SimpleQueryStringQueryBuilder {
	rb.v.FuzzyPrefixLength = &fuzzyprefixlength
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) FuzzyTranspositions(fuzzytranspositions bool) *SimpleQueryStringQueryBuilder {
	rb.v.FuzzyTranspositions = &fuzzytranspositions
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) Lenient(lenient bool) *SimpleQueryStringQueryBuilder {
	rb.v.Lenient = &lenient
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) MinimumShouldMatch(minimumshouldmatch *MinimumShouldMatchBuilder) *SimpleQueryStringQueryBuilder {
	v := minimumshouldmatch.Build()
	rb.v.MinimumShouldMatch = &v
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) Query(query string) *SimpleQueryStringQueryBuilder {
	rb.v.Query = query
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) QueryName_(queryname_ string) *SimpleQueryStringQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *SimpleQueryStringQueryBuilder) QuoteFieldSuffix(quotefieldsuffix string) *SimpleQueryStringQueryBuilder {
	rb.v.QuoteFieldSuffix = &quotefieldsuffix
	return rb
}
