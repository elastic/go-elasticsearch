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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _patternReplaceTokenFilter struct {
	v *types.PatternReplaceTokenFilter
}

func NewPatternReplaceTokenFilter(pattern string) *_patternReplaceTokenFilter {

	tmp := &_patternReplaceTokenFilter{v: types.NewPatternReplaceTokenFilter()}

	tmp.Pattern(pattern)

	return tmp

}

func (s *_patternReplaceTokenFilter) All(all bool) *_patternReplaceTokenFilter {

	s.v.All = &all

	return s
}

func (s *_patternReplaceTokenFilter) Flags(flags string) *_patternReplaceTokenFilter {

	s.v.Flags = &flags

	return s
}

func (s *_patternReplaceTokenFilter) Pattern(pattern string) *_patternReplaceTokenFilter {

	s.v.Pattern = pattern

	return s
}

func (s *_patternReplaceTokenFilter) Replacement(replacement string) *_patternReplaceTokenFilter {

	s.v.Replacement = &replacement

	return s
}

func (s *_patternReplaceTokenFilter) Version(versionstring string) *_patternReplaceTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_patternReplaceTokenFilter) PatternReplaceTokenFilterCaster() *types.PatternReplaceTokenFilter {
	return s.v
}
