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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/noridecompoundmode"
)

// NoriTokenizer type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/_types/analysis/tokenizers.ts#L81-L87
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
				return err
			}

		case "discard_punctuation":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.DiscardPunctuation = &value
			case bool:
				s.DiscardPunctuation = &v
			}

		case "type":
			if err := dec.Decode(&s.Type); err != nil {
				return err
			}

		case "user_dictionary":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.UserDictionary = &o

		case "user_dictionary_rules":
			if err := dec.Decode(&s.UserDictionaryRules); err != nil {
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
