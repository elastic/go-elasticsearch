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

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/googlemodelgardenprovider"
)

// GoogleVertexAIServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/2514615770f18dbb4e3887cc1a279995dbfd0724/specification/inference/_types/CommonTypes.ts#L1548-L1633
type GoogleVertexAIServiceSettings struct {
	// Dimensions For a `text_embedding` task, the number of dimensions the resulting output
	// embeddings should have.
	// By default, the model's standard output dimension is used.
	// Refer to the Google documentation for more information.
	Dimensions *int `json:"dimensions,omitempty"`
	// Location The name of the location to use for the inference task for the Google Vertex
	// AI inference task.
	// For Google Vertex AI, when `provider` is omitted or `google` `location` is
	// mandatory.
	// For Google Model Garden's `completion` and `chat_completion` tasks, when
	// `provider` is a supported non-`google` value - `location` is ignored.
	// Refer to the Google documentation for the list of supported locations.
	Location *string `json:"location,omitempty"`
	// MaxBatchSize Only applicable for the `text_embedding` task type.
	// Controls the batch size of chunked inference requests sent to Google Vertex
	// AI.
	//
	// Setting this parameter lower reduces the risk of exceeding token limits but
	// may result in more API calls. Setting it higher increases throughput but may
	// risk hitting token limits.
	//
	// To estimate a safe `max_batch_size` value, you can use it together with the
	// `max_chunk_size` parameter using the following formula:
	// `max_batch_size ≈ max_chunk_size × 1.3 × 512 ÷ 20000`
	//
	// Where:
	// - `1.3` is an approximate tokens-per-word ratio
	// - `512` is the maximum number of chunks that can be generated per document
	// - `20000` is the Google Vertex AI token limit per request
	//
	// This estimate assumes the worst-case scenario with a document generating the
	// maximum 512 chunks.
	MaxBatchSize *int `json:"max_batch_size,omitempty"`
	// ModelId The name of the model to use for the inference task.
	// For Google Vertex AI `model_id` is mandatory.
	// For Google Model Garden's `completion` and `chat_completion` tasks, when
	// `provider` is a supported non-`google` value - `model_id` will be used for
	// some providers that require it, otherwise - ignored.
	// Refer to the Google documentation for the list of supported models for Google
	// Vertex AI.
	ModelId *string `json:"model_id,omitempty"`
	// ProjectId The name of the project to use for the Google Vertex AI inference task.
	// For Google Vertex AI `project_id` is mandatory.
	// For Google Model Garden's `completion` and `chat_completion` tasks, when
	// `provider` is a supported non-`google` value - `project_id` is ignored.
	ProjectId *string `json:"project_id,omitempty"`
	// Provider The name of the Google Model Garden Provider for `completion` and
	// `chat_completion` tasks.
	// In order for a Google Model Garden endpoint to be used `provider` must be
	// defined and be other than `google`.
	// Modes:
	// - Google Model Garden (third-party models): set `provider` to a supported
	// non-`google` value and provide `url` and/or `streaming_url`.
	// - Google Vertex AI: omit `provider` or set it to `google`. In this mode, do
	// not set `url` or `streaming_url` and Elastic will construct the endpoint url
	// from `location`, `model_id`, and `project_id` parameters.
	Provider *googlemodelgardenprovider.GoogleModelGardenProvider `json:"provider,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Google Vertex AI.
	// By default, the `googlevertexai` service sets the number of requests allowed
	// per minute to 30.000.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// ServiceAccountJson A valid service account in JSON format for the Google Vertex AI API.
	ServiceAccountJson string `json:"service_account_json"`
	// StreamingUrl The URL for streaming `completion` and `chat_completion` requests to a Google
	// Model Garden provider endpoint.
	// If both `streaming_url` and `url` are provided, each is used for its
	// respective mode.
	// If `url` is not provided, `streaming_url` is also used for non-streaming
	// `completion` requests.
	// If `provider` is not provided or set to `google` (Google Vertex AI), do not
	// set `streaming_url` (or `url`).
	// At least one of `streaming_url` or `url` must be provided for Google Model
	// Garden endpoint usage.
	// Certain providers require separate URLs for streaming and non-streaming
	// operations (e.g., Anthropic, Mistral, AI21). Others support both operation
	// types through a single URL (e.g., Meta, Hugging Face).
	// Information on constructing the URL for various providers can be found in the
	// Google Model Garden documentation for the model, or on the endpoint’s `Sample
	// request` page. The request examples also illustrate the proper formatting for
	// the `streaming_url`.
	StreamingUrl *string `json:"streaming_url,omitempty"`
	// Url The URL for non-streaming `completion` requests to a Google Model Garden
	// provider endpoint.
	// If both `url` and `streaming_url` are provided, each is used for its
	// respective mode.
	// If `streaming_url` is not provided, `url` is also used for streaming
	// `completion` and `chat_completion`.
	// If `provider` is not provided or set to `google` (Google Vertex AI), do not
	// set `url` (or `streaming_url`).
	// At least one of `url` or `streaming_url` must be provided for Google Model
	// Garden endpoint usage.
	// Certain providers require separate URLs for streaming and non-streaming
	// operations (e.g., Anthropic, Mistral, AI21). Others support both operation
	// types through a single URL (e.g., Meta, Hugging Face).
	// Information on constructing the URL for various providers can be found in the
	// Google Model Garden documentation for the model, or on the endpoint’s `Sample
	// request` page. The request examples also illustrate the proper formatting for
	// the `url`.
	Url *string `json:"url,omitempty"`
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
			s.Location = &o

		case "max_batch_size":

			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "MaxBatchSize", err)
				}
				s.MaxBatchSize = &value
			case float64:
				f := int(v)
				s.MaxBatchSize = &f
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
			s.ProjectId = &o

		case "provider":
			if err := dec.Decode(&s.Provider); err != nil {
				return fmt.Errorf("%s | %w", "Provider", err)
			}

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

		case "streaming_url":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "StreamingUrl", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.StreamingUrl = &o

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

// NewGoogleVertexAIServiceSettings returns a GoogleVertexAIServiceSettings.
func NewGoogleVertexAIServiceSettings() *GoogleVertexAIServiceSettings {
	r := &GoogleVertexAIServiceSettings{}

	return r
}

type GoogleVertexAIServiceSettingsVariant interface {
	GoogleVertexAIServiceSettingsCaster() *GoogleVertexAIServiceSettings
}

func (s *GoogleVertexAIServiceSettings) GoogleVertexAIServiceSettingsCaster() *GoogleVertexAIServiceSettings {
	return s
}
