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

// DataframeEvaluationRegression type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeEvaluation.ts#L55-L62
type DataframeEvaluationRegression struct {
	// ActualField The field of the index which contains the ground truth. The data type of this
	// field must be numerical.
	ActualField string `json:"actual_field"`
	// Metrics Specifies the metrics that are used for the evaluation. For more information
	// on mse, msle, and huber, consult the Jupyter notebook on regression loss
	// functions.
	Metrics *DataframeEvaluationRegressionMetrics `json:"metrics,omitempty"`
	// PredictedField The field in the index that contains the predicted value, in other words the
	// results of the regression analysis.
	PredictedField string `json:"predicted_field"`
}

func (s *DataframeEvaluationRegression) UnmarshalJSON(data []byte) error {

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

		}
	}
	return nil
}

// NewDataframeEvaluationRegression returns a DataframeEvaluationRegression.
func NewDataframeEvaluationRegression() *DataframeEvaluationRegression {
	r := &DataframeEvaluationRegression{}

	return r
}
