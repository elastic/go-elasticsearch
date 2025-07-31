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

type _dataframeEvaluationRegression struct {
	v *types.DataframeEvaluationRegression
}

// Regression evaluation evaluates the results of a regression analysis which
// outputs a prediction of values.
func NewDataframeEvaluationRegression() *_dataframeEvaluationRegression {

	return &_dataframeEvaluationRegression{v: types.NewDataframeEvaluationRegression()}

}

func (s *_dataframeEvaluationRegression) ActualField(field string) *_dataframeEvaluationRegression {

	s.v.ActualField = field

	return s
}

func (s *_dataframeEvaluationRegression) Metrics(metrics types.DataframeEvaluationRegressionMetricsVariant) *_dataframeEvaluationRegression {

	s.v.Metrics = metrics.DataframeEvaluationRegressionMetricsCaster()

	return s
}

func (s *_dataframeEvaluationRegression) PredictedField(field string) *_dataframeEvaluationRegression {

	s.v.PredictedField = field

	return s
}

func (s *_dataframeEvaluationRegression) DataframeEvaluationContainerCaster() *types.DataframeEvaluationContainer {
	container := types.NewDataframeEvaluationContainer()

	container.Regression = s.v

	return container
}

func (s *_dataframeEvaluationRegression) DataframeEvaluationRegressionCaster() *types.DataframeEvaluationRegression {
	return s.v
}
