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
// https://github.com/elastic/elasticsearch-specification/tree/55f8d05b44cea956ae5ceddfcb02770ea2a24ff6

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/matchtype"
)

type _dynamicTemplate struct {
	v *types.DynamicTemplate
}

func NewDynamicTemplate() *_dynamicTemplate {
	return &_dynamicTemplate{v: types.NewDynamicTemplate()}
}

func (s *_dynamicTemplate) Mapping(property types.PropertyVariant) *_dynamicTemplate {

	s.v.Mapping = *property.PropertyCaster()

	return s
}

func (s *_dynamicTemplate) Match(matches ...string) *_dynamicTemplate {

	s.v.Match = make([]string, len(matches))
	s.v.Match = matches

	return s
}

func (s *_dynamicTemplate) MatchMappingType(matchmappingtypes ...string) *_dynamicTemplate {

	s.v.MatchMappingType = make([]string, len(matchmappingtypes))
	s.v.MatchMappingType = matchmappingtypes

	return s
}

func (s *_dynamicTemplate) MatchPattern(matchpattern matchtype.MatchType) *_dynamicTemplate {

	s.v.MatchPattern = &matchpattern
	return s
}

func (s *_dynamicTemplate) PathMatch(pathmatches ...string) *_dynamicTemplate {

	s.v.PathMatch = make([]string, len(pathmatches))
	s.v.PathMatch = pathmatches

	return s
}

func (s *_dynamicTemplate) PathUnmatch(pathunmatches ...string) *_dynamicTemplate {

	s.v.PathUnmatch = make([]string, len(pathunmatches))
	s.v.PathUnmatch = pathunmatches

	return s
}

func (s *_dynamicTemplate) Runtime(runtime types.RuntimeFieldVariant) *_dynamicTemplate {

	s.v.Runtime = runtime.RuntimeFieldCaster()

	return s
}

func (s *_dynamicTemplate) Unmatch(unmatches ...string) *_dynamicTemplate {

	s.v.Unmatch = make([]string, len(unmatches))
	s.v.Unmatch = unmatches

	return s
}

func (s *_dynamicTemplate) UnmatchMappingType(unmatchmappingtypes ...string) *_dynamicTemplate {

	s.v.UnmatchMappingType = make([]string, len(unmatchmappingtypes))
	s.v.UnmatchMappingType = unmatchmappingtypes

	return s
}

func (s *_dynamicTemplate) DynamicTemplateCaster() *types.DynamicTemplate {
	return s.v
}
