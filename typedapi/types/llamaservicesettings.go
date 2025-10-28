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
// https://github.com/elastic/elasticsearch-specification/tree/d520d9e8cf14cad487de5e0654007686c395b494

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/llamasimilaritytype"
)

// LlamaServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/d520d9e8cf14cad487de5e0654007686c395b494/specification/inference/_types/CommonTypes.ts#L1654-L1684
type LlamaServiceSettings struct {
	// MaxInputTokens For a `text_embedding` task, the maximum number of tokens per input before
	// chunking occurs.
	MaxInputTokens *int `json:"max_input_tokens,omitempty"`
	// ModelId The name of the model to use for the inference task.
	// Refer to the Llama downloading models documentation for different ways of
	// getting a list of available models and downloading them.
	// Service has been tested and confirmed to be working with the following
	// models:
	// * For `text_embedding` task - `all-MiniLM-L6-v2`.
	// * For `completion` and `chat_completion` tasks - `llama3.2:3b`.
	ModelId string `json:"model_id"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// the Llama API.
	// By default, the `llama` service sets the number of requests allowed per
	// minute to 3000.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Similarity For a `text_embedding` task, the similarity measure. One of cosine,
	// dot_product, l2_norm.
	Similarity *llamasimilaritytype.LlamaSimilarityType `json:"similarity,omitempty"`
	// Url The URL endpoint of the Llama stack endpoint.
	// URL must contain:
	// * For `text_embedding` task - `/v1/inference/embeddings`.
	// * For `completion` and `chat_completion` tasks -
	// `/v1/openai/v1/chat/completions`.
	Url string `json:"url"`
}

func (s *LlamaServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "max_input_tokens":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxInputTokens", err)
				}
				s.MaxInputTokens = &value
			case float64:
				f := int(v)
				s.MaxInputTokens = &f
			}

		case "model_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ModelId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ModelId = o

		case "rate_limit":
			if err := dec.Decode(&s.RateLimit); err != nil {
				return fmt.Errorf("%s | %w", "RateLimit", err)
			}

		case "similarity":
			if err := dec.Decode(&s.Similarity); err != nil {
				return fmt.Errorf("%s | %w", "Similarity", err)
			}

		case "url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Url", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Url = o

		}
	}
	return nil
}

// NewLlamaServiceSettings returns a LlamaServiceSettings.
func NewLlamaServiceSettings() *LlamaServiceSettings {
	r := &LlamaServiceSettings{}

	return r
}

type LlamaServiceSettingsVariant interface {
	LlamaServiceSettingsCaster() *LlamaServiceSettings
}

func (s *LlamaServiceSettings) LlamaServiceSettingsCaster() *LlamaServiceSettings {
	return s
}
