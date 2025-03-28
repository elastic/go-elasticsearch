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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/rangerelation"
)

type _untypedRangeQuery struct {
	k string
	v *types.UntypedRangeQuery
}

// Returns documents that contain terms within a provided range.
func NewUntypedRangeQuery(key string) *_untypedRangeQuery {
	return &_untypedRangeQuery{
		k: key,
		v: types.NewUntypedRangeQuery(),
	}
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_untypedRangeQuery) Boost(boost float32) *_untypedRangeQuery {

	s.v.Boost = &boost

	return s
}

// Date format used to convert `date` values in the query.
func (s *_untypedRangeQuery) Format(dateformat string) *_untypedRangeQuery {

	s.v.Format = &dateformat

	return s
}

func (s *_untypedRangeQuery) From(from json.RawMessage) *_untypedRangeQuery {

	s.v.From = &from

	return s
}

// Greater than.
func (s *_untypedRangeQuery) Gt(gt json.RawMessage) *_untypedRangeQuery {

	s.v.Gt = gt

	return s
}

// Greater than or equal to.
func (s *_untypedRangeQuery) Gte(gte json.RawMessage) *_untypedRangeQuery {

	s.v.Gte = gte

	return s
}

// Less than.
func (s *_untypedRangeQuery) Lt(lt json.RawMessage) *_untypedRangeQuery {

	s.v.Lt = lt

	return s
}

// Less than or equal to.
func (s *_untypedRangeQuery) Lte(lte json.RawMessage) *_untypedRangeQuery {

	s.v.Lte = lte

	return s
}

func (s *_untypedRangeQuery) QueryName_(queryname_ string) *_untypedRangeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Indicates how the range query matches values for `range` fields.
func (s *_untypedRangeQuery) Relation(relation rangerelation.RangeRelation) *_untypedRangeQuery {

	s.v.Relation = &relation
	return s
}

// Coordinated Universal Time (UTC) offset or IANA time zone used to convert
// `date` values in the query to UTC.
func (s *_untypedRangeQuery) TimeZone(timezone string) *_untypedRangeQuery {

	s.v.TimeZone = &timezone

	return s
}

func (s *_untypedRangeQuery) To(to json.RawMessage) *_untypedRangeQuery {

	s.v.To = &to

	return s
}

func (s *_untypedRangeQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_untypedRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_untypedRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_untypedRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleUntypedRangeQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleUntypedRangeQuery() *_untypedRangeQuery {
	return &_untypedRangeQuery{
		k: "",
		v: types.NewUntypedRangeQuery(),
	}
}

func (s *_untypedRangeQuery) UntypedRangeQueryCaster() *types.UntypedRangeQuery {
	return s.v.UntypedRangeQueryCaster()
}
