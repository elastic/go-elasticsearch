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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pivot struct {
	v *types.Pivot
}

func NewPivot() *_pivot {

	return &_pivot{v: types.NewPivot()}

}

func (s *_pivot) Aggregations(aggregations map[string]types.Aggregations) *_pivot {

	s.v.Aggregations = aggregations
	return s
}

func (s *_pivot) AddAggregation(key string, value types.AggregationsVariant) *_pivot {

	var tmp map[string]types.Aggregations
	if s.v.Aggregations == nil {
		s.v.Aggregations = make(map[string]types.Aggregations)
	} else {
		tmp = s.v.Aggregations
	}

	tmp[key] = *value.AggregationsCaster()

	s.v.Aggregations = tmp
	return s
}

func (s *_pivot) GroupBy(groupby map[string]types.PivotGroupByContainer) *_pivot {

	s.v.GroupBy = groupby
	return s
}

func (s *_pivot) AddGroupBy(key string, value types.PivotGroupByContainerVariant) *_pivot {

	var tmp map[string]types.PivotGroupByContainer
	if s.v.GroupBy == nil {
		s.v.GroupBy = make(map[string]types.PivotGroupByContainer)
	} else {
		tmp = s.v.GroupBy
	}

	tmp[key] = *value.PivotGroupByContainerCaster()

	s.v.GroupBy = tmp
	return s
}

func (s *_pivot) PivotCaster() *types.Pivot {
	return s.v
}
