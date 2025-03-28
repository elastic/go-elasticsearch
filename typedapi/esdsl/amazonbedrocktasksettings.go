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

type _amazonBedrockTaskSettings struct {
	v *types.AmazonBedrockTaskSettings
}

func NewAmazonBedrockTaskSettings() *_amazonBedrockTaskSettings {

	return &_amazonBedrockTaskSettings{v: types.NewAmazonBedrockTaskSettings()}

}

// For a `completion` task, it sets the maximum number for the output tokens to
// be generated.
func (s *_amazonBedrockTaskSettings) MaxNewTokens(maxnewtokens int) *_amazonBedrockTaskSettings {

	s.v.MaxNewTokens = &maxnewtokens

	return s
}

// For a `completion` task, it is a number between 0.0 and 1.0 that controls the
// apparent creativity of the results.
// At temperature 0.0 the model is most deterministic, at temperature 1.0 most
// random.
// It should not be used if `top_p` or `top_k` is specified.
func (s *_amazonBedrockTaskSettings) Temperature(temperature float32) *_amazonBedrockTaskSettings {

	s.v.Temperature = &temperature

	return s
}

// For a `completion` task, it limits samples to the top-K most likely words,
// balancing coherence and variability.
// It is only available for anthropic, cohere, and mistral providers.
// It is an alternative to `temperature`; it should not be used if `temperature`
// is specified.
func (s *_amazonBedrockTaskSettings) TopK(topk float32) *_amazonBedrockTaskSettings {

	s.v.TopK = &topk

	return s
}

// For a `completion` task, it is a number in the range of 0.0 to 1.0, to
// eliminate low-probability tokens.
// Top-p uses nucleus sampling to select top tokens whose sum of likelihoods
// does not exceed a certain value, ensuring both variety and coherence.
// It is an alternative to `temperature`; it should not be used if `temperature`
// is specified.
func (s *_amazonBedrockTaskSettings) TopP(topp float32) *_amazonBedrockTaskSettings {

	s.v.TopP = &topp

	return s
}

func (s *_amazonBedrockTaskSettings) AmazonBedrockTaskSettingsCaster() *types.AmazonBedrockTaskSettings {
	return s.v
}
