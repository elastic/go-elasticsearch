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

// DataframeAnalysisFeatureProcessorNGramEncoding type.
//
// https://github.com/elastic/elasticsearch-specification/blob/4ab557491062aab5a916a1e274e28c266b0e0708/specification/ml/_types/DataframeAnalytics.ts#L274-L286
type DataframeAnalysisFeatureProcessorNGramEncoding struct {
	Custom *bool `json:"custom,omitempty"`
	// FeaturePrefix The feature name prefix. Defaults to ngram_<start>_<length>.
	FeaturePrefix *string `json:"feature_prefix,omitempty"`
	// Field The name of the text field to encode.
	Field string `json:"field"`
	// Length Specifies the length of the n-gram substring. Defaults to 50. Must be greater
	// than 0.
	Length *int `json:"length,omitempty"`
	// NGrams Specifies which n-grams to gather. It’s an array of integer values where the
	// minimum value is 1, and a maximum value is 5.
	NGrams []int `json:"n_grams"`
	// Start Specifies the zero-indexed start of the n-gram substring. Negative values are
	// allowed for encoding n-grams of string suffixes. Defaults to 0.
	Start *int `json:"start,omitempty"`
}

// NewDataframeAnalysisFeatureProcessorNGramEncoding returns a DataframeAnalysisFeatureProcessorNGramEncoding.
func NewDataframeAnalysisFeatureProcessorNGramEncoding() *DataframeAnalysisFeatureProcessorNGramEncoding {
	r := &DataframeAnalysisFeatureProcessorNGramEncoding{}

	return r
}
