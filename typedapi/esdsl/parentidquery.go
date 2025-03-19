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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _parentIdQuery struct {
	v *types.ParentIdQuery
}

// Returns child documents joined to a specific parent document.
func NewParentIdQuery() *_parentIdQuery {

	return &_parentIdQuery{v: types.NewParentIdQuery()}

}

// Floating point number used to decrease or increase the relevance scores of
// the query.
// Boost values are relative to the default value of 1.0.
// A boost value between 0 and 1.0 decreases the relevance score.
// A value greater than 1.0 increases the relevance score.
func (s *_parentIdQuery) Boost(boost float32) *_parentIdQuery {

	s.v.Boost = &boost

	return s
}

// ID of the parent document.
func (s *_parentIdQuery) Id(id string) *_parentIdQuery {

	s.v.Id = &id

	return s
}

// Indicates whether to ignore an unmapped `type` and not return any documents
// instead of an error.
func (s *_parentIdQuery) IgnoreUnmapped(ignoreunmapped bool) *_parentIdQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_parentIdQuery) QueryName_(queryname_ string) *_parentIdQuery {

	s.v.QueryName_ = &queryname_

	return s
}

// Name of the child relationship mapped for the `join` field.
func (s *_parentIdQuery) Type(relationname string) *_parentIdQuery {

	s.v.Type = &relationname

	return s
}

func (s *_parentIdQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.ParentId = s.v

	return container
}

func (s *_parentIdQuery) ParentIdQueryCaster() *types.ParentIdQuery {
	return s.v
}
