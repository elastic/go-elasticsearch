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

// AzureAiStudioServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L641-L683
type AzureAiStudioServiceSettings struct {
	// ApiKey A valid API key of your Azure AI Studio model deployment.
	// This key can be found on the overview page for your deployment in the
	// management section of your Azure AI Studio account.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	// After creating the inference model, you cannot change the associated API key.
	// If you want to use a different API key, delete the inference model and
	// recreate it with the same name and the updated API key.
	ApiKey string `json:"api_key"`
	// EndpointType The type of endpoint that is available for deployment through Azure AI
	// Studio: `token` or `realtime`.
	// The `token` endpoint type is for "pay as you go" endpoints that are billed
	// per token.
	// The `realtime` endpoint type is for "real-time" endpoints that are billed per
	// hour of usage.
	EndpointType string `json:"endpoint_type"`
	// Provider The model provider for your deployment.
	// Note that some providers may support only certain task types.
	// Supported providers include:
	//
	// * `cohere` - available for `text_embedding` and `completion` task types
	// * `databricks` - available for `completion` task type only
	// * `meta` - available for `completion` task type only
	// * `microsoft_phi` - available for `completion` task type only
	// * `mistral` - available for `completion` task type only
	// * `openai` - available for `text_embedding` and `completion` task types
	Provider string `json:"provider"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Azure AI Studio.
	// By default, the `azureaistudio` service sets the number of requests allowed
	// per minute to 240.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Target The target URL of your Azure AI Studio model deployment.
	// This can be found on the overview page for your deployment in the management
	// section of your Azure AI Studio account.
	Target string `json:"target"`
}

func (s *AzureAiStudioServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "endpoint_type":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "EndpointType", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.EndpointType = o

		case "provider":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Provider", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Provider = o

		case "rate_limit":
			if err := dec.Decode(&s.RateLimit); err != nil {
				return fmt.Errorf("%s | %w", "RateLimit", err)
			}

		case "target":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Target", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Target = o

		}
	}
	return nil
}

// NewAzureAiStudioServiceSettings returns a AzureAiStudioServiceSettings.
func NewAzureAiStudioServiceSettings() *AzureAiStudioServiceSettings {
	r := &AzureAiStudioServiceSettings{}

	return r
}
