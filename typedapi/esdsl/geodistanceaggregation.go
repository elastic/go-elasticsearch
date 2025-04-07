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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/distanceunit"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/geodistancetype"
)

type _geoDistanceAggregation struct {
	v *types.GeoDistanceAggregation
}

// A multi-bucket aggregation that works on `geo_point` fields.
// Evaluates the distance of each document value from an origin point and
// determines the buckets it belongs to, based on ranges defined in the request.
func NewGeoDistanceAggregation() *_geoDistanceAggregation {

	return &_geoDistanceAggregation{v: types.NewGeoDistanceAggregation()}

}

func (s *_geoDistanceAggregation) DistanceType(distancetype geodistancetype.GeoDistanceType) *_geoDistanceAggregation {

	s.v.DistanceType = &distancetype
	return s
}

func (s *_geoDistanceAggregation) Field(field string) *_geoDistanceAggregation {

	s.v.Field = &field

	return s
}

func (s *_geoDistanceAggregation) Origin(geolocation types.GeoLocationVariant) *_geoDistanceAggregation {

	s.v.Origin = *geolocation.GeoLocationCaster()

	return s
}

func (s *_geoDistanceAggregation) Ranges(ranges ...types.AggregationRangeVariant) *_geoDistanceAggregation {

	for _, v := range ranges {

		s.v.Ranges = append(s.v.Ranges, *v.AggregationRangeCaster())

	}
	return s
}

func (s *_geoDistanceAggregation) Unit(unit distanceunit.DistanceUnit) *_geoDistanceAggregation {

	s.v.Unit = &unit
	return s
}

func (s *_geoDistanceAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.GeoDistance = s.v

	return container
}

func (s *_geoDistanceAggregation) GeoDistanceAggregationCaster() *types.GeoDistanceAggregation {
	return s.v
}
