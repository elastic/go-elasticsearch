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

type _azureAiStudioServiceSettings struct {
	v *types.AzureAiStudioServiceSettings
}

func NewAzureAiStudioServiceSettings(apikey string, endpointtype string, provider string, target string) *_azureAiStudioServiceSettings {

	tmp := &_azureAiStudioServiceSettings{v: types.NewAzureAiStudioServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.EndpointType(endpointtype)

	tmp.Provider(provider)

	tmp.Target(target)

	return tmp

}

// A valid API key of your Azure AI Studio model deployment.
// This key can be found on the overview page for your deployment in the
// management section of your Azure AI Studio account.
//
// IMPORTANT: You need to provide the API key only once, during the inference
// model creation.
// The get inference endpoint API does not retrieve your API key.
// After creating the inference model, you cannot change the associated API key.
// If you want to use a different API key, delete the inference model and
// recreate it with the same name and the updated API key.
func (s *_azureAiStudioServiceSettings) ApiKey(apikey string) *_azureAiStudioServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// The type of endpoint that is available for deployment through Azure AI
// Studio: `token` or `realtime`.
// The `token` endpoint type is for "pay as you go" endpoints that are billed
// per token.
// The `realtime` endpoint type is for "real-time" endpoints that are billed per
// hour of usage.
func (s *_azureAiStudioServiceSettings) EndpointType(endpointtype string) *_azureAiStudioServiceSettings {

	s.v.EndpointType = endpointtype

	return s
}

// The model provider for your deployment.
// Note that some providers may support only certain task types.
// Supported providers include:
//
// * `cohere` - available for `text_embedding` and `completion` task types
// * `databricks` - available for `completion` task type only
// * `meta` - available for `completion` task type only
// * `microsoft_phi` - available for `completion` task type only
// * `mistral` - available for `completion` task type only
// * `openai` - available for `text_embedding` and `completion` task types
func (s *_azureAiStudioServiceSettings) Provider(provider string) *_azureAiStudioServiceSettings {

	s.v.Provider = provider

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Azure AI Studio.
// By default, the `azureaistudio` service sets the number of requests allowed
// per minute to 240.
func (s *_azureAiStudioServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_azureAiStudioServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// The target URL of your Azure AI Studio model deployment.
// This can be found on the overview page for your deployment in the management
// section of your Azure AI Studio account.
func (s *_azureAiStudioServiceSettings) Target(target string) *_azureAiStudioServiceSettings {

	s.v.Target = target

	return s
}

func (s *_azureAiStudioServiceSettings) AzureAiStudioServiceSettingsCaster() *types.AzureAiStudioServiceSettings {
	return s.v
}
