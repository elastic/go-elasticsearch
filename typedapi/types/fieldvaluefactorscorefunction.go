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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldvaluefactormodifier"
)

// FieldValueFactorScoreFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L70-L75
type FieldValueFactorScoreFunction struct {
	Factor   *float64                                           `json:"factor,omitempty"`
	Field    Field                                              `json:"field"`
	Missing  *float64                                           `json:"missing,omitempty"`
	Modifier *fieldvaluefactormodifier.FieldValueFactorModifier `json:"modifier,omitempty"`
}

// FieldValueFactorScoreFunctionBuilder holds FieldValueFactorScoreFunction struct and provides a builder API.
type FieldValueFactorScoreFunctionBuilder struct {
	v *FieldValueFactorScoreFunction
}

// NewFieldValueFactorScoreFunction provides a builder for the FieldValueFactorScoreFunction struct.
func NewFieldValueFactorScoreFunctionBuilder() *FieldValueFactorScoreFunctionBuilder {
	r := FieldValueFactorScoreFunctionBuilder{
		&FieldValueFactorScoreFunction{},
	}

	return &r
}

// Build finalize the chain and returns the FieldValueFactorScoreFunction struct
func (rb *FieldValueFactorScoreFunctionBuilder) Build() FieldValueFactorScoreFunction {
	return *rb.v
}

func (rb *FieldValueFactorScoreFunctionBuilder) Factor(factor float64) *FieldValueFactorScoreFunctionBuilder {
	rb.v.Factor = &factor
	return rb
}

func (rb *FieldValueFactorScoreFunctionBuilder) Field(field Field) *FieldValueFactorScoreFunctionBuilder {
	rb.v.Field = field
	return rb
}

func (rb *FieldValueFactorScoreFunctionBuilder) Missing(missing float64) *FieldValueFactorScoreFunctionBuilder {
	rb.v.Missing = &missing
	return rb
}

func (rb *FieldValueFactorScoreFunctionBuilder) Modifier(modifier fieldvaluefactormodifier.FieldValueFactorModifier) *FieldValueFactorScoreFunctionBuilder {
	rb.v.Modifier = &modifier
	return rb
}
