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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noridecompoundmode"
)

// NoriTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/analysis/nori-plugin.ts#L29-L35
type NoriTokenizer struct {
	DecompoundMode      *noridecompoundmode.NoriDecompoundMode `json:"decompound_mode,omitempty"`
	DiscardPunctuation  *bool                                  `json:"discard_punctuation,omitempty"`
	Type                string                                 `json:"type,omitempty"`
	UserDictionary      *string                                `json:"user_dictionary,omitempty"`
	UserDictionaryRules []string                               `json:"user_dictionary_rules,omitempty"`
	Version             *string                                `json:"version,omitempty"`
}

func (s *NoriTokenizer) UnmarshalJSON(data []byte) error {

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

		case "decompound_mode":
			if err := dec.Decode(&s.DecompoundMode); err != nil {
				return fmt.Errorf("%s | %w", "DecompoundMode", err)
			}

		case "discard_punctuation":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "DiscardPunctuation", err)
				}
				s.DiscardPunctuation = &value
			case bool:
				s.DiscardPunctuation = &v
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return fmt.Errorf("%s | %w", "Type", err)
			}

		case "user_dictionary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "UserDictionary", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UserDictionary = &o

		case "user_dictionary_rules":
			if err := dec.Decode(&s.UserDictionaryRules); err != nil {
				return fmt.Errorf("%s | %w", "UserDictionaryRules", err)
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
func (s NoriTokenizer) MarshalJSON() ([]byte, error) {
	type innerNoriTokenizer NoriTokenizer
	tmp := innerNoriTokenizer{
		DecompoundMode:      s.DecompoundMode,
		DiscardPunctuation:  s.DiscardPunctuation,
		Type:                s.Type,
		UserDictionary:      s.UserDictionary,
		UserDictionaryRules: s.UserDictionaryRules,
		Version:             s.Version,
	}

	tmp.Type = "nori_tokenizer"

	return json.Marshal(tmp)
}

// NewNoriTokenizer returns a NoriTokenizer.
func NewNoriTokenizer() *NoriTokenizer {
	r := &NoriTokenizer{}

	return r
}
