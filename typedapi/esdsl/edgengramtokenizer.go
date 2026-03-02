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

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenchar"
)

type _edgeNGramTokenizer struct {
	v *types.EdgeNGramTokenizer
}

func NewEdgeNGramTokenizer() *_edgeNGramTokenizer {

	return &_edgeNGramTokenizer{v: types.NewEdgeNGramTokenizer()}

}

func (s *_edgeNGramTokenizer) CustomTokenChars(customtokenchars string) *_edgeNGramTokenizer {

	s.v.CustomTokenChars = &customtokenchars

	return s
}

func (s *_edgeNGramTokenizer) MaxGram(maxgram int) *_edgeNGramTokenizer {

	s.v.MaxGram = &maxgram

	return s
}

func (s *_edgeNGramTokenizer) MinGram(mingram int) *_edgeNGramTokenizer {

	s.v.MinGram = &mingram

	return s
}

func (s *_edgeNGramTokenizer) TokenChars(tokenchars ...tokenchar.TokenChar) *_edgeNGramTokenizer {

	for _, v := range tokenchars {

		s.v.TokenChars = append(s.v.TokenChars, v)

	}
	return s
}

func (s *_edgeNGramTokenizer) Version(versionstring string) *_edgeNGramTokenizer {

	s.v.Version = &versionstring

	return s
}

func (s *_edgeNGramTokenizer) EdgeNGramTokenizerCaster() *types.EdgeNGramTokenizer {
	return s.v
}
