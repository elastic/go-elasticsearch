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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// TotalFeatureImportanceStatistics type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/TrainedModel.ts#L306-L313
type TotalFeatureImportanceStatistics struct {
	// Max The maximum importance value across all the training data for this feature.
	Max int `json:"max"`
	// MeanMagnitude The average magnitude of this feature across all the training data. This
	// value is the average of the absolute values of the importance for this
	// feature.
	MeanMagnitude Float64 `json:"mean_magnitude"`
	// Min The minimum importance value across all the training data for this feature.
	Min int `json:"min"`
}

func (s *TotalFeatureImportanceStatistics) UnmarshalJSON(data []byte) error {

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

		case "max":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Max", err)
				}
				s.Max = value
			case float64:
				f := int(v)
				s.Max = f
			}

		case "mean_magnitude":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return fmt.Errorf("%s | %w", "MeanMagnitude", err)
				}
				f := Float64(value)
				s.MeanMagnitude = f
			case float64:
				f := Float64(v)
				s.MeanMagnitude = f
			}

		case "min":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Min", err)
				}
				s.Min = value
			case float64:
				f := int(v)
				s.Min = f
			}

		}
	}
	return nil
}

// NewTotalFeatureImportanceStatistics returns a TotalFeatureImportanceStatistics.
func NewTotalFeatureImportanceStatistics() *TotalFeatureImportanceStatistics {
	r := &TotalFeatureImportanceStatistics{}

	return r
}
