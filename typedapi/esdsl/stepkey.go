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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _stepKey struct {
	v *types.StepKey
}

func NewStepKey(phase string) *_stepKey {

	tmp := &_stepKey{v: types.NewStepKey()}

	tmp.Phase(phase)

	return tmp

}

// The optional action to which the index will be moved.
func (s *_stepKey) Action(action string) *_stepKey {

	s.v.Action = &action

	return s
}

// The optional step name to which the index will be moved.
func (s *_stepKey) Name(name string) *_stepKey {

	s.v.Name = &name

	return s
}

func (s *_stepKey) Phase(phase string) *_stepKey {

	s.v.Phase = phase

	return s
}

func (s *_stepKey) StepKeyCaster() *types.StepKey {
	return s.v
}
