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

// ClassificationInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/inference.ts#L80-L95
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

// NewClassificationInferenceOptions returns a ClassificationInferenceOptions.
func NewClassificationInferenceOptions() *ClassificationInferenceOptions {
	r := &ClassificationInferenceOptions{}

	return r
}
