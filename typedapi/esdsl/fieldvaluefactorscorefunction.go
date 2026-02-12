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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldvaluefactormodifier"
)

type _fieldValueFactorScoreFunction struct {
	v *types.FieldValueFactorScoreFunction
}

// Function allows you to use a field from a document to influence the score.
// It’s similar to using the script_score function, however, it avoids the
// overhead of scripting.
func NewFieldValueFactorScoreFunction() *_fieldValueFactorScoreFunction {

	return &_fieldValueFactorScoreFunction{v: types.NewFieldValueFactorScoreFunction()}

}

func (s *_fieldValueFactorScoreFunction) Factor(factor types.Float64) *_fieldValueFactorScoreFunction {

	s.v.Factor = &factor

	return s
}

func (s *_fieldValueFactorScoreFunction) Field(field string) *_fieldValueFactorScoreFunction {

	s.v.Field = field

	return s
}

func (s *_fieldValueFactorScoreFunction) Missing(missing types.Float64) *_fieldValueFactorScoreFunction {

	s.v.Missing = &missing

	return s
}

func (s *_fieldValueFactorScoreFunction) Modifier(modifier fieldvaluefactormodifier.FieldValueFactorModifier) *_fieldValueFactorScoreFunction {

	s.v.Modifier = &modifier
	return s
}

func (s *_fieldValueFactorScoreFunction) FunctionScoreCaster() *types.FunctionScore {
	container := types.NewFunctionScore()

	container.FieldValueFactor = s.v

	return container
}

func (s *_fieldValueFactorScoreFunction) FieldValueFactorScoreFunctionCaster() *types.FieldValueFactorScoreFunction {
	return s.v
}
