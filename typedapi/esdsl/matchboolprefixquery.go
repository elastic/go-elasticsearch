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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type _matchBoolPrefixQuery struct {
	k string
	v *types.MatchBoolPrefixQuery
}

// Analyzes its input and constructs a `bool` query from the terms.
// Each term except the last is used in a `term` query.
// The last term is used in a prefix query.
func NewMatchBoolPrefixQuery(field string, query string) *_matchBoolPrefixQuery {
	tmp := &_matchBoolPrefixQuery{
		k: field,
		v: types.NewMatchBoolPrefixQuery(),
	}

	tmp.Query(query)
	return tmp
}

func (s *_matchBoolPrefixQuery) Analyzer(analyzer string) *_matchBoolPrefixQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_matchBoolPrefixQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_matchBoolPrefixQuery {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

func (s *_matchBoolPrefixQuery) FuzzyRewrite(multitermqueryrewrite string) *_matchBoolPrefixQuery {

	s.v.FuzzyRewrite = &multitermqueryrewrite

	return s
}

func (s *_matchBoolPrefixQuery) FuzzyTranspositions(fuzzytranspositions bool) *_matchBoolPrefixQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

func (s *_matchBoolPrefixQuery) MaxExpansions(maxexpansions int) *_matchBoolPrefixQuery {

	s.v.MaxExpansions = &maxexpansions

	return s
}

func (s *_matchBoolPrefixQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_matchBoolPrefixQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_matchBoolPrefixQuery) Operator(operator operator.Operator) *_matchBoolPrefixQuery {

	s.v.Operator = &operator
	return s
}

func (s *_matchBoolPrefixQuery) PrefixLength(prefixlength int) *_matchBoolPrefixQuery {

	s.v.PrefixLength = &prefixlength

	return s
}

func (s *_matchBoolPrefixQuery) Query(query string) *_matchBoolPrefixQuery {

	s.v.Query = query

	return s
}

func (s *_matchBoolPrefixQuery) Boost(boost float32) *_matchBoolPrefixQuery {

	s.v.Boost = &boost

	return s
}

func (s *_matchBoolPrefixQuery) QueryName_(queryname_ string) *_matchBoolPrefixQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_matchBoolPrefixQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.MatchBoolPrefix = map[string]types.MatchBoolPrefixQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleMatchBoolPrefixQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleMatchBoolPrefixQuery() *_matchBoolPrefixQuery {
	return &_matchBoolPrefixQuery{
		k: "",
		v: types.NewMatchBoolPrefixQuery(),
	}
}

func (s *_matchBoolPrefixQuery) MatchBoolPrefixQueryCaster() *types.MatchBoolPrefixQuery {
	return s.v.MatchBoolPrefixQueryCaster()
}
