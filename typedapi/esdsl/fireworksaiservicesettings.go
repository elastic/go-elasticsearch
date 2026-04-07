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
// https://github.com/elastic/elasticsearch-specification/tree/836fca874204ca4173ae5c36fb6b5107d28d2fc0

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fireworksaisimilaritytype"
)

type _fireworksAIServiceSettings struct {
	v *types.FireworksAIServiceSettings
}

func NewFireworksAIServiceSettings(apikey string, modelid string) *_fireworksAIServiceSettings {

	tmp := &_fireworksAIServiceSettings{v: types.NewFireworksAIServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.ModelId(modelid)

	return tmp

}

func (s *_fireworksAIServiceSettings) ApiKey(apikey string) *_fireworksAIServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_fireworksAIServiceSettings) Dimensions(dimensions int) *_fireworksAIServiceSettings {

	s.v.Dimensions = &dimensions

	return s
}

func (s *_fireworksAIServiceSettings) ModelId(modelid string) *_fireworksAIServiceSettings {

	s.v.ModelId = modelid

	return s
}

func (s *_fireworksAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_fireworksAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_fireworksAIServiceSettings) Similarity(similarity fireworksaisimilaritytype.FireworksAISimilarityType) *_fireworksAIServiceSettings {

	s.v.Similarity = &similarity
	return s
}

func (s *_fireworksAIServiceSettings) Url(url string) *_fireworksAIServiceSettings {

	s.v.Url = &url

	return s
}

func (s *_fireworksAIServiceSettings) FireworksAIServiceSettingsCaster() *types.FireworksAIServiceSettings {
	return s.v
}
