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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/combinedfieldsoperator"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/combinedfieldszeroterms"
)

type _combinedFieldsQuery struct {
	v *types.CombinedFieldsQuery
}

// The `combined_fields` query supports searching multiple text fields as if
// their contents had been indexed into one combined field.
func NewCombinedFieldsQuery(query string) *_combinedFieldsQuery {

	tmp := &_combinedFieldsQuery{v: types.NewCombinedFieldsQuery()}

	tmp.Query(query)

	return tmp

}

func (s *_combinedFieldsQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_combinedFieldsQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

func (s *_combinedFieldsQuery) Boost(boost float32) *_combinedFieldsQuery {

	s.v.Boost = &boost

	return s
}

func (s *_combinedFieldsQuery) Fields(fields ...string) *_combinedFieldsQuery {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

func (s *_combinedFieldsQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_combinedFieldsQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_combinedFieldsQuery) Operator(operator combinedfieldsoperator.CombinedFieldsOperator) *_combinedFieldsQuery {

	s.v.Operator = &operator
	return s
}

func (s *_combinedFieldsQuery) Query(query string) *_combinedFieldsQuery {

	s.v.Query = query

	return s
}

func (s *_combinedFieldsQuery) QueryName_(queryname_ string) *_combinedFieldsQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_combinedFieldsQuery) ZeroTermsQuery(zerotermsquery combinedfieldszeroterms.CombinedFieldsZeroTerms) *_combinedFieldsQuery {

	s.v.ZeroTermsQuery = &zerotermsquery
	return s
}

func (s *_combinedFieldsQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.CombinedFields = s.v

	return container
}

func (s *_combinedFieldsQuery) CombinedFieldsQueryCaster() *types.CombinedFieldsQuery {
	return s.v
}
