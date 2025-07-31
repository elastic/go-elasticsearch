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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
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

func (s *_movingFunctionAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_movingFunctionAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_movingFunctionAggregation) Format(format string) *_movingFunctionAggregation {

	s.v.Format = &format

	return s
}

func (s *_movingFunctionAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_movingFunctionAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_movingFunctionAggregation) Script(script string) *_movingFunctionAggregation {

	s.v.Script = &script

	return s
}

func (s *_movingFunctionAggregation) Shift(shift int) *_movingFunctionAggregation {

	s.v.Shift = &shift

	return s
}

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
