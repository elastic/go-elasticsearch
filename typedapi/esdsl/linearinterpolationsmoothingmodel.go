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

type _linearInterpolationSmoothingModel struct {
	v *types.LinearInterpolationSmoothingModel
}

// A smoothing model that takes the weighted mean of the unigrams, bigrams, and
// trigrams based on user supplied weights (lambdas).
func NewLinearInterpolationSmoothingModel(bigramlambda types.Float64, trigramlambda types.Float64, unigramlambda types.Float64) *_linearInterpolationSmoothingModel {

	tmp := &_linearInterpolationSmoothingModel{v: types.NewLinearInterpolationSmoothingModel()}

	tmp.BigramLambda(bigramlambda)

	tmp.TrigramLambda(trigramlambda)

	tmp.UnigramLambda(unigramlambda)

	return tmp

}

func (s *_linearInterpolationSmoothingModel) BigramLambda(bigramlambda types.Float64) *_linearInterpolationSmoothingModel {

	s.v.BigramLambda = bigramlambda

	return s
}

func (s *_linearInterpolationSmoothingModel) TrigramLambda(trigramlambda types.Float64) *_linearInterpolationSmoothingModel {

	s.v.TrigramLambda = trigramlambda

	return s
}

func (s *_linearInterpolationSmoothingModel) UnigramLambda(unigramlambda types.Float64) *_linearInterpolationSmoothingModel {

	s.v.UnigramLambda = unigramlambda

	return s
}

func (s *_linearInterpolationSmoothingModel) SmoothingModelContainerCaster() *types.SmoothingModelContainer {
	container := types.NewSmoothingModelContainer()

	container.LinearInterpolation = s.v

	return container
}

func (s *_linearInterpolationSmoothingModel) LinearInterpolationSmoothingModelCaster() *types.LinearInterpolationSmoothingModel {
	return s.v
}
