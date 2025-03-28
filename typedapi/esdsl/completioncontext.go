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

type _completionContext struct {
	v *types.CompletionContext
}

func NewCompletionContext() *_completionContext {

	return &_completionContext{v: types.NewCompletionContext()}

}

// The factor by which the score of the suggestion should be boosted.
// The score is computed by multiplying the boost with the suggestion weight.
func (s *_completionContext) Boost(boost types.Float64) *_completionContext {

	s.v.Boost = &boost

	return s
}

// The value of the category to filter/boost on.
func (s *_completionContext) Context(context types.ContextVariant) *_completionContext {

	s.v.Context = *context.ContextCaster()

	return s
}

// An array of precision values at which neighboring geohashes should be taken
// into account.
// Precision value can be a distance value (`5m`, `10km`, etc.) or a raw geohash
// precision (`1`..`12`).
// Defaults to generating neighbors for index time precision level.
func (s *_completionContext) Neighbours(neighbours ...types.GeoHashPrecisionVariant) *_completionContext {

	for _, v := range neighbours {

		s.v.Neighbours = append(s.v.Neighbours, *v.GeoHashPrecisionCaster())

	}
	return s
}

// The precision of the geohash to encode the query geo point.
// Can be specified as a distance value (`5m`, `10km`, etc.), or as a raw
// geohash precision (`1`..`12`).
// Defaults to index time precision level.
func (s *_completionContext) Precision(geohashprecision types.GeoHashPrecisionVariant) *_completionContext {

	s.v.Precision = *geohashprecision.GeoHashPrecisionCaster()

	return s
}

// Whether the category value should be treated as a prefix or not.
func (s *_completionContext) Prefix(prefix bool) *_completionContext {

	s.v.Prefix = &prefix

	return s
}

func (s *_completionContext) CompletionContextCaster() *types.CompletionContext {
	return s.v
}
