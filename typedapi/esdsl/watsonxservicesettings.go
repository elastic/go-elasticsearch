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

// A valid API key of your Watsonx account.
// You can find your Watsonx API keys or you can create a new one on the API
// keys page.
//
// IMPORTANT: You need to provide the API key only once, during the inference
// model creation.
// The get inference endpoint API does not retrieve your API key.
// After creating the inference model, you cannot change the associated API key.
// If you want to use a different API key, delete the inference model and
// recreate it with the same name and the updated API key.
func (s *_watsonxServiceSettings) ApiKey(apikey string) *_watsonxServiceSettings {

	s.v.ApiKey = apikey

	return s
}

// A version parameter that takes a version date in the format of `YYYY-MM-DD`.
// For the active version data parameters, refer to the Wastonx documentation.
func (s *_watsonxServiceSettings) ApiVersion(apiversion string) *_watsonxServiceSettings {

	s.v.ApiVersion = apiversion

	return s
}

// The name of the model to use for the inference task.
// Refer to the IBM Embedding Models section in the Watsonx documentation for
// the list of available text embedding models.
func (s *_watsonxServiceSettings) ModelId(modelid string) *_watsonxServiceSettings {

	s.v.ModelId = modelid

	return s
}

// The identifier of the IBM Cloud project to use for the inference task.
func (s *_watsonxServiceSettings) ProjectId(projectid string) *_watsonxServiceSettings {

	s.v.ProjectId = projectid

	return s
}

// This setting helps to minimize the number of rate limit errors returned from
// Watsonx.
// By default, the `watsonxai` service sets the number of requests allowed per
// minute to 120.
func (s *_watsonxServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_watsonxServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

// The URL of the inference endpoint that you created on Watsonx.
func (s *_watsonxServiceSettings) Url(url string) *_watsonxServiceSettings {

	s.v.Url = url

	return s
}

func (s *_watsonxServiceSettings) WatsonxServiceSettingsCaster() *types.WatsonxServiceSettings {
	return s.v
}
