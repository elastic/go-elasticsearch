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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type _numericDecayFunction struct {
	v *types.NumericDecayFunction
}

// Function that scores a document with a linear decay, depending on the
// distance of a numeric field value of the document from an origin.
func NewNumericDecayFunction() *_numericDecayFunction {

	return &_numericDecayFunction{v: types.NewNumericDecayFunction()}

}

func (s *_numericDecayFunction) DecayFunctionBasedoubledouble(decayfunctionbasedoubledouble map[string]types.DecayPlacementdoubledouble) *_numericDecayFunction {

	s.v.DecayFunctionBasedoubledouble = decayfunctionbasedoubledouble
	return s
}

func (s *_numericDecayFunction) AddDecayFunctionBasedoubledouble(key string, value types.DecayPlacementdoubledoubleVariant) *_numericDecayFunction {

	var tmp map[string]types.DecayPlacementdoubledouble
	if s.v.DecayFunctionBasedoubledouble == nil {
		s.v.DecayFunctionBasedoubledouble = make(map[string]types.DecayPlacementdoubledouble)
	} else {
		tmp = s.v.DecayFunctionBasedoubledouble
	}

	tmp[key] = *value.DecayPlacementdoubledoubleCaster()

	s.v.DecayFunctionBasedoubledouble = tmp
	return s
}

func (s *_numericDecayFunction) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *_numericDecayFunction {

	s.v.MultiValueMode = &multivaluemode
	return s
}

func (s *_numericDecayFunction) FunctionScoreCaster() *types.FunctionScore {
	container := types.NewFunctionScore()

	container.Linear = s.v

	return container
}

func (s *_numericDecayFunction) NumericDecayFunctionCaster() *types.NumericDecayFunction {
	return s.v
}
