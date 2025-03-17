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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

type _matchQuery struct {
	k string
	v *types.MatchQuery
}

// Returns roles that match a provided text, number, date or boolean value.
// The provided text is analyzed before matching.
func NewMatchQuery(field string, query string) *_matchQuery {
	tmp := &_matchQuery{
		k: field,
		v: types.NewMatchQuery(),
	}

	tmp.Query(query)
	return tmp
}

// Analyzer used to convert the text in the query value into tokens.
func (s *_matchQuery) Analyzer(analyzer string) *_matchQuery {

	s.v.Analyzer = &analyzer

	return s
}

// If `true`, match phrase queries are automatically created for multi-term
// synonyms.
func (s *_matchQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_matchQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_matchQuery) Boost(boost float32) *_matchQuery {

	s.v.Boost = &boost

	return s
}

func (s *_matchQuery) CutoffFrequency(cutofffrequency types.Float64) *_matchQuery {

	s.v.CutoffFrequency = &cutofffrequency

	return s
}

// Maximum edit distance allowed for matching.
func (s *_matchQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_matchQuery {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

// Method used to rewrite the query.
func (s *_matchQuery) FuzzyRewrite(multitermqueryrewrite string) *_matchQuery {

	s.v.FuzzyRewrite = &multitermqueryrewrite

	return s
}

// If `true`, edits for fuzzy matching include transpositions of two adjacent
// characters (for example, `ab` to `ba`).
func (s *_matchQuery) FuzzyTranspositions(fuzzytranspositions bool) *_matchQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

// If `true`, format-based errors, such as providing a text query value for a
// numeric field, are ignored.
func (s *_matchQuery) Lenient(lenient bool) *_matchQuery {

	s.v.Lenient = &lenient

	return s
}

// Maximum number of terms to which the query will expand.
func (s *_matchQuery) MaxExpansions(maxexpansions int) *_matchQuery {

	s.v.MaxExpansions = &maxexpansions

	return s
}

// Minimum number of clauses that must match for a document to be returned.
func (s *_matchQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_matchQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

// Boolean logic used to interpret text in the query value.
func (s *_matchQuery) Operator(operator operator.Operator) *_matchQuery {

	s.v.Operator = &operator
	return s
}

// Number of beginning characters left unchanged for fuzzy matching.
func (s *_matchQuery) PrefixLength(prefixlength int) *_matchQuery {

	s.v.PrefixLength = &prefixlength

	return s
}

// Text, number, boolean value or date you wish to find in the provided field.
func (s *_matchQuery) Query(query string) *_matchQuery {

	s.v.Query = query

	return s
}

func (s *_matchQuery) QueryName_(queryname_ string) *_matchQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Indicates whether no documents are returned if the `analyzer` removes all
// tokens, such as when using a `stop` filter.
func (s *_matchQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_matchQuery {

	s.v.ZeroTermsQuery = &zerotermsquery
	return s
}

func (s *_matchQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Match = map[string]types.MatchQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_matchQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Match = map[string]types.MatchQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_matchQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Match = map[string]types.MatchQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_matchQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Match = map[string]types.MatchQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleMatchQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleMatchQuery() *_matchQuery {
	return &_matchQuery{
		k: "",
		v: types.NewMatchQuery(),
	}
}

func (s *_matchQuery) MatchQueryCaster() *types.MatchQuery {
	return s.v.MatchQueryCaster()
}
