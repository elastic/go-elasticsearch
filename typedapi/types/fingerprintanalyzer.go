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

// FingerprintAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/907d11a72a6bfd37b777d526880c56202889609e/specification/_types/analysis/analyzers.ts#L37-L64
type FingerprintAnalyzer struct {
	// MaxOutputSize The maximum token size to emit. Tokens larger than this size will be
	// discarded.
	// Defaults to `255`
	MaxOutputSize *int `json:"max_output_size,omitempty"`
	// Separator The character to use to concatenate the terms.
	// Defaults to a space.
	Separator *string `json:"separator,omitempty"`
	// Stopwords A pre-defined stop words list like `_english_` or an array containing a list
	// of stop words.
	// Defaults to `_none_`.
	Stopwords StopWords `json:"stopwords,omitempty"`
	// StopwordsPath The path to a file containing stop words.
	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *FingerprintAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "max_output_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxOutputSize", err)
				}
				s.MaxOutputSize = &value
			case float64:
				f := int(v)
				s.MaxOutputSize = &f
			}

		case "separator":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Separator", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Separator = &o

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
func (s FingerprintAnalyzer) MarshalJSON() ([]byte, error) {
	type innerFingerprintAnalyzer FingerprintAnalyzer
	tmp := innerFingerprintAnalyzer{
		MaxOutputSize: s.MaxOutputSize,
		Separator:     s.Separator,
		Stopwords:     s.Stopwords,
		StopwordsPath: s.StopwordsPath,
		Type:          s.Type,
		Version:       s.Version,
	}

	tmp.Type = "fingerprint"

	return json.Marshal(tmp)
}

// NewFingerprintAnalyzer returns a FingerprintAnalyzer.
func NewFingerprintAnalyzer() *FingerprintAnalyzer {
	r := &FingerprintAnalyzer{}

	return r
}

type FingerprintAnalyzerVariant interface {
	FingerprintAnalyzerCaster() *FingerprintAnalyzer
}

func (s *FingerprintAnalyzer) FingerprintAnalyzerCaster() *FingerprintAnalyzer {
	return s
}

func (s *FingerprintAnalyzer) AnalyzerCaster() *Analyzer {
	o := Analyzer(s)
	return &o
}
