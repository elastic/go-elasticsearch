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
// https://github.com/elastic/elasticsearch-specification/tree/86f41834c7bb975159a38a73be8a9d930010d673

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsSetQuery struct {
	k string
	v *types.TermsSetQuery
}

// Returns documents that contain a minimum number of exact terms in a provided
// field.
// To return a document, a required number of terms must exactly match the field
// values, including whitespace and capitalization.
func NewTermsSetQuery(key string) *_termsSetQuery {
	return &_termsSetQuery{
		k: key,
		v: types.NewTermsSetQuery(),
	}
}

func (s *_termsSetQuery) Boost(boost float32) *_termsSetQuery {

	s.v.Boost = &boost

	return s
}

func (s *_termsSetQuery) MinimumShouldMatch(minimumshouldmatch types.MinimumShouldMatchVariant) *_termsSetQuery {

	s.v.MinimumShouldMatch = *minimumshouldmatch.MinimumShouldMatchCaster()

	return s
}

func (s *_termsSetQuery) MinimumShouldMatchField(field string) *_termsSetQuery {

	s.v.MinimumShouldMatchField = &field

	return s
}

func (s *_termsSetQuery) MinimumShouldMatchScript(minimumshouldmatchscript types.ScriptVariant) *_termsSetQuery {

	s.v.MinimumShouldMatchScript = minimumshouldmatchscript.ScriptCaster()

	return s
}

func (s *_termsSetQuery) QueryName_(queryname_ string) *_termsSetQuery {

	s.v.QueryName_ = &queryname_

	return s
}

func (s *_termsSetQuery) Terms(terms ...types.FieldValueVariant) *_termsSetQuery {

	for _, v := range terms {

		s.v.Terms = append(s.v.Terms, *v.FieldValueCaster())

	}
	return s
}

func (s *_termsSetQuery) QueryCaster() *types.Query {
	container := types.NewQuery()
	container.TermsSet = map[string]types.TermsSetQuery{
		s.k: *s.v,
	}
	return container
}

// NewSingleTermsSetQuery should be used when you want to
// create a single key dictionary without specifying the key in the
// constructor. Usually key is already defined within the parent container.
func NewSingleTermsSetQuery() *_termsSetQuery {
	return &_termsSetQuery{
		k: "",
		v: types.NewTermsSetQuery(),
	}
}

func (s *_termsSetQuery) TermsSetQueryCaster() *types.TermsSetQuery {
	return s.v.TermsSetQueryCaster()
}
