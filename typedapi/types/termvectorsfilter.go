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

// TermVectorsFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a0da620389f06553c0727f98f95e40dbb564fcca/specification/_global/termvectors/types.ts#L49-L57
type TermVectorsFilter struct {
	MaxDocFreq    *int `json:"max_doc_freq,omitempty"`
	MaxNumTerms   *int `json:"max_num_terms,omitempty"`
	MaxTermFreq   *int `json:"max_term_freq,omitempty"`
	MaxWordLength *int `json:"max_word_length,omitempty"`
	MinDocFreq    *int `json:"min_doc_freq,omitempty"`
	MinTermFreq   *int `json:"min_term_freq,omitempty"`
	MinWordLength *int `json:"min_word_length,omitempty"`
}

func (s *TermVectorsFilter) UnmarshalJSON(data []byte) error {

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

		case "max_doc_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxDocFreq = &value
			case float64:
				f := int(v)
				s.MaxDocFreq = &f
			}

		case "max_num_terms":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxNumTerms = &value
			case float64:
				f := int(v)
				s.MaxNumTerms = &f
			}

		case "max_term_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxTermFreq = &value
			case float64:
				f := int(v)
				s.MaxTermFreq = &f
			}

		case "max_word_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxWordLength = &value
			case float64:
				f := int(v)
				s.MaxWordLength = &f
			}

		case "min_doc_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinDocFreq = &value
			case float64:
				f := int(v)
				s.MinDocFreq = &f
			}

		case "min_term_freq":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MinTermFreq = &value
			case float64:
				f := int(v)
				s.MinTermFreq = &f
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

		}
	}
	return nil
}

// NewTermVectorsFilter returns a TermVectorsFilter.
func NewTermVectorsFilter() *TermVectorsFilter {
	r := &TermVectorsFilter{}

	return r
}
