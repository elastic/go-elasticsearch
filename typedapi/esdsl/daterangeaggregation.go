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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _dateRangeAggregation struct {
	v *types.DateRangeAggregation
}

// A multi-bucket value source based aggregation that enables the user to define
// a set of date ranges - each representing a bucket.
func NewDateRangeAggregation() *_dateRangeAggregation {

	return &_dateRangeAggregation{v: types.NewDateRangeAggregation()}

}

// The date field whose values are use to build ranges.
func (s *_dateRangeAggregation) Field(field string) *_dateRangeAggregation {

	s.v.Field = &field

	return s
}

// The date format used to format `from` and `to` in the response.
func (s *_dateRangeAggregation) Format(format string) *_dateRangeAggregation {

	s.v.Format = &format

	return s
}

// Set to `true` to associate a unique string key with each bucket and returns
// the ranges as a hash rather than an array.
func (s *_dateRangeAggregation) Keyed(keyed bool) *_dateRangeAggregation {

	s.v.Keyed = &keyed

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_dateRangeAggregation) Missing(missing types.MissingVariant) *_dateRangeAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

// Array of date ranges.
func (s *_dateRangeAggregation) Ranges(ranges ...types.DateRangeExpressionVariant) *_dateRangeAggregation {

	for _, v := range ranges {

		s.v.Ranges = append(s.v.Ranges, *v.DateRangeExpressionCaster())

	}
	return s
}

// Time zone used to convert dates from another time zone to UTC.
func (s *_dateRangeAggregation) TimeZone(timezone string) *_dateRangeAggregation {

	s.v.TimeZone = &timezone

	return s
}

func (s *_dateRangeAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.DateRange = s.v

	return container
}

func (s *_dateRangeAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	container := types.NewApiKeyAggregationContainer()

	container.DateRange = s.v

	return container
}

func (s *_dateRangeAggregation) DateRangeAggregationCaster() *types.DateRangeAggregation {
	return s.v
}
