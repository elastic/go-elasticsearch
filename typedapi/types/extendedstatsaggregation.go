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

// ExtendedStatsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L68-L70
type ExtendedStatsAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
	Sigma   *float64 `json:"sigma,omitempty"`
}

// ExtendedStatsAggregationBuilder holds ExtendedStatsAggregation struct and provides a builder API.
type ExtendedStatsAggregationBuilder struct {
	v *ExtendedStatsAggregation
}

// NewExtendedStatsAggregation provides a builder for the ExtendedStatsAggregation struct.
func NewExtendedStatsAggregationBuilder() *ExtendedStatsAggregationBuilder {
	r := ExtendedStatsAggregationBuilder{
		&ExtendedStatsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the ExtendedStatsAggregation struct
func (rb *ExtendedStatsAggregationBuilder) Build() ExtendedStatsAggregation {
	return *rb.v
}

func (rb *ExtendedStatsAggregationBuilder) Field(field Field) *ExtendedStatsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *ExtendedStatsAggregationBuilder) Format(format string) *ExtendedStatsAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *ExtendedStatsAggregationBuilder) Missing(missing *MissingBuilder) *ExtendedStatsAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *ExtendedStatsAggregationBuilder) Script(script *ScriptBuilder) *ExtendedStatsAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *ExtendedStatsAggregationBuilder) Sigma(sigma float64) *ExtendedStatsAggregationBuilder {
	rb.v.Sigma = &sigma
	return rb
}
