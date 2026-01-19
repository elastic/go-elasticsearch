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
// https://github.com/elastic/elasticsearch-specification/tree/6785a6caa1fa3ca5ab3308963d79dce923a3469f

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _amazonBedrockTaskSettings struct {
	v *types.AmazonBedrockTaskSettings
}

func NewAmazonBedrockTaskSettings() *_amazonBedrockTaskSettings {

	return &_amazonBedrockTaskSettings{v: types.NewAmazonBedrockTaskSettings()}

}

func (s *_amazonBedrockTaskSettings) MaxNewTokens(maxnewtokens int) *_amazonBedrockTaskSettings {

	s.v.MaxNewTokens = &maxnewtokens

	return s
}

func (s *_amazonBedrockTaskSettings) Temperature(temperature float32) *_amazonBedrockTaskSettings {

	s.v.Temperature = &temperature

	return s
}

func (s *_amazonBedrockTaskSettings) TopK(topk float32) *_amazonBedrockTaskSettings {

	s.v.TopK = &topk

	return s
}

func (s *_amazonBedrockTaskSettings) TopP(topp float32) *_amazonBedrockTaskSettings {

	s.v.TopP = &topp

	return s
}

func (s *_amazonBedrockTaskSettings) AmazonBedrockTaskSettingsCaster() *types.AmazonBedrockTaskSettings {
	return s.v
}
