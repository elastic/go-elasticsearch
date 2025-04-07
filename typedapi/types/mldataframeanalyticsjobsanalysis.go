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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// MlDataFrameAnalyticsJobsAnalysis type.
//
// https://github.com/elastic/elasticsearch-specification/blob/60a81659be928bfe6cec53708c7f7613555a5eaf/specification/xpack/usage/types.ts#L194-L198
type MlDataFrameAnalyticsJobsAnalysis struct {
	Classification   *int `json:"classification,omitempty"`
	OutlierDetection *int `json:"outlier_detection,omitempty"`
	Regression       *int `json:"regression,omitempty"`
}

func (s *MlDataFrameAnalyticsJobsAnalysis) UnmarshalJSON(data []byte) error {

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

		case "classification":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Classification", err)
				}
				s.Classification = &value
			case float64:
				f := int(v)
				s.Classification = &f
			}

		case "outlier_detection":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OutlierDetection", err)
				}
				s.OutlierDetection = &value
			case float64:
				f := int(v)
				s.OutlierDetection = &f
			}

		case "regression":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Regression", err)
				}
				s.Regression = &value
			case float64:
				f := int(v)
				s.Regression = &f
			}

		}
	}
	return nil
}

// NewMlDataFrameAnalyticsJobsAnalysis returns a MlDataFrameAnalyticsJobsAnalysis.
func NewMlDataFrameAnalyticsJobsAnalysis() *MlDataFrameAnalyticsJobsAnalysis {
	r := &MlDataFrameAnalyticsJobsAnalysis{}

	return r
}

// false
