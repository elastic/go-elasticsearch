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

// PatternAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/analyzers.ts#L332-L365
type PatternAnalyzer struct {
	// Flags Java regular expression flags. Flags should be pipe-separated, eg
	// "CASE_INSENSITIVE|COMMENTS".
	Flags *string `json:"flags,omitempty"`
	// Lowercase Should terms be lowercased or not.
	// Defaults to `true`.
	Lowercase *bool `json:"lowercase,omitempty"`
	// Pattern A Java regular expression.
	// Defaults to `\W+`.
	Pattern *string `json:"pattern,omitempty"`
	// Stopwords A pre-defined stop words list like `_english_` or an array containing a list
	// of stop words.
	// Defaults to `_none_`.
	Stopwords StopWords `json:"stopwords,omitempty"`
	// StopwordsPath The path to a file containing stop words.
	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *PatternAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "lowercase":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Lowercase", err)
				}
				s.Lowercase = &value
			case bool:
				s.Lowercase = &v
			}

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
			s.Pattern = &o

		case "stopwords":
			if err := dec.Decode(&s.Stopwords); err != nil {
				return fmt.Errorf("%s | %w", "Stopwords", err)
			}

		case "stopwords_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "StopwordsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StopwordsPath = &o

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
func (s PatternAnalyzer) MarshalJSON() ([]byte, error) {
	type innerPatternAnalyzer PatternAnalyzer
	tmp := innerPatternAnalyzer{
		Flags:         s.Flags,
		Lowercase:     s.Lowercase,
		Pattern:       s.Pattern,
		Stopwords:     s.Stopwords,
		StopwordsPath: s.StopwordsPath,
		Type:          s.Type,
		Version:       s.Version,
	}

	tmp.Type = "pattern"

	return json.Marshal(tmp)
}

// NewPatternAnalyzer returns a PatternAnalyzer.
func NewPatternAnalyzer() *PatternAnalyzer {
	r := &PatternAnalyzer{}

	return r
}

type PatternAnalyzerVariant interface {
	PatternAnalyzerCaster() *PatternAnalyzer
}

func (s *PatternAnalyzer) PatternAnalyzerCaster() *PatternAnalyzer {
	return s
}

func (s *PatternAnalyzer) AnalyzerCaster() *Analyzer {
	o := Analyzer(s)
	return &o
}
