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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// AdaptiveAllocationsSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/ml/_types/TrainedModel.ts#L109-L113
type AdaptiveAllocationsSettings struct {
	Enabled                bool `json:"enabled"`
	MaxNumberOfAllocations *int `json:"max_number_of_allocations,omitempty"`
	MinNumberOfAllocations *int `json:"min_number_of_allocations,omitempty"`
}

func (s *AdaptiveAllocationsSettings) UnmarshalJSON(data []byte) error {

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

		case "enabled":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Enabled", err)
				}
				s.Enabled = value
			case bool:
				s.Enabled = v
			}

		case "max_number_of_allocations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxNumberOfAllocations", err)
				}
				s.MaxNumberOfAllocations = &value
			case float64:
				f := int(v)
				s.MaxNumberOfAllocations = &f
			}

		case "min_number_of_allocations":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinNumberOfAllocations", err)
				}
				s.MinNumberOfAllocations = &value
			case float64:
				f := int(v)
				s.MinNumberOfAllocations = &f
			}

		}
	}
	return nil
}

// NewAdaptiveAllocationsSettings returns a AdaptiveAllocationsSettings.
func NewAdaptiveAllocationsSettings() *AdaptiveAllocationsSettings {
	r := &AdaptiveAllocationsSettings{}

	return r
}
