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
// https://github.com/elastic/elasticsearch-specification/tree/4316fc1aa18bb04678b156f23b22c9d3f996f9c9


package types

// PhraseSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L188-L202
type PhraseSuggester struct {
	Analyzer                *string                  `json:"analyzer,omitempty"`
	Collate                 *PhraseSuggestCollate    `json:"collate,omitempty"`
	Confidence              *float64                 `json:"confidence,omitempty"`
	DirectGenerator         []DirectGenerator        `json:"direct_generator,omitempty"`
	Field                   Field                    `json:"field"`
	ForceUnigrams           *bool                    `json:"force_unigrams,omitempty"`
	GramSize                *int                     `json:"gram_size,omitempty"`
	Highlight               *PhraseSuggestHighlight  `json:"highlight,omitempty"`
	MaxErrors               *float64                 `json:"max_errors,omitempty"`
	RealWordErrorLikelihood *float64                 `json:"real_word_error_likelihood,omitempty"`
	Separator               *string                  `json:"separator,omitempty"`
	ShardSize               *int                     `json:"shard_size,omitempty"`
	Size                    *int                     `json:"size,omitempty"`
	Smoothing               *SmoothingModelContainer `json:"smoothing,omitempty"`
	Text                    *string                  `json:"text,omitempty"`
	TokenLimit              *int                     `json:"token_limit,omitempty"`
}

// PhraseSuggesterBuilder holds PhraseSuggester struct and provides a builder API.
type PhraseSuggesterBuilder struct {
	v *PhraseSuggester
}

// NewPhraseSuggester provides a builder for the PhraseSuggester struct.
func NewPhraseSuggesterBuilder() *PhraseSuggesterBuilder {
	r := PhraseSuggesterBuilder{
		&PhraseSuggester{},
	}

	return &r
}

// Build finalize the chain and returns the PhraseSuggester struct
func (rb *PhraseSuggesterBuilder) Build() PhraseSuggester {
	return *rb.v
}

func (rb *PhraseSuggesterBuilder) Analyzer(analyzer string) *PhraseSuggesterBuilder {
	rb.v.Analyzer = &analyzer
	return rb
}

func (rb *PhraseSuggesterBuilder) Collate(collate *PhraseSuggestCollateBuilder) *PhraseSuggesterBuilder {
	v := collate.Build()
	rb.v.Collate = &v
	return rb
}

func (rb *PhraseSuggesterBuilder) Confidence(confidence float64) *PhraseSuggesterBuilder {
	rb.v.Confidence = &confidence
	return rb
}

func (rb *PhraseSuggesterBuilder) DirectGenerator(direct_generator []DirectGeneratorBuilder) *PhraseSuggesterBuilder {
	tmp := make([]DirectGenerator, len(direct_generator))
	for _, value := range direct_generator {
		tmp = append(tmp, value.Build())
	}
	rb.v.DirectGenerator = tmp
	return rb
}

func (rb *PhraseSuggesterBuilder) Field(field Field) *PhraseSuggesterBuilder {
	rb.v.Field = field
	return rb
}

func (rb *PhraseSuggesterBuilder) ForceUnigrams(forceunigrams bool) *PhraseSuggesterBuilder {
	rb.v.ForceUnigrams = &forceunigrams
	return rb
}

func (rb *PhraseSuggesterBuilder) GramSize(gramsize int) *PhraseSuggesterBuilder {
	rb.v.GramSize = &gramsize
	return rb
}

func (rb *PhraseSuggesterBuilder) Highlight(highlight *PhraseSuggestHighlightBuilder) *PhraseSuggesterBuilder {
	v := highlight.Build()
	rb.v.Highlight = &v
	return rb
}

func (rb *PhraseSuggesterBuilder) MaxErrors(maxerrors float64) *PhraseSuggesterBuilder {
	rb.v.MaxErrors = &maxerrors
	return rb
}

func (rb *PhraseSuggesterBuilder) RealWordErrorLikelihood(realworderrorlikelihood float64) *PhraseSuggesterBuilder {
	rb.v.RealWordErrorLikelihood = &realworderrorlikelihood
	return rb
}

func (rb *PhraseSuggesterBuilder) Separator(separator string) *PhraseSuggesterBuilder {
	rb.v.Separator = &separator
	return rb
}

func (rb *PhraseSuggesterBuilder) ShardSize(shardsize int) *PhraseSuggesterBuilder {
	rb.v.ShardSize = &shardsize
	return rb
}

func (rb *PhraseSuggesterBuilder) Size(size int) *PhraseSuggesterBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *PhraseSuggesterBuilder) Smoothing(smoothing *SmoothingModelContainerBuilder) *PhraseSuggesterBuilder {
	v := smoothing.Build()
	rb.v.Smoothing = &v
	return rb
}

func (rb *PhraseSuggesterBuilder) Text(text string) *PhraseSuggesterBuilder {
	rb.v.Text = &text
	return rb
}

func (rb *PhraseSuggesterBuilder) TokenLimit(tokenlimit int) *PhraseSuggesterBuilder {
	rb.v.TokenLimit = &tokenlimit
	return rb
}
