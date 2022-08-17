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

// FormatMetricAggregationBase type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L40-L42
type FormatMetricAggregationBase struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// FormatMetricAggregationBaseBuilder holds FormatMetricAggregationBase struct and provides a builder API.
type FormatMetricAggregationBaseBuilder struct {
	v *FormatMetricAggregationBase
}

// NewFormatMetricAggregationBase provides a builder for the FormatMetricAggregationBase struct.
func NewFormatMetricAggregationBaseBuilder() *FormatMetricAggregationBaseBuilder {
	r := FormatMetricAggregationBaseBuilder{
		&FormatMetricAggregationBase{},
	}

	return &r
}

// Build finalize the chain and returns the FormatMetricAggregationBase struct
func (rb *FormatMetricAggregationBaseBuilder) Build() FormatMetricAggregationBase {
	return *rb.v
}

func (rb *FormatMetricAggregationBaseBuilder) Field(field Field) *FormatMetricAggregationBaseBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *FormatMetricAggregationBaseBuilder) Format(format string) *FormatMetricAggregationBaseBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *FormatMetricAggregationBaseBuilder) Missing(missing *MissingBuilder) *FormatMetricAggregationBaseBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *FormatMetricAggregationBaseBuilder) Script(script *ScriptBuilder) *FormatMetricAggregationBaseBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
