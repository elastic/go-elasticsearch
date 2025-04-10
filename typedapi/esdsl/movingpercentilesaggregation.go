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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _movingPercentilesAggregation struct {
	v *types.MovingPercentilesAggregation
}

// Given an ordered series of percentiles, "slides" a window across those
// percentiles and computes cumulative percentiles.
func NewMovingPercentilesAggregation() *_movingPercentilesAggregation {

	return &_movingPercentilesAggregation{v: types.NewMovingPercentilesAggregation()}

}

func (s *_movingPercentilesAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_movingPercentilesAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_movingPercentilesAggregation) Format(format string) *_movingPercentilesAggregation {

	s.v.Format = &format

	return s
}

func (s *_movingPercentilesAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_movingPercentilesAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_movingPercentilesAggregation) Keyed(keyed bool) *_movingPercentilesAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_movingPercentilesAggregation) Shift(shift int) *_movingPercentilesAggregation {

	s.v.Shift = &shift

	return s
}

func (s *_movingPercentilesAggregation) Window(window int) *_movingPercentilesAggregation {

	s.v.Window = &window

	return s
}

func (s *_movingPercentilesAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MovingPercentiles = s.v

	return container
}

func (s *_movingPercentilesAggregation) MovingPercentilesAggregationCaster() *types.MovingPercentilesAggregation {
	return s.v
}
