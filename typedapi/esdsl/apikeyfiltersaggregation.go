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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _apiKeyFiltersAggregation struct {
	v *types.ApiKeyFiltersAggregation
}

// A multi-bucket aggregation where each bucket contains the documents that
// match a query.
func NewApiKeyFiltersAggregation() *_apiKeyFiltersAggregation {

	return &_apiKeyFiltersAggregation{v: types.NewApiKeyFiltersAggregation()}

}

func (s *_apiKeyFiltersAggregation) Filters(bucketsapikeyquerycontainer types.BucketsApiKeyQueryContainerVariant) *_apiKeyFiltersAggregation {

	s.v.Filters = *bucketsapikeyquerycontainer.BucketsApiKeyQueryContainerCaster()

	return s
}

func (s *_apiKeyFiltersAggregation) Keyed(keyed bool) *_apiKeyFiltersAggregation {

	s.v.Keyed = &keyed

	return s
}

func (s *_apiKeyFiltersAggregation) OtherBucket(otherbucket bool) *_apiKeyFiltersAggregation {

	s.v.OtherBucket = &otherbucket

	return s
}

func (s *_apiKeyFiltersAggregation) OtherBucketKey(otherbucketkey string) *_apiKeyFiltersAggregation {

	s.v.OtherBucketKey = &otherbucketkey

	return s
}

func (s *_apiKeyFiltersAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	container := types.NewApiKeyAggregationContainer()

	container.Filters = s.v

	return container
}

func (s *_apiKeyFiltersAggregation) ApiKeyFiltersAggregationCaster() *types.ApiKeyFiltersAggregation {
	return s.v
}
