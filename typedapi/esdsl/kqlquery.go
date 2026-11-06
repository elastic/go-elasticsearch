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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _kqlQuery struct {
	v *types.KqlQuery
}

// Returns documents matching a provided Kibana Query Language (KQL) expression.
// The expression is parsed and rewritten into standard Query DSL.
func NewKqlQuery(query string) *_kqlQuery {

	tmp := &_kqlQuery{v: types.NewKqlQuery()}

	tmp.Query(query)

	return tmp

}

func (s *_kqlQuery) CaseInsensitive(caseinsensitive bool) *_kqlQuery {

	s.v.CaseInsensitive = &caseinsensitive

	return s
}

func (s *_kqlQuery) DefaultField(field string) *_kqlQuery {

	s.v.DefaultField = &field

	return s
}

func (s *_kqlQuery) Query(query string) *_kqlQuery {

	s.v.Query = query

	return s
}

func (s *_kqlQuery) TimeZone(timezone string) *_kqlQuery {

	s.v.TimeZone = &timezone

	return s
}

func (s *_kqlQuery) Boost(boost float32) *_kqlQuery {

	s.v.Boost = &boost

	return s
}

func (s *_kqlQuery) QueryName_(queryname_ string) *_kqlQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_kqlQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Kql = s.v

	return container
}

func (s *_kqlQuery) KqlQueryCaster() *types.KqlQuery {
	return s.v
}
