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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _geohexGridAggregation struct {
	v *types.GeohexGridAggregation
}

// A multi-bucket aggregation that groups `geo_point` and `geo_shape` values
// into buckets that represent a grid. Each cell corresponds to a H3 cell index
// and is labeled using the H3Index representation.
func NewGeohexGridAggregation() *_geohexGridAggregation {

	return &_geohexGridAggregation{v: types.NewGeohexGridAggregation()}

}

func (s *_geohexGridAggregation) Bounds(geobounds types.GeoBoundsVariant) *_geohexGridAggregation {

	s.v.Bounds = *geobounds.GeoBoundsCaster()

	return s
}

func (s *_geohexGridAggregation) Field(field string) *_geohexGridAggregation {

	s.v.Field = field

	return s
}

func (s *_geohexGridAggregation) Precision(precision int) *_geohexGridAggregation {

	s.v.Precision = &precision

	return s
}

func (s *_geohexGridAggregation) ShardSize(shardsize int) *_geohexGridAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_geohexGridAggregation) Size(size int) *_geohexGridAggregation {

	s.v.Size = &size

	return s
}

func (s *_geohexGridAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.GeohexGrid = s.v

	return container
}

func (s *_geohexGridAggregation) GeohexGridAggregationCaster() *types.GeohexGridAggregation {
	return s.v
}
