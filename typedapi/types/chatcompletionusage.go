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

// Token usage statistics for the request.
//
// https://github.com/elastic/elasticsearch-specification/blob/3c28a85a71c6bf5dbfe27404e91d79340ad84f59/specification/inference/_types/Results.ts#L371-L397
type ChatCompletionUsage struct {
	// CompletionTokens The number of tokens in the generated completion.
	CompletionTokens int `json:"completion_tokens"`
	// CompletionTokensDetails Breakdown of the tokens used in the completion. Omitted when no details are
	// available.
	CompletionTokensDetails *ChatCompletionCompletionTokensDetails `json:"completion_tokens_details,omitempty"`
	// PromptTokens The number of tokens in the prompt.
	PromptTokens int `json:"prompt_tokens"`
	// PromptTokensDetails Breakdown of the tokens used in the prompt. Omitted when no details are
	// available.
	PromptTokensDetails *ChatCompletionPromptTokensDetails `json:"prompt_tokens_details,omitempty"`
	// TotalTokens The total number of tokens used (prompt + completion).
	TotalTokens int `json:"total_tokens"`
}

func (s *ChatCompletionUsage) UnmarshalJSON(data []byte) error {

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

		case "completion_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "CompletionTokens", err)
				}
				s.CompletionTokens = value
			case float64:
				f := int(v)
				s.CompletionTokens = f
			}

		case "completion_tokens_details":
			if err := dec.Decode(&s.CompletionTokensDetails); err != nil {
				return fmt.Errorf("%s | %w", "CompletionTokensDetails", err)
			}

		case "prompt_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "PromptTokens", err)
				}
				s.PromptTokens = value
			case float64:
				f := int(v)
				s.PromptTokens = f
			}

		case "prompt_tokens_details":
			if err := dec.Decode(&s.PromptTokensDetails); err != nil {
				return fmt.Errorf("%s | %w", "PromptTokensDetails", err)
			}

		case "total_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TotalTokens", err)
				}
				s.TotalTokens = value
			case float64:
				f := int(v)
				s.TotalTokens = f
			}

		}
	}
	return nil
}

// NewChatCompletionUsage returns a ChatCompletionUsage.
func NewChatCompletionUsage() *ChatCompletionUsage {
	r := &ChatCompletionUsage{}

	return r
}
