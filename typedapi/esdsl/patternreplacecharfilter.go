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
// https://github.com/elastic/elasticsearch-specification/tree/c6ef5fbc736f1dd6256c2babc92e07bf150cadb9

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _patternReplaceCharFilter struct {
	v *types.PatternReplaceCharFilter
}

func NewPatternReplaceCharFilter(pattern string) *_patternReplaceCharFilter {

	tmp := &_patternReplaceCharFilter{v: types.NewPatternReplaceCharFilter()}

	tmp.Pattern(pattern)

	return tmp

}

func (s *_patternReplaceCharFilter) Flags(flags string) *_patternReplaceCharFilter {

	s.v.Flags = &flags

	return s
}

func (s *_patternReplaceCharFilter) Pattern(pattern string) *_patternReplaceCharFilter {

	s.v.Pattern = pattern

	return s
}

func (s *_patternReplaceCharFilter) Replacement(replacement string) *_patternReplaceCharFilter {

	s.v.Replacement = &replacement

	return s
}

func (s *_patternReplaceCharFilter) Version(versionstring string) *_patternReplaceCharFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_patternReplaceCharFilter) PatternReplaceCharFilterCaster() *types.PatternReplaceCharFilter {
	return s.v
}
