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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _statsBucketAggregation struct {
	v *types.StatsBucketAggregation
}

// A sibling pipeline aggregation which calculates a variety of stats across all
// bucket of a specified metric in a sibling aggregation.
func NewStatsBucketAggregation() *_statsBucketAggregation {

	return &_statsBucketAggregation{v: types.NewStatsBucketAggregation()}

}

func (s *_statsBucketAggregation) BucketsPath(bucketspath types.BucketsPathVariant) *_statsBucketAggregation {

	s.v.BucketsPath = *bucketspath.BucketsPathCaster()

	return s
}

func (s *_statsBucketAggregation) Format(format string) *_statsBucketAggregation {

	s.v.Format = &format

	return s
}

func (s *_statsBucketAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_statsBucketAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_statsBucketAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.StatsBucket = s.v

	return container
}

func (s *_statsBucketAggregation) StatsBucketAggregationCaster() *types.StatsBucketAggregation {
	return s.v
}
