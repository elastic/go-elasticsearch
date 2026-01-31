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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _changePointAggregation struct {
	v *types.ChangePointAggregation
}

// A sibling pipeline that detects, spikes, dips, and change points in a metric.
// Given a distribution of values provided by the sibling multi-bucket
// aggregation,
// this aggregation indicates the bucket of any spike or dip and/or the bucket
// at which
// the largest change in the distribution of values, if they are statistically
// significant.
// There must be at least 22 bucketed values. Fewer than 1,000 is preferred.
func NewChangePointAggregation() *_changePointAggregation {

	return &_changePointAggregation{v: types.NewChangePointAggregation()}

}

func (s *_changePointAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_changePointAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_changePointAggregation) Format(format string) *_changePointAggregation {

	s.v.Format = &format

	return s
}

func (s *_changePointAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_changePointAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_changePointAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.ChangePoint = s.v

	return container
}

func (s *_changePointAggregation) ChangePointAggregationCaster() *types.ChangePointAggregation {
	return s.v
}
