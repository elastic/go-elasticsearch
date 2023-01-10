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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// InferenceResponseResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/inference.ts#L412-L459
type InferenceResponseResult struct {
	// Entities If the model is trained for named entity recognition (NER) tasks, the
	// response contains the recognized entities.
	Entities []TrainedModelEntities `json:"entities,omitempty"`
	// FeatureImportance The feature importance for the inference results. Relevant only for
	// classification or regression models
	FeatureImportance []TrainedModelInferenceFeatureImportance `json:"feature_importance,omitempty"`
	// IsTruncated Indicates whether the input text was truncated to meet the model's maximum
	// sequence length limit. This property
	// is present only when it is true.
	IsTruncated *bool `json:"is_truncated,omitempty"`
	// PredictedValue If the model is trained for a text classification or zero shot classification
	// task, the response is the
	// predicted class.
	// For named entity recognition (NER) tasks, it contains the annotated text
	// output.
	// For fill mask tasks, it contains the top prediction for replacing the mask
	// token.
	// For text embedding tasks, it contains the raw numerical text embedding
	// values.
	// For regression models, its a numerical value
	// For classification models, it may be an integer, double, boolean or string
	// depending on prediction type
	PredictedValue []PredictedValue `json:"predicted_value,omitempty"`
	// PredictedValueSequence For fill mask tasks, the response contains the input text sequence with the
	// mask token replaced by the predicted
	// value.
	// Additionally
	PredictedValueSequence *string `json:"predicted_value_sequence,omitempty"`
	// PredictionProbability Specifies a probability for the predicted value.
	PredictionProbability *float64 `json:"prediction_probability,omitempty"`
	// PredictionScore Specifies a confidence score for the predicted value.
	PredictionScore *float64 `json:"prediction_score,omitempty"`
	// TopClasses For fill mask, text classification, and zero shot classification tasks, the
	// response contains a list of top
	// class entries.
	TopClasses []TopClassEntry `json:"top_classes,omitempty"`
	// Warning If the request failed, the response contains the reason for the failure.
	Warning *string `json:"warning,omitempty"`
}

// NewInferenceResponseResult returns a InferenceResponseResult.
func NewInferenceResponseResult() *InferenceResponseResult {
	r := &InferenceResponseResult{}

	return r
}
