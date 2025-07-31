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

// DetectorUpdate type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/ml/_types/Detector.ts#L127-L143
type DetectorUpdate struct {
	// CustomRules An array of custom rule objects, which enable you to customize the way
	// detectors operate.
	// For example, a rule may dictate to the detector conditions under which
	// results should be skipped.
	// Kibana refers to custom rules as job rules.
	CustomRules []DetectionRule `json:"custom_rules,omitempty"`
	// Description A description of the detector.
	Description *string `json:"description,omitempty"`
	// DetectorIndex A unique identifier for the detector.
	// This identifier is based on the order of the detectors in the
	// `analysis_config`, starting at zero.
	DetectorIndex int `json:"detector_index"`
}

func (s *DetectorUpdate) UnmarshalJSON(data []byte) error {

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

		case "custom_rules":
			if err := dec.Decode(&s.CustomRules); err != nil {
				return fmt.Errorf("%s | %w", "CustomRules", err)
			}

		case "description":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Description", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Description = &o

		case "detector_index":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DetectorIndex", err)
				}
				s.DetectorIndex = value
			case float64:
				f := int(v)
				s.DetectorIndex = f
			}

		}
	}
	return nil
}

// NewDetectorUpdate returns a DetectorUpdate.
func NewDetectorUpdate() *DetectorUpdate {
	r := &DetectorUpdate{}

	return r
}
