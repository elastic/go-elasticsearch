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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _simpleMovingAverageAggregation struct {
	v *types.SimpleMovingAverageAggregation
}

func NewSimpleMovingAverageAggregation(settings types.EmptyObjectVariant) *_simpleMovingAverageAggregation {

	tmp := &_simpleMovingAverageAggregation{v: types.NewSimpleMovingAverageAggregation()}

	tmp.Settings(settings)

	return tmp

}

func (s *_simpleMovingAverageAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_simpleMovingAverageAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_simpleMovingAverageAggregation) Format(format string) *_simpleMovingAverageAggregation {

	s.v.Format = &format

	return s
}

func (s *_simpleMovingAverageAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_simpleMovingAverageAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_simpleMovingAverageAggregation) Minimize(minimize bool) *_simpleMovingAverageAggregation {

	s.v.Minimize = &minimize

	return s
}

func (s *_simpleMovingAverageAggregation) Predict(predict int) *_simpleMovingAverageAggregation {

	s.v.Predict = &predict

	return s
}

func (s *_simpleMovingAverageAggregation) Settings(settings types.EmptyObjectVariant) *_simpleMovingAverageAggregation {

	s.v.Settings = *settings.EmptyObjectCaster()

	return s
}

func (s *_simpleMovingAverageAggregation) Window(window int) *_simpleMovingAverageAggregation {

	s.v.Window = &window

	return s
}

func (s *_simpleMovingAverageAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MovingAvg = s.v

	return container
}

func (s *_simpleMovingAverageAggregation) SimpleMovingAverageAggregationCaster() *types.SimpleMovingAverageAggregation {
	return s.v
}
