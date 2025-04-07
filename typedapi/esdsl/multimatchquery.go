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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

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

func (s *_multiMatchQuery) Analyzer(analyzer string) *_multiMatchQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_multiMatchQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_multiMatchQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

func (s *_multiMatchQuery) Boost(boost float32) *_multiMatchQuery {

	s.v.Boost = &boost

	return s
}

func (s *_multiMatchQuery) CutoffFrequency(cutofffrequency types.Float64) *_multiMatchQuery {

	s.v.CutoffFrequency = &cutofffrequency

	return s
}

func (s *_multiMatchQuery) Fields(fields ...string) *_multiMatchQuery {

	s.v.Fields = fields

	return s
}

func (s *_multiMatchQuery) Fuzziness(fuzziness types.FuzzinessVariant) *_multiMatchQuery {

	s.v.Fuzziness = *fuzziness.FuzzinessCaster()

	return s
}

func (s *_multiMatchQuery) FuzzyRewrite(multitermqueryrewrite string) *_multiMatchQuery {

	s.v.FuzzyRewrite = &multitermqueryrewrite

	return s
}

func (s *_multiMatchQuery) FuzzyTranspositions(fuzzytranspositions bool) *_multiMatchQuery {

	s.v.FuzzyTranspositions = &fuzzytranspositions

	return s
}

func (s *_multiMatchQuery) Lenient(lenient bool) *_multiMatchQuery {

	s.v.Lenient = &lenient

	return s
}

func (s *_multiMatchQuery) MaxExpansions(maxexpansions int) *_multiMatchQuery {

	s.v.MaxExpansions = &maxexpansions

	return s
}

func (s *_multiMatchQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_multiMatchQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_multiMatchQuery) Operator(operator operator.Operator) *_multiMatchQuery {

	s.v.Operator = &operator
	return s
}

func (s *_multiMatchQuery) PrefixLength(prefixlength int) *_multiMatchQuery {

	s.v.PrefixLength = &prefixlength

	return s
}

func (s *_multiMatchQuery) Query(query string) *_multiMatchQuery {

	s.v.Query = query

	return s
}

func (s *_multiMatchQuery) QueryName_(queryname_ string) *_multiMatchQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_multiMatchQuery) Slop(slop int) *_multiMatchQuery {

	s.v.Slop = &slop

	return s
}

func (s *_multiMatchQuery) TieBreaker(tiebreaker types.Float64) *_multiMatchQuery {

	s.v.TieBreaker = &tiebreaker

	return s
}

func (s *_multiMatchQuery) Type(type_ textquerytype.TextQueryType) *_multiMatchQuery {

	s.v.Type = &type_
	return s
}

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
