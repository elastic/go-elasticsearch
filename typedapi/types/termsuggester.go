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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/stringdistance"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestsort"
)

// TermSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/search/_types/suggester.ts#L256-L269
type TermSuggester struct {
	Analyzer       *string                        `json:"analyzer,omitempty"`
	Field          string                         `json:"field"`
	LowercaseTerms *bool                          `json:"lowercase_terms,omitempty"`
	MaxEdits       *int                           `json:"max_edits,omitempty"`
	MaxInspections *int                           `json:"max_inspections,omitempty"`
	MaxTermFreq    *float32                       `json:"max_term_freq,omitempty"`
	MinDocFreq     *float32                       `json:"min_doc_freq,omitempty"`
	MinWordLength  *int                           `json:"min_word_length,omitempty"`
	PrefixLength   *int                           `json:"prefix_length,omitempty"`
	ShardSize      *int                           `json:"shard_size,omitempty"`
	Size           *int                           `json:"size,omitempty"`
	Sort           *suggestsort.SuggestSort       `json:"sort,omitempty"`
	StringDistance *stringdistance.StringDistance `json:"string_distance,omitempty"`
	SuggestMode    *suggestmode.SuggestMode       `json:"suggest_mode,omitempty"`
	Text           *string                        `json:"text,omitempty"`
}

func (s *TermSuggester) UnmarshalJSON(data []byte) error {

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

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "lowercase_terms":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.LowercaseTerms = &value
			case bool:
				s.LowercaseTerms = &v
			}

		case "max_edits":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxEdits = &value
			case float64:
				f := int(v)
				s.MaxEdits = &f
			}

		case "max_inspections":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxInspections = &value
			case float64:
				f := int(v)
				s.MaxInspections = &f
			}

		case "max_term_freq":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.MaxTermFreq = &f
			case float64:
				f := float32(v)
				s.MaxTermFreq = &f
			}

		case "min_doc_freq":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.MinDocFreq = &f
			case float64:
				f := float32(v)
				s.MinDocFreq = &f
			}

		case "min_word_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinWordLength = &value
			case float64:
				f := int(v)
				s.MinWordLength = &f
			}

		case "prefix_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.PrefixLength = &value
			case float64:
				f := int(v)
				s.PrefixLength = &f
			}

		case "shard_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.ShardSize = &value
			case float64:
				f := int(v)
				s.ShardSize = &f
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

		case "sort":
			if err := dec.Decode(&s.Sort); err != nil {
				return err
			}

		case "string_distance":
			if err := dec.Decode(&s.StringDistance); err != nil {
				return err
			}

		case "suggest_mode":
			if err := dec.Decode(&s.SuggestMode); err != nil {
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

// NewTermSuggester returns a TermSuggester.
func NewTermSuggester() *TermSuggester {
	r := &TermSuggester{}

	return r
}
