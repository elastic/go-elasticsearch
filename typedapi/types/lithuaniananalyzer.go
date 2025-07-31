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

// LithuanianAnalyzer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/analyzers.ts#L221-L226
type LithuanianAnalyzer struct {
	StemExclusion []string  `json:"stem_exclusion,omitempty"`
	Stopwords     StopWords `json:"stopwords,omitempty"`
	StopwordsPath *string   `json:"stopwords_path,omitempty"`
	Type          string    `json:"type,omitempty"`
}

func (s *LithuanianAnalyzer) UnmarshalJSON(data []byte) error {

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

		case "stem_exclusion":
			if err := dec.Decode(&s.StemExclusion); err != nil {
				return fmt.Errorf("%s | %w", "StemExclusion", err)
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
func (s LithuanianAnalyzer) MarshalJSON() ([]byte, error) {
	type innerLithuanianAnalyzer LithuanianAnalyzer
	tmp := innerLithuanianAnalyzer{
		StemExclusion: s.StemExclusion,
		Stopwords:     s.Stopwords,
		StopwordsPath: s.StopwordsPath,
		Type:          s.Type,
	}

	tmp.Type = "lithuanian"

	return json.Marshal(tmp)
}

// NewLithuanianAnalyzer returns a LithuanianAnalyzer.
func NewLithuanianAnalyzer() *LithuanianAnalyzer {
	r := &LithuanianAnalyzer{}

	return r
}
