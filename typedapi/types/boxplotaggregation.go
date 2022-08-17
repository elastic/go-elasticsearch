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

// BoxplotAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L50-L52
type BoxplotAggregation struct {
	Compression *float64 `json:"compression,omitempty"`
	Field       *Field   `json:"field,omitempty"`
	Missing     *Missing `json:"missing,omitempty"`
	Script      *Script  `json:"script,omitempty"`
}

// BoxplotAggregationBuilder holds BoxplotAggregation struct and provides a builder API.
type BoxplotAggregationBuilder struct {
	v *BoxplotAggregation
}

// NewBoxplotAggregation provides a builder for the BoxplotAggregation struct.
func NewBoxplotAggregationBuilder() *BoxplotAggregationBuilder {
	r := BoxplotAggregationBuilder{
		&BoxplotAggregation{},
	}

	return &r
}

// Build finalize the chain and returns the BoxplotAggregation struct
func (rb *BoxplotAggregationBuilder) Build() BoxplotAggregation {
	return *rb.v
}

func (rb *BoxplotAggregationBuilder) Compression(compression float64) *BoxplotAggregationBuilder {
	rb.v.Compression = &compression
	return rb
}

func (rb *BoxplotAggregationBuilder) Field(field Field) *BoxplotAggregationBuilder {
	rb.v.Field = &field
	return rb
}

func (rb *BoxplotAggregationBuilder) Missing(missing *MissingBuilder) *BoxplotAggregationBuilder {
	v := missing.Build()
	rb.v.Missing = &v
	return rb
}

func (rb *BoxplotAggregationBuilder) Script(script *ScriptBuilder) *BoxplotAggregationBuilder {
	v := script.Build()
	rb.v.Script = &v
	return rb
}
