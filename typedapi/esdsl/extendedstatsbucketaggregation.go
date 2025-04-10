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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _extendedStatsBucketAggregation struct {
	v *types.ExtendedStatsBucketAggregation
}

// A sibling pipeline aggregation which calculates a variety of stats across all
// bucket of a specified metric in a sibling aggregation.
func NewExtendedStatsBucketAggregation() *_extendedStatsBucketAggregation {

	return &_extendedStatsBucketAggregation{v: types.NewExtendedStatsBucketAggregation()}

}

func (s *_extendedStatsBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_extendedStatsBucketAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_extendedStatsBucketAggregation) Format(format string) *_extendedStatsBucketAggregation {

	s.v.Format = &format

	return s
}

func (s *_extendedStatsBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_extendedStatsBucketAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_extendedStatsBucketAggregation) Sigma(sigma types.Float64) *_extendedStatsBucketAggregation {

	s.v.Sigma = &sigma

	return s
}

func (s *_extendedStatsBucketAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.ExtendedStatsBucket = s.v

	return container
}

func (s *_extendedStatsBucketAggregation) ExtendedStatsBucketAggregationCaster() *types.ExtendedStatsBucketAggregation {
	return s.v
}
