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

type _anthropicServiceSettings struct {
	v *types.AnthropicServiceSettings
}

func NewAnthropicServiceSettings(apikey string, modelid string) *_anthropicServiceSettings {

	tmp := &_anthropicServiceSettings{v: types.NewAnthropicServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.ModelId(modelid)

	return tmp

}

// A valid API key for the Anthropic API.
func (s *_anthropicServiceSettings) ApiKey(apikey string) *_anthropicServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// The name of the model to use for the inference task.
// Refer to the Anthropic documentation for the list of supported models.
func (s *_anthropicServiceSettings) ModelId(modelid string) *_anthropicServiceSettings {

	s.v.ModelId = modelid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Anthropic.
// By default, the `anthropic` service sets the number of requests allowed per
// minute to 50.
func (s *_anthropicServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_anthropicServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_anthropicServiceSettings) AnthropicServiceSettingsCaster() *types.AnthropicServiceSettings {
	return s.v
}
