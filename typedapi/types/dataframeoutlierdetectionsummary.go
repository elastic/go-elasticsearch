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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"encoding/json"
)

// DataframeOutlierDetectionSummary type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/ml/evaluate_data_frame/types.ts#L24-L29
type DataframeOutlierDetectionSummary struct {
	AucRoc          *DataframeEvaluationSummaryAucRoc   `json:"auc_roc,omitempty"`
	ConfusionMatrix map[string]ConfusionMatrixThreshold `json:"confusion_matrix,omitempty"`
	Precision       map[string]Float64                  `json:"precision,omitempty"`
	Recall          map[string]Float64                  `json:"recall,omitempty"`
}

func (s *DataframeOutlierDetectionSummary) UnmarshalJSON(data []byte) error {

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

		case "auc_roc":
			if err := dec.Decode(&s.AucRoc); err != nil {
				return err
			}

		case "confusion_matrix":
			if s.ConfusionMatrix == nil {
				s.ConfusionMatrix = make(map[string]ConfusionMatrixThreshold, 0)
			}
			if err := dec.Decode(&s.ConfusionMatrix); err != nil {
				return err
			}

		case "precision":
			if s.Precision == nil {
				s.Precision = make(map[string]Float64, 0)
			}
			if err := dec.Decode(&s.Precision); err != nil {
				return err
			}

		case "recall":
			if s.Recall == nil {
				s.Recall = make(map[string]Float64, 0)
			}
			if err := dec.Decode(&s.Recall); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDataframeOutlierDetectionSummary returns a DataframeOutlierDetectionSummary.
func NewDataframeOutlierDetectionSummary() *DataframeOutlierDetectionSummary {
	r := &DataframeOutlierDetectionSummary{
		ConfusionMatrix: make(map[string]ConfusionMatrixThreshold, 0),
		Precision:       make(map[string]Float64, 0),
		Recall:          make(map[string]Float64, 0),
	}

	return r
}
