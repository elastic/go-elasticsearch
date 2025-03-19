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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/termsaggregationcollectmode"
)

type _multiTermsAggregation struct {
	v *types.MultiTermsAggregation
}

// A multi-bucket value source based aggregation where buckets are dynamically
// built - one per unique set of values.
func NewMultiTermsAggregation() *_multiTermsAggregation {

	return &_multiTermsAggregation{v: types.NewMultiTermsAggregation()}

}

// Specifies the strategy for data collection.
func (s *_multiTermsAggregation) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *_multiTermsAggregation {

	s.v.CollectMode = &collectmode
	return s
}

// The minimum number of documents in a bucket for it to be returned.
func (s *_multiTermsAggregation) MinDocCount(mindoccount int64) *_multiTermsAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

// Specifies the sort order of the buckets.
// Defaults to sorting by descending document count.
func (s *_multiTermsAggregation) Order(aggregateorder types.AggregateOrderVariant) *_multiTermsAggregation {

	s.v.Order = *aggregateorder.AggregateOrderCaster()

	return s
}

// The minimum number of documents in a bucket on each shard for it to be
// returned.
func (s *_multiTermsAggregation) ShardMinDocCount(shardmindoccount int64) *_multiTermsAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

// The number of candidate terms produced by each shard.
// By default, `shard_size` will be automatically estimated based on the number
// of shards and the `size` parameter.
func (s *_multiTermsAggregation) ShardSize(shardsize int) *_multiTermsAggregation {

	s.v.ShardSize = &shardsize

	return s
}

// Calculates the doc count error on per term basis.
func (s *_multiTermsAggregation) ShowTermDocCountError(showtermdoccounterror bool) *_multiTermsAggregation {

	s.v.ShowTermDocCountError = &showtermdoccounterror

	return s
}

// The number of term buckets should be returned out of the overall terms list.
func (s *_multiTermsAggregation) Size(size int) *_multiTermsAggregation {

	s.v.Size = &size

	return s
}

// The field from which to generate sets of terms.
func (s *_multiTermsAggregation) Terms(terms ...types.MultiTermLookupVariant) *_multiTermsAggregation {

	for _, v := range terms {

		s.v.Terms = append(s.v.Terms, *v.MultiTermLookupCaster())

	}
	return s
}

func (s *_multiTermsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MultiTerms = s.v

	return container
}

func (s *_multiTermsAggregation) MultiTermsAggregationCaster() *types.MultiTermsAggregation {
	return s.v
}
