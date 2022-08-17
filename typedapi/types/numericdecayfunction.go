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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/multivaluemode"
)

// NumericDecayFunction type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/query_dsl/compound.ts#L88-L90
type NumericDecayFunction struct {
	MultiValueMode       *multivaluemode.MultiValueMode       `json:"multi_value_mode,omitempty"`
	NumericDecayFunction map[Field]DecayPlacementdoubledouble `json:"-"`
}

// MarhsalJSON overrides marshalling for types with additional properties
func (s NumericDecayFunction) MarshalJSON() ([]byte, error) {
	type opt NumericDecayFunction
	// We transform the struct to a map without the embedded additional properties map
	tmp := make(map[string]interface{}, 0)

	data, err := json.Marshal(opt(s))
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return nil, err
	}

	// We inline the additional fields from the underlying map
	for key, value := range s.NumericDecayFunction {
		tmp[string(key)] = value
	}

	data, err = json.Marshal(tmp)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// NumericDecayFunctionBuilder holds NumericDecayFunction struct and provides a builder API.
type NumericDecayFunctionBuilder struct {
	v *NumericDecayFunction
}

// NewNumericDecayFunction provides a builder for the NumericDecayFunction struct.
func NewNumericDecayFunctionBuilder() *NumericDecayFunctionBuilder {
	r := NumericDecayFunctionBuilder{
		&NumericDecayFunction{
			NumericDecayFunction: make(map[Field]DecayPlacementdoubledouble, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the NumericDecayFunction struct
func (rb *NumericDecayFunctionBuilder) Build() NumericDecayFunction {
	return *rb.v
}

func (rb *NumericDecayFunctionBuilder) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *NumericDecayFunctionBuilder {
	rb.v.MultiValueMode = &multivaluemode
	return rb
}

func (rb *NumericDecayFunctionBuilder) NumericDecayFunction(values map[Field]*DecayPlacementdoubledoubleBuilder) *NumericDecayFunctionBuilder {
	tmp := make(map[Field]DecayPlacementdoubledouble, len(values))
	for key, builder := range values {
		tmp[key] = builder.Build()
	}
	rb.v.NumericDecayFunction = tmp
	return rb
}
