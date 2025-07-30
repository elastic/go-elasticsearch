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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _watsonxServiceSettings struct {
	v *types.WatsonxServiceSettings
}

func NewWatsonxServiceSettings(apikey string, apiversion string, modelid string, projectid string, url string) *_watsonxServiceSettings {

	tmp := &_watsonxServiceSettings{v: types.NewWatsonxServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.ApiVersion(apiversion)

	tmp.ModelId(modelid)

	tmp.ProjectId(projectid)

	tmp.Url(url)

	return tmp

}

func (s *_watsonxServiceSettings) ApiKey(apikey string) *_watsonxServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_watsonxServiceSettings) ApiVersion(apiversion string) *_watsonxServiceSettings {

	s.v.ApiVersion = apiversion

	return s
}

func (s *_watsonxServiceSettings) ModelId(modelid string) *_watsonxServiceSettings {

	s.v.ModelId = modelid

	return s
}

func (s *_watsonxServiceSettings) ProjectId(projectid string) *_watsonxServiceSettings {

	s.v.ProjectId = projectid

	return s
}

func (s *_watsonxServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_watsonxServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_watsonxServiceSettings) Url(url string) *_watsonxServiceSettings {

	s.v.Url = url

	return s
}

func (s *_watsonxServiceSettings) WatsonxServiceSettingsCaster() *types.WatsonxServiceSettings {
	return s.v
}
