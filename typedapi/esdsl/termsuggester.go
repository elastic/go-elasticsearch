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
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/stringdistance"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/suggestsort"
)

type _termSuggester struct {
	v *types.TermSuggester
}

// Suggests terms based on edit distance.
func NewTermSuggester() *_termSuggester {

	return &_termSuggester{v: types.NewTermSuggester()}

}

func (s *_termSuggester) Analyzer(analyzer string) *_termSuggester {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_termSuggester) Field(field string) *_termSuggester {

	s.v.Field = field

	return s
}

func (s *_termSuggester) LowercaseTerms(lowercaseterms bool) *_termSuggester {

	s.v.LowercaseTerms = &lowercaseterms

	return s
}

func (s *_termSuggester) MaxEdits(maxedits int) *_termSuggester {

	s.v.MaxEdits = &maxedits

	return s
}

func (s *_termSuggester) MaxInspections(maxinspections int) *_termSuggester {

	s.v.MaxInspections = &maxinspections

	return s
}

func (s *_termSuggester) MaxTermFreq(maxtermfreq float32) *_termSuggester {

	s.v.MaxTermFreq = &maxtermfreq

	return s
}

func (s *_termSuggester) MinDocFreq(mindocfreq float32) *_termSuggester {

	s.v.MinDocFreq = &mindocfreq

	return s
}

func (s *_termSuggester) MinWordLength(minwordlength int) *_termSuggester {

	s.v.MinWordLength = &minwordlength

	return s
}

func (s *_termSuggester) PrefixLength(prefixlength int) *_termSuggester {

	s.v.PrefixLength = &prefixlength

	return s
}

func (s *_termSuggester) ShardSize(shardsize int) *_termSuggester {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_termSuggester) Size(size int) *_termSuggester {

	s.v.Size = &size

	return s
}

func (s *_termSuggester) Sort(sort suggestsort.SuggestSort) *_termSuggester {

	s.v.Sort = &sort
	return s
}

func (s *_termSuggester) StringDistance(stringdistance stringdistance.StringDistance) *_termSuggester {

	s.v.StringDistance = &stringdistance
	return s
}

func (s *_termSuggester) SuggestMode(suggestmode suggestmode.SuggestMode) *_termSuggester {

	s.v.SuggestMode = &suggestmode
	return s
}

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
