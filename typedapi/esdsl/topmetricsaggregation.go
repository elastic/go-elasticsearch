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

type _topMetricsAggregation struct {
	v *types.TopMetricsAggregation
}

// A metric aggregation that selects metrics from the document with the largest
// or smallest sort value.
func NewTopMetricsAggregation() *_topMetricsAggregation {

	return &_topMetricsAggregation{v: types.NewTopMetricsAggregation()}

}

func (s *_topMetricsAggregation) Metrics(metrics ...types.TopMetricsValueVariant) *_topMetricsAggregation {

	s.v.Metrics = make([]types.TopMetricsValue, len(metrics))
	for i, v := range metrics {
		s.v.Metrics[i] = *v.TopMetricsValueCaster()
	}

	return s
}

func (s *_topMetricsAggregation) Size(size int) *_topMetricsAggregation {

	s.v.Size = &size

	return s
}

func (s *_topMetricsAggregation) Sort(sorts ...types.SortCombinationsVariant) *_topMetricsAggregation {

	for _, v := range sorts {
		s.v.Sort = append(s.v.Sort, *v.SortCombinationsCaster())
	}

	return s
}

func (s *_topMetricsAggregation) Field(field string) *_topMetricsAggregation {

	s.v.Field = &field

	return s
}

func (s *_topMetricsAggregation) Missing(missing types.MissingVariant) *_topMetricsAggregation {

	s.v.Missing = *missing.MissingCaster()

	return s
}

func (s *_topMetricsAggregation) Script(script types.ScriptVariant) *_topMetricsAggregation {

	s.v.Script = script.ScriptCaster()

	return s
}

func (s *_topMetricsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.TopMetrics = s.v

	return container
}

func (s *_topMetricsAggregation) TopMetricsAggregationCaster() *types.TopMetricsAggregation {
	return s.v
}
