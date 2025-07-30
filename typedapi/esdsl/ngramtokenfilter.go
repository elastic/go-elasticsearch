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
// https://github.com/elastic/elasticsearch-specification/tree/de4ff9ec1f716256f521d9e30011ad9c284b0dcc

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _nGramTokenFilter struct {
	v *types.NGramTokenFilter
}

func NewNGramTokenFilter() *_nGramTokenFilter {

	return &_nGramTokenFilter{v: types.NewNGramTokenFilter()}

}

func (s *_nGramTokenFilter) MaxGram(maxgram int) *_nGramTokenFilter {

	s.v.MaxGram = &maxgram

	return s
}

func (s *_nGramTokenFilter) MinGram(mingram int) *_nGramTokenFilter {

	s.v.MinGram = &mingram

	return s
}

func (s *_nGramTokenFilter) PreserveOriginal(stringifiedboolean types.StringifiedbooleanVariant) *_nGramTokenFilter {

	s.v.PreserveOriginal = *stringifiedboolean.StringifiedbooleanCaster()

	return s
}

func (s *_nGramTokenFilter) Version(versionstring string) *_nGramTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_nGramTokenFilter) NGramTokenFilterCaster() *types.NGramTokenFilter {
	return s.v
}
