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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sumAggregation struct {
	v *types.SumAggregation
}

// A single-value metrics aggregation that sums numeric values that are
// extracted from the aggregated documents.
func NewSumAggregation() *_sumAggregation {

	return &_sumAggregation{v: types.NewSumAggregation()}

}

func (s *_sumAggregation) Field(field string) *_sumAggregation {

	s.v.Field = &field

	return s
}

func (s *_sumAggregation) Format(format string) *_sumAggregation {

	s.v.Format = &format

	return s
}

func (s *_sumAggregation) Missing(missing types.MissingVariant) *_sumAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_sumAggregation) Script(script types.ScriptVariant) *_sumAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_sumAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Sum = s.v

	return container
}

func (s *_sumAggregation) SumAggregationCaster() *types.SumAggregation {
	return s.v
}
