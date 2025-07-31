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

// AmazonBedrockServiceSettings type.
//
// https://github.com/elastic/elasticsearch-specification/blob/470b4b9aaaa25cae633ec690e54b725c6fc939c7/specification/inference/_types/CommonTypes.ts#L366-L408
type AmazonBedrockServiceSettings struct {
	// AccessKey A valid AWS access key that has permissions to use Amazon Bedrock and access
	// to models for inference requests.
	AccessKey string `json:"access_key"`
	// Model The base model ID or an ARN to a custom model based on a foundational model.
	// The base model IDs can be found in the Amazon Bedrock documentation.
	// Note that the model ID must be available for the provider chosen and your IAM
	// user must have access to the model.
	Model string `json:"model"`
	// Provider The model provider for your deployment.
	// Note that some providers may support only certain task types.
	// Supported providers include:
	//
	// * `amazontitan` - available for `text_embedding` and `completion` task types
	// * `anthropic` - available for `completion` task type only
	// * `ai21labs` - available for `completion` task type only
	// * `cohere` - available for `text_embedding` and `completion` task types
	// * `meta` - available for `completion` task type only
	// * `mistral` - available for `completion` task type only
	Provider *string `json:"provider,omitempty"`
	// RateLimit This setting helps to minimize the number of rate limit errors returned from
	// Watsonx.
	// By default, the `watsonxai` service sets the number of requests allowed per
	// minute to 120.
	RateLimit *RateLimitSetting `json:"rate_limit,omitempty"`
	// Region The region that your model or ARN is deployed in.
	// The list of available regions per model can be found in the Amazon Bedrock
	// documentation.
	Region string `json:"region"`
	// SecretKey A valid AWS secret key that is paired with the `access_key`.
	// For informationg about creating and managing access and secret keys, refer to
	// the AWS documentation.
	SecretKey string `json:"secret_key"`
}

func (s *AmazonBedrockServiceSettings) UnmarshalJSON(data []byte) error {

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

		case "access_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "AccessKey", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.AccessKey = o

		case "model":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Model", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Model = o

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
			s.Provider = &o

		case "rate_limit":
			if err := dec.Decode(&s.RateLimit); err != nil {
				return fmt.Errorf("%s | %w", "RateLimit", err)
			}

		case "region":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Region", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Region = o

		case "secret_key":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "SecretKey", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.SecretKey = o

		}
	}
	return nil
}

// NewAmazonBedrockServiceSettings returns a AmazonBedrockServiceSettings.
func NewAmazonBedrockServiceSettings() *AmazonBedrockServiceSettings {
	r := &AmazonBedrockServiceSettings{}

	return r
}
