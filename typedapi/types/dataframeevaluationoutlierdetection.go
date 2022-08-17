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

// DataframeEvaluationOutlierDetection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeEvaluation.ts#L46-L53
type DataframeEvaluationOutlierDetection struct {
	// ActualField The field of the index which contains the ground truth. The data type of this
	// field can be boolean or integer. If the data type is integer, the value has
	// to be either 0 (false) or 1 (true).
	ActualField Field `json:"actual_field"`
	// Metrics Specifies the metrics that are used for the evaluation.
	Metrics *DataframeEvaluationOutlierDetectionMetrics `json:"metrics,omitempty"`
	// PredictedProbabilityField The field of the index that defines the probability of whether the item
	// belongs to the class in question or not. It’s the field that contains the
	// results of the analysis.
	PredictedProbabilityField Field `json:"predicted_probability_field"`
}

// DataframeEvaluationOutlierDetectionBuilder holds DataframeEvaluationOutlierDetection struct and provides a builder API.
type DataframeEvaluationOutlierDetectionBuilder struct {
	v *DataframeEvaluationOutlierDetection
}

// NewDataframeEvaluationOutlierDetection provides a builder for the DataframeEvaluationOutlierDetection struct.
func NewDataframeEvaluationOutlierDetectionBuilder() *DataframeEvaluationOutlierDetectionBuilder {
	r := DataframeEvaluationOutlierDetectionBuilder{
		&DataframeEvaluationOutlierDetection{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationOutlierDetection struct
func (rb *DataframeEvaluationOutlierDetectionBuilder) Build() DataframeEvaluationOutlierDetection {
	return *rb.v
}

// ActualField The field of the index which contains the ground truth. The data type of this
// field can be boolean or integer. If the data type is integer, the value has
// to be either 0 (false) or 1 (true).

func (rb *DataframeEvaluationOutlierDetectionBuilder) ActualField(actualfield Field) *DataframeEvaluationOutlierDetectionBuilder {
	rb.v.ActualField = actualfield
	return rb
}

// Metrics Specifies the metrics that are used for the evaluation.

func (rb *DataframeEvaluationOutlierDetectionBuilder) Metrics(metrics *DataframeEvaluationOutlierDetectionMetricsBuilder) *DataframeEvaluationOutlierDetectionBuilder {
	v := metrics.Build()
	rb.v.Metrics = &v
	return rb
}

// PredictedProbabilityField The field of the index that defines the probability of whether the item
// belongs to the class in question or not. It’s the field that contains the
// results of the analysis.

func (rb *DataframeEvaluationOutlierDetectionBuilder) PredictedProbabilityField(predictedprobabilityfield Field) *DataframeEvaluationOutlierDetectionBuilder {
	rb.v.PredictedProbabilityField = predictedprobabilityfield
	return rb
}
