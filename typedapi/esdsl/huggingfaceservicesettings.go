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

type _huggingFaceServiceSettings struct {
	v *types.HuggingFaceServiceSettings
}

func NewHuggingFaceServiceSettings(apikey string, url string) *_huggingFaceServiceSettings {

	tmp := &_huggingFaceServiceSettings{v: types.NewHuggingFaceServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.Url(url)

	return tmp

}

// A valid access token for your HuggingFace account.
// You can create or find your access tokens on the HuggingFace settings page.
//
// IMPORTANT: You need to provide the API key only once, during the inference
// model creation.
// The get inference endpoint API does not retrieve your API key.
// After creating the inference model, you cannot change the associated API key.
// If you want to use a different API key, delete the inference model and
// recreate it with the same name and the updated API key.
func (s *_huggingFaceServiceSettings) ApiKey(apikey string) *_huggingFaceServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Hugging Face.
// By default, the `hugging_face` service sets the number of requests allowed
// per minute to 3000.
func (s *_huggingFaceServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_huggingFaceServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// The URL endpoint to use for the requests.
func (s *_huggingFaceServiceSettings) Url(url string) *_huggingFaceServiceSettings {

	s.v.Url = url

	return s
}

func (s *_huggingFaceServiceSettings) HuggingFaceServiceSettingsCaster() *types.HuggingFaceServiceSettings {
	return s.v
}
