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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
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

func (s *_geoLineAggregation) IncludeSort(includesort bool) *_geoLineAggregation {

	s.v.IncludeSort = &includesort

	return s
}

func (s *_geoLineAggregation) Point(point types.GeoLinePointVariant) *_geoLineAggregation {

	s.v.Point = *point.GeoLinePointCaster()

	return s
}

func (s *_geoLineAggregation) Size(size int) *_geoLineAggregation {

	s.v.Size = &size

	return s
}

func (s *_geoLineAggregation) Sort(sort types.GeoLineSortVariant) *_geoLineAggregation {

	s.v.Sort = *sort.GeoLineSortCaster()

	return s
}

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
