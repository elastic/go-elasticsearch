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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
)

type _simpleQueryStringQuery struct {
	v *types.SimpleQueryStringQuery
}

// Returns documents based on a provided query string, using a parser with a
// limited but fault-tolerant syntax.
func NewSimpleQueryStringQuery(query string) *_simpleQueryStringQuery {

	tmp := &_simpleQueryStringQuery{v: types.NewSimpleQueryStringQuery()}

	tmp.Query(query)

	return tmp

}

// If `true`, the query attempts to analyze wildcard terms in the query string.
func (s *_simpleQueryStringQuery) AnalyzeWildcard(analyzewildcard bool) *_simpleQueryStringQuery {

	s.v.AnalyzeWildcard = &analyzewildcard

	return s
}

// Analyzer used to convert text in the query string into tokens.
func (s *_simpleQueryStringQuery) Analyzer(analyzer string) *_simpleQueryStringQuery {

	s.v.Analyzer = &analyzer

	return s
}

// If `true`, the parser creates a match_phrase query for each multi-position
// token.
func (s *_simpleQueryStringQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_simpleQueryStringQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_simpleQueryStringQuery) Boost(boost float32) *_simpleQueryStringQuery {

	s.v.Boost = &boost

	return s
}

// Default boolean logic used to interpret text in the query string if no
// operators are specified.
func (s *_simpleQueryStringQuery) DefaultOperator(defaultoperator operator.Operator) *_simpleQueryStringQuery {

	s.v.DefaultOperator = &defaultoperator
	return s
}

// Array of fields you wish to search.
// Accepts wildcard expressions.
// You also can boost relevance scores for matches to particular fields using a
// caret (`^`) notation.
// Defaults to the `index.query.default_field index` setting, which has a
// default value of `*`.
func (s *_simpleQueryStringQuery) Fields(fields ...string) *_simpleQueryStringQuery {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

// List of enabled operators for the simple query string syntax.
func (s *_simpleQueryStringQuery) Flags(simplequerystringflags types.PipeSeparatedFlagsSimpleQueryStringFlag) *_simpleQueryStringQuery {

	s.v.Flags = simplequerystringflags

	return s
}

// Maximum number of terms to which the query expands for fuzzy matching.
func (s *_simpleQueryStringQuery) FuzzyMaxExpansions(fuzzymaxexpansions int) *_simpleQueryStringQuery {

	s.v.FuzzyMaxExpansions = &fuzzymaxexpansions

	return s
}

// Number of beginning characters left unchanged for fuzzy matching.
func (s *_simpleQueryStringQuery) FuzzyPrefixLength(fuzzyprefixlength int) *_simpleQueryStringQuery {

	s.v.FuzzyPrefixLength = &fuzzyprefixlength

	return s
}

// If `true`, edits for fuzzy matching include transpositions of two adjacent
// characters (for example, `ab` to `ba`).
func (s *_simpleQueryStringQuery) FuzzyTranspositions(fuzzytranspositions bool) *_simpleQueryStringQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

// If `true`, format-based errors, such as providing a text value for a numeric
// field, are ignored.
func (s *_simpleQueryStringQuery) Lenient(lenient bool) *_simpleQueryStringQuery {

	s.v.Lenient = &lenient

	return s
}

// Minimum number of clauses that must match for a document to be returned.
func (s *_simpleQueryStringQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_simpleQueryStringQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

// Query string in the simple query string syntax you wish to parse and use for
// search.
func (s *_simpleQueryStringQuery) Query(query string) *_simpleQueryStringQuery {

	s.v.Query = query

	return s
}

func (s *_simpleQueryStringQuery) QueryName_(queryname_ string) *_simpleQueryStringQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Suffix appended to quoted text in the query string.
func (s *_simpleQueryStringQuery) QuoteFieldSuffix(quotefieldsuffix string) *_simpleQueryStringQuery {

	s.v.QuoteFieldSuffix = &quotefieldsuffix

	return s
}

func (s *_simpleQueryStringQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.SimpleQueryString = s.v

	return container
}

func (s *_simpleQueryStringQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()

	container.SimpleQueryString = s.v

	return container
}

func (s *_simpleQueryStringQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()

	container.SimpleQueryString = s.v

	return container
}

func (s *_simpleQueryStringQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()

	container.SimpleQueryString = s.v

	return container
}

func (s *_simpleQueryStringQuery) SimpleQueryStringQueryCaster() *types.SimpleQueryStringQuery {
	return s.v
}
