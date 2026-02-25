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

type _patternCaptureTokenFilter struct {
	v *types.PatternCaptureTokenFilter
}

func NewPatternCaptureTokenFilter() *_patternCaptureTokenFilter {

	return &_patternCaptureTokenFilter{v: types.NewPatternCaptureTokenFilter()}

}

func (s *_patternCaptureTokenFilter) Patterns(patterns ...string) *_patternCaptureTokenFilter {

	for _, v := range patterns {

		s.v.Patterns = append(s.v.Patterns, v)

	}
	return s
}

func (s *_patternCaptureTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_patternCaptureTokenFilter {

	s.v.PreserveOriginal = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_patternCaptureTokenFilter) Version(versionstring string) *_patternCaptureTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_patternCaptureTokenFilter) PatternCaptureTokenFilterCaster() *types.PatternCaptureTokenFilter {
	return s.v
}
