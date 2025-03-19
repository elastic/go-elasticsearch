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

// This is provide all the types that are part of the union.
type _indicesPrivilegesQuery struct {
	v types.IndicesPrivilegesQuery
}

func NewIndicesPrivilegesQuery() *_indicesPrivilegesQuery {
	return &_indicesPrivilegesQuery{v: nil}
}

func (u *_indicesPrivilegesQuery) String(string string) *_indicesPrivilegesQuery {

	u.v = &string

	return u
}

func (u *_indicesPrivilegesQuery) Query(query types.QueryVariant) *_indicesPrivilegesQuery {

	u.v = &query

	return u
}

// Interface implementation for Query in IndicesPrivilegesQuery union
func (u *_query) IndicesPrivilegesQueryCaster() *types.IndicesPrivilegesQuery {
	t := types.IndicesPrivilegesQuery(u.v)
	return &t
}

func (u *_indicesPrivilegesQuery) RoleTemplateQuery(roletemplatequery types.RoleTemplateQueryVariant) *_indicesPrivilegesQuery {

	u.v = &roletemplatequery

	return u
}

// Interface implementation for RoleTemplateQuery in IndicesPrivilegesQuery union
func (u *_roleTemplateQuery) IndicesPrivilegesQueryCaster() *types.IndicesPrivilegesQuery {
	t := types.IndicesPrivilegesQuery(u.v)
	return &t
}

func (u *_indicesPrivilegesQuery) IndicesPrivilegesQueryCaster() *types.IndicesPrivilegesQuery {
	return &u.v
}
