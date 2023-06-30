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
)

// DataframeEvaluationClassification type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/ml/_types/DataframeEvaluation.ts#L35-L44
type DataframeEvaluationClassification struct {
	// ActualField The field of the index which contains the ground truth. The data type of this
	// field can be boolean or integer. If the data type is integer, the value has
	// to be either 0 (false) or 1 (true).
	ActualField string `json:"actual_field"`
	// Metrics Specifies the metrics that are used for the evaluation.
	Metrics *DataframeEvaluationClassificationMetrics `json:"metrics,omitempty"`
	// PredictedField The field in the index which contains the predicted value, in other words the
	// results of the classification analysis.
	PredictedField *string `json:"predicted_field,omitempty"`
	// TopClassesField The field of the index which is an array of documents of the form {
	// "class_name": XXX, "class_probability": YYY }. This field must be defined as
	// nested in the mappings.
	TopClassesField *string `json:"top_classes_field,omitempty"`
}

func (s *DataframeEvaluationClassification) UnmarshalJSON(data []byte) error {

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

		case "actual_field":
			if err := dec.Decode(&s.ActualField); err != nil {
				return err
			}

		case "metrics":
			if err := dec.Decode(&s.Metrics); err != nil {
				return err
			}

		case "predicted_field":
			if err := dec.Decode(&s.PredictedField); err != nil {
				return err
			}

		case "top_classes_field":
			if err := dec.Decode(&s.TopClassesField); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeEvaluationClassification returns a DataframeEvaluationClassification.
func NewDataframeEvaluationClassification() *DataframeEvaluationClassification {
	r := &DataframeEvaluationClassification{}

	return r
}
