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

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _completionSuggester struct {
	v *types.CompletionSuggester
}

// Provides auto-complete/search-as-you-type functionality.
func NewCompletionSuggester() *_completionSuggester {

	return &_completionSuggester{v: types.NewCompletionSuggester()}

}

// The analyzer to analyze the suggest text with.
// Defaults to the search analyzer of the suggest field.
func (s *_completionSuggester) Analyzer(analyzer string) *_completionSuggester {

	s.v.Analyzer = &analyzer

	return s
}

// A value, geo point object, or a geo hash string to filter or boost the
// suggestion on.
func (s *_completionSuggester) Contexts(contexts map[string][]types.CompletionContext) *_completionSuggester {

	s.v.Contexts = contexts
	return s
}

// The field to fetch the candidate suggestions from.
// Needs to be set globally or per suggestion.
func (s *_completionSuggester) Field(field string) *_completionSuggester {

	s.v.Field = field

	return s
}

// Enables fuzziness, meaning you can have a typo in your search and still get
// results back.
func (s *_completionSuggester) Fuzzy(fuzzy types.SuggestFuzzinessVariant) *_completionSuggester {

	s.v.Fuzzy = fuzzy.SuggestFuzzinessCaster()

	return s
}

// A regex query that expresses a prefix as a regular expression.
func (s *_completionSuggester) Regex(regex types.RegexOptionsVariant) *_completionSuggester {

	s.v.Regex = regex.RegexOptionsCaster()

	return s
}

// The maximum corrections to be returned per suggest text token.
func (s *_completionSuggester) Size(size int) *_completionSuggester {

	s.v.Size = &size

	return s
}

// Whether duplicate suggestions should be filtered out.
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
