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

type _termRangeQuery struct {
	k string
	v *types.TermRangeQuery
}

// Returns users that contain terms within a provided range.
func NewTermRangeQuery(key string) *_termRangeQuery {
	return &_termRangeQuery{
		k: key,
		v: types.NewTermRangeQuery(),
	}
}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_termRangeQuery) Boost(boost float32) *_termRangeQuery {

	s.v.Boost = &boost

	return s
}

func (s *_termRangeQuery) From(from string) *_termRangeQuery {

	s.v.From = &from

	return s
}

// Greater than.
func (s *_termRangeQuery) Gt(gt string) *_termRangeQuery {

	s.v.Gt = &gt

	return s
}

// Greater than or equal to.
func (s *_termRangeQuery) Gte(gte string) *_termRangeQuery {

	s.v.Gte = &gte

	return s
}

// Less than.
func (s *_termRangeQuery) Lt(lt string) *_termRangeQuery {

	s.v.Lt = &lt

	return s
}

// Less than or equal to.
func (s *_termRangeQuery) Lte(lte string) *_termRangeQuery {

	s.v.Lte = &lte

	return s
}

func (s *_termRangeQuery) QueryName_(queryname_ string) *_termRangeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Indicates how the range query matches values for `range` fields.
func (s *_termRangeQuery) Relation(relation rangerelation.RangeRelation) *_termRangeQuery {

	s.v.Relation = &relation
	return s
}

func (s *_termRangeQuery) To(to string) *_termRangeQuery {

	s.v.To = &to

	return s
}

func (s *_termRangeQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_termRangeQuery) ApiKeyQueryContainerCaster() *types.ApiKeyQueryContainer {
	container := types.NewApiKeyQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_termRangeQuery) RoleQueryContainerCaster() *types.RoleQueryContainer {
	container := types.NewRoleQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

func (s *_termRangeQuery) UserQueryContainerCaster() *types.UserQueryContainer {
	container := types.NewUserQueryContainer()
	container.Range = map[string]types.RangeQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleTermRangeQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleTermRangeQuery() *_termRangeQuery {
	return &_termRangeQuery{
		k: "",
		v: types.NewTermRangeQuery(),
	}
}

func (s *_termRangeQuery) TermRangeQueryCaster() *types.TermRangeQuery {
	return s.v.TermRangeQueryCaster()
}
