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
)

// PatternCaptureTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/analysis/token_filters.ts#L278-L282
type PatternCaptureTokenFilter struct {
	Patterns         []string           `json:"patterns"`
	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`
	Type             string             `json:"type,omitempty"`
	Version          *string            `json:"version,omitempty"`
}

func (s *PatternCaptureTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "patterns":
			if err := dec.Decode(&s.Patterns); err != nil {
				return err
			}

		case "preserve_original":
			if err := dec.Decode(&s.PreserveOriginal); err != nil {
				return err
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s PatternCaptureTokenFilter) MarshalJSON() ([]byte, error) {
	type innerPatternCaptureTokenFilter PatternCaptureTokenFilter
	tmp := innerPatternCaptureTokenFilter{
		Patterns:         s.Patterns,
		PreserveOriginal: s.PreserveOriginal,
		Type:             s.Type,
		Version:          s.Version,
	}

	tmp.Type = "pattern_capture"

	return json.Marshal(tmp)
}

// NewPatternCaptureTokenFilter returns a PatternCaptureTokenFilter.
func NewPatternCaptureTokenFilter() *PatternCaptureTokenFilter {
	r := &PatternCaptureTokenFilter{}

	return r
}
