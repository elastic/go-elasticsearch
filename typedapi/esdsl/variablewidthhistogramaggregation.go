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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _variableWidthHistogramAggregation struct {
	v *types.VariableWidthHistogramAggregation
}

// A multi-bucket aggregation similar to the histogram, except instead of
// providing an interval to use as the width of each bucket, a target number of
// buckets is provided.
func NewVariableWidthHistogramAggregation() *_variableWidthHistogramAggregation {

	return &_variableWidthHistogramAggregation{v: types.NewVariableWidthHistogramAggregation()}

}

func (s *_variableWidthHistogramAggregation) Buckets(buckets int) *_variableWidthHistogramAggregation {

	s.v.Buckets = &buckets

	return s
}

func (s *_variableWidthHistogramAggregation) Field(field string) *_variableWidthHistogramAggregation {

	s.v.Field = &field

	return s
}

func (s *_variableWidthHistogramAggregation) InitialBuffer(initialbuffer int) *_variableWidthHistogramAggregation {

	s.v.InitialBuffer = &initialbuffer

	return s
}

func (s *_variableWidthHistogramAggregation) Script(script types.ScriptVariant) *_variableWidthHistogramAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_variableWidthHistogramAggregation) ShardSize(shardsize int) *_variableWidthHistogramAggregation {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_variableWidthHistogramAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.VariableWidthHistogram = s.v

	return container
}

func (s *_variableWidthHistogramAggregation) VariableWidthHistogramAggregationCaster() *types.VariableWidthHistogramAggregation {
	return s.v
}
