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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _constantScoreQuery struct {
	v *types.ConstantScoreQuery
}

// Wraps a filter query and returns every matching document with a relevance
// score equal to the `boost` parameter value.
func NewConstantScoreQuery(filter types.QueryVariant) *_constantScoreQuery {

	tmp := &_constantScoreQuery{v: types.NewConstantScoreQuery()}

	tmp.Filter(filter)

	return tmp

}

func (s *_constantScoreQuery) Boost(boost float32) *_constantScoreQuery {

	s.v.Boost = &boost

	return s
}

func (s *_constantScoreQuery) Filter(filter types.QueryVariant) *_constantScoreQuery {

	s.v.Filter = *filter.QueryCaster()

	return s
}

func (s *_constantScoreQuery) QueryName_(queryname_ string) *_constantScoreQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_constantScoreQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.ConstantScore = s.v

	return container
}

func (s *_constantScoreQuery) ConstantScoreQueryCaster() *types.ConstantScoreQuery {
	return s.v
}
