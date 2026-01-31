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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
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

func (s *_matchPhrasePrefixQuery) Analyzer(analyzer string) *_matchPhrasePrefixQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_matchPhrasePrefixQuery) MaxExpansions(maxexpansions int) *_matchPhrasePrefixQuery {

	s.v.MaxExpansions = &maxexpansions

	return s
}

func (s *_matchPhrasePrefixQuery) Query(query string) *_matchPhrasePrefixQuery {

	s.v.Query = query

	return s
}

func (s *_matchPhrasePrefixQuery) Slop(slop int) *_matchPhrasePrefixQuery {

	s.v.Slop = &slop

	return s
}

func (s *_matchPhrasePrefixQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_matchPhrasePrefixQuery {

	s.v.ZeroTermsQuery = &zerotermsquery
	return s
}

func (s *_matchPhrasePrefixQuery) Boost(boost float32) *_matchPhrasePrefixQuery {

	s.v.Boost = &boost

	return s
}

func (s *_matchPhrasePrefixQuery) QueryName_(queryname_ string) *_matchPhrasePrefixQuery {

	s.v.QueryName_ = &queryname_

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
