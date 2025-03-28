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
// https://github.com/elastic/elasticsearch-specification/tree/cd5cc9962e79198ac2daf9110c00808293977f13

package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
)

type _fieldSuggester struct {
	v *types.FieldSuggester
}

func NewFieldSuggester() *_fieldSuggester {
	return &_fieldSuggester{v: types.NewFieldSuggester()}
}

// AdditionalFieldSuggesterProperty is a single key dictionnary.
// It will replace the current value on each call.
func (s *_fieldSuggester) AdditionalFieldSuggesterProperty(key string, value json.RawMessage) *_fieldSuggester {

	tmp := make(map[string]json.RawMessage)

	tmp[key] = value

	s.v.AdditionalFieldSuggesterProperty = tmp
	return s
}

// Provides auto-complete/search-as-you-type functionality.
func (s *_fieldSuggester) Completion(completion types.CompletionSuggesterVariant) *_fieldSuggester {

	s.v.Completion = completion.CompletionSuggesterCaster()

	return s
}

// Provides access to word alternatives on a per token basis within a certain
// string distance.
func (s *_fieldSuggester) Phrase(phrase types.PhraseSuggesterVariant) *_fieldSuggester {

	s.v.Phrase = phrase.PhraseSuggesterCaster()

	return s
}

// Prefix used to search for suggestions.
func (s *_fieldSuggester) Prefix(prefix string) *_fieldSuggester {

	s.v.Prefix = &prefix

	return s
}

// A prefix expressed as a regular expression.
func (s *_fieldSuggester) Regex(regex string) *_fieldSuggester {

	s.v.Regex = &regex

	return s
}

// Suggests terms based on edit distance.
func (s *_fieldSuggester) Term(term types.TermSuggesterVariant) *_fieldSuggester {

	s.v.Term = term.TermSuggesterCaster()

	return s
}

// The text to use as input for the suggester.
// Needs to be set globally or per suggestion.
func (s *_fieldSuggester) Text(text string) *_fieldSuggester {

	s.v.Text = &text

	return s
}

func (s *_fieldSuggester) FieldSuggesterCaster() *types.FieldSuggester {
	return s.v
}
