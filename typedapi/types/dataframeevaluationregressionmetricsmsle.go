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

// DataframeEvaluationRegressionMetricsMsle type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeEvaluation.ts#L112-L115
type DataframeEvaluationRegressionMetricsMsle struct {
	// Offset Defines the transition point at which you switch from minimizing quadratic
	// error to minimizing quadratic log error. Defaults to 1.
	Offset *float64 `json:"offset,omitempty"`
}

// DataframeEvaluationRegressionMetricsMsleBuilder holds DataframeEvaluationRegressionMetricsMsle struct and provides a builder API.
type DataframeEvaluationRegressionMetricsMsleBuilder struct {
	v *DataframeEvaluationRegressionMetricsMsle
}

// NewDataframeEvaluationRegressionMetricsMsle provides a builder for the DataframeEvaluationRegressionMetricsMsle struct.
func NewDataframeEvaluationRegressionMetricsMsleBuilder() *DataframeEvaluationRegressionMetricsMsleBuilder {
	r := DataframeEvaluationRegressionMetricsMsleBuilder{
		&DataframeEvaluationRegressionMetricsMsle{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationRegressionMetricsMsle struct
func (rb *DataframeEvaluationRegressionMetricsMsleBuilder) Build() DataframeEvaluationRegressionMetricsMsle {
	return *rb.v
}

// Offset Defines the transition point at which you switch from minimizing quadratic
// error to minimizing quadratic log error. Defaults to 1.

func (rb *DataframeEvaluationRegressionMetricsMsleBuilder) Offset(offset float64) *DataframeEvaluationRegressionMetricsMsleBuilder {
	rb.v.Offset = &offset
	return rb
}
