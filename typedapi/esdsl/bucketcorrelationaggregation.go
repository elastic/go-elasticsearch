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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _bucketCorrelationAggregation struct {
	v *types.BucketCorrelationAggregation
}

// A sibling pipeline aggregation which runs a correlation function on the
// configured sibling multi-bucket aggregation.
func NewBucketCorrelationAggregation(function types.BucketCorrelationFunctionVariant) *_bucketCorrelationAggregation {

	tmp := &_bucketCorrelationAggregation{v: types.NewBucketCorrelationAggregation()}

	tmp.Function(function)

	return tmp

}

// Path to the buckets that contain one set of values to correlate.
func (s *_bucketCorrelationAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketCorrelationAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// The correlation function to execute.
func (s *_bucketCorrelationAggregation) Function(function types.BucketCorrelationFunctionVariant) *_bucketCorrelationAggregation {

	s.v.Function = *function.BucketCorrelationFunctionCaster()

	return s
}

func (s *_bucketCorrelationAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.BucketCorrelation = s.v

	return container
}

func (s *_bucketCorrelationAggregation) BucketCorrelationAggregationCaster() *types.BucketCorrelationAggregation {
	return s.v
}
