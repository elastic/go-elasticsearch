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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/cohereembeddingtype"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/coheresimilaritytype"
)

type _cohereServiceSettings struct {
	v *types.CohereServiceSettings
}

func NewCohereServiceSettings(apikey string) *_cohereServiceSettings {

	tmp := &_cohereServiceSettings{v: types.NewCohereServiceSettings()}

	tmp.ApiKey(apikey)

	return tmp

}

// A valid API key for your Cohere account.
// You can find or create your Cohere API keys on the Cohere API key settings
// page.
//
// IMPORTANT: You need to provide the API key only once, during the inference
// model creation.
// The get inference endpoint API does not retrieve your API key.
// After creating the inference model, you cannot change the associated API key.
// If you want to use a different API key, delete the inference model and
// recreate it with the same name and the updated API key.
func (s *_cohereServiceSettings) ApiKey(apikey string) *_cohereServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// For a `text_embedding` task, the types of embeddings you want to get back.
// Use `byte` for signed int8 embeddings (this is a synonym of `int8`).
// Use `float` for the default float embeddings.
// Use `int8` for signed int8 embeddings.
func (s *_cohereServiceSettings) EmbeddingType(embeddingtype cohereembeddingtype.CohereEmbeddingType) *_cohereServiceSettings {

	s.v.EmbeddingType = &embeddingtype
	return s
}

// For a `completion`, `rerank`, or `text_embedding` task, the name of the model
// to use for the inference task.
//
// * For the available `completion` models, refer to the [Cohere command
// docs](https://docs.cohere.com/docs/models#command).
// * For the available `rerank` models, refer to the [Cohere rerank
// docs](https://docs.cohere.com/reference/rerank-1).
// * For the available `text_embedding` models, refer to [Cohere embed
// docs](https://docs.cohere.com/reference/embed).
//
// The default value for a text embedding task is `embed-english-v2.0`.
func (s *_cohereServiceSettings) ModelId(modelid string) *_cohereServiceSettings {

	s.v.ModelId = &modelid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Cohere.
// By default, the `cohere` service sets the number of requests allowed per
// minute to 10000.
func (s *_cohereServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_cohereServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// The similarity measure.
// If the `embedding_type` is `float`, the default value is `dot_product`.
// If the `embedding_type` is `int8` or `byte`, the default value is `cosine`.
func (s *_cohereServiceSettings) Similarity(similarity coheresimilaritytype.CohereSimilarityType) *_cohereServiceSettings {

	s.v.Similarity = &similarity
	return s
}

func (s *_cohereServiceSettings) CohereServiceSettingsCaster() *types.CohereServiceSettings {
	return s.v
}
