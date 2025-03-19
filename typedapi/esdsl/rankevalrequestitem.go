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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _rankEvalRequestItem struct {
	v *types.RankEvalRequestItem
}

func NewRankEvalRequestItem() *_rankEvalRequestItem {

	return &_rankEvalRequestItem{v: types.NewRankEvalRequestItem()}

}

// The search request’s ID, used to group result details later.
func (s *_rankEvalRequestItem) Id(id string) *_rankEvalRequestItem {

	s.v.Id = id

	return s
}

// The search template parameters.
func (s *_rankEvalRequestItem) Params(params map[string]json.RawMessage) *_rankEvalRequestItem {

	s.v.Params = params
	return s
}

func (s *_rankEvalRequestItem) AddParam(key string, value json.RawMessage) *_rankEvalRequestItem {

	var tmp map[string]json.RawMessage
	if s.v.Params == nil {
		s.v.Params = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Params
	}

	tmp[key] = value

	s.v.Params = tmp
	return s
}

// List of document ratings
func (s *_rankEvalRequestItem) Ratings(ratings ...types.DocumentRatingVariant) *_rankEvalRequestItem {

	for _, v := range ratings {

		s.v.Ratings = append(s.v.Ratings, *v.DocumentRatingCaster())

	}
	return s
}

// The query being evaluated.
func (s *_rankEvalRequestItem) Request(request types.RankEvalQueryVariant) *_rankEvalRequestItem {

	s.v.Request = request.RankEvalQueryCaster()

	return s
}

// The search template Id
func (s *_rankEvalRequestItem) TemplateId(id string) *_rankEvalRequestItem {

	s.v.TemplateId = &id

	return s
}

func (s *_rankEvalRequestItem) RankEvalRequestItemCaster() *types.RankEvalRequestItem {
	return s.v
}
