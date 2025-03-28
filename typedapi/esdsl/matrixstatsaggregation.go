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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/sortmode"
)

type _matrixStatsAggregation struct {
	v *types.MatrixStatsAggregation
}

// A numeric aggregation that computes the following statistics over a set of
// document fields: `count`, `mean`, `variance`, `skewness`, `kurtosis`,
// `covariance`, and `covariance`.
func NewMatrixStatsAggregation() *_matrixStatsAggregation {

	return &_matrixStatsAggregation{v: types.NewMatrixStatsAggregation()}

}

// An array of fields for computing the statistics.
func (s *_matrixStatsAggregation) Fields(fields ...string) *_matrixStatsAggregation {

	s.v.Fields = fields

	return s
}

// The value to apply to documents that do not have a value.
// By default, documents without a value are ignored.
func (s *_matrixStatsAggregation) Missing(missing map[string]types.Float64) *_matrixStatsAggregation {

	s.v.Missing = missing
	return s
}

func (s *_matrixStatsAggregation) AddMissing(key string, value types.Float64) *_matrixStatsAggregation {

	var tmp map[string]types.Float64
	if s.v.Missing == nil {
		s.v.Missing = make(map[string]types.Float64)
	} else {
		tmp = s.v.Missing
	}

	tmp[key] = value

	s.v.Missing = tmp
	return s
}

// Array value the aggregation will use for array or multi-valued fields.
func (s *_matrixStatsAggregation) Mode(mode sortmode.SortMode) *_matrixStatsAggregation {

	s.v.Mode = &mode
	return s
}

func (s *_matrixStatsAggregation) AggregationsCaster() *types.Aggregations {
	container := types.NewAggregations()

	container.MatrixStats = s.v

	return container
}

func (s *_matrixStatsAggregation) MatrixStatsAggregationCaster() *types.MatrixStatsAggregation {
	return s.v
}
