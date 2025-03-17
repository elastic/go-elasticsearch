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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ThaiAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ea991724f4dd4f90c496eff547d3cc2e6529f509/specification/_types/analysis/analyzers.ts#L317-L321
type ThaiAnalyzer struct {
	Stopwords     []string `json:"stopwords,omitempty"`
	StopwordsPath *string  `json:"stopwords_path,omitempty"`
	Type          string   `json:"type,omitempty"`
}

func (s *ThaiAnalyzer) UnmarshalJSON(data []byte) error {

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

		}
	}
	return nil
}

// MarshalJSON override marshalling to include literal value
func (s ThaiAnalyzer) MarshalJSON() ([]byte, error) {
	type innerThaiAnalyzer ThaiAnalyzer
	tmp := innerThaiAnalyzer{
		Stopwords:     s.Stopwords,
		StopwordsPath: s.StopwordsPath,
		Type:          s.Type,
	}

	tmp.Type = "thai"

	return json.Marshal(tmp)
}

// NewThaiAnalyzer returns a ThaiAnalyzer.
func NewThaiAnalyzer() *ThaiAnalyzer {
	r := &ThaiAnalyzer{}

	return r
}

// true

type ThaiAnalyzerVariant interface {
	ThaiAnalyzerCaster() *ThaiAnalyzer
}

func (s *ThaiAnalyzer) ThaiAnalyzerCaster() *ThaiAnalyzer {
	return s
}
