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
)

// QueryStringQuery type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/fulltext.ts#L233-L269
type QueryStringQuery struct {
	AllowLeadingWildcard            *bool                        `json:"allow_leading_wildcard,omitempty"`
	AnalyzeWildcard                 *bool                        `json:"analyze_wildcard,omitempty"`
	Analyzer                        *string                      `json:"analyzer,omitempty"`
	AutoGenerateSynonymsPhraseQuery *bool                        `json:"auto_generate_synonyms_phrase_query,omitempty"`
	Boost                           *float32                     `json:"boost,omitempty"`
	DefaultField                    *Field                       `json:"default_field,omitempty"`
	DefaultOperator                 *operator.Operator           `json:"default_operator,omitempty"`
	EnablePositionIncrements        *bool                        `json:"enable_position_increments,omitempty"`
	Escape                          *bool                        `json:"escape,omitempty"`
	Fields                          []Field                      `json:"fields,omitempty"`
	Fuzziness                       *Fuzziness                   `json:"fuzziness,omitempty"`
	FuzzyMaxExpansions              *int                         `json:"fuzzy_max_expansions,omitempty"`
	FuzzyPrefixLength               *int                         `json:"fuzzy_prefix_length,omitempty"`
	FuzzyRewrite                    *MultiTermQueryRewrite       `json:"fuzzy_rewrite,omitempty"`
	FuzzyTranspositions             *bool                        `json:"fuzzy_transpositions,omitempty"`
	Lenient                         *bool                        `json:"lenient,omitempty"`
	MaxDeterminizedStates           *int                         `json:"max_determinized_states,omitempty"`
	MinimumShouldMatch              *MinimumShouldMatch          `json:"minimum_should_match,omitempty"`
	PhraseSlop                      *float64                     `json:"phrase_slop,omitempty"`
	Query                           string                       `json:"query"`
	QueryName_                      *string                      `json:"_name,omitempty"`
	QuoteAnalyzer                   *string                      `json:"quote_analyzer,omitempty"`
	QuoteFieldSuffix                *string                      `json:"quote_field_suffix,omitempty"`
	Rewrite                         *MultiTermQueryRewrite       `json:"rewrite,omitempty"`
	TieBreaker                      *float64                     `json:"tie_breaker,omitempty"`
	TimeZone                        *TimeZone                    `json:"time_zone,omitempty"`
	Type                            *textquerytype.TextQueryType `json:"type,omitempty"`
}

// QueryStringQueryBuilder holds QueryStringQuery struct and provides a builder API.
type QueryStringQueryBuilder struct {
	v *QueryStringQuery
}

// NewQueryStringQuery provides a builder for the QueryStringQuery struct.
func NewQueryStringQueryBuilder() *QueryStringQueryBuilder {
	r := QueryStringQueryBuilder{
		&QueryStringQuery{},
	}

	return &r
}

// Build finalize the chain and returns the QueryStringQuery struct
func (rb *QueryStringQueryBuilder) Build() QueryStringQuery {
	return *rb.v
}

func (rb *QueryStringQueryBuilder) AllowLeadingWildcard(allowleadingwildcard bool) *QueryStringQueryBuilder {
	rb.v.AllowLeadingWildcard = &allowleadingwildcard
	return rb
}

func (rb *QueryStringQueryBuilder) AnalyzeWildcard(analyzewildcard bool) *QueryStringQueryBuilder {
	rb.v.AnalyzeWildcard = &analyzewildcard
	return rb
}

func (rb *QueryStringQueryBuilder) Analyzer(analyzer string) *QueryStringQueryBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *QueryStringQueryBuilder) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *QueryStringQueryBuilder {
	rb.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery
	return rb
}

func (rb *QueryStringQueryBuilder) Boost(boost float32) *QueryStringQueryBuilder {
	rb.v.Boost = &boost
	return rb
}

func (rb *QueryStringQueryBuilder) DefaultField(defaultfield Field) *QueryStringQueryBuilder {
	rb.v.DefaultField = &defaultfield
	return rb
}

func (rb *QueryStringQueryBuilder) DefaultOperator(defaultoperator operator.Operator) *QueryStringQueryBuilder {
	rb.v.DefaultOperator = &defaultoperator
	return rb
}

