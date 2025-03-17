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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldsoperator"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/combinedfieldszeroterms"
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

// If true, match phrase queries are automatically created for multi-term
// synonyms.
func (s *_combinedFieldsQuery) AutoGenerateSynonymsPhraseQuery(autogeneratesynonymsphrasequery bool) *_combinedFieldsQuery {

	s.v.AutoGenerateSynonymsPhraseQuery = &autogeneratesynonymsphrasequery

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_combinedFieldsQuery) Boost(boost float32) *_combinedFieldsQuery {

	s.v.Boost = &boost

	return s
}

// List of fields to search. Field wildcard patterns are allowed. Only `text`
// fields are supported, and they must all have the same search `analyzer`.
func (s *_combinedFieldsQuery) Fields(fields ...string) *_combinedFieldsQuery {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, v)

	}
	return s
}

// Minimum number of clauses that must match for a document to be returned.
func (s *_combinedFieldsQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_combinedFieldsQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

// Boolean logic used to interpret text in the query value.
func (s *_combinedFieldsQuery) Operator(operator combinedfieldsoperator.CombinedFieldsOperator) *_combinedFieldsQuery {

	s.v.Operator = &operator
	return s
}

// Text to search for in the provided `fields`.
// The `combined_fields` query analyzes the provided text before performing a
// search.
func (s *_combinedFieldsQuery) Query(query string) *_combinedFieldsQuery {

	s.v.Query = query

	return s
}

func (s *_combinedFieldsQuery) QueryName_(queryname_ string) *_combinedFieldsQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Indicates whether no documents are returned if the analyzer removes all
// tokens, such as when using a `stop` filter.
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
