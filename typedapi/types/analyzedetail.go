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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// AnalyzeDetail type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/indices/analyze/types.ts#L24-L30
type AnalyzeDetail struct {
	Analyzer       *AnalyzerDetail    `json:"analyzer,omitempty"`
	Charfilters    []CharFilterDetail `json:"charfilters,omitempty"`
	CustomAnalyzer bool               `json:"custom_analyzer"`
	Tokenfilters   []TokenDetail      `json:"tokenfilters,omitempty"`
	Tokenizer      *TokenDetail       `json:"tokenizer,omitempty"`
}

func (s *AnalyzeDetail) UnmarshalJSON(data []byte) error {

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

		case "analyzer":
			if err := dec.Decode(&s.Analyzer); err != nil {
				return err
			}

		case "charfilters":
			if err := dec.Decode(&s.Charfilters); err != nil {
				return err
			}

		case "custom_analyzer":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.CustomAnalyzer = value
			case bool:
				s.CustomAnalyzer = v
			}

		case "tokenfilters":
			if err := dec.Decode(&s.Tokenfilters); err != nil {
				return err
			}

		case "tokenizer":
			if err := dec.Decode(&s.Tokenizer); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewAnalyzeDetail returns a AnalyzeDetail.
func NewAnalyzeDetail() *AnalyzeDetail {
	r := &AnalyzeDetail{}

	return r
}
