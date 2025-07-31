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

// WatsonxServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1684-L1721
type WatsonxServiceSettings struct {
	// ApiKey A valid API key of your Watsonx account.
	// You can find your Watsonx API keys or you can create a new one on the API
	// keys page.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	// After creating the inference model, you cannot change the associated API key.
	// If you want to use a different API key, delete the inference model and
	// recreate it with the same name and the updated API key.
	ApiKey string `json:"api_key"`
	// ApiVersion A version parameter that takes a version date in the format of `YYYY-MM-DD`.
	// For the active version data parameters, refer to the Wastonx documentation.
	ApiVersion string `json:"api_version"`
	// ModelId The name of the model to use for the inference task.
	// Refer to the IBM Embedding Models section in the Watsonx documentation for
	// the list of available text embedding models.
	ModelId string `json:"model_id"`
	// ProjectId The identifier of the IBM Cloud project to use for the inference task.
	ProjectId string `json:"project_id"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Watsonx.
	// By default, the `watsonxai` service sets the number of requests allowed per
	// minute to 120.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Url The URL of the inference endpoint that you created on Watsonx.
	Url string `json:"url"`
}

func (s *WatsonxServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "project_id":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ProjectId", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ProjectId = o

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
			s.Url = o

		}
	}
	return nil
}

// NewWatsonxServiceSettings returns a WatsonxServiceSettings.
func NewWatsonxServiceSettings() *WatsonxServiceSettings {
	r := &WatsonxServiceSettings{}

	return r
}
