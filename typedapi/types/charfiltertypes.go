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
// https://github.com/elastic/elasticsearch-specification/tree/5fea44e006349579bf3561a82e997002e5716117

package types

// CharFilterTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/5fea44e006349579bf3561a82e997002e5716117/specification/cluster/stats/types.ts#L228-L261
type CharFilterTypes struct {
	// AnalyzerTypes Contains statistics about analyzer types used in selected nodes.
	AnalyzerTypes []FieldTypes `json:"analyzer_types"`
	// BuiltInAnalyzers Contains statistics about built-in analyzers used in selected nodes.
	BuiltInAnalyzers []FieldTypes `json:"built_in_analyzers"`
	// BuiltInCharFilters Contains statistics about built-in character filters used in selected nodes.
	BuiltInCharFilters []FieldTypes `json:"built_in_char_filters"`
	// BuiltInFilters Contains statistics about built-in token filters used in selected nodes.
	BuiltInFilters []FieldTypes `json:"built_in_filters"`
	// BuiltInTokenizers Contains statistics about built-in tokenizers used in selected nodes.
	BuiltInTokenizers []FieldTypes `json:"built_in_tokenizers"`
	// CharFilterTypes Contains statistics about character filter types used in selected nodes.
	CharFilterTypes []FieldTypes `json:"char_filter_types"`
	// FilterTypes Contains statistics about token filter types used in selected nodes.
	FilterTypes []FieldTypes `json:"filter_types"`
	// TokenizerTypes Contains statistics about tokenizer types used in selected nodes.
	TokenizerTypes []FieldTypes `json:"tokenizer_types"`
}

// NewCharFilterTypes returns a CharFilterTypes.
func NewCharFilterTypes() *CharFilterTypes {
	r := &CharFilterTypes{}

	return r
}
