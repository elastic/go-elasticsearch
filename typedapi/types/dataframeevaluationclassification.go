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

// DataframeEvaluationClassification type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/DataframeEvaluation.ts#L35-L44
type DataframeEvaluationClassification struct {
	// ActualField The field of the index which contains the ground truth. The data type of this
	// field can be boolean or integer. If the data type is integer, the value has
	// to be either 0 (false) or 1 (true).
	ActualField Field `json:"actual_field"`
	// Metrics Specifies the metrics that are used for the evaluation.
	Metrics *DataframeEvaluationClassificationMetrics `json:"metrics,omitempty"`
	// PredictedField The field in the index which contains the predicted value, in other words the
	// results of the classification analysis.
	PredictedField *Field `json:"predicted_field,omitempty"`
	// TopClassesField The field of the index which is an array of documents of the form {
	// "class_name": XXX, "class_probability": YYY }. This field must be defined as
	// nested in the mappings.
	TopClassesField *Field `json:"top_classes_field,omitempty"`
}

// DataframeEvaluationClassificationBuilder holds DataframeEvaluationClassification struct and provides a builder API.
type DataframeEvaluationClassificationBuilder struct {
	v *DataframeEvaluationClassification
}

// NewDataframeEvaluationClassification provides a builder for the DataframeEvaluationClassification struct.
func NewDataframeEvaluationClassificationBuilder() *DataframeEvaluationClassificationBuilder {
	r := DataframeEvaluationClassificationBuilder{
		&DataframeEvaluationClassification{},
	}

	return &r
}

// Build finalize the chain and returns the DataframeEvaluationClassification struct
func (rb *DataframeEvaluationClassificationBuilder) Build() DataframeEvaluationClassification {
	return *rb.v
}

// ActualField The field of the index which contains the ground truth. The data type of this
// field can be boolean or integer. If the data type is integer, the value has
// to be either 0 (false) or 1 (true).

func (rb *DataframeEvaluationClassificationBuilder) ActualField(actualfield Field) *DataframeEvaluationClassificationBuilder {
	rb.v.ActualField = actualfield
	return rb
}

// Metrics Specifies the metrics that are used for the evaluation.

func (rb *DataframeEvaluationClassificationBuilder) Metrics(metrics *DataframeEvaluationClassificationMetricsBuilder) *DataframeEvaluationClassificationBuilder {
	v := metrics.Build()
	rb.v.Metrics = &v
	return rb
}

// PredictedField The field in the index which contains the predicted value, in other words the
// results of the classification analysis.

func (rb *DataframeEvaluationClassificationBuilder) PredictedField(predictedfield Field) *DataframeEvaluationClassificationBuilder {
	rb.v.PredictedField = &predictedfield
	return rb
}

// TopClassesField The field of the index which is an array of documents of the form {
// "class_name": XXX, "class_probability": YYY }. This field must be defined as
// nested in the mappings.

func (rb *DataframeEvaluationClassificationBuilder) TopClassesField(topclassesfield Field) *DataframeEvaluationClassificationBuilder {
	rb.v.TopClassesField = &topclassesfield
	return rb
}
