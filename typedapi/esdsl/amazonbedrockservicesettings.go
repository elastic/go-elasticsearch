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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _amazonBedrockServiceSettings struct {
	v *types.AmazonBedrockServiceSettings
}

func NewAmazonBedrockServiceSettings(accesskey string, model string, region string, secretkey string) *_amazonBedrockServiceSettings {

	tmp := &_amazonBedrockServiceSettings{v: types.NewAmazonBedrockServiceSettings()}

	tmp.AccessKey(accesskey)

	tmp.Model(model)

	tmp.Region(region)

	tmp.SecretKey(secretkey)

	return tmp

}

// A valid AWS access key that has permissions to use Amazon Bedrock and access
// to models for inference requests.
func (s *_amazonBedrockServiceSettings) AccessKey(accesskey string) *_amazonBedrockServiceSettings {

	s.v.AccessKey = accesskey

	return s
}

// The base model ID or an ARN to a custom model based on a foundational model.
// The base model IDs can be found in the Amazon Bedrock documentation.
// Note that the model ID must be available for the provider chosen and your IAM
// user must have access to the model.
func (s *_amazonBedrockServiceSettings) Model(model string) *_amazonBedrockServiceSettings {

	s.v.Model = model

	return s
}

// The model provider for your deployment.
// Note that some providers may support only certain task types.
// Supported providers include:
//
// * `amazontitan` - available for `text_embedding` and `completion` task types
// * `anthropic` - available for `completion` task type only
// * `ai21labs` - available for `completion` task type only
// * `cohere` - available for `text_embedding` and `completion` task types
// * `meta` - available for `completion` task type only
// * `mistral` - available for `completion` task type only
func (s *_amazonBedrockServiceSettings) Provider(provider string) *_amazonBedrockServiceSettings {

	s.v.Provider = &provider

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Watsonx.
// By default, the `watsonxai` service sets the number of requests allowed per
// minute to 120.
func (s *_amazonBedrockServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_amazonBedrockServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// The region that your model or ARN is deployed in.
// The list of available regions per model can be found in the Amazon Bedrock
// documentation.
func (s *_amazonBedrockServiceSettings) Region(region string) *_amazonBedrockServiceSettings {

	s.v.Region = region

	return s
}

// A valid AWS secret key that is paired with the `access_key`.
// For informationg about creating and managing access and secret keys, refer to
// the AWS documentation.
func (s *_amazonBedrockServiceSettings) SecretKey(secretkey string) *_amazonBedrockServiceSettings {

	s.v.SecretKey = secretkey

	return s
}

func (s *_amazonBedrockServiceSettings) AmazonBedrockServiceSettingsCaster() *types.AmazonBedrockServiceSettings {
	return s.v
}
