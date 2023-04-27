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
// https://github.com/elastic/elasticsearch-specification/tree/a4f7b5a7f95dad95712a6bbce449241cbb84698d

package types

// CharFilterTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/a4f7b5a7f95dad95712a6bbce449241cbb84698d/specification/cluster/stats/types.ts#L133-L142
type CharFilterTypes struct {
	AnalyzerTypes      []FieldTypes `json:"analyzer_types"`
	BuiltInAnalyzers   []FieldTypes `json:"built_in_analyzers"`
	BuiltInCharFilters []FieldTypes `json:"built_in_char_filters"`
	BuiltInFilters     []FieldTypes `json:"built_in_filters"`
	BuiltInTokenizers  []FieldTypes `json:"built_in_tokenizers"`
	CharFilterTypes    []FieldTypes `json:"char_filter_types"`
	FilterTypes        []FieldTypes `json:"filter_types"`
	TokenizerTypes     []FieldTypes `json:"tokenizer_types"`
}

// NewCharFilterTypes returns a CharFilterTypes.
func NewCharFilterTypes() *CharFilterTypes {
	r := &CharFilterTypes{}

	return r
}
