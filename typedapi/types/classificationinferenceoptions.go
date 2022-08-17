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

// ClassificationInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/ml/_types/inference.ts#L80-L95
type ClassificationInferenceOptions struct {
	// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.
	NumTopClasses *int `json:"num_top_classes,omitempty"`
	// NumTopFeatureImportanceValues Specifies the maximum number of feature importance values per document.
	NumTopFeatureImportanceValues *int `json:"num_top_feature_importance_values,omitempty"`
	// PredictionFieldType Specifies the type of the predicted field to write. Acceptable values are:
	// string, number, boolean. When boolean is provided 1.0 is transformed to true
	// and 0.0 to false.
	PredictionFieldType *string `json:"prediction_field_type,omitempty"`
	// ResultsField The field that is added to incoming documents to contain the inference
	// prediction. Defaults to predicted_value.
	ResultsField *string `json:"results_field,omitempty"`
	// TopClassesResultsField Specifies the field to which the top classes are written. Defaults to
	// top_classes.
	TopClassesResultsField *string `json:"top_classes_results_field,omitempty"`
}

// ClassificationInferenceOptionsBuilder holds ClassificationInferenceOptions struct and provides a builder API.
type ClassificationInferenceOptionsBuilder struct {
	v *ClassificationInferenceOptions
}

// NewClassificationInferenceOptions provides a builder for the ClassificationInferenceOptions struct.
func NewClassificationInferenceOptionsBuilder() *ClassificationInferenceOptionsBuilder {
	r := ClassificationInferenceOptionsBuilder{
		&ClassificationInferenceOptions{},
	}

	return &r
}

// Build finalize the chain and returns the ClassificationInferenceOptions struct
func (rb *ClassificationInferenceOptionsBuilder) Build() ClassificationInferenceOptions {
	return *rb.v
}

// NumTopClasses Specifies the number of top class predictions to return. Defaults to 0.

func (rb *ClassificationInferenceOptionsBuilder) NumTopClasses(numtopclasses int) *ClassificationInferenceOptionsBuilder {
	rb.v.NumTopClasses = &numtopclasses
	return rb
}

// NumTopFeatureImportanceValues Specifies the maximum number of feature importance values per document.

func (rb *ClassificationInferenceOptionsBuilder) NumTopFeatureImportanceValues(numtopfeatureimportancevalues int) *ClassificationInferenceOptionsBuilder {
	rb.v.NumTopFeatureImportanceValues = &numtopfeatureimportancevalues
	return rb
}

// PredictionFieldType Specifies the type of the predicted field to write. Acceptable values are:
// string, number, boolean. When boolean is provided 1.0 is transformed to true
// and 0.0 to false.

func (rb *ClassificationInferenceOptionsBuilder) PredictionFieldType(predictionfieldtype string) *ClassificationInferenceOptionsBuilder {
	rb.v.PredictionFieldType = &predictionfieldtype
	return rb
}

// ResultsField The field that is added to incoming documents to contain the inference
// prediction. Defaults to predicted_value.

func (rb *ClassificationInferenceOptionsBuilder) ResultsField(resultsfield string) *ClassificationInferenceOptionsBuilder {
	rb.v.ResultsField = &resultsfield
	return rb
}

// TopClassesResultsField Specifies the field to which the top classes are written. Defaults to
// top_classes.

func (rb *ClassificationInferenceOptionsBuilder) TopClassesResultsField(topclassesresultsfield string) *ClassificationInferenceOptionsBuilder {
	rb.v.TopClassesResultsField = &topclassesresultsfield
	return rb
}
