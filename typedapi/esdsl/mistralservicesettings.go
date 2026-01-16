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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _mistralServiceSettings struct {
	v *types.MistralServiceSettings
}

func NewMistralServiceSettings(apikey string, model string) *_mistralServiceSettings {

	tmp := &_mistralServiceSettings{v: types.NewMistralServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.Model(model)

	return tmp

}

func (s *_mistralServiceSettings) ApiKey(apikey string) *_mistralServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_mistralServiceSettings) MaxInputTokens(maxinputtokens int) *_mistralServiceSettings {

	s.v.MaxInputTokens = &maxinputtokens

	return s
}

func (s *_mistralServiceSettings) Model(model string) *_mistralServiceSettings {

	s.v.Model = model

	return s
}

func (s *_mistralServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_mistralServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_mistralServiceSettings) MistralServiceSettingsCaster() *types.MistralServiceSettings {
	return s.v
}
