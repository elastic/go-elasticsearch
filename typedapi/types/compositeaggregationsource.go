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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// CompositeAggregationSource type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L83-L88
type CompositeAggregationSource struct {
	DateHistogram *DateHistogramAggregation `json:"date_histogram,omitempty"`
	GeotileGrid   *GeoTileGridAggregation   `json:"geotile_grid,omitempty"`
	Histogram     *HistogramAggregation     `json:"histogram,omitempty"`
	Terms         *TermsAggregation         `json:"terms,omitempty"`
}

// CompositeAggregationSourceBuilder holds CompositeAggregationSource struct and provides a builder API.
type CompositeAggregationSourceBuilder struct {
	v *CompositeAggregationSource
}

// NewCompositeAggregationSource provides a builder for the CompositeAggregationSource struct.
func NewCompositeAggregationSourceBuilder() *CompositeAggregationSourceBuilder {
	r := CompositeAggregationSourceBuilder{
		&CompositeAggregationSource{},
	}

	return &r
}

// Build finalize the chain and returns the CompositeAggregationSource struct
func (rb *CompositeAggregationSourceBuilder) Build() CompositeAggregationSource {
	return *rb.v
}

func (rb *CompositeAggregationSourceBuilder) DateHistogram(datehistogram *DateHistogramAggregationBuilder) *CompositeAggregationSourceBuilder {
	v := datehistogram.Build()
	rb.v.DateHistogram = &v
	return rb
}

func (rb *CompositeAggregationSourceBuilder) GeotileGrid(geotilegrid *GeoTileGridAggregationBuilder) *CompositeAggregationSourceBuilder {
	v := geotilegrid.Build()
	rb.v.GeotileGrid = &v
	return rb
}

func (rb *CompositeAggregationSourceBuilder) Histogram(histogram *HistogramAggregationBuilder) *CompositeAggregationSourceBuilder {
	v := histogram.Build()
	rb.v.Histogram = &v
	return rb
}

func (rb *CompositeAggregationSourceBuilder) Terms(terms *TermsAggregationBuilder) *CompositeAggregationSourceBuilder {
	v := terms.Build()
	rb.v.Terms = &v
	return rb
}
