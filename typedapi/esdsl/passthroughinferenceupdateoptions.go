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
// https://github.com/elastic/elasticsearch-specification/tree/bc885996c471cc7c2c7d51cba22aab19867672ac

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _passThroughInferenceUpdateOptions struct {
	v *types.PassThroughInferenceUpdateOptions
}

// Pass through configuration for inference.
func NewPassThroughInferenceUpdateOptions() *_passThroughInferenceUpdateOptions {

	return &_passThroughInferenceUpdateOptions{v: types.NewPassThroughInferenceUpdateOptions()}

}

func (s *_passThroughInferenceUpdateOptions) ResultsField(resultsfield string) *_passThroughInferenceUpdateOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_passThroughInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_passThroughInferenceUpdateOptions {

	s.v.Tokenization = tokenization.NlpTokenizationUpdateOptionsCaster()

	return s
}

func (s *_passThroughInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	container := types.NewInferenceConfigUpdateContainer()

	container.PassThrough = s.v

	return container
}

func (s *_passThroughInferenceUpdateOptions) PassThroughInferenceUpdateOptionsCaster() *types.PassThroughInferenceUpdateOptions {
	return s.v
}
