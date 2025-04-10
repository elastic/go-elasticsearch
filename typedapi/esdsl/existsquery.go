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

type _existsQuery struct {
	v *types.ExistsQuery
}

// Returns roles that contain an indexed value for a field.
func NewExistsQuery() *_existsQuery {

	return &_existsQuery{v: types.NewExistsQuery()}

}

func (s *_existsQuery) Boost(boost float32) *_existsQuery {

	s.v.Boost = &boost

	return s
}

func (s *_existsQuery) Field(field string) *_existsQuery {

	s.v.Field = field

	return s
}

func (s *_existsQuery) QueryName_(queryname_ string) *_existsQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_existsQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Exists = s.v

	return container
}

func (s *_existsQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()

	container.Exists = s.v

	return container
}

func (s *_existsQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()

	container.Exists = s.v

	return container
}

func (s *_existsQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()

	container.Exists = s.v

	return container
}

func (s *_existsQuery) ExistsQueryCaster() *types.ExistsQuery {
	return s.v
}
