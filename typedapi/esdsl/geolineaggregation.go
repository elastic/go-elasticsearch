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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortorder"
)

type _geoLineAggregation struct {
	v *types.GeoLineAggregation
}

func NewGeoLineAggregation(point types.GeoLinePointVariant, sort types.GeoLineSortVariant) *_geoLineAggregation {

	tmp := &_geoLineAggregation{v: types.NewGeoLineAggregation()}

	tmp.Point(point)

	tmp.Sort(sort)

	return tmp

}

// When `true`, returns an additional array of the sort values in the feature
// properties.
func (s *_geoLineAggregation) IncludeSort(includesort bool) *_geoLineAggregation {

	s.v.IncludeSort = &includesort

	return s
}

// The name of the geo_point field.
func (s *_geoLineAggregation) Point(point types.GeoLinePointVariant) *_geoLineAggregation {

	s.v.Point = *point.GeoLinePointCaster()

	return s
}

// The maximum length of the line represented in the aggregation.
// Valid sizes are between 1 and 10000.
func (s *_geoLineAggregation) Size(size int) *_geoLineAggregation {

	s.v.Size = &size

	return s
}

// The name of the numeric field to use as the sort key for ordering the points.
// When the `geo_line` aggregation is nested inside a `time_series` aggregation,
// this field defaults to `@timestamp`, and any other value will result in
// error.
func (s *_geoLineAggregation) Sort(sort types.GeoLineSortVariant) *_geoLineAggregation {

	s.v.Sort = *sort.GeoLineSortCaster()

	return s
}

// The order in which the line is sorted (ascending or descending).
func (s *_geoLineAggregation) SortOrder(sortorder sortorder.SortOrder) *_geoLineAggregation {

	s.v.SortOrder = &sortorder
	return s
}

func (s *_geoLineAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Line = s.v

	return container
}

func (s *_geoLineAggregation) GeoLineAggregationCaster() *types.GeoLineAggregation {
	return s.v
}
