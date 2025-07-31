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

// AmazonBedrockTaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L410-L434
type AmazonBedrockTaskSettings struct {
	// MaxNewTokens For a `completion` task, it sets the maximum number for the output tokens to
	// be generated.
	MaxNewTokens *int `json:"max_new_tokens,omitempty"`
	// Temperature For a `completion` task, it is a number between 0.0 and 1.0 that controls the
	// apparent creativity of the results.
	// At temperature 0.0 the model is most deterministic, at temperature 1.0 most
	// random.
	// It should not be used if `top_p` or `top_k` is specified.
	Temperature *float32 `json:"temperature,omitempty"`
	// TopK For a `completion` task, it limits samples to the top-K most likely words,
	// balancing coherence and variability.
	// It is only available for anthropic, cohere, and mistral providers.
	// It is an alternative to `temperature`; it should not be used if `temperature`
	// is specified.
	TopK *float32 `json:"top_k,omitempty"`
	// TopP For a `completion` task, it is a number in the range of 0.0 to 1.0, to
	// eliminate low-probability tokens.
	// Top-p uses nucleus sampling to select top tokens whose sum of likelihoods
	// does not exceed a certain value, ensuring both variety and coherence.
	// It is an alternative to `temperature`; it should not be used if `temperature`
	// is specified.
	TopP *float32 `json:"top_p,omitempty"`
}

func (s *AmazonBedrockTaskSettings) UnmarshalJSON(data []byte) error {

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

		case "max_new_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxNewTokens", err)
				}
				s.MaxNewTokens = &value
			case float64:
				f := int(v)
				s.MaxNewTokens = &f
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
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopK", err)
				}
				f := float32(value)
				s.TopK = &f
			case float64:
				f := float32(v)
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

// NewAmazonBedrockTaskSettings returns a AmazonBedrockTaskSettings.
func NewAmazonBedrockTaskSettings() *AmazonBedrockTaskSettings {
	r := &AmazonBedrockTaskSettings{}

	return r
}
