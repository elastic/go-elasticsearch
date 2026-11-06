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
// https://github.com/elastic/elasticsearch-specification/tree/3c28a85a71c6bf5dbfe27404e91d79340ad84f59

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// Breakdown of tokens used in the request.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/_types/Results.ts#L361-L369
type ChatCompletionCompletionTokensDetails struct {
	// ReasoningTokens The number of tokens used for reasoning by the model.
	ReasoningTokens *int `json:"reasoning_tokens,omitempty"`
}

func (s *ChatCompletionCompletionTokensDetails) UnmarshalJSON(data []byte) error {

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

		case "reasoning_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "ReasoningTokens", err)
				}
				s.ReasoningTokens = &value
			case float64:
				f := int(v)
				s.ReasoningTokens = &f
			}

		}
	}
	return nil
}

// NewChatCompletionCompletionTokensDetails returns a ChatCompletionCompletionTokensDetails.
func NewChatCompletionCompletionTokensDetails() *ChatCompletionCompletionTokensDetails {
	r := &ChatCompletionCompletionTokensDetails{}

	return r
}
