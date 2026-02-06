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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _compositeAggregationSource struct {
	v *types.CompositeAggregationSource
}

func NewCompositeAggregationSource() *_compositeAggregationSource {
	return &_compositeAggregationSource{v: types.NewCompositeAggregationSource()}
}

func (s *_compositeAggregationSource) DateHistogram(datehistogram types.CompositeDateHistogramAggregationVariant) *_compositeAggregationSource {

	s.v.DateHistogram = datehistogram.CompositeDateHistogramAggregationCaster()

	return s
}

func (s *_compositeAggregationSource) GeotileGrid(geotilegrid types.CompositeGeoTileGridAggregationVariant) *_compositeAggregationSource {

	s.v.GeotileGrid = geotilegrid.CompositeGeoTileGridAggregationCaster()

	return s
}

func (s *_compositeAggregationSource) Histogram(histogram types.CompositeHistogramAggregationVariant) *_compositeAggregationSource {

	s.v.Histogram = histogram.CompositeHistogramAggregationCaster()

	return s
}

func (s *_compositeAggregationSource) Terms(terms types.CompositeTermsAggregationVariant) *_compositeAggregationSource {

	s.v.Terms = terms.CompositeTermsAggregationCaster()

	return s
}

func (s *_compositeAggregationSource) CompositeAggregationSourceCaster() *types.CompositeAggregationSource {
	return s.v
}
