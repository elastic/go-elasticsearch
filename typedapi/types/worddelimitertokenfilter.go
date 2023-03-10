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
// https://github.com/elastic/elasticsearch-specification/tree/4ab557491062aab5a916a1e274e28c266b0e0708

package types

// WordDelimiterTokenFilter type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/_types/analysis/token_filters.ts#L131-L146
type WordDelimiterTokenFilter struct {
	CatenateAll           *bool    `json:"catenate_all,omitempty"`
	CatenateNumbers       *bool    `json:"catenate_numbers,omitempty"`
	CatenateWords         *bool    `json:"catenate_words,omitempty"`
	GenerateNumberParts   *bool    `json:"generate_number_parts,omitempty"`
	GenerateWordParts     *bool    `json:"generate_word_parts,omitempty"`
	PreserveOriginal      *bool    `json:"preserve_original,omitempty"`
	ProtectedWords        []string `json:"protected_words,omitempty"`
	ProtectedWordsPath    *string  `json:"protected_words_path,omitempty"`
	SplitOnCaseChange     *bool    `json:"split_on_case_change,omitempty"`
	SplitOnNumerics       *bool    `json:"split_on_numerics,omitempty"`
	StemEnglishPossessive *bool    `json:"stem_english_possessive,omitempty"`
	Type                  string   `json:"type,omitempty"`
	TypeTable             []string `json:"type_table,omitempty"`
	TypeTablePath         *string  `json:"type_table_path,omitempty"`
	Version               *string  `json:"version,omitempty"`
}

// NewWordDelimiterTokenFilter returns a WordDelimiterTokenFilter.
func NewWordDelimiterTokenFilter() *WordDelimiterTokenFilter {
	r := &WordDelimiterTokenFilter{}

	r.Type = "word_delimiter"

	return r
}
