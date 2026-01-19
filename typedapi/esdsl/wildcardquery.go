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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _wildcardQuery struct {
	k string
	v *types.WildcardQuery
}

// Returns documents that contain terms matching a wildcard pattern.
func NewWildcardQuery(field string, value string) *_wildcardQuery {
	tmp := &_wildcardQuery{
		k: field,
		v: types.NewWildcardQuery(),
	}

	tmp.Value(value)
	return tmp
}

func (s *_wildcardQuery) CaseInsensitive(caseinsensitive bool) *_wildcardQuery {

	s.v.CaseInsensitive = &caseinsensitive

	return s
}

func (s *_wildcardQuery) Rewrite(multitermqueryrewrite string) *_wildcardQuery {

	s.v.Rewrite = &multitermqueryrewrite

	return s
}

func (s *_wildcardQuery) Value(value string) *_wildcardQuery {

	s.v.Value = &value

	return s
}

func (s *_wildcardQuery) Wildcard(wildcard string) *_wildcardQuery {

	s.v.Wildcard = &wildcard

	return s
}

func (s *_wildcardQuery) Boost(boost float32) *_wildcardQuery {

	s.v.Boost = &boost

	return s
}

func (s *_wildcardQuery) QueryName_(queryname_ string) *_wildcardQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_wildcardQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Wildcard = map[string]types.WildcardQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_wildcardQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Wildcard = map[string]types.WildcardQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_wildcardQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Wildcard = map[string]types.WildcardQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_wildcardQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Wildcard = map[string]types.WildcardQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleWildcardQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleWildcardQuery() *_wildcardQuery {
	return &_wildcardQuery{
		k: "",
		v: types.NewWildcardQuery(),
	}
}

func (s *_wildcardQuery) WildcardQueryCaster() *types.WildcardQuery {
	return s.v.WildcardQueryCaster()
}
