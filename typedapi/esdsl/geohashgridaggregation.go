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

type _geoHashGridAggregation struct {
	v *types.GeoHashGridAggregation
}

// A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
// into buckets that represent a grid.
// Each cell is labeled using a geohash which is of user-definable precision.
func NewGeoHashGridAggregation() *_geoHashGridAggregation {

	return &_geoHashGridAggregation{v: types.NewGeoHashGridAggregation()}

}

// The bounding box to filter the points in each bucket.
func (s *_geoHashGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_geoHashGridAggregation {

	s.v.Bounds = *geobounds.GeoBoundsCaster()

	return s
}

// Field containing indexed `geo_point` or `geo_shape` values.
// If the field contains an array, `geohash_grid` aggregates all array values.
func (s *_geoHashGridAggregation) Field(field string) *_geoHashGridAggregation {

	s.v.Field = &field

	return s
}

// The string length of the geohashes used to define cells/buckets in the
// results.
func (s *_geoHashGridAggregation) Precision(geohashprecision types.GeoHashPrecisionVariant) *_geoHashGridAggregation {

	s.v.Precision = *geohashprecision.GeoHashPrecisionCaster()

	return s
}

// Allows for more accurate counting of the top cells returned in the final
// result the aggregation.
// Defaults to returning `max(10,(size x number-of-shards))` buckets from each
// shard.
func (s *_geoHashGridAggregation) ShardSize(shardsize int) *_geoHashGridAggregation {

	s.v.ShardSize = &shardsize

	return s
}

// The maximum number of geohash buckets to return.
func (s *_geoHashGridAggregation) Size(size int) *_geoHashGridAggregation {

	s.v.Size = &size

	return s
}

func (s *_geoHashGridAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.GeohashGrid = s.v

	return container
}

func (s *_geoHashGridAggregation) GeoHashGridAggregationCaster() *types.GeoHashGridAggregation {
	return s.v
}
