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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

// CardinalityAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/aggregations/metric.ts#L54-L57
type CardinalityAggregation struct {
	Field              *Field   `json:"field,omitempty"`
	Missing            *Missing `json:"missing,omitempty"`
	PrecisionThreshold *int     `json:"precision_threshold,omitempty"`
	Rehash             *bool    `json:"rehash,omitempty"`
	Script             *Script  `json:"script,omitempty"`
}

// CardinalityAggregationBuilder holds CardinalityAggregation struct and provides a builder API.
type CardinalityAggregationBuilder struct {
	v *CardinalityAggregation
}

// NewCardinalityAggregation provides a builder for the CardinalityAggregation struct.
func NewCardinalityAggregationBuilder() *CardinalityAggregationBuilder {
	r := CardinalityAggregationBuilder{
		&CardinalityAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the CardinalityAggregation struct
func (rb *CardinalityAggregationBuilder) Build() CardinalityAggregation {
	return *rb.v
}

func (rb *CardinalityAggregationBuilder) Field(field Field) *CardinalityAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *CardinalityAggregationBuilder) Missing(missing *MissingBuilder) *CardinalityAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *CardinalityAggregationBuilder) PrecisionThreshold(precisionthreshold int) *CardinalityAggregationBuilder {
	rb.v.PrecisionThreshold = &precisionthreshold
	return rb
}

func (rb *CardinalityAggregationBuilder) Rehash(rehash bool) *CardinalityAggregationBuilder {
	rb.v.Rehash = &rehash
	return rb
}

func (rb *CardinalityAggregationBuilder) Script(script *ScriptBuilder) *CardinalityAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
