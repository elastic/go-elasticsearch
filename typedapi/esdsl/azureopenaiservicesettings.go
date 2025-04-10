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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _azureOpenAIServiceSettings struct {
	v *types.AzureOpenAIServiceSettings
}

func NewAzureOpenAIServiceSettings(apiversion string, deploymentid string, resourcename string) *_azureOpenAIServiceSettings {

	tmp := &_azureOpenAIServiceSettings{v: types.NewAzureOpenAIServiceSettings()}

	tmp.ApiVersion(apiversion)

	tmp.DeploymentId(deploymentid)

	tmp.ResourceName(resourcename)

	return tmp

}

func (s *_azureOpenAIServiceSettings) ApiKey(apikey string) *_azureOpenAIServiceSettings {

	s.v.ApiKey = &apikey

	return s
}

func (s *_azureOpenAIServiceSettings) ApiVersion(apiversion string) *_azureOpenAIServiceSettings {

	s.v.ApiVersion = apiversion

	return s
}

func (s *_azureOpenAIServiceSettings) DeploymentId(deploymentid string) *_azureOpenAIServiceSettings {

	s.v.DeploymentId = deploymentid

	return s
}

func (s *_azureOpenAIServiceSettings) EntraId(entraid string) *_azureOpenAIServiceSettings {

	s.v.EntraId = &entraid

	return s
}

func (s *_azureOpenAIServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_azureOpenAIServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_azureOpenAIServiceSettings) ResourceName(resourcename string) *_azureOpenAIServiceSettings {

	s.v.ResourceName = resourcename

	return s
}

func (s *_azureOpenAIServiceSettings) AzureOpenAIServiceSettingsCaster() *types.AzureOpenAIServiceSettings {
	return s.v
}
