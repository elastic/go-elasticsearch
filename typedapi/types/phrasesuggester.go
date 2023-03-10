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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// PhraseSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_global/search/_types/suggester.ts#L191-L205
type PhraseSuggester struct {
	Analyzer                *string                  `json:"analyzer,omitempty"`
	Collate                 *PhraseSuggestCollate    `json:"collate,omitempty"`
	Confidence              *Float64                 `json:"confidence,omitempty"`
	DirectGenerator         []DirectGenerator        `json:"direct_generator,omitempty"`
	Field                   string                   `json:"field"`
	ForceUnigrams           *bool                    `json:"force_unigrams,omitempty"`
	GramSize                *int                     `json:"gram_size,omitempty"`
	Highlight               *PhraseSuggestHighlight  `json:"highlight,omitempty"`
	MaxErrors               *Float64                 `json:"max_errors,omitempty"`
	RealWordErrorLikelihood *Float64                 `json:"real_word_error_likelihood,omitempty"`
	Separator               *string                  `json:"separator,omitempty"`
	ShardSize               *int                     `json:"shard_size,omitempty"`
	Size                    *int                     `json:"size,omitempty"`
	Smoothing               *SmoothingModelContainer `json:"smoothing,omitempty"`
	Text                    *string                  `json:"text,omitempty"`
	TokenLimit              *int                     `json:"token_limit,omitempty"`
}

// NewPhraseSuggester returns a PhraseSuggester.
func NewPhraseSuggester() *PhraseSuggester {
	r := &PhraseSuggester{}

	return r
}
