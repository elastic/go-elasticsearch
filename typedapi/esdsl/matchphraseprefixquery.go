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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/zerotermsquery"
)

type _matchPhrasePrefixQuery struct {
	k string
	v *types.MatchPhrasePrefixQuery
}

// Returns documents that contain the words of a provided text, in the same
// order as provided.
// The last term of the provided text is treated as a prefix, matching any words
// that begin with that term.
func NewMatchPhrasePrefixQuery(field string, query string) *_matchPhrasePrefixQuery {
	tmp := &_matchPhrasePrefixQuery{
		k: field,
		v: types.NewMatchPhrasePrefixQuery(),
	}

	tmp.Query(query)
	return tmp
}

// Analyzer used to convert text in the query value into tokens.
func (s *_matchPhrasePrefixQuery) Analyzer(analyzer string) *_matchPhrasePrefixQuery {

	s.v.Analyzer = &analyzer

	return s
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_matchPhrasePrefixQuery) Boost(boost float32) *_matchPhrasePrefixQuery {

	s.v.Boost = &boost

	return s
}

// Maximum number of terms to which the last provided term of the query value
// will expand.
func (s *_matchPhrasePrefixQuery) MaxExpansions(maxexpansions int) *_matchPhrasePrefixQuery {

	s.v.MaxExpansions = &maxexpansions

	return s
}

// Text you wish to find in the provided field.
func (s *_matchPhrasePrefixQuery) Query(query string) *_matchPhrasePrefixQuery {

	s.v.Query = query

	return s
}

func (s *_matchPhrasePrefixQuery) QueryName_(queryname_ string) *_matchPhrasePrefixQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Maximum number of positions allowed between matching tokens.
func (s *_matchPhrasePrefixQuery) Slop(slop int) *_matchPhrasePrefixQuery {

	s.v.Slop = &slop

	return s
}

// Indicates whether no documents are returned if the analyzer removes all
// tokens, such as when using a `stop` filter.
func (s *_matchPhrasePrefixQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_matchPhrasePrefixQuery {

	s.v.ZeroTermsQuery = &zerotermsquery
	return s
}

func (s *_matchPhrasePrefixQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.MatchPhrasePrefix = map[string]types.MatchPhrasePrefixQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleMatchPhrasePrefixQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleMatchPhrasePrefixQuery() *_matchPhrasePrefixQuery {
	return &_matchPhrasePrefixQuery{
		k: "",
		v: types.NewMatchPhrasePrefixQuery(),
	}
}

func (s *_matchPhrasePrefixQuery) MatchPhrasePrefixQueryCaster() *types.MatchPhrasePrefixQuery {
	return s.v.MatchPhrasePrefixQueryCaster()
}
