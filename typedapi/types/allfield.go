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

// AllField type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/mapping/meta-fields.ts#L29-L40
type AllField struct {
	Analyzer                 string `json:"analyzer"`
	Enabled                  bool   `json:"enabled"`
	OmitNorms                bool   `json:"omit_norms"`
	SearchAnalyzer           string `json:"search_analyzer"`
	Similarity               string `json:"similarity"`
	Store                    bool   `json:"store"`
	StoreTermVectorOffsets   bool   `json:"store_term_vector_offsets"`
	StoreTermVectorPayloads  bool   `json:"store_term_vector_payloads"`
	StoreTermVectorPositions bool   `json:"store_term_vector_positions"`
	StoreTermVectors         bool   `json:"store_term_vectors"`
}

// AllFieldBuilder holds AllField struct and provides a builder API.
type AllFieldBuilder struct {
	v *AllField
}

// NewAllField provides a builder for the AllField struct.
func NewAllFieldBuilder() *AllFieldBuilder {
	r := AllFieldBuilder{
		&AllField{},
	}

	return &r
}

// Build finalize the chain and returns the AllField struct
func (rb *AllFieldBuilder) Build() AllField {
	return *rb.v
}

func (rb *AllFieldBuilder) Analyzer(analyzer string) *AllFieldBuilder {
	rb.v.Analyzer = analyzer
	return rb
}

func (rb *AllFieldBuilder) Enabled(enabled bool) *AllFieldBuilder {
	rb.v.Enabled = enabled
	return rb
}

func (rb *AllFieldBuilder) OmitNorms(omitnorms bool) *AllFieldBuilder {
	rb.v.OmitNorms = omitnorms
	return rb
}

func (rb *AllFieldBuilder) SearchAnalyzer(searchanalyzer string) *AllFieldBuilder {
	rb.v.SearchAnalyzer = searchanalyzer
	return rb
}

func (rb *AllFieldBuilder) Similarity(similarity string) *AllFieldBuilder {
	rb.v.Similarity = similarity
	return rb
}

func (rb *AllFieldBuilder) Store(store bool) *AllFieldBuilder {
	rb.v.Store = store
	return rb
}

func (rb *AllFieldBuilder) StoreTermVectorOffsets(storetermvectoroffsets bool) *AllFieldBuilder {
	rb.v.StoreTermVectorOffsets = storetermvectoroffsets
	return rb
}

func (rb *AllFieldBuilder) StoreTermVectorPayloads(storetermvectorpayloads bool) *AllFieldBuilder {
	rb.v.StoreTermVectorPayloads = storetermvectorpayloads
	return rb
}

func (rb *AllFieldBuilder) StoreTermVectorPositions(storetermvectorpositions bool) *AllFieldBuilder {
	rb.v.StoreTermVectorPositions = storetermvectorpositions
	return rb
}

func (rb *AllFieldBuilder) StoreTermVectors(storetermvectors bool) *AllFieldBuilder {
	rb.v.StoreTermVectors = storetermvectors
	return rb
}
