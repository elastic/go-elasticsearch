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

type _anthropicTaskSettings struct {
	v *types.AnthropicTaskSettings
}

func NewAnthropicTaskSettings(maxtokens int) *_anthropicTaskSettings {

	tmp := &_anthropicTaskSettings{v: types.NewAnthropicTaskSettings()}

	tmp.MaxTokens(maxtokens)

	return tmp

}

// For a `completion` task, it is the maximum number of tokens to generate
// before stopping.
func (s *_anthropicTaskSettings) MaxTokens(maxtokens int) *_anthropicTaskSettings {

	s.v.MaxTokens = maxtokens

	return s
}

// For a `completion` task, it is the amount of randomness injected into the
// response.
// For more details about the supported range, refer to Anthropic documentation.
func (s *_anthropicTaskSettings) Temperature(temperature float32) *_anthropicTaskSettings {

	s.v.Temperature = &temperature

	return s
}

// For a `completion` task, it specifies to only sample from the top K options
// for each subsequent token.
// It is recommended for advanced use cases only.
// You usually only need to use `temperature`.
func (s *_anthropicTaskSettings) TopK(topk int) *_anthropicTaskSettings {

	s.v.TopK = &topk

	return s
}

// For a `completion` task, it specifies to use Anthropic's nucleus sampling.
// In nucleus sampling, Anthropic computes the cumulative distribution over all
// the options for each subsequent token in decreasing probability order and
// cuts it off once it reaches the specified probability.
// You should either alter `temperature` or `top_p`, but not both.
// It is recommended for advanced use cases only.
// You usually only need to use `temperature`.
func (s *_anthropicTaskSettings) TopP(topp float32) *_anthropicTaskSettings {

	s.v.TopP = &topp

	return s
}

func (s *_anthropicTaskSettings) AnthropicTaskSettingsCaster() *types.AnthropicTaskSettings {
	return s.v
}
