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

// DataframeEvaluationRegressionMetricsHuber type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeEvaluation.ts#L117-L120
type DataframeEvaluationRegressionMetricsHuber struct {
	// Delta Approximates 1/2 (prediction - actual)2 for values much less than delta and
	// approximates a straight line with slope delta for values much larger than
	// delta. Defaults to 1. Delta needs to be greater than 0.
	Delta *float64 `json:"delta,omitempty"`
}

// DataframeEvaluationRegressionMetricsHuberBuilder holds DataframeEvaluationRegressionMetricsHuber struct and provides a builder API.
type DataframeEvaluationRegressionMetricsHuberBuilder struct {
	v *DataframeEvaluationRegressionMetricsHuber
}

// NewDataframeEvaluationRegressionMetricsHuber provides a builder for the DataframeEvaluationRegressionMetricsHuber struct.
func NewDataframeEvaluationRegressionMetricsHuberBuilder() *DataframeEvaluationRegressionMetricsHuberBuilder {
	r := DataframeEvaluationRegressionMetricsHuberBuilder{
		&DataframeEvaluationRegressionMetricsHuber{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationRegressionMetricsHuber struct
func (rb *DataframeEvaluationRegressionMetricsHuberBuilder) Build() DataframeEvaluationRegressionMetricsHuber {
	return *rb.v
}

// Delta Approximates 1/2 (prediction - actual)2 for values much less than delta and
// approximates a straight line with slope delta for values much larger than
// delta. Defaults to 1. Delta needs to be greater than 0.

func (rb *DataframeEvaluationRegressionMetricsHuberBuilder) Delta(delta float64) *DataframeEvaluationRegressionMetricsHuberBuilder {
	rb.v.Delta = &delta
	return rb
}
