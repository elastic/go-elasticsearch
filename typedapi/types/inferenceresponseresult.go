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

// InferenceResponseResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L412-L459
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

// InferenceResponseResultBuilder holds InferenceResponseResult struct and provides a builder API.
type InferenceResponseResultBuilder struct {
	v *InferenceResponseResult
}

// NewInferenceResponseResult provides a builder for the InferenceResponseResult struct.
func NewInferenceResponseResultBuilder() *InferenceResponseResultBuilder {
	r := InferenceResponseResultBuilder{
		&InferenceResponseResult{},
	}

	return &r
}

// Build finalize the chain and returns the InferenceResponseResult struct
func (rb *InferenceResponseResultBuilder) Build() InferenceResponseResult {
	return *rb.v
}

// Entities If the model is trained for named entity recognition (NER) tasks, the
// response contains the recognized entities.

func (rb *InferenceResponseResultBuilder) Entities(entities []TrainedModelEntitiesBuilder) *InferenceResponseResultBuilder {
	tmp := make([]TrainedModelEntities, len(entities))
	for _, value := range entities {
		tmp = append(tmp, value.Build())
	}
	rb.v.Entities = tmp
	return rb
}

// FeatureImportance The feature importance for the inference results. Relevant only for
// classification or regression models

func (rb *InferenceResponseResultBuilder) FeatureImportance(feature_importance []TrainedModelInferenceFeatureImportanceBuilder) *InferenceResponseResultBuilder {
	tmp := make([]TrainedModelInferenceFeatureImportance, len(feature_importance))
	for _, value := range feature_importance {
		tmp = append(tmp, value.Build())
	}
	rb.v.FeatureImportance = tmp
	return rb
}

// IsTruncated Indicates whether the input text was truncated to meet the model's maximum
// sequence length limit. This property
// is present only when it is true.

func (rb *InferenceResponseResultBuilder) IsTruncated(istruncated bool) *InferenceResponseResultBuilder {
	rb.v.IsTruncated = &istruncated
	return rb
}

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

func (rb *InferenceResponseResultBuilder) PredictedValue(predicted_value ...PredictedValue) *InferenceResponseResultBuilder {
	rb.v.PredictedValue = predicted_value
	return rb
}

// PredictedValueSequence For fill mask tasks, the response contains the input text sequence with the
// mask token replaced by the predicted
// value.
// Additionally

func (rb *InferenceResponseResultBuilder) PredictedValueSequence(predictedvaluesequence string) *InferenceResponseResultBuilder {
	rb.v.PredictedValueSequence = &predictedvaluesequence
	return rb
}

// PredictionProbability Specifies a probability for the predicted value.

func (rb *InferenceResponseResultBuilder) PredictionProbability(predictionprobability float64) *InferenceResponseResultBuilder {
	rb.v.PredictionProbability = &predictionprobability
	return rb
}

// PredictionScore Specifies a confidence score for the predicted value.

func (rb *InferenceResponseResultBuilder) PredictionScore(predictionscore float64) *InferenceResponseResultBuilder {
	rb.v.PredictionScore = &predictionscore
	return rb
}

// TopClasses For fill mask, text classification, and zero shot classification tasks, the
// response contains a list of top
// class entries.

func (rb *InferenceResponseResultBuilder) TopClasses(top_classes []TopClassEntryBuilder) *InferenceResponseResultBuilder {
	tmp := make([]TopClassEntry, len(top_classes))
	for _, value := range top_classes {
		tmp = append(tmp, value.Build())
	}
	rb.v.TopClasses = tmp
	return rb
}

// Warning If the request failed, the response contains the reason for the failure.

func (rb *InferenceResponseResultBuilder) Warning(warning string) *InferenceResponseResultBuilder {
	rb.v.Warning = &warning
	return rb
}
