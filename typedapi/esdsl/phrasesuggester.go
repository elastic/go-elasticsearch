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
// https://github.com/elastic/elasticsearch-specification/tree/c75a0abec670d027d13eb8d6f23374f86621c76b

package esdsl

import "github.com/elastic/go-elasticsearch/v8/typedapi/types"

type _phraseSuggester struct {
	v *types.PhraseSuggester
}

// Provides access to word alternatives on a per token basis within a certain
// string distance.
func NewPhraseSuggester() *_phraseSuggester {

	return &_phraseSuggester{v: types.NewPhraseSuggester()}

}

// The analyzer to analyze the suggest text with.
// Defaults to the search analyzer of the suggest field.
func (s *_phraseSuggester) Analyzer(analyzer string) *_phraseSuggester {

	s.v.Analyzer = &analyzer

	return s
}

// Checks each suggestion against the specified query to prune suggestions for
// which no matching docs exist in the index.
func (s *_phraseSuggester) Collate(collate types.PhraseSuggestCollateVariant) *_phraseSuggester {

	s.v.Collate = collate.PhraseSuggestCollateCaster()

	return s
}

// Defines a factor applied to the input phrases score, which is used as a
// threshold for other suggest candidates.
// Only candidates that score higher than the threshold will be included in the
// result.
func (s *_phraseSuggester) Confidence(confidence types.Float64) *_phraseSuggester {

	s.v.Confidence = &confidence

	return s
}

// A list of candidate generators that produce a list of possible terms per term
// in the given text.
func (s *_phraseSuggester) DirectGenerator(directgenerators ...types.DirectGeneratorVariant) *_phraseSuggester {

	for _, v := range directgenerators {

		s.v.DirectGenerator = append(s.v.DirectGenerator, *v.DirectGeneratorCaster())

	}
	return s
}

// The field to fetch the candidate suggestions from.
// Needs to be set globally or per suggestion.
func (s *_phraseSuggester) Field(field string) *_phraseSuggester {

	s.v.Field = field

	return s
}

func (s *_phraseSuggester) ForceUnigrams(forceunigrams bool) *_phraseSuggester {

	s.v.ForceUnigrams = &forceunigrams

	return s
}

// Sets max size of the n-grams (shingles) in the field.
// If the field doesn’t contain n-grams (shingles), this should be omitted or
// set to `1`.
// If the field uses a shingle filter, the `gram_size` is set to the
// `max_shingle_size` if not explicitly set.
func (s *_phraseSuggester) GramSize(gramsize int) *_phraseSuggester {

	s.v.GramSize = &gramsize

	return s
}

// Sets up suggestion highlighting.
// If not provided, no highlighted field is returned.
func (s *_phraseSuggester) Highlight(highlight types.PhraseSuggestHighlightVariant) *_phraseSuggester {

	s.v.Highlight = highlight.PhraseSuggestHighlightCaster()

	return s
}

// The maximum percentage of the terms considered to be misspellings in order to
// form a correction.
// This method accepts a float value in the range `[0..1)` as a fraction of the
// actual query terms or a number `>=1` as an absolute number of query terms.
func (s *_phraseSuggester) MaxErrors(maxerrors types.Float64) *_phraseSuggester {

	s.v.MaxErrors = &maxerrors

	return s
}

// The likelihood of a term being misspelled even if the term exists in the
// dictionary.
func (s *_phraseSuggester) RealWordErrorLikelihood(realworderrorlikelihood types.Float64) *_phraseSuggester {

	s.v.RealWordErrorLikelihood = &realworderrorlikelihood

	return s
}

// The separator that is used to separate terms in the bigram field.
// If not set, the whitespace character is used as a separator.
func (s *_phraseSuggester) Separator(separator string) *_phraseSuggester {

	s.v.Separator = &separator

	return s
}

// Sets the maximum number of suggested terms to be retrieved from each
// individual shard.
func (s *_phraseSuggester) ShardSize(shardsize int) *_phraseSuggester {

	s.v.ShardSize = &shardsize

	return s
}

// The maximum corrections to be returned per suggest text token.
func (s *_phraseSuggester) Size(size int) *_phraseSuggester {

	s.v.Size = &size

	return s
}

// The smoothing model used to balance weight between infrequent grams (grams
// (shingles) are not existing in the index) and frequent grams (appear at least
// once in the index).
// The default model is Stupid Backoff.
func (s *_phraseSuggester) Smoothing(smoothing types.SmoothingModelContainerVariant) *_phraseSuggester {

	s.v.Smoothing = smoothing.SmoothingModelContainerCaster()

	return s
}

// The text/query to provide suggestions for.
func (s *_phraseSuggester) Text(text string) *_phraseSuggester {

	s.v.Text = &text

	return s
}

func (s *_phraseSuggester) TokenLimit(tokenlimit int) *_phraseSuggester {

	s.v.TokenLimit = &tokenlimit

	return s
}

func (s *_phraseSuggester) FieldSuggesterCaster() *types.FieldSuggester {
	container := types.NewFieldSuggester()

	container.Phrase = s.v

	return container
}

func (s *_phraseSuggester) PhraseSuggesterCaster() *types.PhraseSuggester {
	return s.v
}
