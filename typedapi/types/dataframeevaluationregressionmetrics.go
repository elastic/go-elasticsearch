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

package types

import (
	"encoding/json"
)

// DataframeEvaluationRegressionMetrics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/ml/_types/DataframeEvaluation.ts#L92-L110
type DataframeEvaluationRegressionMetrics struct {
	// Huber Pseudo Huber loss function.
	Huber *DataframeEvaluationRegressionMetricsHuber `json:"huber,omitempty"`
	// Mse Average squared difference between the predicted values and the actual
	// (ground truth) value. For more information, read this wiki article.
	Mse map[string]json.RawMessage `json:"mse,omitempty"`
	// Msle Average squared difference between the logarithm of the predicted values and
	// the logarithm of the actual (ground truth) value.
	Msle *DataframeEvaluationRegressionMetricsMsle `json:"msle,omitempty"`
	// RSquared Proportion of the variance in the dependent variable that is predictable from
	// the independent variables.
	RSquared map[string]json.RawMessage `json:"r_squared,omitempty"`
}

// NewDataframeEvaluationRegressionMetrics returns a DataframeEvaluationRegressionMetrics.
func NewDataframeEvaluationRegressionMetrics() *DataframeEvaluationRegressionMetrics {
	r := &DataframeEvaluationRegressionMetrics{
		Mse:      make(map[string]json.RawMessage),
		RSquared: make(map[string]json.RawMessage),
	}

	return r
}

type DataframeEvaluationRegressionMetricsVariant interface {
	DataframeEvaluationRegressionMetricsCaster() *DataframeEvaluationRegressionMetrics
}

func (s *DataframeEvaluationRegressionMetrics) DataframeEvaluationRegressionMetricsCaster() *DataframeEvaluationRegressionMetrics {
	return s
}
