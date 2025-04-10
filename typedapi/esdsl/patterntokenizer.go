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

type _patternTokenizer struct {
	v *types.PatternTokenizer
}

func NewPatternTokenizer() *_patternTokenizer {

	return &_patternTokenizer{v: types.NewPatternTokenizer()}

}

func (s *_patternTokenizer) Flags(flags string) *_patternTokenizer {

	s.v.Flags = &flags

	return s
}

func (s *_patternTokenizer) Group(group int) *_patternTokenizer {

	s.v.Group = &group

	return s
}

func (s *_patternTokenizer) Pattern(pattern string) *_patternTokenizer {

	s.v.Pattern = &pattern

	return s
}

func (s *_patternTokenizer) Version(versionstring string) *_patternTokenizer {

	s.v.Version = &versionstring

	return s
}

func (s *_patternTokenizer) PatternTokenizerCaster() *types.PatternTokenizer {
	return s.v
}
