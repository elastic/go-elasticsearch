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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _pathHierarchyTokenizer struct {
	v *types.PathHierarchyTokenizer
}

func NewPathHierarchyTokenizer() *_pathHierarchyTokenizer {

	return &_pathHierarchyTokenizer{v: types.NewPathHierarchyTokenizer()}

}

func (s *_pathHierarchyTokenizer) BufferSize(stringifiedinteger types.StringifiedintegerVariant) *_pathHierarchyTokenizer {

	s.v.BufferSize = *stringifiedinteger.StringifiedintegerCaster()

	return s
}

func (s *_pathHierarchyTokenizer) Delimiter(delimiter string) *_pathHierarchyTokenizer {

	s.v.Delimiter = &delimiter

	return s
}

func (s *_pathHierarchyTokenizer) Replacement(replacement string) *_pathHierarchyTokenizer {

	s.v.Replacement = &replacement

	return s
}

func (s *_pathHierarchyTokenizer) Reverse(stringifiedboolean types.StringifiedbooleanVariant) *_pathHierarchyTokenizer {

	s.v.Reverse = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_pathHierarchyTokenizer) Skip(stringifiedinteger types.StringifiedintegerVariant) *_pathHierarchyTokenizer {

	s.v.Skip = *stringifiedinteger.StringifiedintegerCaster()

	return s
}

func (s *_pathHierarchyTokenizer) Version(versionstring string) *_pathHierarchyTokenizer {

	s.v.Version = &versionstring

	return s
}

func (s *_pathHierarchyTokenizer) PathHierarchyTokenizerCaster() *types.PathHierarchyTokenizer {
	return s.v
}
