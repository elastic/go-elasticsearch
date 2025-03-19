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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

type _movingPercentilesAggregation struct {
	v *types.MovingPercentilesAggregation
}

// Given an ordered series of percentiles, "slides" a window across those
// percentiles and computes cumulative percentiles.
func NewMovingPercentilesAggregation() *_movingPercentilesAggregation {

	return &_movingPercentilesAggregation{v: types.NewMovingPercentilesAggregation()}

}

// Path to the buckets that contain one set of values to correlate.
func (s *_movingPercentilesAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_movingPercentilesAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// `DecimalFormat` pattern for the output value.
// If specified, the formatted value is returned in the aggregation’s
// `value_as_string` property.
func (s *_movingPercentilesAggregation) Format(format string) *_movingPercentilesAggregation {

	s.v.Format = &format

	return s
}

// Policy to apply when gaps are found in the data.
func (s *_movingPercentilesAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_movingPercentilesAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_movingPercentilesAggregation) Keyed(keyed bool) *_movingPercentilesAggregation {

	s.v.Keyed = &keyed

	return s
}

// By default, the window consists of the last n values excluding the current
// bucket.
// Increasing `shift` by 1, moves the starting window position by 1 to the
// right.
func (s *_movingPercentilesAggregation) Shift(shift int) *_movingPercentilesAggregation {

	s.v.Shift = &shift

	return s
}

// The size of window to "slide" across the histogram.
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
