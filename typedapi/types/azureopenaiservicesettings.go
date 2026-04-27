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
)

// AzureOpenAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/eb2e22fb2ac404e676d19bcc7bb089647f029026/specification/inference/_types/CommonTypes.ts#L937-L1024
type AzureOpenAIServiceSettings struct {
	// ApiKey A valid API key for your Azure OpenAI account.
	//
	// IMPORTANT: You must specify either `api_key`, `entra_id`, or `client_secret`.
	// If you do not provide one or you provide more than one of them, you will
	// receive an error when you try to create your endpoint.
	ApiKey *string `json:"api_key,omitempty"`
	// ApiVersion The Azure API version ID to use. It is recommended to use the latest
	// supported non-preview version.
	ApiVersion string `json:"api_version"`
	// ClientId For OAuth 2.0 authentication using the client credentials grant flow. The
	// application ID that's assigned to your app.
	//
	// IMPORTANT: To configure OAuth 2.0, you must specify client_id, scopes,
	// tenant_id, and client_secret together. If one of the fields is missing, you
	// will receive an error when you try to create your endpoint.
	ClientId *string `json:"client_id,omitempty"`
	// ClientSecret For OAuth 2.0 authentication using the client credentials grant flow. The
	// application secret that you created in the Microsoft app registration portal
	// for your app.
	//
	// IMPORTANT: You must specify either `api_key`, `entra_id`, or `client_secret`.
	// If you do not provide one or you provide more than one of them, you will
	// receive an error when you try to create your endpoint.
	//
	// IMPORTANT: To configure OAuth 2.0, you must specify client_id, scopes,
	// tenant_id, and client_secret together. If one of the fields is missing, you
	// will receive an error when you try to create your endpoint.
	ClientSecret *string `json:"client_secret,omitempty"`
	// DeploymentId The deployment name of your deployed models. Your Azure OpenAI deployments
	// can be found though the Azure OpenAI Studio portal that is linked to your
	// subscription.
	DeploymentId string `json:"deployment_id"`
	// EntraId A valid Microsoft Entra token.
	//
	// IMPORTANT: You must specify either `api_key`, `entra_id`, or `client_secret`.
	// If you do not provide one or you provide more than one of them, you will
	// receive an error when you try to create your endpoint.
	EntraId *string `json:"entra_id,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Azure. The `azureopenai` service sets a default number of requests allowed
	// per minute depending on the task type. For `text_embedding`, it is set to
	// `1440`. For `completion` and `chat_completion`, it is set to `120`.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// ResourceName The name of your Azure OpenAI resource. You can find this from the list of
	// resources in the Azure Portal for your subscription.
	ResourceName string `json:"resource_name"`
	// Scopes For OAuth 2.0 authentication using the client credentials grant flow. The
	// resource identifier (application ID URI) of the resource you want, suffixed
	// with .default For example:
	//
	//	"scopes": [
	//	  "https://cognitiveservices.azure.com/.default"
	//	]
	//
	// IMPORTANT: To configure OAuth 2.0, you must specify client_id, scopes,
	// tenant_id, and client_secret together. If one of the fields is missing, you
	// will receive an error when you try to create your endpoint.
	Scopes []string `json:"scopes,omitempty"`
	// TenantId For OAuth 2.0 authentication using the client credentials grant flow. The
	// directory tenant the application plans to operate against.
	//
	// IMPORTANT: To configure OAuth 2.0, you must specify client_id, scopes,
	// tenant_id, and client_secret together. If one of the fields is missing, you
	// will receive an error when you try to create your endpoint.
	TenantId *string `json:"tenant_id,omitempty"`
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

		case "client_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ClientId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ClientId = &o

		case "client_secret":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ClientSecret", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ClientSecret = &o

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

		case "scopes":
			if err := dec.Decode(&s.Scopes); err != nil {
				return fmt.Errorf("%s | %w", "Scopes", err)
			}

		case "tenant_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TenantId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TenantId = &o

		}
	}
	return nil
}

// NewAzureOpenAIServiceSettings returns a AzureOpenAIServiceSettings.
func NewAzureOpenAIServiceSettings() *AzureOpenAIServiceSettings {
	r := &AzureOpenAIServiceSettings{}

	return r
}

type AzureOpenAIServiceSettingsVariant interface {
	AzureOpenAIServiceSettingsCaster() *AzureOpenAIServiceSettings
}

func (s *AzureOpenAIServiceSettings) AzureOpenAIServiceSettingsCaster() *AzureOpenAIServiceSettings {
	return s
}
