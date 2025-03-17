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

type _derivativeAggregation struct {
	v *types.DerivativeAggregation
}

// A parent pipeline aggregation which calculates the derivative of a specified
// metric in a parent `histogram` or `date_histogram` aggregation.
func NewDerivativeAggregation() *_derivativeAggregation {

	return &_derivativeAggregation{v: types.NewDerivativeAggregation()}

}

// Path to the buckets that contain one set of values to correlate.
func (s *_derivativeAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_derivativeAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// `DecimalFormat` pattern for the output value.
// If specified, the formatted value is returned in the aggregation’s
// `value_as_string` property.
func (s *_derivativeAggregation) Format(format string) *_derivativeAggregation {

	s.v.Format = &format

	return s
}

// Policy to apply when gaps are found in the data.
func (s *_derivativeAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_derivativeAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_derivativeAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Derivative = s.v

	return container
}

func (s *_derivativeAggregation) DerivativeAggregationCaster() *types.DerivativeAggregation {
	return s.v
}
