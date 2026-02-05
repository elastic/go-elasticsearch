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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _bucketSelectorAggregation struct {
	v *types.BucketSelectorAggregation
}

// A parent pipeline aggregation which runs a script to determine whether the
// current bucket will be retained in the parent multi-bucket aggregation.
func NewBucketSelectorAggregation() *_bucketSelectorAggregation {

	return &_bucketSelectorAggregation{v: types.NewBucketSelectorAggregation()}

}

func (s *_bucketSelectorAggregation) Script(script types.ScriptVariant) *_bucketSelectorAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_bucketSelectorAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_bucketSelectorAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_bucketSelectorAggregation) Format(format string) *_bucketSelectorAggregation {

	s.v.Format = &format

	return s
}

func (s *_bucketSelectorAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_bucketSelectorAggregation {

	s.v.GapPolicy = &gappolicy
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
