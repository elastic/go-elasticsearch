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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geoTileGridAggregation struct {
	v *types.GeoTileGridAggregation
}

// A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
// into buckets that represent a grid.
// Each cell corresponds to a map tile as used by many online map sites.
func NewGeoTileGridAggregation() *_geoTileGridAggregation {

	return &_geoTileGridAggregation{v: types.NewGeoTileGridAggregation()}

}

func (s *_geoTileGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_geoTileGridAggregation {

	s.v.Bounds = *geobounds.GeoBoundsCaster()

	return s
}

func (s *_geoTileGridAggregation) Field(field string) *_geoTileGridAggregation {

	s.v.Field = &field

	return s
}

func (s *_geoTileGridAggregation) Precision(geotileprecision int) *_geoTileGridAggregation {

	s.v.Precision = &geotileprecision

	return s
}

func (s *_geoTileGridAggregation) ShardSize(shardsize int) *_geoTileGridAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_geoTileGridAggregation) Size(size int) *_geoTileGridAggregation {

	s.v.Size = &size

	return s
}

func (s *_geoTileGridAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.GeotileGrid = s.v

	return container
}

func (s *_geoTileGridAggregation) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	container := types.NewPivotGroupByContainer()

	container.GeotileGrid = s.v

	return container
}

func (s *_geoTileGridAggregation) GeoTileGridAggregationCaster() *types.GeoTileGridAggregation {
	return s.v
}
