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

// PatternReplaceTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L382-L391
type PatternReplaceTokenFilter struct {
	// All If `true`, all substrings matching the pattern parameter’s regular expression
	// are replaced. If `false`, the filter replaces only the first matching
	// substring in each token. Defaults to `true`.
	All   *bool   `json:"all,omitempty"`
	Flags *string `json:"flags,omitempty"`
	// Pattern Regular expression, written in Java’s regular expression syntax. The filter
	// replaces token substrings matching this pattern with the substring in the
	// `replacement` parameter.
	Pattern string `json:"pattern"`
	// Replacement Replacement substring. Defaults to an empty substring (`""`).
	Replacement *string `json:"replacement,omitempty"`
	Type        string  `json:"type,omitempty"`
	Version     *string `json:"version,omitempty"`
}

func (s *PatternReplaceTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "all":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "All", err)
				}
				s.All = &value
			case bool:
				s.All = &v
			}

		case "flags":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Flags", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Flags = &o

		case "pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Pattern", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Pattern = o

		case "replacement":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Replacement", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Replacement = &o

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
func (s PatternReplaceTokenFilter) MarshalJSON() ([]byte, error) {
	type innerPatternReplaceTokenFilter PatternReplaceTokenFilter
	tmp := innerPatternReplaceTokenFilter{
		All:         s.All,
		Flags:       s.Flags,
		Pattern:     s.Pattern,
		Replacement: s.Replacement,
		Type:        s.Type,
		Version:     s.Version,
	}

	tmp.Type = "pattern_replace"

	return json.Marshal(tmp)
}

// NewPatternReplaceTokenFilter returns a PatternReplaceTokenFilter.
func NewPatternReplaceTokenFilter() *PatternReplaceTokenFilter {
	r := &PatternReplaceTokenFilter{}

	return r
}
