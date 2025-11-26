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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _functionScore struct {
	v *types.FunctionScore
}

func NewFunctionScore() *_functionScore {
	return &_functionScore{v: types.NewFunctionScore()}
}

func (s *_functionScore) Exp(decayfunction types.DecayFunctionVariant) *_functionScore {

	s.v.Exp = *decayfunction.DecayFunctionCaster()

	return s
}

func (s *_functionScore) FieldValueFactor(fieldvaluefactor types.FieldValueFactorScoreFunctionVariant) *_functionScore {

	s.v.FieldValueFactor = fieldvaluefactor.FieldValueFactorScoreFunctionCaster()

	return s
}

func (s *_functionScore) Filter(filter types.QueryVariant) *_functionScore {

	s.v.Filter = filter.QueryCaster()

	return s
}

func (s *_functionScore) Gauss(decayfunction types.DecayFunctionVariant) *_functionScore {

	s.v.Gauss = *decayfunction.DecayFunctionCaster()

	return s
}

func (s *_functionScore) Linear(decayfunction types.DecayFunctionVariant) *_functionScore {

	s.v.Linear = *decayfunction.DecayFunctionCaster()

	return s
}

func (s *_functionScore) RandomScore(randomscore types.RandomScoreFunctionVariant) *_functionScore {

	s.v.RandomScore = randomscore.RandomScoreFunctionCaster()

	return s
}

func (s *_functionScore) ScriptScore(scriptscore types.ScriptScoreFunctionVariant) *_functionScore {

	s.v.ScriptScore = scriptscore.ScriptScoreFunctionCaster()

	return s
}

func (s *_functionScore) Weight(weight types.Float64) *_functionScore {

	s.v.Weight = &weight

	return s
}

func (s *_functionScore) FunctionScoreCaster() *types.FunctionScore {
	return s.v
}
