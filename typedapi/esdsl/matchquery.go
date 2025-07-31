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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
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

func (s *_matchQuery) Analyzer(analyzer string) *_matchQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_matchQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_matchQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

func (s *_matchQuery) Boost(boost float32) *_matchQuery {

	s.v.Boost = &boost

	return s
}

func (s *_matchQuery) CutoffFrequency(cutofffrequency types.Float64) *_matchQuery {

	s.v.CutoffFrequency = &cutofffrequency

	return s
}

func (s *_matchQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_matchQuery {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

func (s *_matchQuery) FuzzyRewrite(multitermqueryrewrite string) *_matchQuery {

	s.v.FuzzyRewrite = &multitermqueryrewrite

	return s
}

func (s *_matchQuery) FuzzyTranspositions(fuzzytranspositions bool) *_matchQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

func (s *_matchQuery) Lenient(lenient bool) *_matchQuery {

	s.v.Lenient = &lenient

	return s
}

func (s *_matchQuery) MaxExpansions(maxexpansions int) *_matchQuery {

	s.v.MaxExpansions = &maxexpansions

	return s
}

func (s *_matchQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_matchQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_matchQuery) Operator(operator operator.Operator) *_matchQuery {

	s.v.Operator = &operator
	return s
}

func (s *_matchQuery) PrefixLength(prefixlength int) *_matchQuery {

	s.v.PrefixLength = &prefixlength

	return s
}

func (s *_matchQuery) Query(query string) *_matchQuery {

	s.v.Query = query

	return s
}

func (s *_matchQuery) QueryName_(queryname_ string) *_matchQuery {

	s.v.QueryName_ = &queryname_

	return s
}

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
