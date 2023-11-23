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

// KuromojiIterationMarkCharFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_types/analysis/kuromoji-plugin.ts#L31-L35
type KuromojiIterationMarkCharFilter struct {
	NormalizeKana  bool    `json:"normalize_kana"`
	NormalizeKanji bool    `json:"normalize_kanji"`
	Type           string  `json:"type,omitempty"`
	Version        *string `json:"version,omitempty"`
}

func (s *KuromojiIterationMarkCharFilter) UnmarshalJSON(data []byte) error {

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

		case "normalize_kana":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.NormalizeKana = value
			case bool:
				s.NormalizeKana = v
			}

		case "normalize_kanji":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.NormalizeKanji = value
			case bool:
				s.NormalizeKanji = v
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
func (s KuromojiIterationMarkCharFilter) MarshalJSON() ([]byte, error) {
	type innerKuromojiIterationMarkCharFilter KuromojiIterationMarkCharFilter
	tmp := innerKuromojiIterationMarkCharFilter{
		NormalizeKana:  s.NormalizeKana,
		NormalizeKanji: s.NormalizeKanji,
		Type:           s.Type,
		Version:        s.Version,
	}

	tmp.Type = "kuromoji_iteration_mark"

	return json.Marshal(tmp)
}

// NewKuromojiIterationMarkCharFilter returns a KuromojiIterationMarkCharFilter.
func NewKuromojiIterationMarkCharFilter() *KuromojiIterationMarkCharFilter {
	r := &KuromojiIterationMarkCharFilter{}

	return r
}
