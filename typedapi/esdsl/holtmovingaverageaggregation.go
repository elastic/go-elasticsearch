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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _holtMovingAverageAggregation struct {
	v *types.HoltMovingAverageAggregation
}

func NewHoltMovingAverageAggregation(settings types.HoltLinearModelSettingsVariant) *_holtMovingAverageAggregation {

	tmp := &_holtMovingAverageAggregation{v: types.NewHoltMovingAverageAggregation()}

	tmp.Settings(settings)

	return tmp

}

func (s *_holtMovingAverageAggregation) Settings(settings types.HoltLinearModelSettingsVariant) *_holtMovingAverageAggregation {

	s.v.Settings = *settings.HoltLinearModelSettingsCaster()

	return s
}

func (s *_holtMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_holtMovingAverageAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_holtMovingAverageAggregation) Format(format string) *_holtMovingAverageAggregation {

	s.v.Format = &format

	return s
}

func (s *_holtMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_holtMovingAverageAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_holtMovingAverageAggregation) Minimize(minimize bool) *_holtMovingAverageAggregation {

	s.v.Minimize = &minimize

	return s
}

func (s *_holtMovingAverageAggregation) Predict(predict int) *_holtMovingAverageAggregation {

	s.v.Predict = &predict

	return s
}

func (s *_holtMovingAverageAggregation) Window(window int) *_holtMovingAverageAggregation {

	s.v.Window = &window

	return s
}

func (s *_holtMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MovingAvg = s.v

	return container
}

func (s *_holtMovingAverageAggregation) HoltMovingAverageAggregationCaster() *types.HoltMovingAverageAggregation {
	return s.v
}
