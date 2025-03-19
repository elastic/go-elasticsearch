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

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _functionScore struct {
	v *types.FunctionScore
}

func NewFunctionScore() *_functionScore {
	return &_functionScore{v: types.NewFunctionScore()}
}

// AdditionalFunctionScoreProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_functionScore) AdditionalFunctionScoreProperty(key string, value json.RawMessage) *_functionScore {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalFunctionScoreProperty = tmp
	return s
}

// Function that scores a document with a exponential decay, depending on the
// distance of a numeric field value of the document from an origin.
func (s *_functionScore) Exp(decayfunction types.DecayFunctionVariant) *_functionScore {

	s.v.Exp = *decayfunction.DecayFunctionCaster()

	return s
}

// Function allows you to use a field from a document to influence the score.
// It’s similar to using the script_score function, however, it avoids the
// overhead of scripting.
func (s *_functionScore) FieldValueFactor(fieldvaluefactor types.FieldValueFactorScoreFunctionVariant) *_functionScore {

	s.v.FieldValueFactor = fieldvaluefactor.FieldValueFactorScoreFunctionCaster()

	return s
}

func (s *_functionScore) Filter(filter types.QueryVariant) *_functionScore {

	s.v.Filter = filter.QueryCaster()

	return s
}

// Function that scores a document with a normal decay, depending on the
// distance of a numeric field value of the document from an origin.
func (s *_functionScore) Gauss(decayfunction types.DecayFunctionVariant) *_functionScore {

	s.v.Gauss = *decayfunction.DecayFunctionCaster()

	return s
}

// Function that scores a document with a linear decay, depending on the
// distance of a numeric field value of the document from an origin.
func (s *_functionScore) Linear(decayfunction types.DecayFunctionVariant) *_functionScore {

	s.v.Linear = *decayfunction.DecayFunctionCaster()

	return s
}

// Generates scores that are uniformly distributed from 0 up to but not
// including 1.
// In case you want scores to be reproducible, it is possible to provide a
// `seed` and `field`.
func (s *_functionScore) RandomScore(randomscore types.RandomScoreFunctionVariant) *_functionScore {

	s.v.RandomScore = randomscore.RandomScoreFunctionCaster()

	return s
}

// Enables you to wrap another query and customize the scoring of it optionally
// with a computation derived from other numeric field values in the doc using a
// script expression.
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
