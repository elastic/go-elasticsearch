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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/fieldvaluefactormodifier"
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

// Optional factor to multiply the field value with.
func (s *_fieldValueFactorScoreFunction) Factor(factor types.Float64) *_fieldValueFactorScoreFunction {

	s.v.Factor = &factor

	return s
}

// Field to be extracted from the document.
func (s *_fieldValueFactorScoreFunction) Field(field string) *_fieldValueFactorScoreFunction {

	s.v.Field = field

	return s
}

// Value used if the document doesn’t have that field.
// The modifier and factor are still applied to it as though it were read from
// the document.
func (s *_fieldValueFactorScoreFunction) Missing(missing types.Float64) *_fieldValueFactorScoreFunction {

	s.v.Missing = &missing

	return s
}

// Modifier to apply to the field value.
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
