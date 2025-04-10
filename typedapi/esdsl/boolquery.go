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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _boolQuery struct {
	v *types.BoolQuery
}

// matches documents matching boolean combinations of other queries.
func NewBoolQuery() *_boolQuery {

	return &_boolQuery{v: types.NewBoolQuery()}

}

func (s *_boolQuery) Boost(boost float32) *_boolQuery {

	s.v.Boost = &boost

	return s
}

func (s *_boolQuery) Filter(filters ...types.QueryVariant) *_boolQuery {

	s.v.Filter = make([]types.Query, len(filters))
	for i, v := range filters {
		s.v.Filter[i] = *v.QueryCaster()
	}

	return s
}

func (s *_boolQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_boolQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_boolQuery) Must(musts ...types.QueryVariant) *_boolQuery {

	s.v.Must = make([]types.Query, len(musts))
	for i, v := range musts {
		s.v.Must[i] = *v.QueryCaster()
	}

	return s
}

func (s *_boolQuery) MustNot(mustnots ...types.QueryVariant) *_boolQuery {

	s.v.MustNot = make([]types.Query, len(mustnots))
	for i, v := range mustnots {
		s.v.MustNot[i] = *v.QueryCaster()
	}

	return s
}

func (s *_boolQuery) QueryName_(queryname_ string) *_boolQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_boolQuery) Should(shoulds ...types.QueryVariant) *_boolQuery {

	s.v.Should = make([]types.Query, len(shoulds))
	for i, v := range shoulds {
		s.v.Should[i] = *v.QueryCaster()
	}

	return s
}

func (s *_boolQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Bool = s.v

	return container
}

func (s *_boolQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()

	container.Bool = s.v

	return container
}

func (s *_boolQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()

	container.Bool = s.v

	return container
}

func (s *_boolQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()

	container.Bool = s.v

	return container
}

func (s *_boolQuery) BoolQueryCaster() *types.BoolQuery {
	return s.v
}
