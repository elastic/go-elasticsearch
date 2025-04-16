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

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _completionSuggester struct {
	v *types.CompletionSuggester
}

// Provides auto-complete/search-as-you-type functionality.
func NewCompletionSuggester() *_completionSuggester {

	return &_completionSuggester{v: types.NewCompletionSuggester()}

}

func (s *_completionSuggester) Analyzer(analyzer string) *_completionSuggester {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_completionSuggester) Contexts(contexts map[string][]types.CompletionContext) *_completionSuggester {

	s.v.Contexts = contexts
	return s
}

func (s *_completionSuggester) Field(field string) *_completionSuggester {

	s.v.Field = field

	return s
}

func (s *_completionSuggester) Fuzzy(fuzzy types.SuggestFuzzinessVariant) *_completionSuggester {

	s.v.Fuzzy = fuzzy.SuggestFuzzinessCaster()

	return s
}

func (s *_completionSuggester) Regex(regex types.RegexOptionsVariant) *_completionSuggester {

	s.v.Regex = regex.RegexOptionsCaster()

	return s
}

func (s *_completionSuggester) Size(size int) *_completionSuggester {

	s.v.Size = &size

	return s
}

func (s *_completionSuggester) SkipDuplicates(skipduplicates bool) *_completionSuggester {

	s.v.SkipDuplicates = &skipduplicates

	return s
}

func (s *_completionSuggester) FieldSuggesterCaster() *types.FieldSuggester {
	container := types.NewFieldSuggester()

	container.Completion = s.v

	return container
}

func (s *_completionSuggester) CompletionSuggesterCaster() *types.CompletionSuggester {
	return s.v
}
