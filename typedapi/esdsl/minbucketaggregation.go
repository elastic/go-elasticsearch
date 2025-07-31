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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _minBucketAggregation struct {
	v *types.MinBucketAggregation
}

// A sibling pipeline aggregation which identifies the bucket(s) with the
// minimum value of a specified metric in a sibling aggregation and outputs both
// the value and the key(s) of the bucket(s).
func NewMinBucketAggregation() *_minBucketAggregation {

	return &_minBucketAggregation{v: types.NewMinBucketAggregation()}

}

func (s *_minBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_minBucketAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_minBucketAggregation) Format(format string) *_minBucketAggregation {

	s.v.Format = &format

	return s
}

func (s *_minBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_minBucketAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_minBucketAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MinBucket = s.v

	return container
}

func (s *_minBucketAggregation) MinBucketAggregationCaster() *types.MinBucketAggregation {
	return s.v
}
