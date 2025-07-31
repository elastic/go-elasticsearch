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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _numberRangeQuery struct {
	k string
	v *types.NumberRangeQuery
}

// Returns documents that contain terms within a provided range.
func NewNumberRangeQuery(key string) *_numberRangeQuery {
	return &_numberRangeQuery{
		k: key,
		v: types.NewNumberRangeQuery(),
	}
}

func (s *_numberRangeQuery) Boost(boost float32) *_numberRangeQuery {

	s.v.Boost = &boost

	return s
}

func (s *_numberRangeQuery) Gt(gt types.Float64) *_numberRangeQuery {

	s.v.Gt = &gt

	return s
}

func (s *_numberRangeQuery) Gte(gte types.Float64) *_numberRangeQuery {

	s.v.Gte = &gte

	return s
}

func (s *_numberRangeQuery) Lt(lt types.Float64) *_numberRangeQuery {

	s.v.Lt = &lt

	return s
}

func (s *_numberRangeQuery) Lte(lte types.Float64) *_numberRangeQuery {

	s.v.Lte = &lte

	return s
}

func (s *_numberRangeQuery) QueryName_(queryname_ string) *_numberRangeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_numberRangeQuery) Relation(relation rangerelation.RangeRelation) *_numberRangeQuery {

	s.v.Relation = &relation
	return s
}

func (s *_numberRangeQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_numberRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_numberRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_numberRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleNumberRangeQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleNumberRangeQuery() *_numberRangeQuery {
	return &_numberRangeQuery{
		k: "",
		v: types.NewNumberRangeQuery(),
	}
}

func (s *_numberRangeQuery) NumberRangeQueryCaster() *types.NumberRangeQuery {
	return s.v.NumberRangeQueryCaster()
}
