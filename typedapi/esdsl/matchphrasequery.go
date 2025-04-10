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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/zerotermsquery"
)

type _matchPhraseQuery struct {
	k string
	v *types.MatchPhraseQuery
}

// Analyzes the text and creates a phrase query out of the analyzed text.
func NewMatchPhraseQuery(field string, query string) *_matchPhraseQuery {
	tmp := &_matchPhraseQuery{
		k: field,
		v: types.NewMatchPhraseQuery(),
	}

	tmp.Query(query)
	return tmp
}

func (s *_matchPhraseQuery) Analyzer(analyzer string) *_matchPhraseQuery {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_matchPhraseQuery) Boost(boost float32) *_matchPhraseQuery {

	s.v.Boost = &boost

	return s
}

func (s *_matchPhraseQuery) Query(query string) *_matchPhraseQuery {

	s.v.Query = query

	return s
}

func (s *_matchPhraseQuery) QueryName_(queryname_ string) *_matchPhraseQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_matchPhraseQuery) Slop(slop int) *_matchPhraseQuery {

	s.v.Slop = &slop

	return s
}

func (s *_matchPhraseQuery) ZeroTermsQuery(zerotermsquery zerotermsquery.ZeroTermsQuery) *_matchPhraseQuery {

	s.v.ZeroTermsQuery = &zerotermsquery
	return s
}

func (s *_matchPhraseQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.MatchPhrase = map[string]types.MatchPhraseQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleMatchPhraseQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleMatchPhraseQuery() *_matchPhraseQuery {
	return &_matchPhraseQuery{
		k: "",
		v: types.NewMatchPhraseQuery(),
	}
}

func (s *_matchPhraseQuery) MatchPhraseQueryCaster() *types.MatchPhraseQuery {
	return s.v.MatchPhraseQueryCaster()
}
