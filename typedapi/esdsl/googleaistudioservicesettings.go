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

type _googleAiStudioServiceSettings struct {
	v *types.GoogleAiStudioServiceSettings
}

func NewGoogleAiStudioServiceSettings(apikey string, modelid string) *_googleAiStudioServiceSettings {

	tmp := &_googleAiStudioServiceSettings{v: types.NewGoogleAiStudioServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.ModelId(modelid)

	return tmp

}

// A valid API key of your Google Gemini account.
func (s *_googleAiStudioServiceSettings) ApiKey(apikey string) *_googleAiStudioServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// The name of the model to use for the inference task.
// Refer to the Google documentation for the list of supported models.
func (s *_googleAiStudioServiceSettings) ModelId(modelid string) *_googleAiStudioServiceSettings {

	s.v.ModelId = modelid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Google AI Studio.
// By default, the `googleaistudio` service sets the number of requests allowed
// per minute to 360.
func (s *_googleAiStudioServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_googleAiStudioServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_googleAiStudioServiceSettings) GoogleAiStudioServiceSettingsCaster() *types.GoogleAiStudioServiceSettings {
	return s.v
}
