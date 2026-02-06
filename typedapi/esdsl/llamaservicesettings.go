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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/llamasimilaritytype"
)

type _llamaServiceSettings struct {
	v *types.LlamaServiceSettings
}

func NewLlamaServiceSettings(modelid string, url string) *_llamaServiceSettings {

	tmp := &_llamaServiceSettings{v: types.NewLlamaServiceSettings()}

	tmp.ModelId(modelid)

	tmp.Url(url)

	return tmp

}

func (s *_llamaServiceSettings) MaxInputTokens(maxinputtokens int) *_llamaServiceSettings {

	s.v.MaxInputTokens = &maxinputtokens

	return s
}

func (s *_llamaServiceSettings) ModelId(modelid string) *_llamaServiceSettings {

	s.v.ModelId = modelid

	return s
}

func (s *_llamaServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_llamaServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_llamaServiceSettings) Similarity(similarity llamasimilaritytype.LlamaSimilarityType) *_llamaServiceSettings {

	s.v.Similarity = &similarity
	return s
}

func (s *_llamaServiceSettings) Url(url string) *_llamaServiceSettings {

	s.v.Url = url

	return s
}

func (s *_llamaServiceSettings) LlamaServiceSettingsCaster() *types.LlamaServiceSettings {
	return s.v
}
