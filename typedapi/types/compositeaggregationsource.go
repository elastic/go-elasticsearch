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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

// CompositeAggregationSource type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/bucket.ts#L151-L168
type CompositeAggregationSource struct {
	// DateHistogram A date histogram aggregation.
	DateHistogram *CompositeDateHistogramAggregation `json:"date_histogram,omitempty"`
	// GeotileGrid A geotile grid aggregation.
	GeotileGrid *CompositeGeoTileGridAggregation `json:"geotile_grid,omitempty"`
	// Histogram A histogram aggregation.
	Histogram *CompositeHistogramAggregation `json:"histogram,omitempty"`
	// Terms A terms aggregation.
	Terms *CompositeTermsAggregation `json:"terms,omitempty"`
}

// NewCompositeAggregationSource returns a CompositeAggregationSource.
func NewCompositeAggregationSource() *CompositeAggregationSource {
	r := &CompositeAggregationSource{}

	return r
}
