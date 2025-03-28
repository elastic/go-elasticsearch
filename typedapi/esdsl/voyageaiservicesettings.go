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

type _voyageAIServiceSettings struct {
	v *types.VoyageAIServiceSettings
}

func NewVoyageAIServiceSettings(modelid string) *_voyageAIServiceSettings {

	tmp := &_voyageAIServiceSettings{v: types.NewVoyageAIServiceSettings()}

	tmp.ModelId(modelid)

	return tmp

}

// The number of dimensions for resulting output embeddings.
// This setting maps to `output_dimension` in the VoyageAI documentation.
// Only for the `text_embedding` task type.
func (s *_voyageAIServiceSettings) Dimensions(dimensions int) *_voyageAIServiceSettings {

	s.v.Dimensions = &dimensions

	return s
}

// The data type for the embeddings to be returned.
// This setting maps to `output_dtype` in the VoyageAI documentation.
// Permitted values: float, int8, bit.
// `int8` is a synonym of `byte` in the VoyageAI documentation.
// `bit` is a synonym of `binary` in the VoyageAI documentation.
// Only for the `text_embedding` task type.
func (s *_voyageAIServiceSettings) EmbeddingType(embeddingtype float32) *_voyageAIServiceSettings {

	s.v.EmbeddingType = &embeddingtype

	return s
}

// The name of the model to use for the inference task.
// Refer to the VoyageAI documentation for the list of available text embedding
// and rerank models.
func (s *_voyageAIServiceSettings) ModelId(modelid string) *_voyageAIServiceSettings {

	s.v.ModelId = modelid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// VoyageAI.
// The `voyageai` service sets a default number of requests allowed per minute
// depending on the task type.
// For both `text_embedding` and `rerank`, it is set to `2000`.
func (s *_voyageAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_voyageAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_voyageAIServiceSettings) VoyageAIServiceSettingsCaster() *types.VoyageAIServiceSettings {
	return s.v
}
