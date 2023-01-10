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
// https://github.com/elastic/elasticsearch-specification/tree/7f49eec1f23a5ae155001c058b3196d85981d5c2


package types

// TermVectorsFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/_global/termvectors/types.ts#L49-L57
type TermVectorsFilter struct {
	MaxDocFreq    *int `json:"max_doc_freq,omitempty"`
	MaxNumTerms   *int `json:"max_num_terms,omitempty"`
	MaxTermFreq   *int `json:"max_term_freq,omitempty"`
	MaxWordLength *int `json:"max_word_length,omitempty"`
	MinDocFreq    *int `json:"min_doc_freq,omitempty"`
	MinTermFreq   *int `json:"min_term_freq,omitempty"`
	MinWordLength *int `json:"min_word_length,omitempty"`
}

// NewTermVectorsFilter returns a TermVectorsFilter.
func NewTermVectorsFilter() *TermVectorsFilter {
	r := &TermVectorsFilter{}

	return r
}
