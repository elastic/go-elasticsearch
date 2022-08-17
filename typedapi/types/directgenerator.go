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

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/suggestmode"
)

// DirectGenerator type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L163-L175
type DirectGenerator struct {
	Field          Field                    `json:"field"`
	MaxEdits       *int                     `json:"max_edits,omitempty"`
	MaxInspections *float32                 `json:"max_inspections,omitempty"`
	MaxTermFreq    *float32                 `json:"max_term_freq,omitempty"`
	MinDocFreq     *float32                 `json:"min_doc_freq,omitempty"`
	MinWordLength  *int                     `json:"min_word_length,omitempty"`
	PostFilter     *string                  `json:"post_filter,omitempty"`
	PreFilter      *string                  `json:"pre_filter,omitempty"`
	PrefixLength   *int                     `json:"prefix_length,omitempty"`
	Size           *int                     `json:"size,omitempty"`
	SuggestMode    *suggestmode.SuggestMode `json:"suggest_mode,omitempty"`
}

// DirectGeneratorBuilder holds DirectGenerator struct and provides a builder API.
type DirectGeneratorBuilder struct {
	v *DirectGenerator
}

// NewDirectGenerator provides a builder for the DirectGenerator struct.
func NewDirectGeneratorBuilder() *DirectGeneratorBuilder {
	r := DirectGeneratorBuilder{
		&DirectGenerator{},
	}

	return &r
}

// Build finalize the chain and returns the DirectGenerator struct
func (rb *DirectGeneratorBuilder) Build() DirectGenerator {
	return *rb.v
}

func (rb *DirectGeneratorBuilder) Field(field Field) *DirectGeneratorBuilder {
	rb.v.Field = field
	return rb
}

func (rb *DirectGeneratorBuilder) MaxEdits(maxedits int) *DirectGeneratorBuilder {
	rb.v.MaxEdits = &maxedits
	return rb
}

func (rb *DirectGeneratorBuilder) MaxInspections(maxinspections float32) *DirectGeneratorBuilder {
	rb.v.MaxInspections = &maxinspections
	return rb
}

func (rb *DirectGeneratorBuilder) MaxTermFreq(maxtermfreq float32) *DirectGeneratorBuilder {
	rb.v.MaxTermFreq = &maxtermfreq
	return rb
}

func (rb *DirectGeneratorBuilder) MinDocFreq(mindocfreq float32) *DirectGeneratorBuilder {
	rb.v.MinDocFreq = &mindocfreq
	return rb
}

func (rb *DirectGeneratorBuilder) MinWordLength(minwordlength int) *DirectGeneratorBuilder {
	rb.v.MinWordLength = &minwordlength
	return rb
}

func (rb *DirectGeneratorBuilder) PostFilter(postfilter string) *DirectGeneratorBuilder {
	rb.v.PostFilter = &postfilter
	return rb
}

func (rb *DirectGeneratorBuilder) PreFilter(prefilter string) *DirectGeneratorBuilder {
	rb.v.PreFilter = &prefilter
	return rb
}

func (rb *DirectGeneratorBuilder) PrefixLength(prefixlength int) *DirectGeneratorBuilder {
	rb.v.PrefixLength = &prefixlength
	return rb
}

func (rb *DirectGeneratorBuilder) Size(size int) *DirectGeneratorBuilder {
	rb.v.Size = &size
	return rb
}

func (rb *DirectGeneratorBuilder) SuggestMode(suggestmode suggestmode.SuggestMode) *DirectGeneratorBuilder {
	rb.v.SuggestMode = &suggestmode
	return rb
}
