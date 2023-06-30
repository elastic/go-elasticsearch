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
// https://github.com/elastic/elasticsearch-specification/tree/a0da620389f06553c0727f98f95e40dbb564fcca

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"
)

// CompletionSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/search/_types/suggester.ts#L130-L135
type CompletionSuggester struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Contexts       map[string][]CompletionContext `json:"contexts,omitempty"`
	Field          string                         `json:"field"`
	Fuzzy          *SuggestFuzziness              `json:"fuzzy,omitempty"`
	Regex          *RegexOptions                  `json:"regex,omitempty"`
	Size           *int                           `json:"size,omitempty"`
	SkipDuplicates *bool                          `json:"skip_duplicates,omitempty"`
}

func (s *CompletionSuggester) UnmarshalJSON(data []byte) error {

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
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Analyzer = &o

		case "contexts":
			if s.Contexts == nil {
				s.Contexts = make(map[string][]CompletionContext, 0)
			}
			rawMsg := make(map[string]json.RawMessage, 0)
			dec.Decode(&rawMsg)
			for key, value := range rawMsg {
				switch {
				case bytes.HasPrefix(value, []byte("\"")), bytes.HasPrefix(value, []byte("{")):
					o := NewCompletionContext()
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return err
					}
					s.Contexts[key] = append(s.Contexts[key], *o)
				default:
					o := []CompletionContext{}
					err := json.NewDecoder(bytes.NewReader(value)).Decode(&o)
					if err != nil {
						return err
					}
					s.Contexts[key] = o
				}
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "fuzzy":
			if err := dec.Decode(&s.Fuzzy); err != nil {
				return err
			}

		case "regex":
			if err := dec.Decode(&s.Regex); err != nil {
				return err
			}

		case "size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Size = &value
			case float64:
				f := int(v)
				s.Size = &f
			}

		case "skip_duplicates":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.SkipDuplicates = &value
			case bool:
				s.SkipDuplicates = &v
			}

		}
	}
	return nil
}

// NewCompletionSuggester returns a CompletionSuggester.
func NewCompletionSuggester() *CompletionSuggester {
	r := &CompletionSuggester{
		Contexts: make(map[string][]CompletionContext, 0),
	}

	return r
}
