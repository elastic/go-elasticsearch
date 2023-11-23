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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/stringdistance"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestsort"
)

// TermSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/_global/search/_types/suggester.ts#L503-L565
type TermSuggester struct {
	// Analyzer The analyzer to analyze the suggest text with.
	// Defaults to the search analyzer of the suggest field.
	Analyzer *string `json:"analyzer,omitempty"`
	// Field The field to fetch the candidate suggestions from.
	// Needs to be set globally or per suggestion.
	Field          string `json:"field"`
	LowercaseTerms *bool  `json:"lowercase_terms,omitempty"`
	// MaxEdits The maximum edit distance candidate suggestions can have in order to be
	// considered as a suggestion.
	// Can only be `1` or `2`.
	MaxEdits *int `json:"max_edits,omitempty"`
	// MaxInspections A factor that is used to multiply with the shard_size in order to inspect
	// more candidate spelling corrections on the shard level.
	// Can improve accuracy at the cost of performance.
	MaxInspections *int `json:"max_inspections,omitempty"`
	// MaxTermFreq The maximum threshold in number of documents in which a suggest text token
	// can exist in order to be included.
	// Can be a relative percentage number (for example `0.4`) or an absolute number
	// to represent document frequencies.
	// If a value higher than 1 is specified, then fractional can not be specified.
	MaxTermFreq *float32 `json:"max_term_freq,omitempty"`
	// MinDocFreq The minimal threshold in number of documents a suggestion should appear in.
	// This can improve quality by only suggesting high frequency terms.
	// Can be specified as an absolute number or as a relative percentage of number
	// of documents.
	// If a value higher than 1 is specified, then the number cannot be fractional.
	MinDocFreq *float32 `json:"min_doc_freq,omitempty"`
	// MinWordLength The minimum length a suggest text term must have in order to be included.
	MinWordLength *int `json:"min_word_length,omitempty"`
	// PrefixLength The number of minimal prefix characters that must match in order be a
	// candidate for suggestions.
	// Increasing this number improves spellcheck performance.
	PrefixLength *int `json:"prefix_length,omitempty"`
	// ShardSize Sets the maximum number of suggestions to be retrieved from each individual
	// shard.
	ShardSize *int `json:"shard_size,omitempty"`
	// Size The maximum corrections to be returned per suggest text token.
	Size *int `json:"size,omitempty"`
	// Sort Defines how suggestions should be sorted per suggest text term.
	Sort *suggestsort.SuggestSort `json:"sort,omitempty"`
	// StringDistance The string distance implementation to use for comparing how similar suggested
	// terms are.
	StringDistance *stringdistance.StringDistance `json:"string_distance,omitempty"`
	// SuggestMode Controls what suggestions are included or controls for what suggest text
	// terms, suggestions should be suggested.
	SuggestMode *suggestmode.SuggestMode `json:"suggest_mode,omitempty"`
	// Text The suggest text.
	// Needs to be set globally or per suggestion.
	Text *string `json:"text,omitempty"`
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
