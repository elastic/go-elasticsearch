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

type _statsAggregation struct {
	v *types.StatsAggregation
}

// A multi-value metrics aggregation that computes stats over numeric values
// extracted from the aggregated documents.
func NewStatsAggregation() *_statsAggregation {

	return &_statsAggregation{v: types.NewStatsAggregation()}

}

func (s *_statsAggregation) Field(field string) *_statsAggregation {

	s.v.Field = &field

	return s
}

func (s *_statsAggregation) Format(format string) *_statsAggregation {

	s.v.Format = &format

	return s
}

func (s *_statsAggregation) Missing(missing types.MissingVariant) *_statsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_statsAggregation) Script(script types.ScriptVariant) *_statsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_statsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.Stats = s.v

	return container
}

func (s *_statsAggregation) StatsAggregationCaster() *types.StatsAggregation {
	return s.v
}
