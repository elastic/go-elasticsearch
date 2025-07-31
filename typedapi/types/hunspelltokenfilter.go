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

// HunspellTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L268-L280
type HunspellTokenFilter struct {
	// Dedup If `true`, duplicate tokens are removed from the filter’s output. Defaults to
	// `true`.
	Dedup *bool `json:"dedup,omitempty"`
	// Dictionary One or more `.dic` files (e.g, `en_US.dic`, my_custom.dic) to use for the
	// Hunspell dictionary.
	// By default, the `hunspell` filter uses all `.dic` files in the
	// `<$ES_PATH_CONF>/hunspell/<locale>` directory specified using the `lang`,
	// `language`, or `locale` parameter.
	Dictionary *string `json:"dictionary,omitempty"`
	// Locale Locale directory used to specify the `.aff` and `.dic` files for a Hunspell
	// dictionary.
	Locale string `json:"locale"`
	// LongestOnly If `true`, only the longest stemmed version of each token is included in the
	// output. If `false`, all stemmed versions of the token are included. Defaults
	// to `false`.
	LongestOnly *bool   `json:"longest_only,omitempty"`
	Type        string  `json:"type,omitempty"`
	Version     *string `json:"version,omitempty"`
}

func (s *HunspellTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "dedup":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Dedup", err)
				}
				s.Dedup = &value
			case bool:
				s.Dedup = &v
			}

		case "dictionary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Dictionary", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Dictionary = &o

		case "locale", "lang", "language":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Locale", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Locale = o

		case "longest_only":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "LongestOnly", err)
				}
				s.LongestOnly = &value
			case bool:
				s.LongestOnly = &v
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
func (s HunspellTokenFilter) MarshalJSON() ([]byte, error) {
	type innerHunspellTokenFilter HunspellTokenFilter
	tmp := innerHunspellTokenFilter{
		Dedup:       s.Dedup,
		Dictionary:  s.Dictionary,
		Locale:      s.Locale,
		LongestOnly: s.LongestOnly,
		Type:        s.Type,
		Version:     s.Version,
	}

	tmp.Type = "hunspell"

	return json.Marshal(tmp)
}

// NewHunspellTokenFilter returns a HunspellTokenFilter.
func NewHunspellTokenFilter() *HunspellTokenFilter {
	r := &HunspellTokenFilter{}

	return r
}
