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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// MatchQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L133-L158
type MatchQuery struct {
	Analyzer                        *string                        `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                          `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                       `json:"boost,omitempty"`
	CutoffFrequency                 *float64                       `json:"cutoff_frequency,omitempty"`
	Fuzziness                       *Fuzziness                     `json:"fuzziness,omitempty"`
	FuzzyRewrite                    *MultiTermQueryRewrite         `json:"fuzzy_rewrite,omitempty"`
	FuzzyTranspositions             *bool                          `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool                          `json:"lenient,omitempty"`
	MaxExpansions                   *int                           `json:"max_expansions,omitempty"`
	MinimumShouldMatch              *MinimumShouldMatch            `json:"minimum_should_match,omitempty"`
	Operator                        *operator.Operator             `json:"operator,omitempty"`
	PrefixLength                    *int                           `json:"prefix_length,omitempty"`
	Query                           string                         `json:"query"`
	QueryName_                      *string                        `json:"_name,omitempty"`
	ZeroTermsQuery                  *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

// MatchQueryBuilder holds MatchQuery struct and provides a builder API.
type MatchQueryBuilder struct {
	v *MatchQuery
}

// NewMatchQuery provides a builder for the MatchQuery struct.
func NewMatchQueryBuilder() *MatchQueryBuilder {
	r := MatchQueryBuilder{
		&MatchQuery{},
	}

	return &r
}

// Build finalize the chain and returns the MatchQuery struct
func (rb *MatchQueryBuilder) Build() MatchQuery {
	return *rb.v
}

func (rb *MatchQueryBuilder) Analyzer(analyzer string) *MatchQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *MatchQueryBuilder) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *MatchQueryBuilder {
	rb.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery
	return rb
}

func (rb *MatchQueryBuilder) Boost(boost float32) *MatchQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *MatchQueryBuilder) CutoffFrequency(cutofffrequency float64) *MatchQueryBuilder {
	rb.v.CutoffFrequency = &cutofffrequency
	return rb
}

func (rb *MatchQueryBuilder) Fuzziness(fuzziness *FuzzinessBuilder) *MatchQueryBuilder {
	v := fuzziness.Build()
	rb.v.Fuzziness = &v
	return rb
}

func (rb *MatchQueryBuilder) FuzzyRewrite(fuzzyrewrite MultiTermQueryRewrite) *MatchQueryBuilder {
	rb.v.FuzzyRewrite = &fuzzyrewrite
	return rb
}

func (rb *MatchQueryBuilder) FuzzyTranspositions(fuzzytranspositions bool) *MatchQueryBuilder {
	rb.v.FuzzyTranspositions = &fuzzytranspositions
	return rb
}

func (rb *MatchQueryBuilder) Lenient(lenient bool) *MatchQueryBuilder {
	rb.v.Lenient = &lenient
	return rb
}

func (rb *MatchQueryBuilder) MaxExpansions(maxexpansions int) *MatchQueryBuilder {
	rb.v.MaxExpansions = &maxexpansions
	return rb
}

func (rb *MatchQueryBuilder) MinimumShouldMatch(minimumshouldmatch *MinimumShouldMatchBuilder) *MatchQueryBuilder {
	v := minimumshouldmatch.Build()
	rb.v.MinimumShouldMatch = &v
	return rb
}

func (rb *MatchQueryBuilder) Operator(operator operator.Operator) *MatchQueryBuilder {
	rb.v.Operator = &operator
	return rb
}

func (rb *MatchQueryBuilder) PrefixLength(prefixlength int) *MatchQueryBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

func (rb *MatchQueryBuilder) Query(arg string) *MatchQueryBuilder {
	rb.v.Query = arg
	return rb
}

func (rb *MatchQueryBuilder) QueryName_(queryname_ string) *MatchQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *MatchQueryBuilder) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *MatchQueryBuilder {
	rb.v.ZeroTermsQuery = &zerotermsquery
	return rb
}
