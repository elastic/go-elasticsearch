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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// DataframeEvaluationOutlierDetection type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeEvaluation.ts#L46-L53
type DataframeEvaluationOutlierDetection struct {
	// ActualField The field of the index which contains the ground truth. The data type of this
	// field can be boolean or integer. If the data type is integer, the value has
	// to be either 0 (false) or 1 (true).
	ActualField string `json:"actual_field"`
	// Metrics Specifies the metrics that are used for the evaluation.
	Metrics *DataframeEvaluationOutlierDetectionMetrics `json:"metrics,omitempty"`
	// PredictedProbabilityField The field of the index that defines the probability of whether the item
	// belongs to the class in question or not. It’s the field that contains the
	// results of the analysis.
	PredictedProbabilityField string `json:"predicted_probability_field"`
}

func (s *DataframeEvaluationOutlierDetection) UnmarshalJSON(data []byte) error {

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

		case "predicted_probability_field":
			if err := dec.Decode(&s.PredictedProbabilityField); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeEvaluationOutlierDetection returns a DataframeEvaluationOutlierDetection.
func NewDataframeEvaluationOutlierDetection() *DataframeEvaluationOutlierDetection {
	r := &DataframeEvaluationOutlierDetection{}

	return r
}
