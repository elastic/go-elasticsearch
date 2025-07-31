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

// DictionaryDecompounderTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L63-L65
type DictionaryDecompounderTokenFilter struct {
	// MaxSubwordSize Maximum subword character length. Longer subword tokens are excluded from the
	// output. Defaults to `15`.
	MaxSubwordSize *int `json:"max_subword_size,omitempty"`
	// MinSubwordSize Minimum subword character length. Shorter subword tokens are excluded from
	// the output. Defaults to `2`.
	MinSubwordSize *int `json:"min_subword_size,omitempty"`
	// MinWordSize Minimum word character length. Shorter word tokens are excluded from the
	// output. Defaults to `5`.
	MinWordSize *int `json:"min_word_size,omitempty"`
	// OnlyLongestMatch If `true`, only include the longest matching subword. Defaults to `false`.
	OnlyLongestMatch *bool   `json:"only_longest_match,omitempty"`
	Type             string  `json:"type,omitempty"`
	Version          *string `json:"version,omitempty"`
	// WordList A list of subwords to look for in the token stream. If found, the subword is
	// included in the token output.
	// Either this parameter or `word_list_path` must be specified.
	WordList []string `json:"word_list,omitempty"`
	// WordListPath Path to a file that contains a list of subwords to find in the token stream.
	// If found, the subword is included in the token output.
	// This path must be absolute or relative to the config location, and the file
	// must be UTF-8 encoded. Each token in the file must be separated by a line
	// break.
	// Either this parameter or `word_list` must be specified.
	WordListPath *string `json:"word_list_path,omitempty"`
}

func (s *DictionaryDecompounderTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "max_subword_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxSubwordSize", err)
				}
				s.MaxSubwordSize = &value
			case float64:
				f := int(v)
				s.MaxSubwordSize = &f
			}

		case "min_subword_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinSubwordSize", err)
				}
				s.MinSubwordSize = &value
			case float64:
				f := int(v)
				s.MinSubwordSize = &f
			}

		case "min_word_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MinWordSize", err)
				}
				s.MinWordSize = &value
			case float64:
				f := int(v)
				s.MinWordSize = &f
			}

		case "only_longest_match":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OnlyLongestMatch", err)
				}
				s.OnlyLongestMatch = &value
			case bool:
				s.OnlyLongestMatch = &v
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		case "word_list":
			if err := dec.Decode(&s.WordList); err != nil {
				return fmt.Errorf("%s | %w", "WordList", err)
			}

		case "word_list_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "WordListPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.WordListPath = &o

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s DictionaryDecompounderTokenFilter) MarshalJSON() ([]byte, error) {
	type innerDictionaryDecompounderTokenFilter DictionaryDecompounderTokenFilter
	tmp := innerDictionaryDecompounderTokenFilter{
		MaxSubwordSize:   s.MaxSubwordSize,
		MinSubwordSize:   s.MinSubwordSize,
		MinWordSize:      s.MinWordSize,
		OnlyLongestMatch: s.OnlyLongestMatch,
		Type:             s.Type,
		Version:          s.Version,
		WordList:         s.WordList,
		WordListPath:     s.WordListPath,
	}

	tmp.Type = "dictionary_decompounder"

	return json.Marshal(tmp)
}

// NewDictionaryDecompounderTokenFilter returns a DictionaryDecompounderTokenFilter.
func NewDictionaryDecompounderTokenFilter() *DictionaryDecompounderTokenFilter {
	r := &DictionaryDecompounderTokenFilter{}

	return r
}
