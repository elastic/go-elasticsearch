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

// InvertedIndex type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/indices/field_usage_stats/IndicesFieldUsageStatsResponse.ts#L65-L73
type InvertedIndex struct {
	Offsets         uint `json:"offsets"`
	Payloads        uint `json:"payloads"`
	Positions       uint `json:"positions"`
	Postings        uint `json:"postings"`
	Proximity       uint `json:"proximity"`
	TermFrequencies uint `json:"term_frequencies"`
	Terms           uint `json:"terms"`
}

// InvertedIndexBuilder holds InvertedIndex struct and provides a builder API.
type InvertedIndexBuilder struct {
	v *InvertedIndex
}

// NewInvertedIndex provides a builder for the InvertedIndex struct.
func NewInvertedIndexBuilder() *InvertedIndexBuilder {
	r := InvertedIndexBuilder{
		&InvertedIndex{},
	}

	return &r
}

// Build finalize the chain and returns the InvertedIndex struct
func (rb *InvertedIndexBuilder) Build() InvertedIndex {
	return *rb.v
}

func (rb *InvertedIndexBuilder) Offsets(offsets uint) *InvertedIndexBuilder {
	rb.v.Offsets = offsets
	return rb
}

func (rb *InvertedIndexBuilder) Payloads(payloads uint) *InvertedIndexBuilder {
	rb.v.Payloads = payloads
	return rb
}

func (rb *InvertedIndexBuilder) Positions(positions uint) *InvertedIndexBuilder {
	rb.v.Positions = positions
	return rb
}

func (rb *InvertedIndexBuilder) Postings(postings uint) *InvertedIndexBuilder {
	rb.v.Postings = postings
	return rb
}

func (rb *InvertedIndexBuilder) Proximity(proximity uint) *InvertedIndexBuilder {
	rb.v.Proximity = proximity
	return rb
}

func (rb *InvertedIndexBuilder) TermFrequencies(termfrequencies uint) *InvertedIndexBuilder {
	rb.v.TermFrequencies = termfrequencies
	return rb
}

func (rb *InvertedIndexBuilder) Terms(terms uint) *InvertedIndexBuilder {
	rb.v.Terms = terms
	return rb
}
