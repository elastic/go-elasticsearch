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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// StopAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/60a81659be928bfe6cec53708c7f7613555a5eaf/specification/_types/analysis/analyzers.ts#L404-L419
type StopAnalyzer struct {
	// Stopwords A pre-defined stop words list like `_english_` or an array containing a list
	// of stop words.
	// Defaults to `_none_`.
	Stopwords []string `json:"stopwords,omitempty"`
	// StopwordsPath The path to a file containing stop words.
	StopwordsPath *string `json:"stopwords_path,omitempty"`
	Type          string  `json:"type,omitempty"`
	Version       *string `json:"version,omitempty"`
}

func (s *StopAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "stopwords":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Stopwords", err)
				}

				s.Stopwords = append(s.Stopwords, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Stopwords); err != nil {
					return fmt.Errorf("%s | %w", "Stopwords", err)
				}
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
func (s StopAnalyzer) MarshalJSON() ([]byte, error) {
	type innerStopAnalyzer StopAnalyzer
	tmp := innerStopAnalyzer{
		Stopwords:     s.Stopwords,
		StopwordsPath: s.StopwordsPath,
		Type:          s.Type,
		Version:       s.Version,
	}

	tmp.Type = "stop"

	return json.Marshal(tmp)
}

// NewStopAnalyzer returns a StopAnalyzer.
func NewStopAnalyzer() *StopAnalyzer {
	r := &StopAnalyzer{}

	return r
}

// true

type StopAnalyzerVariant interface {
	StopAnalyzerCaster() *StopAnalyzer
}

func (s *StopAnalyzer) StopAnalyzerCaster() *StopAnalyzer {
	return s
}
