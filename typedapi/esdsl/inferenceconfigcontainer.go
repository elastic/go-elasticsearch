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
// https://github.com/elastic/elasticsearch-specification/tree/beeb1dc688bcc058488dcc45d9cbd2cd364e9943

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _inferenceConfigContainer struct {
	v *types.InferenceConfigContainer
}

func NewInferenceConfigContainer() *_inferenceConfigContainer {
	return &_inferenceConfigContainer{v: types.NewInferenceConfigContainer()}
}

// AdditionalInferenceConfigContainerProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_inferenceConfigContainer) AdditionalInferenceConfigContainerProperty(key string, value json.RawMessage) *_inferenceConfigContainer {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalInferenceConfigContainerProperty = tmp
	return s
}

func (s *_inferenceConfigContainer) Classification(classification types.ClassificationInferenceOptionsVariant) *_inferenceConfigContainer {

	s.v.Classification = classification.ClassificationInferenceOptionsCaster()

	return s
}

func (s *_inferenceConfigContainer) Regression(regression types.RegressionInferenceOptionsVariant) *_inferenceConfigContainer {

	s.v.Regression = regression.RegressionInferenceOptionsCaster()

	return s
}

func (s *_inferenceConfigContainer) InferenceConfigContainerCaster() *types.InferenceConfigContainer {
	return s.v
}
