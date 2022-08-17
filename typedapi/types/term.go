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

// Term type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_global/termvectors/types.ts#L34-L40
type Term struct {
	DocFreq  *int     `json:"doc_freq,omitempty"`
	Score    *float64 `json:"score,omitempty"`
	TermFreq int      `json:"term_freq"`
	Tokens   []Token  `json:"tokens,omitempty"`
	Ttf      *int     `json:"ttf,omitempty"`
}

// TermBuilder holds Term struct and provides a builder API.
type TermBuilder struct {
	v *Term
}

// NewTerm provides a builder for the Term struct.
func NewTermBuilder() *TermBuilder {
	r := TermBuilder{
		&Term{},
	}

	return &r
}

// Build finalize the chain and returns the Term struct
func (rb *TermBuilder) Build() Term {
	return *rb.v
}

func (rb *TermBuilder) DocFreq(docfreq int) *TermBuilder {
	rb.v.DocFreq = &docfreq
	return rb
}

func (rb *TermBuilder) Score(score float64) *TermBuilder {
	rb.v.Score = &score
	return rb
}

func (rb *TermBuilder) TermFreq(termfreq int) *TermBuilder {
	rb.v.TermFreq = termfreq
	return rb
}

func (rb *TermBuilder) Tokens(tokens []TokenBuilder) *TermBuilder {
	tmp := make([]Token, len(tokens))
	for _, value := range tokens {
		tmp = append(tmp, value.Build())
	}
	rb.v.Tokens = tmp
	return rb
}

func (rb *TermBuilder) Ttf(ttf int) *TermBuilder {
	rb.v.Ttf = &ttf
	return rb
}
