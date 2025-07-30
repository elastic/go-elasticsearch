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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type _dateDecayFunction struct {
	v *types.DateDecayFunction
}

// Function that scores a document with a linear decay, depending on the
// distance of a numeric field value of the document from an origin.
func NewDateDecayFunction() *_dateDecayFunction {

	return &_dateDecayFunction{v: types.NewDateDecayFunction()}

}

func (s *_dateDecayFunction) DecayFunctionBaseDateMathDuration(decayfunctionbasedatemathduration map[string]types.DecayPlacementDateMathDuration) *_dateDecayFunction {

	s.v.DecayFunctionBaseDateMathDuration = decayfunctionbasedatemathduration
	return s
}

func (s *_dateDecayFunction) AddDecayFunctionBaseDateMathDuration(key string, value types.DecayPlacementDateMathDurationVariant) *_dateDecayFunction {

	var tmp map[string]types.DecayPlacementDateMathDuration
	if s.v.DecayFunctionBaseDateMathDuration == nil {
		s.v.DecayFunctionBaseDateMathDuration = make(map[string]types.DecayPlacementDateMathDuration)
	} else {
		tmp = s.v.DecayFunctionBaseDateMathDuration
	}

	tmp[key] = *value.DecayPlacementDateMathDurationCaster()

	s.v.DecayFunctionBaseDateMathDuration = tmp
	return s
}

func (s *_dateDecayFunction) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *_dateDecayFunction {

	s.v.MultiValueMode = &multivaluemode
	return s
}

func (s *_dateDecayFunction) FunctionScoreCaster() *types.FunctionScore {
	container := types.NewFunctionScore()

	container.Linear = s.v

	return container
}

func (s *_dateDecayFunction) DateDecayFunctionCaster() *types.DateDecayFunction {
	return s.v
}
