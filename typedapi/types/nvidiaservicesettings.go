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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/nvidiasimilaritytype"
)

// NvidiaServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/inference/_types/CommonTypes.ts#L1970-L2008
type NvidiaServiceSettings struct {
	// ApiKey A valid API key for your Nvidia endpoint.
	// Can be found in `API Keys` section of Nvidia account settings.
	ApiKey string `json:"api_key"`
	// MaxInputTokens For a `text_embedding` task, the maximum number of tokens per input. Inputs
	// exceeding this value are truncated prior to sending to the Nvidia API.
	MaxInputTokens *int `json:"max_input_tokens,omitempty"`
	// ModelId The name of the model to use for the inference task.
	// Refer to the model's documentation for the name if needed.
	// Service has been tested and confirmed to be working with the following
	// models:
	//
	// * For `text_embedding` task - `nvidia/llama-3.2-nv-embedqa-1b-v2`.
	// * For `completion` and `chat_completion` tasks -
	// `microsoft/phi-3-mini-128k-instruct`.
	// * For `rerank` task - `nv-rerank-qa-mistral-4b:1`.
	// Service doesn't support `text_embedding` task `baai/bge-m3` and
	// `nvidia/nvclip` models due to them not recognizing the `input_type`
	// parameter.
	ModelId string `json:"model_id"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// the Nvidia API.
	// By default, the `nvidia` service sets the number of requests allowed per
	// minute to 3000.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Similarity For a `text_embedding` task, the similarity measure. One of cosine,
	// dot_product, l2_norm.
	Similarity *nvidiasimilaritytype.NvidiaSimilarityType `json:"similarity,omitempty"`
	// Url The URL of the Nvidia model endpoint. If not provided, the default endpoint
	// URL is used depending on the task type:
	//
	// * For `text_embedding` task -
	// `https://integrate.api.nvidia.com/v1/embeddings`.
	// * For `completion` and `chat_completion` tasks -
	// `https://integrate.api.nvidia.com/v1/chat/completions`.
	// * For `rerank` task -
	// `https://ai.api.nvidia.com/v1/retrieval/nvidia/reranking`.
	Url *string `json:"url,omitempty"`
}

func (s *NvidiaServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "api_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ApiKey", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ApiKey = o

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
			s.Url = &o

		}
	}
	return nil
}

// NewNvidiaServiceSettings returns a NvidiaServiceSettings.
func NewNvidiaServiceSettings() *NvidiaServiceSettings {
	r := &NvidiaServiceSettings{}

	return r
}

type NvidiaServiceSettingsVariant interface {
	NvidiaServiceSettingsCaster() *NvidiaServiceSettings
}

func (s *NvidiaServiceSettings) NvidiaServiceSettingsCaster() *NvidiaServiceSettings {
	return s
}
