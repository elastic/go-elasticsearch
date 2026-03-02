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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _shapeQuery struct {
	v *types.ShapeQuery
}

// Queries documents that contain fields indexed using the `shape` type.
func NewShapeQuery() *_shapeQuery {

	return &_shapeQuery{v: types.NewShapeQuery()}

}

func (s *_shapeQuery) IgnoreUnmapped(ignoreunmapped bool) *_shapeQuery {

	s.v.IgnoreUnmapped = &ignoreunmapped

	return s
}

func (s *_shapeQuery) ShapeQuery(shapequery map[string]types.ShapeFieldQuery) *_shapeQuery {

	s.v.ShapeQuery = shapequery
	return s
}

func (s *_shapeQuery) AddShapeQuery(key string, value types.ShapeFieldQueryVariant) *_shapeQuery {

	var tmp map[string]types.ShapeFieldQuery
	if s.v.ShapeQuery == nil {
		s.v.ShapeQuery = make(map[string]types.ShapeFieldQuery)
	} else {
		tmp = s.v.ShapeQuery
	}

	tmp[key] = *value.ShapeFieldQueryCaster()

	s.v.ShapeQuery = tmp
	return s
}

func (s *_shapeQuery) Boost(boost float32) *_shapeQuery {

	s.v.Boost = &boost

	return s
}

func (s *_shapeQuery) QueryName_(queryname_ string) *_shapeQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_shapeQuery) QueryCaster() *types.Query {
	container := types.NewQuery()

	container.Shape = s.v

	return container
}

func (s *_shapeQuery) ShapeQueryCaster() *types.ShapeQuery {
	return s.v
}
