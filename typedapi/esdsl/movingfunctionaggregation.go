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

type _movingFunctionAggregation struct {
	v *types.MovingFunctionAggregation
}

// Given an ordered series of data, "slides" a window across the data and runs a
// custom script on each window of data.
// For convenience, a number of common functions are predefined such as `min`,
// `max`, and moving averages.
func NewMovingFunctionAggregation() *_movingFunctionAggregation {

	return &_movingFunctionAggregation{v: types.NewMovingFunctionAggregation()}

}

// Path to the buckets that contain one set of values to correlate.
func (s *_movingFunctionAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_movingFunctionAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// `DecimalFormat` pattern for the output value.
// If specified, the formatted value is returned in the aggregation’s
// `value_as_string` property.
func (s *_movingFunctionAggregation) Format(format string) *_movingFunctionAggregation {

	s.v.Format = &format

	return s
}

// Policy to apply when gaps are found in the data.
func (s *_movingFunctionAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_movingFunctionAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

// The script that should be executed on each window of data.
func (s *_movingFunctionAggregation) Script(script string) *_movingFunctionAggregation {

	s.v.Script = &script

	return s
}

// By default, the window consists of the last n values excluding the current
// bucket.
// Increasing `shift` by 1, moves the starting window position by 1 to the
// right.
func (s *_movingFunctionAggregation) Shift(shift int) *_movingFunctionAggregation {

	s.v.Shift = &shift

	return s
}

// The size of window to "slide" across the histogram.
func (s *_movingFunctionAggregation) Window(window int) *_movingFunctionAggregation {

	s.v.Window = &window

	return s
}

func (s *_movingFunctionAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MovingFn = s.v

	return container
}

func (s *_movingFunctionAggregation) MovingFunctionAggregationCaster() *types.MovingFunctionAggregation {
	return s.v
}
