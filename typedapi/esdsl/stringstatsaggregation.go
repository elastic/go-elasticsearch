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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _stringStatsAggregation struct {
	v *types.StringStatsAggregation
}

// A multi-value metrics aggregation that computes statistics over string values
// extracted from the aggregated documents.
func NewStringStatsAggregation() *_stringStatsAggregation {

	return &_stringStatsAggregation{v: types.NewStringStatsAggregation()}

}

// The field on which to run the aggregation.
func (s *_stringStatsAggregation) Field(field string) *_stringStatsAggregation {

	s.v.Field = &field

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_stringStatsAggregation) Missing(missing types.MissingVariant) *_stringStatsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_stringStatsAggregation) Script(script types.ScriptVariant) *_stringStatsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

// Shows the probability distribution for all characters.
func (s *_stringStatsAggregation) ShowDistribution(showdistribution bool) *_stringStatsAggregation {

	s.v.ShowDistribution = &showdistribution

	return s
}

func (s *_stringStatsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.StringStats = s.v

	return container
}

func (s *_stringStatsAggregation) StringStatsAggregationCaster() *types.StringStatsAggregation {
	return s.v
}
