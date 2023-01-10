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

// CategorizationAnalyzerDefinition type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/Analysis.ts#L127-L140
type CategorizationAnalyzerDefinition struct {
	// CharFilter One or more character filters. In addition to the built-in character filters,
	// other plugins can provide more character filters. If this property is not
	// specified, no character filters are applied prior to categorization. If you
	// are customizing some other aspect of the analyzer and you need to achieve the
	// equivalent of `categorization_filters` (which are not permitted when some
	// other aspect of the analyzer is customized), add them here as pattern replace
	// character filters.
	CharFilter []CharFilter `json:"char_filter,omitempty"`
	// Filter One or more token filters. In addition to the built-in token filters, other
	// plugins can provide more token filters. If this property is not specified, no
	// token filters are applied prior to categorization.
	Filter []TokenFilter `json:"filter,omitempty"`
	// Tokenizer The name or definition of the tokenizer to use after character filters are
	// applied. This property is compulsory if `categorization_analyzer` is
	// specified as an object. Machine learning provides a tokenizer called
	// `ml_standard` that tokenizes in a way that has been determined to produce
	// good categorization results on a variety of log file formats for logs in
	// English. If you want to use that tokenizer but change the character or token
	// filters, specify "tokenizer": "ml_standard" in your
	// `categorization_analyzer`. Additionally, the `ml_classic` tokenizer is
	// available, which tokenizes in the same way as the non-customizable tokenizer
	// in old versions of the product (before 6.2). `ml_classic` was the default
	// categorization tokenizer in versions 6.2 to 7.13, so if you need
	// categorization identical to the default for jobs created in these versions,
	// specify "tokenizer": "ml_classic" in your `categorization_analyzer`.
	Tokenizer *Tokenizer `json:"tokenizer,omitempty"`
}

// NewCategorizationAnalyzerDefinition returns a CategorizationAnalyzerDefinition.
func NewCategorizationAnalyzerDefinition() *CategorizationAnalyzerDefinition {
	r := &CategorizationAnalyzerDefinition{}

	return r
}
