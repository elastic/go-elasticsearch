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

type _fillMaskInferenceUpdateOptions struct {
	v *types.FillMaskInferenceUpdateOptions
}

// Fill mask configuration for inference.
func NewFillMaskInferenceUpdateOptions() *_fillMaskInferenceUpdateOptions {

	return &_fillMaskInferenceUpdateOptions{v: types.NewFillMaskInferenceUpdateOptions()}

}

func (s *_fillMaskInferenceUpdateOptions) NumTopClasses(numtopclasses int) *_fillMaskInferenceUpdateOptions {

	s.v.NumTopClasses = &numtopclasses

	return s
}

func (s *_fillMaskInferenceUpdateOptions) ResultsField(resultsfield string) *_fillMaskInferenceUpdateOptions {

	s.v.ResultsField = &resultsfield

	return s
}

func (s *_fillMaskInferenceUpdateOptions) Tokenization(tokenization types.NlpTokenizationUpdateOptionsVariant) *_fillMaskInferenceUpdateOptions {

	s.v.Tokenization = tokenization.NlpTokenizationUpdateOptionsCaster()

	return s
}

func (s *_fillMaskInferenceUpdateOptions) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	container := types.NewInferenceConfigUpdateContainer()

	container.FillMask = s.v

	return container
}

func (s *_fillMaskInferenceUpdateOptions) FillMaskInferenceUpdateOptionsCaster() *types.FillMaskInferenceUpdateOptions {
	return s.v
}
