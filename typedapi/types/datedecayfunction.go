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
// https://github.com/elastic/elasticsearch-specification/tree/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741

package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
)

// DateDecayFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/1b56d7e58f5c59f05d1641c6d6a8117c5e01d741/specification/_types/query_dsl/compound.ts#L92-L94
type DateDecayFunction struct {
	DateDecayFunction map[Field]DecayPlacementDateMathDuration `json:"DateDecayFunction,omitempty"`
	MultiValueMode    *multivaluemode.MultiValueMode           `json:"multi_value_mode,omitempty"`
}

// DateDecayFunctionBuilder holds DateDecayFunction struct and provides a builder API.
type DateDecayFunctionBuilder struct {
	v *DateDecayFunction
}

// NewDateDecayFunction provides a builder for the DateDecayFunction struct.
func NewDateDecayFunctionBuilder() *DateDecayFunctionBuilder {
	r := DateDecayFunctionBuilder{
		&DateDecayFunction{
			DateDecayFunction: make(map[Field]DecayPlacementDateMathDuration, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DateDecayFunction struct
func (rb *DateDecayFunctionBuilder) Build() DateDecayFunction {
	return *rb.v
}

func (rb *DateDecayFunctionBuilder) DateDecayFunction(values map[Field]*DecayPlacementDateMathDurationBuilder) *DateDecayFunctionBuilder {
	tmp := make(map[Field]DecayPlacementDateMathDuration, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.DateDecayFunction = tmp
	return rb
}

func (rb *DateDecayFunctionBuilder) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *DateDecayFunctionBuilder {
	rb.v.MultiValueMode = &multivaluemode
	return rb
}
