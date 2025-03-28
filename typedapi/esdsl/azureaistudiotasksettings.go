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

type _azureAiStudioTaskSettings struct {
	v *types.AzureAiStudioTaskSettings
}

func NewAzureAiStudioTaskSettings() *_azureAiStudioTaskSettings {

	return &_azureAiStudioTaskSettings{v: types.NewAzureAiStudioTaskSettings()}

}

// For a `completion` task, instruct the inference process to perform sampling.
// It has no effect unless `temperature` or `top_p` is specified.
func (s *_azureAiStudioTaskSettings) DoSample(dosample float32) *_azureAiStudioTaskSettings {

	s.v.DoSample = &dosample

	return s
}

// For a `completion` task, provide a hint for the maximum number of output
// tokens to be generated.
func (s *_azureAiStudioTaskSettings) MaxNewTokens(maxnewtokens int) *_azureAiStudioTaskSettings {

	s.v.MaxNewTokens = &maxnewtokens

	return s
}

// For a `completion` task, control the apparent creativity of generated
// completions with a sampling temperature.
// It must be a number in the range of 0.0 to 2.0.
// It should not be used if `top_p` is specified.
func (s *_azureAiStudioTaskSettings) Temperature(temperature float32) *_azureAiStudioTaskSettings {

	s.v.Temperature = &temperature

	return s
}

// For a `completion` task, make the model consider the results of the tokens
// with nucleus sampling probability.
// It is an alternative value to `temperature` and must be a number in the range
// of 0.0 to 2.0.
// It should not be used if `temperature` is specified.
func (s *_azureAiStudioTaskSettings) TopP(topp float32) *_azureAiStudioTaskSettings {

	s.v.TopP = &topp

	return s
}

// For a `text_embedding` task, specify the user issuing the request.
// This information can be used for abuse detection.
func (s *_azureAiStudioTaskSettings) User(user string) *_azureAiStudioTaskSettings {

	s.v.User = &user

	return s
}

func (s *_azureAiStudioTaskSettings) AzureAiStudioTaskSettingsCaster() *types.AzureAiStudioTaskSettings {
	return s.v
}
