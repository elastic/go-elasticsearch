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

// CommonGramsTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/token_filters.ts#L219-L235
type CommonGramsTokenFilter struct {
	// CommonWords A list of tokens. The filter generates bigrams for these tokens.
	// Either this or the `common_words_path` parameter is required.
	CommonWords []string `json:"common_words,omitempty"`
	// CommonWordsPath Path to a file containing a list of tokens. The filter generates bigrams for
	// these tokens.
	// This path must be absolute or relative to the `config` location. The file
	// must be UTF-8 encoded. Each token in the file must be separated by a line
	// break.
	// Either this or the `common_words` parameter is required.
	CommonWordsPath *string `json:"common_words_path,omitempty"`
	// IgnoreCase If `true`, matches for common words matching are case-insensitive. Defaults
	// to `false`.
	IgnoreCase *bool `json:"ignore_case,omitempty"`
	// QueryMode If `true`, the filter excludes the following tokens from the output:
	// - Unigrams for common words
	// - Unigrams for terms followed by common words
	// Defaults to `false`. We recommend enabling this parameter for search
	// analyzers.
	QueryMode *bool   `json:"query_mode,omitempty"`
	Type      string  `json:"type,omitempty"`
	Version   *string `json:"version,omitempty"`
}

func (s *CommonGramsTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "common_words":
			if err := dec.Decode(&s.CommonWords); err != nil {
				return fmt.Errorf("%s | %w", "CommonWords", err)
			}

		case "common_words_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "CommonWordsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CommonWordsPath = &o

		case "ignore_case":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreCase", err)
				}
				s.IgnoreCase = &value
			case bool:
				s.IgnoreCase = &v
			}

		case "query_mode":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "QueryMode", err)
				}
				s.QueryMode = &value
			case bool:
				s.QueryMode = &v
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
func (s CommonGramsTokenFilter) MarshalJSON() ([]byte, error) {
	type innerCommonGramsTokenFilter CommonGramsTokenFilter
	tmp := innerCommonGramsTokenFilter{
		CommonWords:     s.CommonWords,
		CommonWordsPath: s.CommonWordsPath,
		IgnoreCase:      s.IgnoreCase,
		QueryMode:       s.QueryMode,
		Type:            s.Type,
		Version:         s.Version,
	}

	tmp.Type = "common_grams"

	return json.Marshal(tmp)
}

// NewCommonGramsTokenFilter returns a CommonGramsTokenFilter.
func NewCommonGramsTokenFilter() *CommonGramsTokenFilter {
	r := &CommonGramsTokenFilter{}

	return r
}
