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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _frequentItemSetsAggregation struct {
	v *types.FrequentItemSetsAggregation
}

// A bucket aggregation which finds frequent item sets, a form of association
// rules mining that identifies items that often occur together.
func NewFrequentItemSetsAggregation() *_frequentItemSetsAggregation {

	return &_frequentItemSetsAggregation{v: types.NewFrequentItemSetsAggregation()}

}

// Fields to analyze.
func (s *_frequentItemSetsAggregation) Fields(fields ...types.FrequentItemSetsFieldVariant) *_frequentItemSetsAggregation {

	for _, v := range fields {

		s.v.Fields = append(s.v.Fields, *v.FrequentItemSetsFieldCaster())

	}
	return s
}

// Query that filters documents from analysis.
func (s *_frequentItemSetsAggregation) Filter(filter types.QueryVariant) *_frequentItemSetsAggregation {

	s.v.Filter = filter.QueryCaster()

	return s
}

// The minimum size of one item set.
func (s *_frequentItemSetsAggregation) MinimumSetSize(minimumsetsize int) *_frequentItemSetsAggregation {

	s.v.MinimumSetSize = &minimumsetsize

	return s
}

// The minimum support of one item set.
func (s *_frequentItemSetsAggregation) MinimumSupport(minimumsupport types.Float64) *_frequentItemSetsAggregation {

	s.v.MinimumSupport = &minimumsupport

	return s
}

// The number of top item sets to return.
func (s *_frequentItemSetsAggregation) Size(size int) *_frequentItemSetsAggregation {

	s.v.Size = &size

	return s
}

func (s *_frequentItemSetsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.FrequentItemSets = s.v

	return container
}

func (s *_frequentItemSetsAggregation) FrequentItemSetsAggregationCaster() *types.FrequentItemSetsAggregation {
	return s.v
}
