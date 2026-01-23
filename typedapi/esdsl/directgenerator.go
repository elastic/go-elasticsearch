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
// https://github.com/elastic/elasticsearch-specification/tree/b1811e10a0722431d79d1c234dd412ff47d8656f

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestmode"
)

type _directGenerator struct {
	v *types.DirectGenerator
}

func NewDirectGenerator() *_directGenerator {

	return &_directGenerator{v: types.NewDirectGenerator()}

}

func (s *_directGenerator) Field(field string) *_directGenerator {

	s.v.Field = field

	return s
}

func (s *_directGenerator) MaxEdits(maxedits int) *_directGenerator {

	s.v.MaxEdits = &maxedits

	return s
}

func (s *_directGenerator) MaxInspections(maxinspections float32) *_directGenerator {

	s.v.MaxInspections = &maxinspections

	return s
}

func (s *_directGenerator) MaxTermFreq(maxtermfreq float32) *_directGenerator {

	s.v.MaxTermFreq = &maxtermfreq

	return s
}

func (s *_directGenerator) MinDocFreq(mindocfreq float32) *_directGenerator {

	s.v.MinDocFreq = &mindocfreq

	return s
}

func (s *_directGenerator) MinWordLength(minwordlength int) *_directGenerator {

	s.v.MinWordLength = &minwordlength

	return s
}

func (s *_directGenerator) PostFilter(postfilter string) *_directGenerator {

	s.v.PostFilter = &postfilter

	return s
}

func (s *_directGenerator) PreFilter(prefilter string) *_directGenerator {

	s.v.PreFilter = &prefilter

	return s
}

func (s *_directGenerator) PrefixLength(prefixlength int) *_directGenerator {

	s.v.PrefixLength = &prefixlength

	return s
}

func (s *_directGenerator) Size(size int) *_directGenerator {

	s.v.Size = &size

	return s
}

func (s *_directGenerator) SuggestMode(suggestmode suggestmode.SuggestMode) *_directGenerator {

	s.v.SuggestMode = &suggestmode
	return s
}

func (s *_directGenerator) DirectGeneratorCaster() *types.DirectGenerator {
	return s.v
}
