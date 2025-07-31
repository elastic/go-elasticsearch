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

// OpenAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1554-L1596
type OpenAIServiceSettings struct {
	// ApiKey A valid API key of your OpenAI account.
	// You can find your OpenAI API keys in your OpenAI account under the API keys
	// section.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	// After creating the inference model, you cannot change the associated API key.
	// If you want to use a different API key, delete the inference model and
	// recreate it with the same name and the updated API key.
	ApiKey string `json:"api_key"`
	// Dimensions The number of dimensions the resulting output embeddings should have.
	// It is supported only in `text-embedding-3` and later models.
	// If it is not set, the OpenAI defined default for the model is used.
	Dimensions *int `json:"dimensions,omitempty"`
	// ModelId The name of the model to use for the inference task.
	// Refer to the OpenAI documentation for the list of available text embedding
	// models.
	ModelId string `json:"model_id"`
	// OrganizationId The unique identifier for your organization.
	// You can find the Organization ID in your OpenAI account under *Settings >
	// Organizations*.
	OrganizationId *string `json:"organization_id,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// OpenAI.
	// The `openai` service sets a default number of requests allowed per minute
	// depending on the task type.
	// For `text_embedding`, it is set to `3000`.
	// For `completion`, it is set to `500`.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Url The URL endpoint to use for the requests.
	// It can be changed for testing purposes.
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
