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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _apiKeyAggregationContainer struct {
	v *types.ApiKeyAggregationContainer
}

func NewApiKeyAggregationContainer() *_apiKeyAggregationContainer {
	return &_apiKeyAggregationContainer{v: types.NewApiKeyAggregationContainer()}
}

// AdditionalApiKeyAggregationContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_apiKeyAggregationContainer) AdditionalApiKeyAggregationContainerProperty(key string, value json.RawMessage) *_apiKeyAggregationContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalApiKeyAggregationContainerProperty = tmp
	return s
}

func (s *_apiKeyAggregationContainer) Aggregations(aggregations map[string]types.ApiKeyAggregationContainer) *_apiKeyAggregationContainer {

	s.v.Aggregations = aggregations
	return s
}

func (s *_apiKeyAggregationContainer) AddAggregation(key string, value types.ApiKeyAggregationContainerVariant) *_apiKeyAggregationContainer {

	var tmp map[string]types.ApiKeyAggregationContainer
	if s.v.Aggregations == nil {
		s.v.Aggregations = make(map[string]types.ApiKeyAggregationContainer)
	} else {
		tmp = s.v.Aggregations
	}

	tmp[key] = *value.ApiKeyAggregationContainerCaster()

	s.v.Aggregations = tmp
	return s
}

func (s *_apiKeyAggregationContainer) Cardinality(cardinality types.CardinalityAggregationVariant) *_apiKeyAggregationContainer {

	s.v.Cardinality = cardinality.CardinalityAggregationCaster()

	return s
}

func (s *_apiKeyAggregationContainer) Composite(composite types.CompositeAggregationVariant) *_apiKeyAggregationContainer {

	s.v.Composite = composite.CompositeAggregationCaster()

	return s
}

func (s *_apiKeyAggregationContainer) DateRange(daterange types.DateRangeAggregationVariant) *_apiKeyAggregationContainer {

	s.v.DateRange = daterange.DateRangeAggregationCaster()

	return s
}

func (s *_apiKeyAggregationContainer) Filter(filter types.ApiKeyQueryContainerVariant) *_apiKeyAggregationContainer {

	s.v.Filter = filter.ApiKeyQueryContainerCaster()

	return s
}

func (s *_apiKeyAggregationContainer) Filters(filters types.ApiKeyFiltersAggregationVariant) *_apiKeyAggregationContainer {

	s.v.Filters = filters.ApiKeyFiltersAggregationCaster()

	return s
}

func (s *_apiKeyAggregationContainer) Meta(metadata types.MetadataVariant) *_apiKeyAggregationContainer {

	s.v.Meta = *metadata.MetadataCaster()

	return s
}

func (s *_apiKeyAggregationContainer) Missing(missing types.MissingAggregationVariant) *_apiKeyAggregationContainer {

	s.v.Missing = missing.MissingAggregationCaster()

	return s
}

func (s *_apiKeyAggregationContainer) Range(range_ types.RangeAggregationVariant) *_apiKeyAggregationContainer {

	s.v.Range = range_.RangeAggregationCaster()

	return s
}

func (s *_apiKeyAggregationContainer) Terms(terms types.TermsAggregationVariant) *_apiKeyAggregationContainer {

	s.v.Terms = terms.TermsAggregationCaster()

	return s
}

func (s *_apiKeyAggregationContainer) ValueCount(valuecount types.ValueCountAggregationVariant) *_apiKeyAggregationContainer {

	s.v.ValueCount = valuecount.ValueCountAggregationCaster()

	return s
}

func (s *_apiKeyAggregationContainer) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	return s.v
}
