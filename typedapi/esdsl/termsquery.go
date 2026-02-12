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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsQuery struct {
	v *types.TermsQuery
}

// Returns documents that contain one or more exact terms in a provided field.
// To return a document, one or more terms must exactly match a field value,
// including whitespace and capitalization.
func NewTermsQuery() *_termsQuery {

	return &_termsQuery{v: types.NewTermsQuery()}

}

func (s *_termsQuery) TermsQuery(termsquery map[string]types.TermsQueryField) *_termsQuery {

	s.v.TermsQuery = termsquery
	return s
}

func (s *_termsQuery) AddTermsQuery(key string, value types.TermsQueryFieldVariant) *_termsQuery {

	var tmp map[string]types.TermsQueryField
	if s.v.TermsQuery == nil {
		s.v.TermsQuery = make(map[string]types.TermsQueryField)
	} else {
		tmp = s.v.TermsQuery
	}

	tmp[key] = *value.TermsQueryFieldCaster()

	s.v.TermsQuery = tmp
	return s
}

func (s *_termsQuery) Boost(boost float32) *_termsQuery {

	s.v.Boost = &boost

	return s
}

func (s *_termsQuery) QueryName_(queryname_ string) *_termsQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_termsQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Terms = s.v

	return container
}

func (s *_termsQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()

	container.Terms = s.v

	return container
}

func (s *_termsQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()

	container.Terms = s.v

	return container
}

func (s *_termsQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()

	container.Terms = s.v

	return container
}

func (s *_termsQuery) TermsQueryCaster() *types.TermsQuery {
	return s.v
}
