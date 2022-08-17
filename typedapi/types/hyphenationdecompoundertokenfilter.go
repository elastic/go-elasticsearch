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

// HyphenationDecompounderTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L57-L59
type HyphenationDecompounderTokenFilter struct {
	HyphenationPatternsPath *string        `json:"hyphenation_patterns_path,omitempty"`
	MaxSubwordSize          *int           `json:"max_subword_size,omitempty"`
	MinSubwordSize          *int           `json:"min_subword_size,omitempty"`
	MinWordSize             *int           `json:"min_word_size,omitempty"`
	OnlyLongestMatch        *bool          `json:"only_longest_match,omitempty"`
	Type                    string         `json:"type,omitempty"`
	Version                 *VersionString `json:"version,omitempty"`
	WordList                []string       `json:"word_list,omitempty"`
	WordListPath            *string        `json:"word_list_path,omitempty"`
}

// HyphenationDecompounderTokenFilterBuilder holds HyphenationDecompounderTokenFilter struct and provides a builder API.
type HyphenationDecompounderTokenFilterBuilder struct {
	v *HyphenationDecompounderTokenFilter
}

// NewHyphenationDecompounderTokenFilter provides a builder for the HyphenationDecompounderTokenFilter struct.
func NewHyphenationDecompounderTokenFilterBuilder() *HyphenationDecompounderTokenFilterBuilder {
	r := HyphenationDecompounderTokenFilterBuilder{
		&HyphenationDecompounderTokenFilter{},
	}

	r.v.Type = "hyphenation_decompounder"

	return &r
}

// Build finalize the chain and returns the HyphenationDecompounderTokenFilter struct
func (rb *HyphenationDecompounderTokenFilterBuilder) Build() HyphenationDecompounderTokenFilter {
	return *rb.v
}

func (rb *HyphenationDecompounderTokenFilterBuilder) HyphenationPatternsPath(hyphenationpatternspath string) *HyphenationDecompounderTokenFilterBuilder {
	rb.v.HyphenationPatternsPath = &hyphenationpatternspath
	return rb
}

func (rb *HyphenationDecompounderTokenFilterBuilder) MaxSubwordSize(maxsubwordsize int) *HyphenationDecompounderTokenFilterBuilder {
	rb.v.MaxSubwordSize = &maxsubwordsize
	return rb
}

func (rb *HyphenationDecompounderTokenFilterBuilder) MinSubwordSize(minsubwordsize int) *HyphenationDecompounderTokenFilterBuilder {
	rb.v.MinSubwordSize = &minsubwordsize
	return rb
}

func (rb *HyphenationDecompounderTokenFilterBuilder) MinWordSize(minwordsize int) *HyphenationDecompounderTokenFilterBuilder {
	rb.v.MinWordSize = &minwordsize
	return rb
}

func (rb *HyphenationDecompounderTokenFilterBuilder) OnlyLongestMatch(onlylongestmatch bool) *HyphenationDecompounderTokenFilterBuilder {
	rb.v.OnlyLongestMatch = &onlylongestmatch
	return rb
}

func (rb *HyphenationDecompounderTokenFilterBuilder) Version(version VersionString) *HyphenationDecompounderTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}

func (rb *HyphenationDecompounderTokenFilterBuilder) WordList(word_list ...string) *HyphenationDecompounderTokenFilterBuilder {
	rb.v.WordList = word_list
	return rb
}

func (rb *HyphenationDecompounderTokenFilterBuilder) WordListPath(wordlistpath string) *HyphenationDecompounderTokenFilterBuilder {
	rb.v.WordListPath = &wordlistpath
	return rb
}
