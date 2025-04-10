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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/cohereembeddingtype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/coheresimilaritytype"
)

type _cohereServiceSettings struct {
	v *types.CohereServiceSettings
}

func NewCohereServiceSettings(apikey string) *_cohereServiceSettings {

	tmp := &_cohereServiceSettings{v: types.NewCohereServiceSettings()}

	tmp.ApiKey(apikey)

	return tmp

}

func (s *_cohereServiceSettings) ApiKey(apikey string) *_cohereServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_cohereServiceSettings) EmbeddingType(embeddingtype cohereembeddingtype.CohereEmbeddingType) *_cohereServiceSettings {

	s.v.EmbeddingType = &embeddingtype
	return s
}

func (s *_cohereServiceSettings) ModelId(modelid string) *_cohereServiceSettings {

	s.v.ModelId = &modelid

	return s
}

func (s *_cohereServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_cohereServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_cohereServiceSettings) Similarity(similarity coheresimilaritytype.CohereSimilarityType) *_cohereServiceSettings {

	s.v.Similarity = &similarity
	return s
}

func (s *_cohereServiceSettings) CohereServiceSettingsCaster() *types.CohereServiceSettings {
	return s.v
}
