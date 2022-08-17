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

// PhraseSuggestOption type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/search/_types/suggester.ts#L86-L90
type PhraseSuggestOption struct {
	Highlighted string  `json:"highlighted"`
	Score       float64 `json:"score"`
	Text        string  `json:"text"`
}

// PhraseSuggestOptionBuilder holds PhraseSuggestOption struct and provides a builder API.
type PhraseSuggestOptionBuilder struct {
	v *PhraseSuggestOption
}

// NewPhraseSuggestOption provides a builder for the PhraseSuggestOption struct.
func NewPhraseSuggestOptionBuilder() *PhraseSuggestOptionBuilder {
	r := PhraseSuggestOptionBuilder{
		&PhraseSuggestOption{},
	}

	return &r
}

// Build finalize the chain and returns the PhraseSuggestOption struct
func (rb *PhraseSuggestOptionBuilder) Build() PhraseSuggestOption {
	return *rb.v
}

func (rb *PhraseSuggestOptionBuilder) Highlighted(highlighted string) *PhraseSuggestOptionBuilder {
	rb.v.Highlighted = highlighted
	return rb
}

func (rb *PhraseSuggestOptionBuilder) Score(score float64) *PhraseSuggestOptionBuilder {
	rb.v.Score = score
	return rb
}

func (rb *PhraseSuggestOptionBuilder) Text(text string) *PhraseSuggestOptionBuilder {
	rb.v.Text = text
	return rb
}
