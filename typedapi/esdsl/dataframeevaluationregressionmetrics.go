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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _dataframeEvaluationRegressionMetrics struct {
	v *types.DataframeEvaluationRegressionMetrics
}

func NewDataframeEvaluationRegressionMetrics() *_dataframeEvaluationRegressionMetrics {

	return &_dataframeEvaluationRegressionMetrics{v: types.NewDataframeEvaluationRegressionMetrics()}

}

func (s *_dataframeEvaluationRegressionMetrics) Huber(huber types.DataframeEvaluationRegressionMetricsHuberVariant) *_dataframeEvaluationRegressionMetrics {

	s.v.Huber = huber.DataframeEvaluationRegressionMetricsHuberCaster()

	return s
}

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

func (s *_dataframeEvaluationRegressionMetrics) Msle(msle types.DataframeEvaluationRegressionMetricsMsleVariant) *_dataframeEvaluationRegressionMetrics {

	s.v.Msle = msle.DataframeEvaluationRegressionMetricsMsleCaster()

	return s
}

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
