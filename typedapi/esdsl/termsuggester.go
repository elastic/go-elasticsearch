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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/stringdistance"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestsort"
)

type _termSuggester struct {
	v *types.TermSuggester
}

// Suggests terms based on edit distance.
func NewTermSuggester() *_termSuggester {

	return &_termSuggester{v: types.NewTermSuggester()}

}

// The analyzer to analyze the suggest text with.
// Defaults to the search analyzer of the suggest field.
func (s *_termSuggester) Analyzer(analyzer string) *_termSuggester {

	s.v.Analyzer = &analyzer

	return s
}

// The field to fetch the candidate suggestions from.
// Needs to be set globally or per suggestion.
func (s *_termSuggester) Field(field string) *_termSuggester {

	s.v.Field = field

	return s
}

func (s *_termSuggester) LowercaseTerms(lowercaseterms bool) *_termSuggester {

	s.v.LowercaseTerms = &lowercaseterms

	return s
}

// The maximum edit distance candidate suggestions can have in order to be
// considered as a suggestion.
// Can only be `1` or `2`.
func (s *_termSuggester) MaxEdits(maxedits int) *_termSuggester {

	s.v.MaxEdits = &maxedits

	return s
}

// A factor that is used to multiply with the shard_size in order to inspect
// more candidate spelling corrections on the shard level.
// Can improve accuracy at the cost of performance.
func (s *_termSuggester) MaxInspections(maxinspections int) *_termSuggester {

	s.v.MaxInspections = &maxinspections

	return s
}

// The maximum threshold in number of documents in which a suggest text token
// can exist in order to be included.
// Can be a relative percentage number (for example `0.4`) or an absolute number
// to represent document frequencies.
// If a value higher than 1 is specified, then fractional can not be specified.
func (s *_termSuggester) MaxTermFreq(maxtermfreq float32) *_termSuggester {

	s.v.MaxTermFreq = &maxtermfreq

	return s
}

// The minimal threshold in number of documents a suggestion should appear in.
// This can improve quality by only suggesting high frequency terms.
// Can be specified as an absolute number or as a relative percentage of number
// of documents.
// If a value higher than 1 is specified, then the number cannot be fractional.
func (s *_termSuggester) MinDocFreq(mindocfreq float32) *_termSuggester {

	s.v.MinDocFreq = &mindocfreq

	return s
}

// The minimum length a suggest text term must have in order to be included.
func (s *_termSuggester) MinWordLength(minwordlength int) *_termSuggester {

	s.v.MinWordLength = &minwordlength

	return s
}

// The number of minimal prefix characters that must match in order be a
// candidate for suggestions.
// Increasing this number improves spellcheck performance.
func (s *_termSuggester) PrefixLength(prefixlength int) *_termSuggester {

	s.v.PrefixLength = &prefixlength

	return s
}

// Sets the maximum number of suggestions to be retrieved from each individual
// shard.
func (s *_termSuggester) ShardSize(shardsize int) *_termSuggester {

	s.v.ShardSize = &shardsize

	return s
}

// The maximum corrections to be returned per suggest text token.
func (s *_termSuggester) Size(size int) *_termSuggester {

	s.v.Size = &size

	return s
}

// Defines how suggestions should be sorted per suggest text term.
func (s *_termSuggester) Sort(sort suggestsort.SuggestSort) *_termSuggester {

	s.v.Sort = &sort
	return s
}

// The string distance implementation to use for comparing how similar suggested
// terms are.
func (s *_termSuggester) StringDistance(stringdistance stringdistance.StringDistance) *_termSuggester {

	s.v.StringDistance = &stringdistance
	return s
}

// Controls what suggestions are included or controls for what suggest text
// terms, suggestions should be suggested.
func (s *_termSuggester) SuggestMode(suggestmode suggestmode.SuggestMode) *_termSuggester {

	s.v.SuggestMode = &suggestmode
	return s
}

// The suggest text.
// Needs to be set globally or per suggestion.
func (s *_termSuggester) Text(text string) *_termSuggester {

	s.v.Text = &text

	return s
}

func (s *_termSuggester) FieldSuggesterCaster() *types.FieldSuggester {
	container := types.NewFieldSuggester()

	container.Term = s.v

	return container
}

func (s *_termSuggester) TermSuggesterCaster() *types.TermSuggester {
	return s.v
}
