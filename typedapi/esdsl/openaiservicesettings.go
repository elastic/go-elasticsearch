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

type _openAIServiceSettings struct {
	v *types.OpenAIServiceSettings
}

func NewOpenAIServiceSettings(apikey string, modelid string) *_openAIServiceSettings {

	tmp := &_openAIServiceSettings{v: types.NewOpenAIServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.ModelId(modelid)

	return tmp

}

// A valid API key of your OpenAI account.
// You can find your OpenAI API keys in your OpenAI account under the API keys
// section.
//
// IMPORTANT: You need to provide the API key only once, during the inference
// model creation.
// The get inference endpoint API does not retrieve your API key.
// After creating the inference model, you cannot change the associated API key.
// If you want to use a different API key, delete the inference model and
// recreate it with the same name and the updated API key.
func (s *_openAIServiceSettings) ApiKey(apikey string) *_openAIServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// The number of dimensions the resulting output embeddings should have.
// It is supported only in `text-embedding-3` and later models.
// If it is not set, the OpenAI defined default for the model is used.
func (s *_openAIServiceSettings) Dimensions(dimensions int) *_openAIServiceSettings {

	s.v.Dimensions = &dimensions

	return s
}

// The name of the model to use for the inference task.
// Refer to the OpenAI documentation for the list of available text embedding
// models.
func (s *_openAIServiceSettings) ModelId(modelid string) *_openAIServiceSettings {

	s.v.ModelId = modelid

	return s
}

// The unique identifier for your organization.
// You can find the Organization ID in your OpenAI account under *Settings >
// Organizations*.
func (s *_openAIServiceSettings) OrganizationId(organizationid string) *_openAIServiceSettings {

	s.v.OrganizationId = &organizationid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// OpenAI.
// The `openai` service sets a default number of requests allowed per minute
// depending on the task type.
// For `text_embedding`, it is set to `3000`.
// For `completion`, it is set to `500`.
func (s *_openAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_openAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// The URL endpoint to use for the requests.
// It can be changed for testing purposes.
func (s *_openAIServiceSettings) Url(url string) *_openAIServiceSettings {

	s.v.Url = &url

	return s
}

func (s *_openAIServiceSettings) OpenAIServiceSettingsCaster() *types.OpenAIServiceSettings {
	return s.v
}
