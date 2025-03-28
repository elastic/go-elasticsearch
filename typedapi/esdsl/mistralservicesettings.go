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

type _mistralServiceSettings struct {
	v *types.MistralServiceSettings
}

func NewMistralServiceSettings(apikey string, model string) *_mistralServiceSettings {

	tmp := &_mistralServiceSettings{v: types.NewMistralServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.Model(model)

	return tmp

}

// A valid API key of your Mistral account.
// You can find your Mistral API keys or you can create a new one on the API
// Keys page.
//
// IMPORTANT: You need to provide the API key only once, during the inference
// model creation.
// The get inference endpoint API does not retrieve your API key.
// After creating the inference model, you cannot change the associated API key.
// If you want to use a different API key, delete the inference model and
// recreate it with the same name and the updated API key.
func (s *_mistralServiceSettings) ApiKey(apikey string) *_mistralServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// The maximum number of tokens per input before chunking occurs.
func (s *_mistralServiceSettings) MaxInputTokens(maxinputtokens int) *_mistralServiceSettings {

	s.v.MaxInputTokens = &maxinputtokens

	return s
}

// The name of the model to use for the inference task.
// Refer to the Mistral models documentation for the list of available text
// embedding models.
func (s *_mistralServiceSettings) Model(model string) *_mistralServiceSettings {

	s.v.Model = model

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// the Mistral API.
// By default, the `mistral` service sets the number of requests allowed per
// minute to 240.
func (s *_mistralServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_mistralServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_mistralServiceSettings) MistralServiceSettingsCaster() *types.MistralServiceSettings {
	return s.v
}
