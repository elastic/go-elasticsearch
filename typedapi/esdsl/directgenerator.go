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
	"github.com/elastic/go-elasticsearch/v8/typedapi/types"
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
)

type _directGenerator struct {
	v *types.DirectGenerator
}

func NewDirectGenerator() *_directGenerator {

	return &_directGenerator{v: types.NewDirectGenerator()}

}

// The field to fetch the candidate suggestions from.
// Needs to be set globally or per suggestion.
func (s *_directGenerator) Field(field string) *_directGenerator {

	s.v.Field = field

	return s
}

// The maximum edit distance candidate suggestions can have in order to be
// considered as a suggestion.
// Can only be `1` or `2`.
func (s *_directGenerator) MaxEdits(maxedits int) *_directGenerator {

	s.v.MaxEdits = &maxedits

	return s
}

// A factor that is used to multiply with the shard_size in order to inspect
// more candidate spelling corrections on the shard level.
// Can improve accuracy at the cost of performance.
func (s *_directGenerator) MaxInspections(maxinspections float32) *_directGenerator {

	s.v.MaxInspections = &maxinspections

	return s
}

// The maximum threshold in number of documents in which a suggest text token
// can exist in order to be included.
// This can be used to exclude high frequency terms — which are usually spelled
// correctly — from being spellchecked.
// Can be a relative percentage number (for example `0.4`) or an absolute number
// to represent document frequencies.
// If a value higher than 1 is specified, then fractional can not be specified.
func (s *_directGenerator) MaxTermFreq(maxtermfreq float32) *_directGenerator {

	s.v.MaxTermFreq = &maxtermfreq

	return s
}

// The minimal threshold in number of documents a suggestion should appear in.
// This can improve quality by only suggesting high frequency terms.
// Can be specified as an absolute number or as a relative percentage of number
// of documents.
// If a value higher than 1 is specified, the number cannot be fractional.
func (s *_directGenerator) MinDocFreq(mindocfreq float32) *_directGenerator {

	s.v.MinDocFreq = &mindocfreq

	return s
}

// The minimum length a suggest text term must have in order to be included.
func (s *_directGenerator) MinWordLength(minwordlength int) *_directGenerator {

	s.v.MinWordLength = &minwordlength

	return s
}

// A filter (analyzer) that is applied to each of the generated tokens before
// they are passed to the actual phrase scorer.
func (s *_directGenerator) PostFilter(postfilter string) *_directGenerator {

	s.v.PostFilter = &postfilter

	return s
}

// A filter (analyzer) that is applied to each of the tokens passed to this
// candidate generator.
// This filter is applied to the original token before candidates are generated.
func (s *_directGenerator) PreFilter(prefilter string) *_directGenerator {

	s.v.PreFilter = &prefilter

	return s
}

// The number of minimal prefix characters that must match in order be a
// candidate suggestions.
// Increasing this number improves spellcheck performance.
func (s *_directGenerator) PrefixLength(prefixlength int) *_directGenerator {

	s.v.PrefixLength = &prefixlength

	return s
}

// The maximum corrections to be returned per suggest text token.
func (s *_directGenerator) Size(size int) *_directGenerator {

	s.v.Size = &size

	return s
}

// Controls what suggestions are included on the suggestions generated on each
// shard.
func (s *_directGenerator) SuggestMode(suggestmode suggestmode.SuggestMode) *_directGenerator {

	s.v.SuggestMode = &suggestmode
	return s
}

func (s *_directGenerator) DirectGeneratorCaster() *types.DirectGenerator {
	return s.v
}
