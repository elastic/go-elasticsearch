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

// GeoBoundsAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L72-L74
type GeoBoundsAggregation struct {
	Field         *Field   `json:"field,omitempty"`
	Missing       *Missing `json:"missing,omitempty"`
	Script        *Script  `json:"script,omitempty"`
	WrapLongitude *bool    `json:"wrap_longitude,omitempty"`
}

// GeoBoundsAggregationBuilder holds GeoBoundsAggregation struct and provides a builder API.
type GeoBoundsAggregationBuilder struct {
	v *GeoBoundsAggregation
}

// NewGeoBoundsAggregation provides a builder for the GeoBoundsAggregation struct.
func NewGeoBoundsAggregationBuilder() *GeoBoundsAggregationBuilder {
	r := GeoBoundsAggregationBuilder{
		&GeoBoundsAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the GeoBoundsAggregation struct
func (rb *GeoBoundsAggregationBuilder) Build() GeoBoundsAggregation {
	return *rb.v
}

func (rb *GeoBoundsAggregationBuilder) Field(field Field) *GeoBoundsAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *GeoBoundsAggregationBuilder) Missing(missing *MissingBuilder) *GeoBoundsAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *GeoBoundsAggregationBuilder) Script(script *ScriptBuilder) *GeoBoundsAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}

func (rb *GeoBoundsAggregationBuilder) WrapLongitude(wraplongitude bool) *GeoBoundsAggregationBuilder {
	rb.v.WrapLongitude = &wraplongitude
	return rb
}
