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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

// MultiMatchQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L191-L217
type MultiMatchQuery struct {
	Analyzer                        *string                        `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                          `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                       `json:"boost,omitempty"`
	CutoffFrequency                 *float64                       `json:"cutoff_frequency,omitempty"`
	Fields                          *Fields                        `json:"fields,omitempty"`
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
	Slop                            *int                           `json:"slop,omitempty"`
	TieBreaker                      *float64                       `json:"tie_breaker,omitempty"`
	Type                            *textquerytype.TextQueryType   `json:"type,omitempty"`
	ZeroTermsQuery                  *zerotermsquery.ZeroTermsQuery `json:"zero_terms_query,omitempty"`
}

// MultiMatchQueryBuilder holds MultiMatchQuery struct and provides a builder API.
type MultiMatchQueryBuilder struct {
	v *MultiMatchQuery
}

// NewMultiMatchQuery provides a builder for the MultiMatchQuery struct.
func NewMultiMatchQueryBuilder() *MultiMatchQueryBuilder {
	r := MultiMatchQueryBuilder{
		&MultiMatchQuery{},
	}

	return &r
}

// Build finalize the chain and returns the MultiMatchQuery struct
func (rb *MultiMatchQueryBuilder) Build() MultiMatchQuery {
	return *rb.v
}

func (rb *MultiMatchQueryBuilder) Analyzer(analyzer string) *MultiMatchQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *MultiMatchQueryBuilder) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *MultiMatchQueryBuilder {
	rb.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery
	return rb
}

func (rb *MultiMatchQueryBuilder) Boost(boost float32) *MultiMatchQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *MultiMatchQueryBuilder) CutoffFrequency(cutofffrequency float64) *MultiMatchQueryBuilder {
	rb.v.CutoffFrequency = &cutofffrequency
	return rb
}

func (rb *MultiMatchQueryBuilder) Fields(fields *FieldsBuilder) *MultiMatchQueryBuilder {
	v := fields.Build()
	rb.v.Fields = &v
	return rb
}

func (rb *MultiMatchQueryBuilder) Fuzziness(fuzziness *FuzzinessBuilder) *MultiMatchQueryBuilder {
	v := fuzziness.Build()
	rb.v.Fuzziness = &v
	return rb
}

func (rb *MultiMatchQueryBuilder) FuzzyRewrite(fuzzyrewrite MultiTermQueryRewrite) *MultiMatchQueryBuilder {
	rb.v.FuzzyRewrite = &fuzzyrewrite
	return rb
}

func (rb *MultiMatchQueryBuilder) FuzzyTranspositions(fuzzytranspositions bool) *MultiMatchQueryBuilder {
	rb.v.FuzzyTranspositions = &fuzzytranspositions
	return rb
}

func (rb *MultiMatchQueryBuilder) Lenient(lenient bool) *MultiMatchQueryBuilder {
	rb.v.Lenient = &lenient
	return rb
}

func (rb *MultiMatchQueryBuilder) MaxExpansions(maxexpansions int) *MultiMatchQueryBuilder {
	rb.v.MaxExpansions = &maxexpansions
	return rb
}

func (rb *MultiMatchQueryBuilder) MinimumShouldMatch(minimumshouldmatch *MinimumShouldMatchBuilder) *MultiMatchQueryBuilder {
	v := minimumshouldmatch.Build()
	rb.v.MinimumShouldMatch = &v
	return rb
}

func (rb *MultiMatchQueryBuilder) Operator(operator operator.Operator) *MultiMatchQueryBuilder {
	rb.v.Operator = &operator
	return rb
}

func (rb *MultiMatchQueryBuilder) PrefixLength(prefixlength int) *MultiMatchQueryBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

func (rb *MultiMatchQueryBuilder) Query(query string) *MultiMatchQueryBuilder {
	rb.v.Query = query
	return rb
}

func (rb *MultiMatchQueryBuilder) QueryName_(queryname_ string) *MultiMatchQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *MultiMatchQueryBuilder) Slop(slop int) *MultiMatchQueryBuilder {
	rb.v.Slop = &slop
	return rb
}

func (rb *MultiMatchQueryBuilder) TieBreaker(tiebreaker float64) *MultiMatchQueryBuilder {
	rb.v.TieBreaker = &tiebreaker
	return rb
}

func (rb *MultiMatchQueryBuilder) Type_(type_ textquerytype.TextQueryType) *MultiMatchQueryBuilder {
	rb.v.Type = &type_
	return rb
}

func (rb *MultiMatchQueryBuilder) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *MultiMatchQueryBuilder {
	rb.v.ZeroTermsQuery = &zerotermsquery
	return rb
}
