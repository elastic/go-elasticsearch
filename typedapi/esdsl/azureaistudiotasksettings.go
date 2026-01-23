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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _azureAiStudioTaskSettings struct {
	v *types.AzureAiStudioTaskSettings
}

func NewAzureAiStudioTaskSettings() *_azureAiStudioTaskSettings {

	return &_azureAiStudioTaskSettings{v: types.NewAzureAiStudioTaskSettings()}

}

func (s *_azureAiStudioTaskSettings) DoSample(dosample float32) *_azureAiStudioTaskSettings {

	s.v.DoSample = &dosample

	return s
}

func (s *_azureAiStudioTaskSettings) MaxNewTokens(maxnewtokens int) *_azureAiStudioTaskSettings {

	s.v.MaxNewTokens = &maxnewtokens

	return s
}

func (s *_azureAiStudioTaskSettings) ReturnDocuments(returndocuments bool) *_azureAiStudioTaskSettings {

	s.v.ReturnDocuments = &returndocuments

	return s
}

func (s *_azureAiStudioTaskSettings) Temperature(temperature float32) *_azureAiStudioTaskSettings {

	s.v.Temperature = &temperature

	return s
}

func (s *_azureAiStudioTaskSettings) TopN(topn int) *_azureAiStudioTaskSettings {

	s.v.TopN = &topn

	return s
}

func (s *_azureAiStudioTaskSettings) TopP(topp float32) *_azureAiStudioTaskSettings {

	s.v.TopP = &topp

	return s
}

func (s *_azureAiStudioTaskSettings) User(user string) *_azureAiStudioTaskSettings {

	s.v.User = &user

	return s
}

func (s *_azureAiStudioTaskSettings) AzureAiStudioTaskSettingsCaster() *types.AzureAiStudioTaskSettings {
	return s.v
}
