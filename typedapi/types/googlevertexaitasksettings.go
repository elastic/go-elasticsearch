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
// https://github.com/elastic/elasticsearch-specification/tree/aa1459fbdcaf57c653729142b3b6e9982373bb1c

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// GoogleVertexAITaskSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/aa1459fbdcaf57c653729142b3b6e9982373bb1c/specification/inference/_types/CommonTypes.ts#L1515-L1538
type GoogleVertexAITaskSettings struct {
	// AutoTruncate For a `text_embedding` task, truncate inputs longer than the maximum token
	// length automatically.
	AutoTruncate *bool `json:"auto_truncate,omitempty"`
	// MaxTokens For `completion` and `chat_completion` tasks, specifies the `max_tokens`
	// value for requests sent to the Google Model Garden `anthropic` provider.
	// If `provider` is not set to `anthropic`, this field is ignored.
	// If `max_tokens` is specified - it must be a positive integer. If not
	// specified, the default value of 1024 is used.
	// Anthropic models require `max_tokens` to be set for each request. Please
	// refer to the Anthropic documentation for more information.
	MaxTokens *int `json:"max_tokens,omitempty"`
	// ThinkingConfig For a `completion` or `chat_completion` task, allows configuration of the
	// thinking features for the model.
	// Refer to the Google documentation for the allowable configurations for each
	// model type.
	ThinkingConfig *ThinkingConfig `json:"thinking_config,omitempty"`
	// TopN For a `rerank` task, the number of the top N documents that should be
	// returned.
	TopN *int `json:"top_n,omitempty"`
}

func (s *GoogleVertexAITaskSettings) UnmarshalJSON(data []byte) error {

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

		case "auto_truncate":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "AutoTruncate", err)
				}
				s.AutoTruncate = &value
			case bool:
				s.AutoTruncate = &v
			}

		case "max_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxTokens", err)
				}
				s.MaxTokens = &value
			case float64:
				f := int(v)
				s.MaxTokens = &f
			}

		case "thinking_config":
			if err := dec.Decode(&s.ThinkingConfig); err != nil {
				return fmt.Errorf("%s | %w", "ThinkingConfig", err)
			}

		case "top_n":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "TopN", err)
				}
				s.TopN = &value
			case float64:
				f := int(v)
				s.TopN = &f
			}

		}
	}
	return nil
}

// NewGoogleVertexAITaskSettings returns a GoogleVertexAITaskSettings.
func NewGoogleVertexAITaskSettings() *GoogleVertexAITaskSettings {
	r := &GoogleVertexAITaskSettings{}

	return r
}

type GoogleVertexAITaskSettingsVariant interface {
	GoogleVertexAITaskSettingsCaster() *GoogleVertexAITaskSettings
}

func (s *GoogleVertexAITaskSettings) GoogleVertexAITaskSettingsCaster() *GoogleVertexAITaskSettings {
	return s
}
