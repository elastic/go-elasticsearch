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

// StandardDeviationBounds type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/aggregations/Aggregate.ts#L281-L288
type StandardDeviationBounds struct {
	Lower           *Float64 `json:"lower,omitempty"`
	LowerPopulation *Float64 `json:"lower_population,omitempty"`
	LowerSampling   *Float64 `json:"lower_sampling,omitempty"`
	Upper           *Float64 `json:"upper,omitempty"`
	UpperPopulation *Float64 `json:"upper_population,omitempty"`
	UpperSampling   *Float64 `json:"upper_sampling,omitempty"`
}

func (s *StandardDeviationBounds) UnmarshalJSON(data []byte) error {

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

		case "lower":
			if err := dec.Decode(&s.Lower); err != nil {
				return fmt.Errorf("%s | %w", "Lower", err)
			}

		case "lower_population":
			if err := dec.Decode(&s.LowerPopulation); err != nil {
				return fmt.Errorf("%s | %w", "LowerPopulation", err)
			}

		case "lower_sampling":
			if err := dec.Decode(&s.LowerSampling); err != nil {
				return fmt.Errorf("%s | %w", "LowerSampling", err)
			}

		case "upper":
			if err := dec.Decode(&s.Upper); err != nil {
				return fmt.Errorf("%s | %w", "Upper", err)
			}

		case "upper_population":
			if err := dec.Decode(&s.UpperPopulation); err != nil {
				return fmt.Errorf("%s | %w", "UpperPopulation", err)
			}

		case "upper_sampling":
			if err := dec.Decode(&s.UpperSampling); err != nil {
				return fmt.Errorf("%s | %w", "UpperSampling", err)
			}

		}
	}
	return nil
}

// NewStandardDeviationBounds returns a StandardDeviationBounds.
func NewStandardDeviationBounds() *StandardDeviationBounds {
	r := &StandardDeviationBounds{}

	return r
}
