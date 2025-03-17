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
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _dataframeEvaluationRegressionMetrics struct {
	v *types.DataframeEvaluationRegressionMetrics
}

func NewDataframeEvaluationRegressionMetrics() *_dataframeEvaluationRegressionMetrics {

	return &_dataframeEvaluationRegressionMetrics{v: types.NewDataframeEvaluationRegressionMetrics()}

}

// Pseudo Huber loss function.
func (s *_dataframeEvaluationRegressionMetrics) Huber(huber types.DataframeEvaluationRegressionMetricsHuberVariant) *_dataframeEvaluationRegressionMetrics {

	s.v.Huber = huber.DataframeEvaluationRegressionMetricsHuberCaster()

	return s
}

// Average squared difference between the predicted values and the actual
// (ground truth) value. For more information, read this wiki article.
func (s *_dataframeEvaluationRegressionMetrics) Mse(mse map[string]json.RawMessage) *_dataframeEvaluationRegressionMetrics {

	s.v.Mse = mse
	return s
}

func (s *_dataframeEvaluationRegressionMetrics) AddMse(key string, value json.RawMessage) *_dataframeEvaluationRegressionMetrics {

	var tmp map[string]json.RawMessage
	if s.v.Mse == nil {
		s.v.Mse = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.Mse
	}

	tmp[key] = value

	s.v.Mse = tmp
	return s
}

// Average squared difference between the logarithm of the predicted values and
// the logarithm of the actual (ground truth) value.
func (s *_dataframeEvaluationRegressionMetrics) Msle(msle types.DataframeEvaluationRegressionMetricsMsleVariant) *_dataframeEvaluationRegressionMetrics {

	s.v.Msle = msle.DataframeEvaluationRegressionMetricsMsleCaster()

	return s
}

// Proportion of the variance in the dependent variable that is predictable from
// the independent variables.
func (s *_dataframeEvaluationRegressionMetrics) RSquared(rsquared map[string]json.RawMessage) *_dataframeEvaluationRegressionMetrics {

	s.v.RSquared = rsquared
	return s
}

func (s *_dataframeEvaluationRegressionMetrics) AddRSquared(key string, value json.RawMessage) *_dataframeEvaluationRegressionMetrics {

	var tmp map[string]json.RawMessage
	if s.v.RSquared == nil {
		s.v.RSquared = make(map[string]json.RawMessage)
	} else {
		tmp = s.v.RSquared
	}

	tmp[key] = value

	s.v.RSquared = tmp
	return s
}

func (s *_dataframeEvaluationRegressionMetrics) DataframeEvaluationRegressionMetricsCaster() *types.DataframeEvaluationRegressionMetrics {
	return s.v
}
