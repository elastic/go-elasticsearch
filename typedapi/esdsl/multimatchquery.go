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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/textquerytype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

type _multiMatchQuery struct {
	v *types.MultiMatchQuery
}

// Enables you to search for a provided text, number, date or boolean value
// across multiple fields.
// The provided text is analyzed before matching.
func NewMultiMatchQuery(query string) *_multiMatchQuery {

	tmp := &_multiMatchQuery{v: types.NewMultiMatchQuery()}

	tmp.Query(query)

	return tmp

}

// Analyzer used to convert the text in the query value into tokens.
func (s *_multiMatchQuery) Analyzer(analyzer string) *_multiMatchQuery {

	s.v.Analyzer = &analyzer

	return s
}

// If `true`, match phrase queries are automatically created for multi-term
// synonyms.
func (s *_multiMatchQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_multiMatchQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_multiMatchQuery) Boost(boost float32) *_multiMatchQuery {

	s.v.Boost = &boost

	return s
}

func (s *_multiMatchQuery) CutoffFrequency(cutofffrequency types.Float64) *_multiMatchQuery {

	s.v.CutoffFrequency = &cutofffrequency

	return s
}

// The fields to be queried.
// Defaults to the `index.query.default_field` index settings, which in turn
// defaults to `*`.
func (s *_multiMatchQuery) Fields(fields ...string) *_multiMatchQuery {

	s.v.Fields = fields

	return s
}

// Maximum edit distance allowed for matching.
func (s *_multiMatchQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_multiMatchQuery {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

// Method used to rewrite the query.
func (s *_multiMatchQuery) FuzzyRewrite(multitermqueryrewrite string) *_multiMatchQuery {

	s.v.FuzzyRewrite = &multitermqueryrewrite

	return s
}

// If `true`, edits for fuzzy matching include transpositions of two adjacent
// characters (for example, `ab` to `ba`).
// Can be applied to the term subqueries constructed for all terms but the final
// term.
func (s *_multiMatchQuery) FuzzyTranspositions(fuzzytranspositions bool) *_multiMatchQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

// If `true`, format-based errors, such as providing a text query value for a
// numeric field, are ignored.
func (s *_multiMatchQuery) Lenient(lenient bool) *_multiMatchQuery {

	s.v.Lenient = &lenient

	return s
}

// Maximum number of terms to which the query will expand.
func (s *_multiMatchQuery) MaxExpansions(maxexpansions int) *_multiMatchQuery {

	s.v.MaxExpansions = &maxexpansions

	return s
}

// Minimum number of clauses that must match for a document to be returned.
func (s *_multiMatchQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_multiMatchQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

// Boolean logic used to interpret text in the query value.
func (s *_multiMatchQuery) Operator(operator operator.Operator) *_multiMatchQuery {

	s.v.Operator = &operator
	return s
}

// Number of beginning characters left unchanged for fuzzy matching.
func (s *_multiMatchQuery) PrefixLength(prefixlength int) *_multiMatchQuery {

	s.v.PrefixLength = &prefixlength

	return s
}

// Text, number, boolean value or date you wish to find in the provided field.
func (s *_multiMatchQuery) Query(query string) *_multiMatchQuery {

	s.v.Query = query

	return s
}

func (s *_multiMatchQuery) QueryName_(queryname_ string) *_multiMatchQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Maximum number of positions allowed between matching tokens.
func (s *_multiMatchQuery) Slop(slop int) *_multiMatchQuery {

	s.v.Slop = &slop

	return s
}

// Determines how scores for each per-term blended query and scores across
// groups are combined.
func (s *_multiMatchQuery) TieBreaker(tiebreaker types.Float64) *_multiMatchQuery {

	s.v.TieBreaker = &tiebreaker

	return s
}

// How `the` multi_match query is executed internally.
func (s *_multiMatchQuery) Type(type_ textquerytype.TextQueryType) *_multiMatchQuery {

	s.v.Type = &type_
	return s
}

// Indicates whether no documents are returned if the `analyzer` removes all
// tokens, such as when using a `stop` filter.
func (s *_multiMatchQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_multiMatchQuery {

	s.v.ZeroTermsQuery = &zerotermsquery
	return s
}

func (s *_multiMatchQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.MultiMatch = s.v

	return container
}

func (s *_multiMatchQuery) MultiMatchQueryCaster() *types.MultiMatchQuery {
	return s.v
}
