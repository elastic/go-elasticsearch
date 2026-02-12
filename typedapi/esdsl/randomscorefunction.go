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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _randomScoreFunction struct {
	v *types.RandomScoreFunction
}

// Generates scores that are uniformly distributed from 0 up to but not
// including 1.
// In case you want scores to be reproducible, it is possible to provide a
// `seed` and `field`.
func NewRandomScoreFunction() *_randomScoreFunction {

	return &_randomScoreFunction{v: types.NewRandomScoreFunction()}

}

func (s *_randomScoreFunction) Field(field string) *_randomScoreFunction {

	s.v.Field = &field

	return s
}

func (s *_randomScoreFunction) Seed(seed string) *_randomScoreFunction {

	s.v.Seed = &seed

	return s
}

func (s *_randomScoreFunction) FunctionScoreCaster() *types.FunctionScore {
	container := types.NewFunctionScore()

	container.RandomScore = s.v

	return container
}

func (s *_randomScoreFunction) RandomScoreFunctionCaster() *types.RandomScoreFunction {
	return s.v
}
