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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tokenchar"
)

// NGramTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/analysis/tokenizers.ts#L39-L45
type NGramTokenizer struct {
	CustomTokenChars *string               `json:"custom_token_chars,omitempty"`
	MaxGram          int                   `json:"max_gram"`
	MinGram          int                   `json:"min_gram"`
	TokenChars       []tokenchar.TokenChar `json:"token_chars"`
	Type             string                `json:"type,omitempty"`
	Version          *string               `json:"version,omitempty"`
}

func (s *NGramTokenizer) UnmarshalJSON(data []byte) error {

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

		case "custom_token_chars":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.CustomTokenChars = &o

		case "max_gram":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxGram = value
			case float64:
				f := int(v)
				s.MaxGram = f
			}

		case "min_gram":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinGram = value
			case float64:
				f := int(v)
				s.MinGram = f
			}

		case "token_chars":
			if err := dec.Decode(&s.TokenChars); err != nil {
				return err
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

// MarshalJSON override marshalling to include literal value
func (s NGramTokenizer) MarshalJSON() ([]byte, error) {
	type innerNGramTokenizer NGramTokenizer
	tmp := innerNGramTokenizer{
		CustomTokenChars: s.CustomTokenChars,
		MaxGram:          s.MaxGram,
		MinGram:          s.MinGram,
		TokenChars:       s.TokenChars,
		Type:             s.Type,
		Version:          s.Version,
	}

	tmp.Type = "ngram"

	return json.Marshal(tmp)
}

// NewNGramTokenizer returns a NGramTokenizer.
func NewNGramTokenizer() *NGramTokenizer {
	r := &NGramTokenizer{}

	return r
}
