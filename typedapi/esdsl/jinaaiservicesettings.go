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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/jinaaisimilaritytype"
)

type _jinaAIServiceSettings struct {
	v *types.JinaAIServiceSettings
}

func NewJinaAIServiceSettings(apikey string) *_jinaAIServiceSettings {

	tmp := &_jinaAIServiceSettings{v: types.NewJinaAIServiceSettings()}

	tmp.ApiKey(apikey)

	return tmp

}

// A valid API key of your JinaAI account.
//
// IMPORTANT: You need to provide the API key only once, during the inference
// model creation.
// The get inference endpoint API does not retrieve your API key.
// After creating the inference model, you cannot change the associated API key.
// If you want to use a different API key, delete the inference model and
// recreate it with the same name and the updated API key.
func (s *_jinaAIServiceSettings) ApiKey(apikey string) *_jinaAIServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// The name of the model to use for the inference task.
// For a `rerank` task, it is required.
// For a `text_embedding` task, it is optional.
func (s *_jinaAIServiceSettings) ModelId(modelid string) *_jinaAIServiceSettings {

	s.v.ModelId = &modelid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// JinaAI.
// By default, the `jinaai` service sets the number of requests allowed per
// minute to 2000 for all task types.
func (s *_jinaAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_jinaAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// For a `text_embedding` task, the similarity measure. One of cosine,
// dot_product, l2_norm.
// The default values varies with the embedding type.
// For example, a float embedding type uses a `dot_product` similarity measure
// by default.
func (s *_jinaAIServiceSettings) Similarity(similarity jinaaisimilaritytype.JinaAISimilarityType) *_jinaAIServiceSettings {

	s.v.Similarity = &similarity
	return s
}

func (s *_jinaAIServiceSettings) JinaAIServiceSettingsCaster() *types.JinaAIServiceSettings {
	return s.v
}
