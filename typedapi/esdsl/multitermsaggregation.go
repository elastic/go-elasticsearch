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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationcollectmode"
)

type _multiTermsAggregation struct {
	v *types.MultiTermsAggregation
}

// A multi-bucket value source based aggregation where buckets are dynamically
// built - one per unique set of values.
func NewMultiTermsAggregation() *_multiTermsAggregation {

	return &_multiTermsAggregation{v: types.NewMultiTermsAggregation()}

}

func (s *_multiTermsAggregation) CollectMode(collectmode termsaggregationcollectmode.TermsAggregationCollectMode) *_multiTermsAggregation {

	s.v.CollectMode = &collectmode
	return s
}

func (s *_multiTermsAggregation) MinDocCount(mindoccount int64) *_multiTermsAggregation {

	s.v.MinDocCount = &mindoccount

	return s
}

func (s *_multiTermsAggregation) Order(aggregateorder types.AggregateOrderVariant) *_multiTermsAggregation {

	s.v.Order = *aggregateorder.AggregateOrderCaster()

	return s
}

func (s *_multiTermsAggregation) ShardMinDocCount(shardmindoccount int64) *_multiTermsAggregation {

	s.v.ShardMinDocCount = &shardmindoccount

	return s
}

func (s *_multiTermsAggregation) ShardSize(shardsize int) *_multiTermsAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_multiTermsAggregation) ShowTermDocCountError(showtermdoccounterror bool) *_multiTermsAggregation {

	s.v.ShowTermDocCountError = &showtermdoccounterror

	return s
}

func (s *_multiTermsAggregation) Size(size int) *_multiTermsAggregation {

	s.v.Size = &size

	return s
}

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
