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

// WordDelimiterTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L201-L203
type WordDelimiterTokenFilter struct {
	// CatenateAll If `true`, the filter produces catenated tokens for chains of alphanumeric
	// characters separated by non-alphabetic delimiters. Defaults to `false`.
	CatenateAll *bool `json:"catenate_all,omitempty"`
	// CatenateNumbers If `true`, the filter produces catenated tokens for chains of numeric
	// characters separated by non-alphabetic delimiters. Defaults to `false`.
	CatenateNumbers *bool `json:"catenate_numbers,omitempty"`
	// CatenateWords If `true`, the filter produces catenated tokens for chains of alphabetical
	// characters separated by non-alphabetic delimiters. Defaults to `false`.
	CatenateWords *bool `json:"catenate_words,omitempty"`
	// GenerateNumberParts If `true`, the filter includes tokens consisting of only numeric characters
	// in the output. If `false`, the filter excludes these tokens from the output.
	// Defaults to `true`.
	GenerateNumberParts *bool `json:"generate_number_parts,omitempty"`
	// GenerateWordParts If `true`, the filter includes tokens consisting of only alphabetical
	// characters in the output. If `false`, the filter excludes these tokens from
	// the output. Defaults to `true`.
	GenerateWordParts *bool `json:"generate_word_parts,omitempty"`
	// PreserveOriginal If `true`, the filter includes the original version of any split tokens in
	// the output. This original version includes non-alphanumeric delimiters.
	// Defaults to `false`.
	PreserveOriginal Stringifiedboolean `json:"preserve_original,omitempty"`
	// ProtectedWords Array of tokens the filter won’t split.
	ProtectedWords []string `json:"protected_words,omitempty"`
	// ProtectedWordsPath Path to a file that contains a list of tokens the filter won’t split.
	// This path must be absolute or relative to the `config` location, and the file
	// must be UTF-8 encoded. Each token in the file must be separated by a line
	// break.
	ProtectedWordsPath *string `json:"protected_words_path,omitempty"`
	// SplitOnCaseChange If `true`, the filter splits tokens at letter case transitions. For example:
	// camelCase -> [ camel, Case ]. Defaults to `true`.
	SplitOnCaseChange *bool `json:"split_on_case_change,omitempty"`
	// SplitOnNumerics If `true`, the filter splits tokens at letter-number transitions. For
	// example: j2se -> [ j, 2, se ]. Defaults to `true`.
	SplitOnNumerics *bool `json:"split_on_numerics,omitempty"`
	// StemEnglishPossessive If `true`, the filter removes the English possessive (`'s`) from the end of
	// each token. For example: O'Neil's -> [ O, Neil ]. Defaults to `true`.
	StemEnglishPossessive *bool  `json:"stem_english_possessive,omitempty"`
	Type                  string `json:"type,omitempty"`
	// TypeTable Array of custom type mappings for characters. This allows you to map
	// non-alphanumeric characters as numeric or alphanumeric to avoid splitting on
	// those characters.
	TypeTable []string `json:"type_table,omitempty"`
	// TypeTablePath Path to a file that contains custom type mappings for characters. This allows
	// you to map non-alphanumeric characters as numeric or alphanumeric to avoid
	// splitting on those characters.
	TypeTablePath *string `json:"type_table_path,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *WordDelimiterTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "catenate_all":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CatenateAll", err)
				}
				s.CatenateAll = &value
			case bool:
				s.CatenateAll = &v
			}

		case "catenate_numbers":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CatenateNumbers", err)
				}
				s.CatenateNumbers = &value
			case bool:
				s.CatenateNumbers = &v
			}

		case "catenate_words":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CatenateWords", err)
				}
				s.CatenateWords = &value
			case bool:
				s.CatenateWords = &v
			}

		case "generate_number_parts":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "GenerateNumberParts", err)
				}
				s.GenerateNumberParts = &value
			case bool:
				s.GenerateNumberParts = &v
			}

		case "generate_word_parts":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "GenerateWordParts", err)
				}
				s.GenerateWordParts = &value
			case bool:
				s.GenerateWordParts = &v
			}

		case "preserve_original":
			if err := dec.Decode(&s.PreserveOriginal); err != nil {
				return fmt.Errorf("%s | %w", "PreserveOriginal", err)
			}

		case "protected_words":
			if err := dec.Decode(&s.ProtectedWords); err != nil {
				return fmt.Errorf("%s | %w", "ProtectedWords", err)
			}

		case "protected_words_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ProtectedWordsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ProtectedWordsPath = &o

		case "split_on_case_change":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SplitOnCaseChange", err)
				}
				s.SplitOnCaseChange = &value
			case bool:
				s.SplitOnCaseChange = &v
			}

		case "split_on_numerics":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "SplitOnNumerics", err)
				}
				s.SplitOnNumerics = &value
			case bool:
				s.SplitOnNumerics = &v
			}

		case "stem_english_possessive":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "StemEnglishPossessive", err)
				}
				s.StemEnglishPossessive = &value
			case bool:
				s.StemEnglishPossessive = &v
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "type_table":
			if err := dec.Decode(&s.TypeTable); err != nil {
				return fmt.Errorf("%s | %w", "TypeTable", err)
			}

		case "type_table_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TypeTablePath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TypeTablePath = &o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return fmt.Errorf("%s | %w", "Version", err)
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s WordDelimiterTokenFilter) MarshalJSON() ([]byte, error) {
	type innerWordDelimiterTokenFilter WordDelimiterTokenFilter
	tmp := innerWordDelimiterTokenFilter{
		CatenateAll:           s.CatenateAll,
		CatenateNumbers:       s.CatenateNumbers,
		CatenateWords:         s.CatenateWords,
		GenerateNumberParts:   s.GenerateNumberParts,
		GenerateWordParts:     s.GenerateWordParts,
		PreserveOriginal:      s.PreserveOriginal,
		ProtectedWords:        s.ProtectedWords,
		ProtectedWordsPath:    s.ProtectedWordsPath,
		SplitOnCaseChange:     s.SplitOnCaseChange,
		SplitOnNumerics:       s.SplitOnNumerics,
		StemEnglishPossessive: s.StemEnglishPossessive,
		Type:                  s.Type,
		TypeTable:             s.TypeTable,
		TypeTablePath:         s.TypeTablePath,
		Version:               s.Version,
	}

	tmp.Type = "word_delimiter"

	return json.Marshal(tmp)
}

// NewWordDelimiterTokenFilter returns a WordDelimiterTokenFilter.
func NewWordDelimiterTokenFilter() *WordDelimiterTokenFilter {
	r := &WordDelimiterTokenFilter{}

	return r
}
