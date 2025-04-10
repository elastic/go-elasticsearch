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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/textquerytype"
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

func (s *_queryStringQuery) AllowLeadingWildcard(allowleadingwildcard bool) *_queryStringQuery {

	s.v.AllowLeadingWildcard = &allowleadingwildcard

	return s
}

func (s *_queryStringQuery) AnalyzeWildcard(analyzewildcard bool) *_queryStringQuery {

	s.v.AnalyzeWildcard = &analyzewildcard

	return s
}

func (s *_queryStringQuery) Analyzer(analyzer string) *_queryStringQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_queryStringQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_queryStringQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

func (s *_queryStringQuery) Boost(boost float32) *_queryStringQuery {

	s.v.Boost = &boost

	return s
}

func (s *_queryStringQuery) DefaultField(field string) *_queryStringQuery {

	s.v.DefaultField = &field

	return s
}

func (s *_queryStringQuery) DefaultOperator(defaultoperator operator.Operator) *_queryStringQuery {

	s.v.DefaultOperator = &defaultoperator
	return s
}

func (s *_queryStringQuery) EnablePositionIncrements(enablepositionincrements bool) *_queryStringQuery {

	s.v.EnablePositionIncrements = &enablepositionincrements

	return s
}

func (s *_queryStringQuery) Escape(escape bool) *_queryStringQuery {

	s.v.Escape = &escape

	return s
}

func (s *_queryStringQuery) Fields(fields ...string) *_queryStringQuery {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

func (s *_queryStringQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_queryStringQuery {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

func (s *_queryStringQuery) FuzzyMaxExpansions(fuzzymaxexpansions int) *_queryStringQuery {

	s.v.FuzzyMaxExpansions = &fuzzymaxexpansions

	return s
}

func (s *_queryStringQuery) FuzzyPrefixLength(fuzzyprefixlength int) *_queryStringQuery {

	s.v.FuzzyPrefixLength = &fuzzyprefixlength

	return s
}

func (s *_queryStringQuery) FuzzyRewrite(multitermqueryrewrite string) *_queryStringQuery {

	s.v.FuzzyRewrite = &multitermqueryrewrite

	return s
}

func (s *_queryStringQuery) FuzzyTranspositions(fuzzytranspositions bool) *_queryStringQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

func (s *_queryStringQuery) Lenient(lenient bool) *_queryStringQuery {

	s.v.Lenient = &lenient

	return s
}

func (s *_queryStringQuery) MaxDeterminizedStates(maxdeterminizedstates int) *_queryStringQuery {

	s.v.MaxDeterminizedStates = &maxdeterminizedstates

	return s
}

func (s *_queryStringQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_queryStringQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_queryStringQuery) PhraseSlop(phraseslop types.Float64) *_queryStringQuery {

	s.v.PhraseSlop = &phraseslop

	return s
}

func (s *_queryStringQuery) Query(query string) *_queryStringQuery {

	s.v.Query = query

	return s
}

func (s *_queryStringQuery) QueryName_(queryname_ string) *_queryStringQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_queryStringQuery) QuoteAnalyzer(quoteanalyzer string) *_queryStringQuery {

	s.v.QuoteAnalyzer = &quoteanalyzer

	return s
}

func (s *_queryStringQuery) QuoteFieldSuffix(quotefieldsuffix string) *_queryStringQuery {

	s.v.QuoteFieldSuffix = &quotefieldsuffix

	return s
}

func (s *_queryStringQuery) Rewrite(multitermqueryrewrite string) *_queryStringQuery {

	s.v.Rewrite = &multitermqueryrewrite

	return s
}

func (s *_queryStringQuery) TieBreaker(tiebreaker types.Float64) *_queryStringQuery {

	s.v.TieBreaker = &tiebreaker

	return s
}

func (s *_queryStringQuery) TimeZone(timezone string) *_queryStringQuery {

	s.v.TimeZone = &timezone

	return s
}

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
