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
// https://github.com/elastic/elasticsearch-specification/tree/3a94b6715915b1e9311724a2614c643368eece90

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
)

// HuggingFaceServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/3a94b6715915b1e9311724a2614c643368eece90/specification/inference/_types/CommonTypes.ts#L942-L963
type HuggingFaceServiceSettings struct {
	// ApiKey A valid access token for your HuggingFace account.
	// You can create or find your access tokens on the HuggingFace settings page.
	//
	// IMPORTANT: You need to provide the API key only once, during the inference
	// model creation.
	// The get inference endpoint API does not retrieve your API key.
	// After creating the inference model, you cannot change the associated API key.
	// If you want to use a different API key, delete the inference model and
	// recreate it with the same name and the updated API key.
	ApiKey string `json:"api_key"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Hugging Face.
	// By default, the `hugging_face` service sets the number of requests allowed
	// per minute to 3000.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Url The URL endpoint to use for the requests.
	Url string `json:"url"`
}

func (s *HuggingFaceServiceSettings) UnmarshalJSON(data []byte) error {

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

// NewHuggingFaceServiceSettings returns a HuggingFaceServiceSettings.
func NewHuggingFaceServiceSettings() *HuggingFaceServiceSettings {
	r := &HuggingFaceServiceSettings{}

	return r
}
