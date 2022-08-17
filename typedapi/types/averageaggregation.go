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

// AverageAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L48-L48
type AverageAggregation struct {
	Field   *Field   `json:"field,omitempty"`
	Format  *string  `json:"format,omitempty"`
	Missing *Missing `json:"missing,omitempty"`
	Script  *Script  `json:"script,omitempty"`
}

// AverageAggregationBuilder holds AverageAggregation struct and provides a builder API.
type AverageAggregationBuilder struct {
	v *AverageAggregation
}

// NewAverageAggregation provides a builder for the AverageAggregation struct.
func NewAverageAggregationBuilder() *AverageAggregationBuilder {
	r := AverageAggregationBuilder{
		&AverageAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the AverageAggregation struct
func (rb *AverageAggregationBuilder) Build() AverageAggregation {
	return *rb.v
}

func (rb *AverageAggregationBuilder) Field(field Field) *AverageAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *AverageAggregationBuilder) Format(format string) *AverageAggregationBuilder {
	rb.v.Format = &format
	return rb
}

func (rb *AverageAggregationBuilder) Missing(missing *MissingBuilder) *AverageAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *AverageAggregationBuilder) Script(script *ScriptBuilder) *AverageAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
