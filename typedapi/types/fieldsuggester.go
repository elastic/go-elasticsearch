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

// FieldSuggester type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L103-L117
type FieldSuggester struct {
	Completion *CompletionSuggester `json:"completion,omitempty"`
	Phrase     *PhraseSuggester     `json:"phrase,omitempty"`
	Prefix     *string              `json:"prefix,omitempty"`
	Regex      *string              `json:"regex,omitempty"`
	Term       *TermSuggester       `json:"term,omitempty"`
	Text       *string              `json:"text,omitempty"`
}

// FieldSuggesterBuilder holds FieldSuggester struct and provides a builder API.
type FieldSuggesterBuilder struct {
	v *FieldSuggester
}

// NewFieldSuggester provides a builder for the FieldSuggester struct.
func NewFieldSuggesterBuilder() *FieldSuggesterBuilder {
	r := FieldSuggesterBuilder{
		&FieldSuggester{},
	}

	return &r
}

// Build finalize the chain and returns the FieldSuggester struct
func (rb *FieldSuggesterBuilder) Build() FieldSuggester {
	return *rb.v
}

func (rb *FieldSuggesterBuilder) Completion(completion *CompletionSuggesterBuilder) *FieldSuggesterBuilder {
	v := completion.Build()
	rb.v.Completion = &v
	return rb
}

func (rb *FieldSuggesterBuilder) Phrase(phrase *PhraseSuggesterBuilder) *FieldSuggesterBuilder {
	v := phrase.Build()
	rb.v.Phrase = &v
	return rb
}

func (rb *FieldSuggesterBuilder) Prefix(prefix string) *FieldSuggesterBuilder {
	rb.v.Prefix = &prefix
	return rb
}

func (rb *FieldSuggesterBuilder) Regex(regex string) *FieldSuggesterBuilder {
	rb.v.Regex = &regex
	return rb
}

func (rb *FieldSuggesterBuilder) Term(term *TermSuggesterBuilder) *FieldSuggesterBuilder {
	v := term.Build()
	rb.v.Term = &v
	return rb
}

func (rb *FieldSuggesterBuilder) Text(text string) *FieldSuggesterBuilder {
	rb.v.Text = &text
	return rb
}
