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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/googlemodelgardenprovider"
)

type _googleVertexAIServiceSettings struct {
	v *types.GoogleVertexAIServiceSettings
}

func NewGoogleVertexAIServiceSettings(serviceaccountjson string) *_googleVertexAIServiceSettings {

	tmp := &_googleVertexAIServiceSettings{v: types.NewGoogleVertexAIServiceSettings()}

	tmp.ServiceAccountJson(serviceaccountjson)

	return tmp

}

func (s *_googleVertexAIServiceSettings) Dimensions(dimensions int) *_googleVertexAIServiceSettings {

	s.v.Dimensions = &dimensions

	return s
}

func (s *_googleVertexAIServiceSettings) Location(location string) *_googleVertexAIServiceSettings {

	s.v.Location = &location

	return s
}

func (s *_googleVertexAIServiceSettings) ModelId(modelid string) *_googleVertexAIServiceSettings {

	s.v.ModelId = &modelid

	return s
}

func (s *_googleVertexAIServiceSettings) ProjectId(projectid string) *_googleVertexAIServiceSettings {

	s.v.ProjectId = &projectid

	return s
}

func (s *_googleVertexAIServiceSettings) Provider(provider googlemodelgardenprovider.GoogleModelGardenProvider) *_googleVertexAIServiceSettings {

	s.v.Provider = &provider
	return s
}

func (s *_googleVertexAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_googleVertexAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_googleVertexAIServiceSettings) ServiceAccountJson(serviceaccountjson string) *_googleVertexAIServiceSettings {

	s.v.ServiceAccountJson = serviceaccountjson

	return s
}

func (s *_googleVertexAIServiceSettings) StreamingUrl(streamingurl string) *_googleVertexAIServiceSettings {

	s.v.StreamingUrl = &streamingurl

	return s
}

func (s *_googleVertexAIServiceSettings) Url(url string) *_googleVertexAIServiceSettings {

	s.v.Url = &url

	return s
}

func (s *_googleVertexAIServiceSettings) GoogleVertexAIServiceSettingsCaster() *types.GoogleVertexAIServiceSettings {
	return s.v
}
