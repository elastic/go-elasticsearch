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
// https://github.com/elastic/elasticsearch-specification/tree/eb2e22fb2ac404e676d19bcc7bb089647f029026

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _openAITaskSettings struct {
	v *types.OpenAITaskSettings
}

func NewOpenAITaskSettings() *_openAITaskSettings {

	return &_openAITaskSettings{v: types.NewOpenAITaskSettings()}

}

func (s *_openAITaskSettings) Headers(headers map[string]string) *_openAITaskSettings {

	s.v.Headers = headers
	return s
}

func (s *_openAITaskSettings) AddHeader(key string, value string) *_openAITaskSettings {

	var tmp map[string]string
	if s.v.Headers == nil {
		s.v.Headers = make(map[string]string)
	} else {
		tmp = s.v.Headers
	}

	tmp[key] = value

	s.v.Headers = tmp
	return s
}

func (s *_openAITaskSettings) User(user string) *_openAITaskSettings {

	s.v.User = &user

	return s
}

func (s *_openAITaskSettings) OpenAITaskSettingsCaster() *types.OpenAITaskSettings {
	return s.v
}
