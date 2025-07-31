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

// StandardAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/analyzers.ts#L382-L402
type StandardAnalyzer struct {
	// MaxTokenLength The maximum token length. If a token is seen that exceeds this length then it
	// is split at `max_token_length` intervals.
	// Defaults to `255`.
	MaxTokenLength *int `json:"max_token_length,omitempty"`
	// Stopwords A pre-defined stop words list like `_english_` or an array containing a list
	// of stop words.
	// Defaults to `_none_`.
	Stopwords StopWords `json:"stopwords,omitempty"`
	// StopwordsPath The path to a file containing stop words.
	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
}

func (s *StandardAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "max_token_length":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxTokenLength", err)
				}
				s.MaxTokenLength = &value
			case float64:
				f := int(v)
				s.MaxTokenLength = &f
			}

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

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s StandardAnalyzer) MarshalJSON() ([]byte, error) {
	type innerStandardAnalyzer StandardAnalyzer
	tmp := innerStandardAnalyzer{
		MaxTokenLength: s.MaxTokenLength,
		Stopwords:      s.Stopwords,
		StopwordsPath:  s.StopwordsPath,
		Type:           s.Type,
	}

	tmp.Type = "standard"

	return json.Marshal(tmp)
}

// NewStandardAnalyzer returns a StandardAnalyzer.
func NewStandardAnalyzer() *StandardAnalyzer {
	r := &StandardAnalyzer{}

	return r
}

type StandardAnalyzerVariant interface {
	StandardAnalyzerCaster() *StandardAnalyzer
}

func (s *StandardAnalyzer) StandardAnalyzerCaster() *StandardAnalyzer {
	return s
}

func (s *StandardAnalyzer) AnalyzerCaster() *Analyzer {
	o := Analyzer(s)
	return &o
}
