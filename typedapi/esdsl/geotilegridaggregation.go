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

type _geoTileGridAggregation struct {
	v *types.GeoTileGridAggregation
}

// A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
// into buckets that represent a grid.
// Each cell corresponds to a map tile as used by many online map sites.
func NewGeoTileGridAggregation() *_geoTileGridAggregation {

	return &_geoTileGridAggregation{v: types.NewGeoTileGridAggregation()}

}

// A bounding box to filter the geo-points or geo-shapes in each bucket.
func (s *_geoTileGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_geoTileGridAggregation {

	s.v.Bounds = *geobounds.GeoBoundsCaster()

	return s
}

// Field containing indexed `geo_point` or `geo_shape` values.
// If the field contains an array, `geotile_grid` aggregates all array values.
func (s *_geoTileGridAggregation) Field(field string) *_geoTileGridAggregation {

	s.v.Field = &field

	return s
}

// Integer zoom of the key used to define cells/buckets in the results.
// Values outside of the range [0,29] will be rejected.
func (s *_geoTileGridAggregation) Precision(geotileprecision int) *_geoTileGridAggregation {

	s.v.Precision = &geotileprecision

	return s
}

// Allows for more accurate counting of the top cells returned in the final
// result the aggregation.
// Defaults to returning `max(10,(size x number-of-shards))` buckets from each
// shard.
func (s *_geoTileGridAggregation) ShardSize(shardsize int) *_geoTileGridAggregation {

	s.v.ShardSize = &shardsize

	return s
}

// The maximum number of buckets to return.
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
