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

// StemmerOverrideTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L417-L423
type StemmerOverrideTokenFilter struct {
	// Rules A list of mapping rules to use.
	Rules []string `json:"rules,omitempty"`
	// RulesPath A path (either relative to `config` location, or absolute) to a list of
	// mappings.
	RulesPath *string `json:"rules_path,omitempty"`
	Type      string  `json:"type,omitempty"`
	Version   *string `json:"version,omitempty"`
}

func (s *StemmerOverrideTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "rules":
			if err := dec.Decode(&s.Rules); err != nil {
				return fmt.Errorf("%s | %w", "Rules", err)
			}

		case "rules_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RulesPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RulesPath = &o

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
func (s StemmerOverrideTokenFilter) MarshalJSON() ([]byte, error) {
	type innerStemmerOverrideTokenFilter StemmerOverrideTokenFilter
	tmp := innerStemmerOverrideTokenFilter{
		Rules:     s.Rules,
		RulesPath: s.RulesPath,
		Type:      s.Type,
		Version:   s.Version,
	}

	tmp.Type = "stemmer_override"

	return json.Marshal(tmp)
}

// NewStemmerOverrideTokenFilter returns a StemmerOverrideTokenFilter.
func NewStemmerOverrideTokenFilter() *StemmerOverrideTokenFilter {
	r := &StemmerOverrideTokenFilter{}

	return r
}
