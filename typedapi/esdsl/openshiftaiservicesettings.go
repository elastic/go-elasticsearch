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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/openshiftaisimilaritytype"
)

type _openShiftAiServiceSettings struct {
	v *types.OpenShiftAiServiceSettings
}

func NewOpenShiftAiServiceSettings(apikey string, url string) *_openShiftAiServiceSettings {

	tmp := &_openShiftAiServiceSettings{v: types.NewOpenShiftAiServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.Url(url)

	return tmp

}

func (s *_openShiftAiServiceSettings) ApiKey(apikey string) *_openShiftAiServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_openShiftAiServiceSettings) MaxInputTokens(maxinputtokens int) *_openShiftAiServiceSettings {

	s.v.MaxInputTokens = &maxinputtokens

	return s
}

func (s *_openShiftAiServiceSettings) ModelId(modelid string) *_openShiftAiServiceSettings {

	s.v.ModelId = &modelid

	return s
}

func (s *_openShiftAiServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_openShiftAiServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_openShiftAiServiceSettings) Similarity(similarity openshiftaisimilaritytype.OpenShiftAiSimilarityType) *_openShiftAiServiceSettings {

	s.v.Similarity = &similarity
	return s
}

func (s *_openShiftAiServiceSettings) Url(url string) *_openShiftAiServiceSettings {

	s.v.Url = url

	return s
}

func (s *_openShiftAiServiceSettings) OpenShiftAiServiceSettingsCaster() *types.OpenShiftAiServiceSettings {
	return s.v
}
