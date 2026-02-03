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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
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

func (s *_simpleQueryStringQuery) AnalyzeWildcard(analyzewildcard bool) *_simpleQueryStringQuery {

	s.v.AnalyzeWildcard = &analyzewildcard

	return s
}

func (s *_simpleQueryStringQuery) Analyzer(analyzer string) *_simpleQueryStringQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_simpleQueryStringQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_simpleQueryStringQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

func (s *_simpleQueryStringQuery) DefaultOperator(defaultoperator operator.Operator) *_simpleQueryStringQuery {

	s.v.DefaultOperator = &defaultoperator
	return s
}

func (s *_simpleQueryStringQuery) Fields(fields ...string) *_simpleQueryStringQuery {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

func (s *_simpleQueryStringQuery) Flags(simplequerystringflags types.PipeSeparatedFlagsSimpleQueryStringFlag) *_simpleQueryStringQuery {

	s.v.Flags = simplequerystringflags

	return s
}

func (s *_simpleQueryStringQuery) FuzzyMaxExpansions(fuzzymaxexpansions int) *_simpleQueryStringQuery {

	s.v.FuzzyMaxExpansions = &fuzzymaxexpansions

	return s
}

func (s *_simpleQueryStringQuery) FuzzyPrefixLength(fuzzyprefixlength int) *_simpleQueryStringQuery {

	s.v.FuzzyPrefixLength = &fuzzyprefixlength

	return s
}

func (s *_simpleQueryStringQuery) FuzzyTranspositions(fuzzytranspositions bool) *_simpleQueryStringQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

func (s *_simpleQueryStringQuery) Lenient(lenient bool) *_simpleQueryStringQuery {

	s.v.Lenient = &lenient

	return s
}

func (s *_simpleQueryStringQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_simpleQueryStringQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_simpleQueryStringQuery) Query(query string) *_simpleQueryStringQuery {

	s.v.Query = query

	return s
}

func (s *_simpleQueryStringQuery) QuoteFieldSuffix(quotefieldsuffix string) *_simpleQueryStringQuery {

	s.v.QuoteFieldSuffix = &quotefieldsuffix

	return s
}

func (s *_simpleQueryStringQuery) Boost(boost float32) *_simpleQueryStringQuery {

	s.v.Boost = &boost

	return s
}

func (s *_simpleQueryStringQuery) QueryName_(queryname_ string) *_simpleQueryStringQuery {

	s.v.QueryName_ = &queryname_

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
