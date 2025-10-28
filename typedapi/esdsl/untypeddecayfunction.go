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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/multivaluemode"
)

type _untypedDecayFunction struct {
	v *types.UntypedDecayFunction
}

// Function that scores a document with a linear decay, depending on the
// distance of a numeric field value of the document from an origin.
func NewUntypedDecayFunction() *_untypedDecayFunction {

	return &_untypedDecayFunction{v: types.NewUntypedDecayFunction()}

}

func (s *_untypedDecayFunction) DecayFunctionBase(decayfunctionbase map[string]types.DecayPlacement) *_untypedDecayFunction {

	s.v.DecayFunctionBase = decayfunctionbase
	return s
}

func (s *_untypedDecayFunction) AddDecayFunctionBase(key string, value types.DecayPlacementVariant) *_untypedDecayFunction {

	var tmp map[string]types.DecayPlacement
	if s.v.DecayFunctionBase == nil {
		s.v.DecayFunctionBase = make(map[string]types.DecayPlacement)
	} else {
		tmp = s.v.DecayFunctionBase
	}

	tmp[key] = *value.DecayPlacementCaster()

	s.v.DecayFunctionBase = tmp
	return s
}

func (s *_untypedDecayFunction) MultiValueMode(multivaluemode multivaluemode.MultiValueMode) *_untypedDecayFunction {

	s.v.MultiValueMode = &multivaluemode
	return s
}

func (s *_untypedDecayFunction) FunctionScoreCaster() *types.FunctionScore {
	container := types.NewFunctionScore()

	container.Linear = s.v

	return container
}

func (s *_untypedDecayFunction) UntypedDecayFunctionCaster() *types.UntypedDecayFunction {
	return s.v
}
