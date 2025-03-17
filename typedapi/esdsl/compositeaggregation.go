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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _compositeAggregation struct {
	v *types.CompositeAggregation
}

// A multi-bucket aggregation that creates composite buckets from different
// sources.
// Unlike the other multi-bucket aggregations, you can use the `composite`
// aggregation to paginate *all* buckets from a multi-level aggregation
// efficiently.
func NewCompositeAggregation() *_compositeAggregation {

	return &_compositeAggregation{v: types.NewCompositeAggregation()}

}

// When paginating, use the `after_key` value returned in the previous response
// to retrieve the next page.
func (s *_compositeAggregation) After(compositeaggregatekey types.CompositeAggregateKeyVariant) *_compositeAggregation {

	s.v.After = *compositeaggregatekey.CompositeAggregateKeyCaster()

	return s
}

// The number of composite buckets that should be returned.
func (s *_compositeAggregation) Size(size int) *_compositeAggregation {

	s.v.Size = &size

	return s
}

// The value sources used to build composite buckets.
// Keys are returned in the order of the `sources` definition.
func (s *_compositeAggregation) Sources(sources []map[string]types.CompositeAggregationSource) *_compositeAggregation {

	s.v.Sources = sources

	return s
}

func (s *_compositeAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Composite = s.v

	return container
}

func (s *_compositeAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	container := types.NewApiKeyAggregationContainer()

	container.Composite = s.v

	return container
}

func (s *_compositeAggregation) CompositeAggregationCaster() *types.CompositeAggregation {
	return s.v
}
