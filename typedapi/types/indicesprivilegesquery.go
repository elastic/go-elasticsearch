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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// IndicesPrivilegesQuery holds the union for the following types:
//
//	QueryContainer
//	RoleTemplateQuery
//	string
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/security/_types/Privileges.ts#L106-L114
type IndicesPrivilegesQuery interface{}

// IndicesPrivilegesQueryBuilder holds IndicesPrivilegesQuery struct and provides a builder API.
type IndicesPrivilegesQueryBuilder struct {
	v IndicesPrivilegesQuery
}

// NewIndicesPrivilegesQuery provides a builder for the IndicesPrivilegesQuery struct.
func NewIndicesPrivilegesQueryBuilder() *IndicesPrivilegesQueryBuilder {
	return &IndicesPrivilegesQueryBuilder{}
}

// Build finalize the chain and returns the IndicesPrivilegesQuery struct
func (u *IndicesPrivilegesQueryBuilder) Build() IndicesPrivilegesQuery {
	return u.v
}

func (u *IndicesPrivilegesQueryBuilder) QueryContainer(querycontainer *QueryContainerBuilder) *IndicesPrivilegesQueryBuilder {
	v := querycontainer.Build()
	u.v = &v
	return u
}

func (u *IndicesPrivilegesQueryBuilder) RoleTemplateQuery(roletemplatequery *RoleTemplateQueryBuilder) *IndicesPrivilegesQueryBuilder {
	v := roletemplatequery.Build()
	u.v = &v
	return u
}

func (u *IndicesPrivilegesQueryBuilder) String(string string) *IndicesPrivilegesQueryBuilder {
	u.v = &string
	return u
}
