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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/gappolicy"
)

type _cumulativeSumAggregation struct {
	v *types.CumulativeSumAggregation
}

// A parent pipeline aggregation which calculates the cumulative sum of a
// specified metric in a parent `histogram` or `date_histogram` aggregation.
func NewCumulativeSumAggregation() *_cumulativeSumAggregation {

	return &_cumulativeSumAggregation{v: types.NewCumulativeSumAggregation()}

}

// Path to the buckets that contain one set of values to correlate.
func (s *_cumulativeSumAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_cumulativeSumAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// `DecimalFormat` pattern for the output value.
// If specified, the formatted value is returned in the aggregation’s
// `value_as_string` property.
func (s *_cumulativeSumAggregation) Format(format string) *_cumulativeSumAggregation {

	s.v.Format = &format

	return s
}

// Policy to apply when gaps are found in the data.
func (s *_cumulativeSumAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_cumulativeSumAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_cumulativeSumAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.CumulativeSum = s.v

	return container
}

func (s *_cumulativeSumAggregation) CumulativeSumAggregationCaster() *types.CumulativeSumAggregation {
	return s.v
}
