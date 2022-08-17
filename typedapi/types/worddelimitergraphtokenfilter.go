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

// WordDelimiterGraphTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/_types/analysis/token_filters.ts#L148-L165
type WordDelimiterGraphTokenFilter struct {
	AdjustOffsets         *bool          `json:"adjust_offsets,omitempty"`
	CatenateAll           *bool          `json:"catenate_all,omitempty"`
	CatenateNumbers       *bool          `json:"catenate_numbers,omitempty"`
	CatenateWords         *bool          `json:"catenate_words,omitempty"`
	GenerateNumberParts   *bool          `json:"generate_number_parts,omitempty"`
	GenerateWordParts     *bool          `json:"generate_word_parts,omitempty"`
	IgnoreKeywords        *bool          `json:"ignore_keywords,omitempty"`
	PreserveOriginal      *bool          `json:"preserve_original,omitempty"`
	ProtectedWords        []string       `json:"protected_words,omitempty"`
	ProtectedWordsPath    *string        `json:"protected_words_path,omitempty"`
	SplitOnCaseChange     *bool          `json:"split_on_case_change,omitempty"`
	SplitOnNumerics       *bool          `json:"split_on_numerics,omitempty"`
	StemEnglishPossessive *bool          `json:"stem_english_possessive,omitempty"`
	Type                  string         `json:"type,omitempty"`
	TypeTable             []string       `json:"type_table,omitempty"`
	TypeTablePath         *string        `json:"type_table_path,omitempty"`
	Version               *VersionString `json:"version,omitempty"`
}

// WordDelimiterGraphTokenFilterBuilder holds WordDelimiterGraphTokenFilter struct and provides a builder API.
type WordDelimiterGraphTokenFilterBuilder struct {
	v *WordDelimiterGraphTokenFilter
}

// NewWordDelimiterGraphTokenFilter provides a builder for the WordDelimiterGraphTokenFilter struct.
func NewWordDelimiterGraphTokenFilterBuilder() *WordDelimiterGraphTokenFilterBuilder {
	r := WordDelimiterGraphTokenFilterBuilder{
		&WordDelimiterGraphTokenFilter{},
	}

	r.v.Type = "word_delimiter_graph"

	return &r
}

// Build finalize the chain and returns the WordDelimiterGraphTokenFilter struct
func (rb *WordDelimiterGraphTokenFilterBuilder) Build() WordDelimiterGraphTokenFilter {
	return *rb.v
}

func (rb *WordDelimiterGraphTokenFilterBuilder) AdjustOffsets(adjustoffsets bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.AdjustOffsets = &adjustoffsets
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) CatenateAll(catenateall bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.CatenateAll = &catenateall
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) CatenateNumbers(catenatenumbers bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.CatenateNumbers = &catenatenumbers
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) CatenateWords(catenatewords bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.CatenateWords = &catenatewords
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) GenerateNumberParts(generatenumberparts bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.GenerateNumberParts = &generatenumberparts
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) GenerateWordParts(generatewordparts bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.GenerateWordParts = &generatewordparts
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) IgnoreKeywords(ignorekeywords bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.IgnoreKeywords = &ignorekeywords
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) PreserveOriginal(preserveoriginal bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.PreserveOriginal = &preserveoriginal
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) ProtectedWords(protected_words ...string) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.ProtectedWords = protected_words
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) ProtectedWordsPath(protectedwordspath string) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.ProtectedWordsPath = &protectedwordspath
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) SplitOnCaseChange(splitoncasechange bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.SplitOnCaseChange = &splitoncasechange
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) SplitOnNumerics(splitonnumerics bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.SplitOnNumerics = &splitonnumerics
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) StemEnglishPossessive(stemenglishpossessive bool) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.StemEnglishPossessive = &stemenglishpossessive
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) TypeTable(type_table ...string) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.TypeTable = type_table
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) TypeTablePath(typetablepath string) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.TypeTablePath = &typetablepath
	return rb
}

func (rb *WordDelimiterGraphTokenFilterBuilder) Version(version VersionString) *WordDelimiterGraphTokenFilterBuilder {
	rb.v.Version = &version
	return rb
}
