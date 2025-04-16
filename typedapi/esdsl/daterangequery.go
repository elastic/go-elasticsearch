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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _dateRangeQuery struct {
	k string
	v *types.DateRangeQuery
}

// Returns documents that contain terms within a provided range.
func NewDateRangeQuery(key string) *_dateRangeQuery {
	return &_dateRangeQuery{
		k: key,
		v: types.NewDateRangeQuery(),
	}
}

func (s *_dateRangeQuery) Boost(boost float32) *_dateRangeQuery {

	s.v.Boost = &boost

	return s
}

func (s *_dateRangeQuery) Format(dateformat string) *_dateRangeQuery {

	s.v.Format = &dateformat

	return s
}

func (s *_dateRangeQuery) From(from string) *_dateRangeQuery {

	s.v.From = &from

	return s
}

func (s *_dateRangeQuery) Gt(datemath string) *_dateRangeQuery {

	s.v.Gt = &datemath

	return s
}

func (s *_dateRangeQuery) Gte(datemath string) *_dateRangeQuery {

	s.v.Gte = &datemath

	return s
}

func (s *_dateRangeQuery) Lt(datemath string) *_dateRangeQuery {

	s.v.Lt = &datemath

	return s
}

func (s *_dateRangeQuery) Lte(datemath string) *_dateRangeQuery {

	s.v.Lte = &datemath

	return s
}

func (s *_dateRangeQuery) QueryName_(queryname_ string) *_dateRangeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_dateRangeQuery) Relation(relation rangerelation.RangeRelation) *_dateRangeQuery {

	s.v.Relation = &relation
	return s
}

func (s *_dateRangeQuery) TimeZone(timezone string) *_dateRangeQuery {

	s.v.TimeZone = &timezone

	return s
}

func (s *_dateRangeQuery) To(to string) *_dateRangeQuery {

	s.v.To = &to

	return s
}

func (s *_dateRangeQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_dateRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_dateRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_dateRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleDateRangeQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleDateRangeQuery() *_dateRangeQuery {
	return &_dateRangeQuery{
		k: "",
		v: types.NewDateRangeQuery(),
	}
}

func (s *_dateRangeQuery) DateRangeQueryCaster() *types.DateRangeQuery {
	return s.v.DateRangeQueryCaster()
}
