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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _maxAggregation struct {
	v *types.MaxAggregation
}

// A single-value metrics aggregation that returns the maximum value among the
// numeric values extracted from the aggregated documents.
func NewMaxAggregation() *_maxAggregation {

	return &_maxAggregation{v: types.NewMaxAggregation()}

}

func (s *_maxAggregation) Field(field string) *_maxAggregation {

	s.v.Field = &field

	return s
}

func (s *_maxAggregation) Format(format string) *_maxAggregation {

	s.v.Format = &format

	return s
}

func (s *_maxAggregation) Missing(missing types.MissingVariant) *_maxAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_maxAggregation) Script(script types.ScriptVariant) *_maxAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_maxAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Max = s.v

	return container
}

func (s *_maxAggregation) MaxAggregationCaster() *types.MaxAggregation {
	return s.v
}
