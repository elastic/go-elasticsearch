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

// TermSuggestOption type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L92-L96
type TermSuggestOption struct {
	Freq  int64   `json:"freq"`
	Score float64 `json:"score"`
	Text  string  `json:"text"`
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

func (rb *TermSuggestOptionBuilder) Freq(freq int64) *TermSuggestOptionBuilder {
	rb.v.Freq = freq
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
