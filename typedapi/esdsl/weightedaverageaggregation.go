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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/valuetype"
)

type _weightedAverageAggregation struct {
	v *types.WeightedAverageAggregation
}

// A single-value metrics aggregation that computes the weighted average of
// numeric values that are extracted from the aggregated documents.
func NewWeightedAverageAggregation() *_weightedAverageAggregation {

	return &_weightedAverageAggregation{v: types.NewWeightedAverageAggregation()}

}

// A numeric response formatter.
func (s *_weightedAverageAggregation) Format(format string) *_weightedAverageAggregation {

	s.v.Format = &format

	return s
}

// Configuration for the field that provides the values.
func (s *_weightedAverageAggregation) Value(value types.WeightedAverageValueVariant) *_weightedAverageAggregation {

	s.v.Value = value.WeightedAverageValueCaster()

	return s
}

func (s *_weightedAverageAggregation) ValueType(valuetype valuetype.ValueType) *_weightedAverageAggregation {

	s.v.ValueType = &valuetype
	return s
}

// Configuration for the field or script that provides the weights.
func (s *_weightedAverageAggregation) Weight(weight types.WeightedAverageValueVariant) *_weightedAverageAggregation {

	s.v.Weight = weight.WeightedAverageValueCaster()

	return s
}

func (s *_weightedAverageAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.WeightedAvg = s.v

	return container
}

func (s *_weightedAverageAggregation) WeightedAverageAggregationCaster() *types.WeightedAverageAggregation {
	return s.v
}
