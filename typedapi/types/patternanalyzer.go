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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

import (
	"bytes"
	"errors"
	"io"

	"strconv"

	"encoding/json"
)

// PatternAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/_types/analysis/analyzers.ts#L74-L81
type PatternAnalyzer struct {
	Flags     *string  `json:"flags,omitempty"`
	Lowercase *bool    `json:"lowercase,omitempty"`
	Pattern   string   `json:"pattern"`
	Stopwords []string `json:"stopwords,omitempty"`
	Type      string   `json:"type,omitempty"`
	Version   *string  `json:"version,omitempty"`
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
				return err
			}
			o := string(tmp)
			s.Flags = &o

		case "lowercase":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.Lowercase = &value
			case bool:
				s.Lowercase = &v
			}

		case "pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp)
			s.Pattern = o

		case "stopwords":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return err
				}

				s.Stopwords = append(s.Stopwords, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Stopwords); err != nil {
					return err
				}
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "version":
			if err := dec.Decode(&s.Version); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewPatternAnalyzer returns a PatternAnalyzer.
func NewPatternAnalyzer() *PatternAnalyzer {
	r := &PatternAnalyzer{}

	r.Type = "pattern"

	return r
}
