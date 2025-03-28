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

type _googleVertexAIServiceSettings struct {
	v *types.GoogleVertexAIServiceSettings
}

func NewGoogleVertexAIServiceSettings(location string, modelid string, projectid string, serviceaccountjson string) *_googleVertexAIServiceSettings {

	tmp := &_googleVertexAIServiceSettings{v: types.NewGoogleVertexAIServiceSettings()}

	tmp.Location(location)

	tmp.ModelId(modelid)

	tmp.ProjectId(projectid)

	tmp.ServiceAccountJson(serviceaccountjson)

	return tmp

}

// The name of the location to use for the inference task.
// Refer to the Google documentation for the list of supported locations.
func (s *_googleVertexAIServiceSettings) Location(location string) *_googleVertexAIServiceSettings {

	s.v.Location = location

	return s
}

// The name of the model to use for the inference task.
// Refer to the Google documentation for the list of supported models.
func (s *_googleVertexAIServiceSettings) ModelId(modelid string) *_googleVertexAIServiceSettings {

	s.v.ModelId = modelid

	return s
}

// The name of the project to use for the inference task.
func (s *_googleVertexAIServiceSettings) ProjectId(projectid string) *_googleVertexAIServiceSettings {

	s.v.ProjectId = projectid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Google Vertex AI.
// By default, the `googlevertexai` service sets the number of requests allowed
// per minute to 30.000.
func (s *_googleVertexAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_googleVertexAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// A valid service account in JSON format for the Google Vertex AI API.
func (s *_googleVertexAIServiceSettings) ServiceAccountJson(serviceaccountjson string) *_googleVertexAIServiceSettings {

	s.v.ServiceAccountJson = serviceaccountjson

	return s
}

func (s *_googleVertexAIServiceSettings) GoogleVertexAIServiceSettingsCaster() *types.GoogleVertexAIServiceSettings {
	return s.v
}
