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
// https://github.com/elastic/elasticsearch-specification/tree/2514615770f18dbb4e3887cc1a279995dbfd0724

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _completionContext struct {
	v *types.CompletionContext
}

func NewCompletionContext() *_completionContext {

	return &_completionContext{v: types.NewCompletionContext()}

}

func (s *_completionContext) Boost(boost types.Float64) *_completionContext {

	s.v.Boost = &boost

	return s
}

func (s *_completionContext) Context(context types.ContextVariant) *_completionContext {

	s.v.Context = *context.ContextCaster()

	return s
}

func (s *_completionContext) Neighbours(neighbours ...types.GeoHashPrecisionVariant) *_completionContext {

	for _, v := range neighbours {

		s.v.Neighbours = append(s.v.Neighbours, *v.GeoHashPrecisionCaster())

	}
	return s
}

func (s *_completionContext) Precision(geohashprecision types.GeoHashPrecisionVariant) *_completionContext {

	s.v.Precision = *geohashprecision.GeoHashPrecisionCaster()

	return s
}

func (s *_completionContext) Prefix(prefix bool) *_completionContext {

	s.v.Prefix = &prefix

	return s
}

func (s *_completionContext) CompletionContextCaster() *types.CompletionContext {
	return s.v
}
