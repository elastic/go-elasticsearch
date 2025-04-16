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
// https://github.com/elastic/elasticsearch-specification/tree/cbfcc73d01310bed2a480ec35aaef98138b598e5

package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _phraseSuggester struct {
	v *types.PhraseSuggester
}

// Provides access to word alternatives on a per token basis within a certain
// string distance.
func NewPhraseSuggester() *_phraseSuggester {

	return &_phraseSuggester{v: types.NewPhraseSuggester()}

}

func (s *_phraseSuggester) Analyzer(analyzer string) *_phraseSuggester {

	s.v.Analyzer = &analyzer

	return s
}

func (s *_phraseSuggester) Collate(collate types.PhraseSuggestCollateVariant) *_phraseSuggester {

	s.v.Collate = collate.PhraseSuggestCollateCaster()

	return s
}

func (s *_phraseSuggester) Confidence(confidence types.Float64) *_phraseSuggester {

	s.v.Confidence = &confidence

	return s
}

func (s *_phraseSuggester) DirectGenerator(directgenerators ...types.DirectGeneratorVariant) *_phraseSuggester {

	for _, v := range directgenerators {

		s.v.DirectGenerator = append(s.v.DirectGenerator, *v.DirectGeneratorCaster())

	}
	return s
}

func (s *_phraseSuggester) Field(field string) *_phraseSuggester {

	s.v.Field = field

	return s
}

func (s *_phraseSuggester) ForceUnigrams(forceunigrams bool) *_phraseSuggester {

	s.v.ForceUnigrams = &forceunigrams

	return s
}

func (s *_phraseSuggester) GramSize(gramsize int) *_phraseSuggester {

	s.v.GramSize = &gramsize

	return s
}

func (s *_phraseSuggester) Highlight(highlight types.PhraseSuggestHighlightVariant) *_phraseSuggester {

	s.v.Highlight = highlight.PhraseSuggestHighlightCaster()

	return s
}

func (s *_phraseSuggester) MaxErrors(maxerrors types.Float64) *_phraseSuggester {

	s.v.MaxErrors = &maxerrors

	return s
}

func (s *_phraseSuggester) RealWordErrorLikelihood(realworderrorlikelihood types.Float64) *_phraseSuggester {

	s.v.RealWordErrorLikelihood = &realworderrorlikelihood

	return s
}

func (s *_phraseSuggester) Separator(separator string) *_phraseSuggester {

	s.v.Separator = &separator

	return s
}

func (s *_phraseSuggester) ShardSize(shardsize int) *_phraseSuggester {

	s.v.ShardSize = &shardsize

	return s
}

func (s *_phraseSuggester) Size(size int) *_phraseSuggester {

	s.v.Size = &size

	return s
}

func (s *_phraseSuggester) Smoothing(smoothing types.SmoothingModelContainerVariant) *_phraseSuggester {

	s.v.Smoothing = smoothing.SmoothingModelContainerCaster()

	return s
}

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
