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
// https://github.com/elastic/elasticsearch-specification/tree/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
)

// KuromojiPartOfSpeechTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/26d0e2015b6bb2b1e0c549a4f1abeca6da16e89c/specification/_types/analysis/kuromoji-plugin.ts#L37-L40
type KuromojiPartOfSpeechTokenFilter struct {
	Stoptags []string `json:"stoptags"`
	Type     string   `json:"type,omitempty"`
	Version  *string  `json:"version,omitempty"`
}

func (s *KuromojiPartOfSpeechTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "stoptags":
			if err := dec.Decode(&s.Stoptags); err != nil {
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
func (s KuromojiPartOfSpeechTokenFilter) MarshalJSON() ([]byte, error) {
	type innerKuromojiPartOfSpeechTokenFilter KuromojiPartOfSpeechTokenFilter
	tmp := innerKuromojiPartOfSpeechTokenFilter{
		Stoptags: s.Stoptags,
		Type:     s.Type,
		Version:  s.Version,
	}

	tmp.Type = "kuromoji_part_of_speech"

	return json.Marshal(tmp)
}

// NewKuromojiPartOfSpeechTokenFilter returns a KuromojiPartOfSpeechTokenFilter.
func NewKuromojiPartOfSpeechTokenFilter() *KuromojiPartOfSpeechTokenFilter {
	r := &KuromojiPartOfSpeechTokenFilter{}

	return r
}
