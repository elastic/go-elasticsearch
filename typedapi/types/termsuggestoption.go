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
// https://github.com/elastic/elasticsearch-specification/tree/9b556a1c9fd30159115d6c15226d0cac53a1d1a7


package types

// TermSuggestOption type.
//
// https://github.com/elastic/elasticsearch-specification/blob/9b556a1c9fd30159115d6c15226d0cac53a1d1a7/specification/_global/search/_types/suggester.ts#L93-L99
type TermSuggestOption struct {
	CollateMatch *bool   `json:"collate_match,omitempty"`
	Freq         int64   `json:"freq"`
	Highlighted  *string `json:"highlighted,omitempty"`
	Score        float64 `json:"score"`
	Text         string  `json:"text"`
}

// TermSuggestOptionBuilder holds TermSuggestOption struct and provides a builder API.
type TermSuggestOptionBuilder struct {
	v *TermSuggestOption
}

// NewTermSuggestOption provides a builder for the TermSuggestOption struct.
func NewTermSuggestOptionBuilder() *TermSuggestOptionBuilder {
	r := TermSuggestOptionBuilder{
		&TermSuggestOption{},
	}

	return &r
}

// Build finalize the chain and returns the TermSuggestOption struct
func (rb *TermSuggestOptionBuilder) Build() TermSuggestOption {
	return *rb.v
}

func (rb *TermSuggestOptionBuilder) CollateMatch(collatematch bool) *TermSuggestOptionBuilder {
	rb.v.CollateMatch = &collatematch
	return rb
}

func (rb *TermSuggestOptionBuilder) Freq(freq int64) *TermSuggestOptionBuilder {
	rb.v.Freq = freq
	return rb
}

func (rb *TermSuggestOptionBuilder) Highlighted(highlighted string) *TermSuggestOptionBuilder {
	rb.v.Highlighted = &highlighted
	return rb
}

func (rb *TermSuggestOptionBuilder) Score(score float64) *TermSuggestOptionBuilder {
	rb.v.Score = score
	return rb
}

func (rb *TermSuggestOptionBuilder) Text(text string) *TermSuggestOptionBuilder {
	rb.v.Text = text
	return rb
}
