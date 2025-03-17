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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _idsQuery struct {
	v *types.IdsQuery
}

// Returns roles based on their IDs.
// This query uses role document IDs stored in the `_id` field.
func NewIdsQuery() *_idsQuery {

	return &_idsQuery{v: types.NewIdsQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_idsQuery) Boost(boost float32) *_idsQuery {

	s.v.Boost = &boost

	return s
}

func (s *_idsQuery) QueryName_(queryname_ string) *_idsQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// An array of document IDs.
func (s *_idsQuery) Values(ids ...string) *_idsQuery {

	s.v.Values = ids

	return s
}

func (s *_idsQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Ids = s.v

	return container
}

func (s *_idsQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()

	container.Ids = s.v

	return container
}

func (s *_idsQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()

	container.Ids = s.v

	return container
}

func (s *_idsQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()

	container.Ids = s.v

	return container
}

func (s *_idsQuery) IdsQueryCaster() *types.IdsQuery {
	return s.v
}
