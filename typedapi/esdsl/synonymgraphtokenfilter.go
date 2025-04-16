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
// https://github.com/elastic/elasticsearch-specification/tree/52c473efb1fb5320a5bac12572d0b285882862fb

package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/synonymformat"
)

type _synonymGraphTokenFilter struct {
	v *types.SynonymGraphTokenFilter
}

func NewSynonymGraphTokenFilter() *_synonymGraphTokenFilter {

	return &_synonymGraphTokenFilter{v: types.NewSynonymGraphTokenFilter()}

}

func (s *_synonymGraphTokenFilter) Expand(expand bool) *_synonymGraphTokenFilter {

	s.v.Expand = &expand

	return s
}

func (s *_synonymGraphTokenFilter) Format(format synonymformat.SynonymFormat) *_synonymGraphTokenFilter {

	s.v.Format = &format
	return s
}

func (s *_synonymGraphTokenFilter) Lenient(lenient bool) *_synonymGraphTokenFilter {

	s.v.Lenient = &lenient

	return s
}

func (s *_synonymGraphTokenFilter) Synonyms(synonyms ...string) *_synonymGraphTokenFilter {

	for _, v := range synonyms {

		s.v.Synonyms = append(s.v.Synonyms, v)

	}
	return s
}

func (s *_synonymGraphTokenFilter) SynonymsPath(synonymspath string) *_synonymGraphTokenFilter {

	s.v.SynonymsPath = &synonymspath

	return s
}

func (s *_synonymGraphTokenFilter) SynonymsSet(synonymsset string) *_synonymGraphTokenFilter {

	s.v.SynonymsSet = &synonymsset

	return s
}

func (s *_synonymGraphTokenFilter) Tokenizer(tokenizer string) *_synonymGraphTokenFilter {

	s.v.Tokenizer = &tokenizer

	return s
}

func (s *_synonymGraphTokenFilter) Updateable(updateable bool) *_synonymGraphTokenFilter {

	s.v.Updateable = &updateable

	return s
}

func (s *_synonymGraphTokenFilter) Version(versionstring string) *_synonymGraphTokenFilter {

	s.v.Version = &versionstring

	return s
}

func (s *_synonymGraphTokenFilter) SynonymGraphTokenFilterCaster() *types.SynonymGraphTokenFilter {
	return s.v
}
