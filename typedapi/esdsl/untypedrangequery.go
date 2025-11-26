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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _untypedRangeQuery struct {
	k string
	v *types.UntypedRangeQuery
}

// Returns roles that contain terms within a provided range.
func NewUntypedRangeQuery(key string) *_untypedRangeQuery {
	return &_untypedRangeQuery{
		k: key,
		v: types.NewUntypedRangeQuery(),
	}
}

func (s *_untypedRangeQuery) Boost(boost float32) *_untypedRangeQuery {

	s.v.Boost = &boost

	return s
}

func (s *_untypedRangeQuery) Format(dateformat string) *_untypedRangeQuery {

	s.v.Format = &dateformat

	return s
}

func (s *_untypedRangeQuery) Gt(gt json.RawMessage) *_untypedRangeQuery {

	s.v.Gt = gt

	return s
}

func (s *_untypedRangeQuery) Gte(gte json.RawMessage) *_untypedRangeQuery {

	s.v.Gte = gte

	return s
}

func (s *_untypedRangeQuery) Lt(lt json.RawMessage) *_untypedRangeQuery {

	s.v.Lt = lt

	return s
}

func (s *_untypedRangeQuery) Lte(lte json.RawMessage) *_untypedRangeQuery {

	s.v.Lte = lte

	return s
}

func (s *_untypedRangeQuery) QueryName_(queryname_ string) *_untypedRangeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_untypedRangeQuery) Relation(relation rangerelation.RangeRelation) *_untypedRangeQuery {

	s.v.Relation = &relation
	return s
}

func (s *_untypedRangeQuery) TimeZone(timezone string) *_untypedRangeQuery {

	s.v.TimeZone = &timezone

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
