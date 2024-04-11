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
// https://github.com/elastic/elasticsearch-specification/tree/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// StandardDeviationBoundsAsString type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5bf86339cd4bda77d07f6eaa6789b72f9c0279b1/specification/_types/aggregations/Aggregate.ts#L269-L276
type StandardDeviationBoundsAsString struct {
	Lower           string `json:"lower"`
	LowerPopulation string `json:"lower_population"`
	LowerSampling   string `json:"lower_sampling"`
	Upper           string `json:"upper"`
	UpperPopulation string `json:"upper_population"`
	UpperSampling   string `json:"upper_sampling"`
}

func (s *StandardDeviationBoundsAsString) UnmarshalJSON(data []byte) error {

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
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Lower", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Lower = o

		case "lower_population":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "LowerPopulation", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LowerPopulation = o

		case "lower_sampling":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "LowerSampling", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.LowerSampling = o

		case "upper":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Upper", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Upper = o

		case "upper_population":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UpperPopulation", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UpperPopulation = o

		case "upper_sampling":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UpperSampling", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UpperSampling = o

		}
	}
	return nil
}

// NewStandardDeviationBoundsAsString returns a StandardDeviationBoundsAsString.
func NewStandardDeviationBoundsAsString() *StandardDeviationBoundsAsString {
	r := &StandardDeviationBoundsAsString{}

	return r
}
