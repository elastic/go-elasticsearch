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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _parentAggregation struct {
	v *types.ParentAggregation
}

// A special single bucket aggregation that selects parent documents that have
// the specified type, as defined in a `join` field.
func NewParentAggregation() *_parentAggregation {

	return &_parentAggregation{v: types.NewParentAggregation()}

}

func (s *_parentAggregation) Type(relationname string) *_parentAggregation {

	s.v.Type = &relationname

	return s
}

func (s *_parentAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Parent = s.v

	return container
}

func (s *_parentAggregation) ParentAggregationCaster() *types.ParentAggregation {
	return s.v
}
