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
	"strconv"
)

// DataframeEvaluationClassificationMetricsAucRoc type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/DataframeEvaluation.ts#L85-L90
type DataframeEvaluationClassificationMetricsAucRoc struct {
	// ClassName Name of the only class that is treated as positive during AUC ROC
	// calculation. Other classes are treated as negative ("one-vs-all" strategy).
	// All the evaluated documents must have class_name in the list of their top
	// classes.
	ClassName *string `json:"class_name,omitempty"`
	// IncludeCurve Whether or not the curve should be returned in addition to the score. Default
	// value is false.
	IncludeCurve *bool `json:"include_curve,omitempty"`
}

func (s *DataframeEvaluationClassificationMetricsAucRoc) UnmarshalJSON(data []byte) error {

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

		case "class_name":
			if err := dec.Decode(&s.ClassName); err != nil {
				return err
			}

		case "include_curve":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IncludeCurve = &value
			case bool:
				s.IncludeCurve = &v
			}

		}
	}
	return nil
}

// NewDataframeEvaluationClassificationMetricsAucRoc returns a DataframeEvaluationClassificationMetricsAucRoc.
func NewDataframeEvaluationClassificationMetricsAucRoc() *DataframeEvaluationClassificationMetricsAucRoc {
	r := &DataframeEvaluationClassificationMetricsAucRoc{}

	return r
}
