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

// DataframeAnalysisFeatureProcessor type.
//
// https://github.com/elastic/elasticsearch-specification/blob/7f49eec1f23a5ae155001c058b3196d85981d5c2/specification/ml/_types/DataframeAnalytics.ts#L246-L258
type DataframeAnalysisFeatureProcessor struct {
	// FrequencyEncoding The configuration information necessary to perform frequency encoding.
	FrequencyEncoding *DataframeAnalysisFeatureProcessorFrequencyEncoding `json:"frequency_encoding,omitempty"`
	// MultiEncoding The configuration information necessary to perform multi encoding. It allows
	// multiple processors to be changed together. This way the output of a
	// processor can then be passed to another as an input.
	MultiEncoding *DataframeAnalysisFeatureProcessorMultiEncoding `json:"multi_encoding,omitempty"`
	// NGramEncoding The configuration information necessary to perform n-gram encoding. Features
	// created by this encoder have the following name format:
	// <feature_prefix>.<ngram><string position>. For example, if the feature_prefix
	// is f, the feature name for the second unigram in a string is f.11.
	NGramEncoding *DataframeAnalysisFeatureProcessorNGramEncoding `json:"n_gram_encoding,omitempty"`
	// OneHotEncoding The configuration information necessary to perform one hot encoding.
	OneHotEncoding *DataframeAnalysisFeatureProcessorOneHotEncoding `json:"one_hot_encoding,omitempty"`
	// TargetMeanEncoding The configuration information necessary to perform target mean encoding.
	TargetMeanEncoding *DataframeAnalysisFeatureProcessorTargetMeanEncoding `json:"target_mean_encoding,omitempty"`
}

// NewDataframeAnalysisFeatureProcessor returns a DataframeAnalysisFeatureProcessor.
func NewDataframeAnalysisFeatureProcessor() *DataframeAnalysisFeatureProcessor {
	r := &DataframeAnalysisFeatureProcessor{}

	return r
}
