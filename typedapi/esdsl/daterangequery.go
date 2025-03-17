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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

type _dateRangeQuery struct {
	k string
	v *types.DateRangeQuery
}

// Returns users that contain terms within a provided range.
func NewDateRangeQuery(key string) *_dateRangeQuery {
	return &_dateRangeQuery{
		k: key,
		v: types.NewDateRangeQuery(),
	}
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_dateRangeQuery) Boost(boost float32) *_dateRangeQuery {

	s.v.Boost = &boost

	return s
}

// Date format used to convert `date` values in the query.
func (s *_dateRangeQuery) Format(dateformat string) *_dateRangeQuery {

	s.v.Format = &dateformat

	return s
}

func (s *_dateRangeQuery) From(from string) *_dateRangeQuery {

	s.v.From = &from

	return s
}

// Greater than.
func (s *_dateRangeQuery) Gt(datemath string) *_dateRangeQuery {

	s.v.Gt = &datemath

	return s
}

// Greater than or equal to.
func (s *_dateRangeQuery) Gte(datemath string) *_dateRangeQuery {

	s.v.Gte = &datemath

	return s
}

// Less than.
func (s *_dateRangeQuery) Lt(datemath string) *_dateRangeQuery {

	s.v.Lt = &datemath

	return s
}

// Less than or equal to.
func (s *_dateRangeQuery) Lte(datemath string) *_dateRangeQuery {

	s.v.Lte = &datemath

	return s
}

func (s *_dateRangeQuery) QueryName_(queryname_ string) *_dateRangeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Indicates how the range query matches values for `range` fields.
func (s *_dateRangeQuery) Relation(relation rangerelation.RangeRelation) *_dateRangeQuery {

	s.v.Relation = &relation
	return s
}

// Coordinated Universal Time (UTC) offset or IANA time zone used to convert
// `date` values in the query to UTC.
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
