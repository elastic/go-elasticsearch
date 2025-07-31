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

// AzureOpenAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L724-L769
type AzureOpenAIServiceSettings struct {
	// ApiKey A valid API key for your Azure OpenAI account.
	// You must specify either `api_key` or `entra_id`.
	// If you do not provide either or you provide both, you will receive an error
	// when you try to create your model.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	// After creating the inference model, you cannot change the associated API key.
	// If you want to use a different API key, delete the inference model and
	// recreate it with the same name and the updated API key.
	ApiKey *string `json:"api_key,omitempty"`
	// ApiVersion The Azure API version ID to use.
	// It is recommended to use the latest supported non-preview version.
	ApiVersion string `json:"api_version"`
	// DeploymentId The deployment name of your deployed models.
	// Your Azure OpenAI deployments can be found though the Azure OpenAI Studio
	// portal that is linked to your subscription.
	DeploymentId string `json:"deployment_id"`
	// EntraId A valid Microsoft Entra token.
	// You must specify either `api_key` or `entra_id`.
	// If you do not provide either or you provide both, you will receive an error
	// when you try to create your model.
	EntraId *string `json:"entra_id,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Azure.
	// The `azureopenai` service sets a default number of requests allowed per
	// minute depending on the task type.
	// For `text_embedding`, it is set to `1440`.
	// For `completion`, it is set to `120`.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// ResourceName The name of your Azure OpenAI resource.
	// You can find this from the list of resources in the Azure Portal for your
	// subscription.
	ResourceName string `json:"resource_name"`
}

func (s *AzureOpenAIServiceSettings) UnmarshalJSON(data []byte) error {

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
			s.ApiKey = &o

		case "api_version":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ApiVersion", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ApiVersion = o

		case "deployment_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "DeploymentId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.DeploymentId = o

		case "entra_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "EntraId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.EntraId = &o

		case "rate_limit":
			if err := dec.Decode(&s.RateLimit); err != nil {
				return fmt.Errorf("%s | %w", "RateLimit", err)
			}

		case "resource_name":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ResourceName", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ResourceName = o

		}
	}
	return nil
}

// NewAzureOpenAIServiceSettings returns a AzureOpenAIServiceSettings.
func NewAzureOpenAIServiceSettings() *AzureOpenAIServiceSettings {
	r := &AzureOpenAIServiceSettings{}

	return r
}
