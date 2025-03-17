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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _laplaceSmoothingModel struct {
	v *types.LaplaceSmoothingModel
}

// A smoothing model that uses an additive smoothing where a constant (typically
// `1.0` or smaller) is added to all counts to balance weights.
func NewLaplaceSmoothingModel(alpha types.Float64) *_laplaceSmoothingModel {

	tmp := &_laplaceSmoothingModel{v: types.NewLaplaceSmoothingModel()}

	tmp.Alpha(alpha)

	return tmp

}

// A constant that is added to all counts to balance weights.
func (s *_laplaceSmoothingModel) Alpha(alpha types.Float64) *_laplaceSmoothingModel {

	s.v.Alpha = alpha

	return s
}

func (s *_laplaceSmoothingModel) SmoothingModelContainerCaster() *types.SmoothingModelContainer {
	container := types.NewSmoothingModelContainer()

	container.Laplace = s.v

	return container
}

func (s *_laplaceSmoothingModel) LaplaceSmoothingModelCaster() *types.LaplaceSmoothingModel {
	return s.v
}
