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

// ConfusionMatrixThreshold type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/evaluate_data_frame/types.ts#L137-L158
type ConfusionMatrixThreshold struct {
	// FalseNegative False Negative
	FalseNegative int `json:"fn"`
	// FalsePositive False Positive
	FalsePositive int `json:"fp"`
	// TrueNegative True Negative
	TrueNegative int `json:"tn"`
	// TruePositive True Positive
	TruePositive int `json:"tp"`
}

func (s *ConfusionMatrixThreshold) UnmarshalJSON(data []byte) error {

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

		case "fn":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FalseNegative = value
			case float64:
				f := int(v)
				s.FalseNegative = f
			}

		case "fp":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.FalsePositive = value
			case float64:
				f := int(v)
				s.FalsePositive = f
			}

		case "tn":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TrueNegative = value
			case float64:
				f := int(v)
				s.TrueNegative = f
			}

		case "tp":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TruePositive = value
			case float64:
				f := int(v)
				s.TruePositive = f
			}

		}
	}
	return nil
}

// NewConfusionMatrixThreshold returns a ConfusionMatrixThreshold.
func NewConfusionMatrixThreshold() *ConfusionMatrixThreshold {
	r := &ConfusionMatrixThreshold{}

	return r
}
