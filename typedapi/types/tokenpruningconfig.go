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

// TokenPruningConfig type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/_types/TokenPruningConfig.ts#L22-L35
type TokenPruningConfig struct {
	// OnlyScorePrunedTokens Whether to only score pruned tokens, vs only scoring kept tokens.
	OnlyScorePrunedTokens *bool `json:"only_score_pruned_tokens,omitempty"`
	// TokensFreqRatioThreshold Tokens whose frequency is more than this threshold times the average
	// frequency of all tokens in the specified field are considered outliers and
	// pruned.
	TokensFreqRatioThreshold *int `json:"tokens_freq_ratio_threshold,omitempty"`
	// TokensWeightThreshold Tokens whose weight is less than this threshold are considered nonsignificant
	// and pruned.
	TokensWeightThreshold *float32 `json:"tokens_weight_threshold,omitempty"`
}

func (s *TokenPruningConfig) UnmarshalJSON(data []byte) error {

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

		case "only_score_pruned_tokens":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "OnlyScorePrunedTokens", err)
				}
				s.OnlyScorePrunedTokens = &value
			case bool:
				s.OnlyScorePrunedTokens = &v
			}

		case "tokens_freq_ratio_threshold":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TokensFreqRatioThreshold", err)
				}
				s.TokensFreqRatioThreshold = &value
			case float64:
				f := int(v)
				s.TokensFreqRatioThreshold = &f
			}

		case "tokens_weight_threshold":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseFloat(v, 32)
				if err != nil {
					return fmt.Errorf("%s | %w", "TokensWeightThreshold", err)
				}
				f := float32(value)
				s.TokensWeightThreshold = &f
			case float64:
				f := float32(v)
				s.TokensWeightThreshold = &f
			}

		}
	}
	return nil
}

// NewTokenPruningConfig returns a TokenPruningConfig.
func NewTokenPruningConfig() *TokenPruningConfig {
	r := &TokenPruningConfig{}

	return r
}
