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
// https://github.com/elastic/elasticsearch-specification/tree/cf6914e80d9c586e872b7d5e9e74ca34905dcf5f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _maxBucketAggregation struct {
	v *types.MaxBucketAggregation
}

// A sibling pipeline aggregation which identifies the bucket(s) with the
// maximum value of a specified metric in a sibling aggregation and outputs both
// the value and the key(s) of the bucket(s).
func NewMaxBucketAggregation() *_maxBucketAggregation {

	return &_maxBucketAggregation{v: types.NewMaxBucketAggregation()}

}

func (s *_maxBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_maxBucketAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_maxBucketAggregation) Format(format string) *_maxBucketAggregation {

	s.v.Format = &format

	return s
}

func (s *_maxBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_maxBucketAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_maxBucketAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MaxBucket = s.v

	return container
}

func (s *_maxBucketAggregation) MaxBucketAggregationCaster() *types.MaxBucketAggregation {
	return s.v
}
