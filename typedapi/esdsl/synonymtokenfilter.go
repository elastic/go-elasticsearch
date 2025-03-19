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
// https://github.com/elastic/elasticsearch-specification/tree/ea991724f4dd4f90c496eff547d3cc2e6529f509

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/synonymformat"
)

type _synonymTokenFilter struct {
	v *types.SynonymTokenFilter
}

func NewSynonymTokenFilter() *_synonymTokenFilter {

	return &_synonymTokenFilter{v: types.NewSynonymTokenFilter()}

}

func (s *_synonymTokenFilter) Expand(expand bool) *_synonymTokenFilter {

	s.v.Expand = &expand

	return s
}

func (s *_synonymTokenFilter) Format(format synonymformat.SynonymFormat) *_synonymTokenFilter {

	s.v.Format = &format
	return s
}

func (s *_synonymTokenFilter) Lenient(lenient bool) *_synonymTokenFilter {

	s.v.Lenient = &lenient

	return s
}

func (s *_synonymTokenFilter) Synonyms(synonyms ...string) *_synonymTokenFilter {

	for _, v := range synonyms {

		s.v.Synonyms = append(s.v.Synonyms, v)

	}
	return s
}

func (s *_synonymTokenFilter) SynonymsPath(synonymspath string) *_synonymTokenFilter {

	s.v.SynonymsPath = &synonymspath

	return s
}

func (s *_synonymTokenFilter) SynonymsSet(synonymsset string) *_synonymTokenFilter {

	s.v.SynonymsSet = &synonymsset

	return s
}

func (s *_synonymTokenFilter) Tokenizer(tokenizer string) *_synonymTokenFilter {

	s.v.Tokenizer = &tokenizer

	return s
}

func (s *_synonymTokenFilter) Updateable(updateable bool) *_synonymTokenFilter {

	s.v.Updateable = &updateable

	return s
}

func (s *_synonymTokenFilter) Version(versionstring string) *_synonymTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_synonymTokenFilter) SynonymTokenFilterCaster() *types.SynonymTokenFilter {
	return s.v
}
