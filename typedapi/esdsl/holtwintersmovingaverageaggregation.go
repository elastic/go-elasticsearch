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

type _holtWintersMovingAverageAggregation struct {
	v *types.HoltWintersMovingAverageAggregation
}

func NewHoltWintersMovingAverageAggregation(settings types.HoltWintersModelSettingsVariant) *_holtWintersMovingAverageAggregation {

	tmp := &_holtWintersMovingAverageAggregation{v: types.NewHoltWintersMovingAverageAggregation()}

	tmp.Settings(settings)

	return tmp

}

// Path to the buckets that contain one set of values to correlate.
func (s *_holtWintersMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_holtWintersMovingAverageAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// `DecimalFormat` pattern for the output value.
// If specified, the formatted value is returned in the aggregation’s
// `value_as_string` property.
func (s *_holtWintersMovingAverageAggregation) Format(format string) *_holtWintersMovingAverageAggregation {

	s.v.Format = &format

	return s
}

// Policy to apply when gaps are found in the data.
func (s *_holtWintersMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_holtWintersMovingAverageAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_holtWintersMovingAverageAggregation) Minimize(minimize bool) *_holtWintersMovingAverageAggregation {

	s.v.Minimize = &minimize

	return s
}

func (s *_holtWintersMovingAverageAggregation) Predict(predict int) *_holtWintersMovingAverageAggregation {

	s.v.Predict = &predict

	return s
}

func (s *_holtWintersMovingAverageAggregation) Settings(settings types.HoltWintersModelSettingsVariant) *_holtWintersMovingAverageAggregation {

	s.v.Settings = *settings.HoltWintersModelSettingsCaster()

	return s
}

func (s *_holtWintersMovingAverageAggregation) Window(window int) *_holtWintersMovingAverageAggregation {

	s.v.Window = &window

	return s
}

func (s *_holtWintersMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MovingAvg = s.v

	return container
}

func (s *_holtWintersMovingAverageAggregation) HoltWintersMovingAverageAggregationCaster() *types.HoltWintersMovingAverageAggregation {
	return s.v
}
