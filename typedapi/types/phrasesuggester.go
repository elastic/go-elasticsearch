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

// PhraseSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/search/_types/suggester.ts#L195-L209
type PhraseSuggester struct {
	Analyzer                *string                  `json:"analyzer,omitempty"`
	Collate                 *PhraseSuggestCollate    `json:"collate,omitempty"`
	Confidence              *Float64                 `json:"confidence,omitempty"`
	DirectGenerator         []DirectGenerator        `json:"direct_generator,omitempty"`
	Field                   string                   `json:"field"`
	ForceUnigrams           *bool                    `json:"force_unigrams,omitempty"`
	GramSize                *int                     `json:"gram_size,omitempty"`
	Highlight               *PhraseSuggestHighlight  `json:"highlight,omitempty"`
	MaxErrors               *Float64                 `json:"max_errors,omitempty"`
	RealWordErrorLikelihood *Float64                 `json:"real_word_error_likelihood,omitempty"`
	Separator               *string                  `json:"separator,omitempty"`
	ShardSize               *int                     `json:"shard_size,omitempty"`
	Size                    *int                     `json:"size,omitempty"`
	Smoothing               *SmoothingModelContainer `json:"smoothing,omitempty"`
	Text                    *string                  `json:"text,omitempty"`
	TokenLimit              *int                     `json:"token_limit,omitempty"`
}

func (s *PhraseSuggester) UnmarshalJSON(data []byte) error {

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

		case "collate":
			if err := dec.Decode(&s.Collate); err != nil {
				return err
			}

		case "confidence":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.Confidence = &f
			case float64:
				f := Float64(v)
				s.Confidence = &f
			}

		case "direct_generator":
			if err := dec.Decode(&s.DirectGenerator); err != nil {
				return err
			}

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
			}

		case "force_unigrams":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.ForceUnigrams = &value
			case bool:
				s.ForceUnigrams = &v
			}

		case "gram_size":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.GramSize = &value
			case float64:
				f := int(v)
				s.GramSize = &f
			}

		case "highlight":
			if err := dec.Decode(&s.Highlight); err != nil {
				return err
			}

		case "max_errors":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.MaxErrors = &f
			case float64:
				f := Float64(v)
				s.MaxErrors = &f
			}

		case "real_word_error_likelihood":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 64)
				if err != nil {
					return err
				}
				f := Float64(value)
				s.RealWordErrorLikelihood = &f
			case float64:
				f := Float64(v)
				s.RealWordErrorLikelihood = &f
			}

		case "separator":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Separator = &o

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

		case "smoothing":
			if err := dec.Decode(&s.Smoothing); err != nil {
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

		case "token_limit":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.TokenLimit = &value
			case float64:
				f := int(v)
				s.TokenLimit = &f
			}

		}
	}
	return nil
}

// NewPhraseSuggester returns a PhraseSuggester.
func NewPhraseSuggester() *PhraseSuggester {
	r := &PhraseSuggester{}

	return r
}
