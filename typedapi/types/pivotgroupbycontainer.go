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

// PivotGroupByContainer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/transform/_types/Transform.ts#L70-L78
type PivotGroupByContainer struct {
	DateHistogram *DateHistogramAggregation `json:"date_histogram,omitempty"`
	GeotileGrid   *GeoTileGridAggregation   `json:"geotile_grid,omitempty"`
	Histogram     *HistogramAggregation     `json:"histogram,omitempty"`
	Terms         *TermsAggregation         `json:"terms,omitempty"`
}

// PivotGroupByContainerBuilder holds PivotGroupByContainer struct and provides a builder API.
type PivotGroupByContainerBuilder struct {
	v *PivotGroupByContainer
}

// NewPivotGroupByContainer provides a builder for the PivotGroupByContainer struct.
func NewPivotGroupByContainerBuilder() *PivotGroupByContainerBuilder {
	r := PivotGroupByContainerBuilder{
		&PivotGroupByContainer{},
	}

	return &r
}

// Build finalize the chain and returns the PivotGroupByContainer struct
func (rb *PivotGroupByContainerBuilder) Build() PivotGroupByContainer {
	return *rb.v
}

func (rb *PivotGroupByContainerBuilder) DateHistogram(datehistogram *DateHistogramAggregationBuilder) *PivotGroupByContainerBuilder {
	v := datehistogram.Build()
	rb.v.DateHistogram = &v
	return rb
}

func (rb *PivotGroupByContainerBuilder) GeotileGrid(geotilegrid *GeoTileGridAggregationBuilder) *PivotGroupByContainerBuilder {
	v := geotilegrid.Build()
	rb.v.GeotileGrid = &v
	return rb
}

func (rb *PivotGroupByContainerBuilder) Histogram(histogram *HistogramAggregationBuilder) *PivotGroupByContainerBuilder {
	v := histogram.Build()
	rb.v.Histogram = &v
	return rb
}

func (rb *PivotGroupByContainerBuilder) Terms(terms *TermsAggregationBuilder) *PivotGroupByContainerBuilder {
	v := terms.Build()
	rb.v.Terms = &v
	return rb
}
