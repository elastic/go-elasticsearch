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
// https://github.com/elastic/elasticsearch-specification/tree/60a81659be928bfe6cec53708c7f7613555a5eaf

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

// This is provide all the types that are part of the union.
type _normalizer struct {
	v types.Normalizer
}

func NewNormalizer() *_normalizer {
	return &_normalizer{v: nil}
}

func (u *_normalizer) LowercaseNormalizer(lowercasenormalizer types.LowercaseNormalizerVariant) *_normalizer {

	u.v = &lowercasenormalizer

	return u
}

// Interface implementation for LowercaseNormalizer in Normalizer union
func (u *_lowercaseNormalizer) NormalizerCaster() *types.Normalizer {
	t := types.Normalizer(u.v)
	return &t
}

func (u *_normalizer) CustomNormalizer(customnormalizer types.CustomNormalizerVariant) *_normalizer {

	u.v = &customnormalizer

	return u
}

// Interface implementation for CustomNormalizer in Normalizer union
func (u *_customNormalizer) NormalizerCaster() *types.Normalizer {
	t := types.Normalizer(u.v)
	return &t
}

func (u *_normalizer) NormalizerCaster() *types.Normalizer {
	return &u.v
}
