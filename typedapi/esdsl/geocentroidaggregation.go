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

type _geoCentroidAggregation struct {
	v *types.GeoCentroidAggregation
}

// A metric aggregation that computes the weighted centroid from all coordinate
// values for geo fields.
func NewGeoCentroidAggregation() *_geoCentroidAggregation {

	return &_geoCentroidAggregation{v: types.NewGeoCentroidAggregation()}

}

func (s *_geoCentroidAggregation) Count(count int64) *_geoCentroidAggregation {

	s.v.Count = &count

	return s
}

func (s *_geoCentroidAggregation) Field(field string) *_geoCentroidAggregation {

	s.v.Field = &field

	return s
}

func (s *_geoCentroidAggregation) Location(geolocation types.GeoLocationVariant) *_geoCentroidAggregation {

	s.v.Location = *geolocation.GeoLocationCaster()

	return s
}

func (s *_geoCentroidAggregation) Missing(missing types.MissingVariant) *_geoCentroidAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_geoCentroidAggregation) Script(script types.ScriptVariant) *_geoCentroidAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_geoCentroidAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.GeoCentroid = s.v

	return container
}

func (s *_geoCentroidAggregation) GeoCentroidAggregationCaster() *types.GeoCentroidAggregation {
	return s.v
}
