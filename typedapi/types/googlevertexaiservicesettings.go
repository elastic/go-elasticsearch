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

// GoogleVertexAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L1332-L1358
type GoogleVertexAIServiceSettings struct {
	// Location The name of the location to use for the inference task.
	// Refer to the Google documentation for the list of supported locations.
	Location string `json:"location"`
	// ModelId The name of the model to use for the inference task.
	// Refer to the Google documentation for the list of supported models.
	ModelId string `json:"model_id"`
	// ProjectId The name of the project to use for the inference task.
	ProjectId string `json:"project_id"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Google Vertex AI.
	// By default, the `googlevertexai` service sets the number of requests allowed
	// per minute to 30.000.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// ServiceAccountJson A valid service account in JSON format for the Google Vertex AI API.
	ServiceAccountJson string `json:"service_account_json"`
}

func (s *GoogleVertexAIServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "location":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Location", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Location = o

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

		case "service_account_json":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "ServiceAccountJson", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.ServiceAccountJson = o

		}
	}
	return nil
}

// NewGoogleVertexAIServiceSettings returns a GoogleVertexAIServiceSettings.
func NewGoogleVertexAIServiceSettings() *GoogleVertexAIServiceSettings {
	r := &GoogleVertexAIServiceSettings{}

	return r
}
