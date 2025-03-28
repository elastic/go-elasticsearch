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

type _azureOpenAIServiceSettings struct {
	v *types.AzureOpenAIServiceSettings
}

func NewAzureOpenAIServiceSettings(apiversion string, deploymentid string, resourcename string) *_azureOpenAIServiceSettings {

	tmp := &_azureOpenAIServiceSettings{v: types.NewAzureOpenAIServiceSettings()}

	tmp.ApiVersion(apiversion)

	tmp.DeploymentId(deploymentid)

	tmp.ResourceName(resourcename)

	return tmp

}

// A valid API key for your Azure OpenAI account.
// You must specify either `api_key` or `entra_id`.
// If you do not provide either or you provide both, you will receive an error
// when you try to create your model.
//
// IMPORTANT: You need to provide the API key only once, during the inference
// model creation.
// The get inference endpoint API does not retrieve your API key.
// After creating the inference model, you cannot change the associated API key.
// If you want to use a different API key, delete the inference model and
// recreate it with the same name and the updated API key.
func (s *_azureOpenAIServiceSettings) ApiKey(apikey string) *_azureOpenAIServiceSettings {

	s.v.ApiKey = &apikey

	return s
}

// The Azure API version ID to use.
// It is recommended to use the latest supported non-preview version.
func (s *_azureOpenAIServiceSettings) ApiVersion(apiversion string) *_azureOpenAIServiceSettings {

	s.v.ApiVersion = apiversion

	return s
}

// The deployment name of your deployed models.
// Your Azure OpenAI deployments can be found though the Azure OpenAI Studio
// portal that is linked to your subscription.
func (s *_azureOpenAIServiceSettings) DeploymentId(deploymentid string) *_azureOpenAIServiceSettings {

	s.v.DeploymentId = deploymentid

	return s
}

// A valid Microsoft Entra token.
// You must specify either `api_key` or `entra_id`.
// If you do not provide either or you provide both, you will receive an error
// when you try to create your model.
func (s *_azureOpenAIServiceSettings) EntraId(entraid string) *_azureOpenAIServiceSettings {

	s.v.EntraId = &entraid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Azure.
// The `azureopenai` service sets a default number of requests allowed per
// minute depending on the task type.
// For `text_embedding`, it is set to `1440`.
// For `completion`, it is set to `120`.
func (s *_azureOpenAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_azureOpenAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// The name of your Azure OpenAI resource.
// You can find this from the list of resources in the Azure Portal for your
// subscription.
func (s *_azureOpenAIServiceSettings) ResourceName(resourcename string) *_azureOpenAIServiceSettings {

	s.v.ResourceName = resourcename

	return s
}

func (s *_azureOpenAIServiceSettings) AzureOpenAIServiceSettingsCaster() *types.AzureOpenAIServiceSettings {
	return s.v
}