func (rb *QueryStringQueryBuilder) EnablePositionIncrements(enablepositionincrements bool) *QueryStringQueryBuilder {
	rb.v.EnablePositionIncrements = &enablepositionincrements
	return rb
}

func (rb *QueryStringQueryBuilder) Escape(escape bool) *QueryStringQueryBuilder {
	rb.v.Escape = &escape
	return rb
}

func (rb *QueryStringQueryBuilder) Fields(fields ...Field) *QueryStringQueryBuilder {
	rb.v.Fields = fields
	return rb
}

func (rb *QueryStringQueryBuilder) Fuzziness(fuzziness *FuzzinessBuilder) *QueryStringQueryBuilder {
	v := fuzziness.Build()
	rb.v.Fuzziness = &v
	return rb
}

func (rb *QueryStringQueryBuilder) FuzzyMaxExpansions(fuzzymaxexpansions int) *QueryStringQueryBuilder {
	rb.v.FuzzyMaxExpansions = &fuzzymaxexpansions
	return rb
}

func (rb *QueryStringQueryBuilder) FuzzyPrefixLength(fuzzyprefixlength int) *QueryStringQueryBuilder {
	rb.v.FuzzyPrefixLength = &fuzzyprefixlength
	return rb
}

func (rb *QueryStringQueryBuilder) FuzzyRewrite(fuzzyrewrite MultiTermQueryRewrite) *QueryStringQueryBuilder {
	rb.v.FuzzyRewrite = &fuzzyrewrite
	return rb
}

func (rb *QueryStringQueryBuilder) FuzzyTranspositions(fuzzytranspositions bool) *QueryStringQueryBuilder {
	rb.v.FuzzyTranspositions = &fuzzytranspositions
	return rb
}

func (rb *QueryStringQueryBuilder) Lenient(lenient bool) *QueryStringQueryBuilder {
	rb.v.Lenient = &lenient
	return rb
}

func (rb *QueryStringQueryBuilder) MaxDeterminizedStates(maxdeterminizedstates int) *QueryStringQueryBuilder {
	rb.v.MaxDeterminizedStates = &maxdeterminizedstates
	return rb
}

func (rb *QueryStringQueryBuilder) MinimumShouldMatch(minimumshouldmatch *MinimumShouldMatchBuilder) *QueryStringQueryBuilder {
	v := minimumshouldmatch.Build()
	rb.v.MinimumShouldMatch = &v
	return rb
}

func (rb *QueryStringQueryBuilder) PhraseSlop(phraseslop float64) *QueryStringQueryBuilder {
	rb.v.PhraseSlop = &phraseslop
	return rb
}

func (rb *QueryStringQueryBuilder) Query(query string) *QueryStringQueryBuilder {
	rb.v.Query = query
	return rb
}

func (rb *QueryStringQueryBuilder) QueryName_(queryname_ string) *QueryStringQueryBuilder {
	rb.v.QueryName_ = &queryname_
	return rb
}

func (rb *QueryStringQueryBuilder) QuoteAnalyzer(quoteanalyzer string) *QueryStringQueryBuilder {
	rb.v.QuoteAnalyzer = &quoteanalyzer
	return rb
}

func (rb *QueryStringQueryBuilder) QuoteFieldSuffix(quotefieldsuffix string) *QueryStringQueryBuilder {
	rb.v.QuoteFieldSuffix = &quotefieldsuffix
	return rb
}

func (rb *QueryStringQueryBuilder) Rewrite(rewrite MultiTermQueryRewrite) *QueryStringQueryBuilder {
	rb.v.Rewrite = &rewrite
	return rb
}

func (rb *QueryStringQueryBuilder) TieBreaker(tiebreaker float64) *QueryStringQueryBuilder {
	rb.v.TieBreaker = &tiebreaker
	return rb
}

func (rb *QueryStringQueryBuilder) TimeZone(timezone TimeZone) *QueryStringQueryBuilder {
	rb.v.TimeZone = &timezone
	return rb
}

func (rb *QueryStringQueryBuilder) Type_(type_ textquerytype.TextQueryType) *QueryStringQueryBuilder {
	rb.v.Type = &type_
	return rb
}
