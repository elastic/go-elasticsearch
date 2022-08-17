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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// DataframeEvaluationRegressionMetrics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeEvaluation.ts#L92-L110
type DataframeEvaluationRegressionMetrics struct {
	// Huber Pseudo Huber loss function.
	Huber *DataframeEvaluationRegressionMetricsHuber `json:"huber,omitempty"`
	// Mse Average squared difference between the predicted values and the actual
	// (ground truth) value. For more information, read this wiki article.
	Mse map[string]interface{} `json:"mse,omitempty"`
	// Msle Average squared difference between the logarithm of the predicted values and
	// the logarithm of the actual (ground truth) value.
	Msle *DataframeEvaluationRegressionMetricsMsle `json:"msle,omitempty"`
	// RSquared Proportion of the variance in the dependent variable that is predictable from
	// the independent variables.
	RSquared map[string]interface{} `json:"r_squared,omitempty"`
}

// DataframeEvaluationRegressionMetricsBuilder holds DataframeEvaluationRegressionMetrics struct and provides a builder API.
type DataframeEvaluationRegressionMetricsBuilder struct {
	v *DataframeEvaluationRegressionMetrics
}

// NewDataframeEvaluationRegressionMetrics provides a builder for the DataframeEvaluationRegressionMetrics struct.
func NewDataframeEvaluationRegressionMetricsBuilder() *DataframeEvaluationRegressionMetricsBuilder {
	r := DataframeEvaluationRegressionMetricsBuilder{
		&DataframeEvaluationRegressionMetrics{
			Mse:      make(map[string]interface{}, 0),
			RSquared: make(map[string]interface{}, 0),
		},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationRegressionMetrics struct
func (rb *DataframeEvaluationRegressionMetricsBuilder) Build() DataframeEvaluationRegressionMetrics {
	return *rb.v
}

// Huber Pseudo Huber loss function.

func (rb *DataframeEvaluationRegressionMetricsBuilder) Huber(huber *DataframeEvaluationRegressionMetricsHuberBuilder) *DataframeEvaluationRegressionMetricsBuilder {
	v := huber.Build()
	rb.v.Huber = &v
	return rb
}

// Mse Average squared difference between the predicted values and the actual
// (ground truth) value. For more information, read this wiki article.

func (rb *DataframeEvaluationRegressionMetricsBuilder) Mse(value map[string]interface{}) *DataframeEvaluationRegressionMetricsBuilder {
	rb.v.Mse = value
	return rb
}

// Msle Average squared difference between the logarithm of the predicted values and
// the logarithm of the actual (ground truth) value.

func (rb *DataframeEvaluationRegressionMetricsBuilder) Msle(msle *DataframeEvaluationRegressionMetricsMsleBuilder) *DataframeEvaluationRegressionMetricsBuilder {
	v := msle.Build()
	rb.v.Msle = &v
	return rb
}

// RSquared Proportion of the variance in the dependent variable that is predictable from
// the independent variables.

func (rb *DataframeEvaluationRegressionMetricsBuilder) RSquared(value map[string]interface{}) *DataframeEvaluationRegressionMetricsBuilder {
	rb.v.RSquared = value
	return rb
}
