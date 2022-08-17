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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cardinalityexecutionmode"
)

// CardinalityAggregation type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/aggregations/metric.ts#L62-L66
type CardinalityAggregation struct {
	ExecutionHint      *cardinalityexecutionmode.CardinalityExecutionMode `json:"execution_hint,omitempty"`
	Field              *Field                                             `json:"field,omitempty"`
	Missing            *Missing                                           `json:"missing,omitempty"`
	PrecisionThreshold *int                                               `json:"precision_threshold,omitempty"`
	Rehash             *bool                                              `json:"rehash,omitempty"`
	Script             *Script                                            `json:"script,omitempty"`
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

func (rb *CardinalityAggregationBuilder) ExecutionHint(executionhint cardinalityexecutionmode.CardinalityExecutionMode) *CardinalityAggregationBuilder {
	rb.v.ExecutionHint = &executionhint
	return rb
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
