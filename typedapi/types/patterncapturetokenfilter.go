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

// PatternCaptureTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L374-L380
type PatternCaptureTokenFilter struct {
	// Patterns A list of regular expressions to match.
	Patterns []string `json:"patterns"`
	// PreserveOriginal If set to `true` (the default) it will emit the original token.
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
				return fmt.Errorf("%s | %w", "Patterns", err)
			}

		case "preserve_original":
			if err := dec.Decode(&s.PreserveOriginal); err != nil {
				return fmt.Errorf("%s | %w", "PreserveOriginal", err)
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
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
