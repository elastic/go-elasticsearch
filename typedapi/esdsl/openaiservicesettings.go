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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _openAIServiceSettings struct {
	v *types.OpenAIServiceSettings
}

func NewOpenAIServiceSettings(apikey string, modelid string) *_openAIServiceSettings {

	tmp := &_openAIServiceSettings{v: types.NewOpenAIServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.ModelId(modelid)

	return tmp

}

func (s *_openAIServiceSettings) ApiKey(apikey string) *_openAIServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_openAIServiceSettings) Dimensions(dimensions int) *_openAIServiceSettings {

	s.v.Dimensions = &dimensions

	return s
}

func (s *_openAIServiceSettings) ModelId(modelid string) *_openAIServiceSettings {

	s.v.ModelId = modelid

	return s
}

func (s *_openAIServiceSettings) OrganizationId(organizationid string) *_openAIServiceSettings {

	s.v.OrganizationId = &organizationid

	return s
}

func (s *_openAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_openAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_openAIServiceSettings) Url(url string) *_openAIServiceSettings {

	s.v.Url = &url

	return s
}

func (s *_openAIServiceSettings) OpenAIServiceSettingsCaster() *types.OpenAIServiceSettings {
	return s.v
}
