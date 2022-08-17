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

// CharFilterTypes type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4316fc1aa18bb04678b156f23b22c9d3f996f9c9/specification/cluster/stats/types.ts#L126-L135
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

// CharFilterTypesBuilder holds CharFilterTypes struct and provides a builder API.
type CharFilterTypesBuilder struct {
	v *CharFilterTypes
}

// NewCharFilterTypes provides a builder for the CharFilterTypes struct.
func NewCharFilterTypesBuilder() *CharFilterTypesBuilder {
	r := CharFilterTypesBuilder{
		&CharFilterTypes{},
	}

	return &r
}

// Build finalize the chain and returns the CharFilterTypes struct
func (rb *CharFilterTypesBuilder) Build() CharFilterTypes {
	return *rb.v
}

func (rb *CharFilterTypesBuilder) AnalyzerTypes(analyzer_types []FieldTypesBuilder) *CharFilterTypesBuilder {
	tmp := make([]FieldTypes, len(analyzer_types))
	for _, value := range analyzer_types {
		tmp = append(tmp, value.Build())
	}
	rb.v.AnalyzerTypes = tmp
	return rb
}

func (rb *CharFilterTypesBuilder) BuiltInAnalyzers(built_in_analyzers []FieldTypesBuilder) *CharFilterTypesBuilder {
	tmp := make([]FieldTypes, len(built_in_analyzers))
	for _, value := range built_in_analyzers {
		tmp = append(tmp, value.Build())
	}
	rb.v.BuiltInAnalyzers = tmp
	return rb
}

func (rb *CharFilterTypesBuilder) BuiltInCharFilters(built_in_char_filters []FieldTypesBuilder) *CharFilterTypesBuilder {
	tmp := make([]FieldTypes, len(built_in_char_filters))
	for _, value := range built_in_char_filters {
		tmp = append(tmp, value.Build())
	}
	rb.v.BuiltInCharFilters = tmp
	return rb
}

func (rb *CharFilterTypesBuilder) BuiltInFilters(built_in_filters []FieldTypesBuilder) *CharFilterTypesBuilder {
	tmp := make([]FieldTypes, len(built_in_filters))
	for _, value := range built_in_filters {
		tmp = append(tmp, value.Build())
	}
	rb.v.BuiltInFilters = tmp
	return rb
}

func (rb *CharFilterTypesBuilder) BuiltInTokenizers(built_in_tokenizers []FieldTypesBuilder) *CharFilterTypesBuilder {
	tmp := make([]FieldTypes, len(built_in_tokenizers))
	for _, value := range built_in_tokenizers {
		tmp = append(tmp, value.Build())
	}
	rb.v.BuiltInTokenizers = tmp
	return rb
}

func (rb *CharFilterTypesBuilder) CharFilterTypes(char_filter_types []FieldTypesBuilder) *CharFilterTypesBuilder {
	tmp := make([]FieldTypes, len(char_filter_types))
	for _, value := range char_filter_types {
		tmp = append(tmp, value.Build())
	}
	rb.v.CharFilterTypes = tmp
	return rb
}

func (rb *CharFilterTypesBuilder) FilterTypes(filter_types []FieldTypesBuilder) *CharFilterTypesBuilder {
	tmp := make([]FieldTypes, len(filter_types))
	for _, value := range filter_types {
		tmp = append(tmp, value.Build())
	}
	rb.v.FilterTypes = tmp
	return rb
}

func (rb *CharFilterTypesBuilder) TokenizerTypes(tokenizer_types []FieldTypesBuilder) *CharFilterTypesBuilder {
	tmp := make([]FieldTypes, len(tokenizer_types))
	for _, value := range tokenizer_types {
		tmp = append(tmp, value.Build())
	}
	rb.v.TokenizerTypes = tmp
	return rb
}
