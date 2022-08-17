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

// DateRangeAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L128-L135
type DateRangeAggregation struct {
	Field    *Field                `json:"field,omitempty"`
	Format   *string               `json:"format,omitempty"`
	Keyed    *bool                 `json:"keyed,omitempty"`
	Meta     *Metadata             `json:"meta,omitempty"`
	Missing  *Missing              `json:"missing,omitempty"`
	Name     *string               `json:"name,omitempty"`
	Ranges   []DateRangeExpression `json:"ranges,omitempty"`
	TimeZone *TimeZone             `json:"time_zone,omitempty"`
}

// DateRangeAggregationBuilder holds DateRangeAggregation struct and provides a builder API.
type DateRangeAggregationBuilder struct {
	v *DateRangeAggregation
}

// NewDateRangeAggregation provides a builder for the DateRangeAggregation struct.
func NewDateRangeAggregationBuilder() *DateRangeAggregationBuilder {
	r := DateRangeAggregationBuilder{
		&DateRangeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the DateRangeAggregation struct
func (rb *DateRangeAggregationBuilder) Build() DateRangeAggregation {
	return *rb.v
}

func (rb *DateRangeAggregationBuilder) Field(field Field) *DateRangeAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *DateRangeAggregationBuilder) Format(format string) *DateRangeAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *DateRangeAggregationBuilder) Keyed(keyed bool) *DateRangeAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

func (rb *DateRangeAggregationBuilder) Meta(meta *MetadataBuilder) *DateRangeAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *DateRangeAggregationBuilder) Missing(missing *MissingBuilder) *DateRangeAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *DateRangeAggregationBuilder) Name(name string) *DateRangeAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *DateRangeAggregationBuilder) Ranges(ranges []DateRangeExpressionBuilder) *DateRangeAggregationBuilder {
	tmp := make([]DateRangeExpression, len(ranges))
	for _, value := range ranges {
		tmp = append(tmp, value.Build())
	}
	rb.v.Ranges = tmp
	return rb
}

func (rb *DateRangeAggregationBuilder) TimeZone(timezone TimeZone) *DateRangeAggregationBuilder {
	rb.v.TimeZone = &timezone
	return rb
}
