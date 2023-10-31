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
	"strconv"
)

// WordDelimiterGraphTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/analysis/token_filters.ts#L149-L166
type WordDelimiterGraphTokenFilter struct {
	AdjustOffsets         *bool              `json:"adjust_offsets,omitempty"`
	CatenateAll           *bool              `json:"catenate_all,omitempty"`
	CatenateNumbers       *bool              `json:"catenate_numbers,omitempty"`
	CatenateWords         *bool              `json:"catenate_words,omitempty"`
	GenerateNumberParts   *bool              `json:"generate_number_parts,omitempty"`
	GenerateWordParts     *bool              `json:"generate_word_parts,omitempty"`
	IgnoreKeywords        *bool              `json:"ignore_keywords,omitempty"`
	PreserveOriginal      Stringifiedboolean `json:"preserve_original,omitempty"`
	ProtectedWords        []string           `json:"protected_words,omitempty"`
	ProtectedWordsPath    *string            `json:"protected_words_path,omitempty"`
	SplitOnCaseChange     *bool              `json:"split_on_case_change,omitempty"`
	SplitOnNumerics       *bool              `json:"split_on_numerics,omitempty"`
	StemEnglishPossessive *bool              `json:"stem_english_possessive,omitempty"`
	Type                  string             `json:"type,omitempty"`
	TypeTable             []string           `json:"type_table,omitempty"`
	TypeTablePath         *string            `json:"type_table_path,omitempty"`
	Version               *string            `json:"version,omitempty"`
}

func (s *WordDelimiterGraphTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "adjust_offsets":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AdjustOffsets = &value
			case bool:
				s.AdjustOffsets = &v
			}

		case "catenate_all":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.CatenateAll = &value
			case bool:
				s.CatenateAll = &v
			}

		case "catenate_numbers":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.CatenateNumbers = &value
			case bool:
				s.CatenateNumbers = &v
			}

		case "catenate_words":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.CatenateWords = &value
			case bool:
				s.CatenateWords = &v
			}

		case "generate_number_parts":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.GenerateNumberParts = &value
			case bool:
				s.GenerateNumberParts = &v
			}

		case "generate_word_parts":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.GenerateWordParts = &value
			case bool:
				s.GenerateWordParts = &v
			}

		case "ignore_keywords":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.IgnoreKeywords = &value
			case bool:
				s.IgnoreKeywords = &v
			}

		case "preserve_original":
			if err := dec.Decode(&s.PreserveOriginal); err != nil {
				return err
			}

		case "protected_words":
			if err := dec.Decode(&s.ProtectedWords); err != nil {
				return err
			}

		case "protected_words_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ProtectedWordsPath = &o

		case "split_on_case_change":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.SplitOnCaseChange = &value
			case bool:
				s.SplitOnCaseChange = &v
			}

		case "split_on_numerics":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.SplitOnNumerics = &value
			case bool:
				s.SplitOnNumerics = &v
			}

		case "stem_english_possessive":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.StemEnglishPossessive = &value
			case bool:
				s.StemEnglishPossessive = &v
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "type_table":
			if err := dec.Decode(&s.TypeTable); err != nil {
				return err
			}

		case "type_table_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TypeTablePath = &o

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s WordDelimiterGraphTokenFilter) MarshalJSON() ([]byte, error) {
	type innerWordDelimiterGraphTokenFilter WordDelimiterGraphTokenFilter
	tmp := innerWordDelimiterGraphTokenFilter{
		AdjustOffsets:         s.AdjustOffsets,
		CatenateAll:           s.CatenateAll,
		CatenateNumbers:       s.CatenateNumbers,
		CatenateWords:         s.CatenateWords,
		GenerateNumberParts:   s.GenerateNumberParts,
		GenerateWordParts:     s.GenerateWordParts,
		IgnoreKeywords:        s.IgnoreKeywords,
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

	tmp.Type = "word_delimiter_graph"

	return json.Marshal(tmp)
}

// NewWordDelimiterGraphTokenFilter returns a WordDelimiterGraphTokenFilter.
func NewWordDelimiterGraphTokenFilter() *WordDelimiterGraphTokenFilter {
	r := &WordDelimiterGraphTokenFilter{}

	return r
}
