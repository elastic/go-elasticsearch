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
// https://github.com/elastic/elasticsearch-specification/tree/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// KeywordMarkerTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2f823ff6fcaa7f3f0f9b990dc90512d8901e5d64/specification/_types/analysis/token_filters.ts#L232-L238
type KeywordMarkerTokenFilter struct {
	IgnoreCase      *bool    `json:"ignore_case,omitempty"`
	Keywords        []string `json:"keywords,omitempty"`
	KeywordsPath    *string  `json:"keywords_path,omitempty"`
	KeywordsPattern *string  `json:"keywords_pattern,omitempty"`
	Type            string   `json:"type,omitempty"`
	Version         *string  `json:"version,omitempty"`
}

func (s *KeywordMarkerTokenFilter) UnmarshalJSON(data []byte) error {

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

		case "ignore_case":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "IgnoreCase", err)
				}
				s.IgnoreCase = &value
			case bool:
				s.IgnoreCase = &v
			}

		case "keywords":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := new(string)
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Keywords", err)
				}

				s.Keywords = append(s.Keywords, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Keywords); err != nil {
					return fmt.Errorf("%s | %w", "Keywords", err)
				}
			}

		case "keywords_path":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "KeywordsPath", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.KeywordsPath = &o

		case "keywords_pattern":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "KeywordsPattern", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.KeywordsPattern = &o

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
func (s KeywordMarkerTokenFilter) MarshalJSON() ([]byte, error) {
	type innerKeywordMarkerTokenFilter KeywordMarkerTokenFilter
	tmp := innerKeywordMarkerTokenFilter{
		IgnoreCase:      s.IgnoreCase,
		Keywords:        s.Keywords,
		KeywordsPath:    s.KeywordsPath,
		KeywordsPattern: s.KeywordsPattern,
		Type:            s.Type,
		Version:         s.Version,
	}

	tmp.Type = "keyword_marker"

	return json.Marshal(tmp)
}

// NewKeywordMarkerTokenFilter returns a KeywordMarkerTokenFilter.
func NewKeywordMarkerTokenFilter() *KeywordMarkerTokenFilter {
	r := &KeywordMarkerTokenFilter{}

	return r
}
