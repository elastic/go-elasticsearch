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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _weightedAverageValue struct {
	v *types.WeightedAverageValue
}

func NewWeightedAverageValue() *_weightedAverageValue {

	return &_weightedAverageValue{v: types.NewWeightedAverageValue()}

}

// The field from which to extract the values or weights.
func (s *_weightedAverageValue) Field(field string) *_weightedAverageValue {

	s.v.Field = &field

	return s
}

// A value or weight to use if the field is missing.
func (s *_weightedAverageValue) Missing(missing types.Float64) *_weightedAverageValue {

	s.v.Missing = &missing

	return s
}

func (s *_weightedAverageValue) Script(script types.ScriptVariant) *_weightedAverageValue {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_weightedAverageValue) WeightedAverageValueCaster() *types.WeightedAverageValue {
	return s.v
}
