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
)

// ExponentialAverageCalculationContext type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/Datafeed.ts#L204-L208
type ExponentialAverageCalculationContext struct {
	IncrementalMetricValueMs     Float64 `json:"incremental_metric_value_ms"`
	LatestTimestamp              *int64  `json:"latest_timestamp,omitempty"`
	PreviousExponentialAverageMs Float64 `json:"previous_exponential_average_ms,omitempty"`
}

func (s *ExponentialAverageCalculationContext) UnmarshalJSON(data []byte) error {

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

		case "incremental_metric_value_ms":
			if err := dec.Decode(&s.IncrementalMetricValueMs); err != nil {
				return fmt.Errorf("%s | %w", "IncrementalMetricValueMs", err)
			}

		case "latest_timestamp":
			if err := dec.Decode(&s.LatestTimestamp); err != nil {
				return fmt.Errorf("%s | %w", "LatestTimestamp", err)
			}

		case "previous_exponential_average_ms":
			if err := dec.Decode(&s.PreviousExponentialAverageMs); err != nil {
				return fmt.Errorf("%s | %w", "PreviousExponentialAverageMs", err)
			}

		}
	}
	return nil
}

// NewExponentialAverageCalculationContext returns a ExponentialAverageCalculationContext.
func NewExponentialAverageCalculationContext() *ExponentialAverageCalculationContext {
	r := &ExponentialAverageCalculationContext{}

	return r
}
