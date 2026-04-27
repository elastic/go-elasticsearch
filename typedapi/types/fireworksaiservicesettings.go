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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fireworksaisimilaritytype"
)

// FireworksAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/inference/_types/CommonTypes.ts#L2042-L2082
type FireworksAIServiceSettings struct {
	// ApiKey A valid API key for your Fireworks AI account. You can find or create your
	// API keys in the Fireworks AI dashboard.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation. The get inference endpoint API does not retrieve your API
	// key.
	ApiKey string `json:"api_key"`
	// Dimensions For a `text_embedding` task, the number of dimensions the resulting output
	// embeddings should have. Variable-length embeddings are supported via this
	// parameter.
	Dimensions *int `json:"dimensions,omitempty"`
	// ModelId The name of the model to use for the inference task. Refer to the Fireworks
	// AI documentation for the list of available models for chat completion,
	// completion, and text embedding. For text embedding, supported models include
	// the Qwen3 embedding family (e.g. `fireworks/qwen3-embedding-8b`) and other
	// models in the Fireworks model library.
	ModelId string `json:"model_id"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// the Fireworks AI API. Rate limit grouping is per API key only. By default,
	// the `fireworksai` service sets the number of requests allowed per minute to
	// 6000.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Similarity For a `text_embedding` task, the similarity measure. One of cosine,
	// dot_product, l2_norm.
	Similarity *fireworksaisimilaritytype.FireworksAISimilarityType `json:"similarity,omitempty"`
	// Url The URL endpoint to use for the requests. If not provided, the default
	// Fireworks AI API endpoint is used.
	Url *string `json:"url,omitempty"`
}

func (s *FireworksAIServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "dimensions":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Dimensions", err)
				}
				s.Dimensions = &value
			case float64:
				f := int(v)
				s.Dimensions = &f
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

// NewFireworksAIServiceSettings returns a FireworksAIServiceSettings.
func NewFireworksAIServiceSettings() *FireworksAIServiceSettings {
	r := &FireworksAIServiceSettings{}

	return r
}

type FireworksAIServiceSettingsVariant interface {
	FireworksAIServiceSettingsCaster() *FireworksAIServiceSettings
}

func (s *FireworksAIServiceSettings) FireworksAIServiceSettingsCaster() *FireworksAIServiceSettings {
	return s
}
