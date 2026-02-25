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

type _termQuery struct {
	k string
	v *types.TermQuery
}

// Returns documents that contain an exact term in a provided field.
// To return a document, the query term must exactly match the queried field's
// value, including whitespace and capitalization.
func NewTermQuery(field string, value types.FieldValueVariant) *_termQuery {
	tmp := &_termQuery{
		k: field,
		v: types.NewTermQuery(),
	}

	tmp.Value(value)
	return tmp
}

func (s *_termQuery) CaseInsensitive(caseinsensitive bool) *_termQuery {

	s.v.CaseInsensitive = &caseinsensitive

	return s
}

func (s *_termQuery) Value(fieldvalue types.FieldValueVariant) *_termQuery {

	s.v.Value = *fieldvalue.FieldValueCaster()

	return s
}

func (s *_termQuery) Boost(boost float32) *_termQuery {

	s.v.Boost = &boost

	return s
}

func (s *_termQuery) QueryName_(queryname_ string) *_termQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_termQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Term = map[string]types.TermQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_termQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Term = map[string]types.TermQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_termQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Term = map[string]types.TermQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_termQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Term = map[string]types.TermQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleTermQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleTermQuery() *_termQuery {
	return &_termQuery{
		k: "",
		v: types.NewTermQuery(),
	}
}

func (s *_termQuery) TermQueryCaster() *types.TermQuery {
	return s.v.TermQueryCaster()
}
