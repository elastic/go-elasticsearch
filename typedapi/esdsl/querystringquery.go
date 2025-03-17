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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
)

type _queryStringQuery struct {
	v *types.QueryStringQuery
}

// Returns documents based on a provided query string, using a parser with a
// strict syntax.
func NewQueryStringQuery(query string) *_queryStringQuery {

	tmp := &_queryStringQuery{v: types.NewQueryStringQuery()}

	tmp.Query(query)

	return tmp

}

// If `true`, the wildcard characters `*` and `?` are allowed as the first
// character of the query string.
func (s *_queryStringQuery) AllowLeadingWildcard(allowleadingwildcard bool) *_queryStringQuery {

	s.v.AllowLeadingWildcard = &allowleadingwildcard

	return s
}

// If `true`, the query attempts to analyze wildcard terms in the query string.
func (s *_queryStringQuery) AnalyzeWildcard(analyzewildcard bool) *_queryStringQuery {

	s.v.AnalyzeWildcard = &analyzewildcard

	return s
}

// Analyzer used to convert text in the query string into tokens.
func (s *_queryStringQuery) Analyzer(analyzer string) *_queryStringQuery {

	s.v.Analyzer = &analyzer

	return s
}

// If `true`, match phrase queries are automatically created for multi-term
// synonyms.
func (s *_queryStringQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_queryStringQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_queryStringQuery) Boost(boost float32) *_queryStringQuery {

	s.v.Boost = &boost

	return s
}

// Default field to search if no field is provided in the query string.
// Supports wildcards (`*`).
// Defaults to the `index.query.default_field` index setting, which has a
// default value of `*`.
func (s *_queryStringQuery) DefaultField(field string) *_queryStringQuery {

	s.v.DefaultField = &field

	return s
}

// Default boolean logic used to interpret text in the query string if no
// operators are specified.
func (s *_queryStringQuery) DefaultOperator(defaultoperator operator.Operator) *_queryStringQuery {

	s.v.DefaultOperator = &defaultoperator
	return s
}

// If `true`, enable position increments in queries constructed from a
// `query_string` search.
func (s *_queryStringQuery) EnablePositionIncrements(enablepositionincrements bool) *_queryStringQuery {

	s.v.EnablePositionIncrements = &enablepositionincrements

	return s
}

func (s *_queryStringQuery) Escape(escape bool) *_queryStringQuery {

	s.v.Escape = &escape

	return s
}

// Array of fields to search. Supports wildcards (`*`).
func (s *_queryStringQuery) Fields(fields ...string) *_queryStringQuery {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

// Maximum edit distance allowed for fuzzy matching.
func (s *_queryStringQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_queryStringQuery {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

// Maximum number of terms to which the query expands for fuzzy matching.
func (s *_queryStringQuery) FuzzyMaxExpansions(fuzzymaxexpansions int) *_queryStringQuery {

	s.v.FuzzyMaxExpansions = &fuzzymaxexpansions

	return s
}

// Number of beginning characters left unchanged for fuzzy matching.
func (s *_queryStringQuery) FuzzyPrefixLength(fuzzyprefixlength int) *_queryStringQuery {

	s.v.FuzzyPrefixLength = &fuzzyprefixlength

	return s
}

// Method used to rewrite the query.
func (s *_queryStringQuery) FuzzyRewrite(multitermqueryrewrite string) *_queryStringQuery {

	s.v.FuzzyRewrite = &multitermqueryrewrite

	return s
}

// If `true`, edits for fuzzy matching include transpositions of two adjacent
// characters (for example, `ab` to `ba`).
func (s *_queryStringQuery) FuzzyTranspositions(fuzzytranspositions bool) *_queryStringQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

// If `true`, format-based errors, such as providing a text value for a numeric
// field, are ignored.
func (s *_queryStringQuery) Lenient(lenient bool) *_queryStringQuery {

	s.v.Lenient = &lenient

	return s
}

// Maximum number of automaton states required for the query.
func (s *_queryStringQuery) MaxDeterminizedStates(maxdeterminizedstates int) *_queryStringQuery {

	s.v.MaxDeterminizedStates = &maxdeterminizedstates

	return s
}

// Minimum number of clauses that must match for a document to be returned.
func (s *_queryStringQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_queryStringQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

// Maximum number of positions allowed between matching tokens for phrases.
func (s *_queryStringQuery) PhraseSlop(phraseslop types.Float64) *_queryStringQuery {

	s.v.PhraseSlop = &phraseslop

	return s
}

// Query string you wish to parse and use for search.
func (s *_queryStringQuery) Query(query string) *_queryStringQuery {

	s.v.Query = query

	return s
}

func (s *_queryStringQuery) QueryName_(queryname_ string) *_queryStringQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Analyzer used to convert quoted text in the query string into tokens.
// For quoted text, this parameter overrides the analyzer specified in the
// `analyzer` parameter.
func (s *_queryStringQuery) QuoteAnalyzer(quoteanalyzer string) *_queryStringQuery {

	s.v.QuoteAnalyzer = &quoteanalyzer

	return s
}

// Suffix appended to quoted text in the query string.
// You can use this suffix to use a different analysis method for exact matches.
func (s *_queryStringQuery) QuoteFieldSuffix(quotefieldsuffix string) *_queryStringQuery {

	s.v.QuoteFieldSuffix = &quotefieldsuffix

	return s
}

// Method used to rewrite the query.
func (s *_queryStringQuery) Rewrite(multitermqueryrewrite string) *_queryStringQuery {

	s.v.Rewrite = &multitermqueryrewrite

	return s
}

// How to combine the queries generated from the individual search terms in the
// resulting `dis_max` query.
func (s *_queryStringQuery) TieBreaker(tiebreaker types.Float64) *_queryStringQuery {

	s.v.TieBreaker = &tiebreaker

	return s
}

// Coordinated Universal Time (UTC) offset or IANA time zone used to convert
// date values in the query string to UTC.
func (s *_queryStringQuery) TimeZone(timezone string) *_queryStringQuery {

	s.v.TimeZone = &timezone

	return s
}

// Determines how the query matches and scores documents.
func (s *_queryStringQuery) Type(type_ textquerytype.TextQueryType) *_queryStringQuery {

	s.v.Type = &type_
	return s
}

func (s *_queryStringQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.QueryString = s.v

	return container
}

func (s *_queryStringQuery) QueryStringQueryCaster() *types.QueryStringQuery {
	return s.v
}
