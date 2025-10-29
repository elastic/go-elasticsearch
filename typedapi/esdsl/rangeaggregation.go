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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rangeAggregation struct {
	v *types.RangeAggregation
}

// A multi-bucket value source based aggregation that enables the user to define
// a set of ranges - each representing a bucket.
func NewRangeAggregation() *_rangeAggregation {

	return &_rangeAggregation{v: types.NewRangeAggregation()}

}

func (s *_rangeAggregation) Field(field string) *_rangeAggregation {

	s.v.Field = &field

	return s
}

func (s *_rangeAggregation) Format(format string) *_rangeAggregation {

	s.v.Format = &format

	return s
}

func (s *_rangeAggregation) Keyed(keyed bool) *_rangeAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_rangeAggregation) Missing(missing int) *_rangeAggregation {

	s.v.Missing = &missing

	return s
}

func (s *_rangeAggregation) Ranges(ranges ...types.AggregationRangeVariant) *_rangeAggregation {

	for _, v := range ranges {

		s.v.Ranges = append(s.v.Ranges, *v.AggregationRangeCaster())

	}
	return s
}

func (s *_rangeAggregation) Script(script types.ScriptVariant) *_rangeAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_rangeAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Range = s.v

	return container
}

func (s *_rangeAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	container := types.NewApiKeyAggregationContainer()

	container.Range = s.v

	return container
}

func (s *_rangeAggregation) RangeAggregationCaster() *types.RangeAggregation {
	return s.v
}
