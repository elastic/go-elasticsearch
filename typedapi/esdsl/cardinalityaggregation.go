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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cardinalityexecutionmode"
)

type _cardinalityAggregation struct {
	v *types.CardinalityAggregation
}

// A single-value metrics aggregation that calculates an approximate count of
// distinct values.
func NewCardinalityAggregation() *_cardinalityAggregation {

	return &_cardinalityAggregation{v: types.NewCardinalityAggregation()}

}

func (s *_cardinalityAggregation) ExecutionHint(executionhint cardinalityexecutionmode.CardinalityExecutionMode) *_cardinalityAggregation {

	s.v.ExecutionHint = &executionhint
	return s
}

func (s *_cardinalityAggregation) Field(field string) *_cardinalityAggregation {

	s.v.Field = &field

	return s
}

func (s *_cardinalityAggregation) Missing(missing types.MissingVariant) *_cardinalityAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_cardinalityAggregation) PrecisionThreshold(precisionthreshold int) *_cardinalityAggregation {

	s.v.PrecisionThreshold = &precisionthreshold

	return s
}

func (s *_cardinalityAggregation) Rehash(rehash bool) *_cardinalityAggregation {

	s.v.Rehash = &rehash

	return s
}

func (s *_cardinalityAggregation) Script(script types.ScriptVariant) *_cardinalityAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_cardinalityAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Cardinality = s.v

	return container
}

func (s *_cardinalityAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	container := types.NewApiKeyAggregationContainer()

	container.Cardinality = s.v

	return container
}

func (s *_cardinalityAggregation) CardinalityAggregationCaster() *types.CardinalityAggregation {
	return s.v
}
