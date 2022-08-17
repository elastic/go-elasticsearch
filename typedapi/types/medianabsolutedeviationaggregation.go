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

// MedianAbsoluteDeviationAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L99-L101
type MedianAbsoluteDeviationAggregation struct {
	Compression *float64 `json:"compression,omitempty"`
	Field       *Field   `json:"field,omitempty"`
	Format      *string  `json:"format,omitempty"`
	Missing     *Missing `json:"missing,omitempty"`
	Script      *Script  `json:"script,omitempty"`
}

// MedianAbsoluteDeviationAggregationBuilder holds MedianAbsoluteDeviationAggregation struct and provides a builder API.
type MedianAbsoluteDeviationAggregationBuilder struct {
	v *MedianAbsoluteDeviationAggregation
}

// NewMedianAbsoluteDeviationAggregation provides a builder for the MedianAbsoluteDeviationAggregation struct.
func NewMedianAbsoluteDeviationAggregationBuilder() *MedianAbsoluteDeviationAggregationBuilder {
	r := MedianAbsoluteDeviationAggregationBuilder{
		&MedianAbsoluteDeviationAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the MedianAbsoluteDeviationAggregation struct
func (rb *MedianAbsoluteDeviationAggregationBuilder) Build() MedianAbsoluteDeviationAggregation {
	return *rb.v
}

func (rb *MedianAbsoluteDeviationAggregationBuilder) Compression(compression float64) *MedianAbsoluteDeviationAggregationBuilder {
	rb.v.Compression = &compression
	return rb
}

func (rb *MedianAbsoluteDeviationAggregationBuilder) Field(field Field) *MedianAbsoluteDeviationAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *MedianAbsoluteDeviationAggregationBuilder) Format(format string) *MedianAbsoluteDeviationAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *MedianAbsoluteDeviationAggregationBuilder) Missing(missing *MissingBuilder) *MedianAbsoluteDeviationAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *MedianAbsoluteDeviationAggregationBuilder) Script(script *ScriptBuilder) *MedianAbsoluteDeviationAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
