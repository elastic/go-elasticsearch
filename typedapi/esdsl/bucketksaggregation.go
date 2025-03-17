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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _bucketKsAggregation struct {
	v *types.BucketKsAggregation
}

// A sibling pipeline aggregation which runs a two sample Kolmogorov–Smirnov
// test ("K-S test") against a provided distribution and the distribution
// implied by the documents counts in the configured sibling aggregation.
func NewBucketKsAggregation() *_bucketKsAggregation {

	return &_bucketKsAggregation{v: types.NewBucketKsAggregation()}

}

// A list of string values indicating which K-S test alternative to calculate.
// The valid values
// are: "greater", "less", "two_sided". This parameter is key for determining
// the K-S statistic used
// when calculating the K-S test. Default value is all possible alternative
// hypotheses.
func (s *_bucketKsAggregation) Alternative(alternatives ...string) *_bucketKsAggregation {

	for _, v := range alternatives {

		s.v.Alternative = append(s.v.Alternative, v)

	}
	return s
}

// Path to the buckets that contain one set of values to correlate.
func (s *_bucketKsAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketKsAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// A list of doubles indicating the distribution of the samples with which to
// compare to the `buckets_path` results.
// In typical usage this is the overall proportion of documents in each bucket,
// which is compared with the actual
// document proportions in each bucket from the sibling aggregation counts. The
// default is to assume that overall
// documents are uniformly distributed on these buckets, which they would be if
// one used equal percentiles of a
// metric to define the bucket end points.
func (s *_bucketKsAggregation) Fractions(fractions ...types.Float64) *_bucketKsAggregation {

	for _, v := range fractions {

		s.v.Fractions = append(s.v.Fractions, v)

	}
	return s
}

// Indicates the sampling methodology when calculating the K-S test. Note, this
// is sampling of the returned values.
// This determines the cumulative distribution function (CDF) points used
// comparing the two samples. Default is
// `upper_tail`, which emphasizes the upper end of the CDF points. Valid options
// are: `upper_tail`, `uniform`,
// and `lower_tail`.
func (s *_bucketKsAggregation) SamplingMethod(samplingmethod string) *_bucketKsAggregation {

	s.v.SamplingMethod = &samplingmethod

	return s
}

func (s *_bucketKsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.BucketCountKsTest = s.v

	return container
}

func (s *_bucketKsAggregation) BucketKsAggregationCaster() *types.BucketKsAggregation {
	return s.v
}
