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

type _dataframeEvaluationRegressionMetricsHuber struct {
	v *types.DataframeEvaluationRegressionMetricsHuber
}

func NewDataframeEvaluationRegressionMetricsHuber() *_dataframeEvaluationRegressionMetricsHuber {

	return &_dataframeEvaluationRegressionMetricsHuber{v: types.NewDataframeEvaluationRegressionMetricsHuber()}

}

// Approximates 1/2 (prediction - actual)2 for values much less than delta and
// approximates a straight line with slope delta for values much larger than
// delta. Defaults to 1. Delta needs to be greater than 0.
func (s *_dataframeEvaluationRegressionMetricsHuber) Delta(delta types.Float64) *_dataframeEvaluationRegressionMetricsHuber {

	s.v.Delta = &delta

	return s
}

func (s *_dataframeEvaluationRegressionMetricsHuber) DataframeEvaluationRegressionMetricsHuberCaster() *types.DataframeEvaluationRegressionMetricsHuber {
	return s.v
}
