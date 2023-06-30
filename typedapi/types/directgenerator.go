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

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
)

// DirectGenerator type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/search/_types/suggester.ts#L170-L182
type DirectGenerator struct {
	Field          string                   `json:"field"`
	MaxEdits       *int                     `json:"max_edits,omitempty"`
	MaxInspections *float32                 `json:"max_inspections,omitempty"`
	MaxTermFreq    *float32                 `json:"max_term_freq,omitempty"`
	MinDocFreq     *float32                 `json:"min_doc_freq,omitempty"`
	MinWordLength  *int                     `json:"min_word_length,omitempty"`
	PostFilter     *string                  `json:"post_filter,omitempty"`
	PreFilter      *string                  `json:"pre_filter,omitempty"`
	PrefixLength   *int                     `json:"prefix_length,omitempty"`
	Size           *int                     `json:"size,omitempty"`
	SuggestMode    *suggestmode.SuggestMode `json:"suggest_mode,omitempty"`
}

func (s *DirectGenerator) UnmarshalJSON(data []byte) error {

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

		case "field":
			if err := dec.Decode(&s.Field); err != nil {
				return err
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
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return err
				}
				f := float32(value)
				s.MaxInspections = &f
			case float64:
				f := float32(v)
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

		case "post_filter":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PostFilter = &o

		case "pre_filter":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return err
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.PreFilter = &o

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

		case "suggest_mode":
			if err := dec.Decode(&s.SuggestMode); err != nil {
				return err
			}

		}
	}
	return nil
}

// NewDirectGenerator returns a DirectGenerator.
func NewDirectGenerator() *DirectGenerator {
	r := &DirectGenerator{}

	return r
}
