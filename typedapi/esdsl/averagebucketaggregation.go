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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

type _averageBucketAggregation struct {
	v *types.AverageBucketAggregation
}

// A sibling pipeline aggregation which calculates the mean value of a specified
// metric in a sibling aggregation.
// The specified metric must be numeric and the sibling aggregation must be a
// multi-bucket aggregation.
func NewAverageBucketAggregation() *_averageBucketAggregation {

	return &_averageBucketAggregation{v: types.NewAverageBucketAggregation()}

}

// Path to the buckets that contain one set of values to correlate.
func (s *_averageBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_averageBucketAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// `DecimalFormat` pattern for the output value.
// If specified, the formatted value is returned in the aggregation’s
// `value_as_string` property.
func (s *_averageBucketAggregation) Format(format string) *_averageBucketAggregation {

	s.v.Format = &format

	return s
}

// Policy to apply when gaps are found in the data.
func (s *_averageBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_averageBucketAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_averageBucketAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.AvgBucket = s.v

	return container
}

func (s *_averageBucketAggregation) AverageBucketAggregationCaster() *types.AverageBucketAggregation {
	return s.v
}
