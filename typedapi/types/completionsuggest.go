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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// CompletionSuggest type.
//
// https://github.com/elastic/elasticsearch-specification/blob/60a81659be928bfe6cec53708c7f7613555a5eaf/specification/_global/search/_types/suggester.ts#L49-L56
type CompletionSuggest struct {
	Length  int                       `json:"length"`
	Offset  int                       `json:"offset"`
	Options []CompletionSuggestOption `json:"options"`
	Text    string                    `json:"text"`
}

func (s *CompletionSuggest) UnmarshalJSON(data []byte) error {

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

		case "length":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Length", err)
				}
				s.Length = value
			case float64:
				f := int(v)
				s.Length = f
			}

		case "offset":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Offset", err)
				}
				s.Offset = value
			case float64:
				f := int(v)
				s.Offset = f
			}

		case "options":
			rawMsg := json.RawMessage{}
			dec.Decode(&rawMsg)
			if !bytes.HasPrefix(rawMsg, []byte("[")) {
				o := NewCompletionSuggestOption()
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&o); err != nil {
					return fmt.Errorf("%s | %w", "Options", err)
				}

				s.Options = append(s.Options, *o)
			} else {
				if err := json.NewDecoder(bytes.NewReader(rawMsg)).Decode(&s.Options); err != nil {
					return fmt.Errorf("%s | %w", "Options", err)
				}
			}

		case "text":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Text", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Text = o

		}
	}
	return nil
}

// NewCompletionSuggest returns a CompletionSuggest.
func NewCompletionSuggest() *CompletionSuggest {
	r := &CompletionSuggest{}

	return r
}

// false
