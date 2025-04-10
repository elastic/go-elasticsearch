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

type _huggingFaceServiceSettings struct {
	v *types.HuggingFaceServiceSettings
}

func NewHuggingFaceServiceSettings(apikey string, url string) *_huggingFaceServiceSettings {

	tmp := &_huggingFaceServiceSettings{v: types.NewHuggingFaceServiceSettings()}

	tmp.ApiKey(apikey)

	tmp.Url(url)

	return tmp

}

func (s *_huggingFaceServiceSettings) ApiKey(apikey string) *_huggingFaceServiceSettings {

	s.v.ApiKey = apikey

	return s
}

func (s *_huggingFaceServiceSettings) RateLimit(ratelimit types.RateLimitSettingVariant) *_huggingFaceServiceSettings {

	s.v.RateLimit = ratelimit.RateLimitSettingCaster()

	return s
}

func (s *_huggingFaceServiceSettings) Url(url string) *_huggingFaceServiceSettings {

	s.v.Url = url

	return s
}

func (s *_huggingFaceServiceSettings) HuggingFaceServiceSettingsCaster() *types.HuggingFaceServiceSettings {
	return s.v
}
