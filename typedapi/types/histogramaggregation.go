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

// HistogramAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L232-L244
type HistogramAggregation struct {
	ExtendedBounds *ExtendedBoundsdouble `json:"extended_bounds,omitempty"`
	Field          *Field                `json:"field,omitempty"`
	Format         *string               `json:"format,omitempty"`
	HardBounds     *ExtendedBoundsdouble `json:"hard_bounds,omitempty"`
	Interval       *float64              `json:"interval,omitempty"`
	Keyed          *bool                 `json:"keyed,omitempty"`
	Meta           *Metadata             `json:"meta,omitempty"`
	MinDocCount    *int                  `json:"min_doc_count,omitempty"`
	Missing        *float64              `json:"missing,omitempty"`
	Name           *string               `json:"name,omitempty"`
	Offset         *float64              `json:"offset,omitempty"`
	Order          *AggregateOrder       `json:"order,omitempty"`
	Script         *Script               `json:"script,omitempty"`
}

// HistogramAggregationBuilder holds HistogramAggregation struct and provides a builder API.
type HistogramAggregationBuilder struct {
	v *HistogramAggregation
}

// NewHistogramAggregation provides a builder for the HistogramAggregation struct.
func NewHistogramAggregationBuilder() *HistogramAggregationBuilder {
	r := HistogramAggregationBuilder{
		&HistogramAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the HistogramAggregation struct
func (rb *HistogramAggregationBuilder) Build() HistogramAggregation {
	return *rb.v
}

func (rb *HistogramAggregationBuilder) ExtendedBounds(extendedbounds *ExtendedBoundsdoubleBuilder) *HistogramAggregationBuilder {
	v := extendedbounds.Build()
	rb.v.ExtendedBounds = &v
	return rb
}

func (rb *HistogramAggregationBuilder) Field(field Field) *HistogramAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *HistogramAggregationBuilder) Format(format string) *HistogramAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *HistogramAggregationBuilder) HardBounds(hardbounds *ExtendedBoundsdoubleBuilder) *HistogramAggregationBuilder {
	v := hardbounds.Build()
	rb.v.HardBounds = &v
	return rb
}

func (rb *HistogramAggregationBuilder) Interval(interval float64) *HistogramAggregationBuilder {
	rb.v.Interval = &interval
	return rb
}

func (rb *HistogramAggregationBuilder) Keyed(keyed bool) *HistogramAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

func (rb *HistogramAggregationBuilder) Meta(meta *MetadataBuilder) *HistogramAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *HistogramAggregationBuilder) MinDocCount(mindoccount int) *HistogramAggregationBuilder {
	rb.v.MinDocCount = &mindoccount
	return rb
}

func (rb *HistogramAggregationBuilder) Missing(missing float64) *HistogramAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

func (rb *HistogramAggregationBuilder) Name(name string) *HistogramAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *HistogramAggregationBuilder) Offset(offset float64) *HistogramAggregationBuilder {
	rb.v.Offset = &offset
	return rb
}

func (rb *HistogramAggregationBuilder) Order(order *AggregateOrderBuilder) *HistogramAggregationBuilder {
	v := order.Build()
	rb.v.Order = &v
	return rb
}

func (rb *HistogramAggregationBuilder) Script(script *ScriptBuilder) *HistogramAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
