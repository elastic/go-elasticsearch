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
// https://github.com/elastic/elasticsearch-specification/tree/d82ef79f6af3e5ddb412e64fc4477ca1833d4a27

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticencoder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticlanguage"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticnametype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/phoneticruletype"
)

type _phoneticTokenFilter struct {
	v *types.PhoneticTokenFilter
}

func NewPhoneticTokenFilter(encoder phoneticencoder.PhoneticEncoder) *_phoneticTokenFilter {

	tmp := &_phoneticTokenFilter{v: types.NewPhoneticTokenFilter()}

	tmp.Encoder(encoder)

	return tmp

}

func (s *_phoneticTokenFilter) Encoder(encoder phoneticencoder.PhoneticEncoder) *_phoneticTokenFilter {

	s.v.Encoder = encoder
	return s
}

func (s *_phoneticTokenFilter) Languageset(languagesets ...phoneticlanguage.PhoneticLanguage) *_phoneticTokenFilter {

	s.v.Languageset = make([]phoneticlanguage.PhoneticLanguage, len(languagesets))
	s.v.Languageset = languagesets

	return s
}

func (s *_phoneticTokenFilter) MaxCodeLen(maxcodelen int) *_phoneticTokenFilter {

	s.v.MaxCodeLen = &maxcodelen

	return s
}

func (s *_phoneticTokenFilter) NameType(nametype phoneticnametype.PhoneticNameType) *_phoneticTokenFilter {

	s.v.NameType = &nametype
	return s
}

func (s *_phoneticTokenFilter) Replace(replace bool) *_phoneticTokenFilter {

	s.v.Replace = &replace

	return s
}

func (s *_phoneticTokenFilter) RuleType(ruletype phoneticruletype.PhoneticRuleType) *_phoneticTokenFilter {

	s.v.RuleType = &ruletype
	return s
}

func (s *_phoneticTokenFilter) Version(versionstring string) *_phoneticTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_phoneticTokenFilter) PhoneticTokenFilterCaster() *types.PhoneticTokenFilter {
	return s.v
}
