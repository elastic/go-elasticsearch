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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenchar"
)

type _nGramTokenizer struct {
	v *types.NGramTokenizer
}

func NewNGramTokenizer() *_nGramTokenizer {

	return &_nGramTokenizer{v: types.NewNGramTokenizer()}

}

func (s *_nGramTokenizer) CustomTokenChars(customtokenchars string) *_nGramTokenizer {

	s.v.CustomTokenChars = &customtokenchars

	return s
}

func (s *_nGramTokenizer) MaxGram(maxgram int) *_nGramTokenizer {

	s.v.MaxGram = &maxgram

	return s
}

func (s *_nGramTokenizer) MinGram(mingram int) *_nGramTokenizer {

	s.v.MinGram = &mingram

	return s
}

func (s *_nGramTokenizer) TokenChars(tokenchars ...tokenchar.TokenChar) *_nGramTokenizer {

	for _, v := range tokenchars {

		s.v.TokenChars = append(s.v.TokenChars, v)

	}
	return s
}

func (s *_nGramTokenizer) Version(versionstring string) *_nGramTokenizer {

	s.v.Version = &versionstring

	return s
}

func (s *_nGramTokenizer) NGramTokenizerCaster() *types.NGramTokenizer {
	return s.v
}
