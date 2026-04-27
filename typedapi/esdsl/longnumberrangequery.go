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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/rangerelation"
)

type _longNumberRangeQuery struct {
	k string
	v *types.LongNumberRangeQuery
}

// Returns documents that contain terms within a provided range.
func NewLongNumberRangeQuery(key string) *_longNumberRangeQuery {
	return &_longNumberRangeQuery{
		k: key,
		v: types.NewLongNumberRangeQuery(),
	}
}

func (s *_longNumberRangeQuery) Boost(boost float32) *_longNumberRangeQuery {

	s.v.Boost = &boost

	return s
}

func (s *_longNumberRangeQuery) Gt(gt int64) *_longNumberRangeQuery {

	s.v.Gt = &gt

	return s
}

func (s *_longNumberRangeQuery) Gte(gte int64) *_longNumberRangeQuery {

	s.v.Gte = &gte

	return s
}

func (s *_longNumberRangeQuery) Lt(lt int64) *_longNumberRangeQuery {

	s.v.Lt = &lt

	return s
}

func (s *_longNumberRangeQuery) Lte(lte int64) *_longNumberRangeQuery {

	s.v.Lte = &lte

	return s
}

func (s *_longNumberRangeQuery) QueryName_(queryname_ string) *_longNumberRangeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_longNumberRangeQuery) Relation(relation rangerelation.RangeRelation) *_longNumberRangeQuery {

	s.v.Relation = &relation
	return s
}

func (s *_longNumberRangeQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_longNumberRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_longNumberRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_longNumberRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleLongNumberRangeQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleLongNumberRangeQuery() *_longNumberRangeQuery {
	return &_longNumberRangeQuery{
		k: "",
		v: types.NewLongNumberRangeQuery(),
	}
}

func (s *_longNumberRangeQuery) LongNumberRangeQueryCaster() *types.LongNumberRangeQuery {
	return s.v.LongNumberRangeQueryCaster()
}
