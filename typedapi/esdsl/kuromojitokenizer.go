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
// https://github.com/elastic/elasticsearch-specification/tree/e585438d116b00ff34643179e6286e402c0bcaaf

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/kuromojitokenizationmode"
)

type _kuromojiTokenizer struct {
	v *types.KuromojiTokenizer
}

func NewKuromojiTokenizer(mode kuromojitokenizationmode.KuromojiTokenizationMode) *_kuromojiTokenizer {

	tmp := &_kuromojiTokenizer{v: types.NewKuromojiTokenizer()}

	tmp.Mode(mode)

	return tmp

}

func (s *_kuromojiTokenizer) DiscardCompoundToken(discardcompoundtoken bool) *_kuromojiTokenizer {

	s.v.DiscardCompoundToken = &discardcompoundtoken

	return s
}

func (s *_kuromojiTokenizer) DiscardPunctuation(discardpunctuation bool) *_kuromojiTokenizer {

	s.v.DiscardPunctuation = &discardpunctuation

	return s
}

func (s *_kuromojiTokenizer) Mode(mode kuromojitokenizationmode.KuromojiTokenizationMode) *_kuromojiTokenizer {

	s.v.Mode = mode
	return s
}

func (s *_kuromojiTokenizer) NbestCost(nbestcost int) *_kuromojiTokenizer {

	s.v.NbestCost = &nbestcost

	return s
}

func (s *_kuromojiTokenizer) NbestExamples(nbestexamples string) *_kuromojiTokenizer {

	s.v.NbestExamples = &nbestexamples

	return s
}

func (s *_kuromojiTokenizer) UserDictionary(userdictionary string) *_kuromojiTokenizer {

	s.v.UserDictionary = &userdictionary

	return s
}

func (s *_kuromojiTokenizer) UserDictionaryRules(userdictionaryrules ...string) *_kuromojiTokenizer {

	for _, v := range userdictionaryrules {

		s.v.UserDictionaryRules = append(s.v.UserDictionaryRules, v)

	}
	return s
}

func (s *_kuromojiTokenizer) Version(versionstring string) *_kuromojiTokenizer {

	s.v.Version = &versionstring

	return s
}

func (s *_kuromojiTokenizer) KuromojiTokenizerCaster() *types.KuromojiTokenizer {
	return s.v
}
