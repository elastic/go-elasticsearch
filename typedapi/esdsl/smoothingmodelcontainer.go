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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _smoothingModelContainer struct {
	v *types.SmoothingModelContainer
}

func NewSmoothingModelContainer() *_smoothingModelContainer {
	return &_smoothingModelContainer{v: types.NewSmoothingModelContainer()}
}

// AdditionalSmoothingModelContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_smoothingModelContainer) AdditionalSmoothingModelContainerProperty(key string, value json.RawMessage) *_smoothingModelContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalSmoothingModelContainerProperty = tmp
	return s
}

// A smoothing model that uses an additive smoothing where a constant (typically
// `1.0` or smaller) is added to all counts to balance weights.
func (s *_smoothingModelContainer) Laplace(laplace types.LaplaceSmoothingModelVariant) *_smoothingModelContainer {

	s.v.Laplace = laplace.LaplaceSmoothingModelCaster()

	return s
}

// A smoothing model that takes the weighted mean of the unigrams, bigrams, and
// trigrams based on user supplied weights (lambdas).
func (s *_smoothingModelContainer) LinearInterpolation(linearinterpolation types.LinearInterpolationSmoothingModelVariant) *_smoothingModelContainer {

	s.v.LinearInterpolation = linearinterpolation.LinearInterpolationSmoothingModelCaster()

	return s
}

// A simple backoff model that backs off to lower order n-gram models if the
// higher order count is `0` and discounts the lower order n-gram model by a
// constant factor.
func (s *_smoothingModelContainer) StupidBackoff(stupidbackoff types.StupidBackoffSmoothingModelVariant) *_smoothingModelContainer {

	s.v.StupidBackoff = stupidbackoff.StupidBackoffSmoothingModelCaster()

	return s
}

func (s *_smoothingModelContainer) SmoothingModelContainerCaster() *types.SmoothingModelContainer {
	return s.v
}
