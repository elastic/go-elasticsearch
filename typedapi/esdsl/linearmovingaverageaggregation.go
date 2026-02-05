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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _linearMovingAverageAggregation struct {
	v *types.LinearMovingAverageAggregation
}

func NewLinearMovingAverageAggregation(settings types.EmptyObjectVariant) *_linearMovingAverageAggregation {

	tmp := &_linearMovingAverageAggregation{v: types.NewLinearMovingAverageAggregation()}

	tmp.Settings(settings)

	return tmp

}

func (s *_linearMovingAverageAggregation) Settings(settings types.EmptyObjectVariant) *_linearMovingAverageAggregation {

	s.v.Settings = *settings.EmptyObjectCaster()

	return s
}

func (s *_linearMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_linearMovingAverageAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_linearMovingAverageAggregation) Format(format string) *_linearMovingAverageAggregation {

	s.v.Format = &format

	return s
}

func (s *_linearMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_linearMovingAverageAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_linearMovingAverageAggregation) Minimize(minimize bool) *_linearMovingAverageAggregation {

	s.v.Minimize = &minimize

	return s
}

func (s *_linearMovingAverageAggregation) Predict(predict int) *_linearMovingAverageAggregation {

	s.v.Predict = &predict

	return s
}

func (s *_linearMovingAverageAggregation) Window(window int) *_linearMovingAverageAggregation {

	s.v.Window = &window

	return s
}

func (s *_linearMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MovingAvg = s.v

	return container
}

func (s *_linearMovingAverageAggregation) LinearMovingAverageAggregationCaster() *types.LinearMovingAverageAggregation {
	return s.v
}
