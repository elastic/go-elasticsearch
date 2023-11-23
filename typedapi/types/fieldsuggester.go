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

// FieldSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/search/_types/suggester.ts#L106-L139
type FieldSuggester struct {
	// Completion Provides auto-complete/search-as-you-type functionality.
	Completion *CompletionSuggester `json:"completion,omitempty"`
	// Phrase Provides access to word alternatives on a per token basis within a certain
	// string distance.
	Phrase *PhraseSuggester `json:"phrase,omitempty"`
	// Prefix Prefix used to search for suggestions.
	Prefix *string `json:"prefix,omitempty"`
	// Regex A prefix expressed as a regular expression.
	Regex *string `json:"regex,omitempty"`
	// Term Suggests terms based on edit distance.
	Term *TermSuggester `json:"term,omitempty"`
	// Text The text to use as input for the suggester.
	// Needs to be set globally or per suggestion.
	Text *string `json:"text,omitempty"`
}

func (s *FieldSuggester) UnmarshalJSON(data []byte) error {

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

		case "completion":
			if err := dec.Decode(&s.Completion); err != nil {
				return err
			}

		case "phrase":
			if err := dec.Decode(&s.Phrase); err != nil {
				return err
			}

		case "prefix":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Prefix = &o

		case "regex":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Regex = &o

		case "term":
			if err := dec.Decode(&s.Term); err != nil {
				return err
			}

		case "text":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Text = &o

		}
	}
	return nil
}

// NewFieldSuggester returns a FieldSuggester.
func NewFieldSuggester() *FieldSuggester {
	r := &FieldSuggester{}

	return r
}
