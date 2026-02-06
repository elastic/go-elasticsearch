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

type _pivotGroupByContainer struct {
	v *types.PivotGroupByContainer
}

func NewPivotGroupByContainer() *_pivotGroupByContainer {
	return &_pivotGroupByContainer{v: types.NewPivotGroupByContainer()}
}

func (s *_pivotGroupByContainer) DateHistogram(datehistogram types.DateHistogramAggregationVariant) *_pivotGroupByContainer {

	s.v.DateHistogram = datehistogram.DateHistogramAggregationCaster()

	return s
}

func (s *_pivotGroupByContainer) GeotileGrid(geotilegrid types.GeoTileGridAggregationVariant) *_pivotGroupByContainer {

	s.v.GeotileGrid = geotilegrid.GeoTileGridAggregationCaster()

	return s
}

func (s *_pivotGroupByContainer) Histogram(histogram types.HistogramAggregationVariant) *_pivotGroupByContainer {

	s.v.Histogram = histogram.HistogramAggregationCaster()

	return s
}

func (s *_pivotGroupByContainer) Terms(terms types.TermsAggregationVariant) *_pivotGroupByContainer {

	s.v.Terms = terms.TermsAggregationCaster()

	return s
}

func (s *_pivotGroupByContainer) PivotGroupByContainerCaster() *types.PivotGroupByContainer {
	return s.v
}
