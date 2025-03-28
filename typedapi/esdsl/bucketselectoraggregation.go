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

type _bucketSelectorAggregation struct {
	v *types.BucketSelectorAggregation
}

// A parent pipeline aggregation which runs a script to determine whether the
// current bucket will be retained in the parent multi-bucket aggregation.
func NewBucketSelectorAggregation() *_bucketSelectorAggregation {

	return &_bucketSelectorAggregation{v: types.NewBucketSelectorAggregation()}

}

// Path to the buckets that contain one set of values to correlate.
func (s *_bucketSelectorAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketSelectorAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

// `DecimalFormat` pattern for the output value.
// If specified, the formatted value is returned in the aggregation’s
// `value_as_string` property.
func (s *_bucketSelectorAggregation) Format(format string) *_bucketSelectorAggregation {

	s.v.Format = &format

	return s
}

// Policy to apply when gaps are found in the data.
func (s *_bucketSelectorAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_bucketSelectorAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

// The script to run for this aggregation.
func (s *_bucketSelectorAggregation) Script(script types.ScriptVariant) *_bucketSelectorAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_bucketSelectorAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.BucketSelector = s.v

	return container
}

func (s *_bucketSelectorAggregation) BucketSelectorAggregationCaster() *types.BucketSelectorAggregation {
	return s.v
}
