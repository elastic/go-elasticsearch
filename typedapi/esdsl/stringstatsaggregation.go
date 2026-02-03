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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stringStatsAggregation struct {
	v *types.StringStatsAggregation
}

// A multi-value metrics aggregation that computes statistics over string values
// extracted from the aggregated documents.
func NewStringStatsAggregation() *_stringStatsAggregation {

	return &_stringStatsAggregation{v: types.NewStringStatsAggregation()}

}

func (s *_stringStatsAggregation) ShowDistribution(showdistribution bool) *_stringStatsAggregation {

	s.v.ShowDistribution = &showdistribution

	return s
}

func (s *_stringStatsAggregation) Field(field string) *_stringStatsAggregation {

	s.v.Field = &field

	return s
}

func (s *_stringStatsAggregation) Missing(missing types.MissingVariant) *_stringStatsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_stringStatsAggregation) Script(script types.ScriptVariant) *_stringStatsAggregation {

	s.v.Script = script.ScriptCaster()

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
