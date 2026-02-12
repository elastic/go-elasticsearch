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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaielementtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/jinaaisimilaritytype"
)

type _jinaAIServiceSettings struct {
	v *types.JinaAIServiceSettings
}

func NewJinaAIServiceSettings(apikey string, modelid string) *_jinaAIServiceSettings {

	tmp := &_jinaAIServiceSettings{v: types.NewJinaAIServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.ModelId(modelid)

	return tmp

}

func (s *_jinaAIServiceSettings) ApiKey(apikey string) *_jinaAIServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_jinaAIServiceSettings) Dimensions(dimensions int) *_jinaAIServiceSettings {

	s.v.Dimensions = &dimensions

	return s
}

func (s *_jinaAIServiceSettings) ElementType(elementtype jinaaielementtype.JinaAIElementType) *_jinaAIServiceSettings {

	s.v.ElementType = &elementtype
	return s
}

func (s *_jinaAIServiceSettings) ModelId(modelid string) *_jinaAIServiceSettings {

	s.v.ModelId = modelid

	return s
}

func (s *_jinaAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_jinaAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_jinaAIServiceSettings) Similarity(similarity jinaaisimilaritytype.JinaAISimilarityType) *_jinaAIServiceSettings {

	s.v.Similarity = &similarity
	return s
}

func (s *_jinaAIServiceSettings) JinaAIServiceSettingsCaster() *types.JinaAIServiceSettings {
	return s.v
}
