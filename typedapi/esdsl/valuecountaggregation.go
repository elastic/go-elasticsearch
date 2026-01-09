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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _valueCountAggregation struct {
	v *types.ValueCountAggregation
}

// A single-value metrics aggregation that counts the number of values that are
// extracted from the aggregated documents.
func NewValueCountAggregation() *_valueCountAggregation {

	return &_valueCountAggregation{v: types.NewValueCountAggregation()}

}

func (s *_valueCountAggregation) Field(field string) *_valueCountAggregation {

	s.v.Field = &field

	return s
}

func (s *_valueCountAggregation) Format(format string) *_valueCountAggregation {

	s.v.Format = &format

	return s
}

func (s *_valueCountAggregation) Missing(missing types.MissingVariant) *_valueCountAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_valueCountAggregation) Script(script types.ScriptVariant) *_valueCountAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_valueCountAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.ValueCount = s.v

	return container
}

func (s *_valueCountAggregation) ApiKeyAggregationContainerCaster() *types.ApiKeyAggregationContainer {
	container := types.NewApiKeyAggregationContainer()

	container.ValueCount = s.v

	return container
}

func (s *_valueCountAggregation) ValueCountAggregationCaster() *types.ValueCountAggregation {
	return s.v
}
