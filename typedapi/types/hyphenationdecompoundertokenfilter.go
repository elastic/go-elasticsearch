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
// https://github.com/elastic/elasticsearch-specification/tree/907d11a72a6bfd37b777d526880c56202889609e

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// HyphenationDecompounderTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/token_filters.ts#L67-L76
type HyphenationDecompounderTokenFilter struct {
	// HyphenationPatternsPath Path to an Apache FOP (Formatting Objects Processor) XML hyphenation pattern
	// file.
	// This path must be absolute or relative to the `config` location. Only FOP
	// v1.2 compatible files are supported.
	HyphenationPatternsPath string `json:"hyphenation_patterns_path"`
	// MaxSubwordSize Maximum subword character length. Longer subword tokens are excluded from the
	// output. Defaults to `15`.
	MaxSubwordSize *int `json:"max_subword_size,omitempty"`
	// MinSubwordSize Minimum subword character length. Shorter subword tokens are excluded from
	// the output. Defaults to `2`.
	MinSubwordSize *int `json:"min_subword_size,omitempty"`
	// MinWordSize Minimum word character length. Shorter word tokens are excluded from the
	// output. Defaults to `5`.
	MinWordSize *int `json:"min_word_size,omitempty"`
	// NoOverlappingMatches If `true`, do not allow overlapping tokens. Defaults to `false`.
	NoOverlappingMatches *bool `json:"no_overlapping_matches,omitempty"`
	// NoSubMatches If `true`, do not match sub tokens in tokens that are in the word list.
	// Defaults to `false`.
	NoSubMatches *bool `json:"no_sub_matches,omitempty"`
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

func (s *HyphenationDecompounderTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "hyphenation_patterns_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "HyphenationPatternsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.HyphenationPatternsPath = o

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

		case "no_overlapping_matches":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NoOverlappingMatches", err)
				}
				s.NoOverlappingMatches = &value
			case bool:
				s.NoOverlappingMatches = &v
			}

		case "no_sub_matches":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "NoSubMatches", err)
				}
				s.NoSubMatches = &value
			case bool:
				s.NoSubMatches = &v
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
func (s HyphenationDecompounderTokenFilter) MarshalJSON() ([]byte, error) {
	type innerHyphenationDecompounderTokenFilter HyphenationDecompounderTokenFilter
	tmp := innerHyphenationDecompounderTokenFilter{
		HyphenationPatternsPath: s.HyphenationPatternsPath,
		MaxSubwordSize:          s.MaxSubwordSize,
		MinSubwordSize:          s.MinSubwordSize,
		MinWordSize:             s.MinWordSize,
		NoOverlappingMatches:    s.NoOverlappingMatches,
		NoSubMatches:            s.NoSubMatches,
		OnlyLongestMatch:        s.OnlyLongestMatch,
		Type:                    s.Type,
		Version:                 s.Version,
		WordList:                s.WordList,
		WordListPath:            s.WordListPath,
	}

	tmp.Type = "hyphenation_decompounder"

	return json.Marshal(tmp)
}

// NewHyphenationDecompounderTokenFilter returns a HyphenationDecompounderTokenFilter.
func NewHyphenationDecompounderTokenFilter() *HyphenationDecompounderTokenFilter {
	r := &HyphenationDecompounderTokenFilter{}

	return r
}

type HyphenationDecompounderTokenFilterVariant interface {
	HyphenationDecompounderTokenFilterCaster() *HyphenationDecompounderTokenFilter
}

func (s *HyphenationDecompounderTokenFilter) HyphenationDecompounderTokenFilterCaster() *HyphenationDecompounderTokenFilter {
	return s
}

func (s *HyphenationDecompounderTokenFilter) TokenFilterDefinitionCaster() *TokenFilterDefinition {
	o := TokenFilterDefinition(s)
	return &o
}
