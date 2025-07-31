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
// https://github.com/elastic/elasticsearch-specification/tree/470b4b9aaaa25cae633ec690e54b725c6fc939c7

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// AnthropicTaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L606-L631
type AnthropicTaskSettings struct {
	// MaxTokens For a `completion` task, it is the maximum number of tokens to generate
	// before stopping.
	MaxTokens int `json:"max_tokens"`
	// Temperature For a `completion` task, it is the amount of randomness injected into the
	// response.
	// For more details about the supported range, refer to Anthropic documentation.
	Temperature *float32 `json:"temperature,omitempty"`
	// TopK For a `completion` task, it specifies to only sample from the top K options
	// for each subsequent token.
	// It is recommended for advanced use cases only.
	// You usually only need to use `temperature`.
	TopK *int `json:"top_k,omitempty"`
	// TopP For a `completion` task, it specifies to use Anthropic's nucleus sampling.
	// In nucleus sampling, Anthropic computes the cumulative distribution over all
	// the options for each subsequent token in decreasing probability order and
	// cuts it off once it reaches the specified probability.
	// You should either alter `temperature` or `top_p`, but not both.
	// It is recommended for advanced use cases only.
	// You usually only need to use `temperature`.
	TopP *float32 `json:"top_p,omitempty"`
}

func (s *AnthropicTaskSettings) UnmarshalJSON(data []byte) error {

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

		case "max_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxTokens", err)
				}
				s.MaxTokens = value
			case float64:
				f := int(v)
				s.MaxTokens = f
			}

		case "temperature":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "Temperature", err)
				}
				f := float32(value)
				s.Temperature = &f
			case float64:
				f := float32(v)
				s.Temperature = &f
			}

		case "top_k":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopK", err)
				}
				s.TopK = &value
			case float64:
				f := int(v)
				s.TopK = &f
			}

		case "top_p":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopP", err)
				}
				f := float32(value)
				s.TopP = &f
			case float64:
				f := float32(v)
				s.TopP = &f
			}

		}
	}
	return nil
}

// NewAnthropicTaskSettings returns a AnthropicTaskSettings.
func NewAnthropicTaskSettings() *AnthropicTaskSettings {
	r := &AnthropicTaskSettings{}

	return r
}
