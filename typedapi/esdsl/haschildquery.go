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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/childscoremode"
)

type _hasChildQuery struct {
	v *types.HasChildQuery
}

// Returns parent documents whose joined child documents match a provided query.
func NewHasChildQuery(query types.QueryVariant) *_hasChildQuery {

	tmp := &_hasChildQuery{v: types.NewHasChildQuery()}

	tmp.Query(query)

	return tmp

}

func (s *_hasChildQuery) Boost(boost float32) *_hasChildQuery {

	s.v.Boost = &boost

	return s
}

func (s *_hasChildQuery) IgnoreUnmapped(ignoreunmapped bool) *_hasChildQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_hasChildQuery) InnerHits(innerhits types.InnerHitsVariant) *_hasChildQuery {

	s.v.InnerHits = innerhits.InnerHitsCaster()

	return s
}

func (s *_hasChildQuery) MaxChildren(maxchildren int) *_hasChildQuery {

	s.v.MaxChildren = &maxchildren

	return s
}

func (s *_hasChildQuery) MinChildren(minchildren int) *_hasChildQuery {

	s.v.MinChildren = &minchildren

	return s
}

func (s *_hasChildQuery) Query(query types.QueryVariant) *_hasChildQuery {

	s.v.Query = *query.QueryCaster()

	return s
}

func (s *_hasChildQuery) QueryName_(queryname_ string) *_hasChildQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_hasChildQuery) ScoreMode(scoremode childscoremode.ChildScoreMode) *_hasChildQuery {

	s.v.ScoreMode = &scoremode
	return s
}

func (s *_hasChildQuery) Type(relationname string) *_hasChildQuery {

	s.v.Type = relationname

	return s
}

func (s *_hasChildQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.HasChild = s.v

	return container
}

func (s *_hasChildQuery) HasChildQueryCaster() *types.HasChildQuery {
	return s.v
}
