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

type _extendedStatsAggregation struct {
	v *types.ExtendedStatsAggregation
}

// A multi-value metrics aggregation that computes stats over numeric values
// extracted from the aggregated documents.
func NewExtendedStatsAggregation() *_extendedStatsAggregation {

	return &_extendedStatsAggregation{v: types.NewExtendedStatsAggregation()}

}

func (s *_extendedStatsAggregation) Sigma(sigma types.Float64) *_extendedStatsAggregation {

	s.v.Sigma = &sigma

	return s
}

func (s *_extendedStatsAggregation) Field(field string) *_extendedStatsAggregation {

	s.v.Field = &field

	return s
}

func (s *_extendedStatsAggregation) Format(format string) *_extendedStatsAggregation {

	s.v.Format = &format

	return s
}

func (s *_extendedStatsAggregation) Missing(missing types.MissingVariant) *_extendedStatsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_extendedStatsAggregation) Script(script types.ScriptVariant) *_extendedStatsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_extendedStatsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.ExtendedStats = s.v

	return container
}

func (s *_extendedStatsAggregation) ExtendedStatsAggregationCaster() *types.ExtendedStatsAggregation {
	return s.v
}
