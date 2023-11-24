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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// DataframeEvaluationSummaryAucRocCurveItem type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/ml/evaluate_data_frame/types.ts#L95-L99
type DataframeEvaluationSummaryAucRocCurveItem struct {
	Fpr       Float64 `json:"fpr"`
	Threshold Float64 `json:"threshold"`
	Tpr       Float64 `json:"tpr"`
}

func (s *DataframeEvaluationSummaryAucRocCurveItem) UnmarshalJSON(data []byte) error {

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

		case "fpr":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Fpr = f
			case float64:
				f := Float64(v)
				s.Fpr = f
			}

		case "threshold":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Threshold = f
			case float64:
				f := Float64(v)
				s.Threshold = f
			}

		case "tpr":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Tpr = f
			case float64:
				f := Float64(v)
				s.Tpr = f
			}

		}
	}
	return nil
}

// NewDataframeEvaluationSummaryAucRocCurveItem returns a DataframeEvaluationSummaryAucRocCurveItem.
func NewDataframeEvaluationSummaryAucRocCurveItem() *DataframeEvaluationSummaryAucRocCurveItem {
	r := &DataframeEvaluationSummaryAucRocCurveItem{}

	return r
}
