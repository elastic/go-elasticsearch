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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// WhitespaceTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/beeb1dc688bcc058488dcc45d9cbd2cd364e9943/specification/_types/analysis/tokenizers.ts#L135-L138
type WhitespaceTokenizer struct {
	MaxTokenLength *int    `json:"max_token_length,omitempty"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *WhitespaceTokenizer) UnmarshalJSON(data []byte) error {

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
func (s WhitespaceTokenizer) MarshalJSON() ([]byte, error) {
	type innerWhitespaceTokenizer WhitespaceTokenizer
	tmp := innerWhitespaceTokenizer{
		MaxTokenLength: s.MaxTokenLength,
		Type:           s.Type,
		Version:        s.Version,
	}

	tmp.Type = "whitespace"

	return json.Marshal(tmp)
}

// NewWhitespaceTokenizer returns a WhitespaceTokenizer.
func NewWhitespaceTokenizer() *WhitespaceTokenizer {
	r := &WhitespaceTokenizer{}

	return r
}

// true

type WhitespaceTokenizerVariant interface {
	WhitespaceTokenizerCaster() *WhitespaceTokenizer
}

func (s *WhitespaceTokenizer) WhitespaceTokenizerCaster() *WhitespaceTokenizer {
	return s
}
