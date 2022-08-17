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

// RangeAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/bucket.ts#L285-L292
type RangeAggregation struct {
	Field   *Field             `json:"field,omitempty"`
	Format  *string            `json:"format,omitempty"`
	Keyed   *bool              `json:"keyed,omitempty"`
	Meta    *Metadata          `json:"meta,omitempty"`
	Missing *int               `json:"missing,omitempty"`
	Name    *string            `json:"name,omitempty"`
	Ranges  []AggregationRange `json:"ranges,omitempty"`
	Script  *Script            `json:"script,omitempty"`
}

// RangeAggregationBuilder holds RangeAggregation struct and provides a builder API.
type RangeAggregationBuilder struct {
	v *RangeAggregation
}

// NewRangeAggregation provides a builder for the RangeAggregation struct.
func NewRangeAggregationBuilder() *RangeAggregationBuilder {
	r := RangeAggregationBuilder{
		&RangeAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the RangeAggregation struct
func (rb *RangeAggregationBuilder) Build() RangeAggregation {
	return *rb.v
}

func (rb *RangeAggregationBuilder) Field(field Field) *RangeAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *RangeAggregationBuilder) Format(format string) *RangeAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *RangeAggregationBuilder) Keyed(keyed bool) *RangeAggregationBuilder {
	rb.v.Keyed = &keyed
	return rb
}

func (rb *RangeAggregationBuilder) Meta(meta *MetadataBuilder) *RangeAggregationBuilder {
	v := meta.Build()
	rb.v.Meta = &v
	return rb
}

func (rb *RangeAggregationBuilder) Missing(missing int) *RangeAggregationBuilder {
	rb.v.Missing = &missing
	return rb
}

func (rb *RangeAggregationBuilder) Name(name string) *RangeAggregationBuilder {
	rb.v.Name = &name
	return rb
}

func (rb *RangeAggregationBuilder) Ranges(ranges []AggregationRangeBuilder) *RangeAggregationBuilder {
	tmp := make([]AggregationRange, len(ranges))
	for _, value := range ranges {
		tmp = append(tmp, value.Build())
	}
	rb.v.Ranges = tmp
	return rb
}

func (rb *RangeAggregationBuilder) Script(script *ScriptBuilder) *RangeAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
