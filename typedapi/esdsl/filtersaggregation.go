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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _filtersAggregation struct {
	v *types.FiltersAggregation
}

// A multi-bucket aggregation where each bucket contains the documents that
// match a query.
func NewFiltersAggregation() *_filtersAggregation {

	return &_filtersAggregation{v: types.NewFiltersAggregation()}

}

func (s *_filtersAggregation) Filters(bucketsquery types.BucketsQueryVariant) *_filtersAggregation {

	s.v.Filters = *bucketsquery.BucketsQueryCaster()

	return s
}

func (s *_filtersAggregation) Keyed(keyed bool) *_filtersAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_filtersAggregation) OtherBucket(otherbucket bool) *_filtersAggregation {

	s.v.OtherBucket = &otherbucket

	return s
}

func (s *_filtersAggregation) OtherBucketKey(otherbucketkey string) *_filtersAggregation {

	s.v.OtherBucketKey = &otherbucketkey

	return s
}

func (s *_filtersAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Filters = s.v

	return container
}

func (s *_filtersAggregation) FiltersAggregationCaster() *types.FiltersAggregation {
	return s.v
}
