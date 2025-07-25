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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _minAggregation struct {
	v *types.MinAggregation
}

// A single-value metrics aggregation that returns the minimum value among
// numeric values extracted from the aggregated documents.
func NewMinAggregation() *_minAggregation {

	return &_minAggregation{v: types.NewMinAggregation()}

}

func (s *_minAggregation) Field(field string) *_minAggregation {

	s.v.Field = &field

	return s
}

func (s *_minAggregation) Format(format string) *_minAggregation {

	s.v.Format = &format

	return s
}

func (s *_minAggregation) Missing(missing types.MissingVariant) *_minAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_minAggregation) Script(script types.ScriptVariant) *_minAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_minAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Min = s.v

	return container
}

func (s *_minAggregation) MinAggregationCaster() *types.MinAggregation {
	return s.v
}
