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
// https://github.com/elastic/elasticsearch-specification/tree/ac9c431ec04149d9048f2b8f9731e3c2f7f38754

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/tokenizationtruncate"
)

// NlpRobertaTokenizationConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/ac9c431ec04149d9048f2b8f9731e3c2f7f38754/specification/ml/_types/inference.ts#L160-L187
type NlpRobertaTokenizationConfig struct {
	// AddPrefixSpace Should the tokenizer prefix input with a space character
	AddPrefixSpace *bool `json:"add_prefix_space,omitempty"`
	// MaxSequenceLength Maximum input sequence length for the model
	MaxSequenceLength *int `json:"max_sequence_length,omitempty"`
	// Span Tokenization spanning options. Special value of -1 indicates no spanning
	// takes place
	Span *int `json:"span,omitempty"`
	// Truncate Should tokenization input be automatically truncated before sending to the
	// model for inference
	Truncate *tokenizationtruncate.TokenizationTruncate `json:"truncate,omitempty"`
	// WithSpecialTokens Is tokenization completed with special tokens
	WithSpecialTokens *bool `json:"with_special_tokens,omitempty"`
}

func (s *NlpRobertaTokenizationConfig) UnmarshalJSON(data []byte) error {

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

		case "add_prefix_space":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.AddPrefixSpace = &value
			case bool:
				s.AddPrefixSpace = &v
			}

		case "max_sequence_length":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.MaxSequenceLength = &value
			case float64:
				f := int(v)
				s.MaxSequenceLength = &f
			}

		case "span":

			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return err
				}
				s.Span = &value
			case float64:
				f := int(v)
				s.Span = &f
			}

		case "truncate":
			if err := dec.Decode(&s.Truncate); err != nil {
				return err
			}

		case "with_special_tokens":
			var tmp interface{}
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return err
				}
				s.WithSpecialTokens = &value
			case bool:
				s.WithSpecialTokens = &v
			}

		}
	}
	return nil
}

// NewNlpRobertaTokenizationConfig returns a NlpRobertaTokenizationConfig.
func NewNlpRobertaTokenizationConfig() *NlpRobertaTokenizationConfig {
	r := &NlpRobertaTokenizationConfig{}

	return r
}
