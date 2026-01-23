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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/functionboostmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/functionscoremode"
)

type _functionScoreQuery struct {
	v *types.FunctionScoreQuery
}

// The `function_score` enables you to modify the score of documents that are
// retrieved by a query.
func NewFunctionScoreQuery() *_functionScoreQuery {

	return &_functionScoreQuery{v: types.NewFunctionScoreQuery()}

}

func (s *_functionScoreQuery) BoostMode(boostmode functionboostmode.FunctionBoostMode) *_functionScoreQuery {

	s.v.BoostMode = &boostmode
	return s
}

func (s *_functionScoreQuery) Functions(functions ...types.FunctionScoreVariant) *_functionScoreQuery {

	for _, v := range functions {

		s.v.Functions = append(s.v.Functions, *v.FunctionScoreCaster())

	}
	return s
}

func (s *_functionScoreQuery) MaxBoost(maxboost types.Float64) *_functionScoreQuery {

	s.v.MaxBoost = &maxboost

	return s
}

func (s *_functionScoreQuery) MinScore(minscore types.Float64) *_functionScoreQuery {

	s.v.MinScore = &minscore

	return s
}

func (s *_functionScoreQuery) Query(query types.QueryVariant) *_functionScoreQuery {

	s.v.Query = query.QueryCaster()

	return s
}

func (s *_functionScoreQuery) ScoreMode(scoremode functionscoremode.FunctionScoreMode) *_functionScoreQuery {

	s.v.ScoreMode = &scoremode
	return s
}

func (s *_functionScoreQuery) Boost(boost float32) *_functionScoreQuery {

	s.v.Boost = &boost

	return s
}

func (s *_functionScoreQuery) QueryName_(queryname_ string) *_functionScoreQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_functionScoreQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.FunctionScore = s.v

	return container
}

func (s *_functionScoreQuery) FunctionScoreQueryCaster() *types.FunctionScoreQuery {
	return s.v
}
