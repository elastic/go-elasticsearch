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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/gappolicy"
)

type _bucketSortAggregation struct {
	v *types.BucketSortAggregation
}

// A parent pipeline aggregation which sorts the buckets of its parent
// multi-bucket aggregation.
func NewBucketSortAggregation() *_bucketSortAggregation {

	return &_bucketSortAggregation{v: types.NewBucketSortAggregation()}

}

func (s *_bucketSortAggregation) From(from int) *_bucketSortAggregation {

	s.v.From = &from

	return s
}

func (s *_bucketSortAggregation) GapPolicy(gappolicy gappolicy.GapPolicy) *_bucketSortAggregation {

	s.v.GapPolicy = &gappolicy
	return s
}

func (s *_bucketSortAggregation) Size(size int) *_bucketSortAggregation {

	s.v.Size = &size

	return s
}

func (s *_bucketSortAggregation) Sort(sorts ...types.SortCombinationsVariant) *_bucketSortAggregation {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

func (s *_bucketSortAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.BucketSort = s.v

	return container
}

func (s *_bucketSortAggregation) BucketSortAggregationCaster() *types.BucketSortAggregation {
	return s.v
}
