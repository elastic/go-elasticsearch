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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openshiftaisimilaritytype"
)

// OpenShiftAiServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/inference/_types/CommonTypes.ts#L2135-L2168
type OpenShiftAiServiceSettings struct {
	// ApiKey A valid API key for your OpenShift AI endpoint.
	// Can be found in `Token authentication` section of model related information.
	ApiKey string `json:"api_key"`
	// MaxInputTokens For a `text_embedding` task, the maximum number of tokens per input before
	// chunking occurs.
	MaxInputTokens *int `json:"max_input_tokens,omitempty"`
	// ModelId The name of the model to use for the inference task.
	// Refer to the hosted model's documentation for the name if needed.
	// Service has been tested and confirmed to be working with the following
	// models:
	// * For `text_embedding` task - `gritlm-7b`.
	// * For `completion` and `chat_completion` tasks - `llama-31-8b-instruct`.
	// * For `rerank` task - `bge-reranker-v2-m3`.
	ModelId *string `json:"model_id,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// the OpenShift AI API.
	// By default, the `openshift_ai` service sets the number of requests allowed
	// per minute to 3000.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Similarity For a `text_embedding` task, the similarity measure. One of cosine,
	// dot_product, l2_norm.
	// If not specified, the default dot_product value is used.
	Similarity *openshiftaisimilaritytype.OpenShiftAiSimilarityType `json:"similarity,omitempty"`
	// Url The URL of the OpenShift AI hosted model endpoint.
	Url string `json:"url"`
}

func (s *OpenShiftAiServiceSettings) UnmarshalJSON(data []byte) error {

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
			s.ModelId = &o

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

// NewOpenShiftAiServiceSettings returns a OpenShiftAiServiceSettings.
func NewOpenShiftAiServiceSettings() *OpenShiftAiServiceSettings {
	r := &OpenShiftAiServiceSettings{}

	return r
}

type OpenShiftAiServiceSettingsVariant interface {
	OpenShiftAiServiceSettingsCaster() *OpenShiftAiServiceSettings
}

func (s *OpenShiftAiServiceSettings) OpenShiftAiServiceSettingsCaster() *OpenShiftAiServiceSettings {
	return s
}
