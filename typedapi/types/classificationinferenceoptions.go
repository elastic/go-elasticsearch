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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// ClassificationInferenceOptions type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/inference.ts#L85-L100
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

func (s *ClassificationInferenceOptions) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "num_top_classes":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumTopClasses = &value
			case float64:
				f := int(v)
				s.NumTopClasses = &f
			}

		case "num_top_feature_importance_values":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.NumTopFeatureImportanceValues = &value
			case float64:
				f := int(v)
				s.NumTopFeatureImportanceValues = &f
			}

		case "prediction_field_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PredictionFieldType = &o

		case "results_field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResultsField = &o

		case "top_classes_results_field":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TopClassesResultsField = &o

		}
	}
	return nil
}

// NewClassificationInferenceOptions returns a ClassificationInferenceOptions.
func NewClassificationInferenceOptions() *ClassificationInferenceOptions {
	r := &ClassificationInferenceOptions{}

	return r
}
