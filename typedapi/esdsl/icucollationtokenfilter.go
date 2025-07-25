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
// https://github.com/elastic/elasticsearch-specification/tree/a0b0db20330063a6d11f7997ff443fd2a1a827d1

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationalternate"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationcasefirst"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationdecomposition"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icucollationstrength"
)

type _icuCollationTokenFilter struct {
	v *types.IcuCollationTokenFilter
}

func NewIcuCollationTokenFilter() *_icuCollationTokenFilter {

	return &_icuCollationTokenFilter{v: types.NewIcuCollationTokenFilter()}

}

func (s *_icuCollationTokenFilter) Alternate(alternate icucollationalternate.IcuCollationAlternate) *_icuCollationTokenFilter {

	s.v.Alternate = &alternate
	return s
}

func (s *_icuCollationTokenFilter) CaseFirst(casefirst icucollationcasefirst.IcuCollationCaseFirst) *_icuCollationTokenFilter {

	s.v.CaseFirst = &casefirst
	return s
}

func (s *_icuCollationTokenFilter) CaseLevel(caselevel bool) *_icuCollationTokenFilter {

	s.v.CaseLevel = &caselevel

	return s
}

func (s *_icuCollationTokenFilter) Country(country string) *_icuCollationTokenFilter {

	s.v.Country = &country

	return s
}

func (s *_icuCollationTokenFilter) Decomposition(decomposition icucollationdecomposition.IcuCollationDecomposition) *_icuCollationTokenFilter {

	s.v.Decomposition = &decomposition
	return s
}

func (s *_icuCollationTokenFilter) HiraganaQuaternaryMode(hiraganaquaternarymode bool) *_icuCollationTokenFilter {

	s.v.HiraganaQuaternaryMode = &hiraganaquaternarymode

	return s
}

func (s *_icuCollationTokenFilter) Language(language string) *_icuCollationTokenFilter {

	s.v.Language = &language

	return s
}

func (s *_icuCollationTokenFilter) Numeric(numeric bool) *_icuCollationTokenFilter {

	s.v.Numeric = &numeric

	return s
}

func (s *_icuCollationTokenFilter) Rules(rules string) *_icuCollationTokenFilter {

	s.v.Rules = &rules

	return s
}

func (s *_icuCollationTokenFilter) Strength(strength icucollationstrength.IcuCollationStrength) *_icuCollationTokenFilter {

	s.v.Strength = &strength
	return s
}

func (s *_icuCollationTokenFilter) VariableTop(variabletop string) *_icuCollationTokenFilter {

	s.v.VariableTop = &variabletop

	return s
}

func (s *_icuCollationTokenFilter) Variant(variant string) *_icuCollationTokenFilter {

	s.v.Variant = &variant

	return s
}

func (s *_icuCollationTokenFilter) Version(versionstring string) *_icuCollationTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_icuCollationTokenFilter) IcuCollationTokenFilterCaster() *types.IcuCollationTokenFilter {
	return s.v
}
