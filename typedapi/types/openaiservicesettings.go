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
// https://github.com/elastic/elasticsearch-specification/tree/37285cbd3fd155f913b50d880b40ec45f9df64b3

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openaisimilaritytype"
)

// OpenAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/37285cbd3fd155f913b50d880b40ec45f9df64b3/specification/inference/_types/CommonTypes.ts#L2152-L2239
type OpenAIServiceSettings struct {
	// ApiKey A valid API key of your OpenAI account. You can find your OpenAI API keys in
	// your OpenAI account under the API keys section.
	//
	// IMPORTANT: You must specify either `api_key` or `client_secret`. If you do
	// not provide one or you provide more than one of them, you will receive an
	// error when you try to create your endpoint.
	ApiKey *string `json:"api_key,omitempty"`
	// ClientId For OAuth 2.0 authorization using the client credentials grant flow. The
	// application ID that's assigned to your app.
	//
	// IMPORTANT: To configure OAuth 2.0, you must specify `client_id`, `scopes`,
	// `token_url`, and `client_secret` together. If one of the fields is missing,
	// you will receive an error when you try to create your endpoint.
	ClientId *string `json:"client_id,omitempty"`
	// ClientSecret For OAuth 2.0 authorization using the client credentials grant flow. The
	// application secret that you created for your app.
	//
	// IMPORTANT: You must specify either `api_key` or `client_secret`. If you do
	// not provide one or you provide more than one of them, you will receive an
	// error when you try to create your endpoint.
	//
	// IMPORTANT: To configure OAuth 2.0, you must specify `client_id`, `scopes`,
	// `token_url`, and `client_secret` together. If one of the fields is missing,
	// you will receive an error when you try to create your endpoint.
	ClientSecret *string `json:"client_secret,omitempty"`
	// Dimensions For a `text_embedding` or `embedding` task, the number of dimensions the
	// resulting output embeddings should have. It is supported only in
	// `text-embedding-3` and later models. If it is not set, the OpenAI defined
	// default for the model is used.
	Dimensions *int `json:"dimensions,omitempty"`
	// ModelId The name of the model to use for the inference task. Refer to the OpenAI
	// documentation for the list of available text embedding models.
	ModelId string `json:"model_id"`
	// OrganizationId The unique identifier for your organization. You can find the Organization ID
	// in your OpenAI account under *Settings > Organizations*.
	OrganizationId *string `json:"organization_id,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// OpenAI. The `openai` service sets a default number of requests allowed per
	// minute depending on the task type. For `text_embedding` and `embedding`, it
	// is set to `3000`. For `completion` and `chat_completion`, it is set to `500`.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Scopes For OAuth 2.0 authorization using the client credentials grant flow. The
	// resource identifier of the resource you want. For example:
	//
	//	"scopes": [
	//	  "scope1",
	//	  "scope2"
	//	]
	//
	// IMPORTANT: To configure OAuth 2.0, you must specify `client_id`, `scopes`,
	// `token_url`, and `client_secret` together. If one of the fields is missing,
	// you will receive an error when you try to create your endpoint.
	Scopes []string `json:"scopes,omitempty"`
	// Similarity For a `text_embedding` or `embedding` task, the similarity measure. One of
	// `cosine`, `dot_product`, `l2_norm`. Defaults to `dot_product`.
	Similarity *openaisimilaritytype.OpenAISimilarityType `json:"similarity,omitempty"`
	// TokenUrl For OAuth 2.0 authorization using the client credentials grant flow. An
	// OAuth2 token endpoint where Elasticsearch sends a request to exchange client
	// credentials for an access token.
	//
	// IMPORTANT: To configure OAuth 2.0, you must specify `client_id`, `scopes`,
	// `token_url`, and `client_secret` together. If one of the fields is missing,
	// you will receive an error when you try to create your endpoint.
	TokenUrl *string `json:"token_url,omitempty"`
	// Url The URL endpoint to use for the requests. It can be changed for testing
	// purposes. Default value is `https://api.openai.com/v1/embeddings` for a
	// `text_embedding` or `embedding` task,
	// `https://api.openai.com/v1/chat/completions` for a `completion` or
	// `chat_completion` task.
	Url *string `json:"url,omitempty"`
}

func (s *OpenAIServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "organization_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "OrganizationId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.OrganizationId = &o

		case "rate_limit":
			if err := dec.Decode(&s.RateLimit); err != nil {
				return fmt.Errorf("%s | %w", "RateLimit", err)
			}

		case "scopes":
			if err := dec.Decode(&s.Scopes); err != nil {
				return fmt.Errorf("%s | %w", "Scopes", err)
			}

		case "similarity":
			if err := dec.Decode(&s.Similarity); err != nil {
				return fmt.Errorf("%s | %w", "Similarity", err)
			}

		case "token_url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "TokenUrl", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.TokenUrl = &o

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

// NewOpenAIServiceSettings returns a OpenAIServiceSettings.
func NewOpenAIServiceSettings() *OpenAIServiceSettings {
	r := &OpenAIServiceSettings{}

	return r
}

type OpenAIServiceSettingsVariant interface {
	OpenAIServiceSettingsCaster() *OpenAIServiceSettings
}

func (s *OpenAIServiceSettings) OpenAIServiceSettingsCaster() *OpenAIServiceSettings {
	return s
}
