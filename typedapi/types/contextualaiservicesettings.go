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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// ContextualAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/e196f9953fa743572ee46884835f1934bce9a16b/specification/inference/_types/CommonTypes.ts#L1217-L1238
type ContextualAIServiceSettings struct {
	// ApiKey A valid API key for your Contexutual AI account.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	ApiKey string `json:"api_key"`
	// ModelId The name of the model to use for the inference task.
	// Refer to the Contextual AI documentation for the list of available rerank
	// models.
	ModelId string `json:"model_id"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Contextual AI.
	// The `contextualai` service sets a default number of requests allowed per
	// minute depending on the task type.
	// For `rerank`, it is set to `1000`.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
}

func (s *ContextualAIServiceSettings) UnmarshalJSON(data []byte) error {

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

		}
	}
	return nil
}

// NewContextualAIServiceSettings returns a ContextualAIServiceSettings.
func NewContextualAIServiceSettings() *ContextualAIServiceSettings {
	r := &ContextualAIServiceSettings{}

	return r
}

type ContextualAIServiceSettingsVariant interface {
	ContextualAIServiceSettingsCaster() *ContextualAIServiceSettings
}

func (s *ContextualAIServiceSettings) ContextualAIServiceSettingsCaster() *ContextualAIServiceSettings {
	return s
}
