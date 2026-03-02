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
// https://github.com/elastic/elasticsearch-specification/tree/e196f9953fa743572ee46884835f1934bce9a16b

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _keepWordsTokenFilter struct {
	v *types.KeepWordsTokenFilter
}

func NewKeepWordsTokenFilter() *_keepWordsTokenFilter {

	return &_keepWordsTokenFilter{v: types.NewKeepWordsTokenFilter()}

}

func (s *_keepWordsTokenFilter) KeepWords(keepwords ...string) *_keepWordsTokenFilter {

	for _, v := range keepwords {

		s.v.KeepWords = append(s.v.KeepWords, v)

	}
	return s
}

func (s *_keepWordsTokenFilter) KeepWordsCase(keepwordscase bool) *_keepWordsTokenFilter {

	s.v.KeepWordsCase = &keepwordscase

	return s
}

func (s *_keepWordsTokenFilter) KeepWordsPath(keepwordspath string) *_keepWordsTokenFilter {

	s.v.KeepWordsPath = &keepwordspath

	return s
}

func (s *_keepWordsTokenFilter) Version(versionstring string) *_keepWordsTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_keepWordsTokenFilter) KeepWordsTokenFilterCaster() *types.KeepWordsTokenFilter {
	return s.v
}
